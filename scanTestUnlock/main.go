package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 定义要扫描的目录
	dir := "/home/code/GoProjects/DnsUnlock/UnlockTest/testUnlock"

	// 初始化文件-函数名map
	fileFuncMap := make(map[string][]string)

	// 遍历目录中的所有Go文件
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			funcName, err := extractUnusedFuncNames(path)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			fileFuncMap[path] = funcName
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
		return
	}

	// 打印文件-函数名map
	for file, funcName := range fileFuncMap {
		//提取文件名
		fileName := filepath.Base(file)
		//去除后缀
		fileName = strings.ReplaceAll(fileName, ".go", "")
		//"4GTV":     "TW4GTV",
		for _, v := range funcName {
			//fmt.Printf("\"%s\":\"%s\",\n", fileName, v)
			//	case "Dazn":
			//		return TS(testUnlock.Dazn(c))
			fmt.Printf("case \"%s\":\n\treturn TS(testUnlock.%s(c))\n", fileName, v)
		}

	}
}

// 提取文件中的未调用函数名
func extractUnusedFuncNames(filePath string) ([]string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	funcDecls := make(map[string]*ast.FuncDecl)
	calledFuncs := make(map[string]bool)

	// 遍历所有的函数声明
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			funcDecls[x.Name.Name] = x
		case *ast.CallExpr:
			if ident, ok := x.Fun.(*ast.Ident); ok {
				calledFuncs[ident.Name] = true
			}
		}
		return true
	})

	// 找到未被调用的函数
	var unusedFuncs []string
	for funcName := range funcDecls {
		if !calledFuncs[funcName] {
			unusedFuncs = append(unusedFuncs, funcName)
		}
	}

	return unusedFuncs, nil
}
