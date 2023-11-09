// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"github.com/metronome/metronome-go/option"
)

// ContractPricingProductService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractPricingProductService]
// method instead.
type ContractPricingProductService struct {
	Options []option.RequestOption
	Gets    *ContractPricingProductGetService
	Lists   *ContractPricingProductListService
	Creates *ContractPricingProductCreateService
	Updates *ContractPricingProductUpdateService
}

// NewContractPricingProductService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractPricingProductService(opts ...option.RequestOption) (r *ContractPricingProductService) {
	r = &ContractPricingProductService{}
	r.Options = opts
	r.Gets = NewContractPricingProductGetService(opts...)
	r.Lists = NewContractPricingProductListService(opts...)
	r.Creates = NewContractPricingProductCreateService(opts...)
	r.Updates = NewContractPricingProductUpdateService(opts...)
	return
}
