package spider

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gotoolkits/htmFetcher/conv"
	"github.com/gotoolkits/htmFetcher/go-phantomjs-fetcher"
)

type MovInfo interface {
	NewRules() Rules
	GetImgUrl(sl *goquery.Selection, r Rules) string
	GetName(sl *goquery.Selection, r Rules) string
	GetDownload(sl *goquery.Selection, r Rules) []string
}

// Spider
type Spider struct {
	Url string // page that spider would deal with
	doc *goquery.Document
}

type Rules struct {
	//	Name     MovInfo
	HtmlRule string
	ImgRule  Rule
	TextRule Rule
	Download Rule
}
type Rule struct {
	Url  string
	Ru   string
	Attr string
}

type MovStor struct {
	Site   string    `json:"site"`
	Lenght int       `json:"lenght"`
	List   []MovAttr `json:"list"`
}

type MovAttr struct {
	ImgUrl   string   `json:"imgUrl"`
	Info     string   `json:"filmName"`
	Download []string `json:"Download"`
}

// Start spider
func CreateSpiderFromUrl(url string) (*Spider, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, fmt.Errorf("url %s, error %s", url, err)
	}
	return &Spider{Url: url, doc: doc}, nil
}

func CreateSpiderFromResponse(r *http.Response) (*Spider, error) {
	doc, err := goquery.NewDocumentFromResponse(r)
	if err != nil {
		return nil, fmt.Errorf("error %s", err)
	}
	return &Spider{doc: doc}, nil
}

func (s *Spider) GetHtml(rule string) ([]string, error) {
	var (
		res = make([]string, 0) //for leaf
		wg  sync.WaitGroup
		mu  sync.Mutex
	)

	s.doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			content, _ := sl.Html()
			//sc := StrConvUTF8(content)
			mu.Lock()
			res = append(res, content)
			mu.Unlock()

		}()
	})
	wg.Wait()
	return res, nil
}

func (s *Spider) GetText(rule string) ([]string, error) {
	var (
		res = make([]string, 0) //for leaf
		wg  sync.WaitGroup
		mu  sync.Mutex
	)

	s.doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			res = append(res, sl.Text())
			mu.Unlock()
		}()
	})
	wg.Wait()
	return res, nil
}

func (s *Spider) GetAttr(rule, attr string) ([]string, error) {
	var (
		res = make([]string, 0) //for leaf
		wg  sync.WaitGroup
		mu  sync.Mutex
	)

	s.doc.Find(rule).Each(func(ix int, sl *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			attr, ok := sl.Attr(attr)
			if ok {
				mu.Lock()
				res = append(res, attr)
				mu.Unlock()
			}
		}()
	})
	wg.Wait()
	return res, nil
}

// 增加多级过滤
func (s *Spider) GetMovAttr(i MovInfo) ([]MovAttr, error) {
	var (
		res = []MovAttr{}
		ma  = MovAttr{}
		wg  sync.WaitGroup
		mu  sync.Mutex
		//ok  bool
	)
	r := i.NewRules()

	start := time.Now()
	s.doc.Find(r.HtmlRule).Each(func(ix int, sl *goquery.Selection) {

		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()

			ma.ImgUrl = i.GetImgUrl(sl, r)
			if ma.ImgUrl != "" {
				ma.Info = i.GetName(sl, r)
				ma.Download = i.GetDownload(sl, r)
				res = append(res, ma)
			}
			mu.Unlock()
		}()
	})
	wg.Wait()
	stop := time.Now()
	t := stop.Sub(start)

	fmt.Println("Spend time Sec:", t.Seconds())
	return res, nil
}

// 增加phantomjs获取url
func (s *Spider) LoadPtJsUrl(url string) *goquery.Document {

	fetcher, err := phantomjs.NewFetcher(2017, nil)
	defer fetcher.ShutDownPhantomJSServer()
	if err != nil {
		panic(err)
	}

	// js_script := "function(){document.getElementById('bs-docs-download').onclik();}"
	// js_run_at := phantomjs.RUN_AT_DOC_END

	resp, err := fetcher.GetWithJS(url, "", "")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Content))
	if err != nil {
		panic(err)
	}

	return doc

}

// 增加GBK编码
func StrConvGBK(str string) string {

	decode := conv.NewDecoder("GBK")
	if decode == nil {
		fmt.Errorf("Could not create decoder for %s", "utf-8")
		return "NULL"
	}

	r := decode.ConvertString(str)
	return r

}

func StrConvUTF8(str string) string {

	decode := conv.NewDecoder("UTF-8")
	if decode == nil {
		fmt.Errorf("Could not create decoder for %s", "utf-8")
		return "NULL"
	}

	r := decode.ConvertString(str)
	return r

}
