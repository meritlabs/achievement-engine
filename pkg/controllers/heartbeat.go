package controllers

import "github.com/gin-gonic/gin"
import "net/http"

// Heartbeat tells outside parties that the service is up and acting normally
// This can be used in the case of load_balancers, DNS round robin, and GeoDNS
func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
