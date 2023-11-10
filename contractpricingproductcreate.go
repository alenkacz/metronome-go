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

// ContractPricingProductCreateService contains methods and other services that
// help with interacting with the metronome API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingProductCreateService] method instead.
type ContractPricingProductCreateService struct {
Options []option.RequestOption
}

// NewContractPricingProductCreateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingProductCreateService(opts ...option.RequestOption) (r *ContractPricingProductCreateService) {
  r = &ContractPricingProductCreateService{}
  r.Options = opts
  return
}

// Create a new product
func (r *ContractPricingProductCreateService) NewProduct(ctx context.Context, body ContractPricingProductCreateNewProductParams, opts ...option.RequestOption) (res *ContractPricingProductCreateNewProductResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contract-pricing/products/create"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type ContractPricingProductCreateNewProductResponse struct {
Data ContractPricingProductCreateNewProductResponseData `json:"data,required"`
JSON contractPricingProductCreateNewProductResponseJSON
}

// contractPricingProductCreateNewProductResponseJSON contains the JSON metadata
// for the struct [ContractPricingProductCreateNewProductResponse]
type contractPricingProductCreateNewProductResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductCreateNewProductResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductCreateNewProductResponseData struct {
ID string `json:"id,required" format:"uuid"`
JSON contractPricingProductCreateNewProductResponseDataJSON
}

// contractPricingProductCreateNewProductResponseDataJSON contains the JSON
// metadata for the struct [ContractPricingProductCreateNewProductResponseData]
type contractPricingProductCreateNewProductResponseDataJSON struct {
ID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductCreateNewProductResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductCreateNewProductParams struct {
// displayed on invoices
Name param.Field[string] `json:"name,required"`
Type param.Field[ContractPricingProductCreateNewProductParamsType] `json:"type,required"`
// Required for USAGE products
BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
// Required for COMPOSITE products
CompositeProductIDs param.Field[[]string] `json:"composite_product_ids" format:"uuid"`
// Required for COMPOSITE products
CompositeTags param.Field[[]string] `json:"composite_tags"`
// Defaults to true
IsRefundable param.Field[bool] `json:"is_refundable"`
NetsuiteInternalItemID param.Field[string] `json:"netsuite_internal_item_id"`
NetsuiteOverageItemID param.Field[string] `json:"netsuite_overage_item_id"`
// Timestamp representing when the product will start existing. If not provided,
// defaults to the current time and the product will be added to the product list
// immediately.
StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
Tags param.Field[[]string] `json:"tags"`
}

func (r ContractPricingProductCreateNewProductParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractPricingProductCreateNewProductParamsType string

const (
  ContractPricingProductCreateNewProductParamsTypeFixed ContractPricingProductCreateNewProductParamsType = "FIXED"
  ContractPricingProductCreateNewProductParamsTypeFixed ContractPricingProductCreateNewProductParamsType = "fixed"
  ContractPricingProductCreateNewProductParamsTypeUsage ContractPricingProductCreateNewProductParamsType = "USAGE"
  ContractPricingProductCreateNewProductParamsTypeUsage ContractPricingProductCreateNewProductParamsType = "usage"
  ContractPricingProductCreateNewProductParamsTypeComposite ContractPricingProductCreateNewProductParamsType = "COMPOSITE"
  ContractPricingProductCreateNewProductParamsTypeComposite ContractPricingProductCreateNewProductParamsType = "composite"
)
