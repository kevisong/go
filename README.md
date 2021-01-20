# go

Go Libs

## http

A simple http requester pkg.

```go
// Quick start
resp, err := http.Get("url")
if err != nil {
    // err handling
}
fmt.Println(string(resp))

// With headers
resp, err := http.SetHeaders(map[string]string{}).Post("url", []byte{})
if err != nil {
    // err handling
}
fmt.Println(string(resp))
```

## log

Log initialization pkg using Logrus.

```go
log.Init(&log.Config{Level: "info", Formatter: "json"})
```
