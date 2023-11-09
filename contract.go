// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"github.com/metronome/metronome-go/option"
)

// ContractService contains methods and other services that help with interacting
// with the metronome API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewContractService] method instead.
type ContractService struct {
	Options         []option.RequestOption
	Gets            *ContractGetService
	Lists           *ContractListService
	Creates         *ContractCreateService
	Amends          *ContractAmendService
	SetUsageFilters *ContractSetUsageFilterService
	UpdateEndDates  *ContractUpdateEndDateService
}

// NewContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractService(opts ...option.RequestOption) (r *ContractService) {
	r = &ContractService{}
	r.Options = opts
	r.Gets = NewContractGetService(opts...)
	r.Lists = NewContractListService(opts...)
	r.Creates = NewContractCreateService(opts...)
	r.Amends = NewContractAmendService(opts...)
	r.SetUsageFilters = NewContractSetUsageFilterService(opts...)
	r.UpdateEndDates = NewContractUpdateEndDateService(opts...)
	return
}
