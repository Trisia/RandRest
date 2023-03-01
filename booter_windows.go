package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

const ApplicationName = "RandRest"

// BootSwitch 切换自启动状态
// status: true - 启动；false - 禁用
func BootSwitch(state bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("启动项查询失败,%v", err)
	}

	if state {
		// 获取当前执行文件绝对路径
		exePath, err := os.Executable()
		if err != nil {
			return err
		}
		exeLoc, _ := filepath.EvalSymlinks(exePath)
		err = k.SetStringValue(ApplicationName, exeLoc)
		if err != nil {
			return fmt.Errorf("无法设置开机启动项, %v", err)
		}
	} else {
		_ = k.DeleteValue(ApplicationName)
	}
	return nil
}

// BootState 查询开机启动状态
func BootState() (bool, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE)
	if err != nil {
		return false, fmt.Errorf("启动项查询失败, %v", err)
	}
	defer k.Close()
	if _, _, err = k.GetStringValue(ApplicationName); err != nil {
		if registry.ErrNotExist == err {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
