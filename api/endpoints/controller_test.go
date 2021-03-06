package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SingleActionTest struct {
	Method         string
	Url            string
	Id             string
	ExpectedStatus int
}

func TestStart(t *testing.T) {
	// Test start single container
	var url = "http://0.0.0.0:8001/api/action/start?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		start(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestStop(t *testing.T) {
	// Test stop single container
	var url = "http://0.0.0.0:8001/api/action/stop?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		stop(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestRestart(t *testing.T) {
	// Test restart single container
	var url = "http://0.0.0.0:8001/api/action/stop?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		restart(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestRemove(t *testing.T) {
	// Test restart single container
	var url = "http://0.0.0.0:8001/api/action/remove?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		restart(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestPause(t *testing.T) {
	// Test pause single container
	var url = "http://0.0.0.0:8001/api/action/pause?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		restart(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestUnpause(t *testing.T) {
	// Test unpause single container
	var url = "http://0.0.0.0:8001/api/action/unpause?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		restart(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestKill(t *testing.T) {
	// Test kill single container
	var url = "http://0.0.0.0:8001/api/action/kill?id="
	stopWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	stopWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := [2]SingleActionTest{stopWithEmptyId, stopWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}
		recorder := httptest.NewRecorder()
		restart(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Expected status %v", test.ExpectedStatus)
		}

		res.Body.Close()
	}
}

func TestGetLogs(t *testing.T) {
	// Test restart single container
	var url = "http://0.0.0.0:8001/api/monitor/logs?id="
	logsWithEmptyId := SingleActionTest{"GET", url, "", http.StatusBadRequest}
	logsWithUnknowId := SingleActionTest{"GET", url, "fakeId", http.StatusNotFound}

	testActions := []SingleActionTest{logsWithEmptyId, logsWithUnknowId}

	for _, test := range testActions {
		req, err := http.NewRequest(test.Method, test.Url+test.Id, nil)

		if err != nil {
			t.Fatalf("could not create request %s", err)
		}

		recorder := httptest.NewRecorder()
		getLogs(recorder, req)

		res := recorder.Result()

		if res.StatusCode != test.ExpectedStatus {
			t.Fatalf("Fail, expected status is %v", test.ExpectedStatus)
		}
	}
}
