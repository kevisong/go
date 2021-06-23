# http

A simple http requester

## Usage

```go
resp, err := http.Get("url")
if err != nil {
    // err handling
}
fmt.Println(string(resp))
```

With headers

```go
resp, err := http.SetHeaders(map[string]string{}).Post("url", []byte{})
if err != nil {
    // err handling
}
fmt.Println(string(resp))
```
