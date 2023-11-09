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

// ContractService contains methods and other services that help with interacting
// with the metronome API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewContractService] method instead.
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

// Get a specific contract
func (r *ContractService) Get(ctx context.Context, body ContractGetParams, opts ...option.RequestOption) (res *ContractGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "contracts/get"
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

type ContractNewResponse struct {
	Data ContractNewResponseData `json:"data,required"`
	JSON contractNewResponseJSON
}

// contractNewResponseJSON contains the JSON metadata for the struct
// [ContractNewResponse]
type contractNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNewResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractNewResponseDataJSON
}

// contractNewResponseDataJSON contains the JSON metadata for the struct
// [ContractNewResponseData]
type contractNewResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponse struct {
	Data []ContractListResponseData `json:"data,required"`
	JSON contractListResponseJSON
}

// contractListResponseJSON contains the JSON metadata for the struct
// [ContractListResponse]
type contractListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseData struct {
	ID         string                              `json:"id,required" format:"uuid"`
	Amendments []ContractListResponseDataAmendment `json:"amendments,required"`
	Current    ContractListResponseDataCurrent     `json:"current,required"`
	CustomerID string                              `json:"customer_id,required" format:"uuid"`
	Initial    ContractListResponseDataInitial     `json:"initial,required"`
	JSON       contractListResponseDataJSON
}

// contractListResponseDataJSON contains the JSON metadata for the struct
// [ContractListResponseData]
type contractListResponseDataJSON struct {
	ID          apijson.Field
	Amendments  apijson.Field
	Current     apijson.Field
	CustomerID  apijson.Field
	Initial     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendment struct {
	ID                      string                                              `json:"id,required" format:"uuid"`
	Commits                 []ContractListResponseDataAmendmentsCommit          `json:"commits,required"`
	CreatedAt               time.Time                                           `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                              `json:"created_by,required"`
	Discounts               []ContractListResponseDataAmendmentsDiscount        `json:"discounts,required"`
	Overrides               []ContractListResponseDataAmendmentsOverride        `json:"overrides,required"`
	ResellerRoyalties       []ContractListResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractListResponseDataAmendmentsScheduledCharge `json:"scheduled_charges,required"`
	StartingAt              time.Time                                           `json:"starting_at,required" format:"date-time"`
	NetsuiteSalesOrderID    string                                              `json:"netsuite_sales_order_id"`
	SalesforceOpportunityID string                                              `json:"salesforce_opportunity_id"`
	JSON                    contractListResponseDataAmendmentJSON
}

// contractListResponseDataAmendmentJSON contains the JSON metadata for the struct
// [ContractListResponseDataAmendment]
type contractListResponseDataAmendmentJSON struct {
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

func (r *ContractListResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsCommit struct {
	ID      string                                           `json:"id,required" format:"uuid"`
	Product ContractListResponseDataAmendmentsCommitsProduct `json:"product,required"`
	Type    ContractListResponseDataAmendmentsCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractListResponseDataAmendmentsCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractListResponseDataAmendmentsCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractListResponseDataAmendmentsCommitsLedger       `json:"ledger"`
	Name                 string                                                  `json:"name"`
	NetsuiteSalesOrderID string                                                  `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractListResponseDataAmendmentsCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                 `json:"rollover_fraction"`
	JSON                 contractListResponseDataAmendmentsCommitJSON
}

// contractListResponseDataAmendmentsCommitJSON contains the JSON metadata for the
// struct [ContractListResponseDataAmendmentsCommit]
type contractListResponseDataAmendmentsCommitJSON struct {
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

func (r *ContractListResponseDataAmendmentsCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataAmendmentsCommitsProductJSON
}

// contractListResponseDataAmendmentsCommitsProductJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsCommitsProduct]
type contractListResponseDataAmendmentsCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsCommitsType string

const (
	ContractListResponseDataAmendmentsCommitsTypePrepaid  ContractListResponseDataAmendmentsCommitsType = "PREPAID"
	ContractListResponseDataAmendmentsCommitsTypePostpaid ContractListResponseDataAmendmentsCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractListResponseDataAmendmentsCommitsAccessSchedule struct {
	ScheduleItems []ContractListResponseDataAmendmentsCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractListResponseDataAmendmentsCommitsAccessScheduleJSON
}

// contractListResponseDataAmendmentsCommitsAccessScheduleJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsCommitsAccessSchedule]
type contractListResponseDataAmendmentsCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractListResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON
}

// contractListResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsAccessScheduleScheduleItem]
type contractListResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractListResponseDataAmendmentsCommitsInvoiceSchedule struct {
	ScheduleItems []ContractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataAmendmentsCommitsInvoiceScheduleJSON
}

// contractListResponseDataAmendmentsCommitsInvoiceScheduleJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsCommitsInvoiceSchedule]
type contractListResponseDataAmendmentsCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
}

// contractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem]
type contractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractListResponseDataAmendmentsCommitsLedger interface {
	implementsContractListResponseDataAmendmentsCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractListResponseDataAmendmentsCommitsLedger)(nil)).Elem(), "")
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                 `json:"amount,required"`
	SegmentID string                                                                                  `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                               `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                              `json:"amount,required"`
	InvoiceID string                                                                                               `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                             `json:"amount,required"`
	NewContractID string                                                                              `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                           `json:"timestamp,required" format:"date-time"`
	Type          ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                               `json:"amount,required"`
	SegmentID string                                                                                `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                             `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                             `json:"amount,required"`
	InvoiceID string                                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                             `json:"amount,required"`
	InvoiceID string                                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractListResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                    `json:"amount,required"`
	Timestamp time.Time                                                                                  `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                               `json:"amount,required"`
	InvoiceID string                                                                                                `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                             `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	InvoiceID string                                                                             `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractListResponseDataAmendmentsCommitsLedger() {
}

type ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractListResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractListResponseDataAmendmentsCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractListResponseDataAmendmentsCommitsRolledOverFromJSON
}

// contractListResponseDataAmendmentsCommitsRolledOverFromJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsCommitsRolledOverFrom]
type contractListResponseDataAmendmentsCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsDiscount struct {
	ID                   string                                              `json:"id,required" format:"uuid"`
	Product              ContractListResponseDataAmendmentsDiscountsProduct  `json:"product,required"`
	Schedule             ContractListResponseDataAmendmentsDiscountsSchedule `json:"schedule,required"`
	Name                 string                                              `json:"name"`
	NetsuiteSalesOrderID string                                              `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataAmendmentsDiscountJSON
}

// contractListResponseDataAmendmentsDiscountJSON contains the JSON metadata for
// the struct [ContractListResponseDataAmendmentsDiscount]
type contractListResponseDataAmendmentsDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataAmendmentsDiscountsProductJSON
}

// contractListResponseDataAmendmentsDiscountsProductJSON contains the JSON
// metadata for the struct [ContractListResponseDataAmendmentsDiscountsProduct]
type contractListResponseDataAmendmentsDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsDiscountsSchedule struct {
	ScheduleItems []ContractListResponseDataAmendmentsDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataAmendmentsDiscountsScheduleJSON
}

// contractListResponseDataAmendmentsDiscountsScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataAmendmentsDiscountsSchedule]
type contractListResponseDataAmendmentsDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataAmendmentsDiscountsScheduleScheduleItemJSON
}

// contractListResponseDataAmendmentsDiscountsScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractListResponseDataAmendmentsDiscountsScheduleScheduleItem]
type contractListResponseDataAmendmentsDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsOverride struct {
	ID            string                                                   `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                                `json:"starting_at,required" format:"date-time"`
	Type          ContractListResponseDataAmendmentsOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                                `json:"ending_before" format:"date-time"`
	Entitled      bool                                                     `json:"entitled"`
	Multiplier    float64                                                  `json:"multiplier"`
	OverwriteRate ContractListResponseDataAmendmentsOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractListResponseDataAmendmentsOverridesProduct       `json:"product"`
	Tags          []string                                                 `json:"tags"`
	JSON          contractListResponseDataAmendmentsOverrideJSON
}

// contractListResponseDataAmendmentsOverrideJSON contains the JSON metadata for
// the struct [ContractListResponseDataAmendmentsOverride]
type contractListResponseDataAmendmentsOverrideJSON struct {
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

func (r *ContractListResponseDataAmendmentsOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsOverridesType string

const (
	ContractListResponseDataAmendmentsOverridesTypeOverwrite  ContractListResponseDataAmendmentsOverridesType = "OVERWRITE"
	ContractListResponseDataAmendmentsOverridesTypeMultiplier ContractListResponseDataAmendmentsOverridesType = "MULTIPLIER"
)

type ContractListResponseDataAmendmentsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                          `json:"price,required"`
	RateType ContractListResponseDataAmendmentsOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractListResponseDataAmendmentsOverridesOverwriteRateJSON
}

// contractListResponseDataAmendmentsOverridesOverwriteRateJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsOverridesOverwriteRate]
type contractListResponseDataAmendmentsOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsOverridesOverwriteRateRateType string

const (
	ContractListResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractListResponseDataAmendmentsOverridesOverwriteRateRateType = "FLAT"
	ContractListResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractListResponseDataAmendmentsOverridesOverwriteRateRateType = "flat"
	ContractListResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractListResponseDataAmendmentsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractListResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractListResponseDataAmendmentsOverridesOverwriteRateRateType = "percentage"
)

type ContractListResponseDataAmendmentsOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataAmendmentsOverridesProductJSON
}

// contractListResponseDataAmendmentsOverridesProductJSON contains the JSON
// metadata for the struct [ContractListResponseDataAmendmentsOverridesProduct]
type contractListResponseDataAmendmentsOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          ContractListResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                          `json:"aws_account_number"`
	AwsOfferID            string                                                          `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                          `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                       `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                         `json:"fraction"`
	GcpAccountID          string                                                          `json:"gcp_account_id"`
	GcpOfferID            string                                                          `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                          `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                         `json:"reseller_contract_value"`
	StartingAt            time.Time                                                       `json:"starting_at" format:"date-time"`
	JSON                  contractListResponseDataAmendmentsResellerRoyaltyJSON
}

// contractListResponseDataAmendmentsResellerRoyaltyJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsResellerRoyalty]
type contractListResponseDataAmendmentsResellerRoyaltyJSON struct {
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

func (r *ContractListResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAws ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
)

type ContractListResponseDataAmendmentsScheduledCharge struct {
	ID       string                                                     `json:"id,required" format:"uuid"`
	Product  ContractListResponseDataAmendmentsScheduledChargesProduct  `json:"product,required"`
	Schedule ContractListResponseDataAmendmentsScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataAmendmentsScheduledChargeJSON
}

// contractListResponseDataAmendmentsScheduledChargeJSON contains the JSON metadata
// for the struct [ContractListResponseDataAmendmentsScheduledCharge]
type contractListResponseDataAmendmentsScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataAmendmentsScheduledChargesProductJSON
}

// contractListResponseDataAmendmentsScheduledChargesProductJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsScheduledChargesProduct]
type contractListResponseDataAmendmentsScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsScheduledChargesSchedule struct {
	ScheduleItems []ContractListResponseDataAmendmentsScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataAmendmentsScheduledChargesScheduleJSON
}

// contractListResponseDataAmendmentsScheduledChargesScheduleJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataAmendmentsScheduledChargesSchedule]
type contractListResponseDataAmendmentsScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataAmendmentsScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
}

// contractListResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataAmendmentsScheduledChargesScheduleScheduleItem]
type contractListResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataAmendmentsScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrent struct {
	Commits                 []ContractListResponseDataCurrentCommit             `json:"commits,required"`
	CreatedAt               time.Time                                           `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                              `json:"created_by,required"`
	Discounts               []ContractListResponseDataCurrentDiscount           `json:"discounts,required"`
	Overrides               []ContractListResponseDataCurrentOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractListResponseDataCurrentResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractListResponseDataCurrentScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                           `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractListResponseDataCurrentTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractListResponseDataCurrentUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                           `json:"ending_before" format:"date-time"`
	Name                    string                                              `json:"name"`
	NetPaymentTermsDays     float64                                             `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                              `json:"netsuite_sales_order_id"`
	RateCardID              string                                              `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                              `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                             `json:"total_contract_value"`
	UsageFilter             ContractListResponseDataCurrentUsageFilter          `json:"usage_filter"`
	JSON                    contractListResponseDataCurrentJSON
}

