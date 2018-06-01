package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/db/stores"
)

// Auth is a token-based auth middleware
func Auth(s stores.SessionsStore, u stores.UsersStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Token")

		session, err := s.GetSession(token)
		if err != nil {
			fmt.Printf("Session not found! Error: %s\n", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := u.GetUser(session.UserID)
		if !user.Verified {
			fmt.Printf("Provided user is not varified! Error: %s\n", err.Error())
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Set the user in the context and move onto the next function in the chain
		c.Set("user", user)

		c.Next()
	}
}