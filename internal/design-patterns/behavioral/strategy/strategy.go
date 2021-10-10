package strategy

import "fmt"

// Strategy pattern enables selecting an algorithm at runtime. Instead of
// implementing a single algorithm directly, code receives run-time
// instructions as to which in a family of algorithms to use.

type StorageStrategy interface {
	Save()
}

type standardStorage struct{}

func (s *standardStorage) Save() {
	fmt.Println("Standard save")
}

type encryptedStorage struct{}

func (e *encryptedStorage) Save() {
	fmt.Println("Encrypted save")
}

// Device is the context
type Device struct {
	StorageStrategy
}

func Run() {
	device := Device{&standardStorage{}}
	device.Save()
	device.StorageStrategy = &encryptedStorage{}
	device.Save()
}
