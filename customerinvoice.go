// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/pagination"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// CustomerInvoiceService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerInvoiceService] method instead.
type CustomerInvoiceService struct {
	Options []option.RequestOption
}

// NewCustomerInvoiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerInvoiceService(opts ...option.RequestOption) (r *CustomerInvoiceService) {
	r = &CustomerInvoiceService{}
	r.Options = opts
	return
}

// Fetch a specific invoice for a given customer.
func (r *CustomerInvoiceService) Get(ctx context.Context, customerID string, invoiceID string, query CustomerInvoiceGetParams, opts ...option.RequestOption) (res *CustomerInvoiceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/invoices/%s", customerID, invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List all invoices for a given customer, optionally filtered by status, date
// range, and/or credit type.
func (r *CustomerInvoiceService) List(ctx context.Context, customerID string, query CustomerInvoiceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/invoices", customerID)
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

// List all invoices for a given customer, optionally filtered by status, date
// range, and/or credit type.
func (r *CustomerInvoiceService) ListAutoPaging(ctx context.Context, customerID string, query CustomerInvoiceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Invoice] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Add a one time charge to the specified invoice
func (r *CustomerInvoiceService) AddCharge(ctx context.Context, customerID string, body CustomerInvoiceAddChargeParams, opts ...option.RequestOption) (res *CustomerInvoiceAddChargeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/addCharge", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Invoice struct {
	ID          string            `json:"id,required" format:"uuid"`
	CreditType  shared.CreditType `json:"credit_type,required"`
	CustomerID  string            `json:"customer_id,required" format:"uuid"`
	LineItems   []InvoiceLineItem `json:"line_items,required"`
	Status      string            `json:"status,required"`
	Total       float64           `json:"total,required"`
	Type        string            `json:"type,required"`
	AmendmentID string            `json:"amendment_id" format:"uuid"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus       InvoiceBillableStatus   `json:"billable_status"`
	ContractCustomFields map[string]string       `json:"contract_custom_fields"`
	ContractID           string                  `json:"contract_id" format:"uuid"`
	CorrectionRecord     InvoiceCorrectionRecord `json:"correction_record"`
	// When the invoice was created (UTC). This field is present for correction
	// invoices only.
	CreatedAt            time.Time              `json:"created_at" format:"date-time"`
	CustomFields         map[string]interface{} `json:"custom_fields"`
	CustomerCustomFields map[string]string      `json:"customer_custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                  `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    InvoiceExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []InvoiceInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt            time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays float64   `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string            `json:"netsuite_sales_order_id"`
	PlanCustomFields     map[string]string `json:"plan_custom_fields"`
	PlanID               string            `json:"plan_id" format:"uuid"`
	PlanName             string            `json:"plan_name"`
	// only present for beta contract invoices with reseller royalties
	ResellerRoyalty InvoiceResellerRoyalty `json:"reseller_royalty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time   `json:"start_timestamp" format:"date-time"`
	Subtotal       float64     `json:"subtotal"`
	JSON           invoiceJSON `json:"-"`
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID                      apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	AmendmentID             apijson.Field
	BillableStatus          apijson.Field
	ContractCustomFields    apijson.Field
	ContractID              apijson.Field
	CorrectionRecord        apijson.Field
	CreatedAt               apijson.Field
	CustomFields            apijson.Field
	CustomerCustomFields    apijson.Field
	EndTimestamp            apijson.Field
	ExternalInvoice         apijson.Field
	InvoiceAdjustments      apijson.Field
	IssuedAt                apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	PlanCustomFields        apijson.Field
	PlanID                  apijson.Field
	PlanName                apijson.Field
	ResellerRoyalty         apijson.Field
	SalesforceOpportunityID apijson.Field
	StartTimestamp          apijson.Field
	Subtotal                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItem struct {
	CreditType         shared.CreditType `json:"credit_type,required"`
	Name               string            `json:"name,required"`
	Total              float64           `json:"total,required"`
	CommitCustomFields map[string]string `json:"commit_custom_fields"`
	// only present for beta contract invoices
	CommitID string `json:"commit_id" format:"uuid"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	CommitNetsuiteItemID string `json:"commit_netsuite_item_id"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id"`
	// only present for beta contract invoices
	CommitSegmentID string `json:"commit_segment_id" format:"uuid"`
	// only present for beta contract invoices
	CommitType   string            `json:"commit_type"`
	CustomFields map[string]string `json:"custom_fields"`
	// only present for beta contract invoices
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value,nullable"`
	// only present for beta contract invoices
	IsProrated bool   `json:"is_prorated"`
	Metadata   string `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd time.Time `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart time.Time `json:"netsuite_invoice_billing_start" format:"date-time"`
	// only present for beta contract invoices. This field's availability is dependent
	// on your client's configuration.
	NetsuiteItemID string `json:"netsuite_item_id"`
	// only present for beta contract invoices
	PostpaidCommit InvoiceLineItemsPostpaidCommit `json:"postpaid_commit"`
	// if presentation groups are used, this will contain the values used to break down
	// the line item
	PresentationGroupValues map[string]string `json:"presentation_group_values"`
	// if pricing groups are used, this will contain the values used to calculate the
	// price
	PricingGroupValues              map[string]string `json:"pricing_group_values"`
	ProductCustomFields             map[string]string `json:"product_custom_fields"`
	ProductID                       string            `json:"product_id" format:"uuid"`
	ProductType                     string            `json:"product_type"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields"`
	// only present for beta contract invoices
	ProfessionalServiceID       string                       `json:"professional_service_id" format:"uuid"`
	Quantity                    float64                      `json:"quantity"`
	ResellerType                InvoiceLineItemsResellerType `json:"reseller_type"`
	ScheduledChargeCustomFields map[string]string            `json:"scheduled_charge_custom_fields"`
	// only present for beta contract invoices
	ScheduledChargeID string `json:"scheduled_charge_id" format:"uuid"`
	// only present for beta contract invoices
	StartingAt   time.Time                     `json:"starting_at" format:"date-time"`
	SubLineItems []InvoiceLineItemsSubLineItem `json:"sub_line_items"`
	// only present for beta contract invoices
	UnitPrice float64             `json:"unit_price"`
	JSON      invoiceLineItemJSON `json:"-"`
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	CreditType                      apijson.Field
	Name                            apijson.Field
	Total                           apijson.Field
	CommitCustomFields              apijson.Field
	CommitID                        apijson.Field
	CommitNetsuiteItemID            apijson.Field
	CommitNetsuiteSalesOrderID      apijson.Field
	CommitSegmentID                 apijson.Field
	CommitType                      apijson.Field
	CustomFields                    apijson.Field
	EndingBefore                    apijson.Field
	GroupKey                        apijson.Field
	GroupValue                      apijson.Field
	IsProrated                      apijson.Field
	Metadata                        apijson.Field
	NetsuiteInvoiceBillingEnd       apijson.Field
	NetsuiteInvoiceBillingStart     apijson.Field
	NetsuiteItemID                  apijson.Field
	PostpaidCommit                  apijson.Field
	PresentationGroupValues         apijson.Field
	PricingGroupValues              apijson.Field
	ProductCustomFields             apijson.Field
	ProductID                       apijson.Field
	ProductType                     apijson.Field
	ProfessionalServiceCustomFields apijson.Field
	ProfessionalServiceID           apijson.Field
	Quantity                        apijson.Field
	ResellerType                    apijson.Field
	ScheduledChargeCustomFields     apijson.Field
	ScheduledChargeID               apijson.Field
	StartingAt                      apijson.Field
	SubLineItems                    apijson.Field
	UnitPrice                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemJSON) RawJSON() string {
	return r.raw
}

// only present for beta contract invoices
type InvoiceLineItemsPostpaidCommit struct {
	ID   string                             `json:"id,required" format:"uuid"`
	JSON invoiceLineItemsPostpaidCommitJSON `json:"-"`
}

// invoiceLineItemsPostpaidCommitJSON contains the JSON metadata for the struct
// [InvoiceLineItemsPostpaidCommit]
type invoiceLineItemsPostpaidCommitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsPostpaidCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsPostpaidCommitJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsResellerType string

const (
	InvoiceLineItemsResellerTypeAws           InvoiceLineItemsResellerType = "AWS"
	InvoiceLineItemsResellerTypeAwsProService InvoiceLineItemsResellerType = "AWS_PRO_SERVICE"
	InvoiceLineItemsResellerTypeGcp           InvoiceLineItemsResellerType = "GCP"
	InvoiceLineItemsResellerTypeGcpProService InvoiceLineItemsResellerType = "GCP_PRO_SERVICE"
)

func (r InvoiceLineItemsResellerType) IsKnown() bool {
	switch r {
	case InvoiceLineItemsResellerTypeAws, InvoiceLineItemsResellerTypeAwsProService, InvoiceLineItemsResellerTypeGcp, InvoiceLineItemsResellerTypeGcpProService:
		return true
	}
	return false
}

type InvoiceLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// The end date for the charge (for seats charges only).
	EndDate time.Time `json:"end_date" format:"date-time"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64 `json:"price"`
	// The start date for the charge (for seats charges only).
	StartDate time.Time `json:"start_date" format:"date-time"`
	// when the current tier started and ends (for tiered charges only)
	TierPeriod InvoiceLineItemsSubLineItemsTierPeriod `json:"tier_period"`
	Tiers      []InvoiceLineItemsSubLineItemsTier     `json:"tiers"`
	JSON       invoiceLineItemsSubLineItemJSON        `json:"-"`
}

// invoiceLineItemsSubLineItemJSON contains the JSON metadata for the struct
// [InvoiceLineItemsSubLineItem]
type invoiceLineItemsSubLineItemJSON struct {
	CustomFields  apijson.Field
	Name          apijson.Field
	Quantity      apijson.Field
	Subtotal      apijson.Field
	ChargeID      apijson.Field
	CreditGrantID apijson.Field
	EndDate       apijson.Field
	Price         apijson.Field
	StartDate     apijson.Field
	TierPeriod    apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

// when the current tier started and ends (for tiered charges only)
type InvoiceLineItemsSubLineItemsTierPeriod struct {
	StartingAt   time.Time                                  `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                  `json:"ending_before" format:"date-time"`
	JSON         invoiceLineItemsSubLineItemsTierPeriodJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsTierPeriodJSON contains the JSON metadata for the
// struct [InvoiceLineItemsSubLineItemsTierPeriod]
type invoiceLineItemsSubLineItemsTierPeriodJSON struct {
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsTierPeriodJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64                              `json:"starting_at,required"`
	Subtotal   float64                              `json:"subtotal,required"`
	JSON       invoiceLineItemsSubLineItemsTierJSON `json:"-"`
}

// invoiceLineItemsSubLineItemsTierJSON contains the JSON metadata for the struct
// [InvoiceLineItemsSubLineItemsTier]
type invoiceLineItemsSubLineItemsTierJSON struct {
	Price       apijson.Field
	Quantity    apijson.Field
	StartingAt  apijson.Field
	Subtotal    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineItemsSubLineItemsTierJSON) RawJSON() string {
	return r.raw
}

// This field's availability is dependent on your client's configuration.
type InvoiceBillableStatus string

const (
	InvoiceBillableStatusBillable   InvoiceBillableStatus = "billable"
	InvoiceBillableStatusUnbillable InvoiceBillableStatus = "unbillable"
)

func (r InvoiceBillableStatus) IsKnown() bool {
	switch r {
	case InvoiceBillableStatusBillable, InvoiceBillableStatusUnbillable:
		return true
	}
	return false
}

type InvoiceCorrectionRecord struct {
	CorrectedInvoiceID       string                                          `json:"corrected_invoice_id,required" format:"uuid"`
	Memo                     string                                          `json:"memo,required"`
	Reason                   string                                          `json:"reason,required"`
	CorrectedExternalInvoice InvoiceCorrectionRecordCorrectedExternalInvoice `json:"corrected_external_invoice"`
	JSON                     invoiceCorrectionRecordJSON                     `json:"-"`
}

// invoiceCorrectionRecordJSON contains the JSON metadata for the struct
// [InvoiceCorrectionRecord]
type invoiceCorrectionRecordJSON struct {
	CorrectedInvoiceID       apijson.Field
	Memo                     apijson.Field
	Reason                   apijson.Field
	CorrectedExternalInvoice apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *InvoiceCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCorrectionRecordJSON) RawJSON() string {
	return r.raw
}

type InvoiceCorrectionRecordCorrectedExternalInvoice struct {
	BillingProviderType InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                                             `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                                          `json:"issued_at_timestamp" format:"date-time"`
	JSON                invoiceCorrectionRecordCorrectedExternalInvoiceJSON                `json:"-"`
}

// invoiceCorrectionRecordCorrectedExternalInvoiceJSON contains the JSON metadata
// for the struct [InvoiceCorrectionRecordCorrectedExternalInvoice]
type invoiceCorrectionRecordCorrectedExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceCorrectionRecordCorrectedExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCorrectionRecordCorrectedExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType string

const (
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace   InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "aws_marketplace"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe           InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "stripe"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite         InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "netsuite"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom           InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "custom"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "azure_marketplace"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "quickbooks_online"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday          InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "workday"
	InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace   InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAwsMarketplace, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeStripe, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeNetsuite, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeCustom, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeAzureMarketplace, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeQuickbooksOnline, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeWorkday, InvoiceCorrectionRecordCorrectedExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus string

const (
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft               InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DRAFT"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized           InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "FINALIZED"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid                InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAID"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible       InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid                InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "VOID"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted             InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "DELETED"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed       InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped             InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SKIPPED"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusSent                InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "SENT"
	InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued              InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus = "QUEUED"
)

func (r InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusDraft, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusFinalized, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusPaid, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusUncollectible, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusVoid, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusDeleted, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusPaymentFailed, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusInvalidRequestError, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusSkipped, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusSent, InvoiceCorrectionRecordCorrectedExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type InvoiceExternalInvoice struct {
	BillingProviderType InvoiceExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      InvoiceExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                    `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                 `json:"issued_at_timestamp" format:"date-time"`
	JSON                invoiceExternalInvoiceJSON                `json:"-"`
}

// invoiceExternalInvoiceJSON contains the JSON metadata for the struct
// [InvoiceExternalInvoice]
type invoiceExternalInvoiceJSON struct {
	BillingProviderType apijson.Field
	ExternalStatus      apijson.Field
	InvoiceID           apijson.Field
	IssuedAtTimestamp   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceExternalInvoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceExternalInvoiceBillingProviderType string

const (
	InvoiceExternalInvoiceBillingProviderTypeAwsMarketplace   InvoiceExternalInvoiceBillingProviderType = "aws_marketplace"
	InvoiceExternalInvoiceBillingProviderTypeStripe           InvoiceExternalInvoiceBillingProviderType = "stripe"
	InvoiceExternalInvoiceBillingProviderTypeNetsuite         InvoiceExternalInvoiceBillingProviderType = "netsuite"
	InvoiceExternalInvoiceBillingProviderTypeCustom           InvoiceExternalInvoiceBillingProviderType = "custom"
	InvoiceExternalInvoiceBillingProviderTypeAzureMarketplace InvoiceExternalInvoiceBillingProviderType = "azure_marketplace"
	InvoiceExternalInvoiceBillingProviderTypeQuickbooksOnline InvoiceExternalInvoiceBillingProviderType = "quickbooks_online"
	InvoiceExternalInvoiceBillingProviderTypeWorkday          InvoiceExternalInvoiceBillingProviderType = "workday"
	InvoiceExternalInvoiceBillingProviderTypeGcpMarketplace   InvoiceExternalInvoiceBillingProviderType = "gcp_marketplace"
)

func (r InvoiceExternalInvoiceBillingProviderType) IsKnown() bool {
	switch r {
	case InvoiceExternalInvoiceBillingProviderTypeAwsMarketplace, InvoiceExternalInvoiceBillingProviderTypeStripe, InvoiceExternalInvoiceBillingProviderTypeNetsuite, InvoiceExternalInvoiceBillingProviderTypeCustom, InvoiceExternalInvoiceBillingProviderTypeAzureMarketplace, InvoiceExternalInvoiceBillingProviderTypeQuickbooksOnline, InvoiceExternalInvoiceBillingProviderTypeWorkday, InvoiceExternalInvoiceBillingProviderTypeGcpMarketplace:
		return true
	}
	return false
}

type InvoiceExternalInvoiceExternalStatus string

const (
	InvoiceExternalInvoiceExternalStatusDraft               InvoiceExternalInvoiceExternalStatus = "DRAFT"
	InvoiceExternalInvoiceExternalStatusFinalized           InvoiceExternalInvoiceExternalStatus = "FINALIZED"
	InvoiceExternalInvoiceExternalStatusPaid                InvoiceExternalInvoiceExternalStatus = "PAID"
	InvoiceExternalInvoiceExternalStatusUncollectible       InvoiceExternalInvoiceExternalStatus = "UNCOLLECTIBLE"
	InvoiceExternalInvoiceExternalStatusVoid                InvoiceExternalInvoiceExternalStatus = "VOID"
	InvoiceExternalInvoiceExternalStatusDeleted             InvoiceExternalInvoiceExternalStatus = "DELETED"
	InvoiceExternalInvoiceExternalStatusPaymentFailed       InvoiceExternalInvoiceExternalStatus = "PAYMENT_FAILED"
	InvoiceExternalInvoiceExternalStatusInvalidRequestError InvoiceExternalInvoiceExternalStatus = "INVALID_REQUEST_ERROR"
	InvoiceExternalInvoiceExternalStatusSkipped             InvoiceExternalInvoiceExternalStatus = "SKIPPED"
	InvoiceExternalInvoiceExternalStatusSent                InvoiceExternalInvoiceExternalStatus = "SENT"
	InvoiceExternalInvoiceExternalStatusQueued              InvoiceExternalInvoiceExternalStatus = "QUEUED"
)

func (r InvoiceExternalInvoiceExternalStatus) IsKnown() bool {
	switch r {
	case InvoiceExternalInvoiceExternalStatusDraft, InvoiceExternalInvoiceExternalStatusFinalized, InvoiceExternalInvoiceExternalStatusPaid, InvoiceExternalInvoiceExternalStatusUncollectible, InvoiceExternalInvoiceExternalStatusVoid, InvoiceExternalInvoiceExternalStatusDeleted, InvoiceExternalInvoiceExternalStatusPaymentFailed, InvoiceExternalInvoiceExternalStatusInvalidRequestError, InvoiceExternalInvoiceExternalStatusSkipped, InvoiceExternalInvoiceExternalStatusSent, InvoiceExternalInvoiceExternalStatusQueued:
		return true
	}
	return false
}

type InvoiceInvoiceAdjustment struct {
	CreditType              shared.CreditType            `json:"credit_type,required"`
	Name                    string                       `json:"name,required"`
	Total                   float64                      `json:"total,required"`
	CreditGrantCustomFields map[string]string            `json:"credit_grant_custom_fields"`
	CreditGrantID           string                       `json:"credit_grant_id"`
	JSON                    invoiceInvoiceAdjustmentJSON `json:"-"`
}

// invoiceInvoiceAdjustmentJSON contains the JSON metadata for the struct
// [InvoiceInvoiceAdjustment]
type invoiceInvoiceAdjustmentJSON struct {
	CreditType              apijson.Field
	Name                    apijson.Field
	Total                   apijson.Field
	CreditGrantCustomFields apijson.Field
	CreditGrantID           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvoiceInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceInvoiceAdjustmentJSON) RawJSON() string {
	return r.raw
}

// only present for beta contract invoices with reseller royalties
type InvoiceResellerRoyalty struct {
	Fraction           string                             `json:"fraction,required"`
	NetsuiteResellerID string                             `json:"netsuite_reseller_id,required"`
	ResellerType       InvoiceResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         InvoiceResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         InvoiceResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               invoiceResellerRoyaltyJSON         `json:"-"`
}

// invoiceResellerRoyaltyJSON contains the JSON metadata for the struct
// [InvoiceResellerRoyalty]
type invoiceResellerRoyaltyJSON struct {
	Fraction           apijson.Field
	NetsuiteResellerID apijson.Field
	ResellerType       apijson.Field
	AwsOptions         apijson.Field
	GcpOptions         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type InvoiceResellerRoyaltyResellerType string

const (
	InvoiceResellerRoyaltyResellerTypeAws           InvoiceResellerRoyaltyResellerType = "AWS"
	InvoiceResellerRoyaltyResellerTypeAwsProService InvoiceResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	InvoiceResellerRoyaltyResellerTypeGcp           InvoiceResellerRoyaltyResellerType = "GCP"
	InvoiceResellerRoyaltyResellerTypeGcpProService InvoiceResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

func (r InvoiceResellerRoyaltyResellerType) IsKnown() bool {
	switch r {
	case InvoiceResellerRoyaltyResellerTypeAws, InvoiceResellerRoyaltyResellerTypeAwsProService, InvoiceResellerRoyaltyResellerTypeGcp, InvoiceResellerRoyaltyResellerTypeGcpProService:
		return true
	}
	return false
}

type InvoiceResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string                               `json:"aws_account_number"`
	AwsOfferID          string                               `json:"aws_offer_id"`
	AwsPayerReferenceID string                               `json:"aws_payer_reference_id"`
	JSON                invoiceResellerRoyaltyAwsOptionsJSON `json:"-"`
}

// invoiceResellerRoyaltyAwsOptionsJSON contains the JSON metadata for the struct
// [InvoiceResellerRoyaltyAwsOptions]
type invoiceResellerRoyaltyAwsOptionsJSON struct {
	AwsAccountNumber    apijson.Field
	AwsOfferID          apijson.Field
	AwsPayerReferenceID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceResellerRoyaltyAwsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceResellerRoyaltyAwsOptionsJSON) RawJSON() string {
	return r.raw
}

type InvoiceResellerRoyaltyGcpOptions struct {
	GcpAccountID string                               `json:"gcp_account_id"`
	GcpOfferID   string                               `json:"gcp_offer_id"`
	JSON         invoiceResellerRoyaltyGcpOptionsJSON `json:"-"`
}

// invoiceResellerRoyaltyGcpOptionsJSON contains the JSON metadata for the struct
// [InvoiceResellerRoyaltyGcpOptions]
type invoiceResellerRoyaltyGcpOptionsJSON struct {
	GcpAccountID apijson.Field
	GcpOfferID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceResellerRoyaltyGcpOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceResellerRoyaltyGcpOptionsJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetResponse struct {
	Data Invoice                        `json:"data,required"`
	JSON customerInvoiceGetResponseJSON `json:"-"`
}

// customerInvoiceGetResponseJSON contains the JSON metadata for the struct
// [CustomerInvoiceGetResponse]
type customerInvoiceGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceAddChargeResponse struct {
	JSON customerInvoiceAddChargeResponseJSON `json:"-"`
}

// customerInvoiceAddChargeResponseJSON contains the JSON metadata for the struct
// [CustomerInvoiceAddChargeResponse]
type customerInvoiceAddChargeResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceAddChargeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerInvoiceAddChargeResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerInvoiceGetParams struct {
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
}

// URLQuery serializes [CustomerInvoiceGetParams]'s query parameters as
// `url.Values`.
func (r CustomerInvoiceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerInvoiceListParams struct {
	// Only return invoices for the specified credit type
	CreditTypeID param.Field[string] `query:"credit_type_id"`
	// RFC 3339 timestamp (exclusive). Invoices will only be returned for billing
	// periods that end before this time.
	EndingBefore param.Field[time.Time] `query:"ending_before" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems param.Field[bool] `query:"skip_zero_qty_line_items"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
	// date_asc.
	Sort param.Field[CustomerInvoiceListParamsSort] `query:"sort"`
	// RFC 3339 timestamp (inclusive). Invoices will only be returned for billing
	// periods that start at or after this time.
	StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
	// Invoice status, e.g. DRAFT, FINALIZED, or VOID
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [CustomerInvoiceListParams]'s query parameters as
// `url.Values`.
func (r CustomerInvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice sort order by issued_at, e.g. date_asc or date_desc. Defaults to
// date_asc.
type CustomerInvoiceListParamsSort string

const (
	CustomerInvoiceListParamsSortDateAsc  CustomerInvoiceListParamsSort = "date_asc"
	CustomerInvoiceListParamsSortDateDesc CustomerInvoiceListParamsSort = "date_desc"
)

func (r CustomerInvoiceListParamsSort) IsKnown() bool {
	switch r {
	case CustomerInvoiceListParamsSortDateAsc, CustomerInvoiceListParamsSortDateDesc:
		return true
	}
	return false
}

type CustomerInvoiceAddChargeParams struct {
	// The Metronome ID of the charge to add to the invoice. Note that the charge must
	// be on a product that is not on the current plan, and the product must have only
	// fixed charges.
	ChargeID param.Field[string] `json:"charge_id,required" format:"uuid"`
	// The Metronome ID of the customer plan to add the charge to.
	CustomerPlanID param.Field[string] `json:"customer_plan_id,required" format:"uuid"`
	Description    param.Field[string] `json:"description,required"`
	// The start_timestamp of the invoice to add the charge to.
	InvoiceStartTimestamp param.Field[time.Time] `json:"invoice_start_timestamp,required" format:"date-time"`
	// The price of the charge. This price will match the currency on the invoice, e.g.
	// USD cents.
	Price    param.Field[float64] `json:"price,required"`
	Quantity param.Field[float64] `json:"quantity,required"`
}

func (r CustomerInvoiceAddChargeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
