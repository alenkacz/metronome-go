// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"time"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
)

// ContractPricingProductGetService contains methods and other services that help
// with interacting with the metronome API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingProductGetService] method instead.
type ContractPricingProductGetService struct {
	Options []option.RequestOption
}

// NewContractPricingProductGetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingProductGetService(opts ...option.RequestOption) (r *ContractPricingProductGetService) {
	r = &ContractPricingProductGetService{}
	r.Options = opts
	return
}

// Get a specific product
func (r *ContractPricingProductGetService) GetProduct(ctx context.Context, body ContractPricingProductGetGetProductParams, opts ...option.RequestOption) (res *ContractPricingProductGetGetProductResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractPricingProductGetGetProductResponse struct {
	Data ContractPricingProductGetGetProductResponseData `json:"data,required"`
	JSON contractPricingProductGetGetProductResponseJSON
}

// contractPricingProductGetGetProductResponseJSON contains the JSON metadata for
// the struct [ContractPricingProductGetGetProductResponse]
type contractPricingProductGetGetProductResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductGetGetProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductGetGetProductResponseData struct {
	ID      string                                                  `json:"id,required" format:"uuid"`
	Current ContractPricingProductGetGetProductResponseDataCurrent  `json:"current,required"`
	Initial ContractPricingProductGetGetProductResponseDataInitial  `json:"initial,required"`
	Type    ContractPricingProductGetGetProductResponseDataType     `json:"type,required"`
	Updates []ContractPricingProductGetGetProductResponseDataUpdate `json:"updates,required"`
	JSON    contractPricingProductGetGetProductResponseDataJSON
}

// contractPricingProductGetGetProductResponseDataJSON contains the JSON metadata
// for the struct [ContractPricingProductGetGetProductResponseData]
type contractPricingProductGetGetProductResponseDataJSON struct {
	ID          apijson.Field
	Current     apijson.Field
	Initial     apijson.Field
	Type        apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductGetGetProductResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductGetGetProductResponseDataCurrent struct {
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
	JSON                   contractPricingProductGetGetProductResponseDataCurrentJSON
}

// contractPricingProductGetGetProductResponseDataCurrentJSON contains the JSON
// metadata for the struct [ContractPricingProductGetGetProductResponseDataCurrent]
type contractPricingProductGetGetProductResponseDataCurrentJSON struct {
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

func (r *ContractPricingProductGetGetProductResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductGetGetProductResponseDataInitial struct {
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
	JSON                   contractPricingProductGetGetProductResponseDataInitialJSON
}

// contractPricingProductGetGetProductResponseDataInitialJSON contains the JSON
// metadata for the struct [ContractPricingProductGetGetProductResponseDataInitial]
type contractPricingProductGetGetProductResponseDataInitialJSON struct {
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

func (r *ContractPricingProductGetGetProductResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductGetGetProductResponseDataType string

const (
	ContractPricingProductGetGetProductResponseDataTypeUsage     ContractPricingProductGetGetProductResponseDataType = "USAGE"
	ContractPricingProductGetGetProductResponseDataTypeComposite ContractPricingProductGetGetProductResponseDataType = "COMPOSITE"
	ContractPricingProductGetGetProductResponseDataTypeFixed     ContractPricingProductGetGetProductResponseDataType = "FIXED"
)

type ContractPricingProductGetGetProductResponseDataUpdate struct {
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
	JSON                   contractPricingProductGetGetProductResponseDataUpdateJSON
}

// contractPricingProductGetGetProductResponseDataUpdateJSON contains the JSON
// metadata for the struct [ContractPricingProductGetGetProductResponseDataUpdate]
type contractPricingProductGetGetProductResponseDataUpdateJSON struct {
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

func (r *ContractPricingProductGetGetProductResponseDataUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductGetGetProductParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r ContractPricingProductGetGetProductParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