// contractListResponseDataCurrentJSON contains the JSON metadata for the struct
// [ContractListResponseDataCurrent]
type contractListResponseDataCurrentJSON struct {
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

func (r *ContractListResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentCommit struct {
	ID      string                                        `json:"id,required" format:"uuid"`
	Product ContractListResponseDataCurrentCommitsProduct `json:"product,required"`
	Type    ContractListResponseDataCurrentCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractListResponseDataCurrentCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractListResponseDataCurrentCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractListResponseDataCurrentCommitsLedger       `json:"ledger"`
	Name                 string                                               `json:"name"`
	NetsuiteSalesOrderID string                                               `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractListResponseDataCurrentCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                              `json:"rollover_fraction"`
	JSON                 contractListResponseDataCurrentCommitJSON
}

// contractListResponseDataCurrentCommitJSON contains the JSON metadata for the
// struct [ContractListResponseDataCurrentCommit]
type contractListResponseDataCurrentCommitJSON struct {
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

func (r *ContractListResponseDataCurrentCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataCurrentCommitsProductJSON
}

// contractListResponseDataCurrentCommitsProductJSON contains the JSON metadata for
// the struct [ContractListResponseDataCurrentCommitsProduct]
type contractListResponseDataCurrentCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentCommitsType string

const (
	ContractListResponseDataCurrentCommitsTypePrepaid  ContractListResponseDataCurrentCommitsType = "PREPAID"
	ContractListResponseDataCurrentCommitsTypePostpaid ContractListResponseDataCurrentCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractListResponseDataCurrentCommitsAccessSchedule struct {
	ScheduleItems []ContractListResponseDataCurrentCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractListResponseDataCurrentCommitsAccessScheduleJSON
}

// contractListResponseDataCurrentCommitsAccessScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataCurrentCommitsAccessSchedule]
type contractListResponseDataCurrentCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractListResponseDataCurrentCommitsAccessScheduleScheduleItemJSON
}

// contractListResponseDataCurrentCommitsAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsAccessScheduleScheduleItem]
type contractListResponseDataCurrentCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractListResponseDataCurrentCommitsInvoiceSchedule struct {
	ScheduleItems []ContractListResponseDataCurrentCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataCurrentCommitsInvoiceScheduleJSON
}

// contractListResponseDataCurrentCommitsInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataCurrentCommitsInvoiceSchedule]
type contractListResponseDataCurrentCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON
}

// contractListResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsInvoiceScheduleScheduleItem]
type contractListResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractListResponseDataCurrentCommitsLedger interface {
	implementsContractListResponseDataCurrentCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractListResponseDataCurrentCommitsLedger)(nil)).Elem(), "")
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                              `json:"amount,required"`
	SegmentID string                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractListResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	InvoiceID string                                                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractListResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                          `json:"amount,required"`
	NewContractID string                                                                           `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type          ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractListResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	SegmentID string                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractListResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                          `json:"amount,required"`
	InvoiceID string                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                          `json:"amount,required"`
	InvoiceID string                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractListResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                 `json:"amount,required"`
	Timestamp time.Time                                                                               `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractListResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                            `json:"amount,required"`
	InvoiceID string                                                                                             `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractListResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	InvoiceID string                                                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractListResponseDataCurrentCommitsLedger() {
}

type ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractListResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractListResponseDataCurrentCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractListResponseDataCurrentCommitsRolledOverFromJSON
}

// contractListResponseDataCurrentCommitsRolledOverFromJSON contains the JSON
// metadata for the struct [ContractListResponseDataCurrentCommitsRolledOverFrom]
type contractListResponseDataCurrentCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentDiscount struct {
	ID                   string                                           `json:"id,required" format:"uuid"`
	Product              ContractListResponseDataCurrentDiscountsProduct  `json:"product,required"`
	Schedule             ContractListResponseDataCurrentDiscountsSchedule `json:"schedule,required"`
	Name                 string                                           `json:"name"`
	NetsuiteSalesOrderID string                                           `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataCurrentDiscountJSON
}

// contractListResponseDataCurrentDiscountJSON contains the JSON metadata for the
// struct [ContractListResponseDataCurrentDiscount]
type contractListResponseDataCurrentDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataCurrentDiscountsProductJSON
}

// contractListResponseDataCurrentDiscountsProductJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentDiscountsProduct]
type contractListResponseDataCurrentDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentDiscountsSchedule struct {
	ScheduleItems []ContractListResponseDataCurrentDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataCurrentDiscountsScheduleJSON
}

// contractListResponseDataCurrentDiscountsScheduleJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentDiscountsSchedule]
type contractListResponseDataCurrentDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataCurrentDiscountsScheduleScheduleItemJSON
}

// contractListResponseDataCurrentDiscountsScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractListResponseDataCurrentDiscountsScheduleScheduleItem]
type contractListResponseDataCurrentDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentOverride struct {
	ID            string                                                `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                             `json:"starting_at,required" format:"date-time"`
	Type          ContractListResponseDataCurrentOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                             `json:"ending_before" format:"date-time"`
	Entitled      bool                                                  `json:"entitled"`
	Multiplier    float64                                               `json:"multiplier"`
	OverwriteRate ContractListResponseDataCurrentOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractListResponseDataCurrentOverridesProduct       `json:"product"`
	Tags          []string                                              `json:"tags"`
	JSON          contractListResponseDataCurrentOverrideJSON
}

// contractListResponseDataCurrentOverrideJSON contains the JSON metadata for the
// struct [ContractListResponseDataCurrentOverride]
type contractListResponseDataCurrentOverrideJSON struct {
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

func (r *ContractListResponseDataCurrentOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentOverridesType string

const (
	ContractListResponseDataCurrentOverridesTypeOverwrite  ContractListResponseDataCurrentOverridesType = "OVERWRITE"
	ContractListResponseDataCurrentOverridesTypeMultiplier ContractListResponseDataCurrentOverridesType = "MULTIPLIER"
)

type ContractListResponseDataCurrentOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                       `json:"price,required"`
	RateType ContractListResponseDataCurrentOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractListResponseDataCurrentOverridesOverwriteRateJSON
}

// contractListResponseDataCurrentOverridesOverwriteRateJSON contains the JSON
// metadata for the struct [ContractListResponseDataCurrentOverridesOverwriteRate]
type contractListResponseDataCurrentOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentOverridesOverwriteRateRateType string

const (
	ContractListResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractListResponseDataCurrentOverridesOverwriteRateRateType = "FLAT"
	ContractListResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractListResponseDataCurrentOverridesOverwriteRateRateType = "flat"
	ContractListResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractListResponseDataCurrentOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractListResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractListResponseDataCurrentOverridesOverwriteRateRateType = "percentage"
)

type ContractListResponseDataCurrentOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataCurrentOverridesProductJSON
}

// contractListResponseDataCurrentOverridesProductJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentOverridesProduct]
type contractListResponseDataCurrentOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentResellerRoyalty struct {
	Fraction              float64                                                      `json:"fraction,required"`
	NetsuiteResellerID    string                                                       `json:"netsuite_reseller_id,required"`
	ResellerType          ContractListResponseDataCurrentResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                    `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                       `json:"aws_account_number"`
	AwsOfferID            string                                                       `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                       `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                    `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                       `json:"gcp_account_id"`
	GcpOfferID            string                                                       `json:"gcp_offer_id"`
	ResellerContractValue float64                                                      `json:"reseller_contract_value"`
	JSON                  contractListResponseDataCurrentResellerRoyaltyJSON
}

// contractListResponseDataCurrentResellerRoyaltyJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentResellerRoyalty]
type contractListResponseDataCurrentResellerRoyaltyJSON struct {
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

func (r *ContractListResponseDataCurrentResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentResellerRoyaltiesResellerType string

const (
	ContractListResponseDataCurrentResellerRoyaltiesResellerTypeAws ContractListResponseDataCurrentResellerRoyaltiesResellerType = "AWS"
	ContractListResponseDataCurrentResellerRoyaltiesResellerTypeGcp ContractListResponseDataCurrentResellerRoyaltiesResellerType = "GCP"
)

type ContractListResponseDataCurrentScheduledCharge struct {
	ID       string                                                  `json:"id,required" format:"uuid"`
	Product  ContractListResponseDataCurrentScheduledChargesProduct  `json:"product,required"`
	Schedule ContractListResponseDataCurrentScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataCurrentScheduledChargeJSON
}

// contractListResponseDataCurrentScheduledChargeJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentScheduledCharge]
type contractListResponseDataCurrentScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataCurrentScheduledChargesProductJSON
}

// contractListResponseDataCurrentScheduledChargesProductJSON contains the JSON
// metadata for the struct [ContractListResponseDataCurrentScheduledChargesProduct]
type contractListResponseDataCurrentScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentScheduledChargesSchedule struct {
	ScheduleItems []ContractListResponseDataCurrentScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataCurrentScheduledChargesScheduleJSON
}

// contractListResponseDataCurrentScheduledChargesScheduleJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataCurrentScheduledChargesSchedule]
type contractListResponseDataCurrentScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataCurrentScheduledChargesScheduleScheduleItemJSON
}

// contractListResponseDataCurrentScheduledChargesScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataCurrentScheduledChargesScheduleScheduleItem]
type contractListResponseDataCurrentScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentTransition struct {
	FromContractID string                                         `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                         `json:"to_contract_id,required" format:"uuid"`
	Type           ContractListResponseDataCurrentTransitionsType `json:"type,required"`
	JSON           contractListResponseDataCurrentTransitionJSON
}

// contractListResponseDataCurrentTransitionJSON contains the JSON metadata for the
// struct [ContractListResponseDataCurrentTransition]
type contractListResponseDataCurrentTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentTransitionsType string

const (
	ContractListResponseDataCurrentTransitionsTypeSupersede ContractListResponseDataCurrentTransitionsType = "SUPERSEDE"
	ContractListResponseDataCurrentTransitionsTypeRenewal   ContractListResponseDataCurrentTransitionsType = "RENEWAL"
)

type ContractListResponseDataCurrentUsageInvoiceSchedule struct {
	Frequency ContractListResponseDataCurrentUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractListResponseDataCurrentUsageInvoiceScheduleJSON
}

// contractListResponseDataCurrentUsageInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataCurrentUsageInvoiceSchedule]
type contractListResponseDataCurrentUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentUsageInvoiceScheduleFrequency string

const (
	ContractListResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractListResponseDataCurrentUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractListResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractListResponseDataCurrentUsageInvoiceScheduleFrequency = "monthly"
	ContractListResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractListResponseDataCurrentUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractListResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractListResponseDataCurrentUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractListResponseDataCurrentUsageFilter struct {
	Current ContractListResponseDataCurrentUsageFilterCurrent  `json:"current,required"`
	Initial ContractListResponseDataCurrentUsageFilterInitial  `json:"initial,required"`
	Updates []ContractListResponseDataCurrentUsageFilterUpdate `json:"updates,required"`
	JSON    contractListResponseDataCurrentUsageFilterJSON
}

// contractListResponseDataCurrentUsageFilterJSON contains the JSON metadata for
// the struct [ContractListResponseDataCurrentUsageFilter]
type contractListResponseDataCurrentUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListResponseDataCurrentUsageFilterCurrentJSON
}

// contractListResponseDataCurrentUsageFilterCurrentJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentUsageFilterCurrent]
type contractListResponseDataCurrentUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListResponseDataCurrentUsageFilterInitialJSON
}

// contractListResponseDataCurrentUsageFilterInitialJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentUsageFilterInitial]
type contractListResponseDataCurrentUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataCurrentUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractListResponseDataCurrentUsageFilterUpdateJSON
}

// contractListResponseDataCurrentUsageFilterUpdateJSON contains the JSON metadata
// for the struct [ContractListResponseDataCurrentUsageFilterUpdate]
type contractListResponseDataCurrentUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataCurrentUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitial struct {
	Commits                 []ContractListResponseDataInitialCommit             `json:"commits,required"`
	CreatedAt               time.Time                                           `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                              `json:"created_by,required"`
	Discounts               []ContractListResponseDataInitialDiscount           `json:"discounts,required"`
	Overrides               []ContractListResponseDataInitialOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractListResponseDataInitialResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractListResponseDataInitialScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                           `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractListResponseDataInitialTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractListResponseDataInitialUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                           `json:"ending_before" format:"date-time"`
	Name                    string                                              `json:"name"`
	NetPaymentTermsDays     float64                                             `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                              `json:"netsuite_sales_order_id"`
	RateCardID              string                                              `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                              `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                             `json:"total_contract_value"`
	UsageFilter             ContractListResponseDataInitialUsageFilter          `json:"usage_filter"`
	JSON                    contractListResponseDataInitialJSON
}

// contractListResponseDataInitialJSON contains the JSON metadata for the struct
// [ContractListResponseDataInitial]
type contractListResponseDataInitialJSON struct {
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

func (r *ContractListResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialCommit struct {
	ID      string                                        `json:"id,required" format:"uuid"`
	Product ContractListResponseDataInitialCommitsProduct `json:"product,required"`
	Type    ContractListResponseDataInitialCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractListResponseDataInitialCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractListResponseDataInitialCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractListResponseDataInitialCommitsLedger       `json:"ledger"`
	Name                 string                                               `json:"name"`
	NetsuiteSalesOrderID string                                               `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractListResponseDataInitialCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                              `json:"rollover_fraction"`
	JSON                 contractListResponseDataInitialCommitJSON
}

// contractListResponseDataInitialCommitJSON contains the JSON metadata for the
// struct [ContractListResponseDataInitialCommit]
type contractListResponseDataInitialCommitJSON struct {
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

func (r *ContractListResponseDataInitialCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataInitialCommitsProductJSON
}

// contractListResponseDataInitialCommitsProductJSON contains the JSON metadata for
// the struct [ContractListResponseDataInitialCommitsProduct]
type contractListResponseDataInitialCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialCommitsType string

const (
	ContractListResponseDataInitialCommitsTypePrepaid  ContractListResponseDataInitialCommitsType = "PREPAID"
	ContractListResponseDataInitialCommitsTypePostpaid ContractListResponseDataInitialCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractListResponseDataInitialCommitsAccessSchedule struct {
	ScheduleItems []ContractListResponseDataInitialCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractListResponseDataInitialCommitsAccessScheduleJSON
}

// contractListResponseDataInitialCommitsAccessScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataInitialCommitsAccessSchedule]
type contractListResponseDataInitialCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractListResponseDataInitialCommitsAccessScheduleScheduleItemJSON
}

// contractListResponseDataInitialCommitsAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsAccessScheduleScheduleItem]
type contractListResponseDataInitialCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractListResponseDataInitialCommitsInvoiceSchedule struct {
	ScheduleItems []ContractListResponseDataInitialCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataInitialCommitsInvoiceScheduleJSON
}

// contractListResponseDataInitialCommitsInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataInitialCommitsInvoiceSchedule]
type contractListResponseDataInitialCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON
}

// contractListResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsInvoiceScheduleScheduleItem]
type contractListResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractListResponseDataInitialCommitsLedger interface {
	implementsContractListResponseDataInitialCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractListResponseDataInitialCommitsLedger)(nil)).Elem(), "")
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                              `json:"amount,required"`
	SegmentID string                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractListResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	InvoiceID string                                                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractListResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                          `json:"amount,required"`
	NewContractID string                                                                           `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type          ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractListResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	SegmentID string                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractListResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                          `json:"amount,required"`
	InvoiceID string                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractListResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                          `json:"amount,required"`
	InvoiceID string                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractListResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                 `json:"amount,required"`
	Timestamp time.Time                                                                               `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractListResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                            `json:"amount,required"`
	InvoiceID string                                                                                             `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractListResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	InvoiceID string                                                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractListResponseDataInitialCommitsLedger() {
}

type ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractListResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractListResponseDataInitialCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractListResponseDataInitialCommitsRolledOverFromJSON
}

// contractListResponseDataInitialCommitsRolledOverFromJSON contains the JSON
// metadata for the struct [ContractListResponseDataInitialCommitsRolledOverFrom]
type contractListResponseDataInitialCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialDiscount struct {
	ID                   string                                           `json:"id,required" format:"uuid"`
	Product              ContractListResponseDataInitialDiscountsProduct  `json:"product,required"`
	Schedule             ContractListResponseDataInitialDiscountsSchedule `json:"schedule,required"`
	Name                 string                                           `json:"name"`
	NetsuiteSalesOrderID string                                           `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataInitialDiscountJSON
}

// contractListResponseDataInitialDiscountJSON contains the JSON metadata for the
// struct [ContractListResponseDataInitialDiscount]
type contractListResponseDataInitialDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataInitialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataInitialDiscountsProductJSON
}

// contractListResponseDataInitialDiscountsProductJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialDiscountsProduct]
type contractListResponseDataInitialDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialDiscountsSchedule struct {
	ScheduleItems []ContractListResponseDataInitialDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataInitialDiscountsScheduleJSON
}

// contractListResponseDataInitialDiscountsScheduleJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialDiscountsSchedule]
type contractListResponseDataInitialDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataInitialDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataInitialDiscountsScheduleScheduleItemJSON
}

// contractListResponseDataInitialDiscountsScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractListResponseDataInitialDiscountsScheduleScheduleItem]
type contractListResponseDataInitialDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialOverride struct {
	ID            string                                                `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                             `json:"starting_at,required" format:"date-time"`
	Type          ContractListResponseDataInitialOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                             `json:"ending_before" format:"date-time"`
	Entitled      bool                                                  `json:"entitled"`
	Multiplier    float64                                               `json:"multiplier"`
	OverwriteRate ContractListResponseDataInitialOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractListResponseDataInitialOverridesProduct       `json:"product"`
	Tags          []string                                              `json:"tags"`
	JSON          contractListResponseDataInitialOverrideJSON
}

// contractListResponseDataInitialOverrideJSON contains the JSON metadata for the
// struct [ContractListResponseDataInitialOverride]
type contractListResponseDataInitialOverrideJSON struct {
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

func (r *ContractListResponseDataInitialOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialOverridesType string

const (
	ContractListResponseDataInitialOverridesTypeOverwrite  ContractListResponseDataInitialOverridesType = "OVERWRITE"
	ContractListResponseDataInitialOverridesTypeMultiplier ContractListResponseDataInitialOverridesType = "MULTIPLIER"
)

type ContractListResponseDataInitialOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                       `json:"price,required"`
	RateType ContractListResponseDataInitialOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractListResponseDataInitialOverridesOverwriteRateJSON
}

// contractListResponseDataInitialOverridesOverwriteRateJSON contains the JSON
// metadata for the struct [ContractListResponseDataInitialOverridesOverwriteRate]
type contractListResponseDataInitialOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataInitialOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialOverridesOverwriteRateRateType string

const (
	ContractListResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractListResponseDataInitialOverridesOverwriteRateRateType = "FLAT"
	ContractListResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractListResponseDataInitialOverridesOverwriteRateRateType = "flat"
	ContractListResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractListResponseDataInitialOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractListResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractListResponseDataInitialOverridesOverwriteRateRateType = "percentage"
)

type ContractListResponseDataInitialOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataInitialOverridesProductJSON
}

// contractListResponseDataInitialOverridesProductJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialOverridesProduct]
type contractListResponseDataInitialOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialResellerRoyalty struct {
	Fraction              float64                                                      `json:"fraction,required"`
	NetsuiteResellerID    string                                                       `json:"netsuite_reseller_id,required"`
	ResellerType          ContractListResponseDataInitialResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                    `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                       `json:"aws_account_number"`
	AwsOfferID            string                                                       `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                       `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                    `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                       `json:"gcp_account_id"`
	GcpOfferID            string                                                       `json:"gcp_offer_id"`
	ResellerContractValue float64                                                      `json:"reseller_contract_value"`
	JSON                  contractListResponseDataInitialResellerRoyaltyJSON
}

// contractListResponseDataInitialResellerRoyaltyJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialResellerRoyalty]
type contractListResponseDataInitialResellerRoyaltyJSON struct {
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

func (r *ContractListResponseDataInitialResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialResellerRoyaltiesResellerType string

const (
	ContractListResponseDataInitialResellerRoyaltiesResellerTypeAws ContractListResponseDataInitialResellerRoyaltiesResellerType = "AWS"
	ContractListResponseDataInitialResellerRoyaltiesResellerTypeGcp ContractListResponseDataInitialResellerRoyaltiesResellerType = "GCP"
)

type ContractListResponseDataInitialScheduledCharge struct {
	ID       string                                                  `json:"id,required" format:"uuid"`
	Product  ContractListResponseDataInitialScheduledChargesProduct  `json:"product,required"`
	Schedule ContractListResponseDataInitialScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractListResponseDataInitialScheduledChargeJSON
}

// contractListResponseDataInitialScheduledChargeJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialScheduledCharge]
type contractListResponseDataInitialScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractListResponseDataInitialScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractListResponseDataInitialScheduledChargesProductJSON
}

// contractListResponseDataInitialScheduledChargesProductJSON contains the JSON
// metadata for the struct [ContractListResponseDataInitialScheduledChargesProduct]
type contractListResponseDataInitialScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialScheduledChargesSchedule struct {
	ScheduleItems []ContractListResponseDataInitialScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractListResponseDataInitialScheduledChargesScheduleJSON
}

// contractListResponseDataInitialScheduledChargesScheduleJSON contains the JSON
// metadata for the struct
// [ContractListResponseDataInitialScheduledChargesSchedule]
type contractListResponseDataInitialScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractListResponseDataInitialScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractListResponseDataInitialScheduledChargesScheduleScheduleItemJSON
}

// contractListResponseDataInitialScheduledChargesScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractListResponseDataInitialScheduledChargesScheduleScheduleItem]
type contractListResponseDataInitialScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialTransition struct {
	FromContractID string                                         `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                         `json:"to_contract_id,required" format:"uuid"`
	Type           ContractListResponseDataInitialTransitionsType `json:"type,required"`
	JSON           contractListResponseDataInitialTransitionJSON
}

// contractListResponseDataInitialTransitionJSON contains the JSON metadata for the
// struct [ContractListResponseDataInitialTransition]
type contractListResponseDataInitialTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractListResponseDataInitialTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialTransitionsType string

const (
	ContractListResponseDataInitialTransitionsTypeSupersede ContractListResponseDataInitialTransitionsType = "SUPERSEDE"
	ContractListResponseDataInitialTransitionsTypeRenewal   ContractListResponseDataInitialTransitionsType = "RENEWAL"
)

type ContractListResponseDataInitialUsageInvoiceSchedule struct {
	Frequency ContractListResponseDataInitialUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractListResponseDataInitialUsageInvoiceScheduleJSON
}

// contractListResponseDataInitialUsageInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractListResponseDataInitialUsageInvoiceSchedule]
type contractListResponseDataInitialUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialUsageInvoiceScheduleFrequency string

const (
	ContractListResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractListResponseDataInitialUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractListResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractListResponseDataInitialUsageInvoiceScheduleFrequency = "monthly"
	ContractListResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractListResponseDataInitialUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractListResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractListResponseDataInitialUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractListResponseDataInitialUsageFilter struct {
	Current ContractListResponseDataInitialUsageFilterCurrent  `json:"current,required"`
	Initial ContractListResponseDataInitialUsageFilterInitial  `json:"initial,required"`
	Updates []ContractListResponseDataInitialUsageFilterUpdate `json:"updates,required"`
	JSON    contractListResponseDataInitialUsageFilterJSON
}

// contractListResponseDataInitialUsageFilterJSON contains the JSON metadata for
// the struct [ContractListResponseDataInitialUsageFilter]
type contractListResponseDataInitialUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListResponseDataInitialUsageFilterCurrentJSON
}

// contractListResponseDataInitialUsageFilterCurrentJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialUsageFilterCurrent]
type contractListResponseDataInitialUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractListResponseDataInitialUsageFilterInitialJSON
}

// contractListResponseDataInitialUsageFilterInitialJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialUsageFilterInitial]
type contractListResponseDataInitialUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractListResponseDataInitialUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractListResponseDataInitialUsageFilterUpdateJSON
}

// contractListResponseDataInitialUsageFilterUpdateJSON contains the JSON metadata
// for the struct [ContractListResponseDataInitialUsageFilterUpdate]
type contractListResponseDataInitialUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractListResponseDataInitialUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractAmendResponse struct {
	Data ContractAmendResponseData `json:"data,required"`
	JSON contractAmendResponseJSON
}

// contractAmendResponseJSON contains the JSON metadata for the struct
// [ContractAmendResponse]
type contractAmendResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractAmendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractAmendResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractAmendResponseDataJSON
}

// contractAmendResponseDataJSON contains the JSON metadata for the struct
// [ContractAmendResponseData]
type contractAmendResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractAmendResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponse struct {
	Data ContractGetResponseData `json:"data,required"`
	JSON contractGetResponseJSON
}

// contractGetResponseJSON contains the JSON metadata for the struct
// [ContractGetResponse]
type contractGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseData struct {
	ID         string                             `json:"id,required" format:"uuid"`
	Amendments []ContractGetResponseDataAmendment `json:"amendments,required"`
	Current    ContractGetResponseDataCurrent     `json:"current,required"`
	CustomerID string                             `json:"customer_id,required" format:"uuid"`
	Initial    ContractGetResponseDataInitial     `json:"initial,required"`
	JSON       contractGetResponseDataJSON
}

// contractGetResponseDataJSON contains the JSON metadata for the struct
// [ContractGetResponseData]
type contractGetResponseDataJSON struct {
	ID          apijson.Field
	Amendments  apijson.Field
	Current     apijson.Field
	CustomerID  apijson.Field
	Initial     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendment struct {
	ID                      string                                             `json:"id,required" format:"uuid"`
	Commits                 []ContractGetResponseDataAmendmentsCommit          `json:"commits,required"`
	CreatedAt               time.Time                                          `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                             `json:"created_by,required"`
	Discounts               []ContractGetResponseDataAmendmentsDiscount        `json:"discounts,required"`
	Overrides               []ContractGetResponseDataAmendmentsOverride        `json:"overrides,required"`
	ResellerRoyalties       []ContractGetResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractGetResponseDataAmendmentsScheduledCharge `json:"scheduled_charges,required"`
	StartingAt              time.Time                                          `json:"starting_at,required" format:"date-time"`
	NetsuiteSalesOrderID    string                                             `json:"netsuite_sales_order_id"`
	SalesforceOpportunityID string                                             `json:"salesforce_opportunity_id"`
	JSON                    contractGetResponseDataAmendmentJSON
}

// contractGetResponseDataAmendmentJSON contains the JSON metadata for the struct
// [ContractGetResponseDataAmendment]
type contractGetResponseDataAmendmentJSON struct {
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

func (r *ContractGetResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsCommit struct {
	ID      string                                          `json:"id,required" format:"uuid"`
	Product ContractGetResponseDataAmendmentsCommitsProduct `json:"product,required"`
	Type    ContractGetResponseDataAmendmentsCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractGetResponseDataAmendmentsCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractGetResponseDataAmendmentsCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractGetResponseDataAmendmentsCommitsLedger       `json:"ledger"`
	Name                 string                                                 `json:"name"`
	NetsuiteSalesOrderID string                                                 `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractGetResponseDataAmendmentsCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                                `json:"rollover_fraction"`
	JSON                 contractGetResponseDataAmendmentsCommitJSON
}

// contractGetResponseDataAmendmentsCommitJSON contains the JSON metadata for the
// struct [ContractGetResponseDataAmendmentsCommit]
type contractGetResponseDataAmendmentsCommitJSON struct {
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

func (r *ContractGetResponseDataAmendmentsCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataAmendmentsCommitsProductJSON
}

// contractGetResponseDataAmendmentsCommitsProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsCommitsProduct]
type contractGetResponseDataAmendmentsCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsCommitsType string

const (
	ContractGetResponseDataAmendmentsCommitsTypePrepaid  ContractGetResponseDataAmendmentsCommitsType = "PREPAID"
	ContractGetResponseDataAmendmentsCommitsTypePostpaid ContractGetResponseDataAmendmentsCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractGetResponseDataAmendmentsCommitsAccessSchedule struct {
	ScheduleItems []ContractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractGetResponseDataAmendmentsCommitsAccessScheduleJSON
}

// contractGetResponseDataAmendmentsCommitsAccessScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataAmendmentsCommitsAccessSchedule]
type contractGetResponseDataAmendmentsCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON
}

// contractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItem]
type contractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractGetResponseDataAmendmentsCommitsInvoiceSchedule struct {
	ScheduleItems []ContractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataAmendmentsCommitsInvoiceScheduleJSON
}

// contractGetResponseDataAmendmentsCommitsInvoiceScheduleJSON contains the JSON
// metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsInvoiceSchedule]
type contractGetResponseDataAmendmentsCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON
}

// contractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem]
type contractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or
// [ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractGetResponseDataAmendmentsCommitsLedger interface {
	implementsContractGetResponseDataAmendmentsCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractGetResponseDataAmendmentsCommitsLedger)(nil)).Elem(), "")
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                                `json:"amount,required"`
	SegmentID string                                                                                 `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                              `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                             `json:"amount,required"`
	InvoiceID string                                                                                              `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                            `json:"amount,required"`
	NewContractID string                                                                             `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type          ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                              `json:"amount,required"`
	SegmentID string                                                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	InvoiceID string                                                                             `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                            `json:"amount,required"`
	InvoiceID string                                                                             `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                          `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractGetResponseDataAmendmentsCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                   `json:"amount,required"`
	Timestamp time.Time                                                                                 `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                              `json:"amount,required"`
	InvoiceID string                                                                                               `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                            `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                           `json:"amount,required"`
	InvoiceID string                                                                            `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractGetResponseDataAmendmentsCommitsLedger() {
}

type ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractGetResponseDataAmendmentsCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractGetResponseDataAmendmentsCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractGetResponseDataAmendmentsCommitsRolledOverFromJSON
}

// contractGetResponseDataAmendmentsCommitsRolledOverFromJSON contains the JSON
// metadata for the struct [ContractGetResponseDataAmendmentsCommitsRolledOverFrom]
type contractGetResponseDataAmendmentsCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsDiscount struct {
	ID                   string                                             `json:"id,required" format:"uuid"`
	Product              ContractGetResponseDataAmendmentsDiscountsProduct  `json:"product,required"`
	Schedule             ContractGetResponseDataAmendmentsDiscountsSchedule `json:"schedule,required"`
	Name                 string                                             `json:"name"`
	NetsuiteSalesOrderID string                                             `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataAmendmentsDiscountJSON
}

// contractGetResponseDataAmendmentsDiscountJSON contains the JSON metadata for the
// struct [ContractGetResponseDataAmendmentsDiscount]
type contractGetResponseDataAmendmentsDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataAmendmentsDiscountsProductJSON
}

// contractGetResponseDataAmendmentsDiscountsProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsDiscountsProduct]
type contractGetResponseDataAmendmentsDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsDiscountsSchedule struct {
	ScheduleItems []ContractGetResponseDataAmendmentsDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataAmendmentsDiscountsScheduleJSON
}

// contractGetResponseDataAmendmentsDiscountsScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataAmendmentsDiscountsSchedule]
type contractGetResponseDataAmendmentsDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataAmendmentsDiscountsScheduleScheduleItemJSON
}

// contractGetResponseDataAmendmentsDiscountsScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractGetResponseDataAmendmentsDiscountsScheduleScheduleItem]
type contractGetResponseDataAmendmentsDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsOverride struct {
	ID            string                                                  `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                               `json:"starting_at,required" format:"date-time"`
	Type          ContractGetResponseDataAmendmentsOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                               `json:"ending_before" format:"date-time"`
	Entitled      bool                                                    `json:"entitled"`
	Multiplier    float64                                                 `json:"multiplier"`
	OverwriteRate ContractGetResponseDataAmendmentsOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractGetResponseDataAmendmentsOverridesProduct       `json:"product"`
	Tags          []string                                                `json:"tags"`
	JSON          contractGetResponseDataAmendmentsOverrideJSON
}

// contractGetResponseDataAmendmentsOverrideJSON contains the JSON metadata for the
// struct [ContractGetResponseDataAmendmentsOverride]
type contractGetResponseDataAmendmentsOverrideJSON struct {
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

func (r *ContractGetResponseDataAmendmentsOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsOverridesType string

const (
	ContractGetResponseDataAmendmentsOverridesTypeOverwrite  ContractGetResponseDataAmendmentsOverridesType = "OVERWRITE"
	ContractGetResponseDataAmendmentsOverridesTypeMultiplier ContractGetResponseDataAmendmentsOverridesType = "MULTIPLIER"
)

type ContractGetResponseDataAmendmentsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                         `json:"price,required"`
	RateType ContractGetResponseDataAmendmentsOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractGetResponseDataAmendmentsOverridesOverwriteRateJSON
}

// contractGetResponseDataAmendmentsOverridesOverwriteRateJSON contains the JSON
// metadata for the struct
// [ContractGetResponseDataAmendmentsOverridesOverwriteRate]
type contractGetResponseDataAmendmentsOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsOverridesOverwriteRateRateType string

const (
	ContractGetResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractGetResponseDataAmendmentsOverridesOverwriteRateRateType = "FLAT"
	ContractGetResponseDataAmendmentsOverridesOverwriteRateRateTypeFlat       ContractGetResponseDataAmendmentsOverridesOverwriteRateRateType = "flat"
	ContractGetResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractGetResponseDataAmendmentsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractGetResponseDataAmendmentsOverridesOverwriteRateRateTypePercentage ContractGetResponseDataAmendmentsOverridesOverwriteRateRateType = "percentage"
)

type ContractGetResponseDataAmendmentsOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataAmendmentsOverridesProductJSON
}

// contractGetResponseDataAmendmentsOverridesProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsOverridesProduct]
type contractGetResponseDataAmendmentsOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                         `json:"aws_account_number"`
	AwsOfferID            string                                                         `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                         `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                      `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                        `json:"fraction"`
	GcpAccountID          string                                                         `json:"gcp_account_id"`
	GcpOfferID            string                                                         `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                         `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                        `json:"reseller_contract_value"`
	StartingAt            time.Time                                                      `json:"starting_at" format:"date-time"`
	JSON                  contractGetResponseDataAmendmentsResellerRoyaltyJSON
}

// contractGetResponseDataAmendmentsResellerRoyaltyJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsResellerRoyalty]
type contractGetResponseDataAmendmentsResellerRoyaltyJSON struct {
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

func (r *ContractGetResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAws ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
)

type ContractGetResponseDataAmendmentsScheduledCharge struct {
	ID       string                                                    `json:"id,required" format:"uuid"`
	Product  ContractGetResponseDataAmendmentsScheduledChargesProduct  `json:"product,required"`
	Schedule ContractGetResponseDataAmendmentsScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataAmendmentsScheduledChargeJSON
}

// contractGetResponseDataAmendmentsScheduledChargeJSON contains the JSON metadata
// for the struct [ContractGetResponseDataAmendmentsScheduledCharge]
type contractGetResponseDataAmendmentsScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataAmendmentsScheduledChargesProductJSON
}

// contractGetResponseDataAmendmentsScheduledChargesProductJSON contains the JSON
// metadata for the struct
// [ContractGetResponseDataAmendmentsScheduledChargesProduct]
type contractGetResponseDataAmendmentsScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsScheduledChargesSchedule struct {
	ScheduleItems []ContractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataAmendmentsScheduledChargesScheduleJSON
}

// contractGetResponseDataAmendmentsScheduledChargesScheduleJSON contains the JSON
// metadata for the struct
// [ContractGetResponseDataAmendmentsScheduledChargesSchedule]
type contractGetResponseDataAmendmentsScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
}

// contractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItem]
type contractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataAmendmentsScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrent struct {
	Commits                 []ContractGetResponseDataCurrentCommit             `json:"commits,required"`
	CreatedAt               time.Time                                          `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                             `json:"created_by,required"`
	Discounts               []ContractGetResponseDataCurrentDiscount           `json:"discounts,required"`
	Overrides               []ContractGetResponseDataCurrentOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractGetResponseDataCurrentResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractGetResponseDataCurrentScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                          `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractGetResponseDataCurrentTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractGetResponseDataCurrentUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                          `json:"ending_before" format:"date-time"`
	Name                    string                                             `json:"name"`
	NetPaymentTermsDays     float64                                            `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                             `json:"netsuite_sales_order_id"`
	RateCardID              string                                             `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                             `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                            `json:"total_contract_value"`
	UsageFilter             ContractGetResponseDataCurrentUsageFilter          `json:"usage_filter"`
	JSON                    contractGetResponseDataCurrentJSON
}

// contractGetResponseDataCurrentJSON contains the JSON metadata for the struct
// [ContractGetResponseDataCurrent]
type contractGetResponseDataCurrentJSON struct {
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

func (r *ContractGetResponseDataCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentCommit struct {
	ID      string                                       `json:"id,required" format:"uuid"`
	Product ContractGetResponseDataCurrentCommitsProduct `json:"product,required"`
	Type    ContractGetResponseDataCurrentCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractGetResponseDataCurrentCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractGetResponseDataCurrentCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractGetResponseDataCurrentCommitsLedger       `json:"ledger"`
	Name                 string                                              `json:"name"`
	NetsuiteSalesOrderID string                                              `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractGetResponseDataCurrentCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                             `json:"rollover_fraction"`
	JSON                 contractGetResponseDataCurrentCommitJSON
}

// contractGetResponseDataCurrentCommitJSON contains the JSON metadata for the
// struct [ContractGetResponseDataCurrentCommit]
type contractGetResponseDataCurrentCommitJSON struct {
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

func (r *ContractGetResponseDataCurrentCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataCurrentCommitsProductJSON
}

// contractGetResponseDataCurrentCommitsProductJSON contains the JSON metadata for
// the struct [ContractGetResponseDataCurrentCommitsProduct]
type contractGetResponseDataCurrentCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentCommitsType string

const (
	ContractGetResponseDataCurrentCommitsTypePrepaid  ContractGetResponseDataCurrentCommitsType = "PREPAID"
	ContractGetResponseDataCurrentCommitsTypePostpaid ContractGetResponseDataCurrentCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractGetResponseDataCurrentCommitsAccessSchedule struct {
	ScheduleItems []ContractGetResponseDataCurrentCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractGetResponseDataCurrentCommitsAccessScheduleJSON
}

// contractGetResponseDataCurrentCommitsAccessScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentCommitsAccessSchedule]
type contractGetResponseDataCurrentCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractGetResponseDataCurrentCommitsAccessScheduleScheduleItemJSON
}

// contractGetResponseDataCurrentCommitsAccessScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsAccessScheduleScheduleItem]
type contractGetResponseDataCurrentCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractGetResponseDataCurrentCommitsInvoiceSchedule struct {
	ScheduleItems []ContractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataCurrentCommitsInvoiceScheduleJSON
}

// contractGetResponseDataCurrentCommitsInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentCommitsInvoiceSchedule]
type contractGetResponseDataCurrentCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON
}

// contractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItem]
type contractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or [ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractGetResponseDataCurrentCommitsLedger interface {
	implementsContractGetResponseDataCurrentCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractGetResponseDataCurrentCommitsLedger)(nil)).Elem(), "")
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                             `json:"amount,required"`
	SegmentID string                                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                          `json:"amount,required"`
	InvoiceID string                                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                         `json:"amount,required"`
	NewContractID string                                                                          `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type          ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                           `json:"amount,required"`
	SegmentID string                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	InvoiceID string                                                                          `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	InvoiceID string                                                                          `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractGetResponseDataCurrentCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                `json:"amount,required"`
	Timestamp time.Time                                                                              `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	InvoiceID string                                                                                            `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                        `json:"amount,required"`
	InvoiceID string                                                                         `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractGetResponseDataCurrentCommitsLedger() {
}

type ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractGetResponseDataCurrentCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractGetResponseDataCurrentCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractGetResponseDataCurrentCommitsRolledOverFromJSON
}

// contractGetResponseDataCurrentCommitsRolledOverFromJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentCommitsRolledOverFrom]
type contractGetResponseDataCurrentCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentDiscount struct {
	ID                   string                                          `json:"id,required" format:"uuid"`
	Product              ContractGetResponseDataCurrentDiscountsProduct  `json:"product,required"`
	Schedule             ContractGetResponseDataCurrentDiscountsSchedule `json:"schedule,required"`
	Name                 string                                          `json:"name"`
	NetsuiteSalesOrderID string                                          `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataCurrentDiscountJSON
}

// contractGetResponseDataCurrentDiscountJSON contains the JSON metadata for the
// struct [ContractGetResponseDataCurrentDiscount]
type contractGetResponseDataCurrentDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataCurrentDiscountsProductJSON
}

// contractGetResponseDataCurrentDiscountsProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataCurrentDiscountsProduct]
type contractGetResponseDataCurrentDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentDiscountsSchedule struct {
	ScheduleItems []ContractGetResponseDataCurrentDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataCurrentDiscountsScheduleJSON
}

// contractGetResponseDataCurrentDiscountsScheduleJSON contains the JSON metadata
// for the struct [ContractGetResponseDataCurrentDiscountsSchedule]
type contractGetResponseDataCurrentDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataCurrentDiscountsScheduleScheduleItemJSON
}

// contractGetResponseDataCurrentDiscountsScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractGetResponseDataCurrentDiscountsScheduleScheduleItem]
type contractGetResponseDataCurrentDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentOverride struct {
	ID            string                                               `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                            `json:"starting_at,required" format:"date-time"`
	Type          ContractGetResponseDataCurrentOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                            `json:"ending_before" format:"date-time"`
	Entitled      bool                                                 `json:"entitled"`
	Multiplier    float64                                              `json:"multiplier"`
	OverwriteRate ContractGetResponseDataCurrentOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractGetResponseDataCurrentOverridesProduct       `json:"product"`
	Tags          []string                                             `json:"tags"`
	JSON          contractGetResponseDataCurrentOverrideJSON
}

// contractGetResponseDataCurrentOverrideJSON contains the JSON metadata for the
// struct [ContractGetResponseDataCurrentOverride]
type contractGetResponseDataCurrentOverrideJSON struct {
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

func (r *ContractGetResponseDataCurrentOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentOverridesType string

const (
	ContractGetResponseDataCurrentOverridesTypeOverwrite  ContractGetResponseDataCurrentOverridesType = "OVERWRITE"
	ContractGetResponseDataCurrentOverridesTypeMultiplier ContractGetResponseDataCurrentOverridesType = "MULTIPLIER"
)

type ContractGetResponseDataCurrentOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                      `json:"price,required"`
	RateType ContractGetResponseDataCurrentOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractGetResponseDataCurrentOverridesOverwriteRateJSON
}

// contractGetResponseDataCurrentOverridesOverwriteRateJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentOverridesOverwriteRate]
type contractGetResponseDataCurrentOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentOverridesOverwriteRateRateType string

const (
	ContractGetResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractGetResponseDataCurrentOverridesOverwriteRateRateType = "FLAT"
	ContractGetResponseDataCurrentOverridesOverwriteRateRateTypeFlat       ContractGetResponseDataCurrentOverridesOverwriteRateRateType = "flat"
	ContractGetResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractGetResponseDataCurrentOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractGetResponseDataCurrentOverridesOverwriteRateRateTypePercentage ContractGetResponseDataCurrentOverridesOverwriteRateRateType = "percentage"
)

type ContractGetResponseDataCurrentOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataCurrentOverridesProductJSON
}

// contractGetResponseDataCurrentOverridesProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataCurrentOverridesProduct]
type contractGetResponseDataCurrentOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentResellerRoyalty struct {
	Fraction              float64                                                     `json:"fraction,required"`
	NetsuiteResellerID    string                                                      `json:"netsuite_reseller_id,required"`
	ResellerType          ContractGetResponseDataCurrentResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                   `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                      `json:"aws_account_number"`
	AwsOfferID            string                                                      `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                      `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                   `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                      `json:"gcp_account_id"`
	GcpOfferID            string                                                      `json:"gcp_offer_id"`
	ResellerContractValue float64                                                     `json:"reseller_contract_value"`
	JSON                  contractGetResponseDataCurrentResellerRoyaltyJSON
}

// contractGetResponseDataCurrentResellerRoyaltyJSON contains the JSON metadata for
// the struct [ContractGetResponseDataCurrentResellerRoyalty]
type contractGetResponseDataCurrentResellerRoyaltyJSON struct {
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

func (r *ContractGetResponseDataCurrentResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentResellerRoyaltiesResellerType string

const (
	ContractGetResponseDataCurrentResellerRoyaltiesResellerTypeAws ContractGetResponseDataCurrentResellerRoyaltiesResellerType = "AWS"
	ContractGetResponseDataCurrentResellerRoyaltiesResellerTypeGcp ContractGetResponseDataCurrentResellerRoyaltiesResellerType = "GCP"
)

type ContractGetResponseDataCurrentScheduledCharge struct {
	ID       string                                                 `json:"id,required" format:"uuid"`
	Product  ContractGetResponseDataCurrentScheduledChargesProduct  `json:"product,required"`
	Schedule ContractGetResponseDataCurrentScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataCurrentScheduledChargeJSON
}

// contractGetResponseDataCurrentScheduledChargeJSON contains the JSON metadata for
// the struct [ContractGetResponseDataCurrentScheduledCharge]
type contractGetResponseDataCurrentScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataCurrentScheduledChargesProductJSON
}

// contractGetResponseDataCurrentScheduledChargesProductJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentScheduledChargesProduct]
type contractGetResponseDataCurrentScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentScheduledChargesSchedule struct {
	ScheduleItems []ContractGetResponseDataCurrentScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataCurrentScheduledChargesScheduleJSON
}

// contractGetResponseDataCurrentScheduledChargesScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentScheduledChargesSchedule]
type contractGetResponseDataCurrentScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataCurrentScheduledChargesScheduleScheduleItemJSON
}

// contractGetResponseDataCurrentScheduledChargesScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataCurrentScheduledChargesScheduleScheduleItem]
type contractGetResponseDataCurrentScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentTransition struct {
	FromContractID string                                        `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                        `json:"to_contract_id,required" format:"uuid"`
	Type           ContractGetResponseDataCurrentTransitionsType `json:"type,required"`
	JSON           contractGetResponseDataCurrentTransitionJSON
}

// contractGetResponseDataCurrentTransitionJSON contains the JSON metadata for the
// struct [ContractGetResponseDataCurrentTransition]
type contractGetResponseDataCurrentTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentTransitionsType string

const (
	ContractGetResponseDataCurrentTransitionsTypeSupersede ContractGetResponseDataCurrentTransitionsType = "SUPERSEDE"
	ContractGetResponseDataCurrentTransitionsTypeRenewal   ContractGetResponseDataCurrentTransitionsType = "RENEWAL"
)

type ContractGetResponseDataCurrentUsageInvoiceSchedule struct {
	Frequency ContractGetResponseDataCurrentUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractGetResponseDataCurrentUsageInvoiceScheduleJSON
}

// contractGetResponseDataCurrentUsageInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataCurrentUsageInvoiceSchedule]
type contractGetResponseDataCurrentUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentUsageInvoiceScheduleFrequency string

const (
	ContractGetResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractGetResponseDataCurrentUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractGetResponseDataCurrentUsageInvoiceScheduleFrequencyMonthly   ContractGetResponseDataCurrentUsageInvoiceScheduleFrequency = "monthly"
	ContractGetResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractGetResponseDataCurrentUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractGetResponseDataCurrentUsageInvoiceScheduleFrequencyQuarterly ContractGetResponseDataCurrentUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractGetResponseDataCurrentUsageFilter struct {
	Current ContractGetResponseDataCurrentUsageFilterCurrent  `json:"current,required"`
	Initial ContractGetResponseDataCurrentUsageFilterInitial  `json:"initial,required"`
	Updates []ContractGetResponseDataCurrentUsageFilterUpdate `json:"updates,required"`
	JSON    contractGetResponseDataCurrentUsageFilterJSON
}

// contractGetResponseDataCurrentUsageFilterJSON contains the JSON metadata for the
// struct [ContractGetResponseDataCurrentUsageFilter]
type contractGetResponseDataCurrentUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetResponseDataCurrentUsageFilterCurrentJSON
}

// contractGetResponseDataCurrentUsageFilterCurrentJSON contains the JSON metadata
// for the struct [ContractGetResponseDataCurrentUsageFilterCurrent]
type contractGetResponseDataCurrentUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetResponseDataCurrentUsageFilterInitialJSON
}

// contractGetResponseDataCurrentUsageFilterInitialJSON contains the JSON metadata
// for the struct [ContractGetResponseDataCurrentUsageFilterInitial]
type contractGetResponseDataCurrentUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataCurrentUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractGetResponseDataCurrentUsageFilterUpdateJSON
}

// contractGetResponseDataCurrentUsageFilterUpdateJSON contains the JSON metadata
// for the struct [ContractGetResponseDataCurrentUsageFilterUpdate]
type contractGetResponseDataCurrentUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataCurrentUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitial struct {
	Commits                 []ContractGetResponseDataInitialCommit             `json:"commits,required"`
	CreatedAt               time.Time                                          `json:"created_at,required" format:"date-time"`
	CreatedBy               string                                             `json:"created_by,required"`
	Discounts               []ContractGetResponseDataInitialDiscount           `json:"discounts,required"`
	Overrides               []ContractGetResponseDataInitialOverride           `json:"overrides,required"`
	ResellerRoyalties       []ContractGetResponseDataInitialResellerRoyalty    `json:"reseller_royalties,required"`
	ScheduledCharges        []ContractGetResponseDataInitialScheduledCharge    `json:"scheduled_charges,required"`
	StartingAt              time.Time                                          `json:"starting_at,required" format:"date-time"`
	Transitions             []ContractGetResponseDataInitialTransition         `json:"transitions,required"`
	UsageInvoiceSchedule    ContractGetResponseDataInitialUsageInvoiceSchedule `json:"usage_invoice_schedule,required"`
	EndingBefore            time.Time                                          `json:"ending_before" format:"date-time"`
	Name                    string                                             `json:"name"`
	NetPaymentTermsDays     float64                                            `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    string                                             `json:"netsuite_sales_order_id"`
	RateCardID              string                                             `json:"rate_card_id" format:"uuid"`
	SalesforceOpportunityID string                                             `json:"salesforce_opportunity_id"`
	TotalContractValue      float64                                            `json:"total_contract_value"`
	UsageFilter             ContractGetResponseDataInitialUsageFilter          `json:"usage_filter"`
	JSON                    contractGetResponseDataInitialJSON
}

// contractGetResponseDataInitialJSON contains the JSON metadata for the struct
// [ContractGetResponseDataInitial]
type contractGetResponseDataInitialJSON struct {
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

func (r *ContractGetResponseDataInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialCommit struct {
	ID      string                                       `json:"id,required" format:"uuid"`
	Product ContractGetResponseDataInitialCommitsProduct `json:"product,required"`
	Type    ContractGetResponseDataInitialCommitsType    `json:"type,required"`
	// Only valid for "PREPAID" commits: The schedule that the customer will gain
	// access to the credits purposed with this commit.
	AccessSchedule ContractGetResponseDataInitialCommitsAccessSchedule `json:"access_schedule"`
	// Only valid for "POSTPAID" commits: The total that the customer commits to
	// consume. Must be >= 0.
	Amount               float64  `json:"amount"`
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableTags       []string `json:"applicable_tags"`
	Description          string   `json:"description"`
	// Only valid for "PREPAID" commits: The schedule that the customer will be
	// invoiced for this commit.
	InvoiceSchedule ContractGetResponseDataInitialCommitsInvoiceSchedule `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger               []ContractGetResponseDataInitialCommitsLedger       `json:"ledger"`
	Name                 string                                              `json:"name"`
	NetsuiteSalesOrderID string                                              `json:"netsuite_sales_order_id"`
	RolledOverFrom       ContractGetResponseDataInitialCommitsRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction     float64                                             `json:"rollover_fraction"`
	JSON                 contractGetResponseDataInitialCommitJSON
}

// contractGetResponseDataInitialCommitJSON contains the JSON metadata for the
// struct [ContractGetResponseDataInitialCommit]
type contractGetResponseDataInitialCommitJSON struct {
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

func (r *ContractGetResponseDataInitialCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialCommitsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataInitialCommitsProductJSON
}

// contractGetResponseDataInitialCommitsProductJSON contains the JSON metadata for
// the struct [ContractGetResponseDataInitialCommitsProduct]
type contractGetResponseDataInitialCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialCommitsType string

const (
	ContractGetResponseDataInitialCommitsTypePrepaid  ContractGetResponseDataInitialCommitsType = "PREPAID"
	ContractGetResponseDataInitialCommitsTypePostpaid ContractGetResponseDataInitialCommitsType = "POSTPAID"
)

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
type ContractGetResponseDataInitialCommitsAccessSchedule struct {
	ScheduleItems []ContractGetResponseDataInitialCommitsAccessScheduleScheduleItem `json:"schedule_items,required"`
	JSON          contractGetResponseDataInitialCommitsAccessScheduleJSON
}

// contractGetResponseDataInitialCommitsAccessScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialCommitsAccessSchedule]
type contractGetResponseDataInitialCommitsAccessScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsAccessSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialCommitsAccessScheduleScheduleItem struct {
	ID           string    `json:"id,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	EndingBefore time.Time `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time `json:"starting_at,required" format:"date-time"`
	JSON         contractGetResponseDataInitialCommitsAccessScheduleScheduleItemJSON
}

// contractGetResponseDataInitialCommitsAccessScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsAccessScheduleScheduleItem]
type contractGetResponseDataInitialCommitsAccessScheduleScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsAccessScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Only valid for "PREPAID" commits: The schedule that the customer will be
// invoiced for this commit.
type ContractGetResponseDataInitialCommitsInvoiceSchedule struct {
	ScheduleItems []ContractGetResponseDataInitialCommitsInvoiceScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataInitialCommitsInvoiceScheduleJSON
}

// contractGetResponseDataInitialCommitsInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialCommitsInvoiceSchedule]
type contractGetResponseDataInitialCommitsInvoiceScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialCommitsInvoiceScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON
}

// contractGetResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsInvoiceScheduleScheduleItem]
type contractGetResponseDataInitialCommitsInvoiceScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsInvoiceScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
// or [ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry].
type ContractGetResponseDataInitialCommitsLedger interface {
	implementsContractGetResponseDataInitialCommitsLedger()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ContractGetResponseDataInitialCommitsLedger)(nil)).Elem(), "")
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                                                             `json:"amount,required"`
	SegmentID string                                                                              `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                           `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart ContractGetResponseDataInitialCommitsLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                          `json:"amount,required"`
	InvoiceID string                                                                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                                        `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction ContractGetResponseDataInitialCommitsLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                                                         `json:"amount,required"`
	NewContractID string                                                                          `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type          ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          contractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover ContractGetResponseDataInitialCommitsLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                                                           `json:"amount,required"`
	SegmentID string                                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration ContractGetResponseDataInitialCommitsLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	InvoiceID string                                                                          `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                                                         `json:"amount,required"`
	InvoiceID string                                                                          `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                                          `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                                       `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited ContractGetResponseDataInitialCommitsLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

type ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                                                `json:"amount,required"`
	Timestamp time.Time                                                                              `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance ContractGetResponseDataInitialCommitsLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

type ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                                                           `json:"amount,required"`
	InvoiceID string                                                                                            `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                                         `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction ContractGetResponseDataInitialCommitsLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

type ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                                                        `json:"amount,required"`
	InvoiceID string                                                                         `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                                                      `json:"timestamp,required" format:"date-time"`
	Type      ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      contractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
}

// contractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON
// contains the JSON metadata for the struct
// [ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry]
type contractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntry) implementsContractGetResponseDataInitialCommitsLedger() {
}

type ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup ContractGetResponseDataInitialCommitsLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

type ContractGetResponseDataInitialCommitsRolledOverFrom struct {
	CommitID   string `json:"commit_id,required" format:"uuid"`
	ContractID string `json:"contract_id,required" format:"uuid"`
	JSON       contractGetResponseDataInitialCommitsRolledOverFromJSON
}

// contractGetResponseDataInitialCommitsRolledOverFromJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialCommitsRolledOverFrom]
type contractGetResponseDataInitialCommitsRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialCommitsRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialDiscount struct {
	ID                   string                                          `json:"id,required" format:"uuid"`
	Product              ContractGetResponseDataInitialDiscountsProduct  `json:"product,required"`
	Schedule             ContractGetResponseDataInitialDiscountsSchedule `json:"schedule,required"`
	Name                 string                                          `json:"name"`
	NetsuiteSalesOrderID string                                          `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataInitialDiscountJSON
}

// contractGetResponseDataInitialDiscountJSON contains the JSON metadata for the
// struct [ContractGetResponseDataInitialDiscount]
type contractGetResponseDataInitialDiscountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialDiscountsProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataInitialDiscountsProductJSON
}

// contractGetResponseDataInitialDiscountsProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataInitialDiscountsProduct]
type contractGetResponseDataInitialDiscountsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialDiscountsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialDiscountsSchedule struct {
	ScheduleItems []ContractGetResponseDataInitialDiscountsScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataInitialDiscountsScheduleJSON
}

// contractGetResponseDataInitialDiscountsScheduleJSON contains the JSON metadata
// for the struct [ContractGetResponseDataInitialDiscountsSchedule]
type contractGetResponseDataInitialDiscountsScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialDiscountsSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialDiscountsScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataInitialDiscountsScheduleScheduleItemJSON
}

// contractGetResponseDataInitialDiscountsScheduleScheduleItemJSON contains the
// JSON metadata for the struct
// [ContractGetResponseDataInitialDiscountsScheduleScheduleItem]
type contractGetResponseDataInitialDiscountsScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialDiscountsScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialOverride struct {
	ID            string                                               `json:"id,required" format:"uuid"`
	StartingAt    time.Time                                            `json:"starting_at,required" format:"date-time"`
	Type          ContractGetResponseDataInitialOverridesType          `json:"type,required"`
	EndingBefore  time.Time                                            `json:"ending_before" format:"date-time"`
	Entitled      bool                                                 `json:"entitled"`
	Multiplier    float64                                              `json:"multiplier"`
	OverwriteRate ContractGetResponseDataInitialOverridesOverwriteRate `json:"overwrite_rate"`
	Product       ContractGetResponseDataInitialOverridesProduct       `json:"product"`
	Tags          []string                                             `json:"tags"`
	JSON          contractGetResponseDataInitialOverrideJSON
}

// contractGetResponseDataInitialOverrideJSON contains the JSON metadata for the
// struct [ContractGetResponseDataInitialOverride]
type contractGetResponseDataInitialOverrideJSON struct {
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

func (r *ContractGetResponseDataInitialOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialOverridesType string

const (
	ContractGetResponseDataInitialOverridesTypeOverwrite  ContractGetResponseDataInitialOverridesType = "OVERWRITE"
	ContractGetResponseDataInitialOverridesTypeMultiplier ContractGetResponseDataInitialOverridesType = "MULTIPLIER"
)

type ContractGetResponseDataInitialOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64                                                      `json:"price,required"`
	RateType ContractGetResponseDataInitialOverridesOverwriteRateRateType `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool `json:"use_list_prices"`
	JSON          contractGetResponseDataInitialOverridesOverwriteRateJSON
}

