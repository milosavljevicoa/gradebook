package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Initialize(r *gin.Engine) {
	initializeLogger("./app.log")
	registerHandlers(r)
}

func initializeLogger(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

func registerHandlers(r *gin.Engine) {
	r.POST("/log", func(c *gin.Context) {
		msg, err := ioutil.ReadAll(c.Request.Body)

		if err != nil || len(msg) == 0 {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		write(string(msg))
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
