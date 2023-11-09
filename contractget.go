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

// ContractGetService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewContractGetService] method
// instead.
type ContractGetService struct {
	Options []option.RequestOption
}

// NewContractGetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractGetService(opts ...option.RequestOption) (r *ContractGetService) {
	r = &ContractGetService{}
	r.Options = opts
	return
}

// Get a specific contract
func (r *ContractGetService) GetContract(ctx context.Context, body ContractGetGetContractParams, opts ...option.RequestOption) (res *ContractGetGetContractResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContractGetGetContractResponse struct {
	Data ContractGetGetContractResponseData `json:"data,required"`
	JSON contractGetGetContractResponseJSON
}

// contractGetGetContractResponseJSON contains the JSON metadata for the struct
// [ContractGetGetContractResponse]
type contractGetGetContractResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseData struct {
	ID         string                                        `json:"id,required" format:"uuid"`
	Amendments []ContractGetGetContractResponseDataAmendment `json:"amendments,required"`
	Current    ContractGetGetContractResponseDataCurrent     `json:"current,required"`
	CustomerID string                                        `json:"customer_id,required" format:"uuid"`
	Initial    ContractGetGetContractResponseDataInitial     `json:"initial,required"`
	JSON       contractGetGetContractResponseDataJSON
}

// contractGetGetContractResponseDataJSON contains the JSON metadata for the struct
// [ContractGetGetContractResponseData]
type contractGetGetContractResponseDataJSON struct {
	ID          apijson.Field
	Amendments  apijson.Field
	Current     apijson.Field
	CustomerID  apijson.Field
	Initial     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendment struct {
	ID                      string                                                        `json:"id,required" format:"uuid"`
	Commits                 []ContractGetGetContractResponseDataAmendmentsCommit          `json:"commits,required"`
	CreatedAt               time.Time                                                     `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                                        `json:"created_by,required"`
	Discounts               []ContractGetGetContractResponseDataAmendmentsDiscount        `json:"discounts,required"`
	Overrides               []ContractGetGetContractResponseDataAmendmentsOverride        `json:"overrides,required"`
	ResellerRoyalties       []ContractGetGetContractResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractGetGetContractResponseDataAmendmentsScheduledCharge `json:"scheduled_charges,required"`
	StartingAt              time.Time                                                     `json:"starting_at,required" format:"date-time"`
	NetsuiteSalesOrderID    string                                                        `json:"netsuite_sales_order_id"`
	SalesforceOpportunityID string                                                        `json:"salesforce_opportunity_id"`
	JSON                    contractGetGetContractResponseDataAmendmentJSON
}

// contractGetGetContractResponseDataAmendmentJSON contains the JSON metadata for
// the struct [ContractGetGetContractResponseDataAmendment]
type contractGetGetContractResponseDataAmendmentJSON struct {
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

func (r *ContractGetGetContractResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsCommit struct {
	ID      string                                                     `json:"id,required" format:"uuid"`
	Product ContractGetGetContractResponseDataAmendmentsCommitsProduct `json:"product,required"`
	Type    ContractGetGetContractResponseDataAmendmentsCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractGetGetContractResponseDataAmendmentsCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractGetGetContractResponseDataAmendmentsCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractGetGetContractResponseDataAmendmentsCommitsLedger       `json:"ledger"`
	Name                 string                                                            `json:"name"`
	NetsuiteSalesOrderID string                                                            `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractGetGetContractResponseDataAmendmentsCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                           `json:"rollover_fraction"`
	JSON                 contractGetGetContractResponseDataAmendmentsCommitJSON
}

// contractGetGetContractResponseDataAmendmentsCommitJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataAmendmentsCommit]
type contractGetGetContractResponseDataAmendmentsCommitJSON struct {
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

func (r *ContractGetGetContractResponseDataAmendmentsCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataAmendmentsCommitsProductJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsProduct]
type contractGetGetContractResponseDataAmendmentsCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsCommitsType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsTypePrepaid  ContractGetGetContractResponseDataAmendmentsCommitsType = "PREPAID"
	ContractGetGetContractResponseDataAmendmentsCommitsTypePostpaid ContractGetGetContractResponseDataAmendmentsCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractGetGetContractResponseDataAmendmentsCommitsAccessSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractGetGetContractResponseDataAmendmentsCommitsAccessScheduleJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsAccessScheduleJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsAccessSchedule]
type contractGetGetContractResponseDataAmendmentsCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItem]
type contractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractGetGetContractResponseDataAmendmentsCommitsInvoiceSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsInvoiceSchedule]
type contractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem]
type contractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractGetGetContractResponseDataAmendmentsCommitsLedger interface {
	implementsContractGetGetContractResponseDataAmendmentsCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractGetGetContractResponseDataAmendmentsCommitsLedger)(nil)).Elem(), "")
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	SegmentID string                                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                        `json:"amount,required"`
	InvoiceID string                                                                                                         `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                                       `json:"amount,required"`
	NewContractID string                                                                                        `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type          ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                                         `json:"amount,required"`
	SegmentID string                                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                                       `json:"amount,required"`
	InvoiceID string                                                                                        `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                     `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractGetGetContractResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                              `json:"amount,required"`
	Timestamp time.Time                                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                         `json:"amount,required"`
	InvoiceID string                                                                                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                       `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractGetGetContractResponseDataAmendmentsCommitsLedger() {
}

type ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractGetGetContractResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractGetGetContractResponseDataAmendmentsCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractGetGetContractResponseDataAmendmentsCommitsRolledOverFromJSON
}

// contractGetGetContractResponseDataAmendmentsCommitsRolledOverFromJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsCommitsRolledOverFrom]
type contractGetGetContractResponseDataAmendmentsCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsDiscount struct {
	ID                   string                                                        `json:"id,required" format:"uuid"`
	Product              ContractGetGetContractResponseDataAmendmentsDiscountsProduct  `json:"product,required"`
	Schedule             ContractGetGetContractResponseDataAmendmentsDiscountsSchedule `json:"schedule,required"`
	Name                 string                                                        `json:"name"`
	NetsuiteSalesOrderID string                                                        `json:"netsuite_sales_order_id"`
	JSON                 contractGetGetContractResponseDataAmendmentsDiscountJSON
}

// contractGetGetContractResponseDataAmendmentsDiscountJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataAmendmentsDiscount]
type contractGetGetContractResponseDataAmendmentsDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataAmendmentsDiscountsProductJSON
}

// contractGetGetContractResponseDataAmendmentsDiscountsProductJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsDiscountsProduct]
type contractGetGetContractResponseDataAmendmentsDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsDiscountsSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataAmendmentsDiscountsScheduleJSON
}

// contractGetGetContractResponseDataAmendmentsDiscountsScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsDiscountsSchedule]
type contractGetGetContractResponseDataAmendmentsDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItem]
type contractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsOverride struct {
	ID            string                                                             `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                          `json:"starting_at,required" format:"date-time"`
	Type          ContractGetGetContractResponseDataAmendmentsOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                          `json:"ending_before" format:"date-time"`
	Entitled      bool                                                               `json:"entitled"`
	Multiplier    float64                                                            `json:"multiplier"`
	OverwriteRate ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractGetGetContractResponseDataAmendmentsOverridesProduct       `json:"product"`
	Tags          []string                                                           `json:"tags"`
	JSON          contractGetGetContractResponseDataAmendmentsOverrideJSON
}