// contractGetResponseDataInitialOverridesOverwriteRateJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialOverridesOverwriteRate]
type contractGetResponseDataInitialOverridesOverwriteRateJSON struct {
	Price         apijson.Field
	RateType      apijson.Field
	UseListPrices apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialOverridesOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialOverridesOverwriteRateRateType string

const (
	ContractGetResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractGetResponseDataInitialOverridesOverwriteRateRateType = "FLAT"
	ContractGetResponseDataInitialOverridesOverwriteRateRateTypeFlat       ContractGetResponseDataInitialOverridesOverwriteRateRateType = "flat"
	ContractGetResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractGetResponseDataInitialOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractGetResponseDataInitialOverridesOverwriteRateRateTypePercentage ContractGetResponseDataInitialOverridesOverwriteRateRateType = "percentage"
)

type ContractGetResponseDataInitialOverridesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataInitialOverridesProductJSON
}

// contractGetResponseDataInitialOverridesProductJSON contains the JSON metadata
// for the struct [ContractGetResponseDataInitialOverridesProduct]
type contractGetResponseDataInitialOverridesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialOverridesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialResellerRoyalty struct {
	Fraction              float64                                                     `json:"fraction,required"`
	NetsuiteResellerID    string                                                      `json:"netsuite_reseller_id,required"`
	ResellerType          ContractGetResponseDataInitialResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                                   `json:"starting_at,required" format:"date-time"`
	AwsAccountNumber      string                                                      `json:"aws_account_number"`
	AwsOfferID            string                                                      `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                      `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                   `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                      `json:"gcp_account_id"`
	GcpOfferID            string                                                      `json:"gcp_offer_id"`
	ResellerContractValue float64                                                     `json:"reseller_contract_value"`
	JSON                  contractGetResponseDataInitialResellerRoyaltyJSON
}

// contractGetResponseDataInitialResellerRoyaltyJSON contains the JSON metadata for
// the struct [ContractGetResponseDataInitialResellerRoyalty]
type contractGetResponseDataInitialResellerRoyaltyJSON struct {
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

func (r *ContractGetResponseDataInitialResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialResellerRoyaltiesResellerType string

const (
	ContractGetResponseDataInitialResellerRoyaltiesResellerTypeAws ContractGetResponseDataInitialResellerRoyaltiesResellerType = "AWS"
	ContractGetResponseDataInitialResellerRoyaltiesResellerTypeGcp ContractGetResponseDataInitialResellerRoyaltiesResellerType = "GCP"
)

type ContractGetResponseDataInitialScheduledCharge struct {
	ID       string                                                 `json:"id,required" format:"uuid"`
	Product  ContractGetResponseDataInitialScheduledChargesProduct  `json:"product,required"`
	Schedule ContractGetResponseDataInitialScheduledChargesSchedule `json:"schedule,required"`
	// displayed on invoices
	Name                 string `json:"name"`
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	JSON                 contractGetResponseDataInitialScheduledChargeJSON
}

// contractGetResponseDataInitialScheduledChargeJSON contains the JSON metadata for
// the struct [ContractGetResponseDataInitialScheduledCharge]
type contractGetResponseDataInitialScheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialScheduledChargesProduct struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON contractGetResponseDataInitialScheduledChargesProductJSON
}

// contractGetResponseDataInitialScheduledChargesProductJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialScheduledChargesProduct]
type contractGetResponseDataInitialScheduledChargesProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialScheduledChargesProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialScheduledChargesSchedule struct {
	ScheduleItems []ContractGetResponseDataInitialScheduledChargesScheduleScheduleItem `json:"schedule_items"`
	JSON          contractGetResponseDataInitialScheduledChargesScheduleJSON
}

// contractGetResponseDataInitialScheduledChargesScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialScheduledChargesSchedule]
type contractGetResponseDataInitialScheduledChargesScheduleJSON struct {
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialScheduledChargesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialScheduledChargesScheduleScheduleItem struct {
	ID        string    `json:"id,required" format:"uuid"`
	Amount    float64   `json:"amount,required"`
	InvoiceID string    `json:"invoice_id,required" format:"uuid"`
	Quantity  float64   `json:"quantity,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	UnitPrice float64   `json:"unit_price,required"`
	JSON      contractGetResponseDataInitialScheduledChargesScheduleScheduleItemJSON
}

