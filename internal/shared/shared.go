// File generated from our OpenAPI spec by Stainless.

package shared

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "time"
  "github.com/tidwall/gjson"
  "github.com/metronome/metronome-go/internal/param"
)

type Commit struct {
ID string `json:"id,required" format:"uuid"`
Product CommitProduct `json:"product,required"`
Type CommitType `json:"type,required"`
// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
AccessSchedule CommitAccessSchedule `json:"access_schedule"`
// Only valid for "POSTPAID" commits: The total that the customer commits to
// consume. Must be >= 0.
Amount float64 `json:"amount"`
ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
ApplicableTags []string `json:"applicable_tags"`
Description string `json:"description"`
// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
InvoiceSchedule SchedulePointInTime `json:"invoice_schedule"`
// A list of ordered events that impact the balance of a commit. For example, an
// invoice deduction or a rollover.
Ledger []CommitLedger `json:"ledger"`
Name string `json:"name"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
RolledOverFrom CommitRolledOverFrom `json:"rolled_over_from"`
RolloverFraction float64 `json:"rollover_fraction"`
JSON commitJSON
}

// commitJSON contains the JSON metadata for the struct [Commit]
type commitJSON struct {
ID apijson.Field
Product apijson.Field
Type apijson.Field
AccessSchedule apijson.Field
Amount apijson.Field
ApplicableProductIDs apijson.Field
ApplicableTags apijson.Field
Description apijson.Field
InvoiceSchedule apijson.Field
Ledger apijson.Field
Name apijson.Field
NetsuiteSalesOrderID apijson.Field
RolledOverFrom apijson.Field
RolloverFraction apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Commit) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CommitProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON commitProductJSON
}

// commitProductJSON contains the JSON metadata for the struct [CommitProduct]
type commitProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CommitType string

const (
  CommitTypePrepaid CommitType = "PREPAID"
  CommitTypePostpaid CommitType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type CommitAccessSchedule struct {
ScheduleItems []CommitAccessScheduleScheduleItem `json:"schedule_items,required"`
JSON commitAccessScheduleJSON
}

// commitAccessScheduleJSON contains the JSON metadata for the struct
// [CommitAccessSchedule]
type commitAccessScheduleJSON struct {
ScheduleItems apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitAccessSchedule) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CommitAccessScheduleScheduleItem struct {
ID string `json:"id,required" format:"uuid"`
Amount float64 `json:"amount,required"`
EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
JSON commitAccessScheduleScheduleItemJSON
}

// commitAccessScheduleScheduleItemJSON contains the JSON metadata for the struct
// [CommitAccessScheduleScheduleItem]
type commitAccessScheduleScheduleItemJSON struct {
ID apijson.Field
Amount apijson.Field
EndingBefore apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry] or
// [CommitLedgerPostpaidCommitTrueupLedgerEntry].
type CommitLedger interface {
  implementsSharedCommitLedger()
}

func init() {
  apijson.RegisterUnion(reflect.TypeOf((*CommitLedger)(nil)).Elem(), "")
}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntry struct {
Amount float64 `json:"amount,required"`
SegmentID string `json:"segment_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
JSON commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPrepaidCommitSegmentStartLedgerEntry]
type commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
Amount apijson.Field
SegmentID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
  CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart CommitLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
Amount float64 `json:"amount,required"`
InvoiceID string `json:"invoice_id,required" format:"uuid"`
SegmentID string `json:"segment_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
JSON commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON contains the
// JSON metadata for the struct
// [CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
Amount apijson.Field
InvoiceID apijson.Field
SegmentID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
  CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type CommitLedgerPrepaidCommitRolloverLedgerEntry struct {
Amount float64 `json:"amount,required"`
NewContractID string `json:"new_contract_id,required" format:"uuid"`
SegmentID string `json:"segment_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
JSON commitLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// commitLedgerPrepaidCommitRolloverLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitRolloverLedgerEntry]
type commitLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
Amount apijson.Field
NewContractID apijson.Field
SegmentID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPrepaidCommitRolloverLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitRolloverLedgerEntryType string

const (
  CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover CommitLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type CommitLedgerPrepaidCommitExpirationLedgerEntry struct {
Amount float64 `json:"amount,required"`
SegmentID string `json:"segment_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
JSON commitLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// commitLedgerPrepaidCommitExpirationLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPrepaidCommitExpirationLedgerEntry]
type commitLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
Amount apijson.Field
SegmentID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPrepaidCommitExpirationLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitExpirationLedgerEntryType string

const (
  CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration CommitLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type CommitLedgerPrepaidCommitCanceledLedgerEntry struct {
Amount float64 `json:"amount,required"`
InvoiceID string `json:"invoice_id,required" format:"uuid"`
SegmentID string `json:"segment_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
JSON commitLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// commitLedgerPrepaidCommitCanceledLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCanceledLedgerEntry]
type commitLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
Amount apijson.Field
InvoiceID apijson.Field
SegmentID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPrepaidCommitCanceledLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitCanceledLedgerEntryType string

const (
  CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled CommitLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type CommitLedgerPrepaidCommitCreditedLedgerEntry struct {
Amount float64 `json:"amount,required"`
InvoiceID string `json:"invoice_id,required" format:"uuid"`
SegmentID string `json:"segment_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
JSON commitLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// commitLedgerPrepaidCommitCreditedLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCreditedLedgerEntry]
type commitLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
Amount apijson.Field
InvoiceID apijson.Field
SegmentID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPrepaidCommitCreditedLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPrepaidCommitCreditedLedgerEntryType string

const (
  CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited CommitLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
Amount float64 `json:"amount,required"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
JSON commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON contains the JSON
// metadata for the struct [CommitLedgerPostpaidCommitInitialBalanceLedgerEntry]
type commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
Amount apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
  CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
Amount float64 `json:"amount,required"`
InvoiceID string `json:"invoice_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
JSON commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON contains the
// JSON metadata for the struct
// [CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
Amount apijson.Field
InvoiceID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
  CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type CommitLedgerPostpaidCommitTrueupLedgerEntry struct {
Amount float64 `json:"amount,required"`
InvoiceID string `json:"invoice_id,required" format:"uuid"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Type CommitLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
JSON commitLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// commitLedgerPostpaidCommitTrueupLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitTrueupLedgerEntry]
type commitLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
Amount apijson.Field
InvoiceID apijson.Field
Timestamp apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r CommitLedgerPostpaidCommitTrueupLedgerEntry) implementsSharedCommitLedger() {}

type CommitLedgerPostpaidCommitTrueupLedgerEntryType string

const (
  CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup CommitLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type CommitRolledOverFrom struct {
CommitID string `json:"commit_id,required" format:"uuid"`
ContractID string `json:"contract_id,required" format:"uuid"`
JSON commitRolledOverFromJSON
}

// commitRolledOverFromJSON contains the JSON metadata for the struct
// [CommitRolledOverFrom]
type commitRolledOverFromJSON struct {
CommitID apijson.Field
ContractID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CommitRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendments struct {
Commits []Commit `json:"commits,required"`
CreatedAt time.Time `json:"created_at,required" format:"date-time"`
CreatedBy string `json:"created_by,required"`
Discounts []Discount `json:"discounts,required"`
Overrides []Override `json:"overrides,required"`
ResellerRoyalties []ContractWithoutAmendmentsResellerRoyalty `json:"reseller_royalties,required"`
ScheduledCharges []ScheduledCharge `json:"scheduled_charges,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
Transitions []ContractWithoutAmendmentsTransition `json:"transitions,required"`
UsageInvoiceSchedule ContractWithoutAmendmentsUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
Name string `json:"name"`
NetPaymentTermsDays float64 `json:"net_payment_terms_days"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
RateCardID string `json:"rate_card_id" format:"uuid"`
SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
TotalContractValue float64 `json:"total_contract_value"`
UsageFilter ContractWithoutAmendmentsUsageFilter `json:"usage_filter"`
JSON contractWithoutAmendmentsJSON
}

// contractWithoutAmendmentsJSON contains the JSON metadata for the struct
// [ContractWithoutAmendments]
type contractWithoutAmendmentsJSON struct {
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

func (r *ContractWithoutAmendments) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsResellerRoyalty struct {
Fraction float64 `json:"fraction,required"`
NetsuiteResellerID string `json:"netsuite_reseller_id,required"`
ResellerType ContractWithoutAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
AwsAccountNumber string `json:"aws_account_number"`
AwsOfferID string `json:"aws_offer_id"`
AwsPayerReferenceID string `json:"aws_payer_reference_id"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
GcpAccountID string `json:"gcp_account_id"`
GcpOfferID string `json:"gcp_offer_id"`
ResellerContractValue float64 `json:"reseller_contract_value"`
JSON contractWithoutAmendmentsResellerRoyaltyJSON
}

// contractWithoutAmendmentsResellerRoyaltyJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsResellerRoyalty]
type contractWithoutAmendmentsResellerRoyaltyJSON struct {
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

func (r *ContractWithoutAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsResellerRoyaltiesResellerType string

const (
  ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws ContractWithoutAmendmentsResellerRoyaltiesResellerType = "AWS"
  ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp ContractWithoutAmendmentsResellerRoyaltiesResellerType = "GCP"
)

type ContractWithoutAmendmentsTransition struct {
FromContractID string `json:"from_contract_id,required" format:"uuid"`
ToContractID string `json:"to_contract_id,required" format:"uuid"`
Type ContractWithoutAmendmentsTransitionsType `json:"type,required"`
JSON contractWithoutAmendmentsTransitionJSON
}

// contractWithoutAmendmentsTransitionJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsTransition]
type contractWithoutAmendmentsTransitionJSON struct {
FromContractID apijson.Field
ToContractID apijson.Field
Type apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsTransition) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsTransitionsType string

const (
  ContractWithoutAmendmentsTransitionsTypeSupersede ContractWithoutAmendmentsTransitionsType = "SUPERSEDE"
  ContractWithoutAmendmentsTransitionsTypeRenewal ContractWithoutAmendmentsTransitionsType = "RENEWAL"
)

type ContractWithoutAmendmentsUsageInvoiceSchedule struct {
Frequency ContractWithoutAmendmentsUsageInvoiceScheduleFrequency `json:"frequency,required"`
JSON contractWithoutAmendmentsUsageInvoiceScheduleJSON
}

// contractWithoutAmendmentsUsageInvoiceScheduleJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageInvoiceSchedule]
type contractWithoutAmendmentsUsageInvoiceScheduleJSON struct {
Frequency apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsUsageInvoiceScheduleFrequency string

const (
  ContractWithoutAmendmentsUsageInvoiceScheduleFrequencyMonthly ContractWithoutAmendmentsUsageInvoiceScheduleFrequency = "monthly"
  ContractWithoutAmendmentsUsageInvoiceScheduleFrequencyQuarterly ContractWithoutAmendmentsUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractWithoutAmendmentsUsageFilter struct {
Current ContractWithoutAmendmentsUsageFilterCurrent `json:"current,required"`
Initial ContractWithoutAmendmentsUsageFilterInitial `json:"initial,required"`
Updates []ContractWithoutAmendmentsUsageFilterUpdate `json:"updates,required"`
JSON contractWithoutAmendmentsUsageFilterJSON
}

// contractWithoutAmendmentsUsageFilterJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsUsageFilter]
type contractWithoutAmendmentsUsageFilterJSON struct {
Current apijson.Field
Initial apijson.Field
Updates apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilter) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsUsageFilterCurrent struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractWithoutAmendmentsUsageFilterCurrentJSON
}

// contractWithoutAmendmentsUsageFilterCurrentJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterCurrent]
type contractWithoutAmendmentsUsageFilterCurrentJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsUsageFilterInitial struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at" format:"date-time"`
JSON contractWithoutAmendmentsUsageFilterInitialJSON
}

// contractWithoutAmendmentsUsageFilterInitialJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterInitial]
type contractWithoutAmendmentsUsageFilterInitialJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ContractWithoutAmendmentsUsageFilterUpdate struct {
GroupKey string `json:"group_key,required"`
GroupValues []string `json:"group_values,required"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
JSON contractWithoutAmendmentsUsageFilterUpdateJSON
}

// contractWithoutAmendmentsUsageFilterUpdateJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterUpdate]
type contractWithoutAmendmentsUsageFilterUpdateJSON struct {
GroupKey apijson.Field
GroupValues apijson.Field
StartingAt apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type Discount struct {
ID string `json:"id,required" format:"uuid"`
Product DiscountProduct `json:"product,required"`
Schedule SchedulePointInTime `json:"schedule,required"`
Name string `json:"name"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
JSON discountJSON
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
ID apijson.Field
Product apijson.Field
Schedule apijson.Field
Name apijson.Field
NetsuiteSalesOrderID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type DiscountProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON discountProductJSON
}

// discountProductJSON contains the JSON metadata for the struct [DiscountProduct]
type discountProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DiscountProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ID struct {
ID string `json:"id,required" format:"uuid"`
JSON idJSON
}

// idJSON contains the JSON metadata for the struct [ID]
type idJSON struct {
ID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ID) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type Override struct {
ID string `json:"id,required" format:"uuid"`
StartingAt time.Time `json:"starting_at,required" format:"date-time"`
Type OverrideType `json:"type,required"`
EndingBefore time.Time `json:"ending_before" format:"date-time"`
Entitled bool `json:"entitled"`
Multiplier float64 `json:"multiplier"`
OverwriteRate Rate `json:"overwrite_rate"`
Product OverrideProduct `json:"product"`
Tags []string `json:"tags"`
JSON overrideJSON
}

// overrideJSON contains the JSON metadata for the struct [Override]
type overrideJSON struct {
ID apijson.Field
StartingAt apijson.Field
Type apijson.Field
EndingBefore apijson.Field
Entitled apijson.Field
Multiplier apijson.Field
OverwriteRate apijson.Field
Product apijson.Field
Tags apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Override) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type OverrideType string

const (
  OverrideTypeOverwrite OverrideType = "OVERWRITE"
  OverrideTypeMultiplier OverrideType = "MULTIPLIER"
)

type OverrideProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON overrideProductJSON
}

// overrideProductJSON contains the JSON metadata for the struct [OverrideProduct]
type overrideProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *OverrideProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type Rate struct {
// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
Price float64 `json:"price,required"`
RateType RateRateType `json:"rate_type,required"`
// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
// using list prices rather than the standard rates for this product on the
// contract.
UseListPrices bool `json:"use_list_prices"`
JSON rateJSON
}

// rateJSON contains the JSON metadata for the struct [Rate]
type rateJSON struct {
Price apijson.Field
RateType apijson.Field
UseListPrices apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Rate) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type RateRateType string

const (
  RateRateTypeFlat RateRateType = "FLAT"
  RateRateTypePercentage RateRateType = "PERCENTAGE"
)

type RateParam struct {
// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
Price param.Field[float64] `json:"price,required"`
RateType param.Field[RateRateType] `json:"rate_type,required"`
// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
// using list prices rather than the standard rates for this product on the
// contract.
UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r RateParam) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type SchedulePointInTime struct {
ScheduleItems []SchedulePointInTimeScheduleItem `json:"schedule_items"`
JSON schedulePointInTimeJSON
}

// schedulePointInTimeJSON contains the JSON metadata for the struct
// [SchedulePointInTime]
type schedulePointInTimeJSON struct {
ScheduleItems apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SchedulePointInTime) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type SchedulePointInTimeScheduleItem struct {
ID string `json:"id,required" format:"uuid"`
Amount float64 `json:"amount,required"`
InvoiceID string `json:"invoice_id,required" format:"uuid"`
Quantity float64 `json:"quantity,required"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
UnitPrice float64 `json:"unit_price,required"`
JSON schedulePointInTimeScheduleItemJSON
}

// schedulePointInTimeScheduleItemJSON contains the JSON metadata for the struct
// [SchedulePointInTimeScheduleItem]
type schedulePointInTimeScheduleItemJSON struct {
ID apijson.Field
Amount apijson.Field
InvoiceID apijson.Field
Quantity apijson.Field
Timestamp apijson.Field
UnitPrice apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *SchedulePointInTimeScheduleItem) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ScheduledCharge struct {
ID string `json:"id,required" format:"uuid"`
Product ScheduledChargeProduct `json:"product,required"`
Schedule SchedulePointInTime `json:"schedule,required"`
// displayed on invoices
Name string `json:"name"`
NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
JSON scheduledChargeJSON
}

// scheduledChargeJSON contains the JSON metadata for the struct [ScheduledCharge]
type scheduledChargeJSON struct {
ID apijson.Field
Product apijson.Field
Schedule apijson.Field
Name apijson.Field
NetsuiteSalesOrderID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ScheduledCharge) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type ScheduledChargeProduct struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
JSON scheduledChargeProductJSON
}

// scheduledChargeProductJSON contains the JSON metadata for the struct
// [ScheduledChargeProduct]
type scheduledChargeProductJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *ScheduledChargeProduct) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}
