package main

import (
	"github.com/WandersonLontra/gopportunities/configs"
	"github.com/WandersonLontra/gopportunities/router"
)
func main(){
	logger := configs.GetLogger("main")
	//Initialize configs
	err := configs.Init()
	if err != nil {
		logger.Errorf("Error on initialize configs: %v ",err)
		panic(err)
	}
	router.InitRouter()
}