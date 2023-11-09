// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
)

// ContractPricingRateCardUpdateService contains methods and other services that
// help with interacting with the metronome API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardUpdateService] method instead.
type ContractPricingRateCardUpdateService struct {
	Options []option.RequestOption
}

// NewContractPricingRateCardUpdateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingRateCardUpdateService(opts ...option.RequestOption) (r *ContractPricingRateCardUpdateService) {
	r = &ContractPricingRateCardUpdateService{}
	r.Options = opts
	return
}

// Update a rate card
func (r *ContractPricingRateCardUpdateService) UpdateRateCard(ctx context.Context, body ContractPricingRateCardUpdateUpdateRateCardParams, opts ...option.RequestOption) (res *ContractPricingRateCardUpdateUpdateRateCardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractPricingRateCardUpdateUpdateRateCardResponse struct {
	Data ContractPricingRateCardUpdateUpdateRateCardResponseData `json:"data,required"`
	JSON contractPricingRateCardUpdateUpdateRateCardResponseJSON
}

// contractPricingRateCardUpdateUpdateRateCardResponseJSON contains the JSON
// metadata for the struct [ContractPricingRateCardUpdateUpdateRateCardResponse]
type contractPricingRateCardUpdateUpdateRateCardResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardUpdateUpdateRateCardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardUpdateUpdateRateCardResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractPricingRateCardUpdateUpdateRateCardResponseDataJSON
}

// contractPricingRateCardUpdateUpdateRateCardResponseDataJSON contains the JSON
// metadata for the struct
// [ContractPricingRateCardUpdateUpdateRateCardResponseData]
type contractPricingRateCardUpdateUpdateRateCardResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardUpdateUpdateRateCardResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardUpdateUpdateRateCardParams struct {
	// ID of the rate card to update
	RateCardID  param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	Description param.Field[string] `json:"description"`
	// Used only in UI/API. It is not exposed to end customers.
	Name param.Field[string] `json:"name"`
}

func (r ContractPricingRateCardUpdateUpdateRateCardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
