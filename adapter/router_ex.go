package adapter

import (
	"context"
	"net/netip"

	"github.com/sagernet/sing-box/common/geoip"
	"github.com/sagernet/sing-box/option"
	dns "github.com/sagernet/sing-dns"
	tun "github.com/sagernet/sing-tun"
	"github.com/sagernet/sing/common/control"

	mdns "github.com/miekg/dns"
)

type RouterEx interface {
	Service

	Outbounds() []Outbound
	Outbound(tag string) (Outbound, bool)
	DefaultOutbound(network string) Outbound

	FakeIPStore() FakeIPStore

	ConnectionRouter

	GeoIPReader() *geoip.Reader
	LoadGeosite(code string) (Rule, error)

	Exchange(ctx context.Context, message *mdns.Msg) (*mdns.Msg, error)
	Lookup(ctx context.Context, domain string, strategy dns.DomainStrategy) ([]netip.Addr, error)
	LookupDefault(ctx context.Context, domain string) ([]netip.Addr, error)
	ClearDNSCache()

	InterfaceFinder() control.InterfaceFinder
	UpdateInterfaces() error
	DefaultInterface() string
	AutoDetectInterface() bool
	AutoDetectInterfaceFunc() control.Func
	DefaultMark() int
	NetworkMonitor() tun.NetworkUpdateMonitor
	InterfaceMonitor() tun.DefaultInterfaceMonitor
	PackageManager() tun.PackageManager
	Rules() []Rule

	ClashServer() ClashServer
	SetClashServer(server ClashServer)

	V2RayServer() V2RayServer
	SetV2RayServer(server V2RayServer)

	ResetNetwork() error

	// for v2bx
	AddInbound(in Inbound) error
	DelInbound(tag string) error
	UpdateDnsRules(rules []option.DNSRule) error
}
