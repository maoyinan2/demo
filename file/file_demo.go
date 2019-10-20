package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
const (
    O_RDONLY | O_WRONLY | O_RDWR | O_APPEND
    O_CREATE // create a new file if none exists.
    O_EXCL   // used with O_CREATE, file must not exist
    O_SYNC   // open for synchronous I/O.
    O_TRUNC  // if possible, truncate file when opened.
)
const (
    SEEK_SET | SEEK_CUR | SEEK_END
)
*/

/* function
Mkdir | MkdirAll（d:\prj\demo全路径创建） | Remove | RemoveAll | Rename
SameFile | Symlink | TempDir | Truncate |
File {
	Create | NewFile | Open | OpenFile | Pipe |
	Chdir | Chmod | Chown | Close | Fd（文件描述符） | Name |
	Read | ReadAt | Readdir | Readdirnames | Write | WriteAt | WriteString |
	Seek | Stat | Sync | Truncate | IsDir | IsRegular | Pem | String |
}
FileInfo{
	Lstat | Stat |
}
*/
func main() {
	/* 调试运行，c:\users\maoyi\go 当前workspace目录
	go build 生成exe，exe所在目录
	*/
	// dir, err := os.Getwd()
	// fmt.Println(dir)
	// err = os.Mkdir("log", os.ModeDir)
	// if err != nil {
	// 	fmt.Printf("error: %s", err)
	// }

	// 打开文件
	var fname string = "log_"
	fname += "1011.txt"
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// 创建文件
	f, err = os.Create(fname)
	if err != nil {
		fmt.Println(err)
	}

	var logPath string = "log"
	if IsExist(logPath) {
		//遍历目录，解析文件名，保存到成员
		fi, err := ioutil.ReadDir(logPath)
		if err != nil {
			fmt.Println("")
			return
		}
		for _, fi := range fi {
			fname := fi.Name()
			fmt.Println(fname[4:8]) // 取“月-日”
		}
	}

}

// IsExist 判断所给路径文件/文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false // 文件不存在
	}
	return s.IsDir()
}
