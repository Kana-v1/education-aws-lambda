package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host       string
	basePath   string
	httpClient http.Client
}

func NewClient(host, token string) *Client {
	return &Client{
		host:       host,
		basePath:   "bot" + token,
		httpClient: *http.DefaultClient,
	}
}

func (c *Client) doRequest(query url.Values) []byte {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, "sendMessage"),
	}

	u.RawQuery = query.Encode()

	fmt.Println(u.String())
	
	response, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func (c *Client) SendMessage(chatID int, message string) {
	q := url.Values{}

	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", message)

	c.doRequest(q)
}
