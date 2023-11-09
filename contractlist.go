// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
)

// ContractListService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractListService] method
// instead.
type ContractListService struct {
	Options []option.RequestOption
}

// NewContractListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractListService(opts ...option.RequestOption) (r *ContractListService) {
	r = &ContractListService{}
	r.Options = opts
	return
}

// List all contracts for a customer
func (r *ContractListService) ListContracts(ctx context.Context, body ContractListListContractsParams, opts ...option.RequestOption) (res *ContractListListContractsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractListListContractsResponse struct {
	Data []ContractListListContractsResponseData `json:"data,required"`
	JSON contractListListContractsResponseJSON
}

// contractListListContractsResponseJSON contains the JSON metadata for the struct
// [ContractListListContractsResponse]
type contractListListContractsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseData struct {
	ID         string                                           `json:"id,required" format:"uuid"`
	Amendments []ContractListListContractsResponseDataAmendment `json:"amendments,required"`
	Current    ContractListListContractsResponseDataCurrent     `json:"current,required"`
	CustomerID string                                           `json:"customer_id,required" format:"uuid"`
	Initial    ContractListListContractsResponseDataInitial     `json:"initial,required"`
	JSON       contractListListContractsResponseDataJSON
}

// contractListListContractsResponseDataJSON contains the JSON metadata for the
// struct [ContractListListContractsResponseData]
type contractListListContractsResponseDataJSON struct {
	ID          apijson.Field
	Amendments  apijson.Field
	Current     apijson.Field
	CustomerID  apijson.Field
	Initial     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendment struct {
	ID                      string                                                           `json:"id,required" format:"uuid"`
	Commits                 []ContractListListContractsResponseDataAmendmentsCommit          `json:"commits,required"`
	CreatedAt               time.Time                                                        `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                                           `json:"created_by,required"`
	Discounts               []ContractListListContractsResponseDataAmendmentsDiscount        `json:"discounts,required"`
	Overrides               []ContractListListContractsResponseDataAmendmentsOverride        `json:"overrides,required"`
	ResellerRoyalties       []ContractListListContractsResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractListListContractsResponseDataAmendmentsScheduledCharge `json:"scheduled_charges,required"`
	StartingAt              time.Time                                                        `json:"starting_at,required" format:"date-time"`
	NetsuiteSalesOrderID    string                                                           `json:"netsuite_sales_order_id"`
	SalesforceOpportunityID string                                                           `json:"salesforce_opportunity_id"`
	JSON                    contractListListContractsResponseDataAmendmentJSON
}

// contractListListContractsResponseDataAmendmentJSON contains the JSON metadata
// for the struct [ContractListListContractsResponseDataAmendment]
type contractListListContractsResponseDataAmendmentJSON struct {
	ID                      apijson.Field
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Discounts               apijson.Field
	Overrides               apijson.Field
	ResellerRoyalties       apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsCommit struct {
	ID      string                                                        `json:"id,required" format:"uuid"`
	Product ContractListListContractsResponseDataAmendmentsCommitsProduct `json:"product,required"`
	Type    ContractListListContractsResponseDataAmendmentsCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractListListContractsResponseDataAmendmentsCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractListListContractsResponseDataAmendmentsCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractListListContractsResponseDataAmendmentsCommitsLedger       `json:"ledger"`
	Name                 string                                                               `json:"name"`
	NetsuiteSalesOrderID string                                                               `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractListListContractsResponseDataAmendmentsCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                              `json:"rollover_fraction"`
	JSON                 contractListListContractsResponseDataAmendmentsCommitJSON
}

// contractListListContractsResponseDataAmendmentsCommitJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataAmendmentsCommit]
type contractListListContractsResponseDataAmendmentsCommitJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Type                 apijson.Field
	AccessSchedule       apijson.Field
	Amount               apijson.Field
	ApplicableProductIDs apijson.Field
	ApplicableTags       apijson.Field
	Description          apijson.Field
	InvoiceSchedule      apijson.Field
	Ledger               apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	RolledOverFrom       apijson.Field
	RolloverFraction     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataAmendmentsCommitsProductJSON
}

// contractListListContractsResponseDataAmendmentsCommitsProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsProduct]
type contractListListContractsResponseDataAmendmentsCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsCommitsType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsTypePrepaid  ContractListListContractsResponseDataAmendmentsCommitsType = "PREPAID"
	ContractListListContractsResponseDataAmendmentsCommitsTypePostpaid ContractListListContractsResponseDataAmendmentsCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractListListContractsResponseDataAmendmentsCommitsAccessSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractListListContractsResponseDataAmendmentsCommitsAccessScheduleJSON
}

// contractListListContractsResponseDataAmendmentsCommitsAccessScheduleJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsAccessSchedule]
type contractListListContractsResponseDataAmendmentsCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON
}

// contractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItem]
type contractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractListListContractsResponseDataAmendmentsCommitsInvoiceSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleJSON
}

// contractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsInvoiceSchedule]
type contractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
}

// contractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem]
type contractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractListListContractsResponseDataAmendmentsCommitsLedger interface {
	implementsContractListListContractsResponseDataAmendmentsCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractListListContractsResponseDataAmendmentsCommitsLedger)(nil)).Elem(), "")
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                              `json:"amount,required"`
	SegmentID string                                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                           `json:"amount,required"`
	InvoiceID string                                                                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                                          `json:"amount,required"`
	NewContractID string                                                                                           `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                                        `json:"timestamp,required" format:"date-time"`
	Type          ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                                            `json:"amount,required"`
	SegmentID string                                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                                          `json:"amount,required"`
	InvoiceID string                                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                                          `json:"amount,required"`
	InvoiceID string                                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractListListContractsResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                                 `json:"amount,required"`
	Timestamp time.Time                                                                                               `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                            `json:"amount,required"`
	InvoiceID string                                                                                                             `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                                         `json:"amount,required"`
	InvoiceID string                                                                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractListListContractsResponseDataAmendmentsCommitsLedger() {
}

type ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractListListContractsResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractListListContractsResponseDataAmendmentsCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractListListContractsResponseDataAmendmentsCommitsRolledOverFromJSON
}

// contractListListContractsResponseDataAmendmentsCommitsRolledOverFromJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsCommitsRolledOverFrom]
type contractListListContractsResponseDataAmendmentsCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsDiscount struct {
	ID                   string                                                           `json:"id,required" format:"uuid"`
	Product              ContractListListContractsResponseDataAmendmentsDiscountsProduct  `json:"product,required"`
	Schedule             ContractListListContractsResponseDataAmendmentsDiscountsSchedule `json:"schedule,required"`
	Name                 string                                                           `json:"name"`
	NetsuiteSalesOrderID string                                                           `json:"netsuite_sales_order_id"`
	JSON                 contractListListContractsResponseDataAmendmentsDiscountJSON
}

// contractListListContractsResponseDataAmendmentsDiscountJSON contains the JSON
// metadata for the struct
// [ContractListListContractsResponseDataAmendmentsDiscount]
type contractListListContractsResponseDataAmendmentsDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataAmendmentsDiscountsProductJSON
}

// contractListListContractsResponseDataAmendmentsDiscountsProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsDiscountsProduct]
type contractListListContractsResponseDataAmendmentsDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsDiscountsSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataAmendmentsDiscountsScheduleJSON
}

// contractListListContractsResponseDataAmendmentsDiscountsScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsDiscountsSchedule]
type contractListListContractsResponseDataAmendmentsDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItemJSON
}

// contractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItem]
type contractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsOverride struct {
	ID            string                                                                `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                             `json:"starting_at,required" format:"date-time"`
	Type          ContractListListContractsResponseDataAmendmentsOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                             `json:"ending_before" format:"date-time"`
	Entitled      bool                                                                  `json:"entitled"`
	Multiplier    float64                                                               `json:"multiplier"`
	OverwriteRate ContractListListContractsResponseDataAmendmentsOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractListListContractsResponseDataAmendmentsOverridesProduct       `json:"product"`
	Tags          []string                                                              `json:"tags"`
	JSON          contractListListContractsResponseDataAmendmentsOverrideJSON
}

// contractListListContractsResponseDataAmendmentsOverrideJSON contains the JSON
// metadata for the struct
// [ContractListListContractsResponseDataAmendmentsOverride]
type contractListListContractsResponseDataAmendmentsOverrideJSON struct {
	ID            apijson.Field
	StartingAt    apijson.Field
	Type          apijson.Field
	EndingBefore  apijson.Field
	Entitled      apijson.Field
	Multiplier    apijson.Field
	OverwriteRate apijson.Field
	Product       apijson.Field
	Tags          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsOverridesType string

const (
	ContractListListContractsResponseDataAmendmentsOverridesTypeOverwrite  ContractListListContractsResponseDataAmendmentsOverridesType = "OVERWRITE"
	ContractListListContractsResponseDataAmendmentsOverridesTypeMultiplier ContractListListContractsResponseDataAmendmentsOverridesType = "MULTIPLIER"
)

type ContractListListContractsResponseDataAmendmentsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                                       `json:"price,required"`
	RateType ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractListListContractsResponseDataAmendmentsOverridesOverwriteRateJSON
}

// contractListListContractsResponseDataAmendmentsOverridesOverwriteRateJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsOverridesOverwriteRate]
type contractListListContractsResponseDataAmendmentsOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateType string

const (
	ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateType = "FLAT"
	ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateType = "flat"
	ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractListListContractsResponseDataAmendmentsOverridesOverwriteRateRateType = "percentage"
)

type ContractListListContractsResponseDataAmendmentsOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataAmendmentsOverridesProductJSON
}

