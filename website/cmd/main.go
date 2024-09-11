package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {
	// 获取命令行参数
	serverAddr := flag.String("serverAddr", "", "HTTP服务器网络地址")
	serverPort := flag.String("serverPort", "8080", "HTTP服务器网络端口")
	flag.Parse()

	// 初始化日志
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

	app := &application{
		infoLog: infoLog,
		errLog:  errLog,
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	server := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
