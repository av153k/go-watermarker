package transport

import (
	"context"

	"github.com/av153k/go-watermarker/api/v1/pb/watermarker"
	"github.com/av153k/go-watermarker/internal"
	"github.com/av153k/go-watermarker/pkg/watermark/endpoints"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	serviceStatus grpctransport.Handler
	addDocument   grpctransport.Handler
	watermark     grpctransport.Handler
	watermarker.UnimplementedWatermarkerServer
}

func NewGRPCServer(endpoint endpoints.Set) watermarker.WatermarkerServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			endpoint.GetEndpoint,
			decodeGRPCGetRequest, decodeGRPCGetResponse,
		),
		status: grpctransport.NewServer(
			endpoint.StatusEndpoint,
			decodeGRPCStatusRequest, decodeGRPCStatusResponse,
		),
		serviceStatus: grpctransport.NewServer(
			endpoint.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest, decodeGRPCServiceStatusResponse,
		),
		addDocument: grpctransport.NewServer(
			endpoint.AddDocumentEndpoint,
			decodeGRPCAddDocumentRequest, decodeGRPCAddDocumentResponse,
		),
		watermark: grpctransport.NewServer(
			endpoint.WatermarkEndpoint,
			decodeGRPCWatermarkRequest, decodeGRPCWatermarkResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, request *watermarker.GetRequest) (*watermarker.GetResponse, error) {
	_, response, err := g.get.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*watermarker.GetResponse), nil
}

func (g *grpcServer) Status(ctx context.Context, request *watermarker.StatusRequest) (*watermarker.StatusResponse, error) {
	_, response, err := g.get.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*watermarker.StatusResponse), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, request *watermarker.ServiceStatusRequest) (*watermarker.ServiceStatusResponse, error) {
	_, response, err := g.get.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*watermarker.ServiceStatusResponse), nil
}

func (g *grpcServer) AddDocument(ctx context.Context, request *watermarker.AddDocumentRequest) (*watermarker.AddDocumentResponse, error) {
	_, response, err := g.get.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*watermarker.AddDocumentResponse), nil
}

func (g *grpcServer) Watermark(ctx context.Context, request *watermarker.WatermarkRequest) (*watermarker.WatermarkResponse, error) {
	_, response, err := g.get.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*watermarker.WatermarkResponse), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*watermarker.GetRequest)
	var filters []internal.Filter
	for _, f := range request.Filters {
		filters = append(filters, internal.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.GetRequest{Filters: filters}, nil
}

func decodeGRPCStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*watermarker.StatusRequest)
	return endpoints.StatusRequest{TicketID: request.TicketID}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCWatermarkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*watermarker.WatermarkRequest)
	return endpoints.WatermarkRequest{TicketID: request.TicketID, Mark: request.Mark}, nil
}

func decodeGRPCAddDocumentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*watermarker.AddDocumentRequest)
	doc := &internal.Document{
		Title:     request.Document.Title,
		Content:   request.Document.Content,
		Author:    request.Document.Author,
		Topic:     request.Document.Topic,
		Watermark: request.Document.Watermark,
	}
	return endpoints.AddDocumentRequest{Document: doc}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*watermarker.GetResponse)
	var docs []internal.Document
	for _, d := range response.Documents {
		doc := internal.Document{
			Content:   d.Content,
			Title:     d.Title,
			Author:    d.Author,
			Topic:     d.Topic,
			Watermark: d.Watermark,
		}
		docs = append(docs, doc)
	}
	return endpoints.GetResponse{Documents: docs}, nil
}

func decodeGRPCStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*watermarker.StatusResponse)
	return endpoints.StatusResponse{Status: internal.Status(response.Status)}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*watermarker.ServiceStatusResponse)
	return endpoints.ServiceStatusResponse{Code: int(response.Code)}, nil
}

func decodeGRPCAddDocumentResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*watermarker.AddDocumentResponse)
	return endpoints.AddDocumentResponse{TicketID: response.TicketID}, nil
}

func decodeGRPCWatermarkResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	response := grpcReply.(*watermarker.WatermarkResponse)
	return endpoints.WatermarkResponse{Code: int(response.Code)}, nil
}
