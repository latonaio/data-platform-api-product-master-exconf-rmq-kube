package requests

type ProductMasterProductDescByBP struct {
	Product 	        *string `json:"Product"`
	BusinessPartner     *int    `json:"BusinessPartner"`
	Language            *string `json:"Language"`
}
