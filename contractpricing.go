// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"github.com/metronome/metronome-go/option"
)

// ContractPricingService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractPricingService] method
// instead.
type ContractPricingService struct {
	Options   []option.RequestOption
	Products  *ContractPricingProductService
	RateCards *ContractPricingRateCardService
}

// NewContractPricingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractPricingService(opts ...option.RequestOption) (r *ContractPricingService) {
	r = &ContractPricingService{}
	r.Options = opts
	r.Products = NewContractPricingProductService(opts...)
	r.RateCards = NewContractPricingRateCardService(opts...)
	return
}
