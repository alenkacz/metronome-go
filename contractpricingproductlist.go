// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "time"
  "github.com/metronome/metronome-go/internal/param"
  "net/url"
  "github.com/metronome/metronome-go/internal/apiquery"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// ContractPricingProductListService contains methods and other services that help
// with interacting with the metronome API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingProductListService] method instead.
type ContractPricingProductListService struct {
Options []option.RequestOption
}

// NewContractPricingProductListService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingProductListService(opts ...option.RequestOption) (r *ContractPricingProductListService) {
  r = &ContractPricingProductListService{}
  r.Options = opts
  return
}

// List products
func (r *ContractPricingProductListService) ListProducts(ctx context.Context, params ContractPricingProductListListProductsParams, opts ...option.RequestOption) (res *ContractPricingProductListListProductsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/products/list"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
  return
}

type ContractPricingProductListListProductsResponse struct {
Data []ContractPricingProductListListProductsResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON contractPricingProductListListProductsResponseJSON
}

// contractPricingProductListListProductsResponseJSON contains the JSON metadata
// for the struct [ContractPricingProductListListProductsResponse]
type contractPricingProductListListProductsResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductListListProductsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductListListProductsResponseData struct {
ID string `json:"id,required" format:"uuid"`
Current ContractPricingProductListListProductsResponseDataCurrent `json:"current,required"`
Initial ContractPricingProductListListProductsResponseDataInitial `json:"initial,required"`
Type ContractPricingProductListListProductsResponseDataType `json:"type,required"`
Updates []ContractPricingProductListListProductsResponseDataUpdate `json:"updates,required"`
JSON contractPricingProductListListProductsResponseDataJSON
}

// contractPricingProductListListProductsResponseDataJSON contains the JSON
// metadata for the struct [ContractPricingProductListListProductsResponseData]
type contractPricingProductListListProductsResponseDataJSON struct {
ID apijson.Field
Current apijson.Field
Initial apijson.Field
Type apijson.Field
Updates apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductListListProductsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductListListProductsResponseDataCurrent struct {
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
IsRefundable bool `json:"is_refundable,required"`
Name string `json:"name,required"`
BillableMetricID string `json:"billable_metric_id"`
CompositeProductIDs []string `json:"composite_product_ids" format:"uuid"`
CompositeTags []string `json:"composite_tags"`
NetsuiteInternalItemID string `json:"netsuite_internal_item_id"`
NetsuiteOverageItemID string `json:"netsuite_overage_item_id"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
Tags []string `json:"tags"`
JSON contractPricingProductListListProductsResponseDataCurrentJSON
}

// contractPricingProductListListProductsResponseDataCurrentJSON contains the JSON
// metadata for the struct
// [ContractPricingProductListListProductsResponseDataCurrent]
type contractPricingProductListListProductsResponseDataCurrentJSON struct {
CreatedAt apijson.Field
CreatedBy apijson.Field
IsRefundable apijson.Field
Name apijson.Field
BillableMetricID apijson.Field
CompositeProductIDs apijson.Field
CompositeTags apijson.Field
NetsuiteInternalItemID apijson.Field
NetsuiteOverageItemID apijson.Field
StartingAt apijson.Field
Tags apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductListListProductsResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductListListProductsResponseDataInitial struct {
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
IsRefundable bool `json:"is_refundable,required"`
Name string `json:"name,required"`
BillableMetricID string `json:"billable_metric_id"`
CompositeProductIDs []string `json:"composite_product_ids" format:"uuid"`
CompositeTags []string `json:"composite_tags"`
NetsuiteInternalItemID string `json:"netsuite_internal_item_id"`
NetsuiteOverageItemID string `json:"netsuite_overage_item_id"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
Tags []string `json:"tags"`
JSON contractPricingProductListListProductsResponseDataInitialJSON
}

// contractPricingProductListListProductsResponseDataInitialJSON contains the JSON
// metadata for the struct
// [ContractPricingProductListListProductsResponseDataInitial]
type contractPricingProductListListProductsResponseDataInitialJSON struct {
CreatedAt apijson.Field
CreatedBy apijson.Field
IsRefundable apijson.Field
Name apijson.Field
BillableMetricID apijson.Field
CompositeProductIDs apijson.Field
CompositeTags apijson.Field
NetsuiteInternalItemID apijson.Field
NetsuiteOverageItemID apijson.Field
StartingAt apijson.Field
Tags apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductListListProductsResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductListListProductsResponseDataType string

const (
  ContractPricingProductListListProductsResponseDataTypeUsage ContractPricingProductListListProductsResponseDataType = "USAGE"
  ContractPricingProductListListProductsResponseDataTypeComposite ContractPricingProductListListProductsResponseDataType = "COMPOSITE"
  ContractPricingProductListListProductsResponseDataTypeFixed ContractPricingProductListListProductsResponseDataType = "FIXED"
)

type ContractPricingProductListListProductsResponseDataUpdate struct {
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
BillableMetricID string `json:"billable_metric_id" format:"uuid"`
CompositeProductIDs []string `json:"composite_product_ids" format:"uuid"`
CompositeTags []string `json:"composite_tags"`
IsRefundable bool `json:"is_refundable"`
Name string `json:"name"`
NetsuiteInternalItemID string `json:"netsuite_internal_item_id"`
NetsuiteOverageItemID string `json:"netsuite_overage_item_id"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
Tags []string `json:"tags"`
JSON contractPricingProductListListProductsResponseDataUpdateJSON
}

// contractPricingProductListListProductsResponseDataUpdateJSON contains the JSON
// metadata for the struct
// [ContractPricingProductListListProductsResponseDataUpdate]
type contractPricingProductListListProductsResponseDataUpdateJSON struct {
CreatedAt apijson.Field
CreatedBy apijson.Field
BillableMetricID apijson.Field
CompositeProductIDs apijson.Field
CompositeTags apijson.Field
IsRefundable apijson.Field
Name apijson.Field
NetsuiteInternalItemID apijson.Field
NetsuiteOverageItemID apijson.Field
StartingAt apijson.Field
Tags apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductListListProductsResponseDataUpdate) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductListListProductsParams struct {
Body param.Field[interface{}] `json:"body,required"`
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
}

func (r ContractPricingProductListListProductsParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ContractPricingProductListListProductsParams]'s query
// parameters as `url.Values`.
func (r ContractPricingProductListListProductsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}
