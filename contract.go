// File generated from our OpenAPI spec by Stainless.

package example

import (
  "github.com/example/example-go/internal/apijson"
  "github.com/example/example-go/internal/shared"
  "time"
  "github.com/example/example-go/internal/param"
  "context"
  "github.com/example/example-go/option"
  "github.com/example/example-go/internal/requestconfig"
)

// ContractService contains methods and other services that help with interacting
// with the example API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewContractService] method instead.
type ContractService struct {
Options []option.RequestOption
}

// NewContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractService(opts ...option.RequestOption) (r *ContractService) {
  r = &ContractService{}
  r.Options = opts
  return
}

// Create a new contract
func (r *ContractService) New(ctx context.Context, body ContractNewParams, opts ...option.RequestOption) (res *ContractNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contracts/create"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Get a specific contract
func (r *ContractService) Get(ctx context.Context, body ContractGetParams, opts ...option.RequestOption) (res *ContractGetResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contracts/get"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// List all contracts for a customer
func (r *ContractService) List(ctx context.Context, body ContractListParams, opts ...option.RequestOption) (res *ContractListResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contracts/list"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Amend a contract
func (r *ContractService) Amend(ctx context.Context, body ContractAmendParams, opts ...option.RequestOption) (res *ContractAmendResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contracts/amend"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Set usage filter for a contract
func (r *ContractService) SetUsageFilter(ctx context.Context, body ContractSetUsageFilterParams, opts ...option.RequestOption) (err error) {
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
  path := "contracts/setUsageFilter"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
  return
}

// Update the end date of a contract
func (r *ContractService) UpdateEndDate(ctx context.Context, body ContractUpdateEndDateParams, opts ...option.RequestOption) (res *ContractUpdateEndDateResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "contracts/updateEndDate"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

type Contract struct {
ID string `json:"id,required" format:"uuid"`
Amendments []ContractAmendment `json:"amendments,required"`
Current ContractCurrent `json:"current,required"`
CustomerID string `json:"customer_id,required" format:"uuid"`
Initial ContractInitial `json:"initial,required"`
JSON contractJSON
}

// contractJSON contains the JSON metadata for the struct [Contract]
type contractJSON struct {
ID apijson.Field
Amendments apijson.Field
Current apijson.Field
CustomerID apijson.Field
Initial apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Contract) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractAmendment struct {
ID string `json:"id,required" format:"uuid"`
Commits []shared.Commit `json:"commits,required"`
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
Discounts []shared.Discount `json:"discounts,required"`
Overrides []shared.Override `json:"overrides,required"`
ResellerRoyalties []ContractAmendmentsResellerRoyalty `json:"reseller_royalties,required"`
ScheduledCharges []ContractAmendmentsScheduledCharge `json:"scheduled_charges,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
JSON contractAmendmentJSON
}

// contractAmendmentJSON contains the JSON metadata for the struct
// [ContractAmendment]
type contractAmendmentJSON struct {
ID apijson.Field
Commits apijson.Field
CreatedAt apijson.Field
CreatedBy apijson.Field
Discounts apijson.Field
Overrides apijson.Field
ResellerRoyalties apijson.Field
ScheduledCharges apijson.Field
StartingAt apijson.Field
NetsuiteSalesOrderID apijson.Field
SalesforceOpportunityID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractAmendment) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractAmendmentsResellerRoyalty struct {
ResellerType ContractAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
AwsAccountNumber string `json:"aws_account_number"`
AwsOfferID string `json:"aws_offer_id"`
AwsPayerReferenceID string `json:"aws_payer_reference_id"`
EndingBefore time.Time `json:"ending_before,nullable" format:"date-time"`
Fraction float64 `json:"fraction"`
GcpAccountID string `json:"gcp_account_id"`
GcpOfferID string `json:"gcp_offer_id"`
NetsuiteResellerID string `json:"netsuite_reseller_id"`
ResellerContractValue float64 `json:"reseller_contract_value"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractAmendmentsResellerRoyaltyJSON
}

// contractAmendmentsResellerRoyaltyJSON contains the JSON metadata for the struct
// [ContractAmendmentsResellerRoyalty]
type contractAmendmentsResellerRoyaltyJSON struct {
ResellerType apijson.Field
AwsAccountNumber apijson.Field
AwsOfferID apijson.Field
AwsPayerReferenceID apijson.Field
EndingBefore apijson.Field
Fraction apijson.Field
GcpAccountID apijson.Field
GcpOfferID apijson.Field
NetsuiteResellerID apijson.Field
ResellerContractValue apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractAmendmentsResellerRoyaltiesResellerType string

const (
  ContractAmendmentsResellerRoyaltiesResellerTypeAws ContractAmendmentsResellerRoyaltiesResellerType = "AWS"
  ContractAmendmentsResellerRoyaltiesResellerTypeGcp ContractAmendmentsResellerRoyaltiesResellerType = "GCP"
)

type ContractAmendmentsScheduledCharge struct {
ID string `json:"id,required" format:"uuid"`
Product ContractAmendmentsScheduledChargesProduct `json:"product,required"`
Schedule shared.SchedulePointInTime `json:"schedule,required"`
// displayed on invoices
Name string `json:"name"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
JSON contractAmendmentsScheduledChargeJSON
}

// contractAmendmentsScheduledChargeJSON contains the JSON metadata for the struct
// [ContractAmendmentsScheduledCharge]
type contractAmendmentsScheduledChargeJSON struct {
ID apijson.Field
Product apijson.Field
Schedule apijson.Field
Name apijson.Field
NetsuiteSalesOrderID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractAmendmentsScheduledCharge) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractAmendmentsScheduledChargesProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON contractAmendmentsScheduledChargesProductJSON
}

// contractAmendmentsScheduledChargesProductJSON contains the JSON metadata for the
// struct [ContractAmendmentsScheduledChargesProduct]
type contractAmendmentsScheduledChargesProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractAmendmentsScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrent struct {
Commits []shared.Commit `json:"commits,required"`
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
Discounts []shared.Discount `json:"discounts,required"`
Overrides []shared.Override `json:"overrides,required"`
ResellerRoyalties []ContractCurrentResellerRoyalty `json:"reseller_royalties,required"`
ScheduledCharges []ContractCurrentScheduledCharge `json:"scheduled_charges,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
Transitions []ContractCurrentTransition `json:"transitions,required"`
UsageInvoiceSchedule ContractCurrentUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
Name string `json:"name"`
NetPaymentTermsDays float64 `json:"net_payment_terms_days"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
RateCardID string `json:"rate_card_id" format:"uuid"`
SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
TotalContractValue float64 `json:"total_contract_value"`
UsageFilter ContractCurrentUsageFilter `json:"usage_filter"`
JSON contractCurrentJSON
}

// contractCurrentJSON contains the JSON metadata for the struct [ContractCurrent]
type contractCurrentJSON struct {
Commits apijson.Field
CreatedAt apijson.Field
CreatedBy apijson.Field
Discounts apijson.Field
Overrides apijson.Field
ResellerRoyalties apijson.Field
ScheduledCharges apijson.Field
StartingAt apijson.Field
Transitions apijson.Field
UsageInvoiceSchedule apijson.Field
EndingBefore apijson.Field
Name apijson.Field
NetPaymentTermsDays apijson.Field
NetsuiteSalesOrderID apijson.Field
RateCardID apijson.Field
SalesforceOpportunityID apijson.Field
TotalContractValue apijson.Field
UsageFilter apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrent) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentResellerRoyalty struct {
Fraction float64 `json:"fraction,required"`
NetsuiteResellerID string `json:"netsuite_reseller_id,required"`
ResellerType ContractCurrentResellerRoyaltiesResellerType `json:"reseller_type,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
AwsAccountNumber string `json:"aws_account_number"`
AwsOfferID string `json:"aws_offer_id"`
AwsPayerReferenceID string `json:"aws_payer_reference_id"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
GcpAccountID string `json:"gcp_account_id"`
GcpOfferID string `json:"gcp_offer_id"`
ResellerContractValue float64 `json:"reseller_contract_value"`
JSON contractCurrentResellerRoyaltyJSON
}

// contractCurrentResellerRoyaltyJSON contains the JSON metadata for the struct
// [ContractCurrentResellerRoyalty]
type contractCurrentResellerRoyaltyJSON struct {
Fraction apijson.Field
NetsuiteResellerID apijson.Field
ResellerType apijson.Field
StartingAt apijson.Field
AwsAccountNumber apijson.Field
AwsOfferID apijson.Field
AwsPayerReferenceID apijson.Field
EndingBefore apijson.Field
GcpAccountID apijson.Field
GcpOfferID apijson.Field
ResellerContractValue apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentResellerRoyaltiesResellerType string

const (
  ContractCurrentResellerRoyaltiesResellerTypeAws ContractCurrentResellerRoyaltiesResellerType = "AWS"
  ContractCurrentResellerRoyaltiesResellerTypeGcp ContractCurrentResellerRoyaltiesResellerType = "GCP"
)

type ContractCurrentScheduledCharge struct {
ID string `json:"id,required" format:"uuid"`
Product ContractCurrentScheduledChargesProduct `json:"product,required"`
Schedule shared.SchedulePointInTime `json:"schedule,required"`
// displayed on invoices
Name string `json:"name"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
JSON contractCurrentScheduledChargeJSON
}

// contractCurrentScheduledChargeJSON contains the JSON metadata for the struct
// [ContractCurrentScheduledCharge]
type contractCurrentScheduledChargeJSON struct {
ID apijson.Field
Product apijson.Field
Schedule apijson.Field
Name apijson.Field
NetsuiteSalesOrderID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentScheduledCharge) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentScheduledChargesProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON contractCurrentScheduledChargesProductJSON
}

// contractCurrentScheduledChargesProductJSON contains the JSON metadata for the
// struct [ContractCurrentScheduledChargesProduct]
type contractCurrentScheduledChargesProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentTransition struct {
FromContractID string `json:"from_contract_id,required" format:"uuid"`
ToContractID string `json:"to_contract_id,required" format:"uuid"`
Type ContractCurrentTransitionsType `json:"type,required"`
JSON contractCurrentTransitionJSON
}

// contractCurrentTransitionJSON contains the JSON metadata for the struct
// [ContractCurrentTransition]
type contractCurrentTransitionJSON struct {
FromContractID apijson.Field
ToContractID apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentTransition) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentTransitionsType string

const (
  ContractCurrentTransitionsTypeSupersede ContractCurrentTransitionsType = "SUPERSEDE"
  ContractCurrentTransitionsTypeRenewal ContractCurrentTransitionsType = "RENEWAL"
)

type ContractCurrentUsageInvoiceSchedule struct {
Frequency ContractCurrentUsageInvoiceScheduleFrequency `json:"frequency,required"`
JSON contractCurrentUsageInvoiceScheduleJSON
}

// contractCurrentUsageInvoiceScheduleJSON contains the JSON metadata for the
// struct [ContractCurrentUsageInvoiceSchedule]
type contractCurrentUsageInvoiceScheduleJSON struct {
Frequency apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentUsageInvoiceScheduleFrequency string

const (
  ContractCurrentUsageInvoiceScheduleFrequencyMonthly ContractCurrentUsageInvoiceScheduleFrequency = "MONTHLY"
  ContractCurrentUsageInvoiceScheduleFrequencyMonthly ContractCurrentUsageInvoiceScheduleFrequency = "monthly"
  ContractCurrentUsageInvoiceScheduleFrequencyQuarterly ContractCurrentUsageInvoiceScheduleFrequency = "QUARTERLY"
  ContractCurrentUsageInvoiceScheduleFrequencyQuarterly ContractCurrentUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractCurrentUsageFilter struct {
Current ContractCurrentUsageFilterCurrent `json:"current,required"`
Initial ContractCurrentUsageFilterInitial `json:"initial,required"`
Updates []ContractCurrentUsageFilterUpdate `json:"updates,required"`
JSON contractCurrentUsageFilterJSON
}

// contractCurrentUsageFilterJSON contains the JSON metadata for the struct
// [ContractCurrentUsageFilter]
type contractCurrentUsageFilterJSON struct {
Current apijson.Field
Initial apijson.Field
Updates apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentUsageFilter) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentUsageFilterCurrent struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractCurrentUsageFilterCurrentJSON
}

// contractCurrentUsageFilterCurrentJSON contains the JSON metadata for the struct
// [ContractCurrentUsageFilterCurrent]
type contractCurrentUsageFilterCurrentJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentUsageFilterInitial struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractCurrentUsageFilterInitialJSON
}

// contractCurrentUsageFilterInitialJSON contains the JSON metadata for the struct
// [ContractCurrentUsageFilterInitial]
type contractCurrentUsageFilterInitialJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractCurrentUsageFilterUpdate struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
JSON contractCurrentUsageFilterUpdateJSON
}

// contractCurrentUsageFilterUpdateJSON contains the JSON metadata for the struct
// [ContractCurrentUsageFilterUpdate]
type contractCurrentUsageFilterUpdateJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractCurrentUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitial struct {
Commits []shared.Commit `json:"commits,required"`
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
Discounts []shared.Discount `json:"discounts,required"`
Overrides []shared.Override `json:"overrides,required"`
ResellerRoyalties []ContractInitialResellerRoyalty `json:"reseller_royalties,required"`
ScheduledCharges []ContractInitialScheduledCharge `json:"scheduled_charges,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
Transitions []ContractInitialTransition `json:"transitions,required"`
UsageInvoiceSchedule ContractInitialUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
Name string `json:"name"`
NetPaymentTermsDays float64 `json:"net_payment_terms_days"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
RateCardID string `json:"rate_card_id" format:"uuid"`
SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
TotalContractValue float64 `json:"total_contract_value"`
UsageFilter ContractInitialUsageFilter `json:"usage_filter"`
JSON contractInitialJSON
}

// contractInitialJSON contains the JSON metadata for the struct [ContractInitial]
type contractInitialJSON struct {
Commits apijson.Field
CreatedAt apijson.Field
CreatedBy apijson.Field
Discounts apijson.Field
Overrides apijson.Field
ResellerRoyalties apijson.Field
ScheduledCharges apijson.Field
StartingAt apijson.Field
Transitions apijson.Field
UsageInvoiceSchedule apijson.Field
EndingBefore apijson.Field
Name apijson.Field
NetPaymentTermsDays apijson.Field
NetsuiteSalesOrderID apijson.Field
RateCardID apijson.Field
SalesforceOpportunityID apijson.Field
TotalContractValue apijson.Field
UsageFilter apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitial) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialResellerRoyalty struct {
Fraction float64 `json:"fraction,required"`
NetsuiteResellerID string `json:"netsuite_reseller_id,required"`
ResellerType ContractInitialResellerRoyaltiesResellerType `json:"reseller_type,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
AwsAccountNumber string `json:"aws_account_number"`
AwsOfferID string `json:"aws_offer_id"`
AwsPayerReferenceID string `json:"aws_payer_reference_id"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
GcpAccountID string `json:"gcp_account_id"`
GcpOfferID string `json:"gcp_offer_id"`
ResellerContractValue float64 `json:"reseller_contract_value"`
JSON contractInitialResellerRoyaltyJSON
}

// contractInitialResellerRoyaltyJSON contains the JSON metadata for the struct
// [ContractInitialResellerRoyalty]
type contractInitialResellerRoyaltyJSON struct {
Fraction apijson.Field
NetsuiteResellerID apijson.Field
ResellerType apijson.Field
StartingAt apijson.Field
AwsAccountNumber apijson.Field
AwsOfferID apijson.Field
AwsPayerReferenceID apijson.Field
EndingBefore apijson.Field
GcpAccountID apijson.Field
GcpOfferID apijson.Field
ResellerContractValue apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialResellerRoyaltiesResellerType string

const (
  ContractInitialResellerRoyaltiesResellerTypeAws ContractInitialResellerRoyaltiesResellerType = "AWS"
  ContractInitialResellerRoyaltiesResellerTypeGcp ContractInitialResellerRoyaltiesResellerType = "GCP"
)

type ContractInitialScheduledCharge struct {
ID string `json:"id,required" format:"uuid"`
Product ContractInitialScheduledChargesProduct `json:"product,required"`
Schedule shared.SchedulePointInTime `json:"schedule,required"`
// displayed on invoices
Name string `json:"name"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
JSON contractInitialScheduledChargeJSON
}

// contractInitialScheduledChargeJSON contains the JSON metadata for the struct
// [ContractInitialScheduledCharge]
type contractInitialScheduledChargeJSON struct {
ID apijson.Field
Product apijson.Field
Schedule apijson.Field
Name apijson.Field
NetsuiteSalesOrderID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialScheduledCharge) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialScheduledChargesProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON contractInitialScheduledChargesProductJSON
}

// contractInitialScheduledChargesProductJSON contains the JSON metadata for the
// struct [ContractInitialScheduledChargesProduct]
type contractInitialScheduledChargesProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialTransition struct {
FromContractID string `json:"from_contract_id,required" format:"uuid"`
ToContractID string `json:"to_contract_id,required" format:"uuid"`
Type ContractInitialTransitionsType `json:"type,required"`
JSON contractInitialTransitionJSON
}

// contractInitialTransitionJSON contains the JSON metadata for the struct
// [ContractInitialTransition]
type contractInitialTransitionJSON struct {
FromContractID apijson.Field
ToContractID apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialTransition) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialTransitionsType string

const (
  ContractInitialTransitionsTypeSupersede ContractInitialTransitionsType = "SUPERSEDE"
  ContractInitialTransitionsTypeRenewal ContractInitialTransitionsType = "RENEWAL"
)

type ContractInitialUsageInvoiceSchedule struct {
Frequency ContractInitialUsageInvoiceScheduleFrequency `json:"frequency,required"`
JSON contractInitialUsageInvoiceScheduleJSON
}

// contractInitialUsageInvoiceScheduleJSON contains the JSON metadata for the
// struct [ContractInitialUsageInvoiceSchedule]
type contractInitialUsageInvoiceScheduleJSON struct {
Frequency apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialUsageInvoiceScheduleFrequency string

const (
  ContractInitialUsageInvoiceScheduleFrequencyMonthly ContractInitialUsageInvoiceScheduleFrequency = "MONTHLY"
  ContractInitialUsageInvoiceScheduleFrequencyMonthly ContractInitialUsageInvoiceScheduleFrequency = "monthly"
  ContractInitialUsageInvoiceScheduleFrequencyQuarterly ContractInitialUsageInvoiceScheduleFrequency = "QUARTERLY"
  ContractInitialUsageInvoiceScheduleFrequencyQuarterly ContractInitialUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractInitialUsageFilter struct {
Current ContractInitialUsageFilterCurrent `json:"current,required"`
Initial ContractInitialUsageFilterInitial `json:"initial,required"`
Updates []ContractInitialUsageFilterUpdate `json:"updates,required"`
JSON contractInitialUsageFilterJSON
}

// contractInitialUsageFilterJSON contains the JSON metadata for the struct
// [ContractInitialUsageFilter]
type contractInitialUsageFilterJSON struct {
Current apijson.Field
Initial apijson.Field
Updates apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialUsageFilter) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialUsageFilterCurrent struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractInitialUsageFilterCurrentJSON
}

// contractInitialUsageFilterCurrentJSON contains the JSON metadata for the struct
// [ContractInitialUsageFilterCurrent]
type contractInitialUsageFilterCurrentJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialUsageFilterInitial struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractInitialUsageFilterInitialJSON
}

// contractInitialUsageFilterInitialJSON contains the JSON metadata for the struct
// [ContractInitialUsageFilterInitial]
type contractInitialUsageFilterInitialJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractInitialUsageFilterUpdate struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
JSON contractInitialUsageFilterUpdateJSON
}

// contractInitialUsageFilterUpdateJSON contains the JSON metadata for the struct
// [ContractInitialUsageFilterUpdate]
type contractInitialUsageFilterUpdateJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractInitialUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractNewResponse struct {
Data shared.ID `json:"data,required"`
JSON contractNewResponseJSON
}

// contractNewResponseJSON contains the JSON metadata for the struct
// [ContractNewResponse]
type contractNewResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponse struct {
Data Contract `json:"data,required"`
JSON contractGetResponseJSON
}

// contractGetResponseJSON contains the JSON metadata for the struct
// [ContractGetResponse]
type contractGetResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractListResponse struct {
Data []Contract `json:"data,required"`
JSON contractListResponseJSON
}

// contractListResponseJSON contains the JSON metadata for the struct
// [ContractListResponse]
type contractListResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractAmendResponse struct {
Data shared.ID `json:"data,required"`
JSON contractAmendResponseJSON
}

// contractAmendResponseJSON contains the JSON metadata for the struct
// [ContractAmendResponse]
type contractAmendResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractAmendResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractUpdateEndDateResponse struct {
Data shared.ID `json:"data,required"`
JSON contractUpdateEndDateResponseJSON
}

// contractUpdateEndDateResponseJSON contains the JSON metadata for the struct
// [ContractUpdateEndDateResponse]
type contractUpdateEndDateResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractNewParams struct {
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// inclusive contract start time
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
Commits param.Field[[]ContractNewParamsCommit] `json:"commits"`
Discounts param.Field[[]ContractNewParamsDiscount] `json:"discounts"`
// exclusive contract end time
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
Name param.Field[string] `json:"name"`
NetPaymentTermsDays param.Field[float64] `json:"net_payment_terms_days"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
Overrides param.Field[[]ContractNewParamsOverride] `json:"overrides"`
RateCardID param.Field[string] `json:"rate_card_id" format:"uuid"`
ResellerRoyalties param.Field[[]ContractNewParamsResellerRoyalty] `json:"reseller_royalties"`
SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
ScheduledCharges param.Field[[]ContractNewParamsScheduledCharge] `json:"scheduled_charges"`
TotalContractValue param.Field[float64] `json:"total_contract_value"`
Transition param.Field[ContractNewParamsTransition] `json:"transition"`
UsageFilter param.Field[ContractNewParamsUsageFilter] `json:"usage_filter"`
UsageInvoiceSchedule param.Field[ContractNewParamsUsageInvoiceSchedule] `json:"usage_invoice_schedule"`
}

func (r ContractNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsCommit struct {
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
Type param.Field[ContractNewParamsCommitsType] `json:"type,required"`
// Required and only valid for "PREPAID" commits: Schedule for distributing the
// commit to the customer.
AccessSchedule param.Field[ContractNewParamsCommitsAccessSchedule] `json:"access_schedule"`
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
InvoiceSchedule param.Field[ContractNewParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
// displayed on invoices
Name param.Field[string] `json:"name"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractNewParamsCommit) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsType string

const (
  ContractNewParamsCommitsTypePrepaid ContractNewParamsCommitsType = "PREPAID"
  ContractNewParamsCommitsTypePrepaid ContractNewParamsCommitsType = "prepaid"
  ContractNewParamsCommitsTypePostpaid ContractNewParamsCommitsType = "POSTPAID"
  ContractNewParamsCommitsTypePostpaid ContractNewParamsCommitsType = "postpaid"
)

// Required and only valid for "PREPAID" commits: Schedule for distributing the
// commit to the customer.
type ContractNewParamsCommitsAccessSchedule struct {
ScheduleItems param.Field[[]ContractNewParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
}

func (r ContractNewParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsAccessScheduleScheduleItem struct {
Amount param.Field[float64] `json:"amount,required"`
// RFC 3339 timestamp (exclusive)
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
// RFC 3339 timestamp (inclusive)
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractNewParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Only valid for "PREPAID" commits: If not provided, this will be a
// "complimentary" commit with no invoice.
type ContractNewParamsCommitsInvoiceSchedule struct {
RecurringSchedule param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
ScheduleItems param.Field[[]ContractNewParamsCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractNewParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsInvoiceScheduleRecurringSchedule struct {
AmountDistribution param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
// RFC 3339 timestamp (exclusive).
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
Frequency param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractNewParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractNewParamsCommitsInvoiceScheduleScheduleItem struct {
// timestamp of the scheduled event
Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
Amount param.Field[float64] `json:"amount"`
Quantity param.Field[float64] `json:"quantity"`
UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsDiscount struct {
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
// Must provide either schedule_items or recurring_schedule
Schedule param.Field[ContractNewParamsDiscountsSchedule] `json:"schedule,required"`
// displayed on invoices
Name param.Field[string] `json:"name"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsDiscount) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractNewParamsDiscountsSchedule struct {
RecurringSchedule param.Field[ContractNewParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
ScheduleItems param.Field[[]ContractNewParamsDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractNewParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsDiscountsScheduleRecurringSchedule struct {
AmountDistribution param.Field[ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
// RFC 3339 timestamp (exclusive).
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
Frequency param.Field[ContractNewParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractNewParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
  ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
  ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
  ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
  ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
  ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
  ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractNewParamsDiscountsScheduleRecurringScheduleFrequency string

const (
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
  ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractNewParamsDiscountsScheduleScheduleItem struct {
// timestamp of the scheduled event
Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
Amount param.Field[float64] `json:"amount"`
Quantity param.Field[float64] `json:"quantity"`
UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsOverride struct {
// RFC 3339 timestamp indicating when the override will start applying (inclusive)
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
Type param.Field[ContractNewParamsOverridesType] `json:"type,required"`
// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
Entitled param.Field[bool] `json:"entitled"`
// Required for MULTIPLIER type. Must be >=0.
Multiplier param.Field[float64] `json:"multiplier"`
// Required for OVERWRITE type.
OverwriteRate param.Field[RateParam] `json:"overwrite_rate"`
// ID of the product whose rate is being overridden
ProductID param.Field[string] `json:"product_id" format:"uuid"`
// tags identifying products whose rates are being overridden
Tags param.Field[[]string] `json:"tags"`
}

func (r ContractNewParamsOverride) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsOverridesType string

const (
  ContractNewParamsOverridesTypeOverwrite ContractNewParamsOverridesType = "OVERWRITE"
  ContractNewParamsOverridesTypeOverwrite ContractNewParamsOverridesType = "overwrite"
  ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "MULTIPLIER"
  ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "multiplier"
)

type ContractNewParamsResellerRoyalty struct {
ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids,required" format:"uuid"`
Fraction param.Field[float64] `json:"fraction,required"`
NetsuiteResellerID param.Field[string] `json:"netsuite_reseller_id,required"`
ResellerType param.Field[ContractNewParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
AwsOptions param.Field[ContractNewParamsResellerRoyaltiesAwsOptions] `json:"aws_options"`
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
GcpOptions param.Field[ContractNewParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
ResellerContractValue param.Field[float64] `json:"reseller_contract_value"`
}

func (r ContractNewParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsResellerRoyaltiesResellerType string

const (
  ContractNewParamsResellerRoyaltiesResellerTypeAws ContractNewParamsResellerRoyaltiesResellerType = "AWS"
  ContractNewParamsResellerRoyaltiesResellerTypeGcp ContractNewParamsResellerRoyaltiesResellerType = "GCP"
)

type ContractNewParamsResellerRoyaltiesAwsOptions struct {
AwsAccountNumber param.Field[string] `json:"aws_account_number"`
AwsOfferID param.Field[string] `json:"aws_offer_id"`
AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractNewParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsResellerRoyaltiesGcpOptions struct {
GcpAccountID param.Field[string] `json:"gcp_account_id"`
GcpOfferID param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractNewParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledCharge struct {
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
// Must provide either schedule_items or recurring_schedule
Schedule param.Field[ContractNewParamsScheduledChargesSchedule] `json:"schedule,required"`
// displayed on invoices
Name param.Field[string] `json:"name"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractNewParamsScheduledChargesSchedule struct {
RecurringSchedule param.Field[ContractNewParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
ScheduleItems param.Field[[]ContractNewParamsScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractNewParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledChargesScheduleRecurringSchedule struct {
AmountDistribution param.Field[ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
// RFC 3339 timestamp (exclusive).
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
Frequency param.Field[ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractNewParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
  ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractNewParamsScheduledChargesScheduleScheduleItem struct {
// timestamp of the scheduled event
Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
Amount param.Field[float64] `json:"amount"`
Quantity param.Field[float64] `json:"quantity"`
UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractNewParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsTransition struct {
FromContractID param.Field[string] `json:"from_contract_id,required" format:"uuid"`
Type param.Field[ContractNewParamsTransitionType] `json:"type,required"`
}

func (r ContractNewParamsTransition) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsTransitionType string

const (
  ContractNewParamsTransitionTypeSupersede ContractNewParamsTransitionType = "SUPERSEDE"
  ContractNewParamsTransitionTypeRenewal ContractNewParamsTransitionType = "RENEWAL"
  ContractNewParamsTransitionTypeSupersede ContractNewParamsTransitionType = "supersede"
  ContractNewParamsTransitionTypeRenewal ContractNewParamsTransitionType = "renewal"
)

type ContractNewParamsUsageFilter struct {
GroupKey param.Field[string] `json:"group_key,required"`
GroupValues param.Field[[]string] `json:"group_values,required"`
StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractNewParamsUsageFilter) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsUsageInvoiceSchedule struct {
Frequency param.Field[ContractNewParamsUsageInvoiceScheduleFrequency] `json:"frequency,required"`
}

func (r ContractNewParamsUsageInvoiceSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractNewParamsUsageInvoiceScheduleFrequency string

const (
  ContractNewParamsUsageInvoiceScheduleFrequencyMonthly ContractNewParamsUsageInvoiceScheduleFrequency = "MONTHLY"
  ContractNewParamsUsageInvoiceScheduleFrequencyMonthly ContractNewParamsUsageInvoiceScheduleFrequency = "monthly"
  ContractNewParamsUsageInvoiceScheduleFrequencyQuarterly ContractNewParamsUsageInvoiceScheduleFrequency = "QUARTERLY"
  ContractNewParamsUsageInvoiceScheduleFrequencyQuarterly ContractNewParamsUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractGetParams struct {
ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// Include commit ledgers in the response. Setting this flag may cause the query to
// be slower.
IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r ContractGetParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractListParams struct {
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// Include commit ledgers in the response. Setting this flag may cause the query to
// be slower.
IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r ContractListParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParams struct {
// ID of the contract to amend
ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
// ID of the customer whose contract is to be amended
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// inclusive start time for the amendment
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
Commits param.Field[[]ContractAmendParamsCommit] `json:"commits"`
Discounts param.Field[[]ContractAmendParamsDiscount] `json:"discounts"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
Overrides param.Field[[]ContractAmendParamsOverride] `json:"overrides"`
ResellerRoyalties param.Field[[]ContractAmendParamsResellerRoyalty] `json:"reseller_royalties"`
SalesforceOpportunityID param.Field[string] `json:"salesforce_opportunity_id"`
ScheduledCharges param.Field[[]ContractAmendParamsScheduledCharge] `json:"scheduled_charges"`
TotalContractValue param.Field[float64] `json:"total_contract_value"`
}

func (r ContractAmendParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommit struct {
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
Type param.Field[ContractAmendParamsCommitsType] `json:"type,required"`
// Required and only valid for "PREPAID" commits: Schedule for distributing the
// commit to the customer.
AccessSchedule param.Field[ContractAmendParamsCommitsAccessSchedule] `json:"access_schedule"`
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
InvoiceSchedule param.Field[ContractAmendParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
// displayed on invoices
Name param.Field[string] `json:"name"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractAmendParamsCommit) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsType string

const (
  ContractAmendParamsCommitsTypePrepaid ContractAmendParamsCommitsType = "PREPAID"
  ContractAmendParamsCommitsTypePrepaid ContractAmendParamsCommitsType = "prepaid"
  ContractAmendParamsCommitsTypePostpaid ContractAmendParamsCommitsType = "POSTPAID"
  ContractAmendParamsCommitsTypePostpaid ContractAmendParamsCommitsType = "postpaid"
)

// Required and only valid for "PREPAID" commits: Schedule for distributing the
// commit to the customer.
type ContractAmendParamsCommitsAccessSchedule struct {
ScheduleItems param.Field[[]ContractAmendParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
}

func (r ContractAmendParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsAccessScheduleScheduleItem struct {
Amount param.Field[float64] `json:"amount,required"`
// RFC 3339 timestamp (exclusive)
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
// RFC 3339 timestamp (inclusive)
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractAmendParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Only valid for "PREPAID" commits: If not provided, this will be a
// "complimentary" commit with no invoice.
type ContractAmendParamsCommitsInvoiceSchedule struct {
RecurringSchedule param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
ScheduleItems param.Field[[]ContractAmendParamsCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractAmendParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule struct {
AmountDistribution param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
// RFC 3339 timestamp (exclusive).
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
Frequency param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendParamsCommitsInvoiceScheduleScheduleItem struct {
// timestamp of the scheduled event
Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
Amount param.Field[float64] `json:"amount"`
Quantity param.Field[float64] `json:"quantity"`
UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsDiscount struct {
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
// Must provide either schedule_items or recurring_schedule
Schedule param.Field[ContractAmendParamsDiscountsSchedule] `json:"schedule,required"`
// displayed on invoices
Name param.Field[string] `json:"name"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsDiscount) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractAmendParamsDiscountsSchedule struct {
RecurringSchedule param.Field[ContractAmendParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
ScheduleItems param.Field[[]ContractAmendParamsDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractAmendParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsDiscountsScheduleRecurringSchedule struct {
AmountDistribution param.Field[ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
// RFC 3339 timestamp (exclusive).
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
Frequency param.Field[ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractAmendParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
  ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
  ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
  ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
  ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
  ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
  ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency string

const (
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendParamsDiscountsScheduleScheduleItem struct {
// timestamp of the scheduled event
Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
Amount param.Field[float64] `json:"amount"`
Quantity param.Field[float64] `json:"quantity"`
UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverride struct {
// RFC 3339 timestamp indicating when the override will start applying (inclusive)
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
Type param.Field[ContractAmendParamsOverridesType] `json:"type,required"`
// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
Entitled param.Field[bool] `json:"entitled"`
// Required for MULTIPLIER type. Must be >=0.
Multiplier param.Field[float64] `json:"multiplier"`
// Required for OVERWRITE type.
OverwriteRate param.Field[RateParam] `json:"overwrite_rate"`
// ID of the product whose rate is being overridden
ProductID param.Field[string] `json:"product_id" format:"uuid"`
// tags identifying products whose rates are being overridden
Tags param.Field[[]string] `json:"tags"`
}

func (r ContractAmendParamsOverride) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverridesType string

const (
  ContractAmendParamsOverridesTypeOverwrite ContractAmendParamsOverridesType = "OVERWRITE"
  ContractAmendParamsOverridesTypeOverwrite ContractAmendParamsOverridesType = "overwrite"
  ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "MULTIPLIER"
  ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "multiplier"
)

type ContractAmendParamsResellerRoyalty struct {
ResellerType param.Field[ContractAmendParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
AwsOptions param.Field[ContractAmendParamsResellerRoyaltiesAwsOptions] `json:"aws_options"`
// Use null to indicate that the existing end timestamp should be removed.
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
Fraction param.Field[float64] `json:"fraction"`
GcpOptions param.Field[ContractAmendParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
NetsuiteResellerID param.Field[string] `json:"netsuite_reseller_id"`
ResellerContractValue param.Field[float64] `json:"reseller_contract_value"`
StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r ContractAmendParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsResellerRoyaltiesResellerType string

const (
  ContractAmendParamsResellerRoyaltiesResellerTypeAws ContractAmendParamsResellerRoyaltiesResellerType = "AWS"
  ContractAmendParamsResellerRoyaltiesResellerTypeGcp ContractAmendParamsResellerRoyaltiesResellerType = "GCP"
)

type ContractAmendParamsResellerRoyaltiesAwsOptions struct {
AwsAccountNumber param.Field[string] `json:"aws_account_number"`
AwsOfferID param.Field[string] `json:"aws_offer_id"`
AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractAmendParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsResellerRoyaltiesGcpOptions struct {
GcpAccountID param.Field[string] `json:"gcp_account_id"`
GcpOfferID param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractAmendParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledCharge struct {
ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
// Must provide either schedule_items or recurring_schedule
Schedule param.Field[ContractAmendParamsScheduledChargesSchedule] `json:"schedule,required"`
// displayed on invoices
Name param.Field[string] `json:"name"`
NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractAmendParamsScheduledChargesSchedule struct {
RecurringSchedule param.Field[ContractAmendParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
ScheduleItems param.Field[[]ContractAmendParamsScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r ContractAmendParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledChargesScheduleRecurringSchedule struct {
AmountDistribution param.Field[ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
// RFC 3339 timestamp (exclusive).
EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
Frequency param.Field[ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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

func (r ContractAmendParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendParamsScheduledChargesScheduleScheduleItem struct {
// timestamp of the scheduled event
Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
Amount param.Field[float64] `json:"amount"`
Quantity param.Field[float64] `json:"quantity"`
UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r ContractAmendParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractSetUsageFilterParams struct {
ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
GroupKey param.Field[string] `json:"group_key,required"`
GroupValues param.Field[[]string] `json:"group_values,required"`
StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r ContractSetUsageFilterParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type ContractUpdateEndDateParams struct {
// ID of the contract to update
ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
// ID of the customer whose contract is to be updated
CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
// RFC 3339 timestamp indicating when the contract will end (exclusive). If not
// provided, the contract will be updated to be open-ended.
EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
}

func (r ContractUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
