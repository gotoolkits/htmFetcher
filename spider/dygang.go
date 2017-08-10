//
//    爬取站点：
//    http://www.dygang.net/
//
//
// // // // // // // // // // // // // // // //

package spider

import "github.com/PuerkitoBio/goquery"

type DyGang struct {
	Url string
}

func (st *DyGang) GetImgUrl(sl *goquery.Selection, r Rules) string {

	url, ok := sl.Find(r.ImgRule).Eq(0).Attr(r.Attr)

	if !ok {
		return ""
	}

	if r.Sub.R != "" {
		sp, _ := CreateSpiderFromUrl(url)
		str, _ := sp.GetAttr(r.Sub.R, r.Sub.Attr)
		return str[0]
	}
	return url

}
func (st *DyGang) GetName(sl *goquery.Selection, r Rules) string {

	str := StrConvGBK(sl.Find(r.TextRule).Eq(0).Text())

	return str
}
