package v5

import (
	"github.com/andmetoo/ozon-api-client/ozon/product/v5/info"
	"net/http"
)

type SubRoutes struct {
	info *info.Info
}

func (c SubRoutes) Info() *info.Info {
	return c.info
}

func New(
	h *http.Client,
	uri string,
) *Product {
	return &Product{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			info: info.New(h, uri+"/info"),
		},
	}
}

type Product struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c Product) SubRoutes() *SubRoutes {
	return c.subRoutes
}
