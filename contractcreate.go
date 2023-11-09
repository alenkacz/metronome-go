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

// ContractCreateService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractCreateService] method
// instead.
type ContractCreateService struct {
	Options []option.RequestOption
}

// NewContractCreateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractCreateService(opts ...option.RequestOption) (r *ContractCreateService) {
	r = &ContractCreateService{}
	r.Options = opts
	return
}

// Create a new contract
func (r *ContractCreateService) NewContract(ctx context.Context, body ContractCreateNewContractParams, opts ...option.RequestOption) (res *ContractCreateNewContractResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractCreateNewContractResponse struct {
	Data ContractCreateNewContractResponseData `json:"data,required"`
	JSON contractCreateNewContractResponseJSON
}

// contractCreateNewContractResponseJSON contains the JSON metadata for the struct
// [ContractCreateNewContractResponse]
type contractCreateNewContractResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractCreateNewContractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractCreateNewContractResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractCreateNewContractResponseDataJSON
}

// contractCreateNewContractResponseDataJSON contains the JSON metadata for the
// struct [ContractCreateNewContractResponseData]
type contractCreateNewContractResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractCreateNewContractResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractCreateNewContractParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive contract start time
	StartingAt param.Field[time.Time]                                 `json:"starting_at,required" format:"date-time"`
	Commits    param.Field[[]ContractCreateNewContractParamsCommit]   `json:"commits"`
	Discounts  param.Field[[]ContractCreateNewContractParamsDiscount] `json:"discounts"`
	// exclusive contract end time
	EndingBefore            param.Field[time.Time]                                           `json:"ending_before" format:"date-time"`
	Name                    param.Field[string]                                              `json:"name"`
	NetPaymentTermsDays     param.Field[float64]                                             `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    param.Field[string]                                              `json:"netsuite_sales_order_id"`
	Overrides               param.Field[[]ContractCreateNewContractParamsOverride]           `json:"overrides"`
	RateCardID              param.Field[string]                                              `json:"rate_card_id" format:"uuid"`
	ResellerRoyalties       param.Field[[]ContractCreateNewContractParamsResellerRoyalty]    `json:"reseller_royalties"`
	SalesforceOpportunityID param.Field[string]                                              `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractCreateNewContractParamsScheduledCharge]    `json:"scheduled_charges"`
	TotalContractValue      param.Field[float64]                                             `json:"total_contract_value"`
	Transition              param.Field[ContractCreateNewContractParamsTransition]           `json:"transition"`
	UsageFilter             param.Field[ContractCreateNewContractParamsUsageFilter]          `json:"usage_filter"`
	UsageInvoiceSchedule    param.Field[ContractCreateNewContractParamsUsageInvoiceSchedule] `json:"usage_invoice_schedule"`
}

