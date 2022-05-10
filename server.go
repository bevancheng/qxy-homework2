package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/Moonlight-Zhao/go-project-example/cotroller"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	//
	runtime.SetBlockProfileRate(1)
	runtime.SetBlockProfileRate(1)
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	//
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/topic/do", func(c *gin.Context) {
		//uid, _ := c.GetPostForm("uid")
		title, _ := c.GetPostForm("title")
		content, _ := c.GetPostForm("content")
		data := cotroller.PublishTopic(title, content)
		//repository.FlushTopicAll("./data/")
		c.JSON(200, data)
	})
	err := r.Run()

	if err != nil {
		return
	}

}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
