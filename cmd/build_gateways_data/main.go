package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	ripplenetwork "github.com/goxrp/ripple-network"
	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/mogo/type/stringsutil"
)

/*

Rebuild `bindata.go` with the following:

$ go-bindata -pkg ripplenetwork -ignore '(LICENSE|csv|go|md|mod|png|sum)$' .

*/

func main() {
	gateways, err := ripplenetwork.GatewaysRebuild()
	if err != nil {
		log.Fatal(err)
	}
	gwSet := ripplenetwork.GatewaysSet{
		Gateways:    gateways,
		TimeCreated: time.Now().UTC()}
	bytes, err := json.MarshalIndent(gwSet, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("gateways.json", bytes, 0600)
	if err != nil {
		log.Fatal(err)
	}

	tbl := table.NewTable("Gateways")
	tbl.Columns = []string{
		"Name", "URL", "Address", "TransferRate", "CurrenciesReceive", "CurrenciesSend", "TimeRetrieved",
	}

	for _, gw := range gateways {
		row := []string{
			gw.Name,
			gw.URL,
			gw.Address,
		}
		if gw.AccountRoot != nil {
			row = append(row, strconv.Itoa(int(gw.AccountRoot.TransferRate)))
		} else {
			row = append(row, "")
		}
		if gw.AccountCurriencies != nil {
			row = append(row,
				strings.Join(stringsutil.SliceCondenseSpace(gw.AccountCurriencies.ReceiveCurrencies, true, true), ","),
				strings.Join(stringsutil.SliceCondenseSpace(gw.AccountCurriencies.SendCurrencies, true, true), ","),
			)
		} else {
			row = append(row, "", "")
		}
		row = append(row, gwSet.TimeCreated.Format(time.RFC3339))
		tbl.Rows = append(tbl.Rows, row)
	}

	err = tbl.WriteCSV("gateways.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
