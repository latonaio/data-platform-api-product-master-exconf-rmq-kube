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
		Product:           data.Product,
		BusinessPartner:   data.BusinessPartner,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
	}
}

func (sdc *BPPlantSDC) ConvertToProductMasterBPPlant() *requests.ProductMasterBPPlant {
	data := sdc.ProductMasterBPPlant
	return &requests.ProductMasterBPPlant{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}
}

func (sdc *StorageLocationSDC) ConvertToProductMasterStorageLocation() *requests.ProductMasterStorageLocation {
	data := sdc.ProductMasterStorageLocation
	return &requests.ProductMasterStorageLocation{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
		StorageLocation: data.StorageLocation,
	}
}
