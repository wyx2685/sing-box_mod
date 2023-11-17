package route

import (
	"context"
	"errors"
	"sync"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/experimental/libbox/platform"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

type RouterEx struct {
	*Router

	actionLock sync.RWMutex
}

func NewRouterEx(ctx context.Context,
	logFactory log.Factory,
	options option.RouteOptions,
	dnsOptions option.DNSOptions,
	ntpOptions option.NTPOptions,
	inbounds []option.Inbound,
	platformInterface platform.Interface,
) (*RouterEx, error) {
	r, err := NewRouter(
		ctx,
		logFactory,
		options,
		dnsOptions,
		ntpOptions,
		inbounds,
		platformInterface,
	)
	if err != nil {
		return nil, err
	}
	return &RouterEx{
		Router: r,
	}, nil
}

func (r *RouterEx) AddInbound(inbound adapter.Inbound) error {
	r.actionLock.Lock()
	defer r.actionLock.Unlock()
	if _, ok := r.inboundByTag[inbound.Tag()]; ok {
		return errors.New("the inbound is exist")
	}
	r.inboundByTag[inbound.Tag()] = inbound
	return nil
}

func (r *RouterEx) DelInbound(tag string) error {
	r.actionLock.Lock()
	defer r.actionLock.Unlock()
	if _, ok := r.inboundByTag[tag]; ok {
		delete(r.inboundByTag, tag)
	} else {
		return errors.New("the inbound not have")
	}
	return nil
}

func (r *RouterEx) UpdateDnsRules(rules []option.DNSRule) error {
	dnsRules := make([]adapter.DNSRule, 0, len(rules))
	for i, rule := range rules {
		dnsRule, err := NewDNSRule(r, r.logger, rule)
		if err != nil {
			return E.Cause(err, "parse dns rule[", i, "]")
		}
		err = dnsRule.Start()
		if err != nil {
			return E.Cause(err, "initialize DNS rule[", i, "]")
		}
		dnsRules = append(dnsRules, dnsRule)
	}
	var tempRules []adapter.DNSRule
	r.actionLock.Lock()
	r.dnsRules = tempRules
	r.dnsRules = dnsRules
	r.actionLock.Unlock()
	for i, rule := range tempRules {
		err := rule.Close()
		if err != nil {
			return E.Cause(err, "closing DNS rule[", i, "]")
		}
	}
	return nil
}