// contractGetGetContractResponseDataAmendmentsOverrideJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataAmendmentsOverride]
type contractGetGetContractResponseDataAmendmentsOverrideJSON struct {
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

func (r *ContractGetGetContractResponseDataAmendmentsOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsOverridesType string

const (
	ContractGetGetContractResponseDataAmendmentsOverridesTypeOverwrite  ContractGetGetContractResponseDataAmendmentsOverridesType = "OVERWRITE"
	ContractGetGetContractResponseDataAmendmentsOverridesTypeMultiplier ContractGetGetContractResponseDataAmendmentsOverridesType = "MULTIPLIER"
)

type ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                                    `json:"price,required"`
	RateType ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractGetGetContractResponseDataAmendmentsOverridesOverwriteRateJSON
}

// contractGetGetContractResponseDataAmendmentsOverridesOverwriteRateJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRate]
type contractGetGetContractResponseDataAmendmentsOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateType string

const (
	ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateType = "FLAT"
	ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateType = "flat"
	ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractGetGetContractResponseDataAmendmentsOverridesOverwriteRateRateType = "percentage"
)

type ContractGetGetContractResponseDataAmendmentsOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataAmendmentsOverridesProductJSON
}

// contractGetGetContractResponseDataAmendmentsOverridesProductJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsOverridesProduct]
type contractGetGetContractResponseDataAmendmentsOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          ContractGetGetContractResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                                    `json:"aws_account_number"`
	AwsOfferID            string                                                                    `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                                    `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                                 `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                                   `json:"fraction"`
	GcpAccountID          string                                                                    `json:"gcp_account_id"`
	GcpOfferID            string                                                                    `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                                    `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                                   `json:"reseller_contract_value"`
	StartingAt            time.Time                                                                 `json:"starting_at" format:"date-time"`
	JSON                  contractGetGetContractResponseDataAmendmentsResellerRoyaltyJSON
}

// contractGetGetContractResponseDataAmendmentsResellerRoyaltyJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsResellerRoyalty]
type contractGetGetContractResponseDataAmendmentsResellerRoyaltyJSON struct {
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

func (r *ContractGetGetContractResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	ContractGetGetContractResponseDataAmendmentsResellerRoyaltiesResellerTypeAws ContractGetGetContractResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractGetGetContractResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp ContractGetGetContractResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
)

type ContractGetGetContractResponseDataAmendmentsScheduledCharge struct {
	ID       string                                                               `json:"id,required" format:"uuid"`
	Product  ContractGetGetContractResponseDataAmendmentsScheduledChargesProduct  `json:"product,required"`
	Schedule ContractGetGetContractResponseDataAmendmentsScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractGetGetContractResponseDataAmendmentsScheduledChargeJSON
}

// contractGetGetContractResponseDataAmendmentsScheduledChargeJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsScheduledCharge]
type contractGetGetContractResponseDataAmendmentsScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataAmendmentsScheduledChargesProductJSON
}

// contractGetGetContractResponseDataAmendmentsScheduledChargesProductJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsScheduledChargesProduct]
type contractGetGetContractResponseDataAmendmentsScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsScheduledChargesSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataAmendmentsScheduledChargesScheduleJSON
}

// contractGetGetContractResponseDataAmendmentsScheduledChargesScheduleJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsScheduledChargesSchedule]
type contractGetGetContractResponseDataAmendmentsScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItem]
type contractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataAmendmentsScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrent struct {
	Commits                 []ContractGetGetContractResponseDataCurrentCommit             `json:"commits,required"`
	CreatedAt               time.Time                                                     `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                                        `json:"created_by,required"`
	Discounts               []ContractGetGetContractResponseDataCurrentDiscount           `json:"discounts,required"`
	Overrides               []ContractGetGetContractResponseDataCurrentOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractGetGetContractResponseDataCurrentResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractGetGetContractResponseDataCurrentScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                                     `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractGetGetContractResponseDataCurrentTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractGetGetContractResponseDataCurrentUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                                     `json:"ending_before" format:"date-time"`
	Name                    string                                                        `json:"name"`
	NetPaymentTermsDays     float64                                                       `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                                        `json:"netsuite_sales_order_id"`
	RateCardID              string                                                        `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                                        `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                                       `json:"total_contract_value"`
	UsageFilter             ContractGetGetContractResponseDataCurrentUsageFilter          `json:"usage_filter"`
	JSON                    contractGetGetContractResponseDataCurrentJSON
}

