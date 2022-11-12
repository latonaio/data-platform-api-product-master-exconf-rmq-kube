package dpfm_api_input_reader

import (
	"data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.ProductMasterGeneral
	return &requests.General{
		Product: data.Product,
	}
}
