package instrumentation

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func InstrumentHandler(pattern string, handler http.Handler) http.Handler {
	return otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		RequestCnt.Add(ctx, 1, metric.WithAttributes(attribute.String("http.route", pattern)))
		handler.ServeHTTP(w, r)
	}), pattern)
}
