package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/app/db"
	"net/http"
	"time"
)

func Parse(c *gin.Context) {
	ip := c.Query("ip")
	loc, err := db.Search(ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "获取 ip 地址信息失败",
			"ok":   false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"msg":      "获取 ip 地址信息成功",
		"location": loc,
		"ok":       true,
	})
}

func Generate(c *gin.Context) {
	tStart := time.Now()
	err := db.GenerateDB("./app/db/source/ip.merge.txt", "./app/db/data/ip2region.xdb")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "生成数据库失败",
			"ok":   false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  fmt.Sprintf("生成数据库成功,耗时: %v", time.Since(tStart)),
		"ok":   true,
	})
}
