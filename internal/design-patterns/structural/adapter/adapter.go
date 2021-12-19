package adapter

import "fmt"

type AWSClient struct{}

func (a *AWSClient) RunInstance(cpu, mem int) error {
	fmt.Println("aws client run")
	return nil
}

type TencentCloudClient struct{}

func (t *TencentCloudClient) RunInstance(cpu, mem int) error {
	fmt.Println("tencent cloud client run")
	return nil
}

type CreateServer interface {
	CreateServer(cpu, mem int) error
}

// AWSClientAdapter is the adatper
type AWSClientAdapter struct {
	client AWSClient
}

// CreateServer implements CreateServer interface
func (a *AWSClientAdapter) CreateServer(cpu, mem int) error {
	return a.client.RunInstance(cpu, mem)
}

// TencentCloudClientAdapter is the adatper
type TencentCloudClientAdapter struct {
	client TencentCloudClient
}

// CreateServer implements CreateServer interface
func (t *TencentCloudClientAdapter) CreateServer(cpu, mem int) error {
	return t.client.RunInstance(cpu, mem)
}
