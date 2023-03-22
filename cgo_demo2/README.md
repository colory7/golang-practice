使用如下命令更改可执行程序LDFLAGS参数 ,LDFLAGS 查找库文件路径

install_name_tool -change libhi.so /Users/mac/save/code/golang/golang-practice/cgo_demo2/lib/libhi.so  ./main

语法: install_name_tool -change 可执行程序中的.so路径  实际的.so路径  可执行程序

查看连接库查找路径

otool -L ./main