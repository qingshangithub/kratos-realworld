package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-realworld/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Username: "boom",
		},
	}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *emptypb.Empty) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) UnfollowUser(ctx context.Context, req *pb.UnfollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *RealWorldService) AddComments(ctx context.Context, req *pb.AddCommentsRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, req *pb.UnfavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *emptypb.Empty) (*pb.ListofTagsReply, error) {
	return &pb.ListofTagsReply{}, nil
}
