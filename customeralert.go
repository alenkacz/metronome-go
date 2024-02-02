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
func (r *CustomerAlertService) List(ctx context.Context, params CustomerAlertListParams, opts ...option.RequestOption) (res *CustomerAlertListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer-alerts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
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
	Data     []CustomerAlertListResponseData `json:"data,required"`
	NextPage string                          `json:"next_page,required,nullable"`
	JSON     customerAlertListResponseJSON   `json:"-"`
}

// customerAlertListResponseJSON contains the JSON metadata for the struct
// [CustomerAlertListResponse]
type customerAlertListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertListResponseData struct {
	Alert CustomerAlertListResponseDataAlert `json:"alert,required"`
	// The status of the customer alert. If the alert is archived, null will be
	// returned.
	CustomerStatus CustomerAlertListResponseDataCustomerStatus `json:"customer_status,required,nullable"`
	JSON           customerAlertListResponseDataJSON           `json:"-"`
}

// customerAlertListResponseDataJSON contains the JSON metadata for the struct
// [CustomerAlertListResponseData]
type customerAlertListResponseDataJSON struct {
	Alert          apijson.Field
	CustomerStatus apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAlertListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertListResponseDataAlert struct {
	// the Metronome ID of the alert
	ID         string                                       `json:"id,required"`
	CreditType CustomerAlertListResponseDataAlertCreditType `json:"credit_type,required,nullable"`
	// Name of the alert
	Name string `json:"name,required"`
	// Status of the alert
	Status CustomerAlertListResponseDataAlertStatus `json:"status,required"`
	// Threshold value of the alert policy
	Threshold float64 `json:"threshold,required"`
	// Type of the alert
	Type CustomerAlertListResponseDataAlertType `json:"type,required"`
	// Timestamp for when the alert was last updated
	UpdatedAt time.Time                              `json:"updated_at,required" format:"date-time"`
	JSON      customerAlertListResponseDataAlertJSON `json:"-"`
}

// customerAlertListResponseDataAlertJSON contains the JSON metadata for the struct
// [CustomerAlertListResponseDataAlert]
type customerAlertListResponseDataAlertJSON struct {
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

func (r *CustomerAlertListResponseDataAlert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerAlertListResponseDataAlertCreditType struct {
	ID   string                                           `json:"id,required" format:"uuid"`
	Name string                                           `json:"name,required"`
	JSON customerAlertListResponseDataAlertCreditTypeJSON `json:"-"`
}

// customerAlertListResponseDataAlertCreditTypeJSON contains the JSON metadata for
// the struct [CustomerAlertListResponseDataAlertCreditType]
type customerAlertListResponseDataAlertCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerAlertListResponseDataAlertCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the alert
type CustomerAlertListResponseDataAlertStatus string

const (
	CustomerAlertListResponseDataAlertStatusEnabled  CustomerAlertListResponseDataAlertStatus = "enabled"
	CustomerAlertListResponseDataAlertStatusArchived CustomerAlertListResponseDataAlertStatus = "archived"
	CustomerAlertListResponseDataAlertStatusDisabled CustomerAlertListResponseDataAlertStatus = "disabled"
)

// Type of the alert
type CustomerAlertListResponseDataAlertType string

const (
	CustomerAlertListResponseDataAlertTypeLowCreditBalanceReached                  CustomerAlertListResponseDataAlertType = "low_credit_balance_reached"
	CustomerAlertListResponseDataAlertTypeSpendThresholdReached                    CustomerAlertListResponseDataAlertType = "spend_threshold_reached"
	CustomerAlertListResponseDataAlertTypeMonthlyInvoiceTotalSpendThresholdReached CustomerAlertListResponseDataAlertType = "monthly_invoice_total_spend_threshold_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingDaysInPlanReached            CustomerAlertListResponseDataAlertType = "low_remaining_days_in_plan_reached"
	CustomerAlertListResponseDataAlertTypeLowRemainingCreditPercentageReached      CustomerAlertListResponseDataAlertType = "low_remaining_credit_percentage_reached"
	CustomerAlertListResponseDataAlertTypeUsageThresholdReached                    CustomerAlertListResponseDataAlertType = "usage_threshold_reached"
)

// The status of the customer alert. If the alert is archived, null will be
// returned.
type CustomerAlertListResponseDataCustomerStatus string

const (
	CustomerAlertListResponseDataCustomerStatusOk         CustomerAlertListResponseDataCustomerStatus = "ok"
	CustomerAlertListResponseDataCustomerStatusInAlarm    CustomerAlertListResponseDataCustomerStatus = "in_alarm"
	CustomerAlertListResponseDataCustomerStatusEvaluating CustomerAlertListResponseDataCustomerStatus = "evaluating"
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
