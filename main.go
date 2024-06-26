package main

import (
	"fmt"

	"github.com/WillianIsami/go_api/config"
	"github.com/WillianIsami/go_api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error %v", err)
		fmt.Print(err)
		return
	}

	router.Initialize()
}