// contractGetGetContractResponseDataCurrentJSON contains the JSON metadata for the
// struct [ContractGetGetContractResponseDataCurrent]
type contractGetGetContractResponseDataCurrentJSON struct {
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

func (r *ContractGetGetContractResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentCommit struct {
	ID      string                                                  `json:"id,required" format:"uuid"`
	Product ContractGetGetContractResponseDataCurrentCommitsProduct `json:"product,required"`
	Type    ContractGetGetContractResponseDataCurrentCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractGetGetContractResponseDataCurrentCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractGetGetContractResponseDataCurrentCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractGetGetContractResponseDataCurrentCommitsLedger       `json:"ledger"`
	Name                 string                                                         `json:"name"`
	NetsuiteSalesOrderID string                                                         `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractGetGetContractResponseDataCurrentCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                        `json:"rollover_fraction"`
	JSON                 contractGetGetContractResponseDataCurrentCommitJSON
}

// contractGetGetContractResponseDataCurrentCommitJSON contains the JSON metadata
// for the struct [ContractGetGetContractResponseDataCurrentCommit]
type contractGetGetContractResponseDataCurrentCommitJSON struct {
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

func (r *ContractGetGetContractResponseDataCurrentCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataCurrentCommitsProductJSON
}

// contractGetGetContractResponseDataCurrentCommitsProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsProduct]
type contractGetGetContractResponseDataCurrentCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentCommitsType string

const (
	ContractGetGetContractResponseDataCurrentCommitsTypePrepaid  ContractGetGetContractResponseDataCurrentCommitsType = "PREPAID"
	ContractGetGetContractResponseDataCurrentCommitsTypePostpaid ContractGetGetContractResponseDataCurrentCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractGetGetContractResponseDataCurrentCommitsAccessSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractGetGetContractResponseDataCurrentCommitsAccessScheduleJSON
}

// contractGetGetContractResponseDataCurrentCommitsAccessScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsAccessSchedule]
type contractGetGetContractResponseDataCurrentCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItem]
type contractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractGetGetContractResponseDataCurrentCommitsInvoiceSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataCurrentCommitsInvoiceScheduleJSON
}

// contractGetGetContractResponseDataCurrentCommitsInvoiceScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsInvoiceSchedule]
type contractGetGetContractResponseDataCurrentCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItem]
type contractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractGetGetContractResponseDataCurrentCommitsLedger interface {
	implementsContractGetGetContractResponseDataCurrentCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractGetGetContractResponseDataCurrentCommitsLedger)(nil)).Elem(), "")
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                        `json:"amount,required"`
	SegmentID string                                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                     `json:"amount,required"`
	InvoiceID string                                                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                   `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                                    `json:"amount,required"`
	NewContractID string                                                                                     `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type          ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	SegmentID string                                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                                    `json:"amount,required"`
	InvoiceID string                                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                                    `json:"amount,required"`
	InvoiceID string                                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractGetGetContractResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                                       `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                                   `json:"amount,required"`
	InvoiceID string                                                                                    `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractGetGetContractResponseDataCurrentCommitsLedger() {
}

type ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractGetGetContractResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractGetGetContractResponseDataCurrentCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractGetGetContractResponseDataCurrentCommitsRolledOverFromJSON
}

// contractGetGetContractResponseDataCurrentCommitsRolledOverFromJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentCommitsRolledOverFrom]
type contractGetGetContractResponseDataCurrentCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentDiscount struct {
	ID                   string                                                     `json:"id,required" format:"uuid"`
	Product              ContractGetGetContractResponseDataCurrentDiscountsProduct  `json:"product,required"`
	Schedule             ContractGetGetContractResponseDataCurrentDiscountsSchedule `json:"schedule,required"`
	Name                 string                                                     `json:"name"`
	NetsuiteSalesOrderID string                                                     `json:"netsuite_sales_order_id"`
	JSON                 contractGetGetContractResponseDataCurrentDiscountJSON
}

// contractGetGetContractResponseDataCurrentDiscountJSON contains the JSON metadata
// for the struct [ContractGetGetContractResponseDataCurrentDiscount]
type contractGetGetContractResponseDataCurrentDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataCurrentDiscountsProductJSON
}

// contractGetGetContractResponseDataCurrentDiscountsProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentDiscountsProduct]
type contractGetGetContractResponseDataCurrentDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentDiscountsSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataCurrentDiscountsScheduleJSON
}

// contractGetGetContractResponseDataCurrentDiscountsScheduleJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentDiscountsSchedule]
type contractGetGetContractResponseDataCurrentDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItem]
type contractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentOverride struct {
	ID            string                                                          `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                       `json:"starting_at,required" format:"date-time"`
	Type          ContractGetGetContractResponseDataCurrentOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                       `json:"ending_before" format:"date-time"`
	Entitled      bool                                                            `json:"entitled"`
	Multiplier    float64                                                         `json:"multiplier"`
	OverwriteRate ContractGetGetContractResponseDataCurrentOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractGetGetContractResponseDataCurrentOverridesProduct       `json:"product"`
	Tags          []string                                                        `json:"tags"`
	JSON          contractGetGetContractResponseDataCurrentOverrideJSON
}

// contractGetGetContractResponseDataCurrentOverrideJSON contains the JSON metadata
// for the struct [ContractGetGetContractResponseDataCurrentOverride]
type contractGetGetContractResponseDataCurrentOverrideJSON struct {
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

func (r *ContractGetGetContractResponseDataCurrentOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentOverridesType string

const (
	ContractGetGetContractResponseDataCurrentOverridesTypeOverwrite  ContractGetGetContractResponseDataCurrentOverridesType = "OVERWRITE"
	ContractGetGetContractResponseDataCurrentOverridesTypeMultiplier ContractGetGetContractResponseDataCurrentOverridesType = "MULTIPLIER"
)

type ContractGetGetContractResponseDataCurrentOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                                 `json:"price,required"`
	RateType ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractGetGetContractResponseDataCurrentOverridesOverwriteRateJSON
}

// contractGetGetContractResponseDataCurrentOverridesOverwriteRateJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentOverridesOverwriteRate]
type contractGetGetContractResponseDataCurrentOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateType string

const (
	ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateType = "FLAT"
	ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateType = "flat"
	ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractGetGetContractResponseDataCurrentOverridesOverwriteRateRateType = "percentage"
)

type ContractGetGetContractResponseDataCurrentOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataCurrentOverridesProductJSON
}

// contractGetGetContractResponseDataCurrentOverridesProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentOverridesProduct]
type contractGetGetContractResponseDataCurrentOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentResellerRoyalty struct {
	Fraction              float64                                                                `json:"fraction,required"`
	NetsuiteResellerID    string                                                                 `json:"netsuite_reseller_id,required"`
	ResellerType          ContractGetGetContractResponseDataCurrentResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                              `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                                 `json:"aws_account_number"`
	AwsOfferID            string                                                                 `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                                 `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                              `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                                 `json:"gcp_account_id"`
	GcpOfferID            string                                                                 `json:"gcp_offer_id"`
	ResellerContractValue float64                                                                `json:"reseller_contract_value"`
	JSON                  contractGetGetContractResponseDataCurrentResellerRoyaltyJSON
}

// contractGetGetContractResponseDataCurrentResellerRoyaltyJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentResellerRoyalty]
type contractGetGetContractResponseDataCurrentResellerRoyaltyJSON struct {
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

func (r *ContractGetGetContractResponseDataCurrentResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentResellerRoyaltiesResellerType string

const (
	ContractGetGetContractResponseDataCurrentResellerRoyaltiesResellerTypeAws ContractGetGetContractResponseDataCurrentResellerRoyaltiesResellerType = "AWS"
	ContractGetGetContractResponseDataCurrentResellerRoyaltiesResellerTypeGcp ContractGetGetContractResponseDataCurrentResellerRoyaltiesResellerType = "GCP"
)

type ContractGetGetContractResponseDataCurrentScheduledCharge struct {
	ID       string                                                            `json:"id,required" format:"uuid"`
	Product  ContractGetGetContractResponseDataCurrentScheduledChargesProduct  `json:"product,required"`
	Schedule ContractGetGetContractResponseDataCurrentScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractGetGetContractResponseDataCurrentScheduledChargeJSON
}

// contractGetGetContractResponseDataCurrentScheduledChargeJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentScheduledCharge]
type contractGetGetContractResponseDataCurrentScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataCurrentScheduledChargesProductJSON
}

// contractGetGetContractResponseDataCurrentScheduledChargesProductJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentScheduledChargesProduct]
type contractGetGetContractResponseDataCurrentScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentScheduledChargesSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataCurrentScheduledChargesScheduleJSON
}

// contractGetGetContractResponseDataCurrentScheduledChargesScheduleJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentScheduledChargesSchedule]
type contractGetGetContractResponseDataCurrentScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItem]
type contractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentTransition struct {
	FromContractID string                                                   `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                                   `json:"to_contract_id,required" format:"uuid"`
	Type           ContractGetGetContractResponseDataCurrentTransitionsType `json:"type,required"`
	JSON           contractGetGetContractResponseDataCurrentTransitionJSON
}

// contractGetGetContractResponseDataCurrentTransitionJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataCurrentTransition]
type contractGetGetContractResponseDataCurrentTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentTransitionsType string

const (
	ContractGetGetContractResponseDataCurrentTransitionsTypeSupersede ContractGetGetContractResponseDataCurrentTransitionsType = "SUPERSEDE"
	ContractGetGetContractResponseDataCurrentTransitionsTypeRenewal   ContractGetGetContractResponseDataCurrentTransitionsType = "RENEWAL"
)

type ContractGetGetContractResponseDataCurrentUsageInvoiceSchedule struct {
	Frequency ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractGetGetContractResponseDataCurrentUsageInvoiceScheduleJSON
}

// contractGetGetContractResponseDataCurrentUsageInvoiceScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentUsageInvoiceSchedule]
type contractGetGetContractResponseDataCurrentUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequency string

const (
	ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequency = "monthly"
	ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractGetGetContractResponseDataCurrentUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractGetGetContractResponseDataCurrentUsageFilter struct {
	Current ContractGetGetContractResponseDataCurrentUsageFilterCurrent  `json:"current,required"`
	Initial ContractGetGetContractResponseDataCurrentUsageFilterInitial  `json:"initial,required"`
	Updates []ContractGetGetContractResponseDataCurrentUsageFilterUpdate `json:"updates,required"`
	JSON    contractGetGetContractResponseDataCurrentUsageFilterJSON
}

// contractGetGetContractResponseDataCurrentUsageFilterJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataCurrentUsageFilter]
type contractGetGetContractResponseDataCurrentUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetGetContractResponseDataCurrentUsageFilterCurrentJSON
}

// contractGetGetContractResponseDataCurrentUsageFilterCurrentJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentUsageFilterCurrent]
type contractGetGetContractResponseDataCurrentUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetGetContractResponseDataCurrentUsageFilterInitialJSON
}

// contractGetGetContractResponseDataCurrentUsageFilterInitialJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataCurrentUsageFilterInitial]
type contractGetGetContractResponseDataCurrentUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataCurrentUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractGetGetContractResponseDataCurrentUsageFilterUpdateJSON
}

// contractGetGetContractResponseDataCurrentUsageFilterUpdateJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataCurrentUsageFilterUpdate]
type contractGetGetContractResponseDataCurrentUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataCurrentUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitial struct {
	Commits                 []ContractGetGetContractResponseDataInitialCommit             `json:"commits,required"`
	CreatedAt               time.Time                                                     `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                                        `json:"created_by,required"`
	Discounts               []ContractGetGetContractResponseDataInitialDiscount           `json:"discounts,required"`
	Overrides               []ContractGetGetContractResponseDataInitialOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractGetGetContractResponseDataInitialResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractGetGetContractResponseDataInitialScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                                     `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractGetGetContractResponseDataInitialTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractGetGetContractResponseDataInitialUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                                     `json:"ending_before" format:"date-time"`
	Name                    string                                                        `json:"name"`
	NetPaymentTermsDays     float64                                                       `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                                        `json:"netsuite_sales_order_id"`
	RateCardID              string                                                        `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                                        `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                                       `json:"total_contract_value"`
	UsageFilter             ContractGetGetContractResponseDataInitialUsageFilter          `json:"usage_filter"`
	JSON                    contractGetGetContractResponseDataInitialJSON
}

// contractGetGetContractResponseDataInitialJSON contains the JSON metadata for the
// struct [ContractGetGetContractResponseDataInitial]
type contractGetGetContractResponseDataInitialJSON struct {
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

func (r *ContractGetGetContractResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialCommit struct {
	ID      string                                                  `json:"id,required" format:"uuid"`
	Product ContractGetGetContractResponseDataInitialCommitsProduct `json:"product,required"`
	Type    ContractGetGetContractResponseDataInitialCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractGetGetContractResponseDataInitialCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractGetGetContractResponseDataInitialCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractGetGetContractResponseDataInitialCommitsLedger       `json:"ledger"`
	Name                 string                                                         `json:"name"`
	NetsuiteSalesOrderID string                                                         `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractGetGetContractResponseDataInitialCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                        `json:"rollover_fraction"`
	JSON                 contractGetGetContractResponseDataInitialCommitJSON
}

// contractGetGetContractResponseDataInitialCommitJSON contains the JSON metadata
// for the struct [ContractGetGetContractResponseDataInitialCommit]
type contractGetGetContractResponseDataInitialCommitJSON struct {
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

func (r *ContractGetGetContractResponseDataInitialCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataInitialCommitsProductJSON
}

// contractGetGetContractResponseDataInitialCommitsProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsProduct]
type contractGetGetContractResponseDataInitialCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialCommitsType string

const (
	ContractGetGetContractResponseDataInitialCommitsTypePrepaid  ContractGetGetContractResponseDataInitialCommitsType = "PREPAID"
	ContractGetGetContractResponseDataInitialCommitsTypePostpaid ContractGetGetContractResponseDataInitialCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractGetGetContractResponseDataInitialCommitsAccessSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractGetGetContractResponseDataInitialCommitsAccessScheduleJSON
}

// contractGetGetContractResponseDataInitialCommitsAccessScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsAccessSchedule]
type contractGetGetContractResponseDataInitialCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItem]
type contractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractGetGetContractResponseDataInitialCommitsInvoiceSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataInitialCommitsInvoiceScheduleJSON
}

// contractGetGetContractResponseDataInitialCommitsInvoiceScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsInvoiceSchedule]
type contractGetGetContractResponseDataInitialCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItem]
type contractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractGetGetContractResponseDataInitialCommitsLedger interface {
	implementsContractGetGetContractResponseDataInitialCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractGetGetContractResponseDataInitialCommitsLedger)(nil)).Elem(), "")
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                        `json:"amount,required"`
	SegmentID string                                                                                         `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                     `json:"amount,required"`
	InvoiceID string                                                                                                      `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                   `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                                    `json:"amount,required"`
	NewContractID string                                                                                     `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type          ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                                      `json:"amount,required"`
	SegmentID string                                                                                       `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                                    `json:"amount,required"`
	InvoiceID string                                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                                    `json:"amount,required"`
	InvoiceID string                                                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractGetGetContractResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                                      `json:"amount,required"`
	InvoiceID string                                                                                                       `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                                    `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                                   `json:"amount,required"`
	InvoiceID string                                                                                    `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractGetGetContractResponseDataInitialCommitsLedger() {
}

type ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractGetGetContractResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractGetGetContractResponseDataInitialCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractGetGetContractResponseDataInitialCommitsRolledOverFromJSON
}

// contractGetGetContractResponseDataInitialCommitsRolledOverFromJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialCommitsRolledOverFrom]
type contractGetGetContractResponseDataInitialCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialDiscount struct {
	ID                   string                                                     `json:"id,required" format:"uuid"`
	Product              ContractGetGetContractResponseDataInitialDiscountsProduct  `json:"product,required"`
	Schedule             ContractGetGetContractResponseDataInitialDiscountsSchedule `json:"schedule,required"`
	Name                 string                                                     `json:"name"`
	NetsuiteSalesOrderID string                                                     `json:"netsuite_sales_order_id"`
	JSON                 contractGetGetContractResponseDataInitialDiscountJSON
}

// contractGetGetContractResponseDataInitialDiscountJSON contains the JSON metadata
// for the struct [ContractGetGetContractResponseDataInitialDiscount]
type contractGetGetContractResponseDataInitialDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataInitialDiscountsProductJSON
}

// contractGetGetContractResponseDataInitialDiscountsProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialDiscountsProduct]
type contractGetGetContractResponseDataInitialDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialDiscountsSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataInitialDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataInitialDiscountsScheduleJSON
}

// contractGetGetContractResponseDataInitialDiscountsScheduleJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialDiscountsSchedule]
type contractGetGetContractResponseDataInitialDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataInitialDiscountsScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataInitialDiscountsScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialDiscountsScheduleScheduleItem]
type contractGetGetContractResponseDataInitialDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialOverride struct {
	ID            string                                                          `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                       `json:"starting_at,required" format:"date-time"`
	Type          ContractGetGetContractResponseDataInitialOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                       `json:"ending_before" format:"date-time"`
	Entitled      bool                                                            `json:"entitled"`
	Multiplier    float64                                                         `json:"multiplier"`
	OverwriteRate ContractGetGetContractResponseDataInitialOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractGetGetContractResponseDataInitialOverridesProduct       `json:"product"`
	Tags          []string                                                        `json:"tags"`
	JSON          contractGetGetContractResponseDataInitialOverrideJSON
}

// contractGetGetContractResponseDataInitialOverrideJSON contains the JSON metadata
// for the struct [ContractGetGetContractResponseDataInitialOverride]
type contractGetGetContractResponseDataInitialOverrideJSON struct {
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

func (r *ContractGetGetContractResponseDataInitialOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialOverridesType string

const (
	ContractGetGetContractResponseDataInitialOverridesTypeOverwrite  ContractGetGetContractResponseDataInitialOverridesType = "OVERWRITE"
	ContractGetGetContractResponseDataInitialOverridesTypeMultiplier ContractGetGetContractResponseDataInitialOverridesType = "MULTIPLIER"
)

type ContractGetGetContractResponseDataInitialOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                                 `json:"price,required"`
	RateType ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractGetGetContractResponseDataInitialOverridesOverwriteRateJSON
}

// contractGetGetContractResponseDataInitialOverridesOverwriteRateJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialOverridesOverwriteRate]
type contractGetGetContractResponseDataInitialOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateType string

const (
	ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateType = "FLAT"
	ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateType = "flat"
	ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractGetGetContractResponseDataInitialOverridesOverwriteRateRateType = "percentage"
)

type ContractGetGetContractResponseDataInitialOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataInitialOverridesProductJSON
}

// contractGetGetContractResponseDataInitialOverridesProductJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialOverridesProduct]
type contractGetGetContractResponseDataInitialOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialResellerRoyalty struct {
	Fraction              float64                                                                `json:"fraction,required"`
	NetsuiteResellerID    string                                                                 `json:"netsuite_reseller_id,required"`
	ResellerType          ContractGetGetContractResponseDataInitialResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                              `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                                 `json:"aws_account_number"`
	AwsOfferID            string                                                                 `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                                 `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                              `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                                 `json:"gcp_account_id"`
	GcpOfferID            string                                                                 `json:"gcp_offer_id"`
	ResellerContractValue float64                                                                `json:"reseller_contract_value"`
	JSON                  contractGetGetContractResponseDataInitialResellerRoyaltyJSON
}

// contractGetGetContractResponseDataInitialResellerRoyaltyJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialResellerRoyalty]
type contractGetGetContractResponseDataInitialResellerRoyaltyJSON struct {
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

func (r *ContractGetGetContractResponseDataInitialResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialResellerRoyaltiesResellerType string

const (
	ContractGetGetContractResponseDataInitialResellerRoyaltiesResellerTypeAws ContractGetGetContractResponseDataInitialResellerRoyaltiesResellerType = "AWS"
	ContractGetGetContractResponseDataInitialResellerRoyaltiesResellerTypeGcp ContractGetGetContractResponseDataInitialResellerRoyaltiesResellerType = "GCP"
)

type ContractGetGetContractResponseDataInitialScheduledCharge struct {
	ID       string                                                            `json:"id,required" format:"uuid"`
	Product  ContractGetGetContractResponseDataInitialScheduledChargesProduct  `json:"product,required"`
	Schedule ContractGetGetContractResponseDataInitialScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractGetGetContractResponseDataInitialScheduledChargeJSON
}

// contractGetGetContractResponseDataInitialScheduledChargeJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialScheduledCharge]
type contractGetGetContractResponseDataInitialScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetGetContractResponseDataInitialScheduledChargesProductJSON
}

// contractGetGetContractResponseDataInitialScheduledChargesProductJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialScheduledChargesProduct]
type contractGetGetContractResponseDataInitialScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialScheduledChargesSchedule struct {
	ScheduleItems []ContractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetGetContractResponseDataInitialScheduledChargesScheduleJSON
}

// contractGetGetContractResponseDataInitialScheduledChargesScheduleJSON contains
// the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialScheduledChargesSchedule]
type contractGetGetContractResponseDataInitialScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItemJSON
}

// contractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItem]
type contractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialTransition struct {
	FromContractID string                                                   `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                                   `json:"to_contract_id,required" format:"uuid"`
	Type           ContractGetGetContractResponseDataInitialTransitionsType `json:"type,required"`
	JSON           contractGetGetContractResponseDataInitialTransitionJSON
}

// contractGetGetContractResponseDataInitialTransitionJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataInitialTransition]
type contractGetGetContractResponseDataInitialTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialTransitionsType string

const (
	ContractGetGetContractResponseDataInitialTransitionsTypeSupersede ContractGetGetContractResponseDataInitialTransitionsType = "SUPERSEDE"
	ContractGetGetContractResponseDataInitialTransitionsTypeRenewal   ContractGetGetContractResponseDataInitialTransitionsType = "RENEWAL"
)

type ContractGetGetContractResponseDataInitialUsageInvoiceSchedule struct {
	Frequency ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractGetGetContractResponseDataInitialUsageInvoiceScheduleJSON
}

// contractGetGetContractResponseDataInitialUsageInvoiceScheduleJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialUsageInvoiceSchedule]
type contractGetGetContractResponseDataInitialUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequency string

const (
	ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequency = "monthly"
	ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractGetGetContractResponseDataInitialUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractGetGetContractResponseDataInitialUsageFilter struct {
	Current ContractGetGetContractResponseDataInitialUsageFilterCurrent  `json:"current,required"`
	Initial ContractGetGetContractResponseDataInitialUsageFilterInitial  `json:"initial,required"`
	Updates []ContractGetGetContractResponseDataInitialUsageFilterUpdate `json:"updates,required"`
	JSON    contractGetGetContractResponseDataInitialUsageFilterJSON
}

// contractGetGetContractResponseDataInitialUsageFilterJSON contains the JSON
// metadata for the struct [ContractGetGetContractResponseDataInitialUsageFilter]
type contractGetGetContractResponseDataInitialUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetGetContractResponseDataInitialUsageFilterCurrentJSON
}

// contractGetGetContractResponseDataInitialUsageFilterCurrentJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialUsageFilterCurrent]
type contractGetGetContractResponseDataInitialUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetGetContractResponseDataInitialUsageFilterInitialJSON
}

// contractGetGetContractResponseDataInitialUsageFilterInitialJSON contains the
// JSON metadata for the struct
// [ContractGetGetContractResponseDataInitialUsageFilterInitial]
type contractGetGetContractResponseDataInitialUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractResponseDataInitialUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractGetGetContractResponseDataInitialUsageFilterUpdateJSON
}

// contractGetGetContractResponseDataInitialUsageFilterUpdateJSON contains the JSON
// metadata for the struct
// [ContractGetGetContractResponseDataInitialUsageFilterUpdate]
type contractGetGetContractResponseDataInitialUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetGetContractResponseDataInitialUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetGetContractParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r ContractGetGetContractParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
