package main

import (
	"crypto/tls"
	// "crypto/x509"
	// "io/ioutil"
	"net/http"
)

// newClient creates an http.Client configured for certificate authorization and
// verification against cert-api.access.redhat.com.
func newClient(cfg *config) (*http.Client, error) {
	// caCert, err := ioutil.ReadFile("/etc/insights-client/cert-api.access.redhat.com.pem")
	// if err != nil {
	// 	return nil, err
	// }
	// caCertPool, err := x509.SystemCertPool()
	// if err != nil {
	// 	return nil, err
	// }
	// caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := tls.Config{
		// RootCAs:      caCertPool,
		MaxVersion:   tls.VersionTLS12, // cloud.redhat.com appears to exhibit this openssl bug https://github.com/openssl/openssl/issues/9767
	}
	
	if cfg.AuthMethod == "CERT" {
		cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
		if err != nil {
			return nil, err
		}

		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	tlsConfig.BuildNameToCertificate()
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	client := http.Client{
		Transport: &transport,
	}
	return &client, nil
}
