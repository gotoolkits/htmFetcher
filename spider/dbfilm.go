//
//    爬取站点：
//    http://www.bd-film.com/zx/index.htm
//
//    通过访问链接URL获取信息,下载地址为js动态生成，需使用
//    phantomjs方式获取
// // // // // // // // // // // // // // // //

package spider

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DdFilm struct {
	Url string
}

func (dd *DdFilm) NewRules() Rules {

	return Rules{
		"#content>div>table>tbody>tr>td",
		Rule{"", "#wx_pic > img.pic", "src"},
		Rule{"", "a:nth-child(3)", "title"},
		Rule{"", "#bs-docs-download > table > tbody > tr > td.bd-address > div > a", "href"}}

}

func (dd *DdFilm) GetImgUrl(sl *goquery.Selection, r Rules) string {

	url, _ := sl.Find("a:nth-child(3)").Eq(0).Attr("href")

	dd.Url = url

	doc, _ := goquery.NewDocument(url)
	str, _ := doc.Find(r.ImgRule.Ru).Eq(0).Attr(r.ImgRule.Attr)

	return str
}

func (dd *DdFilm) GetName(sl *goquery.Selection, r Rules) string {

	str, _ := sl.Find(r.TextRule.Ru).Eq(0).Attr(r.TextRule.Attr)
	return str
}

func (dd *DdFilm) GetDownload(sl *goquery.Selection, r Rules) []string {

	// s, _ := sub.GetAttr(r.Download.Ru, r.Download.Attr)
	var sub Spider = Spider{}
	var s []string

	sub.doc = sub.LoadPtJsUrl(dd.Url)

	htm, _ := sub.GetHtml("#bs-docs-download > table > tbody > tr:nth-child(2) > td.bd-address")

	for _, v := range htm {

		if strings.Contains(v, "pan.baidu.com") {

			site, _ := sub.doc.Find("#bs-docs-download > table > tbody > tr:nth-child(2) > td.bd-address > div > a").Eq(0).Attr("href")
			pwd := sub.doc.Find("#bs-docs-download > table > tbody > tr:nth-child(2) > td.bd-address > div > span").Eq(0).Text()
			s = append(s, site+"#"+pwd)

			continue
		}
		site, _ := sub.doc.Find("#bs-docs-download > table > tbody > tr:nth-child(2) > td.bd-address > div > a").Eq(0).Attr("href")
		s = append(s, site)
	}

	//fmt.Println(s)

	return s
}
