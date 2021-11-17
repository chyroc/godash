package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chyroc/godash/internal/generate/funcs"
)

func main() {
	if len(os.Args) != 4 {
		panic(fmt.Sprintf("长度不为 4"))
	}
	file := os.Args[1] + ".go"
	funcname := os.Args[2]
	typeList := getSlice(os.Args[3])

	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	funcTmpl := funcs.Funcs[funcname]
	if funcTmpl == nil {
		panic(fmt.Sprintf("不存在函数 %s", funcname))
	}

	bodyList := []string{
		getHead(string(bs)),
	}
	for _, v := range typeList {
		body := funcTmpl(v)
		bodyList = append(bodyList, body)
	}

	body := strings.Join(bodyList, "\n\n")
	err = ioutil.WriteFile(file, []byte(body), 0666)
	if err != nil {
		panic(err)
	}
}

func getSlice(s string) []string {
	res := []string{}
	for _, v := range strings.Split(s, ",") {
		res = append(res, strings.TrimSpace(v))
	}
	return res
}

func getHead(s string) string {
	res := []string{}
	for _, v := range strings.Split(s, "\n") {
		if strings.HasPrefix(v, "//go:generate") {
			res = append(res, v)
			return strings.Join(res, "\n")
		} else {
			res = append(res, v)
		}
	}
	panic("unreachable")
}