// contractListListContractsResponseDataAmendmentsOverridesProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsOverridesProduct]
type contractListListContractsResponseDataAmendmentsOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          ContractListListContractsResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                                       `json:"aws_account_number"`
	AwsOfferID            string                                                                       `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                                       `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                                    `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                                      `json:"fraction"`
	GcpAccountID          string                                                                       `json:"gcp_account_id"`
	GcpOfferID            string                                                                       `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                                       `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                                      `json:"reseller_contract_value"`
	StartingAt            time.Time                                                                    `json:"starting_at" format:"date-time"`
	JSON                  contractListListContractsResponseDataAmendmentsResellerRoyaltyJSON
}

// contractListListContractsResponseDataAmendmentsResellerRoyaltyJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsResellerRoyalty]
type contractListListContractsResponseDataAmendmentsResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	Fraction              apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerContractValue apijson.Field
	StartingAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	ContractListListContractsResponseDataAmendmentsResellerRoyaltiesResellerTypeAws ContractListListContractsResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractListListContractsResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp ContractListListContractsResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
)

type ContractListListContractsResponseDataAmendmentsScheduledCharge struct {
	ID       string                                                                  `json:"id,required" format:"uuid"`
	Product  ContractListListContractsResponseDataAmendmentsScheduledChargesProduct  `json:"product,required"`
	Schedule ContractListListContractsResponseDataAmendmentsScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractListListContractsResponseDataAmendmentsScheduledChargeJSON
}

// contractListListContractsResponseDataAmendmentsScheduledChargeJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsScheduledCharge]
type contractListListContractsResponseDataAmendmentsScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataAmendmentsScheduledChargesProductJSON
}

// contractListListContractsResponseDataAmendmentsScheduledChargesProductJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsScheduledChargesProduct]
type contractListListContractsResponseDataAmendmentsScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsScheduledChargesSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataAmendmentsScheduledChargesScheduleJSON
}

// contractListListContractsResponseDataAmendmentsScheduledChargesScheduleJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsScheduledChargesSchedule]
type contractListListContractsResponseDataAmendmentsScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
}

// contractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItem]
type contractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataAmendmentsScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrent struct {
	Commits                 []ContractListListContractsResponseDataCurrentCommit             `json:"commits,required"`
	CreatedAt               time.Time                                                        `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                                           `json:"created_by,required"`
	Discounts               []ContractListListContractsResponseDataCurrentDiscount           `json:"discounts,required"`
	Overrides               []ContractListListContractsResponseDataCurrentOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractListListContractsResponseDataCurrentResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractListListContractsResponseDataCurrentScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                                        `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractListListContractsResponseDataCurrentTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractListListContractsResponseDataCurrentUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                                        `json:"ending_before" format:"date-time"`
	Name                    string                                                           `json:"name"`
	NetPaymentTermsDays     float64                                                          `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                                           `json:"netsuite_sales_order_id"`
	RateCardID              string                                                           `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                                           `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                                          `json:"total_contract_value"`
	UsageFilter             ContractListListContractsResponseDataCurrentUsageFilter          `json:"usage_filter"`
	JSON                    contractListListContractsResponseDataCurrentJSON
}

