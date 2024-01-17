package router

import (
	"context"

	pb "github.com/loak155/microservices/proto/article"
	"github.com/loak155/microservices/services/article/domain"
	"github.com/loak155/microservices/services/article/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type articleGRPCServer struct {
	pb.UnimplementedArticleServiceServer
	au usecase.IArticleUsecase
}

func NewArticleGRPCServer(grpcServer *grpc.Server, au usecase.IArticleUsecase) pb.ArticleServiceServer {
	s := articleGRPCServer{au: au}
	pb.RegisterArticleServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)
	return &s
}

func (s *articleGRPCServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	res := pb.CreateArticleResponse{}
	article := domain.Article{Title: req.Article.Title, Url: req.Article.Url}
	articleRes, err := s.au.CreateArticle(article)
	if err != nil {
		return nil, err
	}
	res.Article = &pb.Article{
		Id:            int32(articleRes.ID),
		Title:         articleRes.Title,
		Url:           articleRes.Url,
		BookmarkCount: int32(article.BookmarkCount),
		CreatedAt:     &timestamppb.Timestamp{Seconds: int64(articleRes.CreatedAt.Unix()), Nanos: int32(articleRes.CreatedAt.Nanosecond())},
		UpdatedAt:     &timestamppb.Timestamp{Seconds: int64(articleRes.UpdatedAt.Unix()), Nanos: int32(articleRes.UpdatedAt.Nanosecond())},
	}

	return &res, nil
}

func (s *articleGRPCServer) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	res := pb.GetArticleResponse{}
	articleRes, err := s.au.GetArticle(int(req.Id))
	if err != nil {
		return nil, err
	}
	res.Article = &pb.Article{
		Id:            int32(articleRes.ID),
		Title:         articleRes.Title,
		Url:           articleRes.Url,
		BookmarkCount: int32(articleRes.BookmarkCount),
		CreatedAt:     &timestamppb.Timestamp{Seconds: int64(articleRes.CreatedAt.Unix()), Nanos: int32(articleRes.CreatedAt.Nanosecond())},
		UpdatedAt:     &timestamppb.Timestamp{Seconds: int64(articleRes.UpdatedAt.Unix()), Nanos: int32(articleRes.UpdatedAt.Nanosecond())},
	}

	return &res, nil
}

func (s *articleGRPCServer) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	res := pb.ListArticlesResponse{}
	articleRes, err := s.au.ListArticles()
	if err != nil {
		return nil, err
	}
	for _, article := range articleRes {
		res.Articles = append(res.Articles, &pb.Article{
			Id:            int32(article.ID),
			Title:         article.Title,
			Url:           article.Url,
			BookmarkCount: int32(article.BookmarkCount),
			CreatedAt:     &timestamppb.Timestamp{Seconds: int64(article.CreatedAt.Unix()), Nanos: int32(article.CreatedAt.Nanosecond())},
			UpdatedAt:     &timestamppb.Timestamp{Seconds: int64(article.UpdatedAt.Unix()), Nanos: int32(article.UpdatedAt.Nanosecond())},
		})
	}

	return &res, nil
}

func (s *articleGRPCServer) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	res := pb.UpdateArticleResponse{}
	article := domain.Article{
		ID:            uint(req.Article.Id),
		Title:         req.Article.Title,
		Url:           req.Article.Url,
		BookmarkCount: uint(req.Article.BookmarkCount),
		CreatedAt:     req.Article.CreatedAt.AsTime(),
		UpdatedAt:     req.Article.UpdatedAt.AsTime(),
	}
	articleRes, err := s.au.UpdateArticle(article)
	if err != nil {
		return nil, err
	}
	res.Success = articleRes

	return &res, nil
}

func (s *articleGRPCServer) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	res := pb.DeleteArticleResponse{}
	articleRes, err := s.au.DeleteArticle(int(req.Id))
	if err != nil {
		return nil, err
	}
	res.Success = articleRes

	return &res, nil
}
