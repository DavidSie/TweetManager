package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DavidSie/TweetManager/internal/repository/dbrepo"
	"github.com/DavidSie/TweetManager/pkg/model"
)

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"tweets with emotions", "/tweets-with-emotions", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for  %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestRepository_TweetsJSON(t *testing.T) {

	testCases := []struct {
		name                         string
		url                          string
		expectedCode                 uint
		isResponseBodyATweetResponse bool
		expectedResponseOK           bool
	}{
		{
			name:                         "Succesful path",
			url:                          "/tweets?symbol=my_symbol&start_date=2024-06-24&end_date=2024-07-09",
			expectedCode:                 http.StatusOK,
			isResponseBodyATweetResponse: true,
			expectedResponseOK:           true,
		},
		{
			name:                         "Missing Symbol",
			url:                          "/tweets",
			expectedCode:                 http.StatusBadRequest,
			isResponseBodyATweetResponse: false,
		},
		{
			name:                         "Missing Start Date",
			url:                          "/tweets?symbol=my_symbol",
			expectedCode:                 http.StatusBadRequest,
			isResponseBodyATweetResponse: false,
		},
		{
			name:                         "Missing End Date",
			url:                          "/tweets?symbol=my_symbol&start_date=2024-06-24&",
			expectedCode:                 http.StatusBadRequest,
			isResponseBodyATweetResponse: false,
		},
		{
			name:                         "Start date misconfigured path",
			url:                          "/tweets?symbol=my_symbol&start_date=SOMETHING&end_date=2024-07-09",
			expectedCode:                 http.StatusInternalServerError,
			isResponseBodyATweetResponse: false,
		},
		{
			name:                         "End date misconfigured path",
			url:                          "/tweets?symbol=my_symbol&start_date=2024-07-09&end_date=SOMETHING",
			expectedCode:                 http.StatusInternalServerError,
			isResponseBodyATweetResponse: false,
		},
		{
			name:                         "Database Error",
			url:                          fmt.Sprintf("/tweets?symbol=%s&start_date=2024-06-24&end_date=2024-07-09", dbrepo.TriggerDBErrorSymbolOnTest),
			expectedCode:                 http.StatusInternalServerError,
			isResponseBodyATweetResponse: false,
		},
	}
	for _, tc := range testCases {
		req, err := http.NewRequest("GET", tc.url, nil)
		if err != nil {
			t.Errorf("Test case %s:\n did not expect error but got one: %s", tc.name, err.Error())
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.TweetsJSON)
		handler.ServeHTTP(rr, req)

		if rr.Code != int(tc.expectedCode) {
			t.Errorf("Test case %s:\n Tweet JSON handler returned wrong response code: got  %d, wanted %d", tc.name, rr.Code, http.StatusOK)
		}
		if tc.isResponseBodyATweetResponse {
			// I'll check only if the data is parsable to TweetResponse Tweets as the result depends on test-repo, which is for testing purposes only
			expectedResponse := model.TweetResponse{}
			err = json.Unmarshal(rr.Body.Bytes(), &expectedResponse)
			if err != nil {
				t.Error(err.Error())
			}
			if expectedResponse.OK != tc.expectedResponseOK {
				t.Errorf("Test case %s:\n expected response.OK to be %v, got %v", tc.name, tc.expectedResponseOK, expectedResponse.OK)
			}
		}
	}

}
