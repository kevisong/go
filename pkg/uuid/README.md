# uuid

UUID generator based on [https://github.com/sony/sonyflake](https://github.com/sony/sonyflake)

## Usage

```go
err := uuid.Init()
if err != nil {
    // error handling
}

id, err := uuid.Generate()
if err != nil {
    // error handling
}

fmt.Println(id)
```
