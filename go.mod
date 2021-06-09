module github.com/tsundata/assistant

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/RichardKnop/machinery/v2 v2.0.10
	github.com/go-ping/ping v0.0.0-20210216210419-25d1413fb7bb
	github.com/go-redis/redis/v8 v8.9.0
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gofiber/fiber/v2 v2.10.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/influxdata/cron v0.0.0-20201006132531-4bb0a200dcbe
	github.com/influxdata/influxdb-client-go/v2 v2.3.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/nats-io/nats-server/v2 v2.1.6 // indirect
	github.com/nats-io/nats.go v1.11.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rollbar/rollbar-go v1.4.0
	github.com/shirou/gopsutil/v3 v3.21.4
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/slack-go/slack v0.9.1
	github.com/sourcegraph/checkup v1.0.1-0.20200721114922-77e7567835d4
	github.com/spaolacci/murmur3 v1.1.0
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.8.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/valyala/fasthttp v1.25.0
	github.com/yuin/goldmark v1.3.7
	go.uber.org/zap v1.17.0
	google.golang.org/grpc v1.37.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
