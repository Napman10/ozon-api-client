package posting

import (
	v2 "github.com/andmetoo/ozon-api-client/ozon/posting/v2"
	v3 "github.com/andmetoo/ozon-api-client/ozon/posting/v3"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Posting {
	return &Posting{
		v2: v2.New(h, uri+"/v2/posting"),
		v3: v3.New(h, uri+"/v3/posting"),
	}
}

type Posting struct {
	v2 *v2.Posting
	v3 *v3.Posting
}

func (p *Posting) V2() *v2.Posting {
	return p.v2
}

func (p *Posting) V3() *v3.Posting {
	return p.v3
}
