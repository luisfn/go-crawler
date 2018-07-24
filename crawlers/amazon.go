package crawlers

import (
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"encoding/json"
)

const (
	namePath     = "//span[@id='productTitle']"
	pricePath    = "//span[@id='priceblock_ourprice']"
	discountPath = "//*[@id='regularprice_savings']/td[2]"
)

type Amazon struct {
}

type Product struct {
	Name               string `json:"name"`
	Price              string `json:"price"`
	DiscountPrice      string `json:"discount_price"`
	DiscountPercentage string `json:"discount_percentage"`
}

func (p *Product) Json() string {
	json, _ := json.Marshal(p)

	return string(json)
}

func (a *Amazon) clearPrice(text string) string {
	regex := regexp.MustCompile(`\d{1,3}(?:[.,]\d{3})*(?:[.,]\d{2})`)
	raw := regex.FindStringSubmatch(text)[0]
	replacer := strings.NewReplacer(",", "", ".", "")
	replace := replacer.Replace(raw)


	return replace[:len(replace)-2] + "." + replace[len(replace)-2:]
}

func (a *Amazon) clearDiscount(text string) (string, string) {
	if text == "" {
		return "0", "0"
	}

	regex := regexp.MustCompile(`(\d{1,3}(?:[.,]\d{3})*(?:[.,]\d{2}))\s\((\d{1,2})`)
	matches := regex.FindStringSubmatch(text)

	discountValue := a.clearPrice(matches[1])
	discountPercentage := matches[2]

	return discountValue, discountPercentage
}

func (a *Amazon) find(doc *html.Node, path string) string {
	node := htmlquery.FindOne(doc, path)

	if node == nil {
		return ""
	}

	return strings.TrimSpace(htmlquery.InnerText(node))
}

func (a *Amazon) Scrap(doc *html.Node) *Product {
	discountPrice, discountPercentage := a.clearDiscount(a.find(doc, discountPath))

	return &Product{
		Name:               a.find(doc, namePath),
		Price:              a.clearPrice(a.find(doc, pricePath)),
		DiscountPrice:      discountPrice,
		DiscountPercentage: discountPercentage,
	}
}
