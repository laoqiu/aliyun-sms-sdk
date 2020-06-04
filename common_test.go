package sms

import (
	"fmt"
	"testing"
)

func TestCommonRequest(t *testing.T) {
	req := CommonRequest{}
	req.SetDefaultParam()
	if req.Version != DefaultCommonRequest["Version"] {
		t.Error("未赋值")
	}
	fmt.Println(req)
}
