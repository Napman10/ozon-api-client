package v2

import (
	"github.com/andmetoo/ozon-api-client/ozon/posting/v2/fbo"
	"net/http"
)

type SubRoutes struct {
	fbo *fbo.FBO
}

func (s SubRoutes) FBO() *fbo.FBO {
	return s.fbo
}

func New(
	h *http.Client,
	uri string,
) *Posting {
	return &Posting{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			fbo: fbo.New(h, uri+"/fbo"),
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
