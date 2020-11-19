package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"jaeger-demo/helper"
)

func main() {
	tracer, closer := helper.InitJaeger("demo-with-multi")
	defer closer.Close()

	opentracing.InitGlobalTracer(tracer)

	ctx := Service1()

	Service2(ctx)
}

func Service1() context.Context {
	ctx := context.Background()
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("service-1")
	span.Finish()
	msg := "service-1"
	fmt.Println(msg)
	span.SetTag("service-name", 1)
	span.LogFields(
		log.String("event", "service1"),
	)
	span.SetBaggageItem("baggage", "service1")

	return opentracing.ContextWithSpan(ctx, span)
}

func Service2(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service-2")
	defer span.Finish()
	span.SetTag("service-name", 2)
	span.LogFields(
		log.String("event", "service2"),
	)
	span.SetBaggageItem("baggage3", "service1")
}
