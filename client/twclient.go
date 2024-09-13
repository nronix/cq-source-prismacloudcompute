package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type APIClientConfig struct {
	ConsoleURL           string `json:"console_url"`
	Project              string `json:"project"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	SkipCertVerification bool   `json:"skip_cert_verification"`
}

// TwClient A connection to Prisma Cloud Compute.
type TwClient struct {
	Config     APIClientConfig
	HTTPClient *http.Client
	JWT        string
}

type ErrResponse struct {
	Err string
}

func (c *TwClient) Initialize(filename string) error {
	c2 := TwClient{}

	if filename != "" {
		var (
			b   []byte
			err error
		)

		b, err = ioutil.ReadFile(filename)

		if err != nil {
			return err
		}

		if err = json.Unmarshal(b, &c2); err != nil {
			return err
		}
	}

	if c.Config.ConsoleURL == "" && c2.Config.ConsoleURL != "" {
		c.Config.ConsoleURL = c2.Config.ConsoleURL
	}

	if c.Config.Project == "" && c2.Config.Project != "" {
		c.Config.Project = c2.Config.Project
	}

	if c.Config.Username == "" && c2.Config.Username != "" {
		c.Config.Username = c2.Config.Username
	}

	if c.Config.Password == "" && c2.Config.Password != "" {
		c.Config.Password = c2.Config.Password
	}

	c.HTTPClient = &http.Client{}

	return c.Authenticate()
}

// Communicate with the Prisma Cloud Compute API.
func (c *TwClient) Request(method, endpoint string, query, data, response interface{}, logger *zerolog.Logger) (err error) {
	parsedURL, err := url.Parse(c.Config.ConsoleURL)
	if err != nil {

		return err
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}
	parsedURL.Path = path.Join(parsedURL.Path, endpoint)

	var buf bytes.Buffer

	if data != nil {
		data_json, err := json.Marshal(data)
		if err != nil {
			return err
		}
		buf = *bytes.NewBuffer(data_json)
	}

	req, err := http.NewRequest(method, parsedURL.String(), &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.JWT)
	req.Header.Set("Content-Type", "application/json")

	// TODO: simplify logic
	if c.Config.Project != "" {
		queryParams := req.URL.Query()
		queryParams.Set("project", c.Config.Project)
		if query != nil {
			if queryMap, ok := query.(map[string]string); ok {
				for key, val := range queryMap {
					queryParams.Add(key, val)
				}
			}
		}
		req.URL.RawQuery = queryParams.Encode()

	} else if query != nil {
		queryParams := req.URL.Query()
		if queryMap, ok := query.(map[string]string); ok {
			for key, val := range queryMap {
				queryParams.Add(key, val)
			}
		}
		req.URL.RawQuery = queryParams.Encode()

	}
	if logger != nil {
		logger.Printf("Query logging", req.URL.RawQuery)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Retry in case backend responds with HTTP 429
	// sleep for 3 seconds before retry
	if res.StatusCode == 429 {
		time.Sleep(3 * time.Second)
		return c.Request(method, endpoint, query, data, &response, logger)
	}

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error reading response body from non-OK response: %s", err)
		}

		var response ErrResponse
		if err = json.Unmarshal(body, &response); err != nil {
			//fmt.Printf(string(body))
			return err
		}

		return fmt.Errorf("Non-OK status: %d (%s)", res.StatusCode, response.Err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(body) > 0 {
		if err = json.Unmarshal(body, response); err != nil {

			return err
		}
	}
	return nil
}

// Authenticate with the Prisma Cloud Compute Console.
func (c *TwClient) Authenticate() (err error) {

	type AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type AuthResponse struct {
		Token string `json:"token"`
	}

	res := AuthResponse{}
	if err := c.Request(http.MethodPost, "/api/v1/authenticate", nil, AuthRequest{c.Config.Username, c.Config.Password}, &res, nil); err != nil {
		return fmt.Errorf("error POSTing to authenticate endpoint: %v", err)
	}
	c.JWT = res.Token
	return nil
}

// CreateTwClient and authenticate.
func APIClient(config APIClientConfig) (*TwClient, error) {
	apiClient := &TwClient{
		Config: config,
	}

	if config.SkipCertVerification {
		apiClient.HTTPClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	} else {
		apiClient.HTTPClient = &http.Client{}
	}

	if err := apiClient.Authenticate(); err != nil {
		return nil, err
	}

	return apiClient, nil
}
