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

// LogReport logging in file
func LogReport(logFolder, report string) {
	// file name an its suffix
	var fileName, fileEtx string

	if logFolder == "error" {
		fileEtx = ".err"
	} else if logFolder == "info" {
		fileEtx = ".info"
	} else {
		fileEtx = ".log"
	}

	dateTimeInfo := strings.Split(time.Now().Format("2006-01-02 15:04:05"), " ")

	// Gets current date and time
	dateInfo := strings.Split(dateTimeInfo[0], "-")
	fileName = dateInfo[0] + dateInfo[1] + dateInfo[2]

	var logFileName string
	if runtime.GOOS == "windows" { // log for windows
		// Gets current folder
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		// Makes the log file recursively with "MkdirAll" function,
		os.MkdirAll(filepath.ToSlash(dir)+"/logs/"+logFolder, 0777)

		// the file name
		logFileName = filepath.ToSlash(dir) + "/logs/" + logFolder + "/" + fileName + fileEtx
	} else { // log for linux
		os.MkdirAll("/tmp/logs/"+logFolder, 0777)
		logFileName = "/tmp/logs/" + logFolder + "/" + fileName + fileEtx
	}

	if len(report) != 0 {
		var logfile *os.File
		defer logfile.Close()
		// logging ontents in file with append model
		logfile, _ = os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		io.WriteString(logfile, "["+dateTimeInfo[0]+" "+dateTimeInfo[1]+"]"+report+"\r\n")
	}
}

// LogInformation call "log.Println"
func LogInformation(info ...interface{}) {
	log.Println(info...)
}

// LogPrintf call "log.Printf"
func LogPrintf(format string, info ...interface{}) {
	log.Printf(format, info...)
}
