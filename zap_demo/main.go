package zap_demo

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()

	simpleHttpGet("www.qq.com")
	simpleHttpGet("https://www.qq.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
