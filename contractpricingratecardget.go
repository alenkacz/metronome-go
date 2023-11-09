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

// ContractPricingRateCardGetService contains methods and other services that help
// with interacting with the metronome API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardGetService] method instead.
type ContractPricingRateCardGetService struct {
	Options []option.RequestOption
}

// NewContractPricingRateCardGetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingRateCardGetService(opts ...option.RequestOption) (r *ContractPricingRateCardGetService) {
	r = &ContractPricingRateCardGetService{}
	r.Options = opts
	return
}

// Get a specific rate card
func (r *ContractPricingRateCardGetService) GetRateCard(ctx context.Context, body ContractPricingRateCardGetGetRateCardParams, opts ...option.RequestOption) (res *ContractPricingRateCardGetGetRateCardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractPricingRateCardGetGetRateCardResponse struct {
	Data ContractPricingRateCardGetGetRateCardResponseData `json:"data,required"`
	JSON contractPricingRateCardGetGetRateCardResponseJSON
}

// contractPricingRateCardGetGetRateCardResponseJSON contains the JSON metadata for
// the struct [ContractPricingRateCardGetGetRateCardResponse]
type contractPricingRateCardGetGetRateCardResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardGetGetRateCardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardGetGetRateCardResponseData struct {
	ID              string                                                                    `json:"id,required" format:"uuid"`
	CreatedAt       time.Time                                                                 `json:"created_at,required" format:"date-time"`
	CreatedBy       string                                                                    `json:"created_by,required"`
	Name            string                                                                    `json:"name,required"`
	RateCardEntries map[string]ContractPricingRateCardGetGetRateCardResponseDataRateCardEntry `json:"rate_card_entries,required"`
	Description     string                                                                    `json:"description"`
	JSON            contractPricingRateCardGetGetRateCardResponseDataJSON
}

// contractPricingRateCardGetGetRateCardResponseDataJSON contains the JSON metadata
// for the struct [ContractPricingRateCardGetGetRateCardResponseData]
type contractPricingRateCardGetGetRateCardResponseDataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	CreatedBy       apijson.Field
	Name            apijson.Field
	RateCardEntries apijson.Field
	Description     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractPricingRateCardGetGetRateCardResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardGetGetRateCardResponseDataRateCardEntry struct {
	Current ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrent  `json:"current,nullable"`
	Updates []ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdate `json:"updates"`
	JSON    contractPricingRateCardGetGetRateCardResponseDataRateCardEntryJSON
}

// contractPricingRateCardGetGetRateCardResponseDataRateCardEntryJSON contains the
// JSON metadata for the struct
// [ContractPricingRateCardGetGetRateCardResponseDataRateCardEntry]
type contractPricingRateCardGetGetRateCardResponseDataRateCardEntryJSON struct {
	Current     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardGetGetRateCardResponseDataRateCardEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrent struct {
	ID           string                                                                          `json:"id" format:"uuid"`
	CreatedAt    time.Time                                                                       `json:"created_at" format:"date-time"`
	CreatedBy    string                                                                          `json:"created_by"`
	EndingBefore time.Time                                                                       `json:"ending_before" format:"date-time"`
	Entitled     bool                                                                            `json:"entitled"`
	Price        float64                                                                         `json:"price"`
	ProductID    string                                                                          `json:"product_id" format:"uuid"`
	RateType     ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentRateType `json:"rate_type"`
	StartingAt   time.Time                                                                       `json:"starting_at" format:"date-time"`
	JSON         contractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentJSON
}

// contractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrent]
type contractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	CreatedBy    apijson.Field
	EndingBefore apijson.Field
	Entitled     apijson.Field
	Price        apijson.Field
	ProductID    apijson.Field
	RateType     apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentRateType string

const (
	ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentRateTypeFlat       ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentRateType = "FLAT"
	ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentRateTypePercentage ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesCurrentRateType = "PERCENTAGE"
)

type ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdate struct {
	ID           string                                                                          `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                                                       `json:"created_at,required" format:"date-time"`
	CreatedBy    string                                                                          `json:"created_by,required"`
	Entitled     bool                                                                            `json:"entitled,required"`
	Price        float64                                                                         `json:"price,required"`
	ProductID    string                                                                          `json:"product_id,required" format:"uuid"`
	RateType     ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdatesRateType `json:"rate_type,required"`
	StartingAt   time.Time                                                                       `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                                       `json:"ending_before" format:"date-time"`
	JSON         contractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdateJSON
}

// contractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdateJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdate]
type contractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdateJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	CreatedBy    apijson.Field
	Entitled     apijson.Field
	Price        apijson.Field
	ProductID    apijson.Field
	RateType     apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdatesRateType string

const (
	ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdatesRateTypeFlat       ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdatesRateType = "FLAT"
	ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdatesRateTypePercentage ContractPricingRateCardGetGetRateCardResponseDataRateCardEntriesUpdatesRateType = "PERCENTAGE"
)

type ContractPricingRateCardGetGetRateCardParams struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r ContractPricingRateCardGetGetRateCardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
