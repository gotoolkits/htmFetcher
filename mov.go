package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gotoolkits/htmFetcher/spider"
)

type siteMatch struct {
	url  string
	Name spider.MovInfo
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

		htm, err := s.GetMovAttr(v.Name)
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
		sm1 siteMatch = siteMatch{"http://www.dygang.net/", &spider.DyGang{}}
		//sm2 siteMatch = siteMatch{"http://www.bd-film.com/zx/index.htm", &spider.DdFilm{}}
		sm3 siteMatch = siteMatch{"http://img.piaowu99.com/", &spider.PiaoHua{}}

		// add others site
		//
		//

	)

	list = append(list, sm1)
	//list = append(list, sm2)
	list = append(list, sm3)

	return list
}
