package HotRemote

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func createWSDialer() (*websocket.Dialer, error) {
	cert, err := LoadX509KeyPair("certificates/cert.pem", "certificates/key.pem")
	if err != nil {
		return nil, err
	}

	// Create headers CA certificate pool and add cert.pem to it
	caCert, err := certificatesDir.ReadFile("certificates/cert.pem")
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			RootCAs:            caCertPool,
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}, nil
}

func LoadX509KeyPair(certFile, keyFile string) (tls.Certificate, error) {
	certPEMBlock, err := certificatesDir.ReadFile(certFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEMBlock, err := certificatesDir.ReadFile(keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
}

func getHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Cookie", "session=abcd")
	headers.Set("Origin", "https://192.168.1.3")
	headers.Set("Sec-WebSocket-Protocol", "lws-bidirectional-protocol")
	return headers
}

func generateHexString(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (c Client) sendRequest(v any) error {
	return c.conn.WriteJSON(map[string]interface{}{
		"Params": v,
	})
}

func (c *Client) send(request Request, response interface{}) error {
	request.SourceBox = c.ID
	request.Token = "LAN"

	err := c.sendRequest(request)
	if err != nil {
		return fmt.Errorf("requests failed: %w", err)
	}

	rawResponse := <-c.results

	err = json.Unmarshal(rawResponse, response)
	if err != nil {
		return err
	}

	return nil
}
