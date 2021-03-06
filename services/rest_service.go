package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bitmovin/bitmovin-go/bitmovin"
)

type RestService struct {
	Bitmovin *bitmovin.Bitmovin
}

func NewRestService(bitmovin *bitmovin.Bitmovin) *RestService {
	return &RestService{
		Bitmovin: bitmovin,
	}
}

func (r *RestService) Create(relativeURL string, input []byte) ([]byte, error) {
	fullURL := *r.Bitmovin.APIBaseURL + relativeURL
	_, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(input))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", *r.Bitmovin.APIKey)
	req.Header.Set("X-Api-Client", ClientName)
	req.Header.Set("X-Api-Client-Version", Version)

	resp, err := r.Bitmovin.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *RestService) Retrieve(relativeURL string) ([]byte, error) {
	fullURL := *r.Bitmovin.APIBaseURL + relativeURL
	_, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", *r.Bitmovin.APIKey)
	req.Header.Set("X-Api-Client", ClientName)
	req.Header.Set("X-Api-Client-Version", Version)

	resp, err := r.Bitmovin.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *RestService) Delete(relativeURL string) ([]byte, error) {
	fullURL := *r.Bitmovin.APIBaseURL + relativeURL
	_, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", fullURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", *r.Bitmovin.APIKey)
	req.Header.Set("X-Api-Client", ClientName)
	req.Header.Set("X-Api-Client-Version", Version)

	resp, err := r.Bitmovin.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//TODO default value version
func (r *RestService) List(relativeURL string, offset int64, limit int64) ([]byte, error) {
	queryParams := fmt.Sprintf("?offset=%v&limit=%v", offset, limit)
	fullURL := *r.Bitmovin.APIBaseURL + relativeURL + queryParams

	req, err := http.NewRequest("GET", fullURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", *r.Bitmovin.APIKey)
	req.Header.Set("X-Api-Client", ClientName)
	req.Header.Set("X-Api-Client-Version", Version)

	resp, err := r.Bitmovin.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *RestService) RetrieveCustomData(relativeURL string) ([]byte, error) {
	fullURL := *r.Bitmovin.APIBaseURL + relativeURL + "/customData"
	_, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", *r.Bitmovin.APIKey)
	req.Header.Set("X-Api-Client", ClientName)
	req.Header.Set("X-Api-Client-Version", Version)

	resp, err := r.Bitmovin.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
