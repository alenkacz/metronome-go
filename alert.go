// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/shared"
  "github.com/metronome/metronome-go/internal/apijson"
  "github.com/metronome/metronome-go/internal/param"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// AlertService contains methods and other services that help with interacting with
// the metronome API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAlertService] method instead.
type AlertService struct {
Options []option.RequestOption
}

// NewAlertService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlertService(opts ...option.RequestOption) (r *AlertService) {
  r = &AlertService{}
  r.Options = opts
  return
}

// Create an alert
func (r *AlertService) New(ctx context.Context, body AlertNewParams, opts ...option.RequestOption) (res *AlertNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "alerts/create"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Archive an alert
func (r *AlertService) Archive(ctx context.Context, body AlertArchiveParams, opts ...option.RequestOption) (res *AlertArchiveResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "alerts/archive"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type AlertNewResponse struct {
Data shared.ID `json:"data,required"`
JSON alertNewResponseJSON
}

// alertNewResponseJSON contains the JSON metadata for the struct
// [AlertNewResponse]
type alertNewResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *AlertNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type AlertArchiveResponse struct {
Data shared.ID `json:"data,required"`
JSON alertArchiveResponseJSON
}

// alertArchiveResponseJSON contains the JSON metadata for the struct
// [AlertArchiveResponse]
type alertArchiveResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *AlertArchiveResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type AlertNewParams struct {
// Type of the alert
AlertType param.Field[AlertNewParamsAlertType] `json:"alert_type,required"`
// Name of the alert
Name param.Field[string] `json:"name,required"`
// Threshold value of the alert policy
Threshold param.Field[float64] `json:"threshold,required"`
// For alerts of type `usage_threshold_reached`, specifies which billable metric to
// track the usage for.
BillableMetricID param.Field[string] `json:"billable_metric_id" format:"uuid"`
CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
// If provided, will create this alert for this specific customer. To create an
// alert for all customers, do not specify `customer_id` or `plan_id`.
CustomerID param.Field[string] `json:"customer_id" format:"uuid"`
// If provided, will create this alert for this specific plan. To create an alert
// for all customers, do not specify `customer_id` or `plan_id`.
PlanID param.Field[string] `json:"plan_id" format:"uuid"`
}

func (r AlertNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Type of the alert
type AlertNewParamsAlertType string

const (
  AlertNewParamsAlertTypeLowCreditBalanceReached AlertNewParamsAlertType = "low_credit_balance_reached"
  AlertNewParamsAlertTypeSpendThresholdReached AlertNewParamsAlertType = "spend_threshold_reached"
  AlertNewParamsAlertTypeMonthlyInvoiceTotalSpendThresholdReached AlertNewParamsAlertType = "monthly_invoice_total_spend_threshold_reached"
  AlertNewParamsAlertTypeLowRemainingDaysInPlanReached AlertNewParamsAlertType = "low_remaining_days_in_plan_reached"
  AlertNewParamsAlertTypeLowRemainingCreditPercentageReached AlertNewParamsAlertType = "low_remaining_credit_percentage_reached"
  AlertNewParamsAlertTypeUsageThresholdReached AlertNewParamsAlertType = "usage_threshold_reached"
)

type AlertArchiveParams struct {
ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r AlertArchiveParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
