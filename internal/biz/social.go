package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type SocialUsecase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo
	pr ProfileRepo

	log *log.Helper
}

type Article struct {
	Slug  string
	Title string
}

func NewSocialUsecase(ar ArticleRepo,
	pr ProfileRepo,
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, pr: pr, log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) CreateArticle(ctx context.Context, in *Article) error {
	return nil
}
