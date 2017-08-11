//
//    爬取站点：
//    http://img.piaowu99.com/
//
//
// // // // // // // // // // // // // // // //

package spider

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	PH_SITE = "http://img.piaowu99.com"
)

type PiaoHua struct {
	Url string
}

func (ph *PiaoHua) NewRules() Rules {

	return Rules{
		"#iml1 > ul:nth-child(2) > li",
		Rule{"", "a.img > img", "src"},
		Rule{"", "a:nth-child(2) > strong > font > font", ""},
		Rule{"a", "#showinfo > table > tbody > tr > td", "href"}}

}

func (ph *PiaoHua) GetImgUrl(sl *goquery.Selection, r Rules) string {
	url, _ := sl.Find(r.ImgRule.Ru).Eq(0).Attr(r.ImgRule.Attr)
	return url
}

func (ph *PiaoHua) GetName(sl *goquery.Selection, r Rules) string {
	str := sl.Find(r.TextRule.Ru).Eq(0).Text()
	return str
}

func (ph *PiaoHua) GetDownload(sl *goquery.Selection, r Rules) []string {

	url, ok := sl.Find(r.Download.Url).Eq(0).Attr(r.Download.Attr)
	url = PH_SITE + url

	if !ok {
		return nil
	}

	doc, _ := goquery.NewDocument(url)
	str, _ := doc.Find(r.Download.Ru).Eq(0).Find(r.Download.Url).Attr(r.Download.Attr)

	s := strings.Split(str, "\n")
	return s
}
