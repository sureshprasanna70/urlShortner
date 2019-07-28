package main

import (
	"fmt"
	"strings"

	soup "github.com/anaskhan96/soup"
)

func checkErr(err error) {
	if err != nil {
		panic(err)

	}
}
func printPrice(trs []soup.Root) []priceStruct {
	prices := make([]priceStruct, len(trs)-2)
	i := 2
	for i = 2; i < len(trs); i++ {
		tds := trs[i].Children()
		prices[i-2] = priceStruct{
			date:  strings.TrimSpace(tds[1].FullText()),
			price: strings.TrimSpace(tds[5].FullText()),
		}
	}
	fmt.Println(prices)
	return prices

}

type priceStruct struct {
	date  string
	price string
}

func main() {
	visitURL := "https://www.livechennai.com/gold_silverrate.asp"
	htmlResponse, err := soup.Get(visitURL)
	checkErr(err)
	doc := soup.HTMLParse(htmlResponse)
	lastUpdate := doc.Find("p", "class", "mob-cont")
	fmt.Println(lastUpdate.FullText())
	priceTable := doc.FindAllStrict("table", "class", "table-price")
	//fmt.Println("Gold price")
	printPrice(priceTable[1].FindAll("tr"))
	//fmt.Println("Silver price")
	//printPrice(priceTable[2].FindAll("td"))
}
