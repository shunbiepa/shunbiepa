package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var countryCapitalMap map[string]int /*创建集合 */
func main() {

	countryCapitalMap = make(map[string]int)

	countryCapitalMap["192.168.1.38"] = 0
	countryCapitalMap["192.168.1.39"] = 0
	countryCapitalMap["192.168.1.41"] = 0
	countryCapitalMap["192.168.1.40"] = 0
	countryCapitalMap["192.168.1.42"] = 0
	countryCapitalMap["192.168.1.43"] = 0
	countryCapitalMap["192.168.1.44"] = 0
	countryCapitalMap["192.168.1.47"] = 0
	countryCapitalMap["192.168.1.49"] = 0
	countryCapitalMap["192.168.1.50"] = 0
	countryCapitalMap["192.168.1.51"] = 0
	countryCapitalMap["192.168.1.52"] = 0
	countryCapitalMap["192.168.1.53"] = 0
	countryCapitalMap["192.168.1.56"] = 0
	countryCapitalMap["192.168.1.57"] = 0
	countryCapitalMap["192.168.1.18"] = 0
	countryCapitalMap["192.168.1.21"] = 0

	go kill()

	for {

		if time.Now().Hour() >= 23 || time.Now().Hour() <= 9 {
			time.Sleep(time.Minute * 30)
			continue
		}
		Tell()
		time.Sleep(time.Minute * 5)
	}

	select {}

	//fmt.Println("hell-wood")
	//BaiduHotSearch()
	//TgBotSendMsg("123")
}

func TgBotSendMsg(msg string) {
	bot, err := tgbotapi.NewBotAPI("5155017882:AAHCLA3VLW6fwmTHIKZ1y_Wj0jj3ZLbRFB0")
	if err != nil {
		log.Panic(err)
	}
	bot.Send(tgbotapi.NewMessage(-529991306, msg))
}

func Tell() {
	for k, i := range countryCapitalMap {
		st := fmt.Sprintf("socks5://%s:8000", k)
		hts := HttpGet(st)
		if strings.Contains(hts, "长沙") {
			countryCapitalMap[k] = i + 1
			return
		}
		if strings.Contains(hts, "失败") {
			countryCapitalMap[k] = i + 1
			return
		}
		countryCapitalMap[k] = 0
	}
	return
}

func kill() {
	for {
		for k, v := range countryCapitalMap {
			if v >= 2 {
				bots := fmt.Sprintf("\"[报警小助手] \n ip端口:%s　出错了\"", k)
				TgBotSendMsg(bots)
			}
		}
		time.Sleep(time.Minute * 5)
	}

}

func HttpGet(socksProxy string) (ss string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(socksProxy)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	// 设置使用代理,以及超时时间 5s
	httpClient := &http.Client{
		Timeout:   5 * time.Second,
		Transport: httpTransport,
	}

	req, err := http.NewRequest("GET", "https://cip.cc", nil)
	if err != nil {
		log.Printf("请求失败,超时了--1--%s", socksProxy)
		ss = fmt.Sprintf("请求失败,超时了--1--%s", socksProxy)
		return
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("请求失败,超时了--2--%s", socksProxy)
		ss = fmt.Sprintf("请求失败,超时了--2--%s", socksProxy)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("document 解析失败了---%s", socksProxy)
		ss = fmt.Sprintf("document 解析失败了---%s", socksProxy)
		return
	}
	//data kq-well
	doc.Find(".data").Each(func(i int, s *goquery.Selection) {
		//  返回结果是
		ss = s.Text()
	})
	return ss
}
