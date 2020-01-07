package app

import (
	"github.com/Teslenk0/utils-go/logger"
	"github.com/gin-gonic/gin"
)

/*
* Every request that our application receives
* will be handled by this router
 */
var (
	router = gin.Default()
)

//StartApplication - Starts with Caps for export purpose
func StartApplication(){
	//Call the function that will handle the routess
	routes()
	logger.Info("about to start the application")

	//Run the server in port 3010
	router.Run(":3010")
}