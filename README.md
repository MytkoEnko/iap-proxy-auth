# IAP Auth Client Library for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/iap-auth-client)](https://goreportcard.com/report/github.com/yourusername/iap-auth-client)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A Go library that simplifies handling authentication to Identity-Aware Proxy (IAP) with the ability to set the proxy authorization header for IAP tokens. It provides two main functionalities:

1. **New Client with IAP Authentication:**
   - Create a new HTTP client with IAP authentication configured.
   - Automatically sets the required headers for IAP token-based authentication.

2. **Update Existing Client:**
   - Update an existing HTTP client's transport with IAP authentication.
   - Allows seamless integration with existing HTTP clients while adding IAP authentication capabilities.

## Installation

To install the library, use the `go get` command:

```bash
go get -u github.com/yourusername/iap-auth-client
