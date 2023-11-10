// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "github.com/metronome/metronome-go/internal/param"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// ContractPricingRateCardMoveRateCardProductService contains methods and other
// services that help with interacting with the metronome API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractPricingRateCardMoveRateCardProductService] method instead.
type ContractPricingRateCardMoveRateCardProductService struct {
Options []option.RequestOption
}

// NewContractPricingRateCardMoveRateCardProductService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewContractPricingRateCardMoveRateCardProductService(opts ...option.RequestOption) (r *ContractPricingRateCardMoveRateCardProductService) {
  r = &ContractPricingRateCardMoveRateCardProductService{}
  r.Options = opts
  return
}

// Updates ordering of specified products
func (r *ContractPricingRateCardMoveRateCardProductService) MoveRateCardProducts(ctx context.Context, body ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParams, opts ...option.RequestOption) (res *ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/moveRateCardProducts"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponse struct {
Data ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseData `json:"data,required"`
JSON contractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseJSON
}

// contractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponse]
type contractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseData struct {
ID string `json:"id,required" format:"uuid"`
JSON contractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseDataJSON
}

// contractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseDataJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseData]
type contractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseDataJSON struct {
ID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardMoveRateCardProductMoveRateCardProductsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParams struct {
ProductMoves param.Field[[]ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParamsProductMove] `json:"product_moves,required"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParamsProductMove struct {
// 0-based index of the new position of the product
Position param.Field[float64] `json:"position,required"`
// ID of the product to move
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParamsProductMove) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
