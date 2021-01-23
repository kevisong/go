# wecom

```go
res, err := wecom.Send("webhook", "hi")
if err != nil {
    logrus.Error(err)
}
logrus.Info(string(res))
```
