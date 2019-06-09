package main

import (
	"github.com/kataras/iris"
	"path/filepath"
	"os"
	"strings"
)

func main() {

	app := iris.New()
	app.RegisterView(iris.HTML("./", ".html"))
	app.StaticWeb("/", "./")

	app.Get("/", func(ctx iris.Context){
		ctx.View("showmovie.html")
	})

	app.Get("/list", func(ctx iris.Context){
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		dir=strings.ReplaceAll(dir, "\\", "/")
		if !strings.HasSuffix(dir,"/"){
			dir=dir+"/"
		}

		list,_:=WalkDir(dir,"mp4")
		liststr:=""
		for _,v:=range(list) {
			tmp:=strings.ReplaceAll(v, dir, "")
			liststr+=tmp+","
		}
		ctx.WriteString(liststr)
	})

	servePort:="80"
	if len(os.Args) == 2 {
		_,errport:=strconv.Atoi(os.Args[1])
		if errport==nil {
			servePort=os.Args[1]
		}
	}

	app.Run(iris.Addr(":"+servePort))
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			fname:=strings.ReplaceAll(filename, "\\", "/")
			files = append(files, fname)
		}

		return nil
	})

	return files, err
}
