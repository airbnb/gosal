package sal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/airbnb/gosal/config"
	"github.com/airbnb/gosal/xpreports"
)

func TestCheckin(t *testing.T) {
	// test values
	serial := "serial"
	values := Data{
		Machine: &xpreports.Machine{
			Facts: &xpreports.MachineFacts{
				CheckinModuleVersion: "version",
			},
			ExtraData: &xpreports.MachineExtraData{
				SerialNumber: serial,
			},
		},
		Sal: &xpreports.Sal{
			Facts: &xpreports.SalFacts{
				CheckinModuleVersion: "version",
			},
			ExtraData: &xpreports.SalExtraData{
				Key: "key",
			},
		},
	}

	// create a fake API endpoint served by the test server
	checkin := func(w http.ResponseWriter, r *http.Request) {
		if have, want := r.URL.Path, checkinPath; have != want {
			t.Errorf("have %s, want %s url path for checkin", have, want)
		}
		checkAuth(t, r)
		if have, want := values.Machine.ExtraData.SerialNumber, serial; have != want {
			t.Errorf("parsing serial from form: have %s, want %s", have, want)
		}
	}

	client, _, teardown := setupAPI(t, checkin)
	defer teardown()

	if err := client.Checkin(&values); err != nil {
		t.Fatal(err)
	}
}

const testPassword = "test"

// check authentication sent by the sal client.
func checkAuth(t *testing.T, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		t.Errorf("could not parse BasicAuth from request")
	}

	if have, want := username, "sal"; have != want {
		t.Errorf("have %s, want %s", have, want)
	}

	if have, want := password, testPassword; have != want {
		t.Errorf("have %s, want %s", have, want)
	}

}

// setupAPI creates a sal client and a temporary server used for testing.
func setupAPI(t *testing.T, h http.HandlerFunc) (client *Client, server *httptest.Server, cleanup func()) {
	server = httptest.NewServer(h)
	conf := &config.Config{
		Key: testPassword,
		URL: server.URL,
	}

	client, err := NewClient(conf)
	if err != nil {
		t.Fatal(err)
	}

	// add anything here that should be cleaned up when the tests are run.
	cleanup = func() {
		server.Close()
	}
	return client, server, cleanup
}
