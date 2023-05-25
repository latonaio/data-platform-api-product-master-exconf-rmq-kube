package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey     string        `json:"connection_key"`
	Result            bool          `json:"result"`
	RedisKey          string        `json:"redis_key"`
	Filepath          string        `json:"filepath"`
	APIStatusCode     int           `json:"api_status_code"`
	RuntimeSessionID  string        `json:"runtime_session_id"`
	BusinessPartnerID *int          `json:"business_partner"`
	ServiceLabel      string        `json:"service_label"`
	ProductMaster     ProductMaster `json:"ProductMaster"`
	APISchema         string        `json:"api_schema"`
	Accepter          []string      `json:"accepter"`
	Deleted           bool          `json:"deleted"`
}

type ProductMaster struct {
	General            *General            `json:"General"`
	ProductDescription *ProductDescription `json:"ProductDescription"`
	BusinessPartner    *BusinessPartner    `json:"BusinessPartner"`
	ProductDescByBP    *ProductDescByBP    `json:"ProductDescByBP"`
	BPPlant            *BPPlant            `json:"BPPlant"`
	StorageLocation    *StorageLocation    `json:"StorageLocation"`
	MRPArea            *MRPArea            `json:"MRPArea"`
	WorkScheduling     *WorkScheduling     `json:"WorkScheduling"`
	Accounting         *Accounting         `json:"Accounting"`
	Tax                *Tax                `json:"Tax"`
	Allergen           *Allergen           `json:"Allergen"`
	Calories           *Calories           `json:"Calories"`
	NutritionalInfo    *NutritionalInfo    `json:"NutritionalInfo"`
	Quality            *Quality            `json:"Quality"`
	StorageBin         *StorageBin         `json:"StorageBin"`
}

type General struct {
	Product       string `json:"Product"`
	ExistenceConf bool   `json:"ExistenceConf"`
}

type ProductDescription struct {
	Product       string `json:"Product"`
	Language      string `json:"Language"`
	ExistenceConf bool   `json:"ExistenceConf"`
}

type BusinessPartner struct {
	Product           string `json:"Product"`
	BusinessPartner   int    `json:"BusinessPartner"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ExistenceConf     bool   `json:"ExistenceConf"`
}

type ProductDescByBP struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Language        string `json:"Language"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type BPPlant struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type StorageLocation struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type MRPArea struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	MRPArea         string `json:"MRPArea"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type WorkScheduling struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type Accounting struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type Tax struct {
	Product            string `json:"Product"`
	Country            string `json:"Country"`
	ProductTaxCategory string `json:"ProductTaxCategory"`
	ExistenceConf      bool   `json:"ExistenceConf"`
}

type Allergen struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Allergen        string `json:"Allergen"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type Calories struct {
	Product            string `json:"Product"`
	BusinessPartner    int    `json:"BusinessPartner"`
	CaloryUnitQuantity int    `json:"CaloryUnitQuantity"`
	ExistenceConf      bool   `json:"ExistenceConf"`
}

type NutritionalInfo struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Nutrient        string `json:"Nutrient"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type Quality struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type StorageBin struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
	StorageBin      string `json:"StorageBin"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}
