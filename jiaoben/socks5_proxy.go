package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type GoogleSheet struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

type AssetList struct {
	ID           string `json:"id"`
	Num          string `json:"num"`
	City         string `json:"city"`
	Phone        string `json:"phone"`
	OfflineCount int    `json:"offlineCount"`
}

var dataList []AssetList

type AgentList struct {
	Proxies []Proxies `json:"proxies"`
}

type Proxies struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

//pi是否上线
func main() {

	s1, s3 := check()
	s2 := reloadData(s3)
	ss := SliceCompleStr(s1, s2)

	fmt.Println(len(ss))
	for _, v := range ss {
		fmt.Println(v)
	}

}

//func StringSliceReflectEqual(a, b []string) (st []string) {
//	return reflect.AppendSlice(a,b)
//	reflect.
//}

func SliceCompleStr(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		m[v]++
	}
	nn := make([]string, 0)
	for value, num := range m {
		if num == 1 {
			nn = append(nn, value)
		}
	}
	return nn

}

// 重新加载资产库
func reloadData(str string) (ss []string) {

	api_key := "AIzaSyAAJ2lbLDfUfp0oWDMR29I9Es498Ei1k_Q"
	spreadsheet_id := "1sWAdfigTggbOU62MKfdzRZDOYAPy_pW7TglfM-Lu-QU"
	tab_name := "4g_agent"

	var url = "https://sheets.googleapis.com/v4/spreadsheets/" +
		spreadsheet_id + "/values/" + tab_name +
		"?alt=json&key=" + api_key

	resp, err := http.Get(url)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var info GoogleSheet

	json.Unmarshal(data, &info)

	for _, v := range info.Values {
		if len(v) > 2 {
			if !strings.Contains(str, v[1]) && v[3] == "可用" {
				ss = append(ss, v[1])
			}
		}

	}

	//fmt.Println(info)

	return
}
func check() (ss []string, st string) {
	postReq, err := http.NewRequest(http.MethodGet, "http://159.75.209.173:7749/api/proxy/tcp", nil)
	if err != nil {
		log.Println("请求接口失败", err)
		return
	}

	postReq.Header.Set("Authorization", "Basic ZnJwX3VzZXI6ZnJwX3VzZXJfMkB3cQ==")
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		log.Println("请求接口失败2", err)
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var mas AgentList
	json.Unmarshal(data, &mas)

	for _, v := range mas.Proxies {

		//if v.Name == "pi_10000000f37a6b82.socks5" {
		//	fmt.Println(v)
		//}
		if v.Status == "offline" && strings.Contains(v.Name, "socks5") {
			str := strings.Split(v.Name, ".")
			st = st + str[0]

		}
		if v.Status == "online" && strings.Contains(v.Name, "socks5") {
			str := strings.Split(v.Name, ".")
			ss = append(ss, str[0])
		}
	}

	return
}
