package adapter

import "testing"

func TestAWSClientAdapter_CreateServer(t *testing.T) {
	var a CreateServer = &AWSClientAdapter{
		client: AWSClient{},
	}
	a.CreateServer(8, 8)
}

func TestTencentCloudClientAdapter_CreateServer(t *testing.T) {
	var a CreateServer = &TencentCloudClientAdapter{
		client: TencentCloudClient{},
	}
	a.CreateServer(8, 8)
}
