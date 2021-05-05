package nvosrestgo

import (
	"fmt"
	"testing"
)

func TestNVOSRestGo(t *testing.T) {
	var IP = "10.36.16.14"
	var username = "admin"
	var password = "admin"
	var port = 8000

	visionOneClient := New(IP, username, password, port)

	//pull device type as a get test
	fmt.Println("NVOS Connected")
	fmt.Println("  --Trying device type")
	device_type, err := visionOneClient.GetSystemProperty("type")
	if err != nil {
		fmt.Println("NVOS Test FAILED")
		fmt.Println(err.Error())
		t.FailNow()
	}

	fmt.Println(device_type)

}
