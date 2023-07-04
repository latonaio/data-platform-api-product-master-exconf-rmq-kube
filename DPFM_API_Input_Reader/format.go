package dpfm_api_input_reader

import (
	"data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *GeneralSDC) ConvertToProductMasterGeneral() *requests.ProductMasterGeneral {
	data := sdc.ProductMasterGeneral
	return &requests.ProductMasterGeneral{
		Product: data.Product,
	}
}

func (sdc *BusinessPartnerSDC) ConvertToProductMasterBusinessPartner() *requests.ProductMasterBusinessPartner {
	data := sdc.ProductMasterBusinessPartner
	return &requests.ProductMasterBusinessPartner{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
	}
}

func (sdc *BPPlantSDC) ConvertToProductMasterBPPlant() *requests.ProductMasterBPPlant {
	data := sdc.ProductMasterBPPlant
	return &requests.ProductMasterBPPlant{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
	}
}

func (sdc *MRPAreaSDC) ConvertToProductMasterMRPArea() *requests.ProductMasterMRPArea {
	data := sdc.ProductMasterMRPArea
	return &requests.ProductMasterMRPArea{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
		MRPArea:		 data.MRPArea,
	}
}

func (sdc *ProductionSDC) ConvertToProductMasterProduction() *requests.ProductMasterProduction {
	data := sdc.ProductMasterProduction
	return &requests.ProductMasterProduction{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
	}
}

func (sdc *StorageLocationSDC) ConvertToProductMasterStorageLocation() *requests.ProductMasterStorageLocation {
	data := sdc.ProductMasterStorageLocation
	return &requests.ProductMasterStorageLocation{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
		StorageLocation: data.StorageLocation,
	}
}

func (sdc *QualitySDC) ConvertToProductMasterQuality() *requests.ProductMasterQuality {
	data := sdc.ProductMasterQuality
	return &requests.ProductMasterQuality{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
	}
}

func (sdc *AccountingSDC) ConvertToProductMasterAccounting() *requests.ProductMasterAccounting {
	data := sdc.ProductMasterAccounting
	return &requests.ProductMasterAccounting{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
	}
}

func (sdc *StorageBinSDC) ConvertToProductMasterStorageBin() *requests.ProductMasterStorageBin {
	data := sdc.ProductMasterStorageBin
	return &requests.ProductMasterStorageBin{
		Product: 		 data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:			 data.Plant,
		StorageLocation: data.StorageLocation,
		StorageBin:		 data.StorageBin,
	}
}

func (sdc *TaxSDC) ConvertToProductMasterTax() *requests.ProductMasterTax {
	data := sdc.ProductMasterTax
	return &requests.ProductMasterTax{
		Product:			data.Product,
		Country:			data.Country,
		ProductTaxCategory: data.ProductTaxCategory,
	}
}

func (sdc *ProductDescriptionSDC) ConvertToProductMasterProductDescription() *requests.ProductMasterProductDescription {
	data := sdc.ProductMasterProductDescription
	return &requests.ProductMasterProductDescription{
		Product:	data.Product,
		Language:	data.Language,
	}
}

func (sdc *ProductDescByBPSDC) ConvertToProductMasterProductDescByBP() *requests.ProductMasterProductDescByBP {
	data := sdc.ProductMasterProductDescByBP
	return &requests.ProductMasterProductDescByBP{
		Product:			data.Product,
		BusinessPartner:	data.BusinessPartner,
		Language:			data.Language,
	}
}
