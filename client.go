// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"os"

	"github.com/metronome/metronome-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the metronome API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options          []option.RequestOption
	ContractPricings *ContractPricingService
	Contracts        *ContractService
}

// NewClient generates a new client with the default option read from the
// environment (METRONOME_API_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("METRONOME_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.ContractPricings = NewContractPricingService(opts...)
	r.Contracts = NewContractService(opts...)

	return
}
