package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var rdb *redis.Client

func main() {
	// 初始化Gin引擎
	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	server.Use(cors.New(corsConfig))

	// 初始化Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // 这里使用容器名称来连接Redis
	})

	// 创建一个路由来处理请求并记录日志
	server.GET("/error-test", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "is an error"})
	})

	server.GET("/logs", listLogs)
	server.POST("/log", createLog)
	server.DELETE("/logs", clearLogs)

	// 运行Gin应用程序
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

type logMessageRequest struct {
	Log string `json:"log" binding:"required"`
}

func createLog(c *gin.Context) {
	var req logMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 向Redis写入日志
	err := logRequest(c.Request, req.Log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged successfully"})
}

func listLogs(c *gin.Context) {
	ctx := context.Background()
	key := "logs"

	// 从Redis列表中获取所有日志条目
	logs, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// reversedLogs := make([]string, len(logs))
	// for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
	// 	logs[i], logs[j] = logs[j], logs[i]
	// }

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}

func clearLogs(c *gin.Context) {
	ctx := context.Background()
	key := "logs"

	// 删除Redis列表中的所有日志
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logs cleared successfully"})
}

func logRequest(req *http.Request, log string) error {
	ctx := context.Background()
	key := "logs"
	// 将日志条目写入Redis列表
	_, err := rdb.LPush(ctx, key, log).Result()
	return err
}