// contractGetResponseDataInitialScheduledChargesScheduleScheduleItemJSON contains
// the JSON metadata for the struct
// [ContractGetResponseDataInitialScheduledChargesScheduleScheduleItem]
type contractGetResponseDataInitialScheduledChargesScheduleScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialScheduledChargesScheduleScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialTransition struct {
	FromContractID string                                        `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                        `json:"to_contract_id,required" format:"uuid"`
	Type           ContractGetResponseDataInitialTransitionsType `json:"type,required"`
	JSON           contractGetResponseDataInitialTransitionJSON
}

// contractGetResponseDataInitialTransitionJSON contains the JSON metadata for the
// struct [ContractGetResponseDataInitialTransition]
type contractGetResponseDataInitialTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialTransitionsType string

const (
	ContractGetResponseDataInitialTransitionsTypeSupersede ContractGetResponseDataInitialTransitionsType = "SUPERSEDE"
	ContractGetResponseDataInitialTransitionsTypeRenewal   ContractGetResponseDataInitialTransitionsType = "RENEWAL"
)

type ContractGetResponseDataInitialUsageInvoiceSchedule struct {
	Frequency ContractGetResponseDataInitialUsageInvoiceScheduleFrequency `json:"frequency,required"`
	JSON      contractGetResponseDataInitialUsageInvoiceScheduleJSON
}

// contractGetResponseDataInitialUsageInvoiceScheduleJSON contains the JSON
// metadata for the struct [ContractGetResponseDataInitialUsageInvoiceSchedule]
type contractGetResponseDataInitialUsageInvoiceScheduleJSON struct {
	Frequency   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialUsageInvoiceSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialUsageInvoiceScheduleFrequency string

const (
	ContractGetResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractGetResponseDataInitialUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractGetResponseDataInitialUsageInvoiceScheduleFrequencyMonthly   ContractGetResponseDataInitialUsageInvoiceScheduleFrequency = "monthly"
	ContractGetResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractGetResponseDataInitialUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractGetResponseDataInitialUsageInvoiceScheduleFrequencyQuarterly ContractGetResponseDataInitialUsageInvoiceScheduleFrequency = "quarterly"
)

type ContractGetResponseDataInitialUsageFilter struct {
	Current ContractGetResponseDataInitialUsageFilterCurrent  `json:"current,required"`
	Initial ContractGetResponseDataInitialUsageFilterInitial  `json:"initial,required"`
	Updates []ContractGetResponseDataInitialUsageFilterUpdate `json:"updates,required"`
	JSON    contractGetResponseDataInitialUsageFilterJSON
}

// contractGetResponseDataInitialUsageFilterJSON contains the JSON metadata for the
// struct [ContractGetResponseDataInitialUsageFilter]
type contractGetResponseDataInitialUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialUsageFilterCurrent struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetResponseDataInitialUsageFilterCurrentJSON
}

// contractGetResponseDataInitialUsageFilterCurrentJSON contains the JSON metadata
// for the struct [ContractGetResponseDataInitialUsageFilterCurrent]
type contractGetResponseDataInitialUsageFilterCurrentJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialUsageFilterCurrent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialUsageFilterInitial struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at" format:"date-time"`
	JSON        contractGetResponseDataInitialUsageFilterInitialJSON
}

