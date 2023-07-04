package requests

type ProductMasterTax struct {
	Product                *string `json:"Product"`
	Country                *string `json:"Country"`
	ProductTaxCategory     *string `json:"ProductTaxCategory"`
}
