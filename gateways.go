package ripplenetwork

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	gorippled "github.com/goxrp/go-rippled"
	"github.com/grokify/gohttp/httpsimple"
)

type Account struct {
	Name                        string                             `json:"name,omitempty"`
	Description                 string                             `json:"description,omitempty"`
	Address                     string                             `json:"address,omitempty"`
	URL                         string                             `json:"url,omitempty"`
	FeesURL                     string                             `json:"feesURL,omitempty"`
	CoinistURL                  string                             `json:"coinistURL,omitempty"`
	AccountRoot                 *gorippled.AccountRoot             `json:"accountRoot,omitempty"`
	AccountRootRetrieved        time.Time                          `json:"accountRootRetrieved,omitempty"`
	AccountCurriencies          *gorippled.AccountCurrenciesResult `json:"accountCurrencies,omitempty"`
	AccountCurrienciesRetrieved time.Time                          `json:"accountCurrenciesRetrieved,omitempty"`
}

func (acct *Account) LoadAccountRoot() error {
	acct.Address = strings.TrimSpace(acct.Address)
	if len(acct.Address) == 0 {
		return errors.New("account has no address")
	}

	req := httpsimple.SimpleRequest{
		Method: http.MethodPost,
		URL:    servers[0].JsonRpcUrl,
		Body:   gorippled.AccountInfoRequestSimple(acct.Address),
		IsJSON: true}

	var acctInfo gorippled.AccountInfoResponse

	sc := httpsimple.NewSimpleClient(nil, servers[0].JsonRpcUrl)

	_, resp, err := sc.DoJSON(req, &acctInfo)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("rippled response status_code [%d]", resp.StatusCode)
	}

	acct.AccountRoot = &acctInfo.Result.AccountData
	acct.AccountRootRetrieved = time.Now().UTC()
	return nil
}

func (acct *Account) LoadAccountCurrencies() error {
	acct.Address = strings.TrimSpace(acct.Address)
	if len(acct.Address) == 0 {
		return errors.New("account has no address")
	}

	req := httpsimple.SimpleRequest{
		Method: http.MethodPost,
		URL:    servers[0].JsonRpcUrl,
		Body:   gorippled.AccountCurrenciesRequestSimple(acct.Address),
		IsJSON: true}

	var res gorippled.AccountCurrenciesResponse

	sc := httpsimple.NewSimpleClient(nil, servers[0].JsonRpcUrl)

	_, resp, err := sc.DoJSON(req, &res)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("rippled response status_code [%d]", resp.StatusCode)
	}

	acct.AccountCurriencies = &res.Result
	acct.AccountCurrienciesRetrieved = time.Now().UTC()
	return nil
}

// https://bitcoin.stackexchange.com/questions/16936/is-there-a-comprehensive-list-of-gateways-ious-and-fees-associated-with-using-t
// https://bitcointalk.org/index.php?topic=155236.msg1646402#msg1646402
// https://coinist.co/ripple/gateways/

type GatewaysSet struct {
	TimeCreated time.Time `json:"timeCreated"`
	Gateways    []Account `json:"gateways"`
}

var gateways = []Account{
	{
		Name:    "Bitstamp",
		Address: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		URL:     "https://www.bitstamp.net/",
	},
	{
		Name:       "Coinex",
		Address:    "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
		URL:        "http://www.coinex.co.nz/",
		FeesURL:    "https://www.coinex.co.nz/Fees",
		CoinistURL: "https://coinist.co/ripple/gateways/coinex/",
	},
	{
		Name:        "Devcoin",
		Description: "emfox (https://bitcointalk.org/index.php?action=profile;u=87884). 10 DVC redemption fee (cost of devcoin transaction fee)",
		Address:     "r3gHXhK1pwZFG9ESiaosxjufEVQjwGuJUd",
		URL:         "http://ripple.d.evco.in/",
		FeesURL:     "https://www.dividendrippler.com/fees",
	},
	{
		Name:    "Dividend Rippler",
		Address: "rfYv1TXnwgDDK4WQNbFALykYuEBnrR4pDX",
		URL:     "https://www.dividendrippler.com/",
		FeesURL: "https://www.dividendrippler.com/fees",
	},
	{
		Name:        "Goodwill LETS",
		Description: "Bob \"Red\" Way",
		Address:     "rGDWKWni6exeneJdNbEZ3nVX3Rrw5VG1p1",
		URL:         "https://ripple.com/forum/viewtopic.php?f=1&t=2895",
	},
	{
		Name:    "ICE",
		Address: "r4H3F9dDaYPFwbrUsusvNAHLz2sEZk4wE5",
		URL:     "https://ripple.com/forum/viewtopic.php?f=1&t=3958",
	},
	{
		Name:    "Justcoin",
		Address: "rJHygWcTLVpSXkowott6kzgZU6viQSVYM1",
		URL:     "https://justcoin.com/",
	},
	{
		Name:    "Peercover",
		Address: "ra9eZxMbJrUcgV8ui7aPc161FgrqWScQxV",
		URL:     "https://www.peercover.com/",
	},
	{
		Name:    "Ripple Israel",
		Address: "rNPRNzBB92BVpAhhZr4iXDTveCgV5Pofm9",
		URL:     "http://rippleisrael.co.il/",
	},
	{
		Name:    "Ripple Union",
		Address: "r3ADD8kXSUKHd6zTCKfnKT3zV9EZHjzp1S",
		URL:     "http://rippleunion.com/",
	},
	{
		Name:    "RippleChina",
		Address: "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA",
		URL:     "http://wg.iripplechina.com/",
	},
	{
		Name:    "RippleCN",
		Address: "rnuF96W4SZoCJmbHYBFoJZpR8eCaxNvekK",
		URL:     "http://ripplecn.com/gateway/",
	},
	{
		Name:    "RippleFund",
		Address: "rE7CNMbxwvTQrqSEjbcXCrqeN6EZga1ApU",
		URL:     "http://ripplefund.me/",
	},
	{
		Name:        "SnapSwap",
		Description: "https://btc2ripple.com/",
		Address:     "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
		URL:         "https://www.snapswap.us/",
	},
	{
		Name:    "TTBit",
		Address: "rGwUWgN5BEg3QGNY3RX2HfYowjUTZdid3E",
		URL:     "https://bitcointalk.org/index.php?topic=149533.0",
	},
	{
		Name:    "WisePass",
		Address: "rPDXxSZcuVL3ZWoyU82bcde3zwvmShkRyF",
		URL:     "https://wisepass.com/",
	},
	{
		Name:        "XRP Interchange",
		Description: "deweller/wellerco https://xrptalk.org/user/803-deweller/, https://bitcointalk.org/index.php?action=profile;u=155537",
		Address:     "rEUNavLM2NPpLW8dKqk3BRhXcZk5Hv7dhj",
		URL:         "http://xrpio.com/",
	},
}

func GatewaysRebuild() ([]Account, error) {
	newGateways := []Account{}
	for _, gw := range gateways {
		err := gw.LoadAccountRoot()
		if err != nil {
			return gateways, err
		}
		err = gw.LoadAccountCurrencies()
		if err != nil {
			return gateways, err
		}
		newGateways = append(newGateways, gw)
	}
	gateways = newGateways
	return gateways, nil
}

func Gateways(inflate bool) ([]Account, error) {
	if inflate {
		bytes, err := Asset("gateways.json")
		if err != nil {
			return []Account{}, err
		}
		gwSet := GatewaysSet{}
		err = json.Unmarshal(bytes, &gwSet)
		if err == nil {
			gateways = gwSet.Gateways
		}
		return gwSet.Gateways, err
	}
	return gateways, nil
}
