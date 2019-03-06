# Golang-Compressor
Golang-Compressor is a lot of things. Creation a daemon for Linux machines to compress videos and upload to S3 and alter a database with the new URL.

### Built With
* [FFmpeg](https://www.ffmpeg.org/) - Converting video and audio has never been so easy.

* [Amazon S3](https://aws.amazon.com/es/s3/) - Storage of objects created to store and recover any volume of data from any location.

### Prerequisites
* FFmpeg
* Golang
* AWS-CLI

### Installing
Follow the next stepts to install correctly.

#### Golang

```
$ wget -q https://storage.googleapis.com/golang/getgo/installer_linux

$ chmod +x installer_linux

$ ./installer_linux

Welcome to the Go installer!
Downloading Go version go1.10 to /home/linuxconfig/.go
This may take a bit of time...
Downloaded!
Setting up GOPATH
GOPATH has been set up!

One more thing! Run `source /home/linuxconfig/.bash_profile` to persist the
new environment variables to your current session, or open a
new shell prompt.

$ source /home/linuxconfig/.bash_profile

$ go version
go version go1.10 linux/amd64
```

##### Golang Libraries

* [Daemon](https://github.com/takama/daemon)
```
$ go get github.com/takama/daemon
```

* [AWS SDK for Go](https://github.com/aws/aws-sdk-go)
```
$ go get github.com/aws/aws-sdk-go
```

#### FFmpeg

```
$ sudo apt install ffmpeg

$ ffmpeg -version

ffmpeg version 4.1.1-0york1~18.04 Copyright (c) 2000-2019 the FFmpeg developers
built with gcc 7 (Ubuntu 7.3.0-27ubuntu1~18.04)
```

#### AWS-CLI
```
$ sudo apt install awscli

$ aws --version

aws-cli/1.14.44 Python/3.6.7 Linux/4.15.0-45-generic botocore/1.8.48

$ aws configure

AWS Access Key ID [****************XXXX]: XXXXXXXXXXXXXXXXXXXX

AWS Secret Access Key [****************XXXX]: XXXXXXXXXXXXXXXXXXXX

Default region name [us-east-1]: us-east-1

Default output format [None]:
```

### Configure

#### Client

Copy client binary in the next path `laravel-project/public/go/client/`
```
$ pwd
/home/vagrant/code/public/go/client
```

#### Library ffmpeeutil

Copy folder in the next path `/home/$username$/go/src/`
```
$ pwd
/home/vagrant/go/src/ffmpegutil
```

Build and instasll library

**Remember, You need to be inside the `/home/$username$/go/src/`**

```
$ go build ffmpegutil
$ go install ffmpegutil
```

#### Daemon

Compile the daemon

```
$ go build main.go
```

Prepare enviroment before start

```
$ cp main /bin/ffmpeg-util-server
$ cd /bin
$ ./ffmpeg-util-server install

cat -f /var/log/video_service.log

INFO: 2019/03/06 14:36:30 Install comprres videos:                                    [  OK  ]
```
### Running the tests

##### Invoke the daemon
```
$ sudo systemctl start ffmpeg-video.service

$ sudo systemctl status ffmpeg-video.service

ffmpeg-video.service - comprres videos
    Loaded: loaded (/etc/systemd/system/ffmpeg-video.service; enabled; vendor preset: enabled)
    Active: active (running) since Wed 2019-03-06 21:16:49 UTC; 3s ago
    Process: 6740 ExecStartPre=/bin/rm -f /var/run/ffmpeg-video.pid (code=exited, status=0/SUCCESS)
    Main PID: 6747 (ffmpeg-util-ser)
    Tasks: 8 (limit: 4915)
    CGroup: /system.slice/ffmpeg-video.service
            +-6747 /bin/ffmpeg-util-server

Mar 06 21:16:49 homestead systemd[1]: Starting comprres videos...
Mar 06 21:16:49 homestead systemd[1]: Started comprres videos.
```
#### Run client

Go to the next path `laravel-project/public/go/client/`

```
$ ./client 33361b3c-0d1e-44b5-93b5-025214070b3b https://vis3dev.s3.amazonaws.com/xxxxxxxxx.mp4
```

##### Watch log

Look the log to know the process
```
$ tail -f /var/log/video_service.log

INFO: 2019/03/06 21:16:49 Listening
INFO: 2019/03/06 14:59:07 File saved for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:07 Session opened for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:08 File uploaded for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:08 File deleted for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:08 URL https://vis3dev.s3.amazonaws.com/xxxxxxxxxx.mp4 for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:09 Login for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:11 Updated for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:11 Logout for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
INFO: 2019/03/06 14:59:11 End connection for response: 33361b3c-0d1e-44b5-93b5-025214070b3b
```

### Author

* **Christian Alejandro Herrejon Villa** - *Application Developer* - [Scio](https://sciodev.com/)

### License

This project is licensed under the GNU License - see the LICENSE file for details