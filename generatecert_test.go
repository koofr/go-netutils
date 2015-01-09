package netutils

import (
	"crypto/tls"
	"crypto/x509"
	"strings"
	"testing"
	"time"
)

func TestGenerateCert(t *testing.T) {
	organization := "Organization"
	host := "localhost"
	validFor := 365 * 24 * time.Hour
	rsaBits := 512

	certBytes, keyBytes, err := GenerateCert(organization, host, validFor, rsaBits)

	if err != nil {
		t.Error(err)
	}

	if !strings.HasPrefix(string(certBytes), "-----BEGIN CERTIFICATE-----") {
		t.Error("Invalid certificate output")
	}

	if !strings.HasPrefix(string(keyBytes), "-----BEGIN RSA PRIVATE KEY-----") {
		t.Error("Invalid certificate output")
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)

	if err != nil {
		t.Error(err)
	}

	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])

	if err != nil {
		t.Error(err)
	}

	if x509Cert.Subject.Organization[0] != organization {
		t.Error("Invalid organization")
	}
}
