// author by @xiaoyusilen

package route

import (
	"github.com/xiaoyusilen/geo-go/repository"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Name string `json:"name"`
}

// Request body
// {
//	"name": "xiaoyusilen"
// }
// Response body
// {
//	"result": true
// }

func (api *RestApi) Register(c *gin.Context) {
	params := RegisterReq{}

	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(200, gin.H{
			"result": false,
		})
		log.Errorf("Get register json err: %s.", err)
		return
	}

	user := &repository.User{
		Name: params.Name,
	}

	res, err := repository.Register(api.Mongo, user, "test")

	if err != nil || !res {
		c.JSON(200, gin.H{
			"result": false,
		})
		log.Errorf("Register failed, err is: %s.", err)
		return
	}

	c.JSON(200, gin.H{
		"result": true,
	})
	log.Debugf("Register successed.")
	return
}
