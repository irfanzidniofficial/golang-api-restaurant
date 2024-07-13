package tracking

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

func createTrackingProvider(url string) (*tracesdk.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-resto"),
			attribute.String("enviroment", "staging"),
		)),
	)
	return tp, nil
}

func Init(url string) error {
	tp, err := createTrackingProvider(url)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tp)
	return nil
}

func CreateSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if ctx==nil{
		ctx= context.Background()
	}
	tr:=otel.Tracer(name)
	ctx, span:=tr.Start(ctx, name)
	return ctx, span

}
