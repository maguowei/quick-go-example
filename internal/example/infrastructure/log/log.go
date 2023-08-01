package log

import (
	"github.com/maguowei/lego/kit/logrus-hooks/fields"
	"github.com/orandin/sentrus"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func NewLogger(appName, appEnv, logPath string) *logrus.Logger {
	f, err := os.OpenFile(filepath.Join(logPath, "app.log"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.MultiWriter(os.Stdout, f))
	logger.AddHook(sentrus.NewHook([]logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}))
	logger.AddHook(fields.NewHook(logrus.Fields{"appName": appName, "appEnv": appEnv}, []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.InfoLevel, logrus.DebugLevel}))
	return logger
}
