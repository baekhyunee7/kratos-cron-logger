# kratos-cron-logger
[cron](https://github.com/robfig/cron) logger adapter for [go-kratos](https://github.com/go-kratos/kratos)

```go
import (
	"os"

	"github.com/baekhyunee7/kratos_cron_logger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

func main() {
	logger := log.NewStdLogger(os.Stdout)
	helper := log.NewHelper(logger)
	_ = cron.New(
		cron.WithChain(
			cron.Recover(kratos_cron_logger.NewCronLoggerAdapter(helper)),
		),
	)
}
```