// contractGetResponseDataInitialUsageFilterInitialJSON contains the JSON metadata
// for the struct [ContractGetResponseDataInitialUsageFilterInitial]
type contractGetResponseDataInitialUsageFilterInitialJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialUsageFilterInitial) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractGetResponseDataInitialUsageFilterUpdate struct {
	GroupKey    string    `json:"group_key,required"`
	GroupValues []string  `json:"group_values,required"`
	StartingAt  time.Time `json:"starting_at,required" format:"date-time"`
	JSON        contractGetResponseDataInitialUsageFilterUpdateJSON
}

// contractGetResponseDataInitialUsageFilterUpdateJSON contains the JSON metadata
// for the struct [ContractGetResponseDataInitialUsageFilterUpdate]
type contractGetResponseDataInitialUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractGetResponseDataInitialUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractUpdateEndDateResponse struct {
	Data ContractUpdateEndDateResponseData `json:"data,required"`
	JSON contractUpdateEndDateResponseJSON
}

// contractUpdateEndDateResponseJSON contains the JSON metadata for the struct
// [ContractUpdateEndDateResponse]
type contractUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractUpdateEndDateResponseData struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON contractUpdateEndDateResponseDataJSON
}

// contractUpdateEndDateResponseDataJSON contains the JSON metadata for the struct
// [ContractUpdateEndDateResponseData]
type contractUpdateEndDateResponseDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractUpdateEndDateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContractNewParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive contract start time
	StartingAt param.Field[time.Time]                   `json:"starting_at,required" format:"date-time"`
	Commits    param.Field[[]ContractNewParamsCommit]   `json:"commits"`
	Discounts  param.Field[[]ContractNewParamsDiscount] `json:"discounts"`
	// exclusive contract end time
	EndingBefore            param.Field[time.Time]                             `json:"ending_before" format:"date-time"`
	Name                    param.Field[string]                                `json:"name"`
	NetPaymentTermsDays     param.Field[float64]                               `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID    param.Field[string]                                `json:"netsuite_sales_order_id"`
	Overrides               param.Field[[]ContractNewParamsOverride]           `json:"overrides"`
	RateCardID              param.Field[string]                                `json:"rate_card_id" format:"uuid"`
	ResellerRoyalties       param.Field[[]ContractNewParamsResellerRoyalty]    `json:"reseller_royalties"`
	SalesforceOpportunityID param.Field[string]                                `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractNewParamsScheduledCharge]    `json:"scheduled_charges"`
	TotalContractValue      param.Field[float64]                               `json:"total_contract_value"`
	Transition              param.Field[ContractNewParamsTransition]           `json:"transition"`
	UsageFilter             param.Field[ContractNewParamsUsageFilter]          `json:"usage_filter"`
	UsageInvoiceSchedule    param.Field[ContractNewParamsUsageInvoiceSchedule] `json:"usage_invoice_schedule"`
}

