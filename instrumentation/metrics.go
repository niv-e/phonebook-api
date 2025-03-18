package instrumentation

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const meterName = "phonebook-api"

var (
	meter      = otel.Meter(meterName)
	RequestCnt metric.Int64Counter
)

func InitMetrics() {
	var err error
	RequestCnt, err = meter.Int64Counter("http.requests",
		metric.WithDescription("The number of HTTP requests"),
		metric.WithUnit("{request}"))
	if err != nil {
		panic(err)
	}
}
