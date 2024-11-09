package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToId(c *gin.Context, userId string) (err error) {
	uid := c.GetString("userId")

	if uid != userId {
		return fmt.Errorf("Unauthorized access to this resource")
	}

	return nil
}
