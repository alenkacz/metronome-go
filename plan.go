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
  "fmt"
)

// PlanService contains methods and other services that help with interacting with
// the metronome API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPlanService] method instead.
type PlanService struct {
Options []option.RequestOption
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
  r = &PlanService{}
  r.Options = opts
  return
}

// List all available plans.
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *PlanListResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "plans"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

// Fetch high level details of a specific plan.
func (r *PlanService) GetDetails(ctx context.Context, planID string, opts ...option.RequestOption) (res *PlanGetDetailsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("planDetails/%s", planID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
  return
}

// Fetches a list of charges of a specific plan.
func (r *PlanService) ListCharges(ctx context.Context, planID string, query PlanListChargesParams, opts ...option.RequestOption) (res *PlanListChargesResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("planDetails/%s/charges", planID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

// Fetches a list of customers on a specific plan (only currently active plans are
// included)
func (r *PlanService) ListCustomers(ctx context.Context, planID string, query PlanListCustomersParams, opts ...option.RequestOption) (res *PlanListCustomersResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("planDetails/%s/customers", planID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

type PlanDetail struct {
ID string `json:"id,required" format:"uuid"`
CustomFields map[string]string `json:"custom_fields,required"`
Name string `json:"name,required"`
CreditGrants []PlanDetailCreditGrant `json:"credit_grants"`
Description string `json:"description"`
Minimums []PlanDetailMinimum `json:"minimums"`
JSON planDetailJSON
}

// planDetailJSON contains the JSON metadata for the struct [PlanDetail]
type planDetailJSON struct {
ID apijson.Field
CustomFields apijson.Field
Name apijson.Field
CreditGrants apijson.Field
Description apijson.Field
Minimums apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanDetail) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanDetailCreditGrant struct {
AmountGranted float64 `json:"amount_granted,required"`
AmountGrantedCreditType PlanDetailCreditGrantsAmountGrantedCreditType `json:"amount_granted_credit_type,required"`
AmountPaid float64 `json:"amount_paid,required"`
AmountPaidCreditType PlanDetailCreditGrantsAmountPaidCreditType `json:"amount_paid_credit_type,required"`
EffectiveDuration float64 `json:"effective_duration,required"`
Name string `json:"name,required"`
Priority string `json:"priority,required"`
SendInvoice bool `json:"send_invoice,required"`
Reason string `json:"reason"`
RecurrenceDuration float64 `json:"recurrence_duration"`
RecurrenceInterval float64 `json:"recurrence_interval"`
JSON planDetailCreditGrantJSON
}

// planDetailCreditGrantJSON contains the JSON metadata for the struct
// [PlanDetailCreditGrant]
type planDetailCreditGrantJSON struct {
AmountGranted apijson.Field
AmountGrantedCreditType apijson.Field
AmountPaid apijson.Field
AmountPaidCreditType apijson.Field
EffectiveDuration apijson.Field
Name apijson.Field
Priority apijson.Field
SendInvoice apijson.Field
Reason apijson.Field
RecurrenceDuration apijson.Field
RecurrenceInterval apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanDetailCreditGrant) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanDetailCreditGrantsAmountGrantedCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON planDetailCreditGrantsAmountGrantedCreditTypeJSON
}

// planDetailCreditGrantsAmountGrantedCreditTypeJSON contains the JSON metadata for
// the struct [PlanDetailCreditGrantsAmountGrantedCreditType]
type planDetailCreditGrantsAmountGrantedCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanDetailCreditGrantsAmountGrantedCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanDetailCreditGrantsAmountPaidCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON planDetailCreditGrantsAmountPaidCreditTypeJSON
}

// planDetailCreditGrantsAmountPaidCreditTypeJSON contains the JSON metadata for
// the struct [PlanDetailCreditGrantsAmountPaidCreditType]
type planDetailCreditGrantsAmountPaidCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanDetailCreditGrantsAmountPaidCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanDetailMinimum struct {
CreditType PlanDetailMinimumsCreditType `json:"credit_type,required"`
Name string `json:"name,required"`
// Used in price ramps. Indicates how many billing periods pass before the charge
// applies.
StartPeriod float64 `json:"start_period,required"`
Value float64 `json:"value,required"`
JSON planDetailMinimumJSON
}

// planDetailMinimumJSON contains the JSON metadata for the struct
// [PlanDetailMinimum]
type planDetailMinimumJSON struct {
CreditType apijson.Field
Name apijson.Field
StartPeriod apijson.Field
Value apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanDetailMinimum) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanDetailMinimumsCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON planDetailMinimumsCreditTypeJSON
}

// planDetailMinimumsCreditTypeJSON contains the JSON metadata for the struct
// [PlanDetailMinimumsCreditType]
type planDetailMinimumsCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanDetailMinimumsCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListResponse struct {
Data []PlanListResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON planListResponseJSON
}

// planListResponseJSON contains the JSON metadata for the struct
// [PlanListResponse]
type planListResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListResponseData struct {
ID string `json:"id,required" format:"uuid"`
Description string `json:"description,required"`
Name string `json:"name,required"`
JSON planListResponseDataJSON
}

// planListResponseDataJSON contains the JSON metadata for the struct
// [PlanListResponseData]
type planListResponseDataJSON struct {
ID apijson.Field
Description apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponse struct {
Data PlanDetail `json:"data,required"`
JSON planGetDetailsResponseJSON
}

// planGetDetailsResponseJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponse]
type planGetDetailsResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponse struct {
Data []PlanListChargesResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON planListChargesResponseJSON
}

// planListChargesResponseJSON contains the JSON metadata for the struct
// [PlanListChargesResponse]
type planListChargesResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponseData struct {
ID string `json:"id,required" format:"uuid"`
ChargeType PlanListChargesResponseDataChargeType `json:"charge_type,required"`
CreditType PlanListChargesResponseDataCreditType `json:"credit_type,required"`
CustomFields map[string]string `json:"custom_fields,required"`
Name string `json:"name,required"`
Prices []PlanListChargesResponseDataPrice `json:"prices,required"`
ProductName string `json:"product_name,required"`
Quantity float64 `json:"quantity"`
// Used in price ramps. Indicates how many billing periods pass before the charge
// applies.
StartPeriod float64 `json:"start_period"`
// Specifies how quantities for usage based charges will be converted.
UnitConversion PlanListChargesResponseDataUnitConversion `json:"unit_conversion"`
JSON planListChargesResponseDataJSON
}

// planListChargesResponseDataJSON contains the JSON metadata for the struct
// [PlanListChargesResponseData]
type planListChargesResponseDataJSON struct {
ID apijson.Field
ChargeType apijson.Field
CreditType apijson.Field
CustomFields apijson.Field
Name apijson.Field
Prices apijson.Field
ProductName apijson.Field
Quantity apijson.Field
StartPeriod apijson.Field
UnitConversion apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponseDataChargeType string

const (
  PlanListChargesResponseDataChargeTypeUsage PlanListChargesResponseDataChargeType = "usage"
  PlanListChargesResponseDataChargeTypeFixed PlanListChargesResponseDataChargeType = "fixed"
  PlanListChargesResponseDataChargeTypeComposite PlanListChargesResponseDataChargeType = "composite"
  PlanListChargesResponseDataChargeTypeMinimum PlanListChargesResponseDataChargeType = "minimum"
  PlanListChargesResponseDataChargeTypeSeat PlanListChargesResponseDataChargeType = "seat"
)

type PlanListChargesResponseDataCreditType struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON planListChargesResponseDataCreditTypeJSON
}

// planListChargesResponseDataCreditTypeJSON contains the JSON metadata for the
// struct [PlanListChargesResponseDataCreditType]
type planListChargesResponseDataCreditTypeJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponseDataPrice struct {
// Used in pricing tiers. Indicates at what metric value the price applies.
Tier float64 `json:"tier,required"`
Value float64 `json:"value,required"`
CollectionInterval float64 `json:"collection_interval"`
CollectionSchedule string `json:"collection_schedule"`
Quantity float64 `json:"quantity"`
JSON planListChargesResponseDataPriceJSON
}

// planListChargesResponseDataPriceJSON contains the JSON metadata for the struct
// [PlanListChargesResponseDataPrice]
type planListChargesResponseDataPriceJSON struct {
Tier apijson.Field
Value apijson.Field
CollectionInterval apijson.Field
CollectionSchedule apijson.Field
Quantity apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponseDataPrice) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// Specifies how quantities for usage based charges will be converted.
type PlanListChargesResponseDataUnitConversion struct {
// The conversion factor
DivisionFactor float64 `json:"division_factor,required"`
// Whether usage should be rounded down or up to the nearest whole number. If null,
// quantity will be rounded to 20 decimal places.
RoundingBehavior PlanListChargesResponseDataUnitConversionRoundingBehavior `json:"rounding_behavior"`
JSON planListChargesResponseDataUnitConversionJSON
}

// planListChargesResponseDataUnitConversionJSON contains the JSON metadata for the
// struct [PlanListChargesResponseDataUnitConversion]
type planListChargesResponseDataUnitConversionJSON struct {
DivisionFactor apijson.Field
RoundingBehavior apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponseDataUnitConversion) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// Whether usage should be rounded down or up to the nearest whole number. If null,
// quantity will be rounded to 20 decimal places.
type PlanListChargesResponseDataUnitConversionRoundingBehavior string

const (
  PlanListChargesResponseDataUnitConversionRoundingBehaviorFloor PlanListChargesResponseDataUnitConversionRoundingBehavior = "floor"
  PlanListChargesResponseDataUnitConversionRoundingBehaviorCeiling PlanListChargesResponseDataUnitConversionRoundingBehavior = "ceiling"
)

type PlanListCustomersResponse struct {
Data []PlanListCustomersResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON planListCustomersResponseJSON
}

// planListCustomersResponseJSON contains the JSON metadata for the struct
// [PlanListCustomersResponse]
type planListCustomersResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListCustomersResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListCustomersResponseData struct {
CustomerDetails CustomerDetail `json:"customer_details,required"`
PlanDetails PlanListCustomersResponseDataPlanDetails `json:"plan_details,required"`
JSON planListCustomersResponseDataJSON
}

// planListCustomersResponseDataJSON contains the JSON metadata for the struct
// [PlanListCustomersResponseData]
type planListCustomersResponseDataJSON struct {
CustomerDetails apijson.Field
PlanDetails apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListCustomersResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListCustomersResponseDataPlanDetails struct {
ID string `json:"id,required" format:"uuid"`
CustomFields map[string]string `json:"custom_fields,required"`
Name string `json:"name,required"`
// The start date of the plan
StartingOn time.Time `json:"starting_on,required" format:"date-time"`
// The end date of the plan
EndingBefore time.Time `json:"ending_before,nullable" format:"date-time"`
JSON planListCustomersResponseDataPlanDetailsJSON
}

// planListCustomersResponseDataPlanDetailsJSON contains the JSON metadata for the
// struct [PlanListCustomersResponseDataPlanDetails]
type planListCustomersResponseDataPlanDetailsJSON struct {
ID apijson.Field
CustomFields apijson.Field
Name apijson.Field
StartingOn apijson.Field
EndingBefore apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *PlanListCustomersResponseDataPlanDetails) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type PlanListParams struct {
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [PlanListParams]'s query parameters as `url.Values`.
func (r PlanListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type PlanListChargesParams struct {
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [PlanListChargesParams]'s query parameters as `url.Values`.
func (r PlanListChargesParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type PlanListCustomersParams struct {
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [PlanListCustomersParams]'s query parameters as
// `url.Values`.
func (r PlanListCustomersParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}
