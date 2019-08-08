package main

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	//"golang.org/x/tools/go/ssa/interp/testdata/src/time"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func fetch(url string) *html.Node {
	log.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.87 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("http get err", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("http status code", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func parseUrls(url string, wg *sync.WaitGroup) { //ch chan bool,
	doc := fetch(url)
	nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
	for _, node := range nodes {
		url := htmlquery.FindOne(node, `./a/@href`)
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
		//log.Println(strings.Split(htmlquery.InnerText(url), "/")[4],
		//	htmlquery.InnerText(title))
		log.Println(htmlquery.InnerText(url),
			htmlquery.InnerText(title))
	}
	time.Sleep(2 * time.Second)
	//ch <- true
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	//ch := make(chan bool)
	for i := 1; i < 11; i++ {
		wg.Add(1)
		go parseUrls("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), &wg) //ch
	}
	wg.Wait()
	//for i := 0; i < 10; i++ {
	//	<-ch
	//}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