// contractListListContractsResponseDataCurrentJSON contains the JSON metadata for
// the struct [ContractListListContractsResponseDataCurrent]
type contractListListContractsResponseDataCurrentJSON struct {
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Discounts               apijson.Field
	Overrides               apijson.Field
	ResellerRoyalties       apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Transitions             apijson.Field
	UsageInvoiceSchedule    apijson.Field
	EndingBefore            apijson.Field
	Name                    apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	RateCardID              apijson.Field
	SalesforceOpportunityID apijson.Field
	TotalContractValue      apijson.Field
	UsageFilter             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentCommit struct {
	ID      string                                                     `json:"id,required" format:"uuid"`
	Product ContractListListContractsResponseDataCurrentCommitsProduct `json:"product,required"`
	Type    ContractListListContractsResponseDataCurrentCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractListListContractsResponseDataCurrentCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractListListContractsResponseDataCurrentCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractListListContractsResponseDataCurrentCommitsLedger       `json:"ledger"`
	Name                 string                                                            `json:"name"`
	NetsuiteSalesOrderID string                                                            `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractListListContractsResponseDataCurrentCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                           `json:"rollover_fraction"`
	JSON                 contractListListContractsResponseDataCurrentCommitJSON
}

// contractListListContractsResponseDataCurrentCommitJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataCurrentCommit]
type contractListListContractsResponseDataCurrentCommitJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Type                 apijson.Field
	AccessSchedule       apijson.Field
	Amount               apijson.Field
	ApplicableProductIDs apijson.Field
	ApplicableTags       apijson.Field
	Description          apijson.Field
	InvoiceSchedule      apijson.Field
	Ledger               apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	RolledOverFrom       apijson.Field
	RolloverFraction     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataCurrentCommitsProductJSON
}

// contractListListContractsResponseDataCurrentCommitsProductJSON contains the JSON
// metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsProduct]
type contractListListContractsResponseDataCurrentCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentCommitsType string

const (
	ContractListListContractsResponseDataCurrentCommitsTypePrepaid  ContractListListContractsResponseDataCurrentCommitsType = "PREPAID"
	ContractListListContractsResponseDataCurrentCommitsTypePostpaid ContractListListContractsResponseDataCurrentCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractListListContractsResponseDataCurrentCommitsAccessSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractListListContractsResponseDataCurrentCommitsAccessScheduleJSON
}

// contractListListContractsResponseDataCurrentCommitsAccessScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsAccessSchedule]
type contractListListContractsResponseDataCurrentCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItemJSON
}

// contractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItem]
type contractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractListListContractsResponseDataCurrentCommitsInvoiceSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataCurrentCommitsInvoiceScheduleJSON
}

// contractListListContractsResponseDataCurrentCommitsInvoiceScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsInvoiceSchedule]
type contractListListContractsResponseDataCurrentCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON
}

// contractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItem]
type contractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractListListContractsResponseDataCurrentCommitsLedger interface {
	implementsContractListListContractsResponseDataCurrentCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractListListContractsResponseDataCurrentCommitsLedger)(nil)).Elem(), "")
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	SegmentID string                                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                        `json:"amount,required"`
	InvoiceID string                                                                                                         `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                                       `json:"amount,required"`
	NewContractID string                                                                                        `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type          ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                                         `json:"amount,required"`
	SegmentID string                                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractListListContractsResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                              `json:"amount,required"`
	Timestamp time.Time                                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                         `json:"amount,required"`
	InvoiceID string                                                                                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                       `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractListListContractsResponseDataCurrentCommitsLedger() {
}

type ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractListListContractsResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractListListContractsResponseDataCurrentCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractListListContractsResponseDataCurrentCommitsRolledOverFromJSON
}

// contractListListContractsResponseDataCurrentCommitsRolledOverFromJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentCommitsRolledOverFrom]
type contractListListContractsResponseDataCurrentCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentDiscount struct {
	ID                   string                                                        `json:"id,required" format:"uuid"`
	Product              ContractListListContractsResponseDataCurrentDiscountsProduct  `json:"product,required"`
	Schedule             ContractListListContractsResponseDataCurrentDiscountsSchedule `json:"schedule,required"`
	Name                 string                                                        `json:"name"`
	NetsuiteSalesOrderID string                                                        `json:"netsuite_sales_order_id"`
	JSON                 contractListListContractsResponseDataCurrentDiscountJSON
}

// contractListListContractsResponseDataCurrentDiscountJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataCurrentDiscount]
type contractListListContractsResponseDataCurrentDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataCurrentDiscountsProductJSON
}

// contractListListContractsResponseDataCurrentDiscountsProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentDiscountsProduct]
type contractListListContractsResponseDataCurrentDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentDiscountsSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataCurrentDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataCurrentDiscountsScheduleJSON
}

// contractListListContractsResponseDataCurrentDiscountsScheduleJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentDiscountsSchedule]
type contractListListContractsResponseDataCurrentDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataCurrentDiscountsScheduleScheduleItemJSON
}

// contractListListContractsResponseDataCurrentDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentDiscountsScheduleScheduleItem]
type contractListListContractsResponseDataCurrentDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentOverride struct {
	ID            string                                                             `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                          `json:"starting_at,required" format:"date-time"`
	Type          ContractListListContractsResponseDataCurrentOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                          `json:"ending_before" format:"date-time"`
	Entitled      bool                                                               `json:"entitled"`
	Multiplier    float64                                                            `json:"multiplier"`
	OverwriteRate ContractListListContractsResponseDataCurrentOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractListListContractsResponseDataCurrentOverridesProduct       `json:"product"`
	Tags          []string                                                           `json:"tags"`
	JSON          contractListListContractsResponseDataCurrentOverrideJSON
}

// contractListListContractsResponseDataCurrentOverrideJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataCurrentOverride]
type contractListListContractsResponseDataCurrentOverrideJSON struct {
	ID            apijson.Field
	StartingAt    apijson.Field
	Type          apijson.Field
	EndingBefore  apijson.Field
	Entitled      apijson.Field
	Multiplier    apijson.Field
	OverwriteRate apijson.Field
	Product       apijson.Field
	Tags          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentOverridesType string

const (
	ContractListListContractsResponseDataCurrentOverridesTypeOverwrite  ContractListListContractsResponseDataCurrentOverridesType = "OVERWRITE"
	ContractListListContractsResponseDataCurrentOverridesTypeMultiplier ContractListListContractsResponseDataCurrentOverridesType = "MULTIPLIER"
)

type ContractListListContractsResponseDataCurrentOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                                    `json:"price,required"`
	RateType ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractListListContractsResponseDataCurrentOverridesOverwriteRateJSON
}

// contractListListContractsResponseDataCurrentOverridesOverwriteRateJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentOverridesOverwriteRate]
type contractListListContractsResponseDataCurrentOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateType string

const (
	ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateType = "FLAT"
	ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateType = "flat"
	ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractListListContractsResponseDataCurrentOverridesOverwriteRateRateType = "percentage"
)

type ContractListListContractsResponseDataCurrentOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataCurrentOverridesProductJSON
}

// contractListListContractsResponseDataCurrentOverridesProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentOverridesProduct]
type contractListListContractsResponseDataCurrentOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentResellerRoyalty struct {
	Fraction              float64                                                                   `json:"fraction,required"`
	NetsuiteResellerID    string                                                                    `json:"netsuite_reseller_id,required"`
	ResellerType          ContractListListContractsResponseDataCurrentResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                                 `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                                    `json:"aws_account_number"`
	AwsOfferID            string                                                                    `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                                    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                                 `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                                    `json:"gcp_account_id"`
	GcpOfferID            string                                                                    `json:"gcp_offer_id"`
	ResellerContractValue float64                                                                   `json:"reseller_contract_value"`
	JSON                  contractListListContractsResponseDataCurrentResellerRoyaltyJSON
}

