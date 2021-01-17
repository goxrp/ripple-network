package ripplenetwork

import (
	"github.com/getkin/kin-openapi/openapi3"
)

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
