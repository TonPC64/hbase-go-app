package main

import (
	"os"

	"github.com/TonPC64/hbase-go-app/repository"
)

func main() {
	// var (
	// 	table    = "messages"
	// 	qulifier = "m"
	// )
	client := repository.NewHbase(os.Getenv("HBASE_HOST"), os.Getenv("HBASE_TABLE"), "f")
	value := `testdata`
	client.Put("threadId", "1", []byte(value))
	client.GetCell("threadId", "1")
}
