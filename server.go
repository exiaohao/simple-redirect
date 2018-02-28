package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

type redirectDict struct {
	Hash string `yaml:"hash"`
	Url  string `yaml:"url"`
}

type redirectConfig struct {
	Redirect []redirectDict `yaml:"REDIRECT"`
}

func initData() {
	rConfigContent, _ := ioutil.ReadFile("settings.yml")
	redirConfig := redirectConfig{}
	err := yaml.Unmarshal(rConfigContent, &redirConfig)

	if err != nil {
		log.Fatalf("Failed to load config")
	}
	fmt.Println(redirConfig)
}

func main() {
	router := gin.Default()

	initData()

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
