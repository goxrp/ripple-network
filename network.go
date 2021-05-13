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
	JsonRpcUrl   string
	WebSocketUrl string
	Notes        string
}

// Servers provides the server info available at:
// https://xrpl.org/get-started-with-the-rippled-api.html
func Servers() []Server {
	return []Server{
		{
			Hostname:     "xrpl.ws",
			Operator:     "XRP Ledger Foundation",
			Network:      NetworkMainnet,
			JsonRpcUrl:   "https://xrpl.ws/",
			WebSocketUrl: "wss://xrpl.ws/",
			Notes:        "Full history server cluster."},
		{
			Hostname:     "s1.ripple.com",
			Operator:     "Ripple",
			Network:      NetworkMainnet,
			JsonRpcUrl:   "https://s1.ripple.com:51234/",
			WebSocketUrl: "wss://s1.ripple.com/",
			Notes:        "General purpose server cluster"},
		{
			Hostname:     "s2.ripple.com",
			Operator:     "Ripple",
			Network:      NetworkMainnet,
			JsonRpcUrl:   "https://s2.ripple.com:51234/",
			WebSocketUrl: "wss://s2.ripple.com/",
			Notes:        "Full-history server cluster"},
		{
			Hostname:     "s.altnet.rippletest.net",
			Operator:     "Ripple",
			Network:      NetworkTestnet,
			JsonRpcUrl:   "https://s.altnet.rippletest.net:51234/",
			WebSocketUrl: "wws://s.altnet.rippletest.net/",
			Notes:        "Testnet public server"},
		{
			Hostname:     "s.devnet.rippletest.net",
			Operator:     "Ripple",
			Network:      NetworkDevnet,
			JsonRpcUrl:   "https://s.devnet.rippletest.net:51234/",
			WebSocketUrl: "wss://s.devnet.rippletest.net/",
			Notes:        "Full-history server cluster"},
	}
}

var serversJsonRpc = openapi3.Servers{
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

func ServersJsonRpc() openapi3.Servers {
	return serversJsonRpc
}

func GetJsonRpcUrl() string {
	return GetMainnetPublicJsonRpcUrl()
}

func GetMainnetPublicJsonRpcUrl() string {
	return serversJsonRpc[0].URL
}

func GetMainnetFullHistoryJsonRpcUrl() string {
	return serversJsonRpc[1].URL
}
