// File generated from our OpenAPI spec by Stainless.

package example

import (
  "github.com/example/example-go/internal/shared"
  "github.com/example/example-go/internal/apijson"
  "github.com/example/example-go/internal/param"
  "context"
  "github.com/example/example-go/option"
  "github.com/example/example-go/internal/requestconfig"
)

// ContractPricingRateCardMoveRateCardProductService contains methods and other
// services that help with interacting with the example API. Note, unlike clients,
// this service does not read variables from the environment automatically. You
// should not instantiate this service directly, and instead use the
// [NewContractPricingRateCardMoveRateCardProductService] method instead.
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
func (r *ContractPricingRateCardMoveRateCardProductService) Update(ctx context.Context, body ContractPricingRateCardMoveRateCardProductUpdateParams, opts ...option.RequestOption) (res *ContractPricingRateCardMoveRateCardProductUpdateResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/moveRateCardProducts"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type ContractPricingRateCardMoveRateCardProductUpdateResponse struct {
Data shared.ID `json:"data,required"`
JSON contractPricingRateCardMoveRateCardProductUpdateResponseJSON
}

// contractPricingRateCardMoveRateCardProductUpdateResponseJSON contains the JSON
// metadata for the struct
// [ContractPricingRateCardMoveRateCardProductUpdateResponse]
type contractPricingRateCardMoveRateCardProductUpdateResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardMoveRateCardProductUpdateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardMoveRateCardProductUpdateParams struct {
ProductMoves param.Field[[]ContractPricingRateCardMoveRateCardProductUpdateParamsProductMove] `json:"product_moves,required"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractPricingRateCardMoveRateCardProductUpdateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardMoveRateCardProductUpdateParamsProductMove struct {
// 0-based index of the new position of the product
Position param.Field[float64] `json:"position,required"`
// ID of the product to move
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r ContractPricingRateCardMoveRateCardProductUpdateParamsProductMove) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
