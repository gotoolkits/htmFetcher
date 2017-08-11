//
//    爬取站点：
//    http://www.dygang.net/
//
//    GBK编码
// // // // // // // // // // // // // // // //

package spider

import (
	"github.com/PuerkitoBio/goquery"
)

type DyGang struct {
	Url string
}

func (dg *DyGang) NewRules() Rules {

	return Rules{
		"#tab1_div_0 > table > tbody > tr > td > table > tbody > tr > td",
		Rule{"", "table>tbody>tr>td>a>img", "src"},
		Rule{"", "a.c2", ""},
		Rule{"a.c2", "#dede_content > table > tbody > tr > td > a", "href"}}
}

func (dg *DyGang) GetImgUrl(sl *goquery.Selection, r Rules) string {
	url, _ := sl.Find(r.ImgRule.Ru).Eq(0).Attr(r.ImgRule.Attr)
	return url
}

func (dg *DyGang) GetName(sl *goquery.Selection, r Rules) string {
	str := sl.Find(r.TextRule.Ru).Eq(0).Text()

	str = StrConvGBK(str)
	return str
}

func (dg *DyGang) GetDownload(sl *goquery.Selection, r Rules) []string {
	url, ok := sl.Find(r.Download.Url).Eq(0).Attr(r.Download.Attr)
	if !ok {
		return nil
	}

	// doc, _ := goquery.NewDocument(url)
	// str := doc.Find(r.Download.Ru).Eq(0).Text()
	// s := strings.Split(str, "\n")

	sp, _ := CreateSpiderFromUrl(url)
	s, _ := sp.GetAttr(r.Download.Ru, r.Download.Attr)

	return s

}
