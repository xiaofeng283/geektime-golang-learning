package article

import (
	"github.com/pkg/errors"

	//"github.com/pkg/errors"
	//"log"
	//"geektime-golang-learning/Week02/global"
)

type Article struct {
	Id int64
	Title string
}

func GetArticleInfo(aid string) (*Article, error)  {
	row, err := GetArticleById(aid)
	if err != nil {
		return nil, errors.Wrap(err, "find article error")
	}

	return &Article{Id: row.Id, Title: row.Title}, nil

}