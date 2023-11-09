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

// ContractPricingProductUpdateService contains methods and other services that
// help with interacting with the metronome API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingProductUpdateService] method instead.
type ContractPricingProductUpdateService struct {
	Options []option.RequestOption
}

// NewContractPricingProductUpdateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingProductUpdateService(opts ...option.RequestOption) (r *ContractPricingProductUpdateService) {
	r = &ContractPricingProductUpdateService{}
	r.Options = opts
	return
}

// Update a product
func (r *ContractPricingProductUpdateService) UpdateProduct(ctx context.Context, body ContractPricingProductUpdateUpdateProductParams, opts ...option.RequestOption) (res *ContractPricingProductUpdateUpdateProductResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/products/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractPricingProductUpdateUpdateProductResponse struct {
	Data ContractPricingProductUpdateUpdateProductResponseData `json:"data,required"`
	JSON contractPricingProductUpdateUpdateProductResponseJSON
}

// contractPricingProductUpdateUpdateProductResponseJSON contains the JSON metadata
// for the struct [ContractPricingProductUpdateUpdateProductResponse]
type contractPricingProductUpdateUpdateProductResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductUpdateUpdateProductResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductUpdateUpdateProductResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractPricingProductUpdateUpdateProductResponseDataJSON
}

// contractPricingProductUpdateUpdateProductResponseDataJSON contains the JSON
// metadata for the struct [ContractPricingProductUpdateUpdateProductResponseData]
type contractPricingProductUpdateUpdateProductResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingProductUpdateUpdateProductResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingProductUpdateUpdateProductParams struct {
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

func (r ContractPricingProductUpdateUpdateProductParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