// contractListListContractsResponseDataCurrentResellerRoyaltyJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentResellerRoyalty]
type contractListListContractsResponseDataCurrentResellerRoyaltyJSON struct {
	Fraction              apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerType          apijson.Field
	StartingAt            apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	ResellerContractValue apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentResellerRoyaltiesResellerType string

const (
	ContractListListContractsResponseDataCurrentResellerRoyaltiesResellerTypeAws ContractListListContractsResponseDataCurrentResellerRoyaltiesResellerType = "AWS"
	ContractListListContractsResponseDataCurrentResellerRoyaltiesResellerTypeGcp ContractListListContractsResponseDataCurrentResellerRoyaltiesResellerType = "GCP"
)

type ContractListListContractsResponseDataCurrentScheduledCharge struct {
	ID       string                                                               `json:"id,required" format:"uuid"`
	Product  ContractListListContractsResponseDataCurrentScheduledChargesProduct  `json:"product,required"`
	Schedule ContractListListContractsResponseDataCurrentScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractListListContractsResponseDataCurrentScheduledChargeJSON
}

// contractListListContractsResponseDataCurrentScheduledChargeJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentScheduledCharge]
type contractListListContractsResponseDataCurrentScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataCurrentScheduledChargesProductJSON
}

// contractListListContractsResponseDataCurrentScheduledChargesProductJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentScheduledChargesProduct]
type contractListListContractsResponseDataCurrentScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentScheduledChargesSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataCurrentScheduledChargesScheduleJSON
}

// contractListListContractsResponseDataCurrentScheduledChargesScheduleJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentScheduledChargesSchedule]
type contractListListContractsResponseDataCurrentScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItemJSON
}

// contractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItem]
type contractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentTransition struct {
	FromContractID string                                                      `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                                      `json:"to_contract_id,required" format:"uuid"`
	Type           ContractListListContractsResponseDataCurrentTransitionsType `json:"type,required"`
	JSON           contractListListContractsResponseDataCurrentTransitionJSON
}

// contractListListContractsResponseDataCurrentTransitionJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataCurrentTransition]
type contractListListContractsResponseDataCurrentTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentTransitionsType string

const (
	ContractListListContractsResponseDataCurrentTransitionsTypeSupersede ContractListListContractsResponseDataCurrentTransitionsType = "SUPERSEDE"
	ContractListListContractsResponseDataCurrentTransitionsTypeRenewal   ContractListListContractsResponseDataCurrentTransitionsType = "RENEWAL"
)

type ContractListListContractsResponseDataCurrentUsageInvoiceSchedule struct {
	Frequency ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractListListContractsResponseDataCurrentUsageInvoiceScheduleJSON
}

// contractListListContractsResponseDataCurrentUsageInvoiceScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentUsageInvoiceSchedule]
type contractListListContractsResponseDataCurrentUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequency string

const (
	ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequency = "monthly"
	ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractListListContractsResponseDataCurrentUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractListListContractsResponseDataCurrentUsageFilter struct {
	Current ContractListListContractsResponseDataCurrentUsageFilterCurrent  `json:"current,required"`
	Initial ContractListListContractsResponseDataCurrentUsageFilterInitial  `json:"initial,required"`
	Updates []ContractListListContractsResponseDataCurrentUsageFilterUpdate `json:"updates,required"`
	JSON    contractListListContractsResponseDataCurrentUsageFilterJSON
}

// contractListListContractsResponseDataCurrentUsageFilterJSON contains the JSON
// metadata for the struct
// [ContractListListContractsResponseDataCurrentUsageFilter]
type contractListListContractsResponseDataCurrentUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListListContractsResponseDataCurrentUsageFilterCurrentJSON
}

// contractListListContractsResponseDataCurrentUsageFilterCurrentJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentUsageFilterCurrent]
type contractListListContractsResponseDataCurrentUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListListContractsResponseDataCurrentUsageFilterInitialJSON
}

// contractListListContractsResponseDataCurrentUsageFilterInitialJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentUsageFilterInitial]
type contractListListContractsResponseDataCurrentUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataCurrentUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractListListContractsResponseDataCurrentUsageFilterUpdateJSON
}

// contractListListContractsResponseDataCurrentUsageFilterUpdateJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataCurrentUsageFilterUpdate]
type contractListListContractsResponseDataCurrentUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataCurrentUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitial struct {
	Commits                 []ContractListListContractsResponseDataInitialCommit             `json:"commits,required"`
	CreatedAt               time.Time                                                        `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                                           `json:"created_by,required"`
	Discounts               []ContractListListContractsResponseDataInitialDiscount           `json:"discounts,required"`
	Overrides               []ContractListListContractsResponseDataInitialOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractListListContractsResponseDataInitialResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractListListContractsResponseDataInitialScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                                        `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractListListContractsResponseDataInitialTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractListListContractsResponseDataInitialUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                                        `json:"ending_before" format:"date-time"`
	Name                    string                                                           `json:"name"`
	NetPaymentTermsDays     float64                                                          `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                                           `json:"netsuite_sales_order_id"`
	RateCardID              string                                                           `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                                           `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                                          `json:"total_contract_value"`
	UsageFilter             ContractListListContractsResponseDataInitialUsageFilter          `json:"usage_filter"`
	JSON                    contractListListContractsResponseDataInitialJSON
}

