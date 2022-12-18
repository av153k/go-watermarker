package endpoints

import (
	"context"
	"errors"
	"os"

	"github.com/av153k/go-watermarker/internal"
	"github.com/av153k/go-watermarker/pkg/watermark"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint     endpoint.Endpoint
}

func NewEndpointSet(service watermark.Service) Set {
	return Set{
		GetEndpoint:           MarkGetEndpoint(service),
		AddDocumentEndpoint:   MarkAddDocumentEndpoint(service),
		StatusEndpoint:        MarkStatusEndpoint(service),
		ServiceStatusEndpoint: MarkServiceStatusEndpoint(service),
		WatermarkEndpoint:     MarkWatermarkEndpoint(service),
	}
}

func MarkGetEndpoint(service watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := service.Get(ctx, req.Filters...)
		if err != nil {
			return GetResponse{Documents: docs, Err: err.Error()}, nil
		}

		return GetResponse{Documents: docs, Err: ""}, nil
	}
}

func MarkAddDocumentEndpoint(service watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddDocumentRequest)
		ticketID, err := service.AddDocument(ctx, req.Document)
		if err != nil {
			return AddDocumentResponse{TicketID: ticketID, Err: err.Error()}, nil
		}

		return AddDocumentResponse{TicketID: ticketID, Err: ""}, nil
	}

}

func MarkStatusEndpoint(service watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(StatusRequest)
		status, err := service.Status(ctx, req.TicketID)
		if err != nil {
			return StatusResponse{Status: status, Err: err.Error()}, nil
		}
		return StatusResponse{Status: status, Err: ""}, nil
	}
}

func MarkWatermarkEndpoint(service watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WatermarkRequest)
		code, err := service.Watermark(ctx, req.TicketID, req.Mark)
		if err != nil {
			return WatermarkResponse{Code: code, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: code, Err: ""}, nil
	}
}

func MarkServiceStatusEndpoint(service watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		code, err := service.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: err.Error()}, nil

	}
}

func (set *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	response, err := set.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}
	getResponse := response.(GetResponse)

	if getResponse.Err != "" {
		return []internal.Document{}, errors.New(getResponse.Err)
	}

	return getResponse.Documents, nil
}

func (set *Set) ServiceStatus(ctx context.Context) (int, error) {
	response, err := set.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	serviceStatusResponse := response.(ServiceStatusResponse)
	if err != nil {
		return serviceStatusResponse.Code, err
	}

	if serviceStatusResponse.Err != "" {
		return serviceStatusResponse.Code, errors.New(serviceStatusResponse.Err)
	}

	return serviceStatusResponse.Code, nil
}

func (set *Set) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	response, err := set.StatusEndpoint(ctx, StatusRequest{TicketID: ticketID})
	statusResponse := response.(StatusResponse)
	if err != nil {
		return statusResponse.Status, err
	}

	if statusResponse.Err != "" {
		return statusResponse.Status, errors.New(statusResponse.Err)
	}

	return statusResponse.Status, nil
}

func (set *Set) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	response, err := set.AddDocumentEndpoint(ctx, AddDocumentRequest{Document: doc})
	addDocumentResponse := response.(AddDocumentResponse)
	if err != nil {
		return addDocumentResponse.TicketID, err
	}

	if addDocumentResponse.Err != "" {
		return addDocumentResponse.TicketID, errors.New(addDocumentResponse.Err)
	}

	return addDocumentResponse.TicketID, nil
}

func (set *Set) Watermakr(ctx context.Context, ticketID, mark string) (int, error) {
	response, err := set.WatermarkEndpoint(ctx, WatermarkRequest{TicketID: ticketID, Mark: mark})
	watermarkResponse := response.(WatermarkResponse)
	if err != nil {
		return watermarkResponse.Code, err
	}

	if watermarkResponse.Err != "" {
		return watermarkResponse.Code, errors.New(watermarkResponse.Err)
	}

	return watermarkResponse.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
