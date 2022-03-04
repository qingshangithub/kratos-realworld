package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "kratos-realworld/api/realworld/v1"
)

//用户相关接口实现
func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) UnfollowUser(ctx context.Context, req *pb.UnfollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

//文章相关接口实现
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
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, req *pb.UnfavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *emptypb.Empty) (*pb.ListofTagsReply, error) {
	return &pb.ListofTagsReply{}, nil
}

//评论相关接口实现
func (s *RealWorldService) AddComments(ctx context.Context, req *pb.AddCommentsRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
