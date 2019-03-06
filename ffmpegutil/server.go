package ffmpegutil

import (
	"net"
	"strconv"
	"strings"
	"time"
)

//IDResponse is the variable to save response
var IDResponse string

//URLVideo is the variable where the video is
var URLVideo string

//NameFile is the file name to save video
var NameFile string

const (
	//Info is a const used as prefix to save in log file
	Info = "INFO: "
	//Error is a const used as prefix to save in log file
	Error = "ERROR: "
)

//Server is a function to set socket
func Server() {
	Init()
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		WriteLog(Error, err.Error())
	}

	WriteLog(Info, "Listening")

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			WriteLog(Error, err.Error())
		}

		NameFile = strconv.Itoa(time.Now().Nanosecond())

		bs := make([]byte, 1024)

		message, _ := conn.Read(bs)

		splitMessage := strings.Split(string(bs[:message]), " ")

		IDResponse, URLVideo = splitMessage[0], splitMessage[1]

		WriteLog(Info, "Request for response: "+IDResponse)

		conn.Close()
		err = Save()

		if err == nil {
			svc, err := AwsSet()
			if err == nil {
				err := AddFileToS3(svc)
				Delete()
				if err == nil {
					GetFileLink()
					err := Login()
					if err == nil {
						UpdateURL()
						Logout()
					}
				}
			}
			WriteLog(Info, "End connection for response: "+IDResponse)
		}
	}
}
