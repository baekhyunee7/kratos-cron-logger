package kratos_cron_logger

import (
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type cronLoggerAdapter struct {
	log *log.Helper
}

func NewCronLoggerAdapter(log *log.Helper) *cronLoggerAdapter {
	return &cronLoggerAdapter{
		log: log,
	}
}

func formatTimes(keysAndValues []interface{}) []interface{} {
	var formattedArgs []interface{}
	for _, arg := range keysAndValues {
		if t, ok := arg.(time.Time); ok {
			arg = t.Format(time.RFC3339)
		}
		formattedArgs = append(formattedArgs, arg)
	}
	return formattedArgs
}

func formatString(numKeysAndValues int) string {
	var sb strings.Builder
	sb.WriteString("%s")
	if numKeysAndValues > 0 {
		sb.WriteString(", ")
	}
	for i := 0; i < numKeysAndValues/2; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("%v=%v")
	}
	return sb.String()
}

func (a *cronLoggerAdapter) Info(msg string, keysAndValues ...interface{}) {
	keysAndValues = formatTimes(keysAndValues)
	a.log.Infof(
		formatString(len(keysAndValues)),
		append([]interface{}{msg}, keysAndValues...)...)
}

func (a *cronLoggerAdapter) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = formatTimes(keysAndValues)
	a.log.Errorf(
		formatString(len(keysAndValues)+2),
		append([]interface{}{msg, "error", err}, keysAndValues...)...)
}
