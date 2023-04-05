package trace

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func Start(c context.Context) (context.Context, trace.Span) {
	pc, _, _, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	list := strings.Split(details.Name(), "/")
	adapter := list[len(list)-3]
	specific := list[len(list)-2]
	names := strings.Split(list[len(list)-1], ".")
	name := list[len(list)-1]
	switch adapter {
	case "transport":
		name = fmt.Sprintf("%s.%s.%s", specific, names[0], names[2])

	case "usecase":
		name = fmt.Sprintf("usecase.%s.%s", specific, names[2])

	case "service":
		name = fmt.Sprintf("%s.%s.%s", specific, names[0], names[2])

	case "module":
		name = fmt.Sprintf("module.%s.%s", specific, names[len(names)-1])
	}

	// use default name
	ctx, span := otel.GetTracerProvider().Tracer(adapter).Start(c, name)

	if span.IsRecording() {
		span.SetAttributes(attribute.String("system.func", details.Name()))
	}

	return ctx, span
}

func RecordError(span trace.Span, err error) {
	if err != nil && span.IsRecording() {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
