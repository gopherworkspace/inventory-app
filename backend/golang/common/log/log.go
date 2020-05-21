package applog

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Log *log.Logger

// Create a new instance of the logger. You can have any number of instances.
func init() {
	Log = log.New()
}

func InitializeLogging() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	Log.Out = os.Stdout

	log.SetFormatter(&log.JSONFormatter{})

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("target/logrus.Log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//defer file.Close()
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to Log to file, using default stderr")
	}
}
