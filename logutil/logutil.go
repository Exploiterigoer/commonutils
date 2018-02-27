package logutil

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// 数据入库日志记录函数
func LogReport(logFolder, report string) {
	// 文件名
	var fileName, fileEtx string

	// 如果没传入mn，那么表示这是一个错误日志
	if logFolder == "error" {
		fileEtx = ".err"
	} else if logFolder == "info" {
		fileEtx = ".info"
	} else {
		fileEtx = ".log"
	}
	// 日志时间
	dateTimeInfo := strings.Split(time.Now().Format("2006-01-02 15:04:05"), " ")

	// 当天的日期、时间
	dateInfo := strings.Split(dateTimeInfo[0], "-")
	fileName = dateInfo[0] + dateInfo[1] + dateInfo[2]

	var logFileName string
	if runtime.GOOS == "windows" { // log for windows
		// 当前文件所在的文件夹
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		// 创建以MN号为名的目录，调用MkdirAll会递归创建目录
		os.MkdirAll(filepath.ToSlash(dir)+"/logs/"+logFolder, 0777)

		// 绝对路径下生成日志文件
		logFileName = filepath.ToSlash(dir) + "/logs/" + logFolder + "/" + fileName + fileEtx
	} else { // log for linux
		// 创建以MN号为名的目录，调用MkdirAll会递归创建目录
		os.MkdirAll("/tmp/logs/"+logFolder, 0777)
		logFileName = "/tmp/logs/" + logFolder + "/" + fileName + fileEtx
	}

	// 把日志信息追加写入到文件
	if len(report) != 0 {
		var logfile *os.File
		defer logfile.Close()
		// 不存在则创建，存在则以读写模式打开文件并追加内容
		logfile, _ = os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		io.WriteString(logfile, "["+dateTimeInfo[0]+" "+dateTimeInfo[1]+"]"+report+"\r\n")
	}
}

func LogInformation(info ...interface{}) {
	log.Println(info...)
}

func LogPrintf(format string, info ...interface{}) {
	log.Printf(format, info...)
}
