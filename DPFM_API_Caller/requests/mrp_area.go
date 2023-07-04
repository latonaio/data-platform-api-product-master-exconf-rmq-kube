package requests

type ProductMasterMRPArea struct {
	Product 		  *string	`json:"Product"`
	BusinessPartner	  *int    	`json:"BusinessPartner"`
	Plant	  		  *string  	`json:"Plant"`
	MRPArea           *string   `json:"MRPArea"`
}
