package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.GreeterUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}
