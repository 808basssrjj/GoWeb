package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
)

//func main() {
//	InitLogger()
//	defer sugarLogger.Sync()
//	simpleHttpGet("www.baidu.com")
//	simpleHttpGet("https://www.baidu.com")
//}
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello liwenzhou.com!")
	})
	_ = r.Run(":7071")
}

//func InitLogger() {
//	//调用zap.NewProduction()/zap.NewDevelopment()或者zap.Example()创建一个Logger。
//	logger, _ = zap.NewProduction()
//	sugarLogger = logger.Sugar()
//}

// InitLogger
// 将日志写入文件而不是终端
// 使用zap.New(…)方法来手动传递所有配置
// Encoder:编码器(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()。
// WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去。
// Log Level：哪种级别的日志将被写入。
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	//logger := zap.New(core)
	//记录调用方函数的详细信息
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	//1.JSON Encoder更改为普通Encoder
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	//2.修改时间格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//return zapcore.AddSync(file)

	// 使用lumberjack 进行日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     //M
		MaxBackups: 5,     //最大备份数量
		MaxAge:     30,    //最大备份天数
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

//func simpleHttpGet(url string) {
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarLogger.Error(
//			"error fetching url",
//			zap.String("url", url),
//			zap.Error(err),
//		)
//	} else {
//		sugarLogger.Info(
//			"success",
//			zap.String("statusCode", resp.Status),
//			zap.Error(err),
//		)
//		_ = resp.Body.Close()
//	}
//}
