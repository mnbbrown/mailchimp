package mailchimp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	APIKey  string
}

func NewClient(apiKey string, client *http.Client) *Client {
	if len(strings.Split(apiKey, "-")) != 2 {
		panic("Mailchimp API Key must be formatted like: xyz-zys")
	}
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{APIKey: apiKey, client: client}
}

type ErrorResponse struct {
	Type   string `json:type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d %s (%s)", e.Status, e.Title, e.Detail)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

func (c *Client) Do(method string, path string, body interface{}) (interface{}, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	dc := strings.Split(c.APIKey, "-")[1]
	apiURL := fmt.Sprintf("https://%s.api.mailchimp.com/3.0%s", dc, path)

	req, err := http.NewRequest(method, apiURL, buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("", c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}
	for h, v := range resp.Header {
		log.Printf("%s: %s", h, v)
	}

	var v interface{}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (c *Client) Subscribe(email string, listId string) (interface{}, error) {
	v, err := c.Do("POST", fmt.Sprintf("/lists/%s/members/", listId), &map[string]string{"email_address": email, "status": "subscribed"})
	if err != nil {
		return v, err
	}
	return v, nil
}
