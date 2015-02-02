package emlauncher

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Client struct {
	Host   string
	ApiKey string
	client *http.Client
}

type Package struct {
	PackageURL       string   `json:"package_url"`
	ApplicationURL   string   `json:"application_url"`
	Id               string   `json:"id"`
	Platform         string   `json:"platform"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	IOSIdentifier    string   `json:"ios_identifier"`
	OriginalFileName string   `json:"original_file_name"`
	FileSize         string   `json:"file_size"`
	Created          string   `json:"created"`
	Tags             []string `json:"tags"`
	InstallCount     int      `json:"install_count"`
}

func New(host, apiKey string) *Client {
	return &Client{host, apiKey, &http.Client{}}
}

func (c *Client) Upload(filePath, title, description, tags string, notify string) (*Package, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	_ = writer.WriteField("api_key", c.ApiKey)
	_ = writer.WriteField("title", title)
	_ = writer.WriteField("description", description)
	_ = writer.WriteField("tags", tags)
	_ = writer.WriteField("notify", notify)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest("POST", c.Host+"/api/upload", body)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var pack Package
	err = json.Unmarshal([]byte(data), &pack)
	return &pack, err
}

func (c *Client) List() ([]Package, error) {
	values := url.Values{}
	values.Add("api_key", c.ApiKey)
	request, _ := http.NewRequest("GET", c.Host+"/api/package_list?"+values.Encode(), nil)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var packages []Package
	err = json.Unmarshal([]byte(data), &packages)
	return packages, err
}

func (c *Client) Delete(packageId string) (*Package, error) {
	values := url.Values{}
	values.Add("api_key", c.ApiKey)
	values.Add("id", packageId)
	request, _ := http.NewRequest("GET", c.Host+"/api/delete?"+values.Encode(), nil)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var pack Package
	err = json.Unmarshal([]byte(data), &pack)
	return &pack, err
}
