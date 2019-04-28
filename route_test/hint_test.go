package route_test

import (
	"../route"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

/*

1,

*/

func TestHello (t *testing.T) {
	var router route.Router

	data := url.Values{}
	data.Set("cards", `{"coords":[[[1,1,1],[2,2,2],[3,3,3]], [[1,1,1],[2,2,2],[3,3,3]], [[1,1,1],[2,2,2],[3,3,3]] ]}`)


	req, err := http.NewRequest("POST", "/hint", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.HintHandler)


	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

}