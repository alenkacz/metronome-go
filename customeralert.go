// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/internal/shared"
	"github.com/Metronome-Industries/metronome-go/option"
)

// CustomerAlertService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerAlertService] method
// instead.
type CustomerAlertService struct {
	Options []option.RequestOption
}

// NewCustomerAlertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerAlertService(opts ...option.RequestOption) (r *CustomerAlertService) {
	r = &CustomerAlertService{}
	r.Options = opts
	return
}

// Get the customer alert status and alert information for the specified customer
// and alert
func (r *CustomerAlertService) Get(ctx context.Context, body CustomerAlertGetParams, opts ...option.RequestOption) (res *CustomerAlertGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer-alerts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch all customer alert statuses and alert information for a customer
func (r *CustomerAlertService) List(ctx context.Context, params CustomerAlertListParams, opts ...option.RequestOption) (res *shared.Page[CustomerAlertListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "customer-alerts/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch all customer alert statuses and alert information for a customer
func (r *CustomerAlertService) ListAutoPaging(ctx context.Context, params CustomerAlertListParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerAlertListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, params, opts...))
}

type CustomerAlertGetResponse struct {
	Data CustomerAlertGetResponseData `json:"data,required"`
	JSON customerAlertGetResponseJSON `json:"-"`
}

// customerAlertGetResponseJSON contains the JSON metadata for the struct
// [CustomerAlertGetResponse]
type customerAlertGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertGetResponseData struct {
	Alert CustomerAlertGetResponseDataAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus CustomerAlertGetResponseDataCustomerStatus `json:"customer_status,required,nullable"`
	JSON           customerAlertGetResponseDataJSON           `json:"-"`
}

// customerAlertGetResponseDataJSON contains the JSON metadata for the struct
// [CustomerAlertGetResponseData]
type customerAlertGetResponseDataJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAlertGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertGetResponseDataAlert struct {
	// the Metronome ID of the alert
	ID         string                                      `json:"id,required"`
	CreditType CustomerAlertGetResponseDataAlertCreditType `json:"credit_type,required,nullable"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status CustomerAlertGetResponseDataAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type CustomerAlertGetResponseDataAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON      customerAlertGetResponseDataAlertJSON `json:"-"`
}

// customerAlertGetResponseDataAlertJSON contains the JSON metadata for the struct
// [CustomerAlertGetResponseDataAlert]
type customerAlertGetResponseDataAlertJSON struct {
	ID          apijson.Field
	CreditType  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Threshold   apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponseDataAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertGetResponseDataAlertCreditType struct {
	ID   string                                          `json:"id,required" format:"uuid"`
	Name string                                          `json:"name,required"`
	JSON customerAlertGetResponseDataAlertCreditTypeJSON `json:"-"`
}

// customerAlertGetResponseDataAlertCreditTypeJSON contains the JSON metadata for
// the struct [CustomerAlertGetResponseDataAlertCreditType]
type customerAlertGetResponseDataAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertGetResponseDataAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the alert
type CustomerAlertGetResponseDataAlertStatus string

const (
	CustomerAlertGetResponseDataAlertStatusEnabled  CustomerAlertGetResponseDataAlertStatus = "enabled"
	CustomerAlertGetResponseDataAlertStatusArchived CustomerAlertGetResponseDataAlertStatus = "archived"
	CustomerAlertGetResponseDataAlertStatusDisabled CustomerAlertGetResponseDataAlertStatus = "disabled"
)

// Type of the alert
type CustomerAlertGetResponseDataAlertType string

const (
	CustomerAlertGetResponseDataAlertTypeLowCreditBalanceReached                  CustomerAlertGetResponseDataAlertType = "low_credit_balance_reached"
	CustomerAlertGetResponseDataAlertTypeSpendThresholdReached                    CustomerAlertGetResponseDataAlertType = "spend_threshold_reached"
	CustomerAlertGetResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached CustomerAlertGetResponseDataAlertType = "monthly_invoice_total_spend_threshold_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingDaysInPlanReached            CustomerAlertGetResponseDataAlertType = "low_remaining_days_in_plan_reached"
	CustomerAlertGetResponseDataAlertTypeLowRemainingCreditPercentageReached      CustomerAlertGetResponseDataAlertType = "low_remaining_credit_percentage_reached"
	CustomerAlertGetResponseDataAlertTypeUsageThresholdReached                    CustomerAlertGetResponseDataAlertType = "usage_threshold_reached"
)

// The status of the customer alert. If the alert is archived, null will be
// returned.
type CustomerAlertGetResponseDataCustomerStatus string

const (
	CustomerAlertGetResponseDataCustomerStatusOk         CustomerAlertGetResponseDataCustomerStatus = "ok"
	CustomerAlertGetResponseDataCustomerStatusInAlarm    CustomerAlertGetResponseDataCustomerStatus = "in_alarm"
	CustomerAlertGetResponseDataCustomerStatusEvaluating CustomerAlertGetResponseDataCustomerStatus = "evaluating"
)

type CustomerAlertListResponse struct {
	Alert CustomerAlertListResponseAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus CustomerAlertListResponseCustomerStatus `json:"customer_status,required,nullable"`
	JSON           customerAlertListResponseJSON           `json:"-"`
}

// customerAlertListResponseJSON contains the JSON metadata for the struct
// [CustomerAlertListResponse]
type customerAlertListResponseJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAlertListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertListResponseAlert struct {
	// the Metronome ID of the alert
	ID         string                                   `json:"id,required"`
	CreditType CustomerAlertListResponseAlertCreditType `json:"credit_type,required,nullable"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status CustomerAlertListResponseAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type CustomerAlertListResponseAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	JSON      customerAlertListResponseAlertJSON `json:"-"`
}

// customerAlertListResponseAlertJSON contains the JSON metadata for the struct
// [CustomerAlertListResponseAlert]
type customerAlertListResponseAlertJSON struct {
	ID          apijson.Field
	CreditType  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Threshold   apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponseAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertListResponseAlertCreditType struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	Name string                                       `json:"name,required"`
	JSON customerAlertListResponseAlertCreditTypeJSON `json:"-"`
}

// customerAlertListResponseAlertCreditTypeJSON contains the JSON metadata for the
// struct [CustomerAlertListResponseAlertCreditType]
type customerAlertListResponseAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponseAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the alert
type CustomerAlertListResponseAlertStatus string

const (
	CustomerAlertListResponseAlertStatusEnabled  CustomerAlertListResponseAlertStatus = "enabled"
	CustomerAlertListResponseAlertStatusArchived CustomerAlertListResponseAlertStatus = "archived"
	CustomerAlertListResponseAlertStatusDisabled CustomerAlertListResponseAlertStatus = "disabled"
)

// Type of the alert
type CustomerAlertListResponseAlertType string

const (
	CustomerAlertListResponseAlertTypeLowCreditBalanceReached                  CustomerAlertListResponseAlertType = "low_credit_balance_reached"
	CustomerAlertListResponseAlertTypeSpendThresholdReached                    CustomerAlertListResponseAlertType = "spend_threshold_reached"
	CustomerAlertListResponseAlertTypeMonthlyInvoiceTotalSpendThresholdReached CustomerAlertListResponseAlertType = "monthly_invoice_total_spend_threshold_reached"
	CustomerAlertListResponseAlertTypeLowRemainingDaysInPlanReached            CustomerAlertListResponseAlertType = "low_remaining_days_in_plan_reached"
	CustomerAlertListResponseAlertTypeLowRemainingCreditPercentageReached      CustomerAlertListResponseAlertType = "low_remaining_credit_percentage_reached"
	CustomerAlertListResponseAlertTypeUsageThresholdReached                    CustomerAlertListResponseAlertType = "usage_threshold_reached"
)

// The status of the customer alert. If the alert is archived, null will be
// returned.
type CustomerAlertListResponseCustomerStatus string

const (
	CustomerAlertListResponseCustomerStatusOk         CustomerAlertListResponseCustomerStatus = "ok"
	CustomerAlertListResponseCustomerStatusInAlarm    CustomerAlertListResponseCustomerStatus = "in_alarm"
	CustomerAlertListResponseCustomerStatusEvaluating CustomerAlertListResponseCustomerStatus = "evaluating"
)

type CustomerAlertGetParams struct {
	// The Metronome ID of the alert
	AlertID param.Field[string] `json:"alert_id,required" format:"uuid"`
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
}

func (r CustomerAlertGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerAlertListParams struct {
	// The Metronome ID of the customer
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optionally filter by alert status. If absent, only enabled alerts will be
	// returned.
	AlertStatuses param.Field[[]CustomerAlertListParamsAlertStatus] `json:"alert_statuses"`
}

func (r CustomerAlertListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CustomerAlertListParams]'s query parameters as
// `url.Values`.
func (r CustomerAlertListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerAlertListParamsAlertStatus string

const (
	CustomerAlertListParamsAlertStatusEnabled  CustomerAlertListParamsAlertStatus = "enabled"
	CustomerAlertListParamsAlertStatusDisabled CustomerAlertListParamsAlertStatus = "disabled"
	CustomerAlertListParamsAlertStatusArchived CustomerAlertListParamsAlertStatus = "archived"
	CustomerAlertListParamsAlertStatusEnabled  CustomerAlertListParamsAlertStatus = "ENABLED"
	CustomerAlertListParamsAlertStatusDisabled CustomerAlertListParamsAlertStatus = "DISABLED"
	CustomerAlertListParamsAlertStatusArchived CustomerAlertListParamsAlertStatus = "ARCHIVED"
	CustomerAlertListParamsAlertStatusEnabled  CustomerAlertListParamsAlertStatus = "Enabled"
	CustomerAlertListParamsAlertStatusDisabled CustomerAlertListParamsAlertStatus = "Disabled"
	CustomerAlertListParamsAlertStatusArchived CustomerAlertListParamsAlertStatus = "Archived"
)
