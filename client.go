//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"

	"strings"
	"time"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const userAgent = "gofish/1.0"
const applicationJSON = "application/json"

// APIClient represents a connection to a Redfish/Swordfish enabled service
// or device.
type APIClient struct {
	// ctx is the context used in the HTTP requests
	ctx context.Context

	// Endpoint is the URL of the *fish service
	endpoint string

	// HTTPClient is for direct http actions
	HTTPClient *http.Client

	// Service is the ServiceRoot of this Redfish instance
	Service *Service

	// Auth information saved for later to be able to log out
	auth *redfish.AuthToken

	// dumpWriter will receive HTTP dumps if non-nil.
	dumpWriter io.Writer
}

// Session holds the session ID and auth token needed to identify an
// authenticated client
type Session struct {
	ID    string
	Token string
}

// ClientConfig holds the settings for establishing a connection.
type ClientConfig struct {
	// Endpoint is the URL of the redfish service
	Endpoint string

	// Username is the optional user name to authenticate with.
	Username string

	// Password is the password to use for authentication.
	Password string

	// Session is an optional session ID+token obtained from a previous session
	// If this is set, it is preferred over Username and Password
	Session *Session

	// Insecure controls whether to enforce SSL certificate validity.
	Insecure bool

	// Controls TLS handshake timeout
	TLSHandshakeTimeout int

	// HTTPClient is the optional client to connect with.
	HTTPClient *http.Client

	// DumpWriter is an optional io.Writer to receive dumps of HTTP
	// requests and responses.
	DumpWriter io.Writer

	// BasicAuth tells the APIClient if basic auth should be used (true) or token based auth must be used (false)
	BasicAuth bool
}

