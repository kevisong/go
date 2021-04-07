# limiter

Based on [https://github.com/influxdata/telegraf/blob/master/internal/limiter/limiter.go](https://github.com/influxdata/telegraf/blob/master/internal/limiter/limiter.go)

## Usage

```go
lmtr := limiter.NewRateLimiter(20, time.Second) // 20/s
defer lmtr.Stop()
for {
    <-lmtr.C
    // Do somthing
}
```
