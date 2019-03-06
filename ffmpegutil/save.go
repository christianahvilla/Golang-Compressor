package ffmpegutil

import (
	"os"
	"os/exec"
)

//Save is to save compressed videos
func Save() error {

	os.MkdirAll("videos", os.ModePerm)

	file := exec.Command("ffmpeg", "-i", URLVideo, "-strict", "-2", "videos/"+NameFile+".mp4")

	err := file.Run()

	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		WriteLog(Info, "File saved for response: "+IDResponse)
	}
	return err
}

//Delete is a function to delete video after save in amazon
func Delete() error {
	err := os.Remove("videos/" + NameFile + ".mp4")

	if err != nil {
		WriteLog(Error, err.Error())
	} else {
		WriteLog(Info, "File deleted for response: "+IDResponse)
	}

	return err
}
