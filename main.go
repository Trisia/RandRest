package main

import (
	"embed"
	"github.com/lxn/walk"
	"image"
	_ "image/png"
)

var (
	//go:embed icon.png
	icoFile embed.FS
	wmain   *walk.MainWindow
	ico     walk.Image // UI图标
	tray    *walk.NotifyIcon
)

func main() {
	// 开机启动
	if boot, _ := BootState(); !boot {
		_ = BootSwitch(true)
	}

	err := ui()
	if err != nil {
		panic(err)
	}
	wmain.Starting().Once(func() {
		go RunTimer()
	})
	// 主线程阻塞，启动UI
	wmain.Run()

	if tray != nil {
		_ = tray.Dispose()
	}
	wmain.Dispose()
}

func ui() error {
	var err error
	wmain, err = walk.NewMainWindowWithName("StandRest")
	if err != nil {
		return err
	}
	// 程序相对路径加载图标
	tray, err = walk.NewNotifyIcon(wmain)
	if err != nil {
		return err
	}
	_ = tray.SetVisible(true)

	ico, err = loadIcon()
	if err != nil {
		return err
	} else {
		wmain.AddDisposable(ico)
		_ = tray.SetIcon(ico)
	}
	restBtn := walk.NewAction()
	_ = restBtn.SetText("休息一下")
	restBtn.Triggered().Attach(Rest)

	workBtn := walk.NewAction()
	_ = workBtn.SetText("重新计时")
	restBtn.Triggered().Attach(Work)

	quitBtn := walk.NewAction()
	_ = quitBtn.SetText("退出")
	quitBtn.Triggered().Attach(func() {
		walk.App().Exit(0)
	})

	actions := tray.ContextMenu().Actions()
	_ = actions.Add(restBtn)
	_ = actions.Add(workBtn)
	_ = actions.Add(quitBtn)

	_ = tray.ShowMessage("RandRest", "时钟程序已经在后台运行。")
	return nil
}

// 从路径加载 图片对象
func loadIcon() (*walk.Icon, error) {
	reader, _ := icoFile.Open("icon.png")
	im, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	img, err := walk.NewIconFromImageForDPI(im, 96)
	if err != nil {
		return nil, err
	}
	return img, err
}
