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
		//zap.L().Error("无法加载图标", zap.Error(err))
		//err = nil
	} else {
		wmain.AddDisposable(ico)
		_ = tray.SetIcon(ico)
	}
	resetBtn := walk.NewAction()
	_ = resetBtn.SetText("重置")
	resetBtn.Triggered().Attach(Reset)
	quitBtn := walk.NewAction()
	_ = quitBtn.SetText("退出")
	quitBtn.Triggered().Attach(func() {
		walk.App().Exit(0)
	})

	actions := tray.ContextMenu().Actions()
	_ = actions.Add(resetBtn)
	_ = actions.Add(quitBtn)
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
