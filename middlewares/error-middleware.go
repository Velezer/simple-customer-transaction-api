package middlewares

import (
	"net/http"
	"sctrans/httperror"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			meta, ok := err.Meta.(httperror.Meta)
			if ok {
				if meta.Data == nil {
					meta.Data = err.Err.Error()
				}
				c.JSON(meta.Code, gin.H{"error": meta.Data})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Err.Error()})
		}
	}
}
