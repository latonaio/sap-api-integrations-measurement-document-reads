package main

import (
	sap_api_caller "sap-api-integrations-measurement-document-reads/SAP_API_Caller"
	"sap-api-integrations-measurement-document-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Measurement_Document_Header_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header",
		}
	}

	caller.AsyncGetMeasurementDocument(
		inoutSDC.MeasurementDocument.MeasurementDocument,
		accepter,
	)
}
