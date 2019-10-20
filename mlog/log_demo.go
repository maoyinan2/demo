package mlog

import (
	"demo/utilitys"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

/* -----------------falgs-------------------
	Ldate         = 1 << iota     // local time zone: 2009/01/23
    Ltime                         // local time zone: 01:23:23
    Lmicroseconds                 // 01:23:23.123123
    Llongfile                     // /a/b/c/d.go:23
    Lshortfile                    // d.go:23
    LUTC                          // 和北京时间差8h
    LstdFlags     = Ldate | Ltime
*/

/*function: log.New | SetFlags | SetOutput | SetPrefix |
 */

// logHelper .
type logHelper struct {
	today      string
	existedLog []string //已存在的log文件日期（按“月日”保存，eg：1011代表10月11日）
	logFile    *os.File //日志文件
}

var (
	mloggerHelper logHelper
	mlogger       *log.Logger
)

func init() {
	const (
		logPath    = "pkxlog"
		filePrefix = "log_"
		fileSuffix = ".txt"
		timeFormat = "0102"
	)
	// 读取log目录文件
	// var logPath string = "log"
	if utilitys.IsExist(logPath) {
		// 解析文件名，保存到成员
		fi, err := ioutil.ReadDir(logPath)
		if err != nil {
			panic("ioutil.ReadDir(logPath) error.")
		}
		for _, fi := range fi {
			fname := fi.Name()
			mloggerHelper.existedLog = append(mloggerHelper.existedLog, fname[4:8]) // 取“月-日”
		}
	} else {
		os.Mkdir(logPath, os.ModeDir)
	}

	// 获取当天日期（按“月日”保存，0101代表格式）
	mloggerHelper.today = time.Now().Format(timeFormat)

	// 判断log目录下是否有当天日志文件，有，打开文件；没有，创建文件
	var bFind bool = false
	// var filePrefix, fileSuffix string = "log_", ".txt"
	for _, date := range mloggerHelper.existedLog {
		// ? lambda
		if date == mloggerHelper.today {
			bFind = true
			break
		}
	}

	// 创建当日日志文件，配置格式
	filename := path.Join(logPath, filePrefix+mloggerHelper.today+fileSuffix)
	if bFind {
		mloggerHelper.logFile, _ = os.OpenFile(filename, os.O_APPEND, 0666)
	} else {
		mloggerHelper.logFile, _ = os.Create(filename)
	}

	mlogger = log.New(mloggerHelper.logFile, "", log.Ltime|log.Lshortfile)

	// go程：5分钟检查一波日志-》1，日志按天分文件；2，保留7天日志；
	fileTicker := time.NewTicker(5 * time.Second)
	go func(t *time.Ticker) {
		defer mloggerHelper.logFile.Close()

		for {
			<-t.C

			d := time.Now().Format(timeFormat)
			if d == mloggerHelper.today {
				continue
			} else {
				mloggerHelper.logFile.Close()

				mloggerHelper.today = d
				mloggerHelper.existedLog = append(mloggerHelper.existedLog, d)
				filename := path.Join(logPath, filePrefix+d+fileSuffix)
				mloggerHelper.logFile, _ = os.Create(filename)

				// mlogger.SetOutput(mloggerHelper.logFile)
				// for i := 0; i < 5; i++ {
				// 	mlogger.Printf("%d", i)
				// }
				// mlogger.Println("------end------")
			}
		}
	}(fileTicker)
}

// Println .
func Println(prefix string, v ...interface{}) {
	/*  0表示*Logger.Output中调用runtime.Caller的源代码文件和行号
	1表示log.Println中调用*Logger.Output的源代码文件和行号
	2表示main中调用log.Println的源代码文件和行号
	*/
	prefix = "[" + prefix + "] "
	mlogger.SetPrefix(prefix)
	mlogger.Output(2, fmt.Sprintln(v...))
}

// Printf .
func Printf(prefix string, format string, v ...interface{}) {
	prefix = "[" + prefix + "] "
	mlogger.SetPrefix(prefix)
	mlogger.Output(2, fmt.Sprintf(format, v...))
}
