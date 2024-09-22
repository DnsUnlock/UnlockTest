package main

import (
	"fmt"
	"os"
	"plugin"
)

type PluginInterface interface {
	Registers() string                               // 用于存储插件系统的注册函数名
	Func() map[string]string                         // 用于存储函数名
	Call(Function string, args []interface{}) string // 用于调用函数
}
type PluginEntranceModel struct {
}

func main() {
	//加载插件
	open, err := plugin.Open("./unlock_test.so")
	if err != nil {
		fmt.Println(err)
		return
	}
	//通过插件名获取插件中的函数
	symbol, err := open.Lookup("PluginEntrance")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 转换为 PluginInterface 接口
	var pluginInstance PluginInterface
	pluginInstance, ok := symbol.(PluginInterface)
	if !ok {
		fmt.Println("Unexpected type from module symbol")
		return
	}

	// 调用插件中的函数
	for k, _ := range pluginInstance.Func() {
		var args []interface{}
		for i, v := range os.Args {
			if i == 0 {
				continue
			}
			args = append(args, v)
		}
		result := pluginInstance.Call(k, args)
		/*
			type Result struct {
				Status int
				Region string
				Info   string
				Err    error
			}
		*/
		fmt.Println(k, ":", result)
	}

}