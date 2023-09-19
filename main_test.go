package main

import (
	"fmt"
	wenxin "gotalk/pkg/ErnieBot-Turbo"
	"testing"
)

func TestMain(t *testing.T) {
	wenxin.GetToken()
	fmt.Println(wenxin.TokenBodys.AccessToken)
}
