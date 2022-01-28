package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-measurement-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		header = append(header, Header{
	MeasurementDocument:            data.MeasurementDocument,
	MeasuringPoint:                 data.MeasuringPoint,
	MeasuringPointPositionNumber:   data.MeasuringPointPositionNumber,
	MsmtRdngDate:                   data.MsmtRdngDate,
	MsmtRdngTime:                   data.MsmtRdngTime,
	Characteristic:                 data.Characteristic,
	MsmtDocumentReferredOrder:      data.MsmtDocumentReferredOrder,
	RefdMaintOrderOpStatusObject:   data.RefdMaintOrderOpStatusObject,
	MaintenanceOrderOperation:      data.MaintenanceOrderOperation,
	MaintenanceOrderSubOperation:   data.MaintenanceOrderSubOperation,
	MsmtIsDoneAfterTaskCompltn:     data.MsmtIsDoneAfterTaskCompltn,
	CharcValueUnit:                 data.CharcValueUnit,
	MeasurementReading:             data.MeasurementReading,
	MeasurementReadingInEntryUoM:   data.MeasurementReadingInEntryUoM,
	MeasurementReadingEntryUoM:     data.MeasurementReadingEntryUoM,
	MeasurementCounterReading:      data.MeasurementCounterReading,
	MsmtCounterReadingDifference:   data.MsmtCounterReadingDifference,
	TotalMsmtRdngIsSetExternally:   data.TotalMsmtRdngIsSetExternally,
	MeasuringPointTargetValue:      data.MeasuringPointTargetValue,
	MsmtValuationCode:              data.MsmtValuationCode,
	MeasurementDocumentText:        data.MeasurementDocumentText,
	MeasurementDocumentHasLongText: data.MeasurementDocumentHasLongText,
	MsmtRdngByUser:                 data.MsmtRdngByUser,
	MsmtRdngStatus:                 data.MsmtRdngStatus,
	MsmtCntrReadingDiffIsEntered:   data.MsmtCntrReadingDiffIsEntered,
	MsmtRdngIsReversed:             data.MsmtRdngIsReversed,
	MsmtCounterReadingIsReplaced:   data.MsmtCounterReadingIsReplaced,
		})
	}

	return header, nil
}
