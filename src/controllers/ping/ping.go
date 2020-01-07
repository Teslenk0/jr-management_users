package ping

import (
"net/http"

"github.com/gin-gonic/gin"
)

//Ping - Function to call GET /ping request
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}