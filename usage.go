// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "time"
  "github.com/metronome/metronome-go/internal/param"
  "net/url"
  "github.com/metronome/metronome-go/internal/apiquery"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// UsageService contains methods and other services that help with interacting with
// the metronome API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUsageService] method instead.
type UsageService struct {
Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
  r = &UsageService{}
  r.Options = opts
  return
}

// Fetch aggregated usage data for multiple customers and billable-metrics, broken
// into intervals of the specified length.
func (r *UsageService) GetUsage(ctx context.Context, params UsageGetUsageParams, opts ...option.RequestOption) (res *UsageGetUsageResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "usage"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
  return
}

// Fetch aggregated usage data for the specified customer, billable-metric, and
// optional group, broken into intervals of the specified length.
func (r *UsageService) GetUsageGroups(ctx context.Context, params UsageGetUsageGroupsParams, opts ...option.RequestOption) (res *UsageGetUsageGroupsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "usage/groups"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
  return
}

type UsageGetUsageResponse struct {
Data []UsageGetUsageResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON usageGetUsageResponseJSON
}

// usageGetUsageResponseJSON contains the JSON metadata for the struct
// [UsageGetUsageResponse]
type usageGetUsageResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UsageGetUsageResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type UsageGetUsageResponseData struct {
BillableMetricID string `json:"billable_metric_id,required" format:"uuid"`
BillableMetricName string `json:"billable_metric_name,required"`
CustomerID string `json:"customer_id,required" format:"uuid"`
EndTimestamp time.Time `json:"end_timestamp,required" format:"date-time"`
StartTimestamp time.Time `json:"start_timestamp,required" format:"date-time"`
Value float64 `json:"value,required,nullable"`
// Values will be either a number or null. Null indicates that there were no
// matches for the group_by value.
Groups map[string]float64 `json:"groups"`
JSON usageGetUsageResponseDataJSON
}

