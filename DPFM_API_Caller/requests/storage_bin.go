package requests

type ProductMasterStorageBin struct {
	Product 		  *string	`json:"Product"`
	BusinessPartner	  *int    	`json:"BusinessPartner"`
	Plant	  		  *string  	`json:"Plant"`
	StorageLocation	  *string   `json:"StorageLocation"`
	StorageBin	  	  *string 	`json:"StorageBin"`
}