package web

import (
	"encoding/json"
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

	jsonStr, _ := json.Marshal(cert)
	err = json.Unmarshal(jsonStr, &certMinimal)
	// lib.CheckFatalError(err)
	certMinimal.Subject = cert.Subject.CommonName
	certMinimal.Issuer = cert.Issuer.CommonName

	return certMinimal

}
