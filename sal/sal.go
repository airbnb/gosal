// Package sal is a client for the Sal server API.
package sal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/salopensource/gosal/reports"
)

// LoadConfig loads Config from a JSON file path.
func LoadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("sal: opening client config file: %s", err)
	}

	var conf Config
	if err = json.Unmarshal(file, &conf); err != nil {
		return nil, fmt.Errorf("sal: unmarshal config json: %s", err)
	}

	return &conf, nil
}

// Config is Sal client config.
type Config struct {
	Key string
	URL string
}

// Client is just an example of what you might collect into a "client" struct.
type Client struct {
	User     string
	Password string

	baseURL *url.URL
}

// NewClient creates a new Sal API Client using Config.
func NewClient(conf *Config) (*Client, error) {
	baseURL, err := url.Parse(conf.URL)
	if err != nil {
		return nil, fmt.Errorf("sal: parsing sal API URL: %s", err)
	}
	client := Client{
		User:     "sal",
		Password: conf.Key,
		baseURL:  baseURL,
	}
	return &client, nil
}

const checkinPath = "/checkin/"

func (c *Client) Checkin(values url.Values) error {
	checkinURL := c.baseURL
	checkinURL.Path = checkinPath
	// Create a new POST request with the urlencoded checkin values
	req, err := http.NewRequest("POST", checkinURL.String(), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}

	// The endpoint requires basic authentication, so set the username/password
	req.SetBasicAuth(c.User, c.Password)

	// We're sending URLEncoded data in the body, so tell the server what to expect
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to checkin: %s", err)
	}
	defer resp.Body.Close()

	// Copy the body to the console in case there is any response
	io.Copy(os.Stdout, resp.Body)
	return nil
}

func SendCheckin() {
	conf, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	// Execute a checkin, providing the data to send to the checkin endpoint
	report := reports.BuildReport(conf.Key)
	fmt.Println(report.Serial)

	client.Checkin(url.Values{
		"serial":      {report.Serial},
		"key":         {report.Key},
		"name":        {report.Name},
		"disk_size":   {report.DiskSize},
		"sal_version": {report.SalVersion},
		"run_uuid":    {report.RunUUID},
	})
}
