package routes

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// StartRouter 启动路由
func StartRouter() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "我是主页",
		})
	})

	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})

	})
	err := r.Run()
	if err != nil {
		return
	}
}

// GetParameter 获取web的get请求参数
func GetParameter() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default_last_name")
		c.String(http.StatusOK, "%s, %s", firstName, lastName)

	})
	err := r.Run()
	if err != nil {
		return
	}
}

// PostBody 获取body里的参数
func PostBody() {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_post_last_name")
		c.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(bodyBytes))
	})
	err := r.Run()
	if err != nil {
		return
	}
}
