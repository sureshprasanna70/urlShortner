package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	soup "github.com/anaskhan96/soup"
	chart "github.com/wcharczuk/go-chart"
)

func graphPlot(prices []chart.Value) {
	sbc := chart.BarChart{
		Title:      "Gold Prices",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis:    chart.StyleShow(),
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Bars: prices,
	}
	buffer := bytes.NewBuffer([]byte{})

	err := sbc.Render(chart.SVG, buffer)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("chart.svg", buffer.Bytes(), 0644)

}
func checkErr(err error) {
	if err != nil {
		panic(err)

	}
}
func barPrice(trs []soup.Root) []chart.Value {
	barValues := make([]chart.Value, len(trs)-2)
	i := 2
	for i = 2; i < len(trs); i++ {
		tds := trs[i].Children()
		trimmedString := strings.TrimSpace(tds[5].FullText())
		floatPrice, _ := strconv.ParseFloat(trimmedString, 64)
		barValues[i-2] = chart.Value{
			Label: strings.TrimSpace(tds[1].FullText()),
			Value: floatPrice,
		}
	}
	return barValues

}

func main() {
	visitURL := "https://www.livechennai.com/gold_silverrate.asp"
	htmlResponse, err := soup.Get(visitURL)
	checkErr(err)
	doc := soup.HTMLParse(htmlResponse)
	lastUpdate := doc.Find("p", "class", "mob-cont")
	fmt.Println(lastUpdate.FullText())
	priceTable := doc.FindAllStrict("table", "class", "table-price")
	prices := barPrice(priceTable[1].FindAll("tr"))
	fmt.Println(prices)
	graphPlot(prices)
}
