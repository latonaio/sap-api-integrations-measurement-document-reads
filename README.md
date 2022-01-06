# sap-api-integrations-measurement-document-reads  
sap-api-integrations-measurement-document-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 計測伝票 データを取得するマイクロサービスです。  
sap-api-integrations-measurement-document-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-measurement-document-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_MEASUREMENTDOCUMENT_0001/overview

## 動作環境
sap-api-integrations-measurement-document-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-measurement-document-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-measurement-document-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_MEASUREMENTDOCUMENT_0001/overview
* APIサービス名(=baseURL): api_measurementdocument/srvd_a2x/sap/measurementdocument/0001/

## 本レポジトリ に 含まれる API名
sap-api-integrations-measurement-document-reads には、次の API をコールするためのリソースが含まれています。  

* MeasurementDocument（計測伝票 - ヘッダ）


## API への 値入力条件 の 初期値
sap-api-integrations-measurement-document-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## SDC レイアウト

* inoutSDC.MeasurementDocument.MeasurementDocument（計測伝票）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "sap.s4.beh.measurementdocument.v1.MeasurementDocument.Created.v1",
	"accepter": ["Header"],
	"measurement_document_no": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.measurementdocument.v1.MeasurementDocument.Created.v1",
	"accepter": ["All"],
	"measurement_document_no": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMeasurementDocument(measurementDocument string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(measurementDocument)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## SAP API Business Hub における API サービス の バージョン と バージョン におけるデータレイアウトの相違

SAP API Business Hub における API サービス のうちの 殆どの API サービス のBASE URLのフォーマットは、"API_(リポジトリ名)_SRV" であり、殆どの API サービス 間 の データレイアウトは統一されています。   
従って、Latona および AION における リソースにおいても、データレイアウトが統一されています。    
一方、本レポジトリ に関わる API である Measurement Document のサービスは、BASE URLのフォーマットが他のAPIサービスと異なります。      
その結果、本レポジトリ内の一部のAPIのデータレイアウトが、他のAPIサービスのものと異なっています。  

#### BASE URLが "API_(リポジトリ名)_SRV" のフォーマットである API サービス の データレイアウト（=responses）  
BASE URLが "API_{リポジトリ名}_SRV" のフォーマットであるAPIサービスのデータレイアウト（=responses）は、例えば、次の通りです。  
```
type ToProductionOrderItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ManufacturingOrder             string      `json:"ManufacturingOrder"`
			ManufacturingOrderItem         string      `json:"ManufacturingOrderItem"`
			ManufacturingOrderCategory     string      `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string      `json:"ManufacturingOrderType"`
			IsCompletelyDelivered          bool        `json:"IsCompletelyDelivered"`
			Material                       string      `json:"Material"`
			ProductionPlant                string      `json:"ProductionPlant"`
			Plant                          string      `json:"Plant"`
			MRPArea                        string      `json:"MRPArea"`
			QuantityDistributionKey        string      `json:"QuantityDistributionKey"`
			MaterialGoodsReceiptDuration   string      `json:"MaterialGoodsReceiptDuration"`
			StorageLocation                string      `json:"StorageLocation"`
			Batch                          string      `json:"Batch"`
			InventoryUsabilityCode         string      `json:"InventoryUsabilityCode"`
			GoodsRecipientName             string      `json:"GoodsRecipientName"`
			UnloadingPointName             string      `json:"UnloadingPointName"`
			MfgOrderItemPlndDeliveryDate   string      `json:"MfgOrderItemPlndDeliveryDate"`
			MfgOrderItemActualDeliveryDate string      `json:"MfgOrderItemActualDeliveryDate"`
			ProductionUnit                 string      `json:"ProductionUnit"`
			MfgOrderItemPlannedTotalQty    string      `json:"MfgOrderItemPlannedTotalQty"`
			MfgOrderItemPlannedScrapQty    string      `json:"MfgOrderItemPlannedScrapQty"`
			MfgOrderItemGoodsReceiptQty    string      `json:"MfgOrderItemGoodsReceiptQty"`
			MfgOrderItemActualDeviationQty string      `json:"MfgOrderItemActualDeviationQty"`
		} `json:"results"`
	} `json:"d"`
}

