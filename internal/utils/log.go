package utils

import (
	"github.com/cihub/seelog"
)

const (
	ErrorLevel  = "error"
	WarnLevel   = "warn"
	NoticeLevel = "notice"
	InfoLevel   = "info"
	DebugLevel  = "debug"
	Access      = "access"
)

var errLevels = []string{
	ErrorLevel,
	WarnLevel,
	NoticeLevel,
	InfoLevel,
	DebugLevel,
	Access,
}

type AdLog struct {
	Runtime map[string]seelog.LoggerInterface
	Work seelog.LoggerInterface
}

func (al *AdLog) Init(runtimeLog,work string) error {
	var err error
	/*
	al.Runtime = make(map[string]seelog.LoggerInterface)
	//runtimeLog 需要读里面的内容
	errorByteConfig, err := ioutil.ReadFile(runtimeLog)
	if err != nil {
		return err
	}
	for _, errLevel := range errLevels {
		errorConfig := strings.Replace(string(errorByteConfig), "s%", errLevel, -1)
		if errLevel == DebugLevel {
			DebugLevelRunTimeLogString = errorConfig
		}
		errorConfig = strings.Replace(errorConfig, "m%", "info", -1)
		al.Runtime[errLevel], err = seelog.LoggerFromConfigAsString(errorConfig)
		if err != nil {
			return err
		}
	}
	 */
	if al.Work, err = seelog.LoggerFromConfigAsFile(work); err != nil {
		return err
	}

	return nil
}

func (al *AdLog) Flush() {
	for _, tmp := range al.Runtime {
		tmp.Flush()
	}
}

var Logger *AdLog
var DebugLevelRunTimeLogString = ""

func init() {
	Logger = new(AdLog)
}
