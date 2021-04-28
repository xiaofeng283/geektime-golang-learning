package article

import (
	"github.com/pkg/errors"
)

func GetArticle(aid string) (article *Article, err error)  {
	articleInfo , err:= GetArticleInfo(aid)
	if err != nil {
		return nil, errors.WithMessage(err, "get article error")
	}
	return articleInfo , err
}