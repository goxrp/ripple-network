package ripplenetwork

type Account struct {
	Name       string `json:"name,omitempty"`
	Address    string `json:"address,omitempty"`
	URL        string `json:"url,omitempty"`
	FeesURL    string `json:"feesURL,omitempty"`
	CoinistURL string `json:"coinistURL,omitempty"`
}

// https://bitcoin.stackexchange.com/questions/16936/is-there-a-comprehensive-list-of-gateways-ious-and-fees-associated-with-using-t#:~:text=There%20is%20no%20official%20list,any%20kind%20of%20comprehensive%20list.
// https://bitcointalk.org/index.php?topic=155236.msg1646402#msg1646402
// https://coinist.co/ripple/gateways/

func getAccounts() []Account {
	return []Account{
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
			Name:    "RippleChina",
			Address: "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA",
			URL:     "http://wg.iripplechina.com/",
		},
		{
			Name:    "WisePass",
			Address: "rPDXxSZcuVL3ZWoyU82bcde3zwvmShkRyF",
			URL:     "https://wisepass.com/",
		},
	}
}
