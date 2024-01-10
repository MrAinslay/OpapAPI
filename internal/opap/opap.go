package opap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.opap.gr/draws/v3.0/"
)

// Client manages communication with the OPAP API.
type Client struct {
	client *http.Client

	BaseURL *url.URL

	Draws *drawsService
}

// NewClient returns a new OPAP API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	c.Draws = &drawsService{
		client: c,
	}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
//
// OPAP REST Services: https://www.opap.gr/en/web-services
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		b := resp.Body
		var buf bytes.Buffer
		r := io.TeeReader(b, &buf)
		if err := json.NewDecoder(r).Decode(v); err != nil {
			return resp, fmt.Errorf("JSON decoding: %v (%s)", err, buf.String())
		}
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	return fmt.Errorf("%v %v: %d %s", r.Request.Method, r.Request.URL, r.StatusCode, string(data))
}

func (c *Client) get(url string, result interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, result)
}

// Game is used to specify which OPAP game to bring results for.
type Game int

// Constants of all the game types the OPAP REST service supports except Propo.
const (
	Pοwerspin = 1110
)

// drawsService handles communication with the DrawsRestServices endpoint.
//
// OPAP REST Services: https://www.opap.gr/en/web-services
type drawsService struct {
	client   *Client
	Endpoint string
}

type draws struct {
	Draw Draw `json:"draw"`
}

type drawsByDate struct {
	Draws struct {
		Draw []Draw `json:"draw"`
	} `json:"draws"`
}

func (s *drawsService) Latest(g Game) (*PasedDraws, *http.Response, error) {
	d := new(PasedDraws)
	u := fmt.Sprint(defaultBaseURL, g, "/draw-id/1/10")
	resp, err := s.client.get(u, d)
	if err != nil {
		return nil, resp, err
	}
	return d, resp, nil
}

func (s *drawsService) ByNumber(g Game, number int) (*Draw, *http.Response, error) {
	d := new(draws)
	u := fmt.Sprint(defaultBaseURL, g, number)
	resp, err := s.client.get(u, d)
	if err != nil {
		return nil, resp, err
	}
	return &d.Draw, resp, nil
}

func (s *drawsService) ByDate(g Game, day, month, year int) ([]Draw, *http.Response, error) {
	d := new(drawsByDate)
	date := fmt.Sprintf("%d-%d-%d", day, month, year)
	u := fmt.Sprint(defaultBaseURL, g, date)
	resp, err := s.client.get(u, d)
	if err != nil {
		return nil, resp, err
	}
	return d.Draws.Draw, resp, nil
}
