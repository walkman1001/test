package main

import (
	//"log"
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func main() {

	fmt.Println("aaa99")
	content, err := rotatelogs.New(
		"/var/log/cli.log"+"-%Y%m%d%H%M",
		rotatelogs.WithLinkName("/var/log/cli.log"), // 生成软链，指向最新日志文件
		//MaxAge and RotationCount cannot be both set  两者不能同时设置
		rotatelogs.WithMaxAge(6*time.Minute), //clear 最小分钟为单位
		//rotatelogs.WithRotationCount(5),        //number 默认7份 大于7份 或到了清理时间 开始清理
		rotatelogs.WithRotationTime(time.Minute), //rotate 最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
	)

	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	logrus.SetOutput(content)

	f := func() {
		for i := 0; i < 100; i++ {
			logrus.WithFields(logrus.Fields{

				"animal": "walrus",

				"number": i,
			}).Info("A walrus appears")

			logrus.Error("xxxxxxxxxxxxxxx")
			logrus.Error("xxxxxxxxxxxxxxx2")
			logrus.Error("xxxxxxxxxxxxxxx3")

			time.Sleep(time.Second)
		}

	}

	for i := 0; i < 1000000; i++ {

		go f()
		time.Sleep(time.Second)
	}
	f()

	time.Sleep(121 * time.Second)
}
