# log

Log initialization pkg for [logrus](https://github.com/sirupsen/logrus).

```go
log.Init(&log.Config{Level: "info", Formatter: "json"})
logrus.WithFields(logrus.Fields{
    "api": "/user",
}).Info("success")
```
