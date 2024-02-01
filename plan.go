// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"fmt"
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
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *shared.Page[PlanListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "plans"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List all available plans.
func (r *PlanService) ListAutoPaging(ctx context.Context, query PlanListParams, opts ...option.RequestOption) *shared.PageAutoPager[PlanListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Fetch high level details of a specific plan.
func (r *PlanService) GetDetails(ctx context.Context, planID string, opts ...option.RequestOption) (res *PlanGetDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("planDetails/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a list of charges of a specific plan.
func (r *PlanService) ListCharges(ctx context.Context, planID string, query PlanListChargesParams, opts ...option.RequestOption) (res *shared.Page[PlanListChargesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("planDetails/%s/charges", planID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches a list of charges of a specific plan.
func (r *PlanService) ListChargesAutoPaging(ctx context.Context, planID string, query PlanListChargesParams, opts ...option.RequestOption) *shared.PageAutoPager[PlanListChargesResponse] {
	return shared.NewPageAutoPager(r.ListCharges(ctx, planID, query, opts...))
}

// Fetches a list of customers on a specific plan (by default, only currently
// active plans are included)
func (r *PlanService) ListCustomers(ctx context.Context, planID string, query PlanListCustomersParams, opts ...option.RequestOption) (res *shared.Page[PlanListCustomersResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("planDetails/%s/customers", planID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches a list of customers on a specific plan (by default, only currently
// active plans are included)
func (r *PlanService) ListCustomersAutoPaging(ctx context.Context, planID string, query PlanListCustomersParams, opts ...option.RequestOption) *shared.PageAutoPager[PlanListCustomersResponse] {
	return shared.NewPageAutoPager(r.ListCustomers(ctx, planID, query, opts...))
}

type PlanListResponse struct {
	ID          string               `json:"id,required" format:"uuid"`
	Description string               `json:"description,required"`
	Name        string               `json:"name,required"`
	JSON        planListResponseJSON `json:"-"`
}

// planListResponseJSON contains the JSON metadata for the struct
// [PlanListResponse]
type planListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponse struct {
	Data PlanGetDetailsResponseData `json:"data,required"`
	JSON planGetDetailsResponseJSON `json:"-"`
}

// planGetDetailsResponseJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponse]
type planGetDetailsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseData struct {
	ID           string                                  `json:"id,required" format:"uuid"`
	CustomFields map[string]string                       `json:"custom_fields,required"`
	Name         string                                  `json:"name,required"`
	CreditGrants []PlanGetDetailsResponseDataCreditGrant `json:"credit_grants"`
	Description  string                                  `json:"description"`
	Minimums     []PlanGetDetailsResponseDataMinimum     `json:"minimums"`
	OverageRates []PlanGetDetailsResponseDataOverageRate `json:"overage_rates"`
	JSON         planGetDetailsResponseDataJSON          `json:"-"`
}

// planGetDetailsResponseDataJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponseData]
type planGetDetailsResponseDataJSON struct {
	ID           apijson.Field
	CustomFields apijson.Field
	Name         apijson.Field
	CreditGrants apijson.Field
	Description  apijson.Field
	Minimums     apijson.Field
	OverageRates apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanGetDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataCreditGrant struct {
	AmountGranted           float64                                                       `json:"amount_granted,required"`
	AmountGrantedCreditType PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType `json:"amount_granted_credit_type,required"`
	AmountPaid              float64                                                       `json:"amount_paid,required"`
	AmountPaidCreditType    PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType    `json:"amount_paid_credit_type,required"`
	EffectiveDuration       float64                                                       `json:"effective_duration,required"`
	Name                    string                                                        `json:"name,required"`
	Priority                string                                                        `json:"priority,required"`
	SendInvoice             bool                                                          `json:"send_invoice,required"`
	Reason                  string                                                        `json:"reason"`
	RecurrenceDuration      float64                                                       `json:"recurrence_duration"`
	RecurrenceInterval      float64                                                       `json:"recurrence_interval"`
	JSON                    planGetDetailsResponseDataCreditGrantJSON                     `json:"-"`
}

// planGetDetailsResponseDataCreditGrantJSON contains the JSON metadata for the
// struct [PlanGetDetailsResponseDataCreditGrant]
type planGetDetailsResponseDataCreditGrantJSON struct {
	AmountGranted           apijson.Field
	AmountGrantedCreditType apijson.Field
	AmountPaid              apijson.Field
	AmountPaidCreditType    apijson.Field
	EffectiveDuration       apijson.Field
	Name                    apijson.Field
	Priority                apijson.Field
	SendInvoice             apijson.Field
	Reason                  apijson.Field
	RecurrenceDuration      apijson.Field
	RecurrenceInterval      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataCreditGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType struct {
	ID   string                                                            `json:"id,required" format:"uuid"`
	Name string                                                            `json:"name,required"`
	JSON planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON contains the
// JSON metadata for the struct
// [PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType]
type planGetDetailsResponseDataCreditGrantsAmountGrantedCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataCreditGrantsAmountGrantedCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType struct {
	ID   string                                                         `json:"id,required" format:"uuid"`
	Name string                                                         `json:"name,required"`
	JSON planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON contains the JSON
// metadata for the struct
// [PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType]
type planGetDetailsResponseDataCreditGrantsAmountPaidCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataCreditGrantsAmountPaidCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataMinimum struct {
	CreditType PlanGetDetailsResponseDataMinimumsCreditType `json:"credit_type,required"`
	Name       string                                       `json:"name,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64                               `json:"start_period,required"`
	Value       float64                               `json:"value,required"`
	JSON        planGetDetailsResponseDataMinimumJSON `json:"-"`
}

// planGetDetailsResponseDataMinimumJSON contains the JSON metadata for the struct
// [PlanGetDetailsResponseDataMinimum]
type planGetDetailsResponseDataMinimumJSON struct {
	CreditType  apijson.Field
	Name        apijson.Field
	StartPeriod apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataMinimumsCreditType struct {
	ID   string                                           `json:"id,required" format:"uuid"`
	Name string                                           `json:"name,required"`
	JSON planGetDetailsResponseDataMinimumsCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataMinimumsCreditTypeJSON contains the JSON metadata for
// the struct [PlanGetDetailsResponseDataMinimumsCreditType]
type planGetDetailsResponseDataMinimumsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataMinimumsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataOverageRate struct {
	CreditType     PlanGetDetailsResponseDataOverageRatesCreditType     `json:"credit_type,required"`
	FiatCreditType PlanGetDetailsResponseDataOverageRatesFiatCreditType `json:"fiat_credit_type,required"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod            float64                                   `json:"start_period,required"`
	ToFiatConversionFactor float64                                   `json:"to_fiat_conversion_factor,required"`
	JSON                   planGetDetailsResponseDataOverageRateJSON `json:"-"`
}

// planGetDetailsResponseDataOverageRateJSON contains the JSON metadata for the
// struct [PlanGetDetailsResponseDataOverageRate]
type planGetDetailsResponseDataOverageRateJSON struct {
	CreditType             apijson.Field
	FiatCreditType         apijson.Field
	StartPeriod            apijson.Field
	ToFiatConversionFactor apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataOverageRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataOverageRatesCreditType struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON planGetDetailsResponseDataOverageRatesCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataOverageRatesCreditTypeJSON contains the JSON metadata
// for the struct [PlanGetDetailsResponseDataOverageRatesCreditType]
type planGetDetailsResponseDataOverageRatesCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataOverageRatesCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanGetDetailsResponseDataOverageRatesFiatCreditType struct {
	ID   string                                                   `json:"id,required" format:"uuid"`
	Name string                                                   `json:"name,required"`
	JSON planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON `json:"-"`
}

// planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON contains the JSON
// metadata for the struct [PlanGetDetailsResponseDataOverageRatesFiatCreditType]
type planGetDetailsResponseDataOverageRatesFiatCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanGetDetailsResponseDataOverageRatesFiatCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponse struct {
	ID           string                            `json:"id,required" format:"uuid"`
	ChargeType   PlanListChargesResponseChargeType `json:"charge_type,required"`
	CreditType   PlanListChargesResponseCreditType `json:"credit_type,required"`
	CustomFields map[string]string                 `json:"custom_fields,required"`
	Name         string                            `json:"name,required"`
	Prices       []PlanListChargesResponsePrice    `json:"prices,required"`
	ProductID    string                            `json:"product_id,required"`
	ProductName  string                            `json:"product_name,required"`
	Quantity     float64                           `json:"quantity"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod float64 `json:"start_period"`
	// Specifies how quantities for usage based charges will be converted.
	UnitConversion PlanListChargesResponseUnitConversion `json:"unit_conversion"`
	JSON           planListChargesResponseJSON           `json:"-"`
}

// planListChargesResponseJSON contains the JSON metadata for the struct
// [PlanListChargesResponse]
type planListChargesResponseJSON struct {
	ID             apijson.Field
	ChargeType     apijson.Field
	CreditType     apijson.Field
	CustomFields   apijson.Field
	Name           apijson.Field
	Prices         apijson.Field
	ProductID      apijson.Field
	ProductName    apijson.Field
	Quantity       apijson.Field
	StartPeriod    apijson.Field
	UnitConversion apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanListChargesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponseChargeType string

const (
	PlanListChargesResponseChargeTypeUsage     PlanListChargesResponseChargeType = "usage"
	PlanListChargesResponseChargeTypeFixed     PlanListChargesResponseChargeType = "fixed"
	PlanListChargesResponseChargeTypeComposite PlanListChargesResponseChargeType = "composite"
	PlanListChargesResponseChargeTypeMinimum   PlanListChargesResponseChargeType = "minimum"
	PlanListChargesResponseChargeTypeSeat      PlanListChargesResponseChargeType = "seat"
)

type PlanListChargesResponseCreditType struct {
	ID   string                                `json:"id,required" format:"uuid"`
	Name string                                `json:"name,required"`
	JSON planListChargesResponseCreditTypeJSON `json:"-"`
}

// planListChargesResponseCreditTypeJSON contains the JSON metadata for the struct
// [PlanListChargesResponseCreditType]
type planListChargesResponseCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListChargesResponseCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListChargesResponsePrice struct {
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier               float64                          `json:"tier,required"`
	Value              float64                          `json:"value,required"`
	CollectionInterval float64                          `json:"collection_interval"`
	CollectionSchedule string                           `json:"collection_schedule"`
	Quantity           float64                          `json:"quantity"`
	JSON               planListChargesResponsePriceJSON `json:"-"`
}

// planListChargesResponsePriceJSON contains the JSON metadata for the struct
// [PlanListChargesResponsePrice]
type planListChargesResponsePriceJSON struct {
	Tier               apijson.Field
	Value              apijson.Field
	CollectionInterval apijson.Field
	CollectionSchedule apijson.Field
	Quantity           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PlanListChargesResponsePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies how quantities for usage based charges will be converted.
type PlanListChargesResponseUnitConversion struct {
	// The conversion factor
	DivisionFactor float64 `json:"division_factor,required"`
	// Whether usage should be rounded down or up to the nearest whole number. If null,
	// quantity will be rounded to 20 decimal places.
	RoundingBehavior PlanListChargesResponseUnitConversionRoundingBehavior `json:"rounding_behavior"`
	JSON             planListChargesResponseUnitConversionJSON             `json:"-"`
}

// planListChargesResponseUnitConversionJSON contains the JSON metadata for the
// struct [PlanListChargesResponseUnitConversion]
type planListChargesResponseUnitConversionJSON struct {
	DivisionFactor   apijson.Field
	RoundingBehavior apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PlanListChargesResponseUnitConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether usage should be rounded down or up to the nearest whole number. If null,
// quantity will be rounded to 20 decimal places.
type PlanListChargesResponseUnitConversionRoundingBehavior string

const (
	PlanListChargesResponseUnitConversionRoundingBehaviorFloor   PlanListChargesResponseUnitConversionRoundingBehavior = "floor"
	PlanListChargesResponseUnitConversionRoundingBehaviorCeiling PlanListChargesResponseUnitConversionRoundingBehavior = "ceiling"
)

type PlanListCustomersResponse struct {
	CustomerDetails PlanListCustomersResponseCustomerDetails `json:"customer_details,required"`
	PlanDetails     PlanListCustomersResponsePlanDetails     `json:"plan_details,required"`
	JSON            planListCustomersResponseJSON            `json:"-"`
}

// planListCustomersResponseJSON contains the JSON metadata for the struct
// [PlanListCustomersResponse]
type planListCustomersResponseJSON struct {
	CustomerDetails apijson.Field
	PlanDetails     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PlanListCustomersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListCustomersResponseCustomerDetails struct {
	// the Metronome ID of the customer
	ID                    string                                                        `json:"id,required" format:"uuid"`
	CurrentBillableStatus PlanListCustomersResponseCustomerDetailsCurrentBillableStatus `json:"current_billable_status,required"`
	CustomFields          map[string]string                                             `json:"custom_fields,required"`
	CustomerConfig        PlanListCustomersResponseCustomerDetailsCustomerConfig        `json:"customer_config,required"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
	// alias) that can be used in usage events
	ExternalID string `json:"external_id,required"`
	// aliases for this customer that can be used instead of the Metronome customer ID
	// in usage events
	IngestAliases []string                                     `json:"ingest_aliases,required"`
	Name          string                                       `json:"name,required"`
	JSON          planListCustomersResponseCustomerDetailsJSON `json:"-"`
}

// planListCustomersResponseCustomerDetailsJSON contains the JSON metadata for the
// struct [PlanListCustomersResponseCustomerDetails]
type planListCustomersResponseCustomerDetailsJSON struct {
	ID                    apijson.Field
	CurrentBillableStatus apijson.Field
	CustomFields          apijson.Field
	CustomerConfig        apijson.Field
	ExternalID            apijson.Field
	IngestAliases         apijson.Field
	Name                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PlanListCustomersResponseCustomerDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListCustomersResponseCustomerDetailsCurrentBillableStatus struct {
	Value       PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue `json:"value,required"`
	EffectiveAt time.Time                                                          `json:"effective_at,nullable" format:"date-time"`
	JSON        planListCustomersResponseCustomerDetailsCurrentBillableStatusJSON  `json:"-"`
}

// planListCustomersResponseCustomerDetailsCurrentBillableStatusJSON contains the
// JSON metadata for the struct
// [PlanListCustomersResponseCustomerDetailsCurrentBillableStatus]
type planListCustomersResponseCustomerDetailsCurrentBillableStatusJSON struct {
	Value       apijson.Field
	EffectiveAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanListCustomersResponseCustomerDetailsCurrentBillableStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue string

const (
	PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValueBillable   PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue = "billable"
	PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValueUnbillable PlanListCustomersResponseCustomerDetailsCurrentBillableStatusValue = "unbillable"
)

type PlanListCustomersResponseCustomerDetailsCustomerConfig struct {
	// The Salesforce account ID for the customer
	SalesforceAccountID string                                                     `json:"salesforce_account_id,required,nullable"`
	JSON                planListCustomersResponseCustomerDetailsCustomerConfigJSON `json:"-"`
}

// planListCustomersResponseCustomerDetailsCustomerConfigJSON contains the JSON
// metadata for the struct [PlanListCustomersResponseCustomerDetailsCustomerConfig]
type planListCustomersResponseCustomerDetailsCustomerConfigJSON struct {
	SalesforceAccountID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PlanListCustomersResponseCustomerDetailsCustomerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlanListCustomersResponsePlanDetails struct {
	ID             string            `json:"id,required" format:"uuid"`
	CustomFields   map[string]string `json:"custom_fields,required"`
	CustomerPlanID string            `json:"customer_plan_id,required" format:"uuid"`
	Name           string            `json:"name,required"`
	// The start date of the plan
	StartingOn time.Time `json:"starting_on,required" format:"date-time"`
	// The end date of the plan
	EndingBefore time.Time                                `json:"ending_before,nullable" format:"date-time"`
	JSON         planListCustomersResponsePlanDetailsJSON `json:"-"`
}

// planListCustomersResponsePlanDetailsJSON contains the JSON metadata for the
// struct [PlanListCustomersResponsePlanDetails]
type planListCustomersResponsePlanDetailsJSON struct {
	ID             apijson.Field
	CustomFields   apijson.Field
	CustomerPlanID apijson.Field
	Name           apijson.Field
	StartingOn     apijson.Field
	EndingBefore   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanListCustomersResponsePlanDetails) UnmarshalJSON(data []byte) (err error) {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanListCustomersParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Status of customers on a given plan. Defaults to `active`.
	//
	// - `all` - Return current, past, and upcoming customers of the plan.
	// - `active` - Return current customers of the plan.
	// - `ended` - Return past customers of the plan.
	// - `upcoming` - Return upcoming customers of the plan.
	//
	// Multiple statuses can be OR'd together using commas, e.g. `active,ended`.
	// **Note:** `ended,upcoming` combination is not yet supported.
	Status param.Field[PlanListCustomersParamsStatus] `query:"status"`
}

// URLQuery serializes [PlanListCustomersParams]'s query parameters as
// `url.Values`.
func (r PlanListCustomersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status of customers on a given plan. Defaults to `active`.
//
// - `all` - Return current, past, and upcoming customers of the plan.
// - `active` - Return current customers of the plan.
// - `ended` - Return past customers of the plan.
// - `upcoming` - Return upcoming customers of the plan.
//
// Multiple statuses can be OR'd together using commas, e.g. `active,ended`.
// **Note:** `ended,upcoming` combination is not yet supported.
type PlanListCustomersParamsStatus string

const (
	PlanListCustomersParamsStatusAll      PlanListCustomersParamsStatus = "all"
	PlanListCustomersParamsStatusActive   PlanListCustomersParamsStatus = "active"
	PlanListCustomersParamsStatusEnded    PlanListCustomersParamsStatus = "ended"
	PlanListCustomersParamsStatusUpcoming PlanListCustomersParamsStatus = "upcoming"
)
