package requests

type ProductMasterBPPlant struct {
	Product         *string `json:"Product"`
	BusinessPartner *int    `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
}
