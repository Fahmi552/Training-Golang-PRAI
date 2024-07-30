package handler

import (
	pb "Assignment4/proto"
	"Assignment4/service"
	"context"
)

type URLShortenerServer struct {
	pb.UnimplementedURLShortenerServer
	service service.URLService
}

func NewURLHandlerServer(service service.URLService) *URLShortenerServer {
	return &URLShortenerServer{service: service}
}

func (s *URLShortenerServer) ShortenURL(ctx context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	shortURL, err := s.service.ShortenURL(ctx, req.OriginalUrl)
	if err != nil {
		return nil, err
	}
	return &pb.ShortenResponse{ShortUrl: shortURL}, nil
}

func (s *URLShortenerServer) GetOriginalURL(ctx context.Context, req *pb.OriginalRequest) (*pb.OriginalResponse, error) {
	originalURL, err := s.service.GetOriginalURL(req.ShortUrl)
	if err != nil {
		return nil, err
	}
	return &pb.OriginalResponse{OriginalUrl: originalURL}, nil
}
