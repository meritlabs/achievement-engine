package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errors "github.com/meritlabs/achievement-engine/pkg/dto"
)

// APP error definition
type appError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSONErrorReporter processes errors after route handler and returns
// JSON with error message
func JSONErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var code int
			message := err.Error()

			switch err.(type) {
			case errors.WrongRequestParametersError:
				code = http.StatusBadRequest
			case errors.BadRequestError:
				code = http.StatusBadRequest
			case errors.ForbiddenError:
				code = http.StatusForbidden
			case errors.UnauthorizedError:
				code = http.StatusUnauthorized
			case errors.NotFoundError:
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			// Put the error into response
			c.AbortWithStatusJSON(code, map[string]string{"error": message})
			return
		}

	}
}
