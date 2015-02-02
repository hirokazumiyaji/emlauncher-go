package emlauncher

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Excepted method %s != POST", r.Method)
		}

		if r.URL.Path != "/api/upload" {
			t.Errorf("Excepted: '%s' != '/api/upload'", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		responseJSON, _ := json.Marshal(Package{
			"http://localhost/emlauncher/package?id=3",
			"http://localhost/emlauncher/app?id=1",
			"3",
			"Android",
			"test upload",
			"upload package via upload api",
			"",
			"emlauncher.apk",
			"5776313",
			"2013-11-29 12:26:19",
			[]string{"test", "upload-api", "android"},
			0})
		io.WriteString(w, string(responseJSON))
	}))

	response, err := New(server.URL, "12345").Upload("./emlauncher.apk", "test upload", "upload package via upload api", "test,upload-api,android", "false")

	if err != nil {
		t.Errorf("Excepted: %s", err)
	}

	if response.ApplicationURL != "http://localhost/emlauncher/app?id=1" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/app?id=1'", response.ApplicationURL)
	}

	if response.PackageURL != "http://localhost/emlauncher/package?id=3" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/package?id=3'", response.PackageURL)
	}

	if response.Id != "3" {
		t.Errorf("Excepted: '%s' != '3'", response.Id)
	}

	if response.Platform != "Android" {
		t.Errorf("Excepted: '%s' != 'Android'", response.Platform)
	}

	if response.Title != "test upload" {
		t.Errorf("Excepted: '%s' != 'test upload'", response.Title)
	}

	if response.Description != "upload package via upload api" {
		t.Errorf("Excepted: '%s' != 'upload package via upload api'", response.Description)
	}

	if response.IOSIdentifier != "" {
		t.Errorf("Excepted: '%s' != ''", response.IOSIdentifier)
	}

	if response.OriginalFileName != "emlauncher.apk" {
		t.Errorf("Excepted: '%s' != 'emlauncher.apk'", response.OriginalFileName)
	}

	if response.FileSize != "5776313" {
		t.Errorf("Excepted: %s != 5776313", response.FileSize)
	}

	if response.Created != "2013-11-29 12:26:19" {
		t.Errorf("Excepted: %s != 2013-11-29 12:26:19", response.Created)
	}

	result := []string{"test", "upload-api", "android"}
	for i := 0; i < len(response.Tags); i++ {
		if response.Tags[i] != result[i] {
			t.Errorf("Excepted: '%s' != '%s'", response.Tags[i], result[i])
		}
	}

	if response.InstallCount != 0 {
		t.Errorf("Excepted: %s != 0", response.InstallCount)
	}
}

func TestList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Excepted method %s != GET", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json; charset=utf-8" {
			t.Errorf("Excepted: %s != application/json; charset=utf-8", r.Header.Get("Content-Type"))
		}

		if r.URL.Path != "/api/package_list" {
			t.Errorf("Excepted: '%s' != '/api/package_list'", r.URL.Path)
		}

		if r.URL.RawQuery != "api_key=12345" {
			t.Errorf("Excepted: '%s' != 'api_key=12345'", r.URL.RawQuery)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		responseJSON, _ := json.Marshal([]Package{
			Package{
				"http://localhost/emlauncher/package?id=3",
				"http://localhost/emlauncher/app?id=1",
				"3",
				"Android",
				"test upload",
				"upload package via upload api",
				"",
				"emlauncher.apk",
				"5776313",
				"2013-11-29 12:26:19",
				[]string{"test", "upload-api", "android"},
				1},
			Package{
				"http://localhost/emlauncher/package?id=1",
				"http://localhost/emlauncher/app?id=1",
				"1",
				"iOS",
				"ipa file test",
				"test package for iPhone",
				"com.klab.playground-sandboxes.test6",
				"emlauncher.ipa",
				"4845763",
				"2013-11-29 09:03:01",
				[]string{"test", "ios"},
				0}})
		io.WriteString(w, string(responseJSON))
	}))

	response, err := New(server.URL, "12345").List()
	if err != nil {
		t.Errorf("Excepted: %s", err)
	}

	if len(response) != 2 {
		t.Errorf("Excepted: %s != 2", len(response))
	}

	p := response[0]
	if p.ApplicationURL != "http://localhost/emlauncher/app?id=1" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/app?id=1'", p.ApplicationURL)
	}

	if p.PackageURL != "http://localhost/emlauncher/package?id=3" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/package?id=3'", p.PackageURL)
	}

	if p.Id != "3" {
		t.Errorf("Excepted: '%s' != '3'", p.Id)
	}

	if p.Platform != "Android" {
		t.Errorf("Excepted: '%s' != 'Android'", p.Platform)
	}

	if p.Title != "test upload" {
		t.Errorf("Excepted: '%s' != 'test upload'", p.Title)
	}

	if p.Description != "upload package via upload api" {
		t.Errorf("Excepted: '%s' != 'upload package via upload api'", p.Description)
	}

	if p.IOSIdentifier != "" {
		t.Errorf("Excepted: '%s' != ''", p.IOSIdentifier)
	}

	if p.OriginalFileName != "emlauncher.apk" {
		t.Errorf("Excepted: '%s' != 'emlauncher.apk'", p.OriginalFileName)
	}

	if p.FileSize != "5776313" {
		t.Errorf("Excepted: %s != 5776313", p.FileSize)
	}

	if p.Created != "2013-11-29 12:26:19" {
		t.Errorf("Excepted: %s != 2013-11-29 12:26:19", p.Created)
	}

	result := []string{"test", "upload-api", "android"}
	for i := 0; i < len(p.Tags); i++ {
		if p.Tags[i] != result[i] {
			t.Errorf("Excepted: '%s' != '%s'", p.Tags[i], result[i])
		}
	}

	if p.InstallCount != 1 {
		t.Errorf("Excepted: %s != 1", p.InstallCount)
	}

	p = response[1]
	if p.ApplicationURL != "http://localhost/emlauncher/app?id=1" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/app?id=1'", p.ApplicationURL)
	}

	if p.PackageURL != "http://localhost/emlauncher/package?id=1" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/package?id=1'", p.PackageURL)
	}

	if p.Id != "1" {
		t.Errorf("Excepted: '%s' != '1'", p.Id)
	}

	if p.Platform != "iOS" {
		t.Errorf("Excepted: '%s' != 'iOS'", p.Platform)
	}

	if p.Title != "ipa file test" {
		t.Errorf("Excepted: '%s' != 'ipa file test'", p.Title)
	}

	if p.Description != "test package for iPhone" {
		t.Errorf("Excepted: '%s' != 'test package for iPhone'", p.Description)
	}

	if p.IOSIdentifier != "com.klab.playground-sandboxes.test6" {
		t.Errorf("Excepted: '%s' != 'com.klab.playground-sandboxes.test6'", p.IOSIdentifier)
	}

	if p.OriginalFileName != "emlauncher.ipa" {
		t.Errorf("Excepted: '%s' != 'emlauncher.ipa'", p.OriginalFileName)
	}

	if p.FileSize != "4845763" {
		t.Errorf("Excepted: %s != 4845763", p.FileSize)
	}

	if p.Created != "2013-11-29 09:03:01" {
		t.Errorf("Excepted: %s != 2013-11-29 09:03:01", p.Created)
	}

	result = []string{"test", "ios"}
	for i := 0; i < len(p.Tags); i++ {
		if p.Tags[i] != result[i] {
			t.Errorf("Excepted: '%s' != '%s'", p.Tags[i], result[i])
		}
	}

	if p.InstallCount != 0 {
		t.Errorf("Excepted: %s != 0", p.InstallCount)
	}

}

func TestDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Excepted method %s != GET", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json; charset=utf-8" {
			t.Errorf("Excepted: %s != application/json; charset=utf-8", r.Header.Get("Content-Type"))
		}

		if r.URL.Path != "/api/delete" {
			t.Errorf("Excepted: '%s' != '/api/delete'", r.URL.Path)
		}

		if r.URL.RawQuery != "api_key=12345&id=3" {
			t.Errorf("Excepted: '%s' != 'api_key=12345&id=3'", r.URL.RawQuery)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		responseJSON, _ := json.Marshal(Package{
			"http://localhost/emlauncher/package?id=3",
			"http://localhost/emlauncher/app?id=1",
			"3",
			"Android",
			"test upload",
			"upload package via upload api",
			"",
			"emlauncher.apk",
			"5776313",
			"2013-11-29 12:26:19",
			[]string{"test", "upload-api", "android"},
			0})
		io.WriteString(w, string(responseJSON))
	}))

	response, err := New(server.URL, "12345").Delete("3")

	if err != nil {
		t.Errorf("Excepted: %s", err)
	}

	if response.ApplicationURL != "http://localhost/emlauncher/app?id=1" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/app?id=1'", response.ApplicationURL)
	}

	if response.PackageURL != "http://localhost/emlauncher/package?id=3" {
		t.Errorf("Excepted: '%s' != 'http://localhost/emlauncher/package?id=3'", response.PackageURL)
	}

	if response.Id != "3" {
		t.Errorf("Excepted: '%s' != '3'", response.Id)
	}

	if response.Platform != "Android" {
		t.Errorf("Excepted: '%s' != 'Android'", response.Platform)
	}

	if response.Title != "test upload" {
		t.Errorf("Excepted: '%s' != 'test upload'", response.Title)
	}

	if response.Description != "upload package via upload api" {
		t.Errorf("Excepted: '%s' != 'upload package via upload api'", response.Description)
	}

	if response.IOSIdentifier != "" {
		t.Errorf("Excepted: '%s' != ''", response.IOSIdentifier)
	}

	if response.OriginalFileName != "emlauncher.apk" {
		t.Errorf("Excepted: '%s' != 'emlauncher.apk'", response.OriginalFileName)
	}

	if response.FileSize != "5776313" {
		t.Errorf("Excepted: %s != 5776313", response.FileSize)
	}

	if response.Created != "2013-11-29 12:26:19" {
		t.Errorf("Excepted: %s != 2013-11-29 12:26:19", response.Created)
	}

	result := []string{"test", "upload-api", "android"}
	for i := 0; i < len(response.Tags); i++ {
		if response.Tags[i] != result[i] {
			t.Errorf("Excepted: '%s' != '%s'", response.Tags[i], result[i])
		}
	}

	if response.InstallCount != 0 {
		t.Errorf("Excepted: %s != 0", response.InstallCount)
	}
}
