# data-platform-api-product-master-exconf-rmq-kube
data-platform-api-product-master-exconf-rmq-kube は、データ連携基盤において、API で 品目マスタの存在性チェックを行うためのマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、品目マスタの存在確認が行われます。

* data-platform-product-master-sql-general-data.sql（データ連携基盤 品目マスタ - 一般データ）
* data-platform-product-master-sql-business-partner-data.sql（データ連携基盤 品目マスタ - ビジネスパートナデータ）
* data-platform-product-master-sql-bp-plant-data.sql（データ連携基盤 品目マスタ - BPプラントデータ）
* data-platform-product-master-sql-mrp-area-data.sql（データ連携基盤 品目マスタ - MRPエリアデータ）
* data-platform-product-master-sql-storage-location-data.sql（データ連携基盤 品目マスタ - 保管場所データ）
* data-platform-product-master-sql-storage-bin-data.sql（データ連携基盤 品目マスタ - 棚番データ）
* data-platform-product-master-sql-production-data.sql（データ連携基盤 品目マスタ - 製造データ）
* data-platform-product-master-sql-quality-data.sql（データ連携基盤 品目マスタ - 品質データ）
* data-platform-product-master-sql-accounting-data.sql（データ連携基盤 品目マスタ - 会計データ）
* data-platform-product-master-sql-tax-data.sql（データ連携基盤 品目マスタ - 税データ）
* data-platform-product-master-sql-product-description-data.sql（データ連携基盤 品目マスタ - 品目説明データ）
* data-platform-product-master-sql-product-desc-by-bp-data.sql（データ連携基盤 品目マスタ - 品目説明ビジネスパートナデータ）

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
data-platform-api-product-master-exconf-rmq-kube では、以下のInputファイルをRabbitMQからJSON形式で受け取ります。  

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"service_label": "PRODUCT_MASTER",
	"ProductMasterGeneral": {
		"Product": "A012"
	},
	"api_schema": "DPFMProductMasterExconf",
	"accepter": ["General"],
	"deleted": false
}
```

## Output
data-platform-api-product-master-exconf-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。グローバル地域の対象値が存在する場合 true、存在しない場合 false、を返します。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
XXX
```