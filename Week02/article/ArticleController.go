package article

import (
	"fmt"
	"net/http"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		aid:= r.FormValue("id")
		Article, err := GetArticle(aid)
		if err != nil {
			fmt.Printf("Article not found , id:%s , %+v", aid,err)
			_, _ = w.Write([]byte(fmt.Sprintf("Article not found , id:%s", aid)))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Id：%s，Title：%s", Article.Id,Article.Title)))
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}
