package ffmpegutil

import (
	"fmt"
	"log"
	"os"
)

var (
	//Log is a var to handle the log messages
	Log *log.Logger
)

const (
	//FileName is the default log
	FileName = "/var/log/video_service.log"
)

//Init is a fuction to create log file
func Init() {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		_, err := os.Create(FileName)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//WriteLog fuction to add new lines to log file
func WriteLog(prefix string, message string) {
	logHandle, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	} else {
		defer logHandle.Close()
		log.SetOutput(logHandle)
		log.SetFlags(log.Ldate | log.Ltime)
		log.SetPrefix(prefix)
		log.Println(message)
	}
}
