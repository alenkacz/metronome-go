// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "github.com/metronome/metronome-go/internal/param"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// ContractPricingRateCardSetRateCardProductsOrderService contains methods and
// other services that help with interacting with the metronome API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractPricingRateCardSetRateCardProductsOrderService] method instead.
type ContractPricingRateCardSetRateCardProductsOrderService struct {
Options []option.RequestOption
}

// NewContractPricingRateCardSetRateCardProductsOrderService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewContractPricingRateCardSetRateCardProductsOrderService(opts ...option.RequestOption) (r *ContractPricingRateCardSetRateCardProductsOrderService) {
  r = &ContractPricingRateCardSetRateCardProductsOrderService{}
  r.Options = opts
  return
}

// Sets the ordering of products within a rate card
func (r *ContractPricingRateCardSetRateCardProductsOrderService) SetRateCardProductsOrder(ctx context.Context, body ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderParams, opts ...option.RequestOption) (res *ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/setRateCardProductsOrder"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponse struct {
Data ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseData `json:"data,required"`
JSON contractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseJSON
}

// contractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponse]
type contractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseData struct {
ID string `json:"id,required" format:"uuid"`
JSON contractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseDataJSON
}

// contractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseDataJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseData]
type contractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseDataJSON struct {
ID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderParams struct {
ProductOrder param.Field[[]string] `json:"product_order,required" format:"uuid"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractPricingRateCardSetRateCardProductsOrderSetRateCardProductsOrderParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
