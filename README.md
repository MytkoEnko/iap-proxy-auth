# IAP Auth Client Library for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/mytkoenko/iap-proxy-auth)](https://goreportcard.com/report/github.com/mytkoenko/iap-proxy-auth)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This library resolves conflicts when using Identity-Aware Proxy (IAP) alongside additional layers of authentication. By default, IAP uses the `Authorization` header for its tokens (source: [googleapis/google-cloud-go](https://github.com/googleapis/google-cloud-go/blob/6cd6a73be87a261729d3b6b45f3d28be93c3fdb3/auth/httptransport/httptransport.go#L179C4-L187)), potentially causing issues with secondary authentication layers that tend to use the same header. This solution moves IAP tokens to the `Proxy-Authorization` header, enabling seamless interaction with both IAP and additional authentication layers. Users can now set the `Authorization` header with credentials for the secondary layer.

 It provides two main functionalities:

   - Creates a new HTTP client with IAP authorization moved to proper header.
   - Updates an existing HTTP client's transport with proper IAP header.


## Installation

Just import it it as any other library

## Usage

Caution: Ensure this code runs in an environment with a Service Account (SA) capable of IAP authentication. It relies on standard Google credential sources, seamlessly obtaining credentials within the appropriate setup.

Create new client:

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/MytkoEnko/iap-proxy-auth"
)

func main() {

	// Context is requird
	ctx := context.Background()

	// IAP client ID of the resource is required
	iapID := "123456789012-abc123def456ghijklmnopqrstuvwxyz.apps.googleusercontent.com"
	// Create an HTTP client with proxied IAP headers.
	client := proxiap.NewIapClient(ctx, iapID)

	// Make a sample request to a resource protected by IAP
	req, err := http.NewRequestWithContext(ctx, "GET", "https://example.com/protected/resource", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Add any necessary headers for your second layer authentication
	req.Header.Set("Authorization", "Bearer your_second_layer_token")

	// Send the request
	resp, err := client.Do(req)

	// Use your client ...
}
```

Update existing client:

```go
package main

import (
	"context"
	"net/http"
	"github.com/MytkoEnko/iap-proxy-auth"
)

func main() {
	// Context is requird
	ctx := context.Background()

	// IAP client ID of the resource is required
	iapID := "123456789012-abc123def456ghijklmnopqrstuvwxyz.apps.googleusercontent.com"
	// Create a new http.Client
	client := &http.Client{}

	// Update cient's transport with proxiap.SetIapTransport()
	proxiap.SetIapTransport(ctx, iapID, *client)

	// Use your client ...
}
```
