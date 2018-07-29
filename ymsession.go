package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type YMSession struct {
	client    *http.Client
	userAgent string
	cookies   map[string]*http.Cookie
	sk        string
}

func NewYMSession(client *http.Client, userAgent string) *YMSession {
	return &YMSession{
		client:    client,
		userAgent: userAgent,
		cookies:   make(map[string]*http.Cookie, 10),
	}
}

func (sess *YMSession) prepareRequest(url string) (req *http.Request, err error) {

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}
	req.Header.Add("user-agent", sess.userAgent)
	req.Header.Add("Accept-Encoding", "gzip")

	for _, cookie := range sess.cookies {
		req.AddCookie(cookie)
	}

	return
}

func (sess *YMSession) Start() error {

	body, err := sess.Request("https://app.market.yandex.ru/")
	if err != nil {
		return err
	}

	sbody := string(body)
	pos := strings.Index(sbody, `"sk":"`)
	if pos < 0 {
		return fmt.Errorf("Expected `sk` in body, but it was not found")
	}
	sk := sbody[pos+6:]
	pos = strings.Index(sk, `"`)
	sess.sk = fmt.Sprintf("%s", sk[:pos])
	sess.cookies["currentRegionId"] = &http.Cookie{Name: "currentRegionId", Value: "213", Path: "/", Domain: "yandex.ru"}

	return nil
}

func (sess *YMSession) Request(url string) (body []byte, err error) {
	if len(sess.sk) != 0 {
		url = url + "&sk=" + sess.sk
	}
	req, _ := sess.prepareRequest(url)

	resp, err := sess.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected 200, but got %d", resp.StatusCode)
	}
	for _, cookie := range resp.Cookies() {
		sess.cookies[cookie.Name] = cookie
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}
	body, err = ioutil.ReadAll(reader)

	return body, err
}

func (sess *YMSession) RequestJSON(url string, dest interface{}) error {
	if len(sess.sk) != 0 {
		url = url + "&sk=" + sess.sk
	}
	req, _ := sess.prepareRequest(url)

	resp, err := sess.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected 200, but got %d", resp.StatusCode)
	}
	for _, cookie := range resp.Cookies() {
		sess.cookies[cookie.Name] = cookie
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	dec := json.NewDecoder(reader)

	return dec.Decode(dest)
}

func (sess *YMSession) GetSearchResults(Nid, CatID, Hid []string, Page int) (resp *YMResultsResp, err error) {

	url := fmt.Sprintf("https://app.market.yandex.ru/api/search/result?local-offers-first=0&onstock=0&numdoc=16&isMultisearch=false&omit-render=1&page=%d", Page)

	if len(Nid) > 0 {
		url = url + "&nid=" + Nid[0]
	}

	if len(CatID) > 0 {
		url = url + "&catID=" + CatID[0]
	}
	if len(Hid) > 0 {
		url = url + "&hid=" + Hid[0]
	}

	resp = &YMResultsResp{}
	if err := sess.RequestJSON(url, &resp); err != nil {
		// log.Println(err)
		return nil, err
	}
	return resp, nil
}

func (sess *YMSession) GetNavigationTree(category int) (resp *YMNavigationResp, err error) {

	resp = &YMNavigationResp{}
	if err := sess.RequestJSON(fmt.Sprintf("https://app.market.yandex.ru/api/cataloger/getNavigationTree?nid=%d", category), resp); err != nil {
		return nil, err
	}

	return resp, nil
}
