package bridge

type ISender interface {
	Send(msg string) error
}

type EmailSender struct{}

func (e *EmailSender) Send(msg string) error {
	return nil
}

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender ISender
}

func (e *ErrorNotification) Notify(msg string) error {
	e.sender.Send(msg)
	return nil
}

func Run() {
	errNotification := ErrorNotification{sender: &EmailSender{}}
	errNotification.Notify("hi")
}
