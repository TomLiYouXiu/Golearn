package main

import (
	"Thought/infra"
	"fmt"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

// ?:Something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com/"))
}