```

#### BASE URL が "api_measurementdocument/srvd_a2x/sap/measurementdocument/0001/" である Measurement Document の APIサービス の データレイアウト（=responses）  
BASE URL が "api_measurementdocument/srvd_a2x/sap/measurementdocument/0001/" である Measurement Document の APIサービス の データレイアウト（=responses）は、例えば、次の通りです。  

```
type Header struct {
	Value             []struct {
		MeasurementDocument            string        `json:"MeasurementDocument"`
		MeasuringPoint                 string        `json:"MeasuringPoint"`
		MeasuringPointPositionNumber   string        `json:"MeasuringPointPositionNumber"`
		MsmtRdngDate                   string        `json:"MsmtRdngDate"`
		MsmtRdngTime                   string        `json:"MsmtRdngTime"`
		Characteristic                 string        `json:"Characteristic"`
		MsmtDocumentReferredOrder      string        `json:"MsmtDocumentReferredOrder"`
		RefdMaintOrderOpStatusObject   string        `json:"RefdMaintOrderOpStatusObject"`
		MaintenanceOrderOperation      string        `json:"MaintenanceOrderOperation"`
		MaintenanceOrderSubOperation   string        `json:"MaintenanceOrderSubOperation"`
		MsmtIsDoneAfterTaskCompltn     bool          `json:"MsmtIsDoneAfterTaskCompltn"`
		CharcValueUnit                 string        `json:"CharcValueUnit"`
		MeasurementReading             float64       `json:"MeasurementReading"`
		MeasurementReadingInEntryUoM   float64       `json:"MeasurementReadingInEntryUoM"`
		MeasurementReadingEntryUoM     string        `json:"MeasurementReadingEntryUoM"`
		MeasurementCounterReading      float64       `json:"MeasurementCounterReading"`
		MsmtCounterReadingDifference   float64       `json:"MsmtCounterReadingDifference"`
		TotalMsmtRdngIsSetExternally   bool          `json:"TotalMsmtRdngIsSetExternally"`
		MeasuringPointTargetValue      int           `json:"MeasuringPointTargetValue"`
		MsmtValuationCode              string        `json:"MsmtValuationCode"`
		MeasurementDocumentText        string        `json:"MeasurementDocumentText"`
		MeasurementDocumentHasLongText bool          `json:"MeasurementDocumentHasLongText"`
		MsmtRdngByUser                 string        `json:"MsmtRdngByUser"`
		MsmtRdngStatus                 string        `json:"MsmtRdngStatus"`
		MsmtCntrReadingDiffIsEntered   bool          `json:"MsmtCntrReadingDiffIsEntered"`
		MsmtRdngIsReversed             bool          `json:"MsmtRdngIsReversed"`
		MsmtCounterReadingIsReplaced   bool          `json:"MsmtCounterReadingIsReplaced"`
	} `json:"value"`
}
```
このように、BASE URLが "API_(リポジトリ名)_SRV" のフォーマットである API サービス の データレイアウトと、 Measurement Document の データレイアウトは、D、Results、Metadata、Value の配列構造を持っているか持っていないかという点が異なります。  

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 計測伝票 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"MeasurementDocument" ～ "MsmtCounterReadingIsReplaced" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-measurement-document-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-measurement-document-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"MeasurementDocument": "1",
			"MeasuringPoint": "3",
			"MeasuringPointPositionNumber": "",
			"MsmtRdngDate": "2020-11-02",
			"MsmtRdngTime": "16:04:23",
			"Characteristic": "ODOMTER_RDNG",
			"MsmtDocumentReferredOrder": "",
			"RefdMaintOrderOpStatusObject": "",
			"MaintenanceOrderOperation": "",
			"MaintenanceOrderSubOperation": "",
			"MsmtIsDoneAfterTaskCompltn": false,
			"CharcValueUnit": "KM",
			"MeasurementReading": 600000,
			"MeasurementReadingInEntryUoM": 600,
			"MeasurementReadingEntryUoM": "KM",
			"MeasurementCounterReading": 600000,
			"MsmtCounterReadingDifference": 600000,
			"TotalMsmtRdngIsSetExternally": false,
			"MeasuringPointTargetValue": 0,
			"MsmtValuationCode": "",
			"MeasurementDocumentText": "",
			"MeasurementDocumentHasLongText": false,
			"MsmtRdngByUser": "CB9980002113",
			"MsmtRdngStatus": "",
			"MsmtCntrReadingDiffIsEntered": false,
			"MsmtRdngIsReversed": false,
			"MsmtCounterReadingIsReplaced": false
		}
	],
	"time": "2022-01-06T15:03:03.807815+09:00"
}
```
