package qlogtracer

import (
	"bufio"
	"fmt"
	"io"
	"liangrun/utils"
	"log"
	"os"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/logging"
	"github.com/lucas-clemente/quic-go/qlog"
	"github.com/sirupsen/logrus"
)

func AddTracer(isServer bool, config *quic.Config) {

	tracer := qlog.NewTracer(func(_ logging.Perspective, connID []byte) io.WriteCloser {
		var qlogDir string
		var filename string
		if isServer {
			qlogDir = utils.ServerEnvMap["qLogDir"]
			filename = fmt.Sprintf("%sserver_%x.qlog", qlogDir, connID)
		} else {
			qlogDir = utils.ClientEnvMap["qLogDir"]
			filename = fmt.Sprintf("%sclient_%x.qlog", qlogDir, connID)
		}
		if err := os.MkdirAll(qlogDir, os.ModePerm); err != nil {
			logrus.Error(err)
		}
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			log.Fatal(err)
		}

		return NewBufferedWriteCloser(bufio.NewWriter(f), f)
	})
	config.Tracer = tracer

}