func (r ContractCreateNewContractParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsCommit struct {
	ProductID param.Field[string]                                     `json:"product_id,required" format:"uuid"`
	Type      param.Field[ContractCreateNewContractParamsCommitsType] `json:"type,required"`
	// Required and only valid for "PREPAID" commits: Schedule for distributing the
	// commit to the customer.
	AccessSchedule param.Field[ContractCreateNewContractParamsCommitsAccessSchedule] `json:"access_schedule"`
	// Required and only valid for "POSTPAID" commits: The total that the customer
	// commits to consume. Must be >= 0.
	Amount param.Field[float64] `json:"amount"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_tags are not provided, the commit applies to all products.
	ApplicableTags param.Field[[]string] `json:"applicable_tags"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// Only valid for "PREPAID" commits: If not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[ContractCreateNewContractParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractCreateNewContractParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsCommitsType string

const (
	ContractCreateNewContractParamsCommitsTypePrepaid  ContractCreateNewContractParamsCommitsType = "PREPAID"
	ContractCreateNewContractParamsCommitsTypePrepaid  ContractCreateNewContractParamsCommitsType = "prepaid"
	ContractCreateNewContractParamsCommitsTypePostpaid ContractCreateNewContractParamsCommitsType = "POSTPAID"
	ContractCreateNewContractParamsCommitsTypePostpaid ContractCreateNewContractParamsCommitsType = "postpaid"
)

// Required and only valid for "PREPAID" commits: Schedule for distributing the
// commit to the customer.
type ContractCreateNewContractParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractCreateNewContractParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
}

func (r ContractCreateNewContractParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractCreateNewContractParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only valid for "PREPAID" commits: If not provided, this will be a
// "complimentary" commit with no invoice.
type ContractCreateNewContractParamsCommitsInvoiceSchedule struct {
	RecurringSchedule param.Field[ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractCreateNewContractParamsCommitsInvoiceScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractCreateNewContractParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                       `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// specified, unit_price and quantity cannot be provided.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractCreateNewContractParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractCreateNewContractParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractCreateNewContractParamsDiscountsSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractCreateNewContractParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractCreateNewContractParamsDiscountsSchedule struct {
	RecurringSchedule param.Field[ContractCreateNewContractParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractCreateNewContractParamsDiscountsScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractCreateNewContractParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                  `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// specified, unit_price and quantity cannot be provided.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractCreateNewContractParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractCreateNewContractParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractCreateNewContractParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time]                                    `json:"starting_at,required" format:"date-time"`
	Type       param.Field[ContractCreateNewContractParamsOverridesType] `json:"type,required"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[ContractCreateNewContractParamsOverridesOverwriteRate] `json:"overwrite_rate"`
	// ID of the product whose rate is being overridden
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// tags identifying products whose rates are being overridden
	Tags param.Field[[]string] `json:"tags"`
}

func (r ContractCreateNewContractParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsOverridesType string

const (
	ContractCreateNewContractParamsOverridesTypeOverwrite  ContractCreateNewContractParamsOverridesType = "OVERWRITE"
	ContractCreateNewContractParamsOverridesTypeOverwrite  ContractCreateNewContractParamsOverridesType = "overwrite"
	ContractCreateNewContractParamsOverridesTypeMultiplier ContractCreateNewContractParamsOverridesType = "MULTIPLIER"
	ContractCreateNewContractParamsOverridesTypeMultiplier ContractCreateNewContractParamsOverridesType = "multiplier"
)

// Required for OVERWRITE type.
type ContractCreateNewContractParamsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    param.Field[float64]                                                       `json:"price,required"`
	RateType param.Field[ContractCreateNewContractParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractCreateNewContractParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsOverridesOverwriteRateRateType string

const (
	ContractCreateNewContractParamsOverridesOverwriteRateRateTypeFlat       ContractCreateNewContractParamsOverridesOverwriteRateRateType = "FLAT"
	ContractCreateNewContractParamsOverridesOverwriteRateRateTypeFlat       ContractCreateNewContractParamsOverridesOverwriteRateRateType = "flat"
	ContractCreateNewContractParamsOverridesOverwriteRateRateTypePercentage ContractCreateNewContractParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractCreateNewContractParamsOverridesOverwriteRateRateTypePercentage ContractCreateNewContractParamsOverridesOverwriteRateRateType = "percentage"
)

type ContractCreateNewContractParamsResellerRoyalty struct {
	ApplicableProductIDs  param.Field[[]string]                                                     `json:"applicable_product_ids,required" format:"uuid"`
	Fraction              param.Field[float64]                                                      `json:"fraction,required"`
	NetsuiteResellerID    param.Field[string]                                                       `json:"netsuite_reseller_id,required"`
	ResellerType          param.Field[ContractCreateNewContractParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	StartingAt            param.Field[time.Time]                                                    `json:"starting_at,required" format:"date-time"`
	AwsOptions            param.Field[ContractCreateNewContractParamsResellerRoyaltiesAwsOptions]   `json:"aws_options"`
	EndingBefore          param.Field[time.Time]                                                    `json:"ending_before" format:"date-time"`
	GcpOptions            param.Field[ContractCreateNewContractParamsResellerRoyaltiesGcpOptions]   `json:"gcp_options"`
	ResellerContractValue param.Field[float64]                                                      `json:"reseller_contract_value"`
}

func (r ContractCreateNewContractParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsResellerRoyaltiesResellerType string

const (
	ContractCreateNewContractParamsResellerRoyaltiesResellerTypeAws ContractCreateNewContractParamsResellerRoyaltiesResellerType = "AWS"
	ContractCreateNewContractParamsResellerRoyaltiesResellerTypeGcp ContractCreateNewContractParamsResellerRoyaltiesResellerType = "GCP"
)

type ContractCreateNewContractParamsResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractCreateNewContractParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractCreateNewContractParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractCreateNewContractParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractCreateNewContractParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractCreateNewContractParamsScheduledChargesSchedule struct {
	RecurringSchedule param.Field[ContractCreateNewContractParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractCreateNewContractParamsScheduledChargesScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractCreateNewContractParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                         `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// specified, unit_price and quantity cannot be provided.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractCreateNewContractParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractCreateNewContractParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractCreateNewContractParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsTransition struct {
	FromContractID param.Field[string]                                        `json:"from_contract_id,required" format:"uuid"`
	Type           param.Field[ContractCreateNewContractParamsTransitionType] `json:"type,required"`
}

func (r ContractCreateNewContractParamsTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsTransitionType string

const (
	ContractCreateNewContractParamsTransitionTypeSupersede ContractCreateNewContractParamsTransitionType = "SUPERSEDE"
	ContractCreateNewContractParamsTransitionTypeRenewal   ContractCreateNewContractParamsTransitionType = "RENEWAL"
	ContractCreateNewContractParamsTransitionTypeSupersede ContractCreateNewContractParamsTransitionType = "supersede"
	ContractCreateNewContractParamsTransitionTypeRenewal   ContractCreateNewContractParamsTransitionType = "renewal"
)

type ContractCreateNewContractParamsUsageFilter struct {
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractCreateNewContractParamsUsageFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsUsageInvoiceSchedule struct {
	Frequency param.Field[ContractCreateNewContractParamsUsageInvoiceScheduleFrequency] `json:"frequency,required"`
}

func (r ContractCreateNewContractParamsUsageInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractCreateNewContractParamsUsageInvoiceScheduleFrequency string

const (
	ContractCreateNewContractParamsUsageInvoiceScheduleFrequencyMonthly   ContractCreateNewContractParamsUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractCreateNewContractParamsUsageInvoiceScheduleFrequencyMonthly   ContractCreateNewContractParamsUsageInvoiceScheduleFrequency = "monthly"
	ContractCreateNewContractParamsUsageInvoiceScheduleFrequencyQuarterly ContractCreateNewContractParamsUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractCreateNewContractParamsUsageInvoiceScheduleFrequencyQuarterly ContractCreateNewContractParamsUsageInvoiceScheduleFrequency = "quarterly"
)
