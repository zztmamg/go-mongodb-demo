package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type apis struct {
	users string
}

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	apis    apis
}

func main() {
	// 获取命令行参数
	serverAddr := flag.String("serverAddr", "", "HTTP服务器网络地址")
	serverPort := flag.Int("serverPort", 8000, "HTTP服务器网络端口")
	usersAPI := flag.String("usersAPI", "http://localhost:4000/api/users", "users API")
	flag.Parse()

	// 初始化日志
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

	app := &application{
		infoLog: infoLog,
		errLog:  errLog,
		apis: apis{
			users: *usersAPI,
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("服务运行于 %s", serverURI)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
