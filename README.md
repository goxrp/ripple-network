# Ripple Network

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This package provides a list of public Rippled servers and Gateways for machine accessible use in CSV, JSON and Go library formats.

A few gateway lists exist online but are not in formats suitable for programmtic access. See the following thread for more info:

https://bitcoin.stackexchange.com/questions/16936/is-there-a-comprehensive-list-of-gateways-ious-and-fees-associated-with-using-t

## Gateways

Known gateways are listed in the [`gateways.csv`](gateways.csv) and [`gateways.json`](gateways.json) files. This is is built from online sources and periodically updated with AccountRoot and AccountCurrencies info.

Gateway information can be programmatically accessed by calling the `Gateways(inflate bool)` function.

Gateway information can be loaded with up to date AccountRoot and AccountCurrencies data by calling `GatewaysRebuild()`.

This is a [the CSV as hosted on GitHub](gateways.csv).

[![](gateways.png)](gateways.csv)

## Rippled Servers

Rippled Servers can be retrieved programmaticaly by calling the `Servers()` function. Currently, this provides programmatic access to the server list hosted on https://xrpl.org.

| Operator              | Network | JSON-RPC URL        | WebSocket URL                  | Notes |
| --------------------- | ------- | ---------------------------------------- | ------------- | -------------- |
| XRP Ledger Foundation | Mainnet | `https://xrpl.ws/`                       | `wss://xrpl.ws/`                 | Full history server cluster |
| Ripple | Mainnet                | `https://s1.ripple.com:51234/`           | `wss://s1.ripple.com/`           | General purpose server cluster |
| Ripple | Mainnet                | `https://s2.ripple.com:51234/`           | `wss://s2.ripple.com/`           | Full-history server cluster |
| Ripple | Testnet                | `https://s.altnet.rippletest.net:51234/` | `wss://s.altnet.rippletest.net/` | Testnet public server |
| Ripple | Mainnet                | `https://s.devnet.rippletest.net:51234/` | `wss://s.devnet.rippletest.net/` | Devnet public server |

### Testing

#### JSON RPC

Perform a quick test with the following curl commands. The [`rippled-postman`](https://github.com/goxrp/rippled-postman) Postman Collection can also be used with Postman.

```bash
$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://xrpl.ws/

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s1.ripple.com:51234/

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s2.ripple.com:51234/

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s.altnet.rippletest.net:51234/ --insecure

$ curl -H 'Content-Type: application/json' -d '{"method":"server_info"}' https://s.devnet.rippletest.net:51234/ --insecure
```

 [build-status-svg]: https://github.com/goxrp/ripple-network/actions/workflows/test.yaml/badge.svg
 [build-status-url]: https://github.com/goxrp/ripple-network/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/goxrp/ripple-network
 [goreport-url]: https://goreportcard.com/report/github.com/goxrp/ripple-network
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/goxrp/ripple-network
 [docs-godoc-url]: https://pkg.go.dev/github.com/goxrp/ripple-network
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/goxrp/ripple-network/blob/master/LICENSE
