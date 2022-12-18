package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/av153k/go-watermarker/internal/util"
	"github.com/av153k/go-watermarker/pkg/watermark/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func NewHttpHandler(endpoint endpoints.Set) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/health", httptransport.NewServer(
		endpoint.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))

	mux.Handle("/status", httptransport.NewServer(
		endpoint.StatusEndpoint, decodeHTTPStatusRequest, encodeResponse,
	))

	mux.Handle(
		"/addDocument",
		httptransport.NewServer(
			endpoint.AddDocumentEndpoint, decodeHTTPAddDocumentRequest, encodeResponse))

	mux.Handle(
		"/get", httptransport.NewServer(
			endpoint.GetEndpoint, decodeHTTPGetRequest, encodeResponse))

	mux.Handle(
		"/watermark", httptransport.NewServer(
			endpoint.WatermarkEndpoint, decodeHTTPWatermarkRequest, encodeResponse))

	return mux
}

func decodeHTTPGetRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var req endpoints.GetRequest
	if request.ContentLength == 0 {
		logger.Log("Get request with no body")
		return req, nil
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.StatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPWatermarkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.WatermarkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPAddDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddDocumentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
