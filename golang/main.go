package main

import (
	"go.uber.org/zap"
)

func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	sugar.Infof("failed to fetch url : %s", "test")
}
