package v3

import (
	"github.com/andmetoo/ozon-api-client/ozon/posting/v3/fbs"
	"net/http"
)

type SubRoutes struct {
	fbs *fbs.FBS
}

func (s SubRoutes) FBS() *fbs.FBS {
	return s.fbs
}

func New(
	h *http.Client,
	uri string,
) *Posting {
	return &Posting{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			fbs: fbs.New(h, uri+"/fbs"),
		},
	}
}

type Posting struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (f Posting) SubRoutes() *SubRoutes {
	return f.subRoutes
}
