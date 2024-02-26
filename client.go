// Package proxiap provides methods and structs to handle authentication
// to resources behind Google's Identity-Aware Proxy (IAP).
package proxiap

import (
	"context"
	"net/http"
)

// NewIapClient returns a new http.Client with a Transport that handles IAP authentication.
// If iapId is empty, it returns nil.
func NewIapClient(ctx context.Context, iapId string) *http.Client {
	if len(iapId) > 0 {
		return &http.Client{Transport: &IapAuthTransport{
			Tokensource: TokensourceInit(ctx, iapId),
		}}
	}
	return nil
}

// SetIapTransport replaces http.Transport in the provided http.Client with
// a one that will handle authentication behind IAP.
// If iapId is empty, it returns the same clie
func SetIapTransport(ctx context.Context, iapId string, client *http.Client) {
	if len(iapId) > 0 {
		client.Transport = &IapAuthTransport{
			Tokensource: TokensourceInit(ctx, iapId),
		}
	}
}
