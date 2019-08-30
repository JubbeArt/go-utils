package lg

import (
	"log"
	"os"
	"path/filepath"
)

var logFile *os.File
var errorFile *os.File

var Log *log.Logger
var Err *log.Logger

func Init(logFolder string) error {
	err := os.MkdirAll(logFolder, 0755)

	if err != nil {
		return err
	}

	logFile, err = os.OpenFile(filepath.Join(logFolder, "log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	errorFile, err = os.OpenFile(filepath.Join(logFolder, "error"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		logFile.Close()
		return err
	}

	Log = log.New(logFile, "", log.LstdFlags)
	Err = log.New(errorFile, "", log.LstdFlags)
	return nil
}

func Close() {
	logFile.Close()
	errorFile.Close()

}
