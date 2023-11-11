// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/apiquery"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/internal/shared"
	"github.com/metronome/metronome-go/option"
)

// ContractPricingProductService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractPricingProductService]
// method instead.
type ContractPricingProductService struct {
	Options []option.RequestOption
}

// NewContractPricingProductService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractPricingProductService(opts ...option.RequestOption) (r *ContractPricingProductService) {
	r = &ContractPricingProductService{}
	r.Options = opts
	return
}

// Create a new product
func (r *ContractPricingProductService) New(ctx context.Context, body ContractPricingProductNewParams, opts ...option.RequestOption) (res *ContractPricingProductNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific product
func (r *ContractPricingProductService) Get(ctx context.Context, body ContractPricingProductGetParams, opts ...option.RequestOption) (res *ContractPricingProductGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a product
func (r *ContractPricingProductService) Update(ctx context.Context, body ContractPricingProductUpdateParams, opts ...option.RequestOption) (res *ContractPricingProductUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List products
func (r *ContractPricingProductService) List(ctx context.Context, params ContractPricingProductListParams, opts ...option.RequestOption) (res *ContractPricingProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ProductListItem struct {
	ID      string                  `json:"id,required" format:"uuid"`
	Current ProductListItemCurrent  `json:"current,required"`
	Initial ProductListItemInitial  `json:"initial,required"`
	Type    ProductListItemType     `json:"type,required"`
	Updates []ProductListItemUpdate `json:"updates,required"`
	JSON    productListItemJSON
}

// productListItemJSON contains the JSON metadata for the struct [ProductListItem]
type productListItemJSON struct {
	ID          apijson.Field
	Current     apijson.Field
	Initial     apijson.Field
	Type        apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListItemCurrent struct {
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy              string    `json:"created_by,required"`
	IsRefundable           bool      `json:"is_refundable,required"`
	Name                   string    `json:"name,required"`
	BillableMetricID       string    `json:"billable_metric_id"`
	CompositeProductIDs    []string  `json:"composite_product_ids" format:"uuid"`
	CompositeTags          []string  `json:"composite_tags"`
	NetsuiteInternalItemID string    `json:"netsuite_internal_item_id"`
	NetsuiteOverageItemID  string    `json:"netsuite_overage_item_id"`
	StartingAt             time.Time `json:"starting_at" format:"date-time"`
	Tags                   []string  `json:"tags"`
	JSON                   productListItemCurrentJSON
}

// productListItemCurrentJSON contains the JSON metadata for the struct
// [ProductListItemCurrent]
type productListItemCurrentJSON struct {
	CreatedAt              apijson.Field
	CreatedBy              apijson.Field
	IsRefundable           apijson.Field
	Name                   apijson.Field
	BillableMetricID       apijson.Field
	CompositeProductIDs    apijson.Field
	CompositeTags          apijson.Field
	NetsuiteInternalItemID apijson.Field
	NetsuiteOverageItemID  apijson.Field
	StartingAt             apijson.Field
	Tags                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProductListItemCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListItemInitial struct {
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy              string    `json:"created_by,required"`
	IsRefundable           bool      `json:"is_refundable,required"`
	Name                   string    `json:"name,required"`
	BillableMetricID       string    `json:"billable_metric_id"`
	CompositeProductIDs    []string  `json:"composite_product_ids" format:"uuid"`
	CompositeTags          []string  `json:"composite_tags"`
	NetsuiteInternalItemID string    `json:"netsuite_internal_item_id"`
	NetsuiteOverageItemID  string    `json:"netsuite_overage_item_id"`
	StartingAt             time.Time `json:"starting_at" format:"date-time"`
	Tags                   []string  `json:"tags"`
	JSON                   productListItemInitialJSON
}

// productListItemInitialJSON contains the JSON metadata for the struct
// [ProductListItemInitial]
type productListItemInitialJSON struct {
	CreatedAt              apijson.Field
	CreatedBy              apijson.Field
	IsRefundable           apijson.Field
	Name                   apijson.Field
	BillableMetricID       apijson.Field
	CompositeProductIDs    apijson.Field
	CompositeTags          apijson.Field
	NetsuiteInternalItemID apijson.Field
	NetsuiteOverageItemID  apijson.Field
	StartingAt             apijson.Field
	Tags                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProductListItemInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListItemType string

const (
	ProductListItemTypeUsage     ProductListItemType = "USAGE"
	ProductListItemTypeComposite ProductListItemType = "COMPOSITE"
	ProductListItemTypeFixed     ProductListItemType = "FIXED"
)

type ProductListItemUpdate struct {
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy              string    `json:"created_by,required"`
	BillableMetricID       string    `json:"billable_metric_id" format:"uuid"`
	CompositeProductIDs    []string  `json:"composite_product_ids" format:"uuid"`
	CompositeTags          []string  `json:"composite_tags"`
	IsRefundable           bool      `json:"is_refundable"`
	Name                   string    `json:"name"`
	NetsuiteInternalItemID string    `json:"netsuite_internal_item_id"`
	NetsuiteOverageItemID  string    `json:"netsuite_overage_item_id"`
	StartingAt             time.Time `json:"starting_at" format:"date-time"`
	Tags                   []string  `json:"tags"`
	JSON                   productListItemUpdateJSON
}

// productListItemUpdateJSON contains the JSON metadata for the struct
// [ProductListItemUpdate]
type productListItemUpdateJSON struct {
	CreatedAt              apijson.Field
	CreatedBy              apijson.Field
	BillableMetricID       apijson.Field
	CompositeProductIDs    apijson.Field
	CompositeTags          apijson.Field
	IsRefundable           apijson.Field
	Name                   apijson.Field
	NetsuiteInternalItemID apijson.Field
	NetsuiteOverageItemID  apijson.Field
	StartingAt             apijson.Field
	Tags                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ProductListItemUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductNewResponse struct {
	Data shared.ID `json:"data,required"`
	JSON contractPricingProductNewResponseJSON
}

// contractPricingProductNewResponseJSON contains the JSON metadata for the struct
// [ContractPricingProductNewResponse]
type contractPricingProductNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductGetResponse struct {
	Data ProductListItem `json:"data,required"`
	JSON contractPricingProductGetResponseJSON
}

// contractPricingProductGetResponseJSON contains the JSON metadata for the struct
// [ContractPricingProductGetResponse]
type contractPricingProductGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductUpdateResponse struct {
	Data shared.ID `json:"data,required"`
	JSON contractPricingProductUpdateResponseJSON
}

// contractPricingProductUpdateResponseJSON contains the JSON metadata for the
// struct [ContractPricingProductUpdateResponse]
type contractPricingProductUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductListResponse struct {
	Data     []ProductListItem `json:"data,required"`
	NextPage string            `json:"next_page,required,nullable"`
	JSON     contractPricingProductListResponseJSON
}

// contractPricingProductListResponseJSON contains the JSON metadata for the struct
// [ContractPricingProductListResponse]
type contractPricingProductListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductNewParams struct {
	// displayed on invoices
	Name param.Field[string]                              `json:"name,required"`
	Type param.Field[ContractPricingProductNewParamsType] `json:"type,required"`
	// Required for USAGE products
	BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
	// Required for COMPOSITE products
	CompositeProductIDs param.Field[[]string] `json:"composite_product_ids" format:"uuid"`
	// Required for COMPOSITE products
	CompositeTags param.Field[[]string] `json:"composite_tags"`
	// Defaults to true
	IsRefundable           param.Field[bool]   `json:"is_refundable"`
	NetsuiteInternalItemID param.Field[string] `json:"netsuite_internal_item_id"`
	NetsuiteOverageItemID  param.Field[string] `json:"netsuite_overage_item_id"`
	// Timestamp representing when the product will start existing. If not provided,
	// defaults to the current time and the product will be added to the product list
	// immediately.
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
	Tags       param.Field[[]string]  `json:"tags"`
}

func (r ContractPricingProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractPricingProductNewParamsType string

const (
	ContractPricingProductNewParamsTypeFixed     ContractPricingProductNewParamsType = "FIXED"
	ContractPricingProductNewParamsTypeFixed     ContractPricingProductNewParamsType = "fixed"
	ContractPricingProductNewParamsTypeUsage     ContractPricingProductNewParamsType = "USAGE"
	ContractPricingProductNewParamsTypeUsage     ContractPricingProductNewParamsType = "usage"
	ContractPricingProductNewParamsTypeComposite ContractPricingProductNewParamsType = "COMPOSITE"
	ContractPricingProductNewParamsTypeComposite ContractPricingProductNewParamsType = "composite"
)

type ContractPricingProductGetParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r ContractPricingProductGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractPricingProductUpdateParams struct {
	// ID of the product to update
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Timestamp representing when the update should go into effect. It must be on an
	// hour boundary (e.g. 1:00, not 1:30).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Available for USAGE products only. If not provided, defaults to product's
	// current billable metric.
	BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
	// Available for COMPOSITE products only. If not provided, defaults to product's
	// current composite_product_ids.
	CompositeProductIDs param.Field[[]string] `json:"composite_product_ids" format:"uuid"`
	// Available for COMPOSITE products only. If not provided, defaults to product's
	// current composite_tags.
	CompositeTags param.Field[[]string] `json:"composite_tags"`
	// Defaults to product's current refundability status
	IsRefundable param.Field[bool] `json:"is_refundable"`
	// displayed on invoices. If not provided, defaults to product's current name.
	Name param.Field[string] `json:"name"`
	// If not provided, defaults to product's current netsuite_internal_item_id.
	NetsuiteInternalItemID param.Field[string] `json:"netsuite_internal_item_id"`
	// Available for USAGE and COMPOSITE products only. If not provided, defaults to
	// product's current netsuite_overage_item_id.
	NetsuiteOverageItemID param.Field[string] `json:"netsuite_overage_item_id"`
	// If not provided, defaults to product's current tags
	Tags param.Field[[]string] `json:"tags"`
}

func (r ContractPricingProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractPricingProductListParams struct {
	Body param.Field[interface{}] `json:"body,required"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

func (r ContractPricingProductListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ContractPricingProductListParams]'s query parameters as
// `url.Values`.
func (r ContractPricingProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
