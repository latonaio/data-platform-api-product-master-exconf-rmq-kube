# data-platform-api-supply-chain-relationship-exconf-rmq-kube
data-platform-api-supply-chain-relationship-exconf-rmq-kube は、データ連携基盤において、API で サプライチェーンリレーションシップの存在性チェックを行うためのマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、サプライチェーンリレーションシップの存在確認が行われます。

* data-platform-supply-chain-relationship-sql-general-data.sql（データ連携基盤 サプライチェーンリレーションシップ - 一般データ）
* data-platform-supply-chain-relationship-sql-delivery-relation-data.sql（データ連携基盤 サプライチェーンリレーションシップ - 入出荷関係データ）
* data-platform-supply-chain-relationship-sql-billing-relation-data.sql（データ連携基盤 サプライチェーンリレーションシップ - 請求関係データ）
* data-platform-supply-chain-relationship-sql-payment-relation-data.sql（データ連携基盤 サプライチェーンリレーションシップ - 支払関係データ）
* data-platform-supply-chain-relationship-sql-delivery-plant-relation-data.sql（データ連携基盤 サプライチェーンリレーションシップ - プラント入出荷関係データ）
* data-platform-supply-chain-relationship-sql-production-plant-relation-data.sql（データ連携基盤 サプライチェーンリレーションシップ - プラント入出荷関係データ）

## caller.go による存在性確認
Input で取得されたファイルに基づいて、caller.go で、 API がコールされます。
caller.go の 以下の箇所が、指定された API をコールするソースコードです。

```
func (e *ExistenceConf) Conf(msg rabbitmq.RabbitmqMessage) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	_, ok := input["GlobalRegion"]
	if ok {
		input := &dpfm_api_input_reader.SDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confGlobalRegion(input)
		goto endProcess
	}

	err = xerrors.Errorf("can not get exconf check target")
endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

```

## Input
data-platform-api-supply-chain-relationship-exconf-rmq-kube では、以下のInputファイルをRabbitMQからJSON形式で受け取ります。  

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"service_label": "ORDERS",
	"ProductMasterGeneral": {
		"Product": "A012"
	},
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
}
```

## Output
data-platform-api-supply-chain-relationship-exconf-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。グローバル地域の対象値が存在する場合 true、存在しない場合 false、を返します。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
XXX
```