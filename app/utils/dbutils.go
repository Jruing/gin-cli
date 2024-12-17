package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"time"
)

var Pgconn *pgx.Conn

func init() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=127.0.0.1  port=5432 user=pqgotest password=123456 dbname=pqgotest sslmode=verify-full")
	if err != nil {

	}
	defer conn.Close(ctx)

	// 创建一个日志文件路径
	logPath := "./storage/logs/server.log"
	// 获取当前时间
	now := time.Now()

	// 计算明天0点的时间
	tomorrowMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

	// 计算明天0点与当前时间的时间差
	//durationUntilTomorrowMidnight := tomorrowMidnight.Sub(now)

	// 创建轮转日志对象
	rotatelogsPath, err := rotatelogs.New(
		logPath+".%Y%m%d%H%M%S",                          // 日志文件命名规则
		rotatelogs.WithLinkName(logPath),                 // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),            // 设置日志最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour),        // 设置日志切割时间间隔
		rotatelogs.WithStartRotationAt(tomorrowMidnight), // 设置起始切割时间为明天0点
	)
	if err != nil {
		fmt.Println("无法创建轮转日志对象:", err)
		return
	}
	// 创建一个 Logrus 日志记录器
	logger := logrus.New()

	// 设置日志格式为 JSON 格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置日志输出到轮转日志文件
	writer := lfshook.WriterMap{
		logrus.InfoLevel:  rotatelogsPath,
		logrus.FatalLevel: rotatelogsPath,
		logrus.ErrorLevel: rotatelogsPath,
		logrus.WarnLevel:  rotatelogsPath,
		logrus.DebugLevel: rotatelogsPath,
		logrus.PanicLevel: rotatelogsPath,
	}

	logger.AddHook(lfshook.NewHook(writer, &logrus.JSONFormatter{}))

	// 记录一条日志
	logger.Infof("This is an info message")
}
