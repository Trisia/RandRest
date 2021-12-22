package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"time"
)

func RunTimer() {
	sitDuration := time.Minute * 35
	standDuration := time.Minute * 5
	for {
		time.Sleep(sitDuration)
		fmt.Println("请站起来吧")

		_ = tray.ShowCustom("请站起来吧", "你已经久坐了35分钟", ico)
		walk.MsgBox(wmain, "请站起来吧", "你已经久坐了35分钟", win.MB_SYSTEMMODAL|walk.MsgBoxOK|walk.MsgBoxIconWarning)
		time.Sleep(standDuration)
		_ = tray.ShowCustom("请坐下", "你已经站立了5分钟", ico)
		walk.MsgBox(wmain, "请坐下", "你已经站立了5分钟", win.MB_SYSTEMMODAL|walk.MsgBoxOK|walk.MsgBoxIconWarning)
	}
}
