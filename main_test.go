package main

import (
	"fmt"
	wenxin "gotalk/pkg/ErnieBot-Turbo"
	"testing"
)

func TestMain(t *testing.T) {
	res1 := wenxin.DoChat("hello")
	fmt.Println("对话结果:", res1)
	res2 := wenxin.DoChat("介绍一下你自己")
	fmt.Println("对话结果:", res2)
}
