package main

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	redisClient, redisErr := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
	if redisErr != nil {
		log.Fatalf("Connect redis server failed: %v", redisErr)
	}
	redisClient.Do("SELECT", os.Getenv("REDIS_DB"))
	defer redisClient.Close()

	router.GET("/:redirectKey", func(c *gin.Context) {
		redirectKey := c.Param("redirectKey")
		// c.String(http.StatusOK, "Hello %s", name)
		resultURL, err := redis.String(redisClient.Do("GET", redirectKey))
		if err != nil {
			c.Redirect(302, "http://caafashion.top/")
		}
		c.Redirect(302, resultURL)
	})

	router.Run(":" + os.Getenv("SERVER_PORT"))
}
