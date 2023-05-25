package main

import (
	"context"
	dpfm_api_caller "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Caller"
	dpfm_api_input_reader "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-exconf-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-product-master-exconf-rmq-kube/config"
	"encoding/json"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}
	defer db.Close()

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), "", nil, -1)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	for msg := range iter {
		output, err := dataCallProcess(ctx, c, db, msg)
		if err != nil {
			l.Error(err)
			data := msg.Data()
			data["result"] = false
			msg.Respond(data)
		}
		msg.Respond(output)
	}
}

func dataCallProcess(
	ctx context.Context,
	c *config.Conf,
	db *database.Mysql,
	rmqMsg rabbitmq.RabbitmqMessage,
) (*dpfm_api_output_formatter.SDC, error) {
	var err error
	defer rmqMsg.Success()
	l := logger.NewLogger()
	sessionId := getBodyHeader(rmqMsg.Data())
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": sessionId})
	conf := dpfm_api_caller.NewExistenceConf(ctx, db, l)

	var input dpfm_api_input_reader.SDC
	var output dpfm_api_output_formatter.SDC

	err = json.Unmarshal(rmqMsg.Raw(), &input)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	err = json.Unmarshal(rmqMsg.Raw(), &output)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	accepter := getAccepter(&input)
	res := conf.Conf(accepter, &input)
	output.ProductMaster = res

	l.JsonParseOut(output)

	return &output, nil
}

func getBodyHeader(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func getAccepter(input *dpfm_api_input_reader.SDC) []string {
	accepter := input.Accepter
	if len(input.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"General", "BusinessPartner", "BPPlant", "StorageLocation", "MRPArea", "WorkScheduling", "Accounting",
			"Allergen", "Calories", "NutritionalInfo", "Quality", "StorageBin", "Tax", "ProductDescription", "ProductDescByBP",
		}
	}
	return accepter
}
