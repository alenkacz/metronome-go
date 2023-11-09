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

// ContractAmendService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractAmendService] method
// instead.
type ContractAmendService struct {
	Options []option.RequestOption
}

// NewContractAmendService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractAmendService(opts ...option.RequestOption) (r *ContractAmendService) {
	r = &ContractAmendService{}
	r.Options = opts
	return
}

// Amend a contract
func (r *ContractAmendService) AmendContract(ctx context.Context, body ContractAmendAmendContractParams, opts ...option.RequestOption) (res *ContractAmendAmendContractResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/amend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractAmendAmendContractResponse struct {
	Data ContractAmendAmendContractResponseData `json:"data,required"`
	JSON contractAmendAmendContractResponseJSON
}

// contractAmendAmendContractResponseJSON contains the JSON metadata for the struct
// [ContractAmendAmendContractResponse]
type contractAmendAmendContractResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractAmendAmendContractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractAmendAmendContractResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractAmendAmendContractResponseDataJSON
}

// contractAmendAmendContractResponseDataJSON contains the JSON metadata for the
// struct [ContractAmendAmendContractResponseData]
type contractAmendAmendContractResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractAmendAmendContractResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractAmendAmendContractParams struct {
	// ID of the contract to amend
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be amended
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive start time for the amendment
	StartingAt              param.Field[time.Time]                                         `json:"starting_at,required" format:"date-time"`
	Commits                 param.Field[[]ContractAmendAmendContractParamsCommit]          `json:"commits"`
	Discounts               param.Field[[]ContractAmendAmendContractParamsDiscount]        `json:"discounts"`
	NetsuiteSalesOrderID    param.Field[string]                                            `json:"netsuite_sales_order_id"`
	Overrides               param.Field[[]ContractAmendAmendContractParamsOverride]        `json:"overrides"`
	ResellerRoyalties       param.Field[[]ContractAmendAmendContractParamsResellerRoyalty] `json:"reseller_royalties"`
	SalesforceOpportunityID param.Field[string]                                            `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractAmendAmendContractParamsScheduledCharge] `json:"scheduled_charges"`
	TotalContractValue      param.Field[float64]                                           `json:"total_contract_value"`
}

func (r ContractAmendAmendContractParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsCommit struct {
	ProductID param.Field[string]                                      `json:"product_id,required" format:"uuid"`
	Type      param.Field[ContractAmendAmendContractParamsCommitsType] `json:"type,required"`
	// Required and only valid for "PREPAID" commits: Schedule for distributing the
	// commit to the customer.
	AccessSchedule param.Field[ContractAmendAmendContractParamsCommitsAccessSchedule] `json:"access_schedule"`
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
	InvoiceSchedule param.Field[ContractAmendAmendContractParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractAmendAmendContractParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsCommitsType string

const (
	ContractAmendAmendContractParamsCommitsTypePrepaid  ContractAmendAmendContractParamsCommitsType = "PREPAID"
	ContractAmendAmendContractParamsCommitsTypePrepaid  ContractAmendAmendContractParamsCommitsType = "prepaid"
	ContractAmendAmendContractParamsCommitsTypePostpaid ContractAmendAmendContractParamsCommitsType = "POSTPAID"
	ContractAmendAmendContractParamsCommitsTypePostpaid ContractAmendAmendContractParamsCommitsType = "postpaid"
)

// Required and only valid for "PREPAID" commits: Schedule for distributing the
// commit to the customer.
type ContractAmendAmendContractParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]ContractAmendAmendContractParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
}

func (r ContractAmendAmendContractParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractAmendAmendContractParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only valid for "PREPAID" commits: If not provided, this will be a
// "complimentary" commit with no invoice.
type ContractAmendAmendContractParamsCommitsInvoiceSchedule struct {
	RecurringSchedule param.Field[ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractAmendAmendContractParamsCommitsInvoiceScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractAmendAmendContractParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                        `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendAmendContractParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractAmendAmendContractParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractAmendAmendContractParamsDiscountsSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendAmendContractParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractAmendAmendContractParamsDiscountsSchedule struct {
	RecurringSchedule param.Field[ContractAmendAmendContractParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractAmendAmendContractParamsDiscountsScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractAmendAmendContractParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                   `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractAmendAmendContractParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendAmendContractParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractAmendAmendContractParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time]                                     `json:"starting_at,required" format:"date-time"`
	Type       param.Field[ContractAmendAmendContractParamsOverridesType] `json:"type,required"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[ContractAmendAmendContractParamsOverridesOverwriteRate] `json:"overwrite_rate"`
	// ID of the product whose rate is being overridden
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// tags identifying products whose rates are being overridden
	Tags param.Field[[]string] `json:"tags"`
}

func (r ContractAmendAmendContractParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsOverridesType string

const (
	ContractAmendAmendContractParamsOverridesTypeOverwrite  ContractAmendAmendContractParamsOverridesType = "OVERWRITE"
	ContractAmendAmendContractParamsOverridesTypeOverwrite  ContractAmendAmendContractParamsOverridesType = "overwrite"
	ContractAmendAmendContractParamsOverridesTypeMultiplier ContractAmendAmendContractParamsOverridesType = "MULTIPLIER"
	ContractAmendAmendContractParamsOverridesTypeMultiplier ContractAmendAmendContractParamsOverridesType = "multiplier"
)

// Required for OVERWRITE type.
type ContractAmendAmendContractParamsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    param.Field[float64]                                                        `json:"price,required"`
	RateType param.Field[ContractAmendAmendContractParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractAmendAmendContractParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsOverridesOverwriteRateRateType string

const (
	ContractAmendAmendContractParamsOverridesOverwriteRateRateTypeFlat       ContractAmendAmendContractParamsOverridesOverwriteRateRateType = "FLAT"
	ContractAmendAmendContractParamsOverridesOverwriteRateRateTypeFlat       ContractAmendAmendContractParamsOverridesOverwriteRateRateType = "flat"
	ContractAmendAmendContractParamsOverridesOverwriteRateRateTypePercentage ContractAmendAmendContractParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractAmendAmendContractParamsOverridesOverwriteRateRateTypePercentage ContractAmendAmendContractParamsOverridesOverwriteRateRateType = "percentage"
)

type ContractAmendAmendContractParamsResellerRoyalty struct {
	ResellerType         param.Field[ContractAmendAmendContractParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	ApplicableProductIDs param.Field[[]string]                                                      `json:"applicable_product_ids" format:"uuid"`
	AwsOptions           param.Field[ContractAmendAmendContractParamsResellerRoyaltiesAwsOptions]   `json:"aws_options"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Field[time.Time]                                                   `json:"ending_before" format:"date-time"`
	Fraction              param.Field[float64]                                                     `json:"fraction"`
	GcpOptions            param.Field[ContractAmendAmendContractParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	NetsuiteResellerID    param.Field[string]                                                      `json:"netsuite_reseller_id"`
	ResellerContractValue param.Field[float64]                                                     `json:"reseller_contract_value"`
	StartingAt            param.Field[time.Time]                                                   `json:"starting_at" format:"date-time"`
}

func (r ContractAmendAmendContractParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsResellerRoyaltiesResellerType string

const (
	ContractAmendAmendContractParamsResellerRoyaltiesResellerTypeAws ContractAmendAmendContractParamsResellerRoyaltiesResellerType = "AWS"
	ContractAmendAmendContractParamsResellerRoyaltiesResellerTypeGcp ContractAmendAmendContractParamsResellerRoyaltiesResellerType = "GCP"
)

type ContractAmendAmendContractParamsResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractAmendAmendContractParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractAmendAmendContractParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractAmendAmendContractParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendAmendContractParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractAmendAmendContractParamsScheduledChargesSchedule struct {
	RecurringSchedule param.Field[ContractAmendAmendContractParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractAmendAmendContractParamsScheduledChargesScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractAmendAmendContractParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                                          `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractAmendAmendContractParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendAmendContractParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractAmendAmendContractParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
