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

// ContractSetUsageFilterService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractSetUsageFilterService]
// method instead.
type ContractSetUsageFilterService struct {
	Options []option.RequestOption
}

// NewContractSetUsageFilterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContractSetUsageFilterService(opts ...option.RequestOption) (r *ContractSetUsageFilterService) {
	r = &ContractSetUsageFilterService{}
	r.Options = opts
	return
}

// Set usage filter for a contract
func (r *ContractSetUsageFilterService) SetUsageFilter(ctx context.Context, body ContractSetUsageFilterSetUsageFilterParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "contracts/setUsageFilter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ContractSetUsageFilterSetUsageFilterParams struct {
	ContractID  param.Field[string]    `json:"contract_id,required" format:"uuid"`
	CustomerID  param.Field[string]    `json:"customer_id,required" format:"uuid"`
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractSetUsageFilterSetUsageFilterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
