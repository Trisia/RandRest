package main

import (
	"fmt"
	"github.com/lxn/walk"
	"time"
)

func RunTimer() {
	sitDuration := time.Minute * 35
	standDuration := time.Second * 5
	for {
		time.Sleep(sitDuration)
		fmt.Println("请站起来吧")
		walk.MsgBox(wmain, "请站起来吧", "你已经久坐了35分钟", walk.MsgBoxTopMost|walk.MsgBoxOK|walk.MsgBoxIconWarning)
		time.Sleep(standDuration)
		fmt.Println("请坐下")
		walk.MsgBox(wmain, "请坐下", "你已经站立了5分钟", walk.MsgBoxTopMost|walk.MsgBoxOK|walk.MsgBoxIconWarning)
	}
}
