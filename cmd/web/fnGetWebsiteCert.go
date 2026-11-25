package web

import (
	jsonv1 "encoding/json"
	json "encoding/json/v2"
	"net/http"
	"net/url"
	"time"
)

func GetWebsiteCertificate(siteUrl string, proxy *string) (certMinimal WebsiteCertificateMinimal) {
	var cl http.Client
	if proxy != nil {
		proxyUrl, err := url.Parse(*proxy)
		if err != nil {
			// TODO handle me
			panic(err)
		}

		cl = http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
			Timeout: 3000 * time.Millisecond,
		}
	} else {
		cl = http.Client{
			// Transport: &http.Transport{
			// 	Proxy: http.ProxyURL(proxyUrl),
			// },
			Timeout: 3000 * time.Millisecond,
		}
	}

	resp, err := cl.Get(siteUrl)
	if err != nil {
		// TODO handle me
		panic(err)
	}

	cert := resp.TLS.PeerCertificates[0]
	// certs := resp.TLS.PeerCertificates

	// lib.JsonMarshalAndPrint(certs)

	// x509.Certificate has no json tags, so its fields are written as
	// SerialNumber, NotAfter and so on, and read back here into camel case
	// tags. That, and marshalling whatever public key type the certificate
	// holds, relies on json v1 behaviour, so v1 semantics are kept for this
	// round trip.
	jsonStr, _ := json.Marshal(cert, jsonv1.DefaultOptionsV1())
	err = json.Unmarshal(jsonStr, &certMinimal, jsonv1.DefaultOptionsV1())
	// lib.CheckFatalError(err)
	certMinimal.Subject = cert.Subject.CommonName
	certMinimal.Issuer = cert.Issuer.CommonName

	return certMinimal

}
