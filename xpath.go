package main

import (
	"fmt"
	"github.com/lestrrat/go-libxml2"
	"log"
	"net/http"
)

type MyResponse struct {
	*http.Response
}

func (response *MyResponse) String() string {
	return fmt.Sprint(response.Response)
}

func request(method string, url string) *MyResponse {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	//request.Header.Add 可以添加多个header选项
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	my_response := &MyResponse{response}
	return my_response

}

func main() {
	response := request("GET", "https://sh.fang.ke.com/loupan")
	if doc, err := libxml2.ParseHTMLReader(response.Body); err != nil {
		log.Fatal(err)
	} else {
		defer doc.Free()
		nodes, err := doc.Find("//div/ul[@class='resblock-list-wrapper']/li/a/@href")
		if err != nil {
			log.Fatal(err)
		} else {
			//fmt.Println(nodes.NodeList()[0].TextContent())
			for _, href := range nodes.NodeList() {
				fmt.Println(href.TextContent())
			}
		}
	}
	defer response.Body.Close()

}
