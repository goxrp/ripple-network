package ripplenetwork

type Account struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address,omitempty"`
	URL         string `json:"url,omitempty"`
	FeesURL     string `json:"feesURL,omitempty"`
	CoinistURL  string `json:"coinistURL,omitempty"`
}

// https://bitcoin.stackexchange.com/questions/16936/is-there-a-comprehensive-list-of-gateways-ious-and-fees-associated-with-using-t
// https://bitcointalk.org/index.php?topic=155236.msg1646402#msg1646402
// https://coinist.co/ripple/gateways/

func Gateways() []Account {
	return []Account{
		{
			Name:    "Bitstamp",
			Address: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			URL:     "https://www.bitstamp.net/",
		},
		{
			Name:       "Coinex",
			Address:    "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			URL:        "http://www.coinex.co.nz",
			FeesURL:    "https://www.coinex.co.nz/Fees",
			CoinistURL: "https://coinist.co/ripple/gateways/coinex/",
		},
		{
			Name:    "Dividend Rippler",
			Address: "rfYv1TXnwgDDK4WQNbFALykYuEBnrR4pDX",
			URL:     "https://www.dividendrippler.com/",
			FeesURL: "https://www.dividendrippler.com/fees",
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
}
