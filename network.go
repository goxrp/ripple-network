package ripplenetwork

import (
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	NetworkMainnet = "Mainnet"
	NetworkTestnet = "Testnet"
	NetworkDevnet  = "Devnet"
)

type Server struct {
	Hostname     string
	Operator     string
	Network      string
	JSONRPCURL   string
	WebSocketURL string
	Notes        string
}

// Servers provides the server info available at:
// https://xrpl.org/get-started-with-the-rippled-api.html
func Servers() []Server {
	return servers
}

var servers = []Server{
	{
		Hostname:     "xrpl.ws",
		Operator:     "XRP Ledger Foundation",
		Network:      NetworkMainnet,
		JSONRPCURL:   "https://xrpl.ws/",
		WebSocketURL: "wss://xrpl.ws/",
		Notes:        "Full history server cluster."},
	{
		Hostname:     "s1.ripple.com",
		Operator:     "Ripple",
		Network:      NetworkMainnet,
		JSONRPCURL:   "https://s1.ripple.com:51234/",
		WebSocketURL: "wss://s1.ripple.com/",
		Notes:        "General purpose server cluster"},
	{
		Hostname:     "s2.ripple.com",
		Operator:     "Ripple",
		Network:      NetworkMainnet,
		JSONRPCURL:   "https://s2.ripple.com:51234/",
		WebSocketURL: "wss://s2.ripple.com/",
		Notes:        "Full-history server cluster"},
	{
		Hostname:     "s.altnet.rippletest.net",
		Operator:     "Ripple",
		Network:      NetworkTestnet,
		JSONRPCURL:   "https://s.altnet.rippletest.net:51234/",
		WebSocketURL: "wws://s.altnet.rippletest.net/",
		Notes:        "Testnet public server"},
	{
		Hostname:     "s.devnet.rippletest.net",
		Operator:     "Ripple",
		Network:      NetworkDevnet,
		JSONRPCURL:   "https://s.devnet.rippletest.net:51234/",
		WebSocketURL: "wss://s.devnet.rippletest.net/",
		Notes:        "Full-history server cluster"},
}

var serversJSONPRC = openapi3.Servers{
	&openapi3.Server{
		URL:         "https://s1.ripple.com:51234",
		Description: "Mainnet Public Cluster"},
	&openapi3.Server{
		URL:         "https://s2.ripple.com:51234",
		Description: "Mainnet Full History Cluster"},
	&openapi3.Server{
		URL:         "https://s.altnet.rippletest.net:51234",
		Description: "Testnet Public Cluster"},
	&openapi3.Server{
		URL:         "https://s.devnet.rippletest.net:51234",
		Description: "Devnet Public Cluster"},
}

func ServersJSONRPC() openapi3.Servers {
	return serversJSONPRC
}

func GetJSONRPCURL() string {
	return GetMainnetPublicJSONRPCURL()
}

func GetMainnetPublicJSONRPCURL() string {
	return serversJSONPRC[0].URL
}

func GetMainnetFullHistoryJSONRPCURL() string {
	return serversJSONPRC[1].URL
}
