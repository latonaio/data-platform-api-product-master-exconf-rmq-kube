package dpfm_api_output_formatter

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func NewOutput(rmqMsg rabbitmq.RabbitmqMessage, exconf interface{}) (*MetaData, error) {
	output := &MetaData{}
	err := json.Unmarshal(rmqMsg.Raw(), output)
	if err != nil {
		return nil, xerrors.Errorf("output Marshal error: %w", err)
	}

	switch exconf := exconf.(type) {
	case *ProductMasterGeneral:
		output.ProductMasterGeneral = exconf
	case *ProductMasterBusinessPartner:
		output.ProductMasterBusinessPartner = exconf
	case *ProductMasterBPPlant:
		output.ProductMasterBPPlant = exconf
	case *ProductMasterMRPArea:
		output.ProductMasterMRPArea = exconf
	case *ProductMasterProduction:
		output.ProductMasterProduction = exconf
	case *ProductMasterStorageLocation:
		output.ProductMasterStorageLocation = exconf
	case *ProductMasterQuality:
		output.ProductMasterQuality = exconf
	case *ProductMasterAccounting:
		output.ProductMasterAccounting = exconf
	case *ProductMasterStorageBin:
		output.ProductMasterStorageBin = exconf
	case *ProductMasterTax:
		output.ProductMasterTax = exconf
	case *ProductMasterProductDescription:
		output.ProductMasterProductDescription = exconf
	case *ProductMasterProductDescByBP:
		output.ProductMasterProductDescByBP = exconf
	default:
		return nil, xerrors.Errorf("unknown type %+v", exconf)
	}

	return output, nil
}
