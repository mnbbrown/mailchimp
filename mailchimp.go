package mailchimp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// A Client manages communication with the Mailchimp API.
type Client struct {
	client  *http.Client
	BaseURL *url.URL
	DC      string
	APIKey  string
}

// NewClient returns a new Mailchimp API client.  If a nil httpClient is
// provided, http.DefaultClient will be used. The apiKey must be in the format xyz-us11.
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	if len(strings.Split(apiKey, "-")) != 2 {
		return nil, errors.New("Mailchimp API Key must be formatted like: xyz-zys")
	}
	dc := strings.Split(apiKey, "-")[1]
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseUrl, _ := url.Parse(fmt.Sprintf("https://%s.api.mailchimp.com/3.0", dc))
	return &Client{APIKey: apiKey, client: httpClient, DC: dc, BaseURL: baseUrl}, nil
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

	apiURL := fmt.Sprintf("%s%s", c.BaseURL.String(), path)

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

	var v interface{}
	err = json.NewDecoder(resp.Body).Decode(&v)
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
