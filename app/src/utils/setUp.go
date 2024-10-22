package utils

import (
	"io"
	"log"
	"os"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

const (
	TRACE_FILE_PATH   = "trace.log"
	INFO_FILE_PATH    = "info.log"
	WARNING_FILE_PATH = "warning.log"
	ERROR_FILE_PATH   = "error.log"
)

var (
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	TRACE = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	WARNING = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

// InitLoggers To initialize loggers
func InitLoggers(u *Utilities) {
	traceLogFile, err1 := os.Create(TRACE_FILE_PATH)
	infoLogFile, err2 := os.Create(INFO_FILE_PATH)
	warningLogFile, err3 := os.Create(WARNING_FILE_PATH)
	errorLogFile, err4 := os.Create(ERROR_FILE_PATH)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Fatalln(u.GetErrorMessage("fail-open-file"))
	}
	initLog(traceLogFile, infoLogFile, warningLogFile, errorLogFile)
	INFO.Println(u.GetMessage("set-up-logger-successfully"))
}

func RedisConnect(u *Utilities) redis.Conn {

	port := u.GetIntConfigValue("general.redis.port")
	host:= u.GetRedisHost()

	uri := fmt.Sprintf("%s:%d", host, port)

	c, err:= redis.Dial("tcp", uri)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}

func ReadInConfig() *viper.Viper {

	v := viper.New()
	v.AddConfigPath("./app/config")
	v.SetConfigType("toml")

	v.SetConfigName("app")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	readInConfigHelper(v, "app.docker")
	readInConfigHelper(v, "app.development")
	readInConfigHelper(v, "app.production")
	readInConfigHelper(v, "app.qa")
	readInConfigHelper(v, "messages")

	return v
}

func readInConfigHelper(v *viper.Viper, fileName string) {

	v.SetConfigName(fileName)
	err := v.MergeInConfig()
	if err != nil {
		log.Fatalln(err)
	}

} 

func GetUtilities(v *viper.Viper) *Utilities {
	return NewUtilites(v)
}

