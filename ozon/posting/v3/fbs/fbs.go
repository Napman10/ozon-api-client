package fbs

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
	"net/http"
)

type FBS struct {
	h   *http.Client
	uri string
}

func New(
	h *http.Client,
	uri string,
) *FBS {
	return &FBS{
		h:   h,
		uri: uri,
	}
}

func (f FBS) List(ctx context.Context, req *ListRequest) (*ListResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ListRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, f.uri+"/list", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ListRequest.NewRequest")
	}

	return request.Send[ListResponse](f.h, r, request.ContentTypeApplicationJson)
}
