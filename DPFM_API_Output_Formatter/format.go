package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneral(input *dpfm_api_input_reader.SDC) General {
	data := input.ProductMaster.General
	return General{
		Product:       *data.Product,
		ExistenceConf: false,
	}
}

func ConvertToProductDescription(input *dpfm_api_input_reader.SDC) ProductDescription {
	data := input.ProductMaster.ProductDescription
	return ProductDescription{
		Product:       *data.Product,
		Language:      *data.Language,
		ExistenceConf: false,
	}
}

func ConvertToBusinessPartner(input *dpfm_api_input_reader.SDC) BusinessPartner {
	data := input.ProductMaster.BusinessPartner
	return BusinessPartner{
		Product:           *data.Product,
		BusinessPartner:   *data.BusinessPartner,
		ValidityEndDate:   *data.ValidityEndDate,
		ValidityStartDate: *data.ValidityStartDate,
		ExistenceConf:     false,
	}
}

func ConvertToProductDescByBP(input *dpfm_api_input_reader.SDC) ProductDescByBP {
	data := input.ProductMaster.ProductDescByBP
	return ProductDescByBP{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Language:        *data.Language,
		ExistenceConf:   false,
	}
}

func ConvertToBPPlant(input *dpfm_api_input_reader.SDC) BPPlant {
	data := input.ProductMaster.BPPlant
	return BPPlant{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		ExistenceConf:   false,
	}
}

func ConvertToStorageLocation(input *dpfm_api_input_reader.SDC) StorageLocation {
	data := input.ProductMaster.StorageLocation
	return StorageLocation{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		StorageLocation: *data.StorageLocation,
		ExistenceConf:   false,
	}
}

func ConvertToMRPArea(input *dpfm_api_input_reader.SDC) MRPArea {
	data := input.ProductMaster.MRPArea
	return MRPArea{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		MRPArea:         *data.MRPArea,
		ExistenceConf:   false,
	}
}

func ConvertToWorkScheduling(input *dpfm_api_input_reader.SDC) WorkScheduling {
	data := input.ProductMaster.WorkScheduling
	return WorkScheduling{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		ExistenceConf:   false,
	}
}

func ConvertToAccounting(input *dpfm_api_input_reader.SDC) Accounting {
	data := input.ProductMaster.Accounting
	return Accounting{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		ExistenceConf:   false,
	}
}

func ConvertToTax(input *dpfm_api_input_reader.SDC) Tax {
	data := input.ProductMaster.Tax
	return Tax{
		Product:            *data.Product,
		Country:            *data.Country,
		ProductTaxCategory: *data.ProductTaxCategory,
		ExistenceConf:      false,
	}
}

func ConvertToAllergen(input *dpfm_api_input_reader.SDC) Allergen {
	data := input.ProductMaster.Allergen
	return Allergen{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Allergen:        *data.Allergen,
		ExistenceConf:   false,
	}
}

func ConvertToCalories(input *dpfm_api_input_reader.SDC) Calories {
	data := input.ProductMaster.Calories
	return Calories{
		Product:            *data.Product,
		BusinessPartner:    *data.BusinessPartner,
		CaloryUnitQuantity: *data.CaloryUnitQuantity,
		ExistenceConf:      false,
	}
}

func ConvertToNutritionalInfo(input *dpfm_api_input_reader.SDC) NutritionalInfo {
	data := input.ProductMaster.NutritionalInfo
	return NutritionalInfo{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Nutrient:        *data.Nutrient,
		ExistenceConf:   false,
	}
}

func ConvertToQuality(input *dpfm_api_input_reader.SDC) Quality {
	data := input.ProductMaster.Quality
	return Quality{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		ExistenceConf:   false,
	}
}

func ConvertToStorageBin(input *dpfm_api_input_reader.SDC) StorageBin {
	data := input.ProductMaster.StorageBin
	return StorageBin{
		Product:         *data.Product,
		BusinessPartner: *data.BusinessPartner,
		Plant:           *data.Plant,
		StorageLocation: *data.StorageLocation,
		StorageBin:      *data.StorageBin,
		ExistenceConf:   false,
	}
}
