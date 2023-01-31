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
	case *ProductMasterStorageLocation:
		output.ProductMasterStorageLocation = exconf
	default:
		return nil, xerrors.Errorf("unknown type %+v", exconf)
	}

	return output, nil
}
