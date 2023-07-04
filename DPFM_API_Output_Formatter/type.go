package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey                                  string                                          `json:"connection_key"`
	Result                                         bool                                            `json:"result"`
	RedisKey                                       string                                          `json:"redis_key"`
	Filepath                                       string                                          `json:"filepath"`
	APIStatusCode                                  int                                             `json:"api_status_code"`
	RuntimeSessionID                               string                                          `json:"runtime_session_id"`
	BusinessPartnerID                              *int                                            `json:"business_partner"`
	ServiceLabel                                   string                                          `json:"service_label"`
	ProductMasterGeneral                		   *ProductMasterGeneral                 		   `json:"ProductMasterGeneral,omitempty"`
	ProductMasterBusinessPartner                   *ProductMasterBusinessPartner        		   `json:"ProductMasterBusinessPartner,omitempty"`
	ProductMasterBPPlant                           *ProductMasterBPPlant        		 		   `json:"ProductMasterBPPlant,omitempty"`
	ProductMasterMRPArea                           *ProductMasterMRPArea        		 		   `json:"ProductMasterMRPArea,omitempty"`
	ProductMasterProduction                        *ProductMasterProduction        		 		   `json:"ProductMasterProduction,omitempty"`
	ProductMasterStorageLocation                   *ProductMasterStorageLocation        		   `json:"ProductMasterStorageLocation,omitempty"`
	ProductMasterQuality              	      	   *ProductMasterQuality        				   `json:"ProductMasterQuality,omitempty"`
	ProductMasterAccounting              	       *ProductMasterAccounting        				   `json:"ProductMasterAccounting,omitempty"`
	ProductMasterStorageBin              	       *ProductMasterStorageBin        				   `json:"ProductMasterStorageBin,omitempty"`
	ProductMasterTax                		   	   *ProductMasterTax                 			   `json:"ProductMasterTax,omitempty"`
	ProductMasterProductDescription                *ProductMasterProductDescription       		   `json:"ProductMasterProductDescription,omitempty"`
	ProductMasterProductDescByBP              	   *ProductMasterProductDescByBP      	 		   `json:"ProductMasterProductDescByBP,omitempty"`
	APISchema                                      string                                          `json:"api_schema"`
	Accepter                                       []string                                        `json:"accepter"`
	Deleted                                        bool                                            `json:"deleted"`
}

type ProductMasterGeneral struct {
	Product					  string	`json:"Product"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterBusinessPartner struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterBPPlant struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterMRPArea struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	MRPArea	  				  string  	`json:"MRPArea"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterProduction struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterStorageLocation struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	StorageLocation	  		  string  	`json:"StorageLocation"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterQuality struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterAccounting struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterStorageBin struct {
	Product					  string	`json:"Product"`
	BusinessPartner	  		  int    	`json:"BusinessPartner"`
	Plant	  				  string  	`json:"Plant"`
	StorageLocation	  		  string  	`json:"StorageLocation"`
	StorageBin	  			  string  	`json:"StorageBin"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterTax struct {
	Product					  string	`json:"Product"`
	Country                   string	`json:"Country"`
	ProductTaxCategory        string	`json:"ProductTaxCategory"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterProductDescription struct {
	Product					  string	`json:"Product"`
	Language				  string	`json:"Language"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}

type ProductMasterProductDescByBP struct {
	Product					  string	`json:"Product"`
	BusinessPartner			  int		`json:"BusinessPartner"`
	Language				  string	`json:"Language"`
	ExistenceConf             bool		`json:"ExistenceConf"`
}
