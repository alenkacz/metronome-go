// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/apiquery"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
)

// CustomerInvoiceService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerInvoiceService] method
// instead.
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
func (r *CustomerInvoiceService) Get(ctx context.Context, customerID string, invoiceID string, opts ...option.RequestOption) (res *CustomerInvoiceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/invoices/%s", customerID, invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all invoices for a given customer, optionally filtered by status, date
// range, and/or credit type.
func (r *CustomerInvoiceService) List(ctx context.Context, customerID string, query CustomerInvoiceListParams, opts ...option.RequestOption) (res *CustomerInvoiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/invoices", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Invoice struct {
	ID               string                  `json:"id,required" format:"uuid"`
	CreditType       InvoiceCreditType       `json:"credit_type,required"`
	CustomerID       string                  `json:"customer_id,required" format:"uuid"`
	LineItems        []InvoiceLineItem       `json:"line_items,required"`
	Status           string                  `json:"status,required"`
	Total            float64                 `json:"total,required"`
	AmendmentID      string                  `json:"amendment_id" format:"uuid"`
	ContractID       string                  `json:"contract_id" format:"uuid"`
	CorrectionRecord InvoiceCorrectionRecord `json:"correction_record"`
	CustomFields     interface{}             `json:"custom_fields"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp       time.Time                  `json:"end_timestamp" format:"date-time"`
	ExternalInvoice    InvoiceExternalInvoice     `json:"external_invoice,nullable"`
	InvoiceAdjustments []InvoiceInvoiceAdjustment `json:"invoice_adjustments"`
	// When the invoice was issued (UTC)
	IssuedAt             time.Time `json:"issued_at" format:"date-time"`
	NetPaymentTermsDays  float64   `json:"net_payment_terms_days"`
	NetsuiteSalesOrderID string    `json:"netsuite_sales_order_id"`
	PlanID               string    `json:"plan_id" format:"uuid"`
	PlanName             string    `json:"plan_name"`
	// only present for beta contract invoices with reseller royalties
	ResellerRoyalty         InvoiceResellerRoyalty `json:"reseller_royalty"`
	SalesforceOpportunityID string                 `json:"salesforce_opportunity_id"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp time.Time `json:"start_timestamp" format:"date-time"`
	Subtotal       float64   `json:"subtotal"`
	JSON           invoiceJSON
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID                      apijson.Field
	CreditType              apijson.Field
	CustomerID              apijson.Field
	LineItems               apijson.Field
	Status                  apijson.Field
	Total                   apijson.Field
	AmendmentID             apijson.Field
	ContractID              apijson.Field
	CorrectionRecord        apijson.Field
	CustomFields            apijson.Field
	EndTimestamp            apijson.Field
	ExternalInvoice         apijson.Field
	InvoiceAdjustments      apijson.Field
	IssuedAt                apijson.Field
	NetPaymentTermsDays     apijson.Field
	NetsuiteSalesOrderID    apijson.Field
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

type InvoiceCreditType struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON invoiceCreditTypeJSON
}

// invoiceCreditTypeJSON contains the JSON metadata for the struct
// [InvoiceCreditType]
type invoiceCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItem struct {
	CreditType InvoiceLineItemsCreditType `json:"credit_type,required"`
	Name       string                     `json:"name,required"`
	Total      float64                    `json:"total,required"`
	// only present for beta contract invoices
	CommitID string `json:"commit_id" format:"uuid"`
	// only present for beta contract invoices
	CommitSegmentID string            `json:"commit_segment_id" format:"uuid"`
	CustomFields    map[string]string `json:"custom_fields"`
	// only present for beta contract invoices
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	GroupKey     string    `json:"group_key"`
	GroupValue   string    `json:"group_value"`
	// only present for beta contract invoices
	NetsuiteItemID string `json:"netsuite_item_id"`
	// only present for beta contract invoices
	PostpaidCommit InvoiceLineItemsPostpaidCommit `json:"postpaid_commit"`
	ProductID      string                         `json:"product_id" format:"uuid"`
	Quantity       float64                        `json:"quantity"`
	ResellerType   InvoiceLineItemsResellerType   `json:"reseller_type"`
	// only present for beta contract invoices
	StartingAt   time.Time                     `json:"starting_at" format:"date-time"`
	SubLineItems []InvoiceLineItemsSubLineItem `json:"sub_line_items"`
	// only present for beta contract invoices
	UnitPrice float64 `json:"unit_price"`
	JSON      invoiceLineItemJSON
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	CreditType      apijson.Field
	Name            apijson.Field
	Total           apijson.Field
	CommitID        apijson.Field
	CommitSegmentID apijson.Field
	CustomFields    apijson.Field
	EndingBefore    apijson.Field
	GroupKey        apijson.Field
	GroupValue      apijson.Field
	NetsuiteItemID  apijson.Field
	PostpaidCommit  apijson.Field
	ProductID       apijson.Field
	Quantity        apijson.Field
	ResellerType    apijson.Field
	StartingAt      apijson.Field
	SubLineItems    apijson.Field
	UnitPrice       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsCreditType struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON invoiceLineItemsCreditTypeJSON
}

// invoiceLineItemsCreditTypeJSON contains the JSON metadata for the struct
// [InvoiceLineItemsCreditType]
type invoiceLineItemsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// only present for beta contract invoices
type InvoiceLineItemsPostpaidCommit struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON invoiceLineItemsPostpaidCommitJSON
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

type InvoiceLineItemsResellerType string

const (
	InvoiceLineItemsResellerTypeAws InvoiceLineItemsResellerType = "AWS"
	InvoiceLineItemsResellerTypeGcp InvoiceLineItemsResellerType = "GCP"
)

type InvoiceLineItemsSubLineItem struct {
	CustomFields  map[string]string `json:"custom_fields,required"`
	Name          string            `json:"name,required"`
	Quantity      float64           `json:"quantity,required"`
	Subtotal      float64           `json:"subtotal,required"`
	ChargeID      string            `json:"charge_id" format:"uuid"`
	CreditGrantID string            `json:"credit_grant_id" format:"uuid"`
	// the unit price for this charge, present only if the charge is not tiered and the
	// quantity is nonzero
	Price float64                            `json:"price"`
	Tiers []InvoiceLineItemsSubLineItemsTier `json:"tiers"`
	JSON  invoiceLineItemsSubLineItemJSON
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
	Price         apijson.Field
	Tiers         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsSubLineItemsTier struct {
	Price    float64 `json:"price,required"`
	Quantity float64 `json:"quantity,required"`
	// at what metric amount this tier begins
	StartingAt float64 `json:"starting_at,required"`
	Subtotal   float64 `json:"subtotal,required"`
	JSON       invoiceLineItemsSubLineItemsTierJSON
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

type InvoiceCorrectionRecord struct {
	Memo   string `json:"memo,required"`
	Reason string `json:"reason,required"`
	JSON   invoiceCorrectionRecordJSON
}

// invoiceCorrectionRecordJSON contains the JSON metadata for the struct
// [InvoiceCorrectionRecord]
type invoiceCorrectionRecordJSON struct {
	Memo        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCorrectionRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceExternalInvoice struct {
	BillingProviderType InvoiceExternalInvoiceBillingProviderType `json:"billing_provider_type,required"`
	ExternalStatus      InvoiceExternalInvoiceExternalStatus      `json:"external_status"`
	InvoiceID           string                                    `json:"invoice_id"`
	IssuedAtTimestamp   time.Time                                 `json:"issued_at_timestamp" format:"date-time"`
	JSON                invoiceExternalInvoiceJSON
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

type InvoiceExternalInvoiceBillingProviderType string

const (
	InvoiceExternalInvoiceBillingProviderTypeAwsMarketplace   InvoiceExternalInvoiceBillingProviderType = "aws_marketplace"
	InvoiceExternalInvoiceBillingProviderTypeStripe           InvoiceExternalInvoiceBillingProviderType = "stripe"
	InvoiceExternalInvoiceBillingProviderTypeNetsuite         InvoiceExternalInvoiceBillingProviderType = "netsuite"
	InvoiceExternalInvoiceBillingProviderTypeCustom           InvoiceExternalInvoiceBillingProviderType = "custom"
	InvoiceExternalInvoiceBillingProviderTypeAzureMarketplace InvoiceExternalInvoiceBillingProviderType = "azure_marketplace"
	InvoiceExternalInvoiceBillingProviderTypeQuickbooksOnline InvoiceExternalInvoiceBillingProviderType = "quickbooks_online"
)

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

type InvoiceInvoiceAdjustment struct {
	CreditType InvoiceInvoiceAdjustmentsCreditType `json:"credit_type,required"`
	Name       string                              `json:"name,required"`
	Total      float64                             `json:"total,required"`
	JSON       invoiceInvoiceAdjustmentJSON
}

// invoiceInvoiceAdjustmentJSON contains the JSON metadata for the struct
// [InvoiceInvoiceAdjustment]
type invoiceInvoiceAdjustmentJSON struct {
	CreditType  apijson.Field
	Name        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceInvoiceAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceInvoiceAdjustmentsCreditType struct {
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	JSON invoiceInvoiceAdjustmentsCreditTypeJSON
}

// invoiceInvoiceAdjustmentsCreditTypeJSON contains the JSON metadata for the
// struct [InvoiceInvoiceAdjustmentsCreditType]
type invoiceInvoiceAdjustmentsCreditTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceInvoiceAdjustmentsCreditType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// only present for beta contract invoices with reseller royalties
type InvoiceResellerRoyalty struct {
	Fraction           string                             `json:"fraction,required"`
	NetsuiteResellerID string                             `json:"netsuite_reseller_id,required"`
	ResellerType       InvoiceResellerRoyaltyResellerType `json:"reseller_type,required"`
	AwsOptions         InvoiceResellerRoyaltyAwsOptions   `json:"aws_options"`
	GcpOptions         InvoiceResellerRoyaltyGcpOptions   `json:"gcp_options"`
	JSON               invoiceResellerRoyaltyJSON
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

type InvoiceResellerRoyaltyResellerType string

const (
	InvoiceResellerRoyaltyResellerTypeAws InvoiceResellerRoyaltyResellerType = "AWS"
	InvoiceResellerRoyaltyResellerTypeGcp InvoiceResellerRoyaltyResellerType = "GCP"
)

type InvoiceResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string `json:"aws_account_number"`
	AwsOfferID          string `json:"aws_offer_id"`
	AwsPayerReferenceID string `json:"aws_payer_reference_id"`
	JSON                invoiceResellerRoyaltyAwsOptionsJSON
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

type InvoiceResellerRoyaltyGcpOptions struct {
	GcpAccountID string `json:"gcp_account_id"`
	GcpOfferID   string `json:"gcp_offer_id"`
	JSON         invoiceResellerRoyaltyGcpOptionsJSON
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

type CustomerInvoiceGetResponse struct {
	Data Invoice `json:"data,required"`
	JSON customerInvoiceGetResponseJSON
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

type CustomerInvoiceListResponse struct {
	Data     []Invoice `json:"data,required"`
	NextPage string    `json:"next_page,required,nullable"`
	JSON     customerInvoiceListResponseJSON
}

// customerInvoiceListResponseJSON contains the JSON metadata for the struct
// [CustomerInvoiceListResponse]
type customerInvoiceListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerInvoiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
