package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type GeneralSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterGeneral struct {
		Product					  *string `json:"Product"`
	} `json:"ProductMasterGeneral"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type BusinessPartnerSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterBusinessPartner struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
	} `json:"ProductMasterBusinessPartner"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type BPPlantSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterBPPlant struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		} `json:"ProductMasterBPPlant"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type MRPAreaSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterMRPArea struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		MRPArea	  				   *string `json:"MRPArea"`
		} `json:"ProductMasterMRPArea"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type ProductionSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterProduction struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		} `json:"ProductMasterProduction"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type StorageLocationSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterStorageLocation struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		StorageLocation	  		   *string `json:"StorageLocation"`
		} `json:"ProductMasterStorageLocation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type QualitySDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterQuality struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		} `json:"ProductMasterQuality"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type AccountingSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterAccounting struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		} `json:"ProductMasterAccounting"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type StorageBinSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterStorageBin struct {
		Product					   *string `json:"Product"`
		BusinessPartner			   *int    `json:"BusinessPartner"`
		Plant	  				   *string `json:"Plant"`
		StorageLocation	  		   *string `json:"StorageLocation"`
		StorageBin	  			   *string `json:"StorageBin"`
		} `json:"ProductMasterStorageBin"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type TaxSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	ProductMasterTax struct {
		Product					  *string `json:"Product"`
		Country                   *string `json:"Country"`
		ProductTaxCategory        *string `json:"ProductTaxCategory"`
	} `json:"ProductMasterTax"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type ProductDescriptionSDC struct {
	ConnectionKey                   string `json:"connection_key"`
	Result                          bool   `json:"result"`
	RedisKey                        string `json:"redis_key"`
	Filepath                        string `json:"filepath"`
	APIStatusCode                   int    `json:"api_status_code"`
	RuntimeSessionID                string `json:"runtime_session_id"`
	BusinessPartner                 *int   `json:"business_partner"`
	ServiceLabel                    string `json:"service_label"`
	ProductMasterProductDescription struct {
		Product					    *string `json:"Product"`
		Language					*string `json:"Language"`
	} `json:"ProductMasterProductDescription"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type ProductDescByBPSDC struct {
	ConnectionKey                   string `json:"connection_key"`
	Result                          bool   `json:"result"`
	RedisKey                        string `json:"redis_key"`
	Filepath                        string `json:"filepath"`
	APIStatusCode                   int    `json:"api_status_code"`
	RuntimeSessionID                string `json:"runtime_session_id"`
	BusinessPartner                 *int   `json:"business_partner"`
	ServiceLabel                    string `json:"service_label"`
	ProductMasterProductDescByBP struct {
		Product					    *string `json:"Product"`
		BusinessPartner			    *int    `json:"BusinessPartner"`
		Language					*string `json:"Language"`
	} `json:"ProductMasterProductDescByBP"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}
