package dpfm_api_input_reader

import (
	"data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToProductMasterGeneral() *requests.ProductMasterGeneral {
	data := sdc.ProductMasterGeneral
	return &requests.ProductMasterGeneral{
		Product: data.Product,
	}
}
