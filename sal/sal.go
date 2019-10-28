// Package sal is a client for the Sal server API.
package sal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/xpreports"
	"github.com/pkg/errors"
)

// Client is what we need to send the POST request.
type Client struct {
	User     string
	Password string

	ServerURL *url.URL
}

// NewClient creates a new Sal API Client using Config.
func NewClient(conf *config.Config) (*Client, error) {
	baseURL, err := url.Parse(conf.URL)
	if err != nil {
		return nil, fmt.Errorf("sal: parsing sal API URL: %s", err)
	}
	client := Client{
		User:      "sal",
		Password:  conf.Key,
		ServerURL: baseURL,
	}

	return &client, nil
}

const checkinPath = "/checkin/"

// Checkin is our POST request
func (c *Client) Checkin(values *Data) error {
	checkinURL := c.ServerURL
	checkinURL.Path = checkinPath

	j, err := json.Marshal(values)
	if err != nil {
		fmt.Println(err)
	}
	// Create a new POST request with the urlencoded checkin values
	req, err := http.NewRequest("POST", checkinURL.String(), bytes.NewBuffer(j))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}

	// The endpoint requires basic authentication, so set the username/password
	req.SetBasicAuth(c.User, c.Password)

	// We're sending URLEncoded data in the body, so tell the server what to expect
	req.Header.Set("Content-Type", "application/json")

	// Configure new http.client with timeouts
	httpclient := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
		},
	}

	// Execute the request
	resp, err := httpclient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to checkin: %s", err)
	}

	defer resp.Body.Close()

	// Copy the body to the console in case there is any response
	io.Copy(os.Stdout, resp.Body)
	return nil
}

// SendCheckin uses Checkin to send our values to Sal
func SendCheckin(conf *config.Config) error {
	client, err := NewClient(conf)
	if err != nil {
		return errors.Wrap(err, "creating gosal client from config")
	}

	// Execute a checkin, providing the data to send to the checkin endpoint
	report, err := xpreports.Build(conf)
	if err != nil {
		return errors.Wrap(err, "build report")
	}

	err = client.Checkin(&Data{
		Machine: report.Machine,
		Sal:     report.Sal,
	})

	return errors.Wrap(err, "checkin")
}

// Data is the collective POST structure
type Data struct {
	Machine *xpreports.Machine
	Sal     *xpreports.Sal
}
