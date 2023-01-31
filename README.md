# data-platform-api-product-master-exconf-rmq-kube
data-platform-api-product-master-exconf-rmq-kube は、データ連携基盤において、API で 品目マスタの存在性チェックを行うためのマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、ビジネスパートナの存在確認が行われます。

* data-platform-product-master-sql-general-data.sql（データ連携基盤 品目マスタ - 一般データ）
* data-platform-product-master-sql-business-partner-data.sql（データ連携基盤 品目マスタ - ビジネスパートナデータ）
* data-platform-product-master-sql-bp-plant-data.sql（データ連携基盤 品目マスタ - BPプラントデータ）
* data-platform-product-master-sql-storage-location-data.sql（データ連携基盤 品目マスタ - 保管場所データ）

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

	_, ok := input["ProductMasterGeneral"]
	if ok {
		input := &dpfm_api_input_reader.GeneralSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterGeneral(input)
		goto endProcess
	}
	_, ok = input["ProductMasterBusinessPartner"]
	if ok {
		input := &dpfm_api_input_reader.BusinessPartnerSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterBusinessPartner(input)
		goto endProcess
	}
	_, ok = input["ProductMasterBPPlant"]
	if ok {
		input := &dpfm_api_input_reader.BPPlantSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterBPPlant(input)
		goto endProcess
	}
	_, ok = input["ProductMasterStorageLocation"]
	if ok {
		input := &dpfm_api_input_reader.StorageLocationSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confProductMasterStorageLocation(input)
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
data-platform-api-product-master-exconf-rmq-kube では、以下のInputファイルをRabbitMQからJSON形式で受け取ります。  

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
data-platform-api-product-master-exconf-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。品目マスタの対象値が存在する場合 true、存在しない場合 false、を返します。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
{
	"cursor": "/Users/latona2/bitbucket/data-platform-api-product-master-exconf-rmq-kube/main.go#L69",
	"function": "main.dataCallProcess",
	"level": "INFO",
	"message": {
		"ProductMasterGeneral": {
			"Product": "A3750",
			"ExistenceConf": true
		}
	},
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"time": "2022-11-14T23:18:48+09:00"
}
```