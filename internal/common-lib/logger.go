package commonlib

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CustomTextFormatter struct {
	SessionID string
}

func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// Get the file and line number where the log was called
	_, filename, line, _ := runtime.Caller(7)

	// Get the filename from the full file path
	codeFile := filepath.Base(filename)

	// Format the log message
	message := fmt.Sprintf("[%s] [%s] [%s] [%s:%d] %s\n",
		entry.Time.Format("2006-01-02 15:04:05"),
		entry.Level.String(),
		f.SessionID,
		codeFile,
		line,
		entry.Message,
	)

	return []byte(message), nil
}

func generateSessionID() string {

	randomUUID := uuid.New()
	return strings.Replace(randomUUID.String(), "-", "", -1)
}

func InitLog() {
	// TODO Replace logrus with zerolog -> https://betterstack.com/community/guides/logging/best-golang-logging-libraries/
	sessionID := generateSessionID()
	customFormatter := &CustomTextFormatter{SessionID: sessionID}
	logrus.SetFormatter(customFormatter)
}