// usageGetUsageResponseDataJSON contains the JSON metadata for the struct
// [UsageGetUsageResponseData]
type usageGetUsageResponseDataJSON struct {
BillableMetricID apijson.Field
BillableMetricName apijson.Field
CustomerID apijson.Field
EndTimestamp apijson.Field
StartTimestamp apijson.Field
Value apijson.Field
Groups apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UsageGetUsageResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type UsageGetUsageGroupsResponse struct {
Data []UsageGetUsageGroupsResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON usageGetUsageGroupsResponseJSON
}

// usageGetUsageGroupsResponseJSON contains the JSON metadata for the struct
// [UsageGetUsageGroupsResponse]
type usageGetUsageGroupsResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UsageGetUsageGroupsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type UsageGetUsageGroupsResponseData struct {
EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
GroupKey string `json:"group_key,required,nullable"`
GroupValue string `json:"group_value,required,nullable"`
StartingOn time.Time `json:"starting_on,required" format:"date-time"`
Value float64 `json:"value,required,nullable"`
JSON usageGetUsageGroupsResponseDataJSON
}

// usageGetUsageGroupsResponseDataJSON contains the JSON metadata for the struct
// [UsageGetUsageGroupsResponseData]
type usageGetUsageGroupsResponseDataJSON struct {
EndingBefore apijson.Field
GroupKey apijson.Field
GroupValue apijson.Field
StartingOn apijson.Field
Value apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *UsageGetUsageGroupsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type UsageGetUsageParams struct {
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
StartingOn param.Field[time.Time] `json:"starting_on,required" format:"date-time"`
// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
WindowSize param.Field[UsageGetUsageParamsWindowSize] `json:"window_size,required"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
// A list of billable metrics to fetch usage for. If absent, all billable metrics
// will be returned.
BillableMetrics param.Field[[]UsageGetUsageParamsBillableMetric] `json:"billable_metrics"`
// A list of Metronome customer IDs to fetch usage for. If absent, usage for all
// customers will be returned.
CustomerIDs param.Field[[]string] `json:"customer_ids" format:"uuid"`
}

func (r UsageGetUsageParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// URLQuery serializes [UsageGetUsageParams]'s query parameters as `url.Values`.
func (r UsageGetUsageParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
type UsageGetUsageParamsWindowSize string

const (
  UsageGetUsageParamsWindowSizeHour UsageGetUsageParamsWindowSize = "hour"
  UsageGetUsageParamsWindowSizeDay UsageGetUsageParamsWindowSize = "day"
  UsageGetUsageParamsWindowSizeNone UsageGetUsageParamsWindowSize = "none"
  UsageGetUsageParamsWindowSizeHour UsageGetUsageParamsWindowSize = "HOUR"
  UsageGetUsageParamsWindowSizeDay UsageGetUsageParamsWindowSize = "DAY"
  UsageGetUsageParamsWindowSizeNone UsageGetUsageParamsWindowSize = "NONE"
  UsageGetUsageParamsWindowSizeHour UsageGetUsageParamsWindowSize = "Hour"
  UsageGetUsageParamsWindowSizeDay UsageGetUsageParamsWindowSize = "Day"
  UsageGetUsageParamsWindowSizeNone UsageGetUsageParamsWindowSize = "None"
)

type UsageGetUsageParamsBillableMetric struct {
ID param.Field[string] `json:"id,required" format:"uuid"`
GroupBy param.Field[UsageGetUsageParamsBillableMetricsGroupBy] `json:"group_by"`
}

func (r UsageGetUsageParamsBillableMetric) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type UsageGetUsageParamsBillableMetricsGroupBy struct {
// The name of the group_by key to use
Key param.Field[string] `json:"key,required"`
// Values of the group_by key to return in the query. If this field is omitted, all
// available values will be returned, up to a maximum of 200.
Values param.Field[[]string] `json:"values"`
}

func (r UsageGetUsageParamsBillableMetricsGroupBy) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type UsageGetUsageGroupsParams struct {
BillableMetricID param.Field[string] `json:"billable_metric_id,required" format:"uuid"`
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
WindowSize param.Field[UsageGetUsageGroupsParamsWindowSize] `json:"window_size,required"`
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
// If true, will return the usage for the current billing period. Will return an
// error if the customer is currently uncontracted or starting_on and ending_before
// are specified when this is true.
CurrentPeriod param.Field[bool] `json:"current_period"`
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
GroupBy param.Field[UsageGetUsageGroupsParamsGroupBy] `json:"group_by"`
StartingOn param.Field[time.Time] `json:"starting_on" format:"date-time"`
}

func (r UsageGetUsageGroupsParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// URLQuery serializes [UsageGetUsageGroupsParams]'s query parameters as
// `url.Values`.
func (r UsageGetUsageGroupsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

// A window_size of "day" or "hour" will return the usage for the specified period
// segmented into daily or hourly aggregates. A window_size of "none" will return a
// single usage aggregate for the entirety of the specified period.
type UsageGetUsageGroupsParamsWindowSize string

const (
  UsageGetUsageGroupsParamsWindowSizeHour UsageGetUsageGroupsParamsWindowSize = "hour"
  UsageGetUsageGroupsParamsWindowSizeDay UsageGetUsageGroupsParamsWindowSize = "day"
  UsageGetUsageGroupsParamsWindowSizeNone UsageGetUsageGroupsParamsWindowSize = "none"
  UsageGetUsageGroupsParamsWindowSizeHour UsageGetUsageGroupsParamsWindowSize = "HOUR"
  UsageGetUsageGroupsParamsWindowSizeDay UsageGetUsageGroupsParamsWindowSize = "DAY"
  UsageGetUsageGroupsParamsWindowSizeNone UsageGetUsageGroupsParamsWindowSize = "NONE"
  UsageGetUsageGroupsParamsWindowSizeHour UsageGetUsageGroupsParamsWindowSize = "Hour"
  UsageGetUsageGroupsParamsWindowSizeDay UsageGetUsageGroupsParamsWindowSize = "Day"
  UsageGetUsageGroupsParamsWindowSizeNone UsageGetUsageGroupsParamsWindowSize = "None"
)

type UsageGetUsageGroupsParamsGroupBy struct {
// The name of the group_by key to use
Key param.Field[string] `json:"key,required"`
// Values of the group_by key to return in the query. Omit this if you'd like all
// values for the key returned.
Values param.Field[[]string] `json:"values"`
}

func (r UsageGetUsageGroupsParamsGroupBy) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
