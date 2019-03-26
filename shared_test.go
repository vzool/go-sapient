package sapient

var originalMessage = []byte(`
	{
		"message": "Salam World!!!",
		"time": 3.14159265359,
		"defer": true
	}
`)

// TestServerPingURL this is URL to PING to an API endpoint
var TestServerPingURL = "http://localhost:8080/chronicle"

// TestServerPublishURL this is a URL to POST to an API endpoint
var TestServerPublishURL = "http://localhost:8080/chronicle/publish"

// TestServerPublicKey this is a Server Public key for Siging
var TestServerPublicKey = "dnt0NT7pUzs6W5LS02gZDX6O_wRJ1FbrOqFZBQrLLvo="

/* func TestToPingSapientClient(t *testing.T) {

	// Create a new HTTP client with a default timeout
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	res, err := client.Get(TestServerPingURL, nil)
	if err != nil {
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, _ := ioutil.ReadAll(res.Body)
	// auth := res.Header.Get(HeaderAuthName)
	signature := res.Header.Get(HeaderSignatureName)

	sig, _ := Base64UrlDecode(signature)

	pk := NewSigningPublicKey(TestServerPublicKey)

	t.Error(
		pk.Verify(body, sig),
	)
}
*/
