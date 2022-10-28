# data-platform-api-product-master-exconf-rmq-kube
data-platform-api-product-master-exconf-rmq-kube は、データ連携基盤において、品目マスタの存在性チェックを行うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 存在確認先テーブル名
以下のsqlファイルに対して、品目マスタの存在確認が行われます。

* data-platform-product-master-general-data.sql（データ連携基盤 品目マスタ - 一般データ）

## existence_check.go による存在性確認
Input で取得されたファイルに基づいて、existence_check.go で、database 内の models 内の data_platform_product_master_general_data.go から品目の対象値の存在性確認が行われます。

```
func (e *ExistencyChecker) Check(data map[string]interface{}) map[string]interface{} {
	existData := map[string]interface{}{
		"ExistenceConf": false,
	}
	if data == nil {
		return existData
	}
	wg := sync.WaitGroup{}

	rawOrderData, ok := data["Orders"] // 変更箇所
	if !ok {
		return existData
	}

	orderData, ok := rawOrderData.(map[string]interface{})
	if !ok {
		return existData
	}

	existData["ExistenceConf"] = true

	for k := range orderData {
		v := fmt.Sprintf("%v", orderData[k])
		notExistKeys := make([]string, 0, len(orderData))
		switch k {
		case "Product": // 変更箇所
			wg.Add(1)
			existData[k] = v
			go func() {
				if !e.checkProduct(v) {
					existData["ExistenceConf"] = false
					notExistKeys = append(notExistKeys, k)
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return existData
}
```

## Input
data-platform-api-product-master-exconf-rmq-kube は、入力ファイルとして、RabbitMQ からのメッセージを JSON 形式で受け取ります。入力ファイルは、Inputsフォルダ内にあります。

```
{
	"connection_key": "response",
	"result": true,
	"redis_key": "abcdefg",
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"service_label": "ORDERS",
	"Product": {
		"Product": "21"
	},
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
}
```


## Output
data-platform-api-product-master-exconf-rmq-kube では、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。ビジネスパートナの対象値が存在する場合 true、存在しない場合 false、を返します。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。

```
{
    "cursor": "/go/src/github.com/latonaio/data-platform-product-master-exconf/main.go#L56",
    "function": "('DPFM_API_ORDERS_SRV', 'creates', 'A_Item', 'Product', 'DPFM_API_PRODUCT_MASTER_SRV', 'exconf');",
    "level": "INFO",
    "service_label": "ORDERS",
    "Product": {
        "Product": "21",
        "ExistenceConf": true
    },
    "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
    "time": "2022-10-15T17:58:07+09:00"
}
```