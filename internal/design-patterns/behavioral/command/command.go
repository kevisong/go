package command

import "fmt"

// the command pattern is a behavioral design pattern in which an
// object is used to encapsulate all information needed to perform
// an action or trigger an event at a later time.

// StorageCommand is the command
type StorageCommand interface {
	Execute()
}

// standardCommand is the command
type standardCommand struct {
	StorageDevice
}

func (s standardCommand) Execute() {
	s.StandardStore()
}

// encryptedCommand is the command
type encryptedCommand struct {
	StorageDevice
}

func (e encryptedCommand) Execute() {
	e.EncryptedStore()
	fmt.Println("Encrypted Execute")
}

// StorageDevice is the receiver interface
type StorageDevice interface {
	StandardStore()
	EncryptedStore()
}

// HDD is the receiver
type HDD struct{}

func (h HDD) StandardStore() {
	fmt.Println("Standard store")
}

func (h HDD) EncryptedStore() {
	fmt.Println("Encrypted store")
}

type Invoker struct {
	StorageCommand
}

func (i Invoker) Execute() {
	i.Execute()
}

func Run() {
	standardCommand := standardCommand{HDD{}}
	encryptedCommand := encryptedCommand{HDD{}}

	invoker := Invoker{standardCommand}
	invoker.Execute()
	invoker.StorageCommand = encryptedCommand
	invoker.Execute()
}