// contractListListContractsResponseDataInitialJSON contains the JSON metadata for
// the struct [ContractListListContractsResponseDataInitial]
type contractListListContractsResponseDataInitialJSON struct {
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Discounts               apijson.Field
	Overrides               apijson.Field
	ResellerRoyalties       apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Transitions             apijson.Field
	UsageInvoiceSchedule    apijson.Field
	EndingBefore            apijson.Field
	Name                    apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	RateCardID              apijson.Field
	SalesforceOpportunityID apijson.Field
	TotalContractValue      apijson.Field
	UsageFilter             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialCommit struct {
	ID      string                                                     `json:"id,required" format:"uuid"`
	Product ContractListListContractsResponseDataInitialCommitsProduct `json:"product,required"`
	Type    ContractListListContractsResponseDataInitialCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractListListContractsResponseDataInitialCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractListListContractsResponseDataInitialCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractListListContractsResponseDataInitialCommitsLedger       `json:"ledger"`
	Name                 string                                                            `json:"name"`
	NetsuiteSalesOrderID string                                                            `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractListListContractsResponseDataInitialCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                           `json:"rollover_fraction"`
	JSON                 contractListListContractsResponseDataInitialCommitJSON
}

// contractListListContractsResponseDataInitialCommitJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataInitialCommit]
type contractListListContractsResponseDataInitialCommitJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Type                 apijson.Field
	AccessSchedule       apijson.Field
	Amount               apijson.Field
	ApplicableProductIDs apijson.Field
	ApplicableTags       apijson.Field
	Description          apijson.Field
	InvoiceSchedule      apijson.Field
	Ledger               apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	RolledOverFrom       apijson.Field
	RolloverFraction     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataInitialCommitsProductJSON
}

// contractListListContractsResponseDataInitialCommitsProductJSON contains the JSON
// metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsProduct]
type contractListListContractsResponseDataInitialCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialCommitsType string

const (
	ContractListListContractsResponseDataInitialCommitsTypePrepaid  ContractListListContractsResponseDataInitialCommitsType = "PREPAID"
	ContractListListContractsResponseDataInitialCommitsTypePostpaid ContractListListContractsResponseDataInitialCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractListListContractsResponseDataInitialCommitsAccessSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractListListContractsResponseDataInitialCommitsAccessScheduleJSON
}

// contractListListContractsResponseDataInitialCommitsAccessScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsAccessSchedule]
type contractListListContractsResponseDataInitialCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItemJSON
}

// contractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItem]
type contractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractListListContractsResponseDataInitialCommitsInvoiceSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataInitialCommitsInvoiceScheduleJSON
}

// contractListListContractsResponseDataInitialCommitsInvoiceScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsInvoiceSchedule]
type contractListListContractsResponseDataInitialCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON
}

// contractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItem]
type contractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractListListContractsResponseDataInitialCommitsLedger interface {
	implementsContractListListContractsResponseDataInitialCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractListListContractsResponseDataInitialCommitsLedger)(nil)).Elem(), "")
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	SegmentID string                                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                        `json:"amount,required"`
	InvoiceID string                                                                                                         `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                                       `json:"amount,required"`
	NewContractID string                                                                                        `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type          ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                                         `json:"amount,required"`
	SegmentID string                                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractListListContractsResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                              `json:"amount,required"`
	Timestamp time.Time                                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                         `json:"amount,required"`
	InvoiceID string                                                                                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                       `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractListListContractsResponseDataInitialCommitsLedger() {
}

type ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractListListContractsResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractListListContractsResponseDataInitialCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractListListContractsResponseDataInitialCommitsRolledOverFromJSON
}

// contractListListContractsResponseDataInitialCommitsRolledOverFromJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialCommitsRolledOverFrom]
type contractListListContractsResponseDataInitialCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialDiscount struct {
	ID                   string                                                        `json:"id,required" format:"uuid"`
	Product              ContractListListContractsResponseDataInitialDiscountsProduct  `json:"product,required"`
	Schedule             ContractListListContractsResponseDataInitialDiscountsSchedule `json:"schedule,required"`
	Name                 string                                                        `json:"name"`
	NetsuiteSalesOrderID string                                                        `json:"netsuite_sales_order_id"`
	JSON                 contractListListContractsResponseDataInitialDiscountJSON
}

// contractListListContractsResponseDataInitialDiscountJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataInitialDiscount]
type contractListListContractsResponseDataInitialDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataInitialDiscountsProductJSON
}

// contractListListContractsResponseDataInitialDiscountsProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialDiscountsProduct]
type contractListListContractsResponseDataInitialDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialDiscountsSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataInitialDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataInitialDiscountsScheduleJSON
}

// contractListListContractsResponseDataInitialDiscountsScheduleJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialDiscountsSchedule]
type contractListListContractsResponseDataInitialDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataInitialDiscountsScheduleScheduleItemJSON
}

// contractListListContractsResponseDataInitialDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialDiscountsScheduleScheduleItem]
type contractListListContractsResponseDataInitialDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialOverride struct {
	ID            string                                                             `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                          `json:"starting_at,required" format:"date-time"`
	Type          ContractListListContractsResponseDataInitialOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                          `json:"ending_before" format:"date-time"`
	Entitled      bool                                                               `json:"entitled"`
	Multiplier    float64                                                            `json:"multiplier"`
	OverwriteRate ContractListListContractsResponseDataInitialOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractListListContractsResponseDataInitialOverridesProduct       `json:"product"`
	Tags          []string                                                           `json:"tags"`
	JSON          contractListListContractsResponseDataInitialOverrideJSON
}

// contractListListContractsResponseDataInitialOverrideJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataInitialOverride]
type contractListListContractsResponseDataInitialOverrideJSON struct {
	ID            apijson.Field
	StartingAt    apijson.Field
	Type          apijson.Field
	EndingBefore  apijson.Field
	Entitled      apijson.Field
	Multiplier    apijson.Field
	OverwriteRate apijson.Field
	Product       apijson.Field
	Tags          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialOverridesType string

const (
	ContractListListContractsResponseDataInitialOverridesTypeOverwrite  ContractListListContractsResponseDataInitialOverridesType = "OVERWRITE"
	ContractListListContractsResponseDataInitialOverridesTypeMultiplier ContractListListContractsResponseDataInitialOverridesType = "MULTIPLIER"
)

type ContractListListContractsResponseDataInitialOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                                    `json:"price,required"`
	RateType ContractListListContractsResponseDataInitialOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractListListContractsResponseDataInitialOverridesOverwriteRateJSON
}

// contractListListContractsResponseDataInitialOverridesOverwriteRateJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialOverridesOverwriteRate]
type contractListListContractsResponseDataInitialOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialOverridesOverwriteRateRateType string

const (
	ContractListListContractsResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractListListContractsResponseDataInitialOverridesOverwriteRateRateType = "FLAT"
	ContractListListContractsResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractListListContractsResponseDataInitialOverridesOverwriteRateRateType = "flat"
	ContractListListContractsResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractListListContractsResponseDataInitialOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractListListContractsResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractListListContractsResponseDataInitialOverridesOverwriteRateRateType = "percentage"
)

type ContractListListContractsResponseDataInitialOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataInitialOverridesProductJSON
}

// contractListListContractsResponseDataInitialOverridesProductJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialOverridesProduct]
type contractListListContractsResponseDataInitialOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialResellerRoyalty struct {
	Fraction              float64                                                                   `json:"fraction,required"`
	NetsuiteResellerID    string                                                                    `json:"netsuite_reseller_id,required"`
	ResellerType          ContractListListContractsResponseDataInitialResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                                 `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                                    `json:"aws_account_number"`
	AwsOfferID            string                                                                    `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                                    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                                 `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                                    `json:"gcp_account_id"`
	GcpOfferID            string                                                                    `json:"gcp_offer_id"`
	ResellerContractValue float64                                                                   `json:"reseller_contract_value"`
	JSON                  contractListListContractsResponseDataInitialResellerRoyaltyJSON
}

