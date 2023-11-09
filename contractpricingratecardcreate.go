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

// ContractPricingRateCardCreateService contains methods and other services that
// help with interacting with the metronome API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewContractPricingRateCardCreateService] method instead.
type ContractPricingRateCardCreateService struct {
	Options []option.RequestOption
}

// NewContractPricingRateCardCreateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewContractPricingRateCardCreateService(opts ...option.RequestOption) (r *ContractPricingRateCardCreateService) {
	r = &ContractPricingRateCardCreateService{}
	r.Options = opts
	return
}

// Create a new rate card
func (r *ContractPricingRateCardCreateService) NewRateCard(ctx context.Context, body ContractPricingRateCardCreateNewRateCardParams, opts ...option.RequestOption) (res *ContractPricingRateCardCreateNewRateCardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contract-pricing/rate-cards/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractPricingRateCardCreateNewRateCardResponse struct {
	Data ContractPricingRateCardCreateNewRateCardResponseData `json:"data,required"`
	JSON contractPricingRateCardCreateNewRateCardResponseJSON
}

// contractPricingRateCardCreateNewRateCardResponseJSON contains the JSON metadata
// for the struct [ContractPricingRateCardCreateNewRateCardResponse]
type contractPricingRateCardCreateNewRateCardResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardCreateNewRateCardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardCreateNewRateCardResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractPricingRateCardCreateNewRateCardResponseDataJSON
}

// contractPricingRateCardCreateNewRateCardResponseDataJSON contains the JSON
// metadata for the struct [ContractPricingRateCardCreateNewRateCardResponseData]
type contractPricingRateCardCreateNewRateCardResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractPricingRateCardCreateNewRateCardResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractPricingRateCardCreateNewRateCardParams struct {
	// Used only in UI/API. It is not exposed to end customers.
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
}

func (r ContractPricingRateCardCreateNewRateCardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
