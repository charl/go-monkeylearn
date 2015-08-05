package gomonkeylearn

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const (
	apiToken = "THISISATESYTAPITOKEN"
)

func TestNewClient(t *testing.T) {
	client := NewClient(apiToken)
	got := reflect.TypeOf(client).Elem().Name()
	want := "Client"
	if got != want {
		t.Errorf("NewClient(%q) == %q, want %q", apiToken, got, want)
	}
}

func TestClientClassifyUnknownClassifierCategory(t *testing.T) {
	client := NewClient(apiToken)
	category := "UNKNOWN CATEGORY"
	got, err := client.Classify(category, []string{"foo bar"})
	if got != nil && err.Error() != fmt.Sprintf("unknown classifier category %s", category) {
		t.Errorf("NewClient(%q) == %q, want nil", apiToken, got)
	}
}

func TestClientClassify(t *testing.T) {
	setup(apiToken)
	defer teardown()

	// Test all classifier categories.
	cases := []struct {
		category string
		in       []string
		want     string
	}{
		{
			"News Categorizer",
			[]string{"First text to classify", "Second text to classify"},
			`{"result": [[{"probability": 0.65, "label": "Arts & Culture"}, {"probability": 0.72, "label": "Books"}]]}`,
		},
	}
	for _, c := range cases {
		mux.HandleFunc(fmt.Sprintf("/classifiers/%s/classify/", client.Categorizers[c.category]), func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, c.want)
		})

		got, err := client.Classify(c.category, c.in)
		if err != nil {
			t.Errorf("Classification category %s returned: %s", c.category, err)
		}
		if string(got) != c.want {
			t.Errorf("Classification category %s: got %s, want %s", c.category, got, c.want)
		}
	}
}
