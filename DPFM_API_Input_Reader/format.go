package dpfm_api_input_reader

import (
	"data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.ProductMaster.General
	return &requests.General{
		Product: data.Product,
	}
}

func (sdc *SDC) ConvertToProductDescription() *requests.ProductDescription {
	data := sdc.ProductMaster.ProductDescription
	return &requests.ProductDescription{
		Product:  data.Product,
		Language: data.Language,
	}
}

func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
	data := sdc.ProductMaster.BusinessPartner
	return &requests.BusinessPartner{
		Product:           data.Product,
		BusinessPartner:   data.BusinessPartner,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
	}
}

func (sdc *SDC) ConvertToProductDescByBP() *requests.ProductDescByBP {
	data := sdc.ProductMaster.ProductDescByBP
	return &requests.ProductDescByBP{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Language:        data.Language,
	}
}

func (sdc *SDC) ConvertToBPPlant() *requests.BPPlant {
	data := sdc.ProductMaster.BPPlant
	return &requests.BPPlant{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}
}

func (sdc *SDC) ConvertToStorageLocation() *requests.StorageLocation {
	data := sdc.ProductMaster.StorageLocation
	return &requests.StorageLocation{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
		StorageLocation: data.StorageLocation,
	}
}

func (sdc *SDC) ConvertToMRPArea() *requests.MRPArea {
	data := sdc.ProductMaster.MRPArea
	return &requests.MRPArea{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
		MRPArea:         data.MRPArea,
	}
}

func (sdc *SDC) ConvertToWorkScheduling() *requests.WorkScheduling {
	data := sdc.ProductMaster.WorkScheduling
	return &requests.WorkScheduling{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	data := sdc.ProductMaster.Accounting
	return &requests.Accounting{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}
}

func (sdc *SDC) ConvertToTax() *requests.Tax {
	data := sdc.ProductMaster.Tax
	return &requests.Tax{
		Product:            data.Product,
		Country:            data.Country,
		ProductTaxCategory: data.ProductTaxCategory,
	}
}

func (sdc *SDC) ConvertToAllergen() *requests.Allergen {
	data := sdc.ProductMaster.Allergen
	return &requests.Allergen{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Allergen:        data.Allergen,
	}
}

func (sdc *SDC) ConvertToCalories() *requests.Calories {
	data := sdc.ProductMaster.Calories
	return &requests.Calories{
		Product:            data.Product,
		BusinessPartner:    data.BusinessPartner,
		CaloryUnitQuantity: data.CaloryUnitQuantity,
	}
}

func (sdc *SDC) ConvertToNutritionalInfo() *requests.NutritionalInfo {
	data := sdc.ProductMaster.NutritionalInfo
	return &requests.NutritionalInfo{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Nutrient:        data.Nutrient,
	}
}

func (sdc *SDC) ConvertToQuality() *requests.Quality {
	data := sdc.ProductMaster.Quality
	return &requests.Quality{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}
}

func (sdc *SDC) ConvertToStorageBin() *requests.StorageBin {
	data := sdc.ProductMaster.StorageBin
	return &requests.StorageBin{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
		StorageLocation: data.StorageLocation,
		StorageBin:      data.StorageBin,
	}
}
