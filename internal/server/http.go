package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
	"kratos-realworld/internal/service"
)

func NewSkipRoutesMatcher() selector.MatchFunc {

	skipRoutes := make(map[string]struct{})
	skipRoutes["/realworld.v1.RealWorld/Login"] = struct{}{}
	skipRoutes["/realworld.v1.RealWorld/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRoutes[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, greeter *service.RealWorldService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),

		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwtc.Token)).Match(NewSkipRoutesMatcher()).Build(),
			//auth.JWTAuth(jwtc.Token),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealWorldHTTPServer(srv, greeter)
	return srv
}
