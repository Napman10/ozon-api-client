package info

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Info {
	return &Info{
		h:   h,
		uri: uri,
	}
}

type Info struct {
	h   *http.Client
	uri string
}

func (c Info) Prices(ctx context.Context, req *PricesRequest) (*PricesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "PricesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/prices", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "PricesRequest.NewRequest")
	}

	return request.Send[PricesResponse](c.h, r, request.ContentTypeApplicationJson)
}
