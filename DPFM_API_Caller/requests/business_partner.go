package requests

type BusinessPartner struct {
	Product           *string `json:"Product"`
	BusinessPartner   *int    `json:"BusinessPartner"`
	ValidityEndDate   *string `json:"ValidityEndDate"`
	ValidityStartDate *string `json:"ValidityStartDate"`
}
