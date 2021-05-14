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
	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/simplego/type/stringsutil"
)

func main() {
	gateways, err := ripplenetwork.Gateways(true)
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
	err = ioutil.WriteFile("gateways.json", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	tbl := table.NewTable()
	tbl.Columns = []string{
		"Name", "URL", "Address", "TransferRate", "CurrenciesReceive", "CurrenciesSend", "TimeRetrieved",
	}

	for _, gw := range gateways {
		row := []string{
			gw.Name,
			gw.URL,
			gw.Address,
			strconv.Itoa(int(gw.AccountRoot.TransferRate)),
			strings.Join(stringsutil.SliceCondenseSpace(gw.AccountCurriencies.ReceiveCurrencies, true, true), ","),
			strings.Join(stringsutil.SliceCondenseSpace(gw.AccountCurriencies.SendCurrencies, true, true), ","),
			gwSet.TimeCreated.Format(time.RFC3339),
		}
		tbl.Records = append(tbl.Records, row)
	}

	err = tbl.WriteCSV("gateways.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
