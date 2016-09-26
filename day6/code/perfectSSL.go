package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// HTTP Strict Transport Security (HSTS) is a web security policy
		// mechanism which helps to protect websites against protocol downgrade
		// attacks and cookie hijacking.
		// HSTS is an IETF standards track protocol and is specified in RFC 6797.
		// The HSTS Policy[2] is communicated by the server to the user agent via
		// an HTTP response header field named "Strict-Transport-Security".
		// HSTS Policy specifies a period of time during which the user agent should
		// only access the server in a secure fashion.
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is a perfect  A+ SSL example server.\n"))
	})
	// https://www.ssllabs.com/downloads/SSL_Server_Rating_Guide.pdf
	//	Start with the score of the strongest cipher.
	//	Add the score of the weakest cipher.
	//	Divide the total by 2.
	// >= 256 bits (e.g., 256) = 100%
	cfg := &tls.Config{
		// Getting 100% requires that the server only supports TLS 1.2 or above.
		MinVersion: tls.VersionTLS12,
		// Key or DH parameter strength >= 4096 bits
		// change the CurvePreferences, to prioritise using the elliptic curve algorithms
		// with an approximate comparable symmetric key size of >=4096 bit.
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		// Disabling HTTP/2 also disables the server cipher preference.
		//This can be corrected by setting the PreferServerCipherSuites option in tls.Config.
		PreferServerCipherSuites: true,
		// Remove from default CipherSuites any that use a cipher smaller than 256bit.
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:      ":4443",
		Handler:   mux,
		TLSConfig: cfg,
		// HTTP/2 was enabled by default in go 1.6, however HTTP/2 mandates the support of the
		// cipher suite TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256. As this is a 128bit cipher,
		// it needs to be removed. The only way to achieve a perfect score now is to disable HTTP/2.
		// Programs that must disable HTTP/2 can do so by setting Transport.TLSNextProto (for clients)
		// or Server.TLSNextProto (for servers) to a non-nil, empty map.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("certificate.pem", "private.key"))
}
