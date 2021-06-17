# log

log initialization lib for [logrus](https://github.com/sirupsen/logrus).

```go
log.Init(&log.Config{Level: "info"})
logrus.WithFields(logrus.Fields{
    "api": "/user",
}).Info("success")
```