// setupClientWithConfig setups the client using the client config
func setupClientWithConfig(ctx context.Context, config *ClientConfig) (c *APIClient, err error) {
	if !strings.HasPrefix(config.Endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}

	client := &APIClient{
		endpoint:   config.Endpoint,
		dumpWriter: config.DumpWriter,
		ctx:        ctx,
	}

	if config.TLSHandshakeTimeout == 0 {
		config.TLSHandshakeTimeout = 10
	}

	if config.HTTPClient == nil {
		defaultTransport := http.DefaultTransport.(*http.Transport)
		transport := &http.Transport{
			Proxy:                 defaultTransport.Proxy,
			DialContext:           defaultTransport.DialContext,
			MaxIdleConns:          defaultTransport.MaxIdleConns,
			IdleConnTimeout:       defaultTransport.IdleConnTimeout,
			ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
			TLSHandshakeTimeout:   time.Duration(config.TLSHandshakeTimeout) * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.Insecure,
			},
		}
		client.HTTPClient = &http.Client{Transport: transport}
	} else {
		client.HTTPClient = config.HTTPClient
	}

	// Fetch the service root
	client.Service, err = ServiceRoot(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// setupClientWithEndpoint setups the client using only the endpoint
func setupClientWithEndpoint(ctx context.Context, endpoint string) (c *APIClient, err error) {
	if !strings.HasPrefix(endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}

	client := &APIClient{
		endpoint: endpoint,
		ctx:      ctx,
	}
	client.HTTPClient = &http.Client{}

	// Fetch the service root
	client.Service, err = ServiceRoot(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// setupClientAuth setups the authentication in the client using the client config
func (c *APIClient) setupClientAuth(config *ClientConfig) error {
	if config.Session != nil {
		c.auth = &redfish.AuthToken{
			Session: config.Session.ID,
			Token:   config.Session.Token,
		}
	} else if config.Username != "" {
		var auth *redfish.AuthToken
		if config.BasicAuth {
			auth = &redfish.AuthToken{
				Username:  config.Username,
				Password:  config.Password,
				BasicAuth: true,
			}
		} else {
			var err error
			auth, err = c.Service.CreateSession(config.Username, config.Password)
			if err != nil {
				return err
			}
		}

		c.auth = auth
	}

	return nil
}

// Connect creates a new client connection to a Redfish service.
func Connect(config ClientConfig) (c *APIClient, err error) { //nolint:gocritic
	return ConnectContext(context.Background(), config)
}

// ConnectContext is the same as Connect, but sets the ctx.
func ConnectContext(ctx context.Context, config ClientConfig) (c *APIClient, err error) { //nolint:gocritic
	client, err := setupClientWithConfig(ctx, &config)
	if err != nil {
		return c, err
	}

	// Authenticate with the service
	err = client.setupClientAuth(&config)
	if err != nil {
		return c, err
	}

	return client, err
}

// ConnectDefault creates an unauthenticated connection to a Redfish service.
func ConnectDefault(endpoint string) (c *APIClient, err error) {
	return ConnectDefaultContext(context.Background(), endpoint)
}

// ConnectDefaultContext is the same as ConnectDefault, but sets the ctx.
func ConnectDefaultContext(ctx context.Context, endpoint string) (c *APIClient, err error) {
	client, err := setupClientWithEndpoint(ctx, endpoint)
	if err != nil {
		return c, err
	}

	return client, err
}

// GetService returns the APIClient's service.
func (c *APIClient) GetService() *Service {
	return c.Service
}

// CloneWithSession will create a new Client with a session instead of basic auth.
func (c *APIClient) CloneWithSession() (*APIClient, error) {
	if c.auth.Session != "" {
		return nil, fmt.Errorf("client already has a session")
	}

	newClient := *c
	newClient.HTTPClient = c.HTTPClient
	service, err := ServiceRoot(&newClient)
	if err != nil {
		return nil, err
	}
	newClient.Service = service

	auth, err := newClient.Service.CreateSession(
		newClient.auth.Username,
		newClient.auth.Password)
	if err != nil {
		return nil, err
	}
	newClient.auth = auth

	return &newClient, err
}

// GetSession retrieves the session data from an initialized APIClient. An error
// is returned if the client is not authenticated.
func (c *APIClient) GetSession() (*Session, error) {
	if c.auth == nil || c.auth.Session == "" {
		return nil, fmt.Errorf("client not authenticated")
	}
	return &Session{
		ID:    c.auth.Session,
		Token: c.auth.Token,
	}, nil
}

// Get performs a GET request against the Redfish service.
func (c *APIClient) Get(url string) (*http.Response, error) {
	return c.GetWithHeaders(url, nil)
}

// GetWithHeaders performs a GET request against the Redfish service but allowing custom headers
func (c *APIClient) GetWithHeaders(url string, customHeaders map[string]string) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	return c.runRequestWithHeaders(http.MethodGet, relativePath, nil, customHeaders)
}

// Post performs a Post request against the Redfish service.
func (c *APIClient) Post(url string, payload interface{}) (*http.Response, error) {
	return c.PostWithHeaders(url, payload, nil)
}

// PostWithHeaders performs a Post request against the Redfish service but allowing custom headers
func (c *APIClient) PostWithHeaders(url string, payload interface{}, customHeaders map[string]string) (*http.Response, error) {
	return c.runRequestWithHeaders(http.MethodPost, url, payload, customHeaders)
}

// PostMultipart performs a Post request against the Redfish service with multipart payload.
func (c *APIClient) PostMultipart(url string, payload map[string]io.Reader) (*http.Response, error) {
	return c.PostMultipartWithHeaders(url, payload, nil)
}

// PostMultipartWithHeadersperforms a Post request against the Redfish service with multipart payload but allowing custom headers
func (c *APIClient) PostMultipartWithHeaders(url string, payload map[string]io.Reader, customHeaders map[string]string) (*http.Response, error) {
	return c.runRequestWithMultipartPayloadWithHeaders(http.MethodPost, url, payload, customHeaders)
}

// Put performs a Put request against the Redfish service.
func (c *APIClient) Put(url string, payload interface{}) (*http.Response, error) {
	return c.PutWithHeaders(url, payload, nil)
}

// PutWithHeaders performs a Put request against the Redfish service but allowing custom headers
func (c *APIClient) PutWithHeaders(url string, payload interface{}, customHeaders map[string]string) (*http.Response, error) {
	return c.runRequestWithHeaders(http.MethodPut, url, payload, customHeaders)
}

// Patch performs a Patch request against the Redfish service.
func (c *APIClient) Patch(url string, payload interface{}) (*http.Response, error) {
	return c.PatchWithHeaders(url, payload, nil)
}

// PatchWithHeaders performs a Patch request against the Redfish service but allowing custom headers
func (c *APIClient) PatchWithHeaders(url string, payload interface{}, customHeaders map[string]string) (*http.Response, error) {
	return c.runRequestWithHeaders(http.MethodPatch, url, payload, customHeaders)
}

// Delete performs a Delete request against the Redfish service
func (c *APIClient) Delete(url string) (*http.Response, error) {
	return c.DeleteWithHeaders(url, nil)
}

// DeleteWithHeaders performs a Delete request against the Redfish service but allowing custom headers
func (c *APIClient) DeleteWithHeaders(url string, customHeaders map[string]string) (*http.Response, error) {
	resp, err := c.runRequestWithHeaders(http.MethodDelete, url, nil, customHeaders)
	if err != nil {
		return nil, err
	}
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return resp, nil
}

// runRequestWithHeaders performs JSON REST calls but allowing custom headers
func (c *APIClient) runRequestWithHeaders(method, url string, payload interface{}, customHeaders map[string]string) (*http.Response, error) {
	if url == "" {
		return nil, fmt.Errorf("unable to execute request, no target provided")
	}

	var payloadBuffer io.ReadSeeker
	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		payloadBuffer = bytes.NewReader(body)
	}

	return c.runRawRequestWithHeaders(method, url, payloadBuffer, applicationJSON, customHeaders)
}

// runRequestWithMultipartPayloadWithHeaders performs REST calls with a multipart payload but allowing custom headers
func (c *APIClient) runRequestWithMultipartPayloadWithHeaders(method, url string, payload map[string]io.Reader, customHeaders map[string]string) (*http.Response, error) {
	if url == "" {
		return nil, fmt.Errorf("unable to execute request, no target provided")
	}

	var payloadBuffer bytes.Buffer
	var err error
	payloadWriter := multipart.NewWriter(&payloadBuffer)
	for key, reader := range payload {
		var partWriter io.Writer
		if file, ok := reader.(*os.File); ok {
			// Add a file stream
			if partWriter, err = payloadWriter.CreateFormFile(key, filepath.Base(file.Name())); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if partWriter, err = createFormField(key, payloadWriter); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(partWriter, reader); err != nil {
			return nil, err
		}
	}
	payloadWriter.Close()

	return c.runRawRequestWithHeaders(method, url, bytes.NewReader(payloadBuffer.Bytes()), payloadWriter.FormDataContentType(), customHeaders)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// createFormField create form field with Content-Type
func createFormField(fieldname string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name=%q`, escapeQuotes(fieldname)))
	h.Set("Content-Type", "application/json")
	return w.CreatePart(h)
}

// runRawRequest actually performs the REST calls
func (c *APIClient) runRawRequest(method, url string, payloadBuffer io.ReadSeeker, contentType string) (*http.Response, error) {
	return c.runRawRequestWithHeaders(method, url, payloadBuffer, contentType, nil)
}

// RunRawRequestWithHeaders actually performs the REST calls but allowing custom headers
func (c *APIClient) RunRawRequestWithHeaders(method, url string, payloadBuffer io.ReadSeeker, contentType string, customHeaders map[string]string) (*http.Response, error) {
	return c.runRawRequestWithHeaders(method, url, payloadBuffer, contentType, customHeaders)
}

// runRawRequestWithHeaders actually performs the REST calls but allowing custom headers
func (c *APIClient) runRawRequestWithHeaders(method, url string, payloadBuffer io.ReadSeeker, contentType string, customHeaders map[string]string) (*http.Response, error) {
	if url == "" {
		return nil, common.ConstructError(0, []byte("unable to execute request, no target provided"))
	}

	endpoint := fmt.Sprintf("%s%s", c.endpoint, url)
	req, err := http.NewRequestWithContext(c.ctx, method, endpoint, payloadBuffer)
	if err != nil {
		return nil, err
	}

	// Add common headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", applicationJSON)

	// Add custom headers
	for k, v := range customHeaders {
		if k == "" && v == "" { // Quick check to avoid empty headers
			continue
		}

		// Set Content-Length custom headers on the request
		// since its ignored when set using Header.Set()
		if strings.EqualFold("Content-Length", k) {
			req.ContentLength, err = strconv.ParseInt(v, 10, 64) //nolint:gomnd // base 10, 64 bit
			if err != nil {
				return nil, common.ConstructError(0, []byte("error parsing custom Content-Length header"))
			}

			continue
		}

		req.Header.Set(k, v)
	}

	// Add content info if present
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// Add auth info if authenticated
	if c.auth != nil {
		if c.auth.Token != "" {
			req.Header.Set("X-Auth-Token", c.auth.Token)
		} else if c.auth.BasicAuth && c.auth.Username != "" && c.auth.Password != "" {
			encodedAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", c.auth.Username, c.auth.Password)))
			req.Header.Set("Authorization", fmt.Sprintf("Basic %v", encodedAuth))
		}
	}
	req.Close = true

	// Dump request if needed.
	if c.dumpWriter != nil {
		if err := c.dumpRequest(req); err != nil {
			return nil, err
		}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Dump response if needed.
	if c.dumpWriter != nil {
		if err := c.dumpResponse(resp); err != nil {
			defer resp.Body.Close()
			return nil, err
		}
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		payload, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, common.ConstructError(0, []byte(err.Error()))
		}
		defer resp.Body.Close()
		return nil, common.ConstructError(resp.StatusCode, payload)
	}

	return resp, err
}

// dumpRequest writes outgoing client requests to dumpWriter
func (c *APIClient) dumpRequest(req *http.Request) error {
	d, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return common.ConstructError(0, []byte(err.Error()))
	}

	d = append(d, '\n')
	_, err = c.dumpWriter.Write(d)
	if err != nil {
		panic(err)
	}

	return nil
}

// dumpRequest writes incoming responses to dumpWriter
func (c *APIClient) dumpResponse(resp *http.Response) error {
	d, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return common.ConstructError(0, []byte(err.Error()))
	}

	d = append(d, '\n')
	_, err = c.dumpWriter.Write(d)
	if err != nil {
		panic(err)
	}

	return nil
}

// Logout will delete any active session. Useful to defer logout when creating
// a new connection.
func (c *APIClient) Logout() {
	if c.Service != nil && c.auth != nil {
		_ = c.Service.DeleteSession(c.auth.Session)
	}
}

// SetDumpWriter sets the client the DumpWriter dynamically
func (c *APIClient) SetDumpWriter(writer io.Writer) {
	c.dumpWriter = writer
}
