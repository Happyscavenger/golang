package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func cli(url string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(request)
	defer response.Body.Close()
	sHtml, _ := ioutil.ReadAll(response.Body)
	html := string(sHtml)
	return html
}

func main() {
	url := "http://www.fangdi.com.cn/service/freshHouse/getHouseDetail.action?houseID=a839ffa2eb459118da146d51e6433842"
	html := cli(url)
	info := make(map[interface{}]interface{})
	houseDetailList := gjson.Get(html, "houseDetailList.0")
	if houseDetailList.Exists() {
		info["room_number"] = houseDetailList.Get("room_number").String()
		info["priv_flarea"] = houseDetailList.Get("priv_flarea").String()
		info["land_use"] = houseDetailList.Get("land_use").String()
		info["cellar_area"] = houseDetailList.Get("cellar_area").String()
		info["plan_co_flarea"] = houseDetailList.Get("plan_co_flarea").String()
		info["status"] = houseDetailList.Get("status").String()
		info["flarea"] = houseDetailList.Get("flarea").String()
		info["flat_style"] = houseDetailList.Get("flat_style").String()
		info["plan_priv_flarea"] = houseDetailList.Get("plan_priv_flarea").String()
		info["plan_cellar_area"] = houseDetailList.Get("plan_cellar_area").String()
		info["co_flarea"] = houseDetailList.Get("co_flarea").String()
		info["floor_name"] = houseDetailList.Get("floor_name").String()
		info["plan_flarea"] = houseDetailList.Get("plan_flarea").String()
	}
	fmt.Println(info)
}
