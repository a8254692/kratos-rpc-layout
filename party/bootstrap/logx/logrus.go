package logx

import (
	"log"
	"os"

	klogrus "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"

	"gitlab.top.slotssprite.com/my/rpc-layout/party/util/xcolor"
)

// NewLogrusLogger ...
func NewLogrusLogger() (klog.Logger, func()) {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp:  true,
		DisableHTMLEscape: true,
		// PrettyPrint:       true,
	})
	return klogrus.NewLogger(logger), func() { log.Printf(xcolor.GREEN, "logrus logger graceful close") }
}
