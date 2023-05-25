package requests

type Allergen struct {
	Product         *string `json:"Product"`
	BusinessPartner *int    `json:"BusinessPartner"`
	Allergen        *string `json:"Allergen"`
}
