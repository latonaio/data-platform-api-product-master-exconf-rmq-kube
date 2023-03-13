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
	_, ok = input["ProductMasterBusinessPartner"]
	if ok {
		input := &dpfm_api_input_reader.BusinessPartnerSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterBusinessPartner(input)
		goto endProcess
	}
	_, ok = input["ProductMasterBPPlant"]
	if ok {
		input := &dpfm_api_input_reader.BPPlantSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterBPPlant(input)
		goto endProcess
	}
	_, ok = input["ProductMasterStorageLocation"]
	if ok {
		input := &dpfm_api_input_reader.StorageLocationSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterStorageLocation(input)
		goto endProcess
	}

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
		Product:       *input.ProductMasterGeneral.Product,
		ExistenceConf: false,
	}

	rows, err := e.db.Query(
		`SELECT Product 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data 
		WHERE Product = ?;`, exconf.Product,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterBusinessPartner(input *dpfm_api_input_reader.BusinessPartnerSDC) *dpfm_api_output_formatter.ProductMasterBusinessPartner {
	exconf := dpfm_api_output_formatter.ProductMasterBusinessPartner{
		ExistenceConf: false,
	}
	if input.ProductMasterBusinessPartner.Product == nil {
		return &exconf
	}
	if input.ProductMasterBusinessPartner.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterBusinessPartner.ValidityEndDate == nil {
		return &exconf
	}
	if input.ProductMasterBusinessPartner.ValidityStartDate == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterBusinessPartner{
		Product:           *input.ProductMasterBusinessPartner.Product,
		BusinessPartner:   *input.ProductMasterBusinessPartner.BusinessPartner,
		ValidityEndDate:   *input.ProductMasterBusinessPartner.ValidityEndDate,
		ValidityStartDate: *input.ProductMasterBusinessPartner.ValidityStartDate,
		ExistenceConf:     false,
	}

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data
		WHERE (Product, BusinessPartner, ValidityEndDate, ValidityStartDate) = (?, ?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.ValidityEndDate, exconf.ValidityStartDate,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterBPPlant(input *dpfm_api_input_reader.BPPlantSDC) *dpfm_api_output_formatter.ProductMasterBPPlant {
	exconf := dpfm_api_output_formatter.ProductMasterBPPlant{
		ExistenceConf: false,
	}
	if input.ProductMasterBPPlant.Product == nil {
		return &exconf
	}
	if input.ProductMasterBPPlant.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterBPPlant.Plant == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterBPPlant{
		Product:         *input.ProductMasterBPPlant.Product,
		BusinessPartner: *input.ProductMasterBPPlant.BusinessPartner,
		Plant:           *input.ProductMasterBPPlant.Plant,
		ExistenceConf:   false,
	}

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) = (?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterStorageLocation(input *dpfm_api_input_reader.StorageLocationSDC) *dpfm_api_output_formatter.ProductMasterStorageLocation {
	exconf := dpfm_api_output_formatter.ProductMasterStorageLocation{
		ExistenceConf: false,
	}
	if input.ProductMasterStorageLocation.Product == nil {
		return &exconf
	}
	if input.ProductMasterStorageLocation.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterStorageLocation.Plant == nil {
		return &exconf
	}
	if input.ProductMasterStorageLocation.StorageLocation == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterStorageLocation{
		Product:         *input.ProductMasterStorageLocation.Product,
		BusinessPartner: *input.ProductMasterStorageLocation.BusinessPartner,
		Plant:           *input.ProductMasterStorageLocation.Plant,
		StorageLocation: *input.ProductMasterStorageLocation.StorageLocation,
		ExistenceConf:   false,
	}

	rows, err := e.db.Query(
		`SELECT ProductMaster
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_location_data
		WHERE (Product, BusinessPartner, Plant, StorageLocation) = (?, ?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant, exconf.StorageLocation,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
