package main

import (
	"context"
	"data-platform-api-product-master-exconf-rmq-kube/config"
	"data-platform-api-product-master-exconf-rmq-kube/database"
	"data-platform-api-product-master-exconf-rmq-kube/formatter"
	"encoding/json"
	"fmt"

	"github.com/latonaio/golang-logging-library/logger"
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

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), "", nil, -1)
	if err != nil {
		l.Fatal(err.Error())
	}
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	for msg := range iter {
		go dataCheckProcess(ctx, c, db, msg)
	}
}

func dataCheckProcess(
	ctx context.Context,
	c *config.Conf,
	db *database.Mysql,
	rmqMsg rabbitmq.RabbitmqMessage,
) {
	defer rmqMsg.Success()
	l := logger.NewLogger()
	data := rmqMsg.Data()
	sessionId := getBodyHeader(data)
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": sessionId})

	exist := (*ExistencyChecker).Check(NewExistencyChecker(ctx, db, l), data)
	rmqMsg.Respond(exist)
	l.Info(exist)

	inputParams, err := inputParamsToMarshall(l, rmqMsg)
	if err != nil {
		l.Fatal("inputParamsToMarshall error")
	}

	l.Info(inputParams.ServiceLabel)
}

func getBodyHeader(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func inputParamsToMarshall(l *logger.Logger, msg rabbitmq.RabbitmqMessage) (*formatter.InputParams, error) {
	raw, err := json.Marshal(msg.Data())
	if err != nil {
		l.Fatal("data marshal error", err)
		return nil, err
	}

	inputParams := formatter.InputParams{}
	err = json.Unmarshal(raw, &inputParams)
	if err != nil {
		l.Fatal("inputServiceParam unmarshal error", err)
		return nil, err
	}

	return &inputParams, nil
}
