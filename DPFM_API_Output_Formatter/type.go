package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey                string                        `json:"connection_key"`
	Result                       bool                          `json:"result"`
	RedisKey                     string                        `json:"redis_key"`
	Filepath                     string                        `json:"filepath"`
	APIStatusCode                int                           `json:"api_status_code"`
	RuntimeSessionID             string                        `json:"runtime_session_id"`
	BusinessPartnerID            *int                          `json:"business_partner"`
	ServiceLabel                 string                        `json:"service_label"`
	ProductMasterGeneral         *ProductMasterGeneral         `json:"ProductMasterGeneral,omitempty"`
	ProductMasterBusinessPartner *ProductMasterBusinessPartner `json:"ProductMasterBusinessPartner,omitempty"`
	ProductMasterBPPlant         *ProductMasterBPPlant         `json:"ProductMasterBPPlant,omitempty"`
	ProductMasterStorageLocation *ProductMasterStorageLocation `json:"ProductMasterStorageLocation,omitempty"`
	APISchema                    string                        `json:"api_schema"`
	Accepter                     []string                      `json:"accepter"`
	Deleted                      bool                          `json:"deleted"`
}

type ProductMasterGeneral struct {
	Product       string `json:"Product"`
	ExistenceConf bool   `json:"ExistenceConf"`
}

type ProductMasterBusinessPartner struct {
	Product           string `json:"Product"`
	BusinessPartner   int    `json:"BusinessPartner"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ExistenceConf     bool   `json:"ExistenceConf"`
}

type ProductMasterBPPlant struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type ProductMasterStorageLocation struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}
