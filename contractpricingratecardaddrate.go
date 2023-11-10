// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "github.com/metronome/metronome-go/internal/param"
  "time"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// ContractPricingRateCardAddRateService contains methods and other services that
// help with interacting with the metronome API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardAddRateService] method instead.
type ContractPricingRateCardAddRateService struct {
Options []option.RequestOption
}

// NewContractPricingRateCardAddRateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingRateCardAddRateService(opts ...option.RequestOption) (r *ContractPricingRateCardAddRateService) {
  r = &ContractPricingRateCardAddRateService{}
  r.Options = opts
  return
}

// Add a new rate
func (r *ContractPricingRateCardAddRateService) AddRate(ctx context.Context, body ContractPricingRateCardAddRateAddRateParams, opts ...option.RequestOption) (res *ContractPricingRateCardAddRateAddRateResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/addRate"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type ContractPricingRateCardAddRateAddRateResponse struct {
Data ContractPricingRateCardAddRateAddRateResponseData `json:"data,required"`
JSON contractPricingRateCardAddRateAddRateResponseJSON
}

// contractPricingRateCardAddRateAddRateResponseJSON contains the JSON metadata for
// the struct [ContractPricingRateCardAddRateAddRateResponse]
type contractPricingRateCardAddRateAddRateResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardAddRateAddRateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardAddRateAddRateResponseData struct {
// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
Price float64 `json:"price,required"`
RateType ContractPricingRateCardAddRateAddRateResponseDataRateType `json:"rate_type,required"`
// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
// using list prices rather than the standard rates for this product on the
// contract.
UseListPrices bool `json:"use_list_prices"`
JSON contractPricingRateCardAddRateAddRateResponseDataJSON
}

// contractPricingRateCardAddRateAddRateResponseDataJSON contains the JSON metadata
// for the struct [ContractPricingRateCardAddRateAddRateResponseData]
type contractPricingRateCardAddRateAddRateResponseDataJSON struct {
Price apijson.Field
RateType apijson.Field
UseListPrices apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardAddRateAddRateResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardAddRateAddRateResponseDataRateType string

const (
  ContractPricingRateCardAddRateAddRateResponseDataRateTypeFlat ContractPricingRateCardAddRateAddRateResponseDataRateType = "FLAT"
  ContractPricingRateCardAddRateAddRateResponseDataRateTypeFlat ContractPricingRateCardAddRateAddRateResponseDataRateType = "flat"
  ContractPricingRateCardAddRateAddRateResponseDataRateTypePercentage ContractPricingRateCardAddRateAddRateResponseDataRateType = "PERCENTAGE"
  ContractPricingRateCardAddRateAddRateResponseDataRateTypePercentage ContractPricingRateCardAddRateAddRateResponseDataRateType = "percentage"
)

type ContractPricingRateCardAddRateAddRateParams struct {
Entitled param.Field[bool] `json:"entitled,required"`
// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
Price param.Field[float64] `json:"price,required"`
// ID of the product to add a rate for
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
RateType param.Field[ContractPricingRateCardAddRateAddRateParamsRateType] `json:"rate_type,required"`
// inclusive effective date
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
// exclusive end date
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
// using list prices rather than the standard rates for this product on the
// contract.
UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractPricingRateCardAddRateAddRateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardAddRateAddRateParamsRateType string

const (
  ContractPricingRateCardAddRateAddRateParamsRateTypeFlat ContractPricingRateCardAddRateAddRateParamsRateType = "FLAT"
  ContractPricingRateCardAddRateAddRateParamsRateTypeFlat ContractPricingRateCardAddRateAddRateParamsRateType = "flat"
  ContractPricingRateCardAddRateAddRateParamsRateTypePercentage ContractPricingRateCardAddRateAddRateParamsRateType = "PERCENTAGE"
  ContractPricingRateCardAddRateAddRateParamsRateTypePercentage ContractPricingRateCardAddRateAddRateParamsRateType = "percentage"
)
