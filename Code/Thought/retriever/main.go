package main

import (
	"Thought/retriever/mock"
	"Thought/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com/")
}
func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

}

const url = "https://www.imooc.com/"

func post(poster Poster) {
	poster.Post("https://www.imooc.com/",
		map[string]string{
			"name":   "liyouxiu",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "anther fake imooc.com",
	})
	return s.Get(url)
}
func main() {
	var r Retriever
	retriever := mock.Retriever{
		"this is fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockretriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockretriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

}
