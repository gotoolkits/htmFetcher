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

	//迭代爬取内容
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
		r1 spider.Rules = spider.Rules{
			"html>body>table>tbody>tr>td>div#tl>div#tab1_div_0>table>tbody>tr>td>table>tbody>tr>td>table>tbody>tr>td",
			"table>tbody>tr>td>a>img",
			"a.c2",
			"src"}
		sm1 siteMatch = siteMatch{"http://www.dygang.net/", r1}
	)
	list = append(list, sm1)

	// add others
	//
	//

	return list
}
