package Log

import (
	"log"
	"os"
)

var g_logger *log.Logger

func GetLogger() *log.Logger{
	if g_logger == nil {
		g_logger = configureInstance()
	}

	return g_logger
}

// 描述符泄露长时间没写入的话关闭描述符
func configureInstance() *log.Logger{
	logName := "bilibili.log"
	logFile, err := os.Create(logName)
	//defer logFile.Close()
	if err != nil {
		log.Fatalln("open file bilibili.log error")
	}

	g_logger = log.New(logFile, "DEBU", log.LstdFlags)
	g_logger.Println("log start")
	return g_logger
}

func RegisterAccountPrefix(account string) string {
	if len(account) == 0 {
		return ""
	}

	return "[" + account + "]"
}

func Debug(message string, vs ...interface{}) {
	GetLogger().SetPrefix("DEBU")
	g_logger.Printf(message, vs)
}

func Info(message string, vs ...interface{}) {
	GetLogger().SetPrefix("INFO")
	g_logger.Printf(message, vs)
}

func Warning(message string, vs ...interface{}){
	GetLogger().SetPrefix("WARN")
	g_logger.Printf(message, vs)
}

func Notice(message string, vs ... interface{}){
	GetLogger().SetPrefix("NOTI")
	g_logger.Printf(message, vs)
}

func Error(message string, vs ...interface{}){
	GetLogger().SetPrefix("ERRO")
	g_logger.Printf(message, vs)
}

func callback(levelId int, level int, message string) {

}
