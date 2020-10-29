package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"bytes"
)

/*****************************************
* 作  用:生成markdown文件的目录
* 创建者:高振
*****************************************/
type CataLogUtil struct {
	//读取文件的更目录，需要指定
	rootDir string
	//需要生成的目录文件名，需要指定
	catalogFileName string
	//需要忽略的文件或者目录名称
	ignoreDir map[string]int
	//目录缓存
	buffer bytes.Buffer

}
func main() {
	ignoreDir:= map[string]int{
		".git": 1,
		"img":1,
		"README.md":1,
		"toc.go":1,
		"生成目录.bat":1,

	}
	clu := CataLogUtil{
		catalogFileName: "README.md",
		rootDir:  ".", //当前目录可以这样配置
		ignoreDir:ignoreDir,
	}
	clu.MakeCatalog()
}
func (self *CataLogUtil) MakeCatalog() {
	//1.读取目录
	self.readDir(0,self.rootDir)
	//2.将目录写入到readme文件中
	self.writeFile(self.catalogFileName,self.buffer.Bytes())
	//3.完成后输出提示
	fmt.Println("生成目录完成")
}

func (self *CataLogUtil) writeFile(filename string,content []byte) {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (self *CataLogUtil) readDir(level int , path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for  _,file := range files {
		//如果需要ignore，则结束本次循环
		_,status := self.ignoreDir[file.Name()]
		if status == true{
			continue
		}

		newPath := path + "/" +file.Name()
		//拼接前缀
		//如果是一级目录，格式为: # 一级目录
		//不是一级目录，则格式为 - 其他目录
		var pre = ""
		if level == 0 {
			pre = "#"
		}else{
			for i := level ; i>0 ; i--{
				pre = pre + " "
			}
			pre = pre + "-"
		}
		//写成markdown格式支持的目录，并写入缓冲区
		var catalogLine = ""
		if file.IsDir(){
			catalogLine = fmt.Sprintf("%s %s\n",pre,file.Name())
			self.buffer.WriteString(catalogLine)
			self.readDir(level+1,newPath)
		}else{
			relativePath := strings.Replace(newPath,self.rootDir+"/","",-1)
			catalogLine =  fmt.Sprintf(" %s <a style='text-decoration:none;' href='%s'>%s</a>\n",pre,relativePath,file.Name())
			self.buffer.WriteString(catalogLine)
		}
	}
}
