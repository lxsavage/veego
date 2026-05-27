package veego

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Controller struct {
	active bool
	key    string
	logger *slog.Logger
	http   *http.Client
}

func NewController(apiKey string, logger *slog.Logger) *Controller {
	return &Controller{
		key:    apiKey,
		logger: logger,
	}
}

func (c *Controller) WithClient(h *http.Client) *Controller {
	c.http = h
	return c
}

// Initialize the controller and get it ready to interact with the API
func (c *Controller) Init() (*Controller, error) {
	if c.active {
		return c, errors.New("instance already open")
	}

	if c.http == nil {
		c.http = http.DefaultClient
	}

	c.active = true
	return c, nil
}

// Send a raw request to the Govee API and return the response deserialized as
// T. this should not be used unless the needed functionality is not
// implemented elsewhere
func Request[T any](c Controller, method, endpoint string, content []byte) (*T, error) {
	if !c.active {
		return nil, errors.New("controller not active")
	}

	uri, err := url.JoinPath(apiV1, endpoint)
	if err != nil {
		return nil, err
	}
	c.logger.Debug(fmt.Sprintf("%s %s", method, uri))
	c.logger.Debug(fmt.Sprintf("Content: %v\n", content))

	var body io.Reader = nil
	if content != nil {
		body = bytes.NewBuffer(content)
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Govee-API-Key", c.key)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		c.logger.Error("request failed", "error", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		r, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("%s\n", r)
		return nil, errors.New("non-success status: " + res.Status)
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		c.logger.Warn("unable to read body, treating as if there isn't one", "warning", err)
		return nil, nil
	}

	if len(response) == 0 {
		return nil, nil
	}

	if res.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("unsupported response type " + res.Header.Get("Content-Type"))
	}

	var r T
	if err := json.Unmarshal(response, &r); err != nil {
		return nil, errors.New("failed to unmarshal JSON content: " + err.Error())
	}

	return &r, nil
}
