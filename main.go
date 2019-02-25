package main

import (
	"github.com/TonPC64/hbase-go-app/repository"
)

func main() {
	// var (
	// 	table    = "messages"
	// 	qulifier = "m"
	// )
	client := repository.NewHbase("hbase", "messages", "m")
	// value := `{"id":"85556","threadId":"885999_34344343","sender":"885999","recipient":"34344343","timestamp":1538363596,"message":{"type":"text","text":"hello world"},"senderType":"user","channelId":"885999"}`
	client.GetCell("threadId", "1")
}
