package send

import (
	"fmt"
	"github.com/salopensource/gosal/reports"
	"github.com/salopensource/gosal/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// CheckinClient is just an example of what you might collect into a "client" struct.
type CheckinClient struct {
	URL      string
	User     string
	Password string
}

func (c *CheckinClient) Checkin(values url.Values) error {
	// Create a new POST request with the urlencoded checkin values
	req, err := http.NewRequest("POST", c.URL, strings.NewReader(values.Encode()))
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
	// Make a fake checkin server that we can pretend to connect to
	// fakeCheckinServer := httptest.NewServer(http.HandlerFunc(logRequest))

	// Create a "client" with the stuff that's the same for every request
	config := &CheckinClient{
		URL:      utils.LoadConfig("./config.json").URL + "/checkin/",
		User:     "sal",
		Password: utils.LoadConfig("./config.json").Key,
	}

	// Execute a checkin, providing the data to send to the checkin endpoint
	report := reports.BuildReport()
	fmt.Println(report.Serial)

	config.Checkin(url.Values{
		"serial":      {report.Serial},
		"key":         {report.Key},
		"name":        {report.Name},
		"disk_size":   {report.DiskSize},
		"sal_version": {report.SalVersion},
		"run_uuid":    {report.RunUUID},
	})
}

// logRequest is just a dummy "in process" thing that pretends to be the remote server and prints out what it gets
// func logRequest(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Path: %q", r.URL.Path)
// 	if err := r.ParseForm(); err != nil {
// 		log.Fatalf("failed to parse form: %s", err)
// 	}
//
// 	username, password, ok := r.BasicAuth()
// 	log.Printf("User auth: %s %q (%v)", username, password, ok)
// 	log.Printf("Server records:")
// 	for k, v := range r.Form {
// 		log.Printf(" - %q = %q", k, v)
// 	}
//
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "OK")
// }
