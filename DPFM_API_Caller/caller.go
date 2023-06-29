package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(msg rabbitmq.RabbitmqMessage) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	_, ok := input["ProductMasterGeneral"]
	if ok {
		input := &dpfm_api_input_reader.GeneralSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterGeneral(input)
		goto endProcess
	}
//	_, ok = input["SupplyChainRelationshipDeliveryRelation"]
//	if ok {
//		input := &dpfm_api_input_reader.DeliveryRelationSDC{}
//		err = json.Unmarshal(msg.Raw(), input)
//		ret = e.confSupplyChainRelationshipDeliveryRelation(input)
//		goto endProcess
//	}
	err = xerrors.Errorf("can not get exconf check target")
endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

func (e *ExistenceConf) confProductMasterGeneral(input *dpfm_api_input_reader.GeneralSDC) *dpfm_api_output_formatter.ProductMasterGeneral {
	exconf := dpfm_api_output_formatter.ProductMasterGeneral{
		ExistenceConf: false,
	}
	if input.ProductMasterGeneral.Product == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterGeneral{
		Product: *input.ProductMasterGeneral.Product,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data 
		WHERE (Product) = (?);`, exconf.Product,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
