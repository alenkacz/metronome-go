// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "context"
  "github.com/metronome/metronome-go/option"
)

// ContractPricingRateCardService contains methods and other services that help
// with interacting with the metronome API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardService] method instead.
type ContractPricingRateCardService struct {
Options []option.RequestOption
Gets *ContractPricingRateCardGetService
Lists *ContractPricingRateCardListService
Creates *ContractPricingRateCardCreateService
Updates *ContractPricingRateCardUpdateService
AddRates *ContractPricingRateCardAddRateService
SetRateCardProductsOrders *ContractPricingRateCardSetRateCardProductsOrderService
MoveRateCardProducts *ContractPricingRateCardMoveRateCardProductService
}

// NewContractPricingRateCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractPricingRateCardService(opts ...option.RequestOption) (r *ContractPricingRateCardService) {
  r = &ContractPricingRateCardService{}
  r.Options = opts
  r.Gets = NewContractPricingRateCardGetService(opts...)
  r.Lists = NewContractPricingRateCardListService(opts...)
  r.Creates = NewContractPricingRateCardCreateService(opts...)
  r.Updates = NewContractPricingRateCardUpdateService(opts...)
  r.AddRates = NewContractPricingRateCardAddRateService(opts...)
  r.SetRateCardProductsOrders = NewContractPricingRateCardSetRateCardProductsOrderService(opts...)
  r.MoveRateCardProducts = NewContractPricingRateCardMoveRateCardProductService(opts...)
  return
}
