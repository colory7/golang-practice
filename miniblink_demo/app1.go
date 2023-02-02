package main

import (
	"github.com/del-xiong/miniblink"
	"log"
)

func main() {
	//设置调试模式
	miniblink.SetDebugMode(true)
	//初始化miniblink模块
	err := miniblink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}
	// 启动1366x920普通浏览器
	view := miniblink.NewWebView(false, 1366, 920)
	// 启动1366x920透明浏览器(只有web界面会显示)
	//view := miniblink.NewWebView(true, 1366, 920)
	// 加载github
	view.LoadURL("https://github.com/del-xiong/miniblink")
	// 设置窗体标题(会被web页面标题覆盖)
	view.SetWindowTitle("miniblink window")
	// 移动到屏幕中心位置
	view.MoveToCenter()
	// 显示窗口
	view.ShowWindow()
	// 开启调试模式(会调起chrome调试页面)
	view.ShowDevTools()
	<-make(chan bool)
}
