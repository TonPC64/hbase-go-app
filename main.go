package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

func main() {
	client := gohbase.NewClient("hbase")

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		getRequest, err := hrpc.NewGetStr(context.Background(), "table", "key")
		getRsp, err := client.Get(getRequest)

		c.JSON(200, gin.H{
			"getRsp": fmt.Sprintln(getRsp),
			"err":    fmt.Sprintln(err),
		})
	})

	r.GET("/insert", func(c *gin.Context) {
		values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}
		putRequest, err := hrpc.NewPutStr(context.Background(), "table", "key", values)
		rsp, err := client.Put(putRequest)

		c.JSON(200, gin.H{
			"rsp": fmt.Sprintln(rsp),
			"err": fmt.Sprintln(err),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
