package box

import (
	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/route"
)

type BoxEx struct {
	*Box

	routerEx adapter.RouterEx
}

func NewEx(options Options) (*BoxEx, error) {
	b, err := New(options)
	if err != nil {
		return nil, err
	}

	return &BoxEx{
		Box:      b,
		routerEx: &route.RouterEx{Router: b.router.(*route.Router)},
	}, nil
}

func (b *BoxEx) RouterEx() adapter.RouterEx {
	return b.routerEx
}

func (b *BoxEx) LogFactory() log.Factory {
	return b.logFactory
}