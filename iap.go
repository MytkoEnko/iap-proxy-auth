// Package proxiap provides methods and structs to handle authentication to resources behind Google's Identity-Aware Proxy (IAP).
package proxiap

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"log"
)

// IapAuthTransport is a custom http.RoundTripper that adds IAP authentication
// to the http.Requests. It assumes a default service account with permission
// to access the resource.
type IapAuthTransport struct {
	Transport    http.RoundTripper
	Tokensource  oauth2.TokenSource
	CurrentToken oauth2.Token
}

// RoundTrip adds the IAP authorization header (obtained from TokenSource)
// to each request before it is sent.
func (t *IapAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req) // per RoundTrip contract
	token, err := t.Tokensource.Token()
	if err != nil {
		log.Panicf("Failed to get token from reuseTokenSource: %e", err)
	}

	if token.AccessToken != "" {
		req.Header.Add("Proxy-Authorization", "Bearer "+token.AccessToken)
	}
	return t.transport().RoundTrip(req)
}

// Transport returns the underlying transport if it is set, otherwise
// it returns the http.DefaultTransport.
func (t *IapAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided http.Request.
func cloneRequest(r *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *r
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

// TokensourceInit initializes a new TokenSource using the provided audience
// and client options. It panics if it encounters an error during initialization.
func TokensourceInit(ctx context.Context, audience string, opts ...option.ClientOption) oauth2.TokenSource {
	tokensource, err := idtoken.NewTokenSource(ctx, audience, opts...)
	if err != nil {
		log.Panicf("Failed to get NewTokenSource: %e", err)
	}
	reuseTokenSource := oauth2.ReuseTokenSource(nil, tokensource)
	return reuseTokenSource
}
