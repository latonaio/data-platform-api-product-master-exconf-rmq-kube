package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Output_Formatter"
)

func (e *ExistenceConf) confProductMasterGeneral(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.General {
	exconf := dpfm_api_output_formatter.General{
		ExistenceConf: false,
	}
	if input.ProductMaster.General.Product == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToGeneral(input)

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

func (e *ExistenceConf) confProductMasterProductDescription(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.ProductDescription {
	exconf := dpfm_api_output_formatter.ProductDescription{
		ExistenceConf: false,
	}
	if input.ProductMaster.ProductDescription.Product == nil {
		return &exconf
	}
	if input.ProductMaster.ProductDescription.Language == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToProductDescription(input)

	rows, err := e.db.Query(
		`SELECT Product 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data 
		WHERE (Product, Language) = (?, ?);`, exconf.Product, exconf.Language,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterBusinessPartner(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.BusinessPartner {
	exconf := dpfm_api_output_formatter.BusinessPartner{
		ExistenceConf: false,
	}
	if input.ProductMaster.BusinessPartner.Product == nil {
		return &exconf
	}
	if input.ProductMaster.BusinessPartner.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.BusinessPartner.ValidityEndDate == nil {
		return &exconf
	}
	if input.ProductMaster.BusinessPartner.ValidityStartDate == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToBusinessPartner(input)

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

func (e *ExistenceConf) confProductMasterProductDescByBP(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.ProductDescByBP {
	exconf := dpfm_api_output_formatter.ProductDescByBP{
		ExistenceConf: false,
	}
	if input.ProductMaster.ProductDescByBP.Product == nil {
		return &exconf
	}
	if input.ProductMaster.ProductDescByBP.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.ProductDescByBP.Language == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToProductDescByBP(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		WHERE (Product, BusinessPartner, Language) = (?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Language,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterBPPlant(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.BPPlant {
	exconf := dpfm_api_output_formatter.BPPlant{
		ExistenceConf: false,
	}
	if input.ProductMaster.BPPlant.Product == nil {
		return &exconf
	}
	if input.ProductMaster.BPPlant.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.BPPlant.Plant == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToBPPlant(input)

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

func (e *ExistenceConf) confProductMasterStorageLocation(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.StorageLocation {
	exconf := dpfm_api_output_formatter.StorageLocation{
		ExistenceConf: false,
	}
	if input.ProductMaster.StorageLocation.Product == nil {
		return &exconf
	}
	if input.ProductMaster.StorageLocation.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.StorageLocation.Plant == nil {
		return &exconf
	}
	if input.ProductMaster.StorageLocation.StorageLocation == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToStorageLocation(input)

	rows, err := e.db.Query(
		`SELECT Product
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

func (e *ExistenceConf) confProductMasterMRPArea(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.MRPArea {
	exconf := dpfm_api_output_formatter.MRPArea{
		ExistenceConf: false,
	}
	if input.ProductMaster.MRPArea.Product == nil {
		return &exconf
	}
	if input.ProductMaster.MRPArea.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.MRPArea.Plant == nil {
		return &exconf
	}
	if input.ProductMaster.MRPArea.MRPArea == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToMRPArea(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_mrp_area_data
		WHERE (Product, BusinessPartner, Plant, MRPArea) = (?, ?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant, exconf.MRPArea,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterWorkScheduling(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.WorkScheduling {
	exconf := dpfm_api_output_formatter.WorkScheduling{
		ExistenceConf: false,
	}
	if input.ProductMaster.WorkScheduling.Product == nil {
		return &exconf
	}
	if input.ProductMaster.WorkScheduling.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.WorkScheduling.Plant == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToWorkScheduling(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_work_scheduling_data
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

func (e *ExistenceConf) confProductMasterAccounting(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.Accounting {
	exconf := dpfm_api_output_formatter.Accounting{
		ExistenceConf: false,
	}
	if input.ProductMaster.Accounting.Product == nil {
		return &exconf
	}
	if input.ProductMaster.Accounting.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.Accounting.Plant == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToAccounting(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_accounting_data
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

func (e *ExistenceConf) confProductMasterQuality(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.Quality {
	exconf := dpfm_api_output_formatter.Quality{
		ExistenceConf: false,
	}
	if input.ProductMaster.Quality.Product == nil {
		return &exconf
	}
	if input.ProductMaster.Quality.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.Quality.Plant == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToQuality(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_quality_data
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

func (e *ExistenceConf) confProductMasterAllergen(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.Allergen {
	exconf := dpfm_api_output_formatter.Allergen{
		ExistenceConf: false,
	}
	if input.ProductMaster.Allergen.Product == nil {
		return &exconf
	}
	if input.ProductMaster.Allergen.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.Allergen.Allergen == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToAllergen(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_allergen_data
		WHERE (Product, BusinessPartner, Allergen) = (?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Allergen,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterNutritionalInfo(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.NutritionalInfo {
	exconf := dpfm_api_output_formatter.NutritionalInfo{
		ExistenceConf: false,
	}
	if input.ProductMaster.NutritionalInfo.Product == nil {
		return &exconf
	}
	if input.ProductMaster.NutritionalInfo.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.NutritionalInfo.Nutrient == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToNutritionalInfo(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_nutritional_info_data
		WHERE (Product, BusinessPartner, Nutrient) = (?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Nutrient,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterCalories(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.Calories {
	exconf := dpfm_api_output_formatter.Calories{
		ExistenceConf: false,
	}
	if input.ProductMaster.Calories.Product == nil {
		return &exconf
	}
	if input.ProductMaster.Calories.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.Calories.CaloryUnitQuantity == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToCalories(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_calories_data
		WHERE (Product, BusinessPartner, CaloryUnitQuantity) = (?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.CaloryUnitQuantity,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterTax(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.Tax {
	exconf := dpfm_api_output_formatter.Tax{
		ExistenceConf: false,
	}
	if input.ProductMaster.Tax.Product == nil {
		return &exconf
	}
	if input.ProductMaster.Tax.Country == nil {
		return &exconf
	}
	if input.ProductMaster.Tax.ProductTaxCategory == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToTax(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE (Product, Country, ProductTaxCategory) = (?, ?, ?);`, exconf.Product, exconf.Country, exconf.ProductTaxCategory,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confProductMasterStorageBin(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.StorageBin {
	exconf := dpfm_api_output_formatter.StorageBin{
		ExistenceConf: false,
	}
	if input.ProductMaster.StorageBin.Product == nil {
		return &exconf
	}
	if input.ProductMaster.StorageBin.BusinessPartner == nil {
		return &exconf
	}
	if input.ProductMaster.StorageBin.Plant == nil {
		return &exconf
	}
	if input.ProductMaster.StorageBin.StorageLocation == nil {
		return &exconf
	}
	if input.ProductMaster.StorageBin.StorageBin == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.ConvertToStorageBin(input)

	rows, err := e.db.Query(
		`SELECT Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_bin_data
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin) = (?, ?, ?, ?, ?);`, exconf.Product, exconf.BusinessPartner, exconf.Plant, exconf.StorageLocation, exconf.StorageBin,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
