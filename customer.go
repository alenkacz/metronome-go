// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "github.com/metronome/metronome-go/internal/shared"
  "time"
  "github.com/metronome/metronome-go/internal/param"
  "net/url"
  "github.com/metronome/metronome-go/internal/apiquery"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
  "fmt"
)

// CustomerService contains methods and other services that help with interacting
// with the metronome API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomerService] method instead.
type CustomerService struct {
Options []option.RequestOption
Plans *CustomerPlanService
Invoices *CustomerInvoiceService
BillingConfig *CustomerBillingConfigService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
  r = &CustomerService{}
  r.Options = opts
  r.Plans = NewCustomerPlanService(opts...)
  r.Invoices = NewCustomerInvoiceService(opts...)
  r.BillingConfig = NewCustomerBillingConfigService(opts...)
  return
}

// Create a new customer
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *CustomerNewResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "customers"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Get a customer by Metronome ID.
func (r *CustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerGetResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("customers/%s", customerID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
  return
}

// List all customers.
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *CustomerListResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "customers"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

// Archive a customer
func (r *CustomerService) Archive(ctx context.Context, body CustomerArchiveParams, opts ...option.RequestOption) (res *CustomerArchiveResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "customers/archive"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// List all billable metrics.
func (r *CustomerService) BillableMetrics(ctx context.Context, customerID string, query CustomerBillableMetricsParams, opts ...option.RequestOption) (res *CustomerBillableMetricsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("customers/%s/billable-metrics", customerID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

// Fetch daily pending costs for the specified customer, broken down by credit type
// and line items. Note: this is not supported for customers whose plan includes a
// UNIQUE-type billable metric.
func (r *CustomerService) Costs(ctx context.Context, customerID string, query CustomerCostsParams, opts ...option.RequestOption) (res *CustomerCostsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("customers/%s/costs", customerID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

// Sets the ingest aliases for a customer. Ingest aliases can be used in the
// `customer_id` field when sending usage events to Metronome. This call is
// idempotent. It fully replaces the set of ingest aliases for the given customer.
func (r *CustomerService) SetIngestAliases(ctx context.Context, customerID string, body CustomerSetIngestAliasesParams, opts ...option.RequestOption) (err error) {
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
  path := fmt.Sprintf("customers/%s/setIngestAliases", customerID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
  return
}

// Updates the specified customer's name.
func (r *CustomerService) SetName(ctx context.Context, customerID string, body CustomerSetNameParams, opts ...option.RequestOption) (res *CustomerSetNameResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := fmt.Sprintf("customers/%s/setName", customerID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
  return
}

// Updates the specified customer's config.
func (r *CustomerService) UpdateConfig(ctx context.Context, customerID string, body CustomerUpdateConfigParams, opts ...option.RequestOption) (err error) {
  opts = append(r.Options[:], opts...)
  opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
  path := fmt.Sprintf("customers/%s/updateConfig", customerID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
  return
}

type Customer struct {
// the Metronome ID of the customer
ID string `json:"id,required" format:"uuid"`
// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
// alias) that can be used in usage events
ExternalID string `json:"external_id,required"`
// aliases for this customer that can be used instead of the Metronome customer ID
// in usage events
IngestAliases []string `json:"ingest_aliases,required"`
Name string `json:"name,required"`
JSON customerJSON
}

// customerJSON contains the JSON metadata for the struct [Customer]
type customerJSON struct {
ID apijson.Field
ExternalID apijson.Field
IngestAliases apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Customer) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerDetail struct {
// the Metronome ID of the customer
ID string `json:"id,required" format:"uuid"`
CustomFields map[string]string `json:"custom_fields,required"`
CustomerConfig CustomerDetailCustomerConfig `json:"customer_config,required"`
// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest
// alias) that can be used in usage events
ExternalID string `json:"external_id,required"`
// aliases for this customer that can be used instead of the Metronome customer ID
// in usage events
IngestAliases []string `json:"ingest_aliases,required"`
Name string `json:"name,required"`
JSON customerDetailJSON
}

// customerDetailJSON contains the JSON metadata for the struct [CustomerDetail]
type customerDetailJSON struct {
ID apijson.Field
CustomFields apijson.Field
CustomerConfig apijson.Field
ExternalID apijson.Field
IngestAliases apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerDetail) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerDetailCustomerConfig struct {
// The Salesforce account ID for the customer
SalesforceAccountID string `json:"salesforce_account_id,required,nullable"`
JSON customerDetailCustomerConfigJSON
}

// customerDetailCustomerConfigJSON contains the JSON metadata for the struct
// [CustomerDetailCustomerConfig]
type customerDetailCustomerConfigJSON struct {
SalesforceAccountID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerDetailCustomerConfig) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerNewResponse struct {
Data Customer `json:"data,required"`
JSON customerNewResponseJSON
}

// customerNewResponseJSON contains the JSON metadata for the struct
// [CustomerNewResponse]
type customerNewResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerNewResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerGetResponse struct {
Data CustomerDetail `json:"data,required"`
JSON customerGetResponseJSON
}

// customerGetResponseJSON contains the JSON metadata for the struct
// [CustomerGetResponse]
type customerGetResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerGetResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponse struct {
Data []CustomerDetail `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON customerListResponseJSON
}

// customerListResponseJSON contains the JSON metadata for the struct
// [CustomerListResponse]
type customerListResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerListResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerArchiveResponse struct {
Data shared.ID `json:"data,required"`
JSON customerArchiveResponseJSON
}

// customerArchiveResponseJSON contains the JSON metadata for the struct
// [CustomerArchiveResponse]
type customerArchiveResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerArchiveResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerBillableMetricsResponse struct {
Data []CustomerBillableMetricsResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON customerBillableMetricsResponseJSON
}

// customerBillableMetricsResponseJSON contains the JSON metadata for the struct
// [CustomerBillableMetricsResponse]
type customerBillableMetricsResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerBillableMetricsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerBillableMetricsResponseData struct {
ID string `json:"id,required" format:"uuid"`
Name string `json:"name,required"`
GroupBy []string `json:"group_by"`
JSON customerBillableMetricsResponseDataJSON
}

// customerBillableMetricsResponseDataJSON contains the JSON metadata for the
// struct [CustomerBillableMetricsResponseData]
type customerBillableMetricsResponseDataJSON struct {
ID apijson.Field
Name apijson.Field
GroupBy apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerBillableMetricsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerCostsResponse struct {
Data []CustomerCostsResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON customerCostsResponseJSON
}

// customerCostsResponseJSON contains the JSON metadata for the struct
// [CustomerCostsResponse]
type customerCostsResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerCostsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerCostsResponseData struct {
CreditTypes map[string]CustomerCostsResponseDataCreditType `json:"credit_types,required"`
EndTimestamp time.Time `json:"end_timestamp,required" format:"date-time"`
StartTimestamp time.Time `json:"start_timestamp,required" format:"date-time"`
JSON customerCostsResponseDataJSON
}

// customerCostsResponseDataJSON contains the JSON metadata for the struct
// [CustomerCostsResponseData]
type customerCostsResponseDataJSON struct {
CreditTypes apijson.Field
EndTimestamp apijson.Field
StartTimestamp apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerCostsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerCostsResponseDataCreditType struct {
Cost float64 `json:"cost"`
LineItemBreakdown []CustomerCostsResponseDataCreditTypesLineItemBreakdown `json:"line_item_breakdown"`
Name string `json:"name"`
JSON customerCostsResponseDataCreditTypeJSON
}

// customerCostsResponseDataCreditTypeJSON contains the JSON metadata for the
// struct [CustomerCostsResponseDataCreditType]
type customerCostsResponseDataCreditTypeJSON struct {
Cost apijson.Field
LineItemBreakdown apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerCostsResponseDataCreditType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerCostsResponseDataCreditTypesLineItemBreakdown struct {
Cost float64 `json:"cost,required"`
Name string `json:"name,required"`
GroupKey string `json:"group_key"`
GroupValue string `json:"group_value,nullable"`
JSON customerCostsResponseDataCreditTypesLineItemBreakdownJSON
}

// customerCostsResponseDataCreditTypesLineItemBreakdownJSON contains the JSON
// metadata for the struct [CustomerCostsResponseDataCreditTypesLineItemBreakdown]
type customerCostsResponseDataCreditTypesLineItemBreakdownJSON struct {
Cost apijson.Field
Name apijson.Field
GroupKey apijson.Field
GroupValue apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerCostsResponseDataCreditTypesLineItemBreakdown) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerSetNameResponse struct {
Data Customer `json:"data,required"`
JSON customerSetNameResponseJSON
}

// customerSetNameResponseJSON contains the JSON metadata for the struct
// [CustomerSetNameResponse]
type customerSetNameResponseJSON struct {
Data apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomerSetNameResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParams struct {
Name param.Field[string] `json:"name,required"`
BillingConfig param.Field[CustomerNewParamsBillingConfig] `json:"billing_config"`
// (deprecated, use ingest_aliases instead) the first ID (Metronome ID or ingest
// alias) that can be used in usage events
ExternalID param.Field[string] `json:"external_id"`
// Aliases that can be used to refer to this customer in usage events
IngestAliases param.Field[[]string] `json:"ingest_aliases"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingConfig struct {
BillingProviderCustomerID param.Field[string] `json:"billing_provider_customer_id,required"`
BillingProviderType param.Field[CustomerNewParamsBillingConfigBillingProviderType] `json:"billing_provider_type,required"`
AwsProductCode param.Field[string] `json:"aws_product_code"`
AwsRegion param.Field[CustomerNewParamsBillingConfigAwsRegion] `json:"aws_region"`
StripeCollectionMethod param.Field[CustomerNewParamsBillingConfigStripeCollectionMethod] `json:"stripe_collection_method"`
}

func (r CustomerNewParamsBillingConfig) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CustomerNewParamsBillingConfigBillingProviderType string

const (
  CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace CustomerNewParamsBillingConfigBillingProviderType = "aws_marketplace"
  CustomerNewParamsBillingConfigBillingProviderTypeStripe CustomerNewParamsBillingConfigBillingProviderType = "stripe"
  CustomerNewParamsBillingConfigBillingProviderTypeNetsuite CustomerNewParamsBillingConfigBillingProviderType = "netsuite"
  CustomerNewParamsBillingConfigBillingProviderTypeCustom CustomerNewParamsBillingConfigBillingProviderType = "custom"
  CustomerNewParamsBillingConfigBillingProviderTypeAzureMarketplace CustomerNewParamsBillingConfigBillingProviderType = "azure_marketplace"
  CustomerNewParamsBillingConfigBillingProviderTypeQuickbooksOnline CustomerNewParamsBillingConfigBillingProviderType = "quickbooks_online"
)

type CustomerNewParamsBillingConfigAwsRegion string

const (
  CustomerNewParamsBillingConfigAwsRegionAfSouth1 CustomerNewParamsBillingConfigAwsRegion = "af-south-1"
  CustomerNewParamsBillingConfigAwsRegionApEast1 CustomerNewParamsBillingConfigAwsRegion = "ap-east-1"
  CustomerNewParamsBillingConfigAwsRegionApNortheast1 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-1"
  CustomerNewParamsBillingConfigAwsRegionApNortheast2 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-2"
  CustomerNewParamsBillingConfigAwsRegionApNortheast3 CustomerNewParamsBillingConfigAwsRegion = "ap-northeast-3"
  CustomerNewParamsBillingConfigAwsRegionApSouth1 CustomerNewParamsBillingConfigAwsRegion = "ap-south-1"
  CustomerNewParamsBillingConfigAwsRegionApSoutheast1 CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-1"
  CustomerNewParamsBillingConfigAwsRegionApSoutheast2 CustomerNewParamsBillingConfigAwsRegion = "ap-southeast-2"
  CustomerNewParamsBillingConfigAwsRegionCaCentral1 CustomerNewParamsBillingConfigAwsRegion = "ca-central-1"
  CustomerNewParamsBillingConfigAwsRegionCnNorth1 CustomerNewParamsBillingConfigAwsRegion = "cn-north-1"
  CustomerNewParamsBillingConfigAwsRegionCnNorthwest1 CustomerNewParamsBillingConfigAwsRegion = "cn-northwest-1"
  CustomerNewParamsBillingConfigAwsRegionEuCentral1 CustomerNewParamsBillingConfigAwsRegion = "eu-central-1"
  CustomerNewParamsBillingConfigAwsRegionEuNorth1 CustomerNewParamsBillingConfigAwsRegion = "eu-north-1"
  CustomerNewParamsBillingConfigAwsRegionEuSouth1 CustomerNewParamsBillingConfigAwsRegion = "eu-south-1"
  CustomerNewParamsBillingConfigAwsRegionEuWest1 CustomerNewParamsBillingConfigAwsRegion = "eu-west-1"
  CustomerNewParamsBillingConfigAwsRegionEuWest2 CustomerNewParamsBillingConfigAwsRegion = "eu-west-2"
  CustomerNewParamsBillingConfigAwsRegionEuWest3 CustomerNewParamsBillingConfigAwsRegion = "eu-west-3"
  CustomerNewParamsBillingConfigAwsRegionMeSouth1 CustomerNewParamsBillingConfigAwsRegion = "me-south-1"
  CustomerNewParamsBillingConfigAwsRegionSaEast1 CustomerNewParamsBillingConfigAwsRegion = "sa-east-1"
  CustomerNewParamsBillingConfigAwsRegionUsEast1 CustomerNewParamsBillingConfigAwsRegion = "us-east-1"
  CustomerNewParamsBillingConfigAwsRegionUsEast2 CustomerNewParamsBillingConfigAwsRegion = "us-east-2"
  CustomerNewParamsBillingConfigAwsRegionUsGovEast1 CustomerNewParamsBillingConfigAwsRegion = "us-gov-east-1"
  CustomerNewParamsBillingConfigAwsRegionUsGovWest1 CustomerNewParamsBillingConfigAwsRegion = "us-gov-west-1"
  CustomerNewParamsBillingConfigAwsRegionUsWest1 CustomerNewParamsBillingConfigAwsRegion = "us-west-1"
  CustomerNewParamsBillingConfigAwsRegionUsWest2 CustomerNewParamsBillingConfigAwsRegion = "us-west-2"
)

type CustomerNewParamsBillingConfigStripeCollectionMethod string

const (
  CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically CustomerNewParamsBillingConfigStripeCollectionMethod = "charge_automatically"
  CustomerNewParamsBillingConfigStripeCollectionMethodSendInvoice CustomerNewParamsBillingConfigStripeCollectionMethod = "send_invoice"
)

type CustomerListParams struct {
// Filter the customer list by customer_id. Up to 100 ids can be provided.
CustomerIDs param.Field[[]string] `query:"customer_ids"`
// Filter the customer list by ingest_alias
IngestAlias param.Field[string] `query:"ingest_alias"`
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
// Filter the customer list by only archived customers.
OnlyArchived param.Field[bool] `query:"only_archived"`
// Filter the customer list by salesforce_account_id. Up to 100 ids can be
// provided.
SalesforceAccountIDs param.Field[[]string] `query:"salesforce_account_ids"`
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type CustomerArchiveParams struct {
ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r CustomerArchiveParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CustomerBillableMetricsParams struct {
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
// If true, the list of metrics will be filtered to just ones that are on the
// customer's current plan
OnCurrentPlan param.Field[bool] `query:"on_current_plan"`
}

// URLQuery serializes [CustomerBillableMetricsParams]'s query parameters as
// `url.Values`.
func (r CustomerBillableMetricsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type CustomerCostsParams struct {
// RFC 3339 timestamp (exclusive)
EndingBefore param.Field[time.Time] `query:"ending_before,required" format:"date-time"`
// RFC 3339 timestamp (inclusive)
StartingOn param.Field[time.Time] `query:"starting_on,required" format:"date-time"`
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
}

// URLQuery serializes [CustomerCostsParams]'s query parameters as `url.Values`.
func (r CustomerCostsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}

type CustomerSetIngestAliasesParams struct {
IngestAliases param.Field[[]string] `json:"ingest_aliases,required"`
}

func (r CustomerSetIngestAliasesParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CustomerSetNameParams struct {
// The new name for the customer
Name param.Field[string] `json:"name,required"`
}

func (r CustomerSetNameParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CustomerUpdateConfigParams struct {
// The Salesforce account ID for the customer
SalesforceAccountID param.Field[string] `json:"salesforce_account_id"`
}

func (r CustomerUpdateConfigParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
