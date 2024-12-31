package storage

import (
	"io"
	"net/http"
)

type Client struct {
	AccessKey   string
	SecretKey   string
	EndpointURL string
}

// Read downloads the content from the specified path and writes it to the provided io.Writer.
func (sc *Client) Read(w io.Writer, path string) (int64, error) {
	url := sc.EndpointURL + "/" + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.SetBasicAuth(sc.AccessKey, sc.SecretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, &httpError{StatusCode: resp.StatusCode, Status: resp.Status}
	}

	return io.Copy(w, resp.Body)
}

// Write uploads the content from the provided io.Reader to the specified path.
func (sc *Client) Write(path string, r io.Reader) error {
	url := sc.EndpointURL + "/" + path

	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		return err
	}

	req.SetBasicAuth(sc.AccessKey, sc.SecretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return &httpError{StatusCode: resp.StatusCode, Status: resp.Status}
	}

	return nil
}

// httpError is used to handle HTTP response errors.
type httpError struct {
	StatusCode int
	Status     string
}

func (e *httpError) Error() string {
	return "HTTP error: " + e.Status
}
