package main

import (
	"context"
	"flag"
	"fmt"
	"go-mongodb-demo/users/pkg/models/mongodb"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	users   *mongodb.UserModel
}

func main() {

	// 获取命令行参数
	serverAddr := flag.String("serverAddr", "", "HTTP服务器网络地址")
	serverPort := flag.Int("serverPort", 4000, "HTTP服务器网络端口")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "MongoDB连接字符串")
	mongoDatabase := flag.String("mongoDatabase", "users", "MongoDB数据库名")
	enableCredentials := flag.Bool("enableCredentials", false, "启用用户认证")
	flag.Parse()

	// 初始化日志
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

	// 创建 MongoDB 客户端配置
	clientOptions := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		clientOptions.Auth = &options.Credential{
			Username: os.Getenv("MONGO_USERNAME"),
			Password: os.Getenv("MONGO_PASSWORD"),
		}
	}

	// 创建超时上下文链
	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	// 建立数据库连接
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// 初始化一个包含依赖项的应用实例
	app := &application{
		infoLog: infoLog,
		errLog:  errLog,
		users: &mongodb.UserModel{
			C: client.Database(*mongoDatabase).Collection("users"),
		},
	}

	// 初始化 HTTP 服务器
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	server := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("users服务运行于 %s", serverURI)
	err = server.ListenAndServe()
	errLog.Fatal(err)
}
