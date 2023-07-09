package MqttServer

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// main entry
func Main() {

	// load config
	if err := ConfigInit(); err != nil {
		log.Panicln(err)
	}

	// init mysql
	if err := SqlUtilInit(); err != nil {
		log.Panicln(err)
	}

	if err := JwtInit(); err != nil {
		log.Panicln(err)
	}

	// run web server
	go func() {
		RunWebServer()
	}()

	// run mqtt server
	if err := MqttServerRun(); err != nil {
		log.Panicln(err)
	}

	// 监听 终止 中断 信号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	// 即将结束 mqtt 服务
	MqttServerClose()
}