func (r ContractNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommit struct {
	ProductID param.Field[string]                       `json:"product_id,required" format:"uuid"`
	Type      param.Field[ContractNewParamsCommitsType] `json:"type,required"`
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
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractNewParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsType string

const (
	ContractNewParamsCommitsTypePrepaid  ContractNewParamsCommitsType = "PREPAID"
	ContractNewParamsCommitsTypePrepaid  ContractNewParamsCommitsType = "prepaid"
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
	ScheduleItems     param.Field[[]ContractNewParamsCommitsInvoiceScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractNewParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                         `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractNewParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractNewParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractNewParamsDiscountsSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractNewParamsDiscountsSchedule struct {
	RecurringSchedule param.Field[ContractNewParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractNewParamsDiscountsScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractNewParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                    `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractNewParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractNewParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractNewParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractNewParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time]                      `json:"starting_at,required" format:"date-time"`
	Type       param.Field[ContractNewParamsOverridesType] `json:"type,required"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[ContractNewParamsOverridesOverwriteRate] `json:"overwrite_rate"`
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
	ContractNewParamsOverridesTypeOverwrite  ContractNewParamsOverridesType = "OVERWRITE"
	ContractNewParamsOverridesTypeOverwrite  ContractNewParamsOverridesType = "overwrite"
	ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "MULTIPLIER"
	ContractNewParamsOverridesTypeMultiplier ContractNewParamsOverridesType = "multiplier"
)

// Required for OVERWRITE type.
type ContractNewParamsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    param.Field[float64]                                         `json:"price,required"`
	RateType param.Field[ContractNewParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractNewParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsOverridesOverwriteRateRateType string

const (
	ContractNewParamsOverridesOverwriteRateRateTypeFlat       ContractNewParamsOverridesOverwriteRateRateType = "FLAT"
	ContractNewParamsOverridesOverwriteRateRateTypeFlat       ContractNewParamsOverridesOverwriteRateRateType = "flat"
	ContractNewParamsOverridesOverwriteRateRateTypePercentage ContractNewParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractNewParamsOverridesOverwriteRateRateTypePercentage ContractNewParamsOverridesOverwriteRateRateType = "percentage"
)

type ContractNewParamsResellerRoyalty struct {
	ApplicableProductIDs  param.Field[[]string]                                       `json:"applicable_product_ids,required" format:"uuid"`
	Fraction              param.Field[float64]                                        `json:"fraction,required"`
	NetsuiteResellerID    param.Field[string]                                         `json:"netsuite_reseller_id,required"`
	ResellerType          param.Field[ContractNewParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	StartingAt            param.Field[time.Time]                                      `json:"starting_at,required" format:"date-time"`
	AwsOptions            param.Field[ContractNewParamsResellerRoyaltiesAwsOptions]   `json:"aws_options"`
	EndingBefore          param.Field[time.Time]                                      `json:"ending_before" format:"date-time"`
	GcpOptions            param.Field[ContractNewParamsResellerRoyaltiesGcpOptions]   `json:"gcp_options"`
	ResellerContractValue param.Field[float64]                                        `json:"reseller_contract_value"`
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
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractNewParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractNewParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractNewParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractNewParamsScheduledChargesSchedule struct {
	RecurringSchedule param.Field[ContractNewParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractNewParamsScheduledChargesScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractNewParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                           `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractNewParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractNewParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsTransition struct {
	FromContractID param.Field[string]                          `json:"from_contract_id,required" format:"uuid"`
	Type           param.Field[ContractNewParamsTransitionType] `json:"type,required"`
}

func (r ContractNewParamsTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractNewParamsTransitionType string

const (
	ContractNewParamsTransitionTypeSupersede ContractNewParamsTransitionType = "SUPERSEDE"
	ContractNewParamsTransitionTypeRenewal   ContractNewParamsTransitionType = "RENEWAL"
	ContractNewParamsTransitionTypeSupersede ContractNewParamsTransitionType = "supersede"
	ContractNewParamsTransitionTypeRenewal   ContractNewParamsTransitionType = "renewal"
)

type ContractNewParamsUsageFilter struct {
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at" format:"date-time"`
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
	ContractNewParamsUsageInvoiceScheduleFrequencyMonthly   ContractNewParamsUsageInvoiceScheduleFrequency = "MONTHLY"
	ContractNewParamsUsageInvoiceScheduleFrequencyMonthly   ContractNewParamsUsageInvoiceScheduleFrequency = "monthly"
	ContractNewParamsUsageInvoiceScheduleFrequencyQuarterly ContractNewParamsUsageInvoiceScheduleFrequency = "QUARTERLY"
	ContractNewParamsUsageInvoiceScheduleFrequencyQuarterly ContractNewParamsUsageInvoiceScheduleFrequency = "quarterly"
)

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
	StartingAt              param.Field[time.Time]                            `json:"starting_at,required" format:"date-time"`
	Commits                 param.Field[[]ContractAmendParamsCommit]          `json:"commits"`
	Discounts               param.Field[[]ContractAmendParamsDiscount]        `json:"discounts"`
	NetsuiteSalesOrderID    param.Field[string]                               `json:"netsuite_sales_order_id"`
	Overrides               param.Field[[]ContractAmendParamsOverride]        `json:"overrides"`
	ResellerRoyalties       param.Field[[]ContractAmendParamsResellerRoyalty] `json:"reseller_royalties"`
	SalesforceOpportunityID param.Field[string]                               `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]ContractAmendParamsScheduledCharge] `json:"scheduled_charges"`
	TotalContractValue      param.Field[float64]                              `json:"total_contract_value"`
}

func (r ContractAmendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommit struct {
	ProductID param.Field[string]                         `json:"product_id,required" format:"uuid"`
	Type      param.Field[ContractAmendParamsCommitsType] `json:"type,required"`
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
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
}

func (r ContractAmendParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsType string

const (
	ContractAmendParamsCommitsTypePrepaid  ContractAmendParamsCommitsType = "PREPAID"
	ContractAmendParamsCommitsTypePrepaid  ContractAmendParamsCommitsType = "prepaid"
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
	ScheduleItems     param.Field[[]ContractAmendParamsCommitsInvoiceScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractAmendParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                           `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractAmendParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractAmendParamsDiscountsSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractAmendParamsDiscountsSchedule struct {
	RecurringSchedule param.Field[ContractAmendParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractAmendParamsDiscountsScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractAmendParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                      `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractAmendParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time]                        `json:"starting_at,required" format:"date-time"`
	Type       param.Field[ContractAmendParamsOverridesType] `json:"type,required"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[ContractAmendParamsOverridesOverwriteRate] `json:"overwrite_rate"`
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
	ContractAmendParamsOverridesTypeOverwrite  ContractAmendParamsOverridesType = "OVERWRITE"
	ContractAmendParamsOverridesTypeOverwrite  ContractAmendParamsOverridesType = "overwrite"
	ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "MULTIPLIER"
	ContractAmendParamsOverridesTypeMultiplier ContractAmendParamsOverridesType = "multiplier"
)

// Required for OVERWRITE type.
type ContractAmendParamsOverridesOverwriteRate struct {
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    param.Field[float64]                                           `json:"price,required"`
	RateType param.Field[ContractAmendParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices param.Field[bool] `json:"use_list_prices"`
}

func (r ContractAmendParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsOverridesOverwriteRateRateType string

const (
	ContractAmendParamsOverridesOverwriteRateRateTypeFlat       ContractAmendParamsOverridesOverwriteRateRateType = "FLAT"
	ContractAmendParamsOverridesOverwriteRateRateTypeFlat       ContractAmendParamsOverridesOverwriteRateRateType = "flat"
	ContractAmendParamsOverridesOverwriteRateRateTypePercentage ContractAmendParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	ContractAmendParamsOverridesOverwriteRateRateTypePercentage ContractAmendParamsOverridesOverwriteRateRateType = "percentage"
)

type ContractAmendParamsResellerRoyalty struct {
	ResellerType         param.Field[ContractAmendParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	ApplicableProductIDs param.Field[[]string]                                         `json:"applicable_product_ids" format:"uuid"`
	AwsOptions           param.Field[ContractAmendParamsResellerRoyaltiesAwsOptions]   `json:"aws_options"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Field[time.Time]                                      `json:"ending_before" format:"date-time"`
	Fraction              param.Field[float64]                                        `json:"fraction"`
	GcpOptions            param.Field[ContractAmendParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	NetsuiteResellerID    param.Field[string]                                         `json:"netsuite_reseller_id"`
	ResellerContractValue param.Field[float64]                                        `json:"reseller_contract_value"`
	StartingAt            param.Field[time.Time]                                      `json:"starting_at" format:"date-time"`
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
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r ContractAmendParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r ContractAmendParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule
	Schedule param.Field[ContractAmendParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name                 param.Field[string] `json:"name"`
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r ContractAmendParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule
type ContractAmendParamsScheduledChargesSchedule struct {
	RecurringSchedule param.Field[ContractAmendParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	ScheduleItems     param.Field[[]ContractAmendParamsScheduledChargesScheduleScheduleItem]    `json:"schedule_items"`
}

func (r ContractAmendParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContractAmendParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                             `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
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
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "divided_rounded"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "each"
)

type ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "monthly"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "quarterly"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "semi_annual"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
	ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "annual"
)

// Either provide amount or provide both unit_price and quantity
type ContractAmendParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	Amount    param.Field[float64]   `json:"amount"`
	Quantity  param.Field[float64]   `json:"quantity"`
	UnitPrice param.Field[float64]   `json:"unit_price"`
}

func (r ContractAmendParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

type ContractSetUsageFilterParams struct {
	ContractID  param.Field[string]    `json:"contract_id,required" format:"uuid"`
	CustomerID  param.Field[string]    `json:"customer_id,required" format:"uuid"`
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
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
