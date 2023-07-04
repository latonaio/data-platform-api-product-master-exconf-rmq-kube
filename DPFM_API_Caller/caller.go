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
	_, ok = input["ProductMasterMRPArea"]
	if ok {
		input := &dpfm_api_input_reader.MRPAreaSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterMRPArea(input)
		goto endProcess
	}
	_, ok = input["ProductMasterStorageLocation"]
	if ok {
		input := &dpfm_api_input_reader.StorageLocationSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterStorageLocation(input)
		goto endProcess
	}
	_, ok = input["ProductMasterStorageBin"]
	if ok {
		input := &dpfm_api_input_reader.StorageBinSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterStorageBin(input)
		goto endProcess
	}
	_, ok = input["ProductMasterProduction"]
	if ok {
		input := &dpfm_api_input_reader.ProductionSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterProduction(input)
		goto endProcess
	}
	_, ok = input["ProductMasterQuality"]
	if ok {
		input := &dpfm_api_input_reader.QualitySDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterQuality(input)
		goto endProcess
	}
	_, ok = input["ProductMasterAccounting"]
	if ok {
		input := &dpfm_api_input_reader.AccountingSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterAccounting(input)
		goto endProcess
	}
	_, ok = input["ProductMasterTax"]
	if ok {
		input := &dpfm_api_input_reader.TaxSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterTax(input)
		goto endProcess
	}
	_, ok := input["ProductMasterProductDescription"]
	if ok {
		input := &dpfm_api_input_reader.ProductDescriptionSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterProductDescription(input)
		goto endProcess
	}
	_, ok := input["ProductMasterProductDescByBP"]
	if ok {
		input := &dpfm_api_input_reader.ProductDescByBPSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterProductDescByBP(input)
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
	exconf = dpfm_api_output_formatter.ProductMasterBusinessPartner{
		Product: 			*input.ProductMasterBusinessPartner.Product,
		BusinessPartner:	*input.ProductMasterBusinessPartner.BusinessPartner,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data 
		WHERE (Product, BusinessPartner) = (?);`, exconf.Product, exconf.BusinessPartner,
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
		Product: 			*input.ProductMasterBPPlant.Product,
		BusinessPartner:	*input.ProductMasterBPPlant.BusinessPartner,
		Plant:				*input.ProductMasterBPPlant.Plant,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data 
		WHERE (Product, BusinessPartner, Plant) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterProduction(input *dpfm_api_input_reader.ProductionSDC) *dpfm_api_output_formatter.ProductMasterProduction {
	exconf := dpfm_api_output_formatter.ProductMasterProduction{
		ExistenceConf: false,
	}
	if input.ProductMasterProduction.Product == nil {
		return &exconf
	}
	if input.ProductMasterProduction.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterProduction.Plant == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.ProductMasterProduction{
		Product: 			*input.ProductMasterProduction.Product,
		BusinessPartner:	*input.ProductMasterProduction.BusinessPartner,
		Plant:				*input.ProductMasterProduction.Plant,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_production_data 
		WHERE (Product, BusinessPartner, Plant) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant,
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
		Product: 			*input.ProductMasterStorageLocation.Product,
		BusinessPartner:	*input.ProductMasterStorageLocation.BusinessPartner,
		Plant:				*input.ProductMasterStorageLocation.Plant,
		StorageLocation:	*input.ProductMasterStorageLocation.StorageLocation,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_location_data 
		WHERE (Product, BusinessPartner, Plant, StorageLocation) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant, exconf.StorageLocation,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterQuality(input *dpfm_api_input_reader.QualitySDC) *dpfm_api_output_formatter.ProductMasterQuality {
	exconf := dpfm_api_output_formatter.ProductMasterQuality{
		ExistenceConf: false,
	}
	if input.ProductMasterQuality.Product == nil {
		return &exconf
	}
	if input.ProductMasterQuality.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterQuality.Plant == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.ProductMasterQuality{
		Product: 			*input.ProductMasterQuality.Product,
		BusinessPartner:	*input.ProductMasterQuality.BusinessPartner,
		Plant:				*input.ProductMasterQuality.Plant,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_quality_data 
		WHERE (Product, BusinessPartner, Plant) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}


func (e *ExistenceConf) confProductMasterAccounting(input *dpfm_api_input_reader.AccountingSDC) *dpfm_api_output_formatter.ProductMasterAccounting {
	exconf := dpfm_api_output_formatter.ProductMasterAccounting{
		ExistenceConf: false,
	}
	if input.ProductMasterAccounting.Product == nil {
		return &exconf
	}
	if input.ProductMasterAccounting.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterAccounting.Plant == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.ProductMasterAccounting{
		Product: 			*input.ProductMasterAccounting.Product,
		BusinessPartner:	*input.ProductMasterAccounting.BusinessPartner,
		Plant:				*input.ProductMasterAccounting.Plant,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_accounting_data 
		WHERE (Product, BusinessPartner, Plant) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterStorageBin(input *dpfm_api_input_reader.StorageBinSDC) *dpfm_api_output_formatter.ProductMasterStorageBin {
	exconf := dpfm_api_output_formatter.ProductMasterStorageBin{
		ExistenceConf: false,
	}
	if input.ProductMasterStorageBin.Product == nil {
		return &exconf
	}
	if input.ProductMasterStorageBin.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterStorageBin.Plant == nil {
		return &exconf
	}
	if input.ProductMasterStorageBin.StorageLocation == nil {
		return &exconf
	}
	if input.ProductMasterStorageBin.StorageBin == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.ProductMasterStorageBin{
		Product: 			*input.ProductMasterStorageBin.Product,
		BusinessPartner:	*input.ProductMasterStorageBin.BusinessPartner,
		Plant:				*input.ProductMasterStorageBin.Plant,
		StorageLocation:	*input.ProductMasterStorageBin.StorageLocation,
		StorageBin:			*input.ProductMasterStorageBin.StorageBin,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_bin_data 
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant, exconf.StorageLocation, exconf.StorageBin,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterTax(input *dpfm_api_input_reader.TaxSDC) *dpfm_api_output_formatter.ProductMasterTax {
	exconf := dpfm_api_output_formatter.ProductMasterTax{
		ExistenceConf: false,
	}
	if input.ProductMasterTax.Product == nil {
		return &exconf
	}
	if input.ProductMasterTax.Country == nil {
		return &exconf
	}
	if input.ProductMasterTax.ProductTaxCategory == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterTax{
		Product:			*input.ProductMasterTax.Product,
		Country:			*input.ProductMasterTax.Country,
		ProductTaxCategory: *input.ProductMasterTax.ProductTaxCategory,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data 
		WHERE (Product, Country, ProductTaxCategory) = (?);`, exconf.Product, exconf.Country, exconf.ProductTaxCategory,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterProductDescription(input *dpfm_api_input_reader.ProductDescriptionSDC) *dpfm_api_output_formatter.ProductMasterProductDescription {
	exconf := dpfm_api_output_formatter.ProductMasterProductDescription{
		ExistenceConf: false,
	}
	if input.ProductMasterProductDescription.Product == nil {
		return &exconf
	}
	if input.ProductMasterProductDescription.Language == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterProductDescription{
		Product:		*input.ProductMasterProductDescription.Product,
		Language:		*input.ProductMasterProductDescription.Language,
		ExistenceConf:	false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data 
		WHERE (Product, Language) = (?);`, exconf.Product, exconf.Language,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterProductDescByBP(input *dpfm_api_input_reader.ProductDescByBPSDC) *dpfm_api_output_formatter.ProductMasterProductDescByBP {
	exconf := dpfm_api_output_formatter.ProductMasterProductDescByBP{
		ExistenceConf: false,
	}
	if input.ProductMasterProductDescByBP.Product == nil {
		return &exconf
	}
	if input.ProductMasterProductDescByBP.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMasterProductDescByBP.Language == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ProductMasterProductDescByBP{
		Product:			*input.ProductMasterProductDescByBP.Product,
		BusinessPartner:	*input.ProductMasterProductDescByBP.BusinessPartner,
		Language:			*input.ProductMasterProductDescByBP.Language,
		ExistenceConf:		false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data 
		WHERE (Product, BusinessPartner, Language) = (?);`, exconf.Product, exconf.BusinessPartner, exconf.Language,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
