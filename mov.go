package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"github.com/PuerkitoBio/goquery"
	"github.com/gotoolkits/htmFetcher/spider"
)

type siteMatch struct {
	url string
	r   spider.Rules
}

func ErrCheck(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func main() {
	var stor spider.MovStor = spider.MovStor{}
	var ss []spider.MovStor

	//获取爬取站点列表与匹配规则
	sm := getSiteList()

	//迭代List爬取内容
	for _, v := range sm {

		s, err := spider.CreateSpiderFromUrl(v.url)
		ErrCheck(err, "create spider from url failed!")

		htm, err := s.GetMovAttr(v.r)
		ErrCheck(err, "get html failed!")

		stor.Site = v.url
		stor.List = htm
		stor.Lenght = len(htm)

		ss = append(ss, stor)

	}

	j, _ := json.MarshalIndent(ss, "", "  ")
	fmt.Println(string(j))

}

// 增加需要爬取的站点地址与DOM匹配规则
func getSiteList() []siteMatch {

	var list []siteMatch

	// add the site matcher info to list
	var (
		//site01
		name1 spider.DyGang = spider.DyGang{}
		r1    spider.Rules  = spider.Rules{
			&name1,
			"html>body>table>tbody>tr>td>div#tl>div#tab1_div_0>table>tbody>tr>td>table>tbody>tr>td>table>tbody>tr>td",
			"table>tbody>tr>td>a>img",
			"a.c2",
			"src",
			spider.Rule{}}
		sm1 siteMatch = siteMatch{"http://www.dygang.net/", r1}

		//site02
		name2 spider.DdFilm = spider.DdFilm{}
		sub2  spider.Rule   = spider.Rule{
			"#wx_pic > img.pic",
			"src"}
		r2 spider.Rules = spider.Rules{
			&name2,
			"#content>div>table>tbody>tr>td",
			"a:nth-child(3)",
			"a:nth-child(3)",
			"title",
			sub2}
		sm2 siteMatch = siteMatch{"http://www.bd-film.com/zx/index.htm", r2}

		// add others site
		//
		//

	)

	list = append(list, sm1)
	list = append(list, sm2)
	// list = append(list, X)

	return list
}
