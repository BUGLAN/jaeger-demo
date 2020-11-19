## jaeger demo

> jaeger 使用示例

## jaeger 服务搭建

使用docker-compose搭建

```yaml
version: '3'

services:
  jaeger:
    container_name: jaegert
    image: jaegertracing/all-in-one:latest
    ports:
      - 5775:5775/udp
      - 6831:6831/udp
      - 6832:6832/udp
      - 5778:5778
      - 16686:16686
      - 14268:14268
      - 14250:14250
      - 9411:9411
```

启动

```yaml
docker-compose up
```

## 介绍

该项目为jaeger demo

#### hello world

```go
package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
    "time"
)

func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(errors.Wrap(err, "jaeger set config fail"))
	}
	return tracer, closer
}

func main() {
	tracer, closer := InitJaeger("demo-with-multi")
	defer closer.Close()

	opentracing.InitGlobalTracer(tracer)

	helloStr := "hello jaeger"
	span := tracer.StartSpan("say-hello")
	span.SetTag("value", helloStr)
	span.LogFields(
		log.String("event", "sayhello"),
		log.String("value", helloStr),
	)
	time.Sleep(time.Millisecond * 2)
	fmt.Println(helloStr)
	
	span.Finish()
}
```

#### 带链路的示例

见 [tracer](./tracer/main.go)

### 全链路追踪


* user_a 服务为http层, 为接口层
* b 服务为grpc层, 为逻辑层
* c 服务为数据操纵层, 为数据库层