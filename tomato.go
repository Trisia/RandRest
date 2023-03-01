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

const (
	StateWorking = 0 // 休息状态
	StateRest    = 1 // 站立状态
)

var (
	counter int32 // 计时器
	state   int32 // 状态
)

var hitBtn *walk.Action

// Work 重置计数器
func Work() {
	atomic.SwapInt32(&counter, 0)
	atomic.SwapInt32(&state, StateWorking)
	_ = tray.SetToolTip(fmt.Sprintf("工作: %d/%d min", counter, sitDuration))
	_ = hitBtn.SetText(fmt.Sprintf("工作: %d/%d min", counter, sitDuration))
}

// Rest 休息
func Rest() {
	atomic.SwapInt32(&counter, 0)
	atomic.SwapInt32(&state, StateRest)
	_ = tray.SetToolTip(fmt.Sprintf("休息: %d/%d min", counter, standDuration))
	_ = hitBtn.SetText(fmt.Sprintf("休息: %d/%d min", counter, standDuration))
}

// RunTimer 启动计时器
func RunTimer() {
	Work()
	for {
		time.Sleep(time.Minute)
		current := atomic.AddInt32(&counter, 1)

		switch atomic.LoadInt32(&state) {
		case StateWorking:
			_ = hitBtn.SetText(fmt.Sprintf("工作: %d/%d min", counter, sitDuration))
			if current < sitDuration {
				_ = tray.SetToolTip(fmt.Sprintf("工作: %d/%d min", counter, sitDuration))
			} else {
				_ = tray.ShowCustom("请站起来吧", "你已经久坐了工作很久了", ico)
				walk.MsgBox(wmain, "请站起来吧", "你已经久坐了工作很久了", win.MB_SYSTEMMODAL|walk.MsgBoxOK|walk.MsgBoxIconWarning)
				Work()
			}
		case StateRest:
			_ = hitBtn.SetToolTip(fmt.Sprintf("休息: %d/%d min", current, standDuration))
			if current < standDuration {
				_ = tray.SetToolTip(fmt.Sprintf("休息: %d/%d min", current, standDuration))
			} else {
				_ = tray.ShowCustom("请接续工作吧", "你已经休息了一段时间了", ico)
				walk.MsgBox(wmain, "请接续工作把", "你已经休息了一段时间了", win.MB_SYSTEMMODAL|walk.MsgBoxOK|walk.MsgBoxIconWarning)
				Rest()
			}
		}
	}
}