// contractListListContractsResponseDataInitialResellerRoyaltyJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialResellerRoyalty]
type contractListListContractsResponseDataInitialResellerRoyaltyJSON struct {
	Fraction              apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerType          apijson.Field
	StartingAt            apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	ResellerContractValue apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialResellerRoyaltiesResellerType string

const (
	ContractListListContractsResponseDataInitialResellerRoyaltiesResellerTypeAws ContractListListContractsResponseDataInitialResellerRoyaltiesResellerType = "AWS"
	ContractListListContractsResponseDataInitialResellerRoyaltiesResellerTypeGcp ContractListListContractsResponseDataInitialResellerRoyaltiesResellerType = "GCP"
)

type ContractListListContractsResponseDataInitialScheduledCharge struct {
	ID       string                                                               `json:"id,required" format:"uuid"`
	Product  ContractListListContractsResponseDataInitialScheduledChargesProduct  `json:"product,required"`
	Schedule ContractListListContractsResponseDataInitialScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractListListContractsResponseDataInitialScheduledChargeJSON
}

// contractListListContractsResponseDataInitialScheduledChargeJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialScheduledCharge]
type contractListListContractsResponseDataInitialScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListListContractsResponseDataInitialScheduledChargesProductJSON
}

// contractListListContractsResponseDataInitialScheduledChargesProductJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialScheduledChargesProduct]
type contractListListContractsResponseDataInitialScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialScheduledChargesSchedule struct {
	ScheduleItems []ContractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListListContractsResponseDataInitialScheduledChargesScheduleJSON
}

// contractListListContractsResponseDataInitialScheduledChargesScheduleJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialScheduledChargesSchedule]
type contractListListContractsResponseDataInitialScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItemJSON
}

// contractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItem]
type contractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialTransition struct {
	FromContractID string                                                      `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                                      `json:"to_contract_id,required" format:"uuid"`
	Type           ContractListListContractsResponseDataInitialTransitionsType `json:"type,required"`
	JSON           contractListListContractsResponseDataInitialTransitionJSON
}

// contractListListContractsResponseDataInitialTransitionJSON contains the JSON
// metadata for the struct [ContractListListContractsResponseDataInitialTransition]
type contractListListContractsResponseDataInitialTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialTransitionsType string

const (
	ContractListListContractsResponseDataInitialTransitionsTypeSupersede ContractListListContractsResponseDataInitialTransitionsType = "SUPERSEDE"
	ContractListListContractsResponseDataInitialTransitionsTypeRenewal   ContractListListContractsResponseDataInitialTransitionsType = "RENEWAL"
)

type ContractListListContractsResponseDataInitialUsageInvoiceSchedule struct {
	Frequency ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractListListContractsResponseDataInitialUsageInvoiceScheduleJSON
}

// contractListListContractsResponseDataInitialUsageInvoiceScheduleJSON contains
// the JSON metadata for the struct
// [ContractListListContractsResponseDataInitialUsageInvoiceSchedule]
type contractListListContractsResponseDataInitialUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequency string

const (
	ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequency = "monthly"
	ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractListListContractsResponseDataInitialUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractListListContractsResponseDataInitialUsageFilter struct {
	Current ContractListListContractsResponseDataInitialUsageFilterCurrent  `json:"current,required"`
	Initial ContractListListContractsResponseDataInitialUsageFilterInitial  `json:"initial,required"`
	Updates []ContractListListContractsResponseDataInitialUsageFilterUpdate `json:"updates,required"`
	JSON    contractListListContractsResponseDataInitialUsageFilterJSON
}

// contractListListContractsResponseDataInitialUsageFilterJSON contains the JSON
// metadata for the struct
// [ContractListListContractsResponseDataInitialUsageFilter]
type contractListListContractsResponseDataInitialUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListListContractsResponseDataInitialUsageFilterCurrentJSON
}

// contractListListContractsResponseDataInitialUsageFilterCurrentJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialUsageFilterCurrent]
type contractListListContractsResponseDataInitialUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListListContractsResponseDataInitialUsageFilterInitialJSON
}

// contractListListContractsResponseDataInitialUsageFilterInitialJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialUsageFilterInitial]
type contractListListContractsResponseDataInitialUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsResponseDataInitialUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractListListContractsResponseDataInitialUsageFilterUpdateJSON
}

// contractListListContractsResponseDataInitialUsageFilterUpdateJSON contains the
// JSON metadata for the struct
// [ContractListListContractsResponseDataInitialUsageFilterUpdate]
type contractListListContractsResponseDataInitialUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListListContractsResponseDataInitialUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListListContractsParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r ContractListListContractsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
