package cmd

import "github.com/gin-gonic/gin"

func Start() {
	engine := gin.Default()
	setRoutes(engine)

	err := engine.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
