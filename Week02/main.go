package main

import (
	"geektime-golang-learning/Week02/article"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/article", article.ArticleHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

