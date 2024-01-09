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

// CustomerPlanService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerPlanService] method
// instead.
type CustomerPlanService struct {
	Options []option.RequestOption
}

// NewCustomerPlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerPlanService(opts ...option.RequestOption) (r *CustomerPlanService) {
	r = &CustomerPlanService{}
	r.Options = opts
	return
}

// List the given customer's plans in reverse-chronological order.
func (r *CustomerPlanService) List(ctx context.Context, customerID string, query CustomerPlanListParams, opts ...option.RequestOption) (res *shared.Page[CustomerPlanListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/%s/plans", customerID)
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

// List the given customer's plans in reverse-chronological order.
func (r *CustomerPlanService) ListAutoPaging(ctx context.Context, customerID string, query CustomerPlanListParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerPlanListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Associate an existing customer with a plan for a specified date range. See the
// [price adjustments documentation](https://docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details on the price adjustments.
func (r *CustomerPlanService) Add(ctx context.Context, customerID string, body CustomerPlanAddParams, opts ...option.RequestOption) (res *CustomerPlanAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/plans/add", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change the end date of a customer's plan.
func (r *CustomerPlanService) End(ctx context.Context, customerID string, customerPlanID string, body CustomerPlanEndParams, opts ...option.RequestOption) (res *CustomerPlanEndResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/plans/%s/end", customerID, customerPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists a customer plans adjustments. See the
// [price adjustments documentation](https://docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details.
func (r *CustomerPlanService) ListPriceAdjustments(ctx context.Context, customerID string, customerPlanID string, query CustomerPlanListPriceAdjustmentsParams, opts ...option.RequestOption) (res *shared.Page[CustomerPlanListPriceAdjustmentsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("customers/%s/plans/%s/priceAdjustments", customerID, customerPlanID)
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

// Lists a customer plans adjustments. See the
// [price adjustments documentation](https://docs.metronome.com/pricing/managing-plans/#price-adjustments)
// for details.
func (r *CustomerPlanService) ListPriceAdjustmentsAutoPaging(ctx context.Context, customerID string, customerPlanID string, query CustomerPlanListPriceAdjustmentsParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerPlanListPriceAdjustmentsResponse] {
	return shared.NewPageAutoPager(r.ListPriceAdjustments(ctx, customerID, customerPlanID, query, opts...))
}

type CustomerPlanListResponse struct {
	// the ID of the customer plan
	ID              string            `json:"id,required" format:"uuid"`
	CustomFields    map[string]string `json:"custom_fields,required"`
	PlanDescription string            `json:"plan_description,required"`
	// the ID of the plan
	PlanID              string                            `json:"plan_id,required" format:"uuid"`
	PlanName            string                            `json:"plan_name,required"`
	StartingOn          time.Time                         `json:"starting_on,required" format:"date-time"`
	EndingBefore        time.Time                         `json:"ending_before" format:"date-time"`
	NetPaymentTermsDays float64                           `json:"net_payment_terms_days"`
	TrialInfo           CustomerPlanListResponseTrialInfo `json:"trial_info"`
	JSON                customerPlanListResponseJSON      `json:"-"`
}

// customerPlanListResponseJSON contains the JSON metadata for the struct
// [CustomerPlanListResponse]
type customerPlanListResponseJSON struct {
	ID                  apijson.Field
	CustomFields        apijson.Field
	PlanDescription     apijson.Field
	PlanID              apijson.Field
	PlanName            apijson.Field
	StartingOn          apijson.Field
	EndingBefore        apijson.Field
	NetPaymentTermsDays apijson.Field
	TrialInfo           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerPlanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanListResponseTrialInfo struct {
	EndingBefore string                                         `json:"ending_before,required"`
	SpendingCaps []CustomerPlanListResponseTrialInfoSpendingCap `json:"spending_caps,required"`
	JSON         customerPlanListResponseTrialInfoJSON          `json:"-"`
}

// customerPlanListResponseTrialInfoJSON contains the JSON metadata for the struct
// [CustomerPlanListResponseTrialInfo]
type customerPlanListResponseTrialInfoJSON struct {
	EndingBefore apijson.Field
	SpendingCaps apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CustomerPlanListResponseTrialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanListResponseTrialInfoSpendingCap struct {
	Amount          float64                                                 `json:"amount,required"`
	AmountRemaining float64                                                 `json:"amount_remaining,required"`
	CreditType      CustomerPlanListResponseTrialInfoSpendingCapsCreditType `json:"credit_type,required"`
	JSON            customerPlanListResponseTrialInfoSpendingCapJSON        `json:"-"`
}

// customerPlanListResponseTrialInfoSpendingCapJSON contains the JSON metadata for
// the struct [CustomerPlanListResponseTrialInfoSpendingCap]
type customerPlanListResponseTrialInfoSpendingCapJSON struct {
	Amount          apijson.Field
	AmountRemaining apijson.Field
	CreditType      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerPlanListResponseTrialInfoSpendingCap) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanListResponseTrialInfoSpendingCapsCreditType struct {
	ID   string                                                      `json:"id,required" format:"uuid"`
	Name string                                                      `json:"name,required"`
	JSON customerPlanListResponseTrialInfoSpendingCapsCreditTypeJSON `json:"-"`
}

// customerPlanListResponseTrialInfoSpendingCapsCreditTypeJSON contains the JSON
// metadata for the struct
// [CustomerPlanListResponseTrialInfoSpendingCapsCreditType]
type customerPlanListResponseTrialInfoSpendingCapsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerPlanListResponseTrialInfoSpendingCapsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanAddResponse struct {
	Data shared.ID                   `json:"data,required"`
	JSON customerPlanAddResponseJSON `json:"-"`
}

// customerPlanAddResponseJSON contains the JSON metadata for the struct
// [CustomerPlanAddResponse]
type customerPlanAddResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerPlanAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanEndResponse struct {
	JSON customerPlanEndResponseJSON `json:"-"`
}

// customerPlanEndResponseJSON contains the JSON metadata for the struct
// [CustomerPlanEndResponse]
type customerPlanEndResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerPlanEndResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanListPriceAdjustmentsResponse struct {
	ChargeID    string                                             `json:"charge_id,required" format:"uuid"`
	ChargeType  CustomerPlanListPriceAdjustmentsResponseChargeType `json:"charge_type,required"`
	Prices      []CustomerPlanListPriceAdjustmentsResponsePrice    `json:"prices,required"`
	StartPeriod float64                                            `json:"start_period,required"`
	Quantity    float64                                            `json:"quantity"`
	JSON        customerPlanListPriceAdjustmentsResponseJSON       `json:"-"`
}

// customerPlanListPriceAdjustmentsResponseJSON contains the JSON metadata for the
// struct [CustomerPlanListPriceAdjustmentsResponse]
type customerPlanListPriceAdjustmentsResponseJSON struct {
	ChargeID    apijson.Field
	ChargeType  apijson.Field
	Prices      apijson.Field
	StartPeriod apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerPlanListPriceAdjustmentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPlanListPriceAdjustmentsResponseChargeType string

const (
	CustomerPlanListPriceAdjustmentsResponseChargeTypeUsage     CustomerPlanListPriceAdjustmentsResponseChargeType = "usage"
	CustomerPlanListPriceAdjustmentsResponseChargeTypeFixed     CustomerPlanListPriceAdjustmentsResponseChargeType = "fixed"
	CustomerPlanListPriceAdjustmentsResponseChargeTypeComposite CustomerPlanListPriceAdjustmentsResponseChargeType = "composite"
	CustomerPlanListPriceAdjustmentsResponseChargeTypeMinimum   CustomerPlanListPriceAdjustmentsResponseChargeType = "minimum"
	CustomerPlanListPriceAdjustmentsResponseChargeTypeSeat      CustomerPlanListPriceAdjustmentsResponseChargeType = "seat"
)

type CustomerPlanListPriceAdjustmentsResponsePrice struct {
	// Determines how the value will be applied.
	AdjustmentType CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType `json:"adjustment_type,required"`
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier  float64                                           `json:"tier"`
	Value float64                                           `json:"value"`
	JSON  customerPlanListPriceAdjustmentsResponsePriceJSON `json:"-"`
}

// customerPlanListPriceAdjustmentsResponsePriceJSON contains the JSON metadata for
// the struct [CustomerPlanListPriceAdjustmentsResponsePrice]
type customerPlanListPriceAdjustmentsResponsePriceJSON struct {
	AdjustmentType apijson.Field
	Tier           apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerPlanListPriceAdjustmentsResponsePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines how the value will be applied.
type CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType string

const (
	CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeFixed      CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "fixed"
	CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeQuantity   CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "quantity"
	CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypePercentage CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "percentage"
	CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentTypeOverride   CustomerPlanListPriceAdjustmentsResponsePricesAdjustmentType = "override"
)

type CustomerPlanListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CustomerPlanListParams]'s query parameters as `url.Values`.
func (r CustomerPlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerPlanAddParams struct {
	PlanID param.Field[string] `json:"plan_id,required" format:"uuid"`
	// RFC 3339 timestamp for when the plan becomes active for this customer. Must be
	// at 0:00 UTC (midnight).
	StartingOn param.Field[time.Time] `json:"starting_on,required" format:"date-time"`
	// RFC 3339 timestamp for when the plan ends (exclusive) for this customer. Must be
	// at 0:00 UTC (midnight).
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Number of days after issuance of invoice after which the invoice is due (e.g.
	// Net 30).
	NetPaymentTermsDays param.Field[float64] `json:"net_payment_terms_days"`
	// A list of price adjustments can be applied on top of the pricing in the plans.
	// See the
	// [price adjustments documentation](https://docs.metronome.com/pricing/managing-plans/#price-adjustments)
	// for details.
	PriceAdjustments param.Field[[]CustomerPlanAddParamsPriceAdjustment] `json:"price_adjustments"`
}

func (r CustomerPlanAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerPlanAddParamsPriceAdjustment struct {
	AdjustmentType param.Field[CustomerPlanAddParamsPriceAdjustmentsAdjustmentType] `json:"adjustment_type,required"`
	ChargeID       param.Field[string]                                              `json:"charge_id,required" format:"uuid"`
	// Used in price ramps. Indicates how many billing periods pass before the charge
	// applies.
	StartPeriod param.Field[float64] `json:"start_period,required"`
	// the overridden quantity for a fixed charge
	Quantity param.Field[float64] `json:"quantity"`
	// Used in pricing tiers. Indicates at what metric value the price applies.
	Tier param.Field[float64] `json:"tier"`
	// The amount of change to a price. Percentage and fixed adjustments can be
	// positive or negative. Percentage-based adjustments should be decimals, e.g.
	// -0.05 for a 5% discount.
	Value param.Field[float64] `json:"value"`
}

func (r CustomerPlanAddParamsPriceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerPlanAddParamsPriceAdjustmentsAdjustmentType string

const (
	CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypePercentage CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "percentage"
	CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeFixed      CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "fixed"
	CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeOverride   CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "override"
	CustomerPlanAddParamsPriceAdjustmentsAdjustmentTypeQuantity   CustomerPlanAddParamsPriceAdjustmentsAdjustmentType = "quantity"
)

type CustomerPlanEndParams struct {
	// RFC 3339 timestamp for when the plan ends (exclusive) for this customer. Must be
	// at 0:00 UTC (midnight). If not provided, the plan end date will be cleared.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// If true, plan end date can be before the last finalized invoice date. Any
	// invoices generated after the plan end date will be voided.
	VoidInvoices param.Field[bool] `json:"void_invoices"`
	// Only applicable when void_invoices is set to true. If true, for every invoice
	// that is voided we will also attempt to void/delete the stripe invoice (if any).
	// Stripe invoices will be voided if finalized or deleted if still in draft state.
	VoidStripeInvoices param.Field[bool] `json:"void_stripe_invoices"`
}

func (r CustomerPlanEndParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerPlanListPriceAdjustmentsParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CustomerPlanListPriceAdjustmentsParams]'s query parameters
// as `url.Values`.
func (r CustomerPlanListPriceAdjustmentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
