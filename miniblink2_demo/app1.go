package main

import (
	gm "gitee.com/aochulai/GoMiniblink"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	ws "gitee.com/aochulai/GoMiniblink/forms/windows"
)

func main() {
	//windows版本初始化
	cs.App = new(ws.Provider).Init()

	//创建一个窗体并设置基本属性
	frm := new(cs.Form).Init()
	frm.SetTitle("普通窗口")
	frm.SetSize(800, 500)

	//创建浏览器控件并设置基本属性
	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetSize(700, 400)

	//添加浏览器控件到窗体
	frm.AddChild(mb)
	//注册回调, EvLoad回调在窗体首次显示前触发
	frm.EvLoad["回调名称"] = func(s cs.GUI) {
		//加载网址
		mb.LoadUri("https://www.baidu.com")
	}
	//将frm作为主窗口打开
	cs.Run(frm)
}
