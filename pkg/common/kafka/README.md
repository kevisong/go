# kafka

```go
go kafka.StartConsumer("localhost:9092", "myTopic", 0, -2, 1)
for {
    select {
    case msg := <-kafka.MsgChan:
        fmt.Println(msg)
    case err := <-kafka.ErrChan:
        logrus.Error(err)
    }
}
```
