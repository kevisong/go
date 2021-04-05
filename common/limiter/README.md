# limiter

## Usage

```go
lmtr := limiter.NewRateLimiter(20, time.Second) // 20/s
defer lmtr.Stop()
for {
    <-lmtr.C
    // Do somthing
}
```
