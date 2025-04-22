package initialize

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
