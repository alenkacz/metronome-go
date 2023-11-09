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

// ContractUpdateEndDateService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractUpdateEndDateService]
// method instead.
type ContractUpdateEndDateService struct {
	Options []option.RequestOption
}

// NewContractUpdateEndDateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractUpdateEndDateService(opts ...option.RequestOption) (r *ContractUpdateEndDateService) {
	r = &ContractUpdateEndDateService{}
	r.Options = opts
	return
}

// Update the end date of a contract
func (r *ContractUpdateEndDateService) UpdateContractEndDate(ctx context.Context, body ContractUpdateEndDateUpdateContractEndDateParams, opts ...option.RequestOption) (res *ContractUpdateEndDateUpdateContractEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractUpdateEndDateUpdateContractEndDateResponse struct {
	Data ContractUpdateEndDateUpdateContractEndDateResponseData `json:"data,required"`
	JSON contractUpdateEndDateUpdateContractEndDateResponseJSON
}

// contractUpdateEndDateUpdateContractEndDateResponseJSON contains the JSON
// metadata for the struct [ContractUpdateEndDateUpdateContractEndDateResponse]
type contractUpdateEndDateUpdateContractEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractUpdateEndDateUpdateContractEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractUpdateEndDateUpdateContractEndDateResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractUpdateEndDateUpdateContractEndDateResponseDataJSON
}

// contractUpdateEndDateUpdateContractEndDateResponseDataJSON contains the JSON
// metadata for the struct [ContractUpdateEndDateUpdateContractEndDateResponseData]
type contractUpdateEndDateUpdateContractEndDateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractUpdateEndDateUpdateContractEndDateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractUpdateEndDateUpdateContractEndDateParams struct {
	// ID of the contract to update
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive). If not
	// provided, the contract will be updated to be open-ended.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
}

func (r ContractUpdateEndDateUpdateContractEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
