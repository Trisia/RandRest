package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"sync/atomic"
	"time"
)

const (
	sitDuration   int32 = 35 // 静坐时间
	standDuration int32 = 5  // 站立时间
)

var (
	duration int32
)

// Reset 重置计数器
func Reset() {
	atomic.SwapInt32(&duration, 0)
	_ = tray.SetToolTip(fmt.Sprintf("静坐: %d/%d min", duration, sitDuration))
}

// RunTimer 启动计时器
func RunTimer() {
	_ = tray.SetToolTip(fmt.Sprintf("静坐: %d/%d min", duration, sitDuration))
	for {
		time.Sleep(time.Minute)
		d := atomic.AddInt32(&duration, 1)
		// 更新静坐时间
		if d < sitDuration {
			_ = tray.SetToolTip(fmt.Sprintf("静坐: %d/%d min", duration, sitDuration))
			continue
		}

		if sitDuration == d {
			_ = tray.ShowCustom("请站起来吧", "你已经久坐了35分钟", ico)
			walk.MsgBox(wmain, "请站起来吧", "你已经久坐了35分钟", win.MB_SYSTEMMODAL|walk.MsgBoxOK|walk.MsgBoxIconWarning)
		}

		// 更新站立时间
		if d >= sitDuration {
			_ = tray.SetToolTip(fmt.Sprintf("站立: %d/%d min", d-sitDuration, standDuration))
		}

		if d >= (sitDuration + standDuration) {
			_ = tray.ShowCustom("请坐下", "你已经站立了5分钟", ico)
			walk.MsgBox(wmain, "请坐下", "你已经站立了5分钟", win.MB_SYSTEMMODAL|walk.MsgBoxOK|walk.MsgBoxIconWarning)
			Reset()
		}
	}
}
