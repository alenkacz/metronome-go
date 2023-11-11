// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "time"
  "github.com/metronome/metronome-go/internal/apijson"
  "github.com/metronome/metronome-go/internal/shared"
  "github.com/metronome/metronome-go/internal/param"
  "net/url"
  "github.com/metronome/metronome-go/internal/apiquery"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// ContractPricingRateCardService contains methods and other services that help
// with interacting with the metronome API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardService] method instead.
type ContractPricingRateCardService struct {
Options []option.RequestOption
}

// NewContractPricingRateCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractPricingRateCardService(opts ...option.RequestOption) (r *ContractPricingRateCardService) {
  r = &ContractPricingRateCardService{}
  r.Options = opts
  return
}

// Create a new rate card
func (r *ContractPricingRateCardService) New(ctx context.Context, body ContractPricingRateCardNewParams, opts ...option.RequestOption) (res *ContractPricingRateCardNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/create"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Get a specific rate card
func (r *ContractPricingRateCardService) Get(ctx context.Context, body ContractPricingRateCardGetParams, opts ...option.RequestOption) (res *ContractPricingRateCardGetResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/get"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Update a rate card
func (r *ContractPricingRateCardService) Update(ctx context.Context, body ContractPricingRateCardUpdateParams, opts ...option.RequestOption) (res *ContractPricingRateCardUpdateResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/update"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// List rate cards
func (r *ContractPricingRateCardService) List(ctx context.Context, params ContractPricingRateCardListParams, opts ...option.RequestOption) (res *ContractPricingRateCardListResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/list"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
  return
}

// Add a new rate
func (r *ContractPricingRateCardService) AddRate(ctx context.Context, body ContractPricingRateCardAddRateParams, opts ...option.RequestOption) (res *ContractPricingRateCardAddRateResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/addRate"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Updates ordering of specified products
func (r *ContractPricingRateCardService) MoveRateCardProducts(ctx context.Context, body ContractPricingRateCardMoveRateCardProductsParams, opts ...option.RequestOption) (res *ContractPricingRateCardMoveRateCardProductsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/moveRateCardProducts"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Sets the ordering of products within a rate card
func (r *ContractPricingRateCardService) SetRateCardProductsOrder(ctx context.Context, body ContractPricingRateCardSetRateCardProductsOrderParams, opts ...option.RequestOption) (res *ContractPricingRateCardSetRateCardProductsOrderResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/rate-cards/setRateCardProductsOrder"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type RateCard struct {
ID string `json:"id,required" format:"uuid"`
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
Name string `json:"name,required"`
RateCardEntries map[string]RateCardRateCardEntry `json:"rate_card_entries,required"`
Description string `json:"description"`
JSON rateCardJSON
}

// rateCardJSON contains the JSON metadata for the struct [RateCard]
type rateCardJSON struct {
ID apijson.Field
CreatedAt apijson.Field
CreatedBy apijson.Field
Name apijson.Field
RateCardEntries apijson.Field
Description apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RateCard) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type RateCardRateCardEntry struct {
Current RateCardRateCardEntriesCurrent `json:"current,nullable"`
Updates []RateCardRateCardEntriesUpdate `json:"updates"`
JSON rateCardRateCardEntryJSON
}

// rateCardRateCardEntryJSON contains the JSON metadata for the struct
// [RateCardRateCardEntry]
type rateCardRateCardEntryJSON struct {
Current apijson.Field
Updates apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RateCardRateCardEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type RateCardRateCardEntriesCurrent struct {
ID string `json:"id" format:"uuid"`
CreatedAt time.Time `json:"created_at" format:"date-time"`
CreatedBy string `json:"created_by"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
Entitled bool `json:"entitled"`
Price float64 `json:"price"`
ProductID string `json:"product_id" format:"uuid"`
RateType RateCardRateCardEntriesCurrentRateType `json:"rate_type"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON rateCardRateCardEntriesCurrentJSON
}

// rateCardRateCardEntriesCurrentJSON contains the JSON metadata for the struct
// [RateCardRateCardEntriesCurrent]
type rateCardRateCardEntriesCurrentJSON struct {
ID apijson.Field
CreatedAt apijson.Field
CreatedBy apijson.Field
EndingBefore apijson.Field
Entitled apijson.Field
Price apijson.Field
ProductID apijson.Field
RateType apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RateCardRateCardEntriesCurrent) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type RateCardRateCardEntriesCurrentRateType string

const (
  RateCardRateCardEntriesCurrentRateTypeFlat RateCardRateCardEntriesCurrentRateType = "FLAT"
  RateCardRateCardEntriesCurrentRateTypePercentage RateCardRateCardEntriesCurrentRateType = "PERCENTAGE"
)

type RateCardRateCardEntriesUpdate struct {
ID string `json:"id,required" format:"uuid"`
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
Entitled bool `json:"entitled,required"`
Price float64 `json:"price,required"`
ProductID string `json:"product_id,required" format:"uuid"`
RateType RateCardRateCardEntriesUpdatesRateType `json:"rate_type,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
JSON rateCardRateCardEntriesUpdateJSON
}

// rateCardRateCardEntriesUpdateJSON contains the JSON metadata for the struct
// [RateCardRateCardEntriesUpdate]
type rateCardRateCardEntriesUpdateJSON struct {
ID apijson.Field
CreatedAt apijson.Field
CreatedBy apijson.Field
Entitled apijson.Field
Price apijson.Field
ProductID apijson.Field
RateType apijson.Field
StartingAt apijson.Field
EndingBefore apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RateCardRateCardEntriesUpdate) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type RateCardRateCardEntriesUpdatesRateType string

const (
  RateCardRateCardEntriesUpdatesRateTypeFlat RateCardRateCardEntriesUpdatesRateType = "FLAT"
  RateCardRateCardEntriesUpdatesRateTypePercentage RateCardRateCardEntriesUpdatesRateType = "PERCENTAGE"
)

type ContractPricingRateCardNewResponse struct {
Data shared.ID `json:"data,required"`
JSON contractPricingRateCardNewResponseJSON
}

// contractPricingRateCardNewResponseJSON contains the JSON metadata for the struct
// [ContractPricingRateCardNewResponse]
type contractPricingRateCardNewResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardGetResponse struct {
Data RateCard `json:"data,required"`
JSON contractPricingRateCardGetResponseJSON
}

// contractPricingRateCardGetResponseJSON contains the JSON metadata for the struct
// [ContractPricingRateCardGetResponse]
type contractPricingRateCardGetResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardUpdateResponse struct {
Data shared.ID `json:"data,required"`
JSON contractPricingRateCardUpdateResponseJSON
}

// contractPricingRateCardUpdateResponseJSON contains the JSON metadata for the
// struct [ContractPricingRateCardUpdateResponse]
type contractPricingRateCardUpdateResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardUpdateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardListResponse struct {
Data []RateCard `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON contractPricingRateCardListResponseJSON
}

// contractPricingRateCardListResponseJSON contains the JSON metadata for the
// struct [ContractPricingRateCardListResponse]
type contractPricingRateCardListResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardAddRateResponse struct {
Data shared.Rate `json:"data,required"`
JSON contractPricingRateCardAddRateResponseJSON
}

// contractPricingRateCardAddRateResponseJSON contains the JSON metadata for the
// struct [ContractPricingRateCardAddRateResponse]
type contractPricingRateCardAddRateResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardAddRateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardMoveRateCardProductsResponse struct {
Data shared.ID `json:"data,required"`
JSON contractPricingRateCardMoveRateCardProductsResponseJSON
}

// contractPricingRateCardMoveRateCardProductsResponseJSON contains the JSON
// metadata for the struct [ContractPricingRateCardMoveRateCardProductsResponse]
type contractPricingRateCardMoveRateCardProductsResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardMoveRateCardProductsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardSetRateCardProductsOrderResponse struct {
Data shared.ID `json:"data,required"`
JSON contractPricingRateCardSetRateCardProductsOrderResponseJSON
}

// contractPricingRateCardSetRateCardProductsOrderResponseJSON contains the JSON
// metadata for the struct
// [ContractPricingRateCardSetRateCardProductsOrderResponse]
type contractPricingRateCardSetRateCardProductsOrderResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardSetRateCardProductsOrderResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardNewParams struct {
// Used only in UI/API. It is not exposed to end customers.
Name param.Field[string] `json:"name,required"`
Description param.Field[string] `json:"description"`
}

func (r ContractPricingRateCardNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardGetParams struct {
ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r ContractPricingRateCardGetParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardUpdateParams struct {
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
Description param.Field[string] `json:"description"`
// Used only in UI/API. It is not exposed to end customers.
Name param.Field[string] `json:"name"`
}

func (r ContractPricingRateCardUpdateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardListParams struct {
Body param.Field[interface{}] `json:"body,required"`
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
}

func (r ContractPricingRateCardListParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ContractPricingRateCardListParams]'s query parameters as
// `url.Values`.
func (r ContractPricingRateCardListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type ContractPricingRateCardAddRateParams struct {
Entitled param.Field[bool] `json:"entitled,required"`
// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
Price param.Field[float64] `json:"price,required"`
// ID of the product to add a rate for
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
RateType param.Field[ContractPricingRateCardAddRateParamsRateType] `json:"rate_type,required"`
// inclusive effective date
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
// exclusive end date
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
// using list prices rather than the standard rates for this product on the
// contract.
UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractPricingRateCardAddRateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardAddRateParamsRateType string

const (
  ContractPricingRateCardAddRateParamsRateTypeFlat ContractPricingRateCardAddRateParamsRateType = "FLAT"
  ContractPricingRateCardAddRateParamsRateTypeFlat ContractPricingRateCardAddRateParamsRateType = "flat"
  ContractPricingRateCardAddRateParamsRateTypePercentage ContractPricingRateCardAddRateParamsRateType = "PERCENTAGE"
  ContractPricingRateCardAddRateParamsRateTypePercentage ContractPricingRateCardAddRateParamsRateType = "percentage"
)

type ContractPricingRateCardMoveRateCardProductsParams struct {
ProductMoves param.Field[[]ContractPricingRateCardMoveRateCardProductsParamsProductMove] `json:"product_moves,required"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractPricingRateCardMoveRateCardProductsParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardMoveRateCardProductsParamsProductMove struct {
// 0-based index of the new position of the product
Position param.Field[float64] `json:"position,required"`
// ID of the product to move
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r ContractPricingRateCardMoveRateCardProductsParamsProductMove) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingRateCardSetRateCardProductsOrderParams struct {
ProductOrder param.Field[[]string] `json:"product_order,required" format:"uuid"`
// ID of the rate card to update
RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
}

func (r ContractPricingRateCardSetRateCardProductsOrderParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
