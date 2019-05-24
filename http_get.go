package main

import (
	"fmt"
	//"github.com/pquerna/ffjson/ffjson"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.fangdi.com.cn/service/freshHouse/getHouseDetail.action?houseID=a839ffa2eb459118da146d51e6433842"
	//resp,err :=http.Get(url)
	//sHtml,_ := ioutil.ReadAll(resp.Body)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(request)
	defer response.Body.Close()
	sHtml, _ := ioutil.ReadAll(response.Body) //返回的是byte数组
	//fmt.Println(string(sHtml))
	info := make(map[interface{}]interface{})
	houseDetailList := gjson.Get(string(sHtml), "houseDetailList.0")
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
	//info["room_number"] = gjson.Get(string(sHtml),"houseDetailList.0.room_number")
	//info["priv_flarea"] = gjson.Get(string(sHtml),"houseDetailList.0.priv_flarea")
	//info["land_use"] = gjson.Get(string(sHtml),"houseDetailList.0.land_use")
	//info["cellar_area"] = gjson.Get(string(sHtml),"houseDetailList.0.cellar_area")
	//info["plan_co_flarea"] = gjson.Get(string(sHtml),"houseDetailList.0.plan_co_flarea")
	//info["status"] = gjson.Get(string(sHtml),"houseDetailList.0.status")
	//info["flarea"] = gjson.Get(string(sHtml),"houseDetailList.0.flarea")
	//info["flat_style"] = gjson.Get(string(sHtml),"houseDetailList.0.flat_style")
	//info["plan_priv_flarea"] = gjson.Get(string(sHtml),"houseDetailList.0.plan_priv_flarea")
	//info["plan_cellar_area"] = gjson.Get(string(sHtml),"houseDetailList.0.plan_cellar_area")
	//info["co_flarea"] = gjson.Get(string(sHtml),"houseDetailList.0.co_flarea")
	//info["floor_name"] = gjson.Get(string(sHtml),"houseDetailList.0.floor_name")
	//info["plan_flarea"] = gjson.Get(string(sHtml),"houseDetailList.0.plan_flarea")
	fmt.Println(info)
}
