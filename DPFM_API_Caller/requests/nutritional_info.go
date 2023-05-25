package requests

type NutritionalInfo struct {
	Product         *string `json:"Product"`
	BusinessPartner *int    `json:"BusinessPartner"`
	Nutrient        *string `json:"Nutrient"`
}
