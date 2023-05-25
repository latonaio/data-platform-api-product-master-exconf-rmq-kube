package requests

type ProductDescByBP struct {
	Product         *string `json:"Product"`
	BusinessPartner *int    `json:"BusinessPartner"`
	Language        *string `json:"Language"`
}
