# Ripple Network

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This package provides a list of public Rippled servers and Gateways.

## Gateways

Known gateways are listed in the [`gateways.json`](gateways.json). This is is built from online sources and periodically updated with AccountRoot info using the script in [`cmd/build_gateways_data/main.go`](cmd/build_gateways_data/main.go).

## Servers

| Hostname                  | Description                  | JSON-RPC Port |
| ------------------------- | ---------------------------- | ------------- |
| `s1.ripple.com`           | Mainnet Public Cluster       | 51234 |
| `xrpl.ws`                 | Mainnet Full History Cluster |  |
| `s2.ripple.com`           | Mainnet Full History Cluster | 51234 |
| `s.altnet.rippletest.net` | Testnet Public Cluster       | 51234 |
| `s.devnet.rippletest.net` | Devnet Public Cluster        | 51234 |

## Testing

### JSON RPC

Test with the following:

```bash
$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s1.ripple.com:51234/

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s2.ripple.com:51234/

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s.altnet.rippletest.net:51234/ --insecure

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s.devnet.rippletest.net:51234/ --insecure
```

 [build-status-svg]: https://github.com/goxrp/ripple-network/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/goxrp/ripple-network/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/goxrp/ripple-network
 [goreport-url]: https://goreportcard.com/report/github.com/goxrp/ripple-network
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/goxrp/ripple-network
 [docs-godoc-url]: https://pkg.go.dev/github.com/goxrp/ripple-network
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/goxrp/ripple-network/blob/master/LICENSE