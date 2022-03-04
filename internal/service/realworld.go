package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-realworld/internal/biz"

	pb "kratos-realworld/api/realworld/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewRealWorldService(uc *biz.GreeterUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}
