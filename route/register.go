// author by @xiaoyusilen

package route

import (
	r "github.com/GoRethink/gorethink"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Name string `json:"name"`
	ID   string
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

	// 解析json
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(200, gin.H{
			"result": false,
		})
		log.Errorf("Get register json err: %s.", err)
		return
	}

	_, err = r.DB("test").Table("test").Insert(map[string]string{
		"name": params.Name,
	}).RunWrite(api.Rethink)

	if err != nil {
		c.JSON(200, gin.H{
			"result": false,
		})
		log.Errorf("Error insert data: %s", err)
		return
	}

	c.JSON(200, gin.H{
		"result": true,
	})
	log.Debugf("Register successed.")
	return
}
