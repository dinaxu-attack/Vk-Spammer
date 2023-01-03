package app

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func ReadTokens() []string {
	file, err := os.Open("./assets/tokens.txt")

	if err != nil {
		fmt.Println("Не могу найти файл assets/tokens.txt")
		os.Exit(0)

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text

}

func Request(vurl string) []byte {
	req, err := http.NewRequest("GET", vurl, nil)

	if err != nil {
		fmt.Println("Ошибка при запросе с прокси")
	}

	var proxyURL url.URL
	proxyURL = url.URL{
		Scheme: "http",
		Host:   getProxy()[2] + ":" + getProxy()[3]}

	auth := fmt.Sprintf("%s:%s", getProxy()[0], getProxy()[1])
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Proxy-Authorization", basic)

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "vk.com")
	req.Header.Add("Origin", "https://vk.com")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("TE", "trailers")
	req.Header.Add("User-Agent", user_agent())
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	transport := &http.Transport{
		Proxy:           http.ProxyURL(&proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	transport.ProxyConnectHeader = req.Header

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при запросе с прокси")
	}
	responc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при запросе с прокси")
	}
	return responc

}

func user_agent() string {

	req, _ := http.Get("https://raw.githubusercontent.com/sergeymlnn/Random-User-Agents-Database/main/useragents.txt")
	useragent, _ := ioutil.ReadAll(req.Body)

	strs := strings.Split(string(useragent), "\n")
	cStrs := len(strs)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	d := strs[r1.Intn(cStrs)]

	return string(d)
}

func getProxy() []string {

	file, err := os.Open("./assets/proxies.txt")

	if err != nil {
		fmt.Println("Не могу найти файл assets/proxies.txt")
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var strs []string

	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	file.Close()
	cStrs := len(strs)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	d := strs[r1.Intn(cStrs)]

	proxy := strings.Split(d, ":")

	return proxy

}
