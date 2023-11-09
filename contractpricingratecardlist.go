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
	"github.com/metronome/metronome-go/option"
)

// ContractPricingRateCardListService contains methods and other services that help
// with interacting with the metronome API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardListService] method instead.
type ContractPricingRateCardListService struct {
	Options []option.RequestOption
}

// NewContractPricingRateCardListService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingRateCardListService(opts ...option.RequestOption) (r *ContractPricingRateCardListService) {
	r = &ContractPricingRateCardListService{}
	r.Options = opts
	return
}

// List rate cards
func (r *ContractPricingRateCardListService) ListRateCards(ctx context.Context, params ContractPricingRateCardListListRateCardsParams, opts ...option.RequestOption) (res *ContractPricingRateCardListListRateCardsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ContractPricingRateCardListListRateCardsResponse struct {
	Data     []ContractPricingRateCardListListRateCardsResponseData `json:"data,required"`
	NextPage string                                                 `json:"next_page,required,nullable"`
	JSON     contractPricingRateCardListListRateCardsResponseJSON
}

// contractPricingRateCardListListRateCardsResponseJSON contains the JSON metadata
// for the struct [ContractPricingRateCardListListRateCardsResponse]
type contractPricingRateCardListListRateCardsResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardListListRateCardsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardListListRateCardsResponseData struct {
	ID              string                                                                       `json:"id,required" format:"uuid"`
	CreatedAt       time.Time                                                                    `json:"created_at,required" format:"date-time"`
	CreatedBy       string                                                                       `json:"created_by,required"`
	Name            string                                                                       `json:"name,required"`
	RateCardEntries map[string]ContractPricingRateCardListListRateCardsResponseDataRateCardEntry `json:"rate_card_entries,required"`
	Description     string                                                                       `json:"description"`
	JSON            contractPricingRateCardListListRateCardsResponseDataJSON
}

// contractPricingRateCardListListRateCardsResponseDataJSON contains the JSON
// metadata for the struct [ContractPricingRateCardListListRateCardsResponseData]
type contractPricingRateCardListListRateCardsResponseDataJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	CreatedBy       apijson.Field
	Name            apijson.Field
	RateCardEntries apijson.Field
	Description     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractPricingRateCardListListRateCardsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardListListRateCardsResponseDataRateCardEntry struct {
	Current ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrent  `json:"current,nullable"`
	Updates []ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdate `json:"updates"`
	JSON    contractPricingRateCardListListRateCardsResponseDataRateCardEntryJSON
}

// contractPricingRateCardListListRateCardsResponseDataRateCardEntryJSON contains
// the JSON metadata for the struct
// [ContractPricingRateCardListListRateCardsResponseDataRateCardEntry]
type contractPricingRateCardListListRateCardsResponseDataRateCardEntryJSON struct {
	Current     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardListListRateCardsResponseDataRateCardEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrent struct {
	ID           string                                                                             `json:"id" format:"uuid"`
	CreatedAt    time.Time                                                                          `json:"created_at" format:"date-time"`
	CreatedBy    string                                                                             `json:"created_by"`
	EndingBefore time.Time                                                                          `json:"ending_before" format:"date-time"`
	Entitled     bool                                                                               `json:"entitled"`
	Price        float64                                                                            `json:"price"`
	ProductID    string                                                                             `json:"product_id" format:"uuid"`
	RateType     ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentRateType `json:"rate_type"`
	StartingAt   time.Time                                                                          `json:"starting_at" format:"date-time"`
	JSON         contractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentJSON
}

// contractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrent]
type contractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentJSON struct {
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

func (r *ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentRateType string

const (
	ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentRateTypeFlat       ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentRateType = "FLAT"
	ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentRateTypePercentage ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesCurrentRateType = "PERCENTAGE"
)

type ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdate struct {
	ID           string                                                                             `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                                                          `json:"created_at,required" format:"date-time"`
	CreatedBy    string                                                                             `json:"created_by,required"`
	Entitled     bool                                                                               `json:"entitled,required"`
	Price        float64                                                                            `json:"price,required"`
	ProductID    string                                                                             `json:"product_id,required" format:"uuid"`
	RateType     ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdatesRateType `json:"rate_type,required"`
	StartingAt   time.Time                                                                          `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                                          `json:"ending_before" format:"date-time"`
	JSON         contractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdateJSON
}

// contractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdateJSON
// contains the JSON metadata for the struct
// [ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdate]
type contractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdateJSON struct {
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

func (r *ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdatesRateType string

const (
	ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdatesRateTypeFlat       ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdatesRateType = "FLAT"
	ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdatesRateTypePercentage ContractPricingRateCardListListRateCardsResponseDataRateCardEntriesUpdatesRateType = "PERCENTAGE"
)

type ContractPricingRateCardListListRateCardsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

func (r ContractPricingRateCardListListRateCardsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ContractPricingRateCardListListRateCardsParams]'s query
// parameters as `url.Values`.
func (r ContractPricingRateCardListListRateCardsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
