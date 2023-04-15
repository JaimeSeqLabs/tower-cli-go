package testutils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
)

type TestResourceReader func(path string)string

func TestResourceFolder(testDataDir string) TestResourceReader {
	return func(path string) string {
		path = filepath.Join(testDataDir, path)
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		return string(data)
	}
}

func SendJSON(w http.ResponseWriter, json string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write([]byte(json))
}

func SendJSONResFile(resReader TestResourceReader, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendJSON(w, resReader(path))
	}
}

func BodyJSONEq(r *http.Request, json string) bool {
	if r.Body == nil {
		return false
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	
	eq, err := AreEqualJSON(json, string(body))
	if err != nil {
		panic(err)
	}

	return eq
}

type ServerMockMap map[string]map[string]http.HandlerFunc

func MockServer(mockMap ServerMockMap) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		if methodMap, pathFound := mockMap[r.URL.Path]; pathFound && methodMap != nil {
			if handler, methodFound := methodMap[r.Method]; methodFound && handler != nil {
				handler.ServeHTTP(w, r)
			}
		}

	}))

	return server
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
