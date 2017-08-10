//
//    爬取站点：
//    http://www.bd-film.com/zx/index.htm
//
//
// // // // // // // // // // // // // // // //

package spider

import "github.com/PuerkitoBio/goquery"

type DdFilm struct {
	Url string
}

func (st *DdFilm) GetImgUrl(sl *goquery.Selection, r Rules) string {

	url, ok := sl.Find(r.ImgRule).Eq(0).Attr("href")

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
func (st *DdFilm) GetName(sl *goquery.Selection, r Rules) string {

	str, ok := sl.Find(r.TextRule).Eq(0).Attr(r.Attr)

	if !ok {
		return ""
	}
	return str
}
