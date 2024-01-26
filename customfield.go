// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/internal/shared"
	"github.com/Metronome-Industries/metronome-go/option"
)

// CustomFieldService contains methods and other services that help with
// interacting with the metronome API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomFieldService] method
// instead.
type CustomFieldService struct {
	Options []option.RequestOption
}

// NewCustomFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomFieldService(opts ...option.RequestOption) (r *CustomFieldService) {
	r = &CustomFieldService{}
	r.Options = opts
	return
}

// Add a key to the allow list for a given entity. There is a 100 character limit
// on custom field keys.
func (r *CustomFieldService) AddKey(ctx context.Context, body CustomFieldAddKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/addKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Deletes one or more custom fields on an instance of a Metronome entity.
func (r *CustomFieldService) DeleteValues(ctx context.Context, body CustomFieldDeleteValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/deleteValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List all active custom field keys, optionally filtered by entity type.
func (r *CustomFieldService) ListKeys(ctx context.Context, params CustomFieldListKeysParams, opts ...option.RequestOption) (res *shared.Page[CustomFieldListKeysResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "customFields/listKeys"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List all active custom field keys, optionally filtered by entity type.
func (r *CustomFieldService) ListKeysAutoPaging(ctx context.Context, params CustomFieldListKeysParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomFieldListKeysResponse] {
	return shared.NewPageAutoPager(r.ListKeys(ctx, params, opts...))
}

// Remove a key from the allow list for a given entity.
func (r *CustomFieldService) RemoveKey(ctx context.Context, body CustomFieldRemoveKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/removeKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sets one or more custom fields on an instance of a Metronome entity. If a
// key/value pair passed in this request matches one already set on the entity, its
// value will be overwritten. Any key/value pairs that exist on the entity that do
// not match those passed in this request will remain untouched. This endpoint is
// transactional and will update all key/value pairs or no key/value pairs. Partial
// updates are not supported. There is a 200 character limit on custom field
// values.
func (r *CustomFieldService) SetValues(ctx context.Context, body CustomFieldSetValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/setValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomFieldListKeysResponse struct {
	EnforceUniqueness bool                              `json:"enforce_uniqueness,required"`
	Entity            CustomFieldListKeysResponseEntity `json:"entity,required"`
	Key               string                            `json:"key,required"`
	JSON              customFieldListKeysResponseJSON   `json:"-"`
}

// customFieldListKeysResponseJSON contains the JSON metadata for the struct
// [CustomFieldListKeysResponse]
type customFieldListKeysResponseJSON struct {
	EnforceUniqueness apijson.Field
	Entity            apijson.Field
	Key               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomFieldListKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomFieldListKeysResponseEntity string

const (
	CustomFieldListKeysResponseEntityCharge         CustomFieldListKeysResponseEntity = "charge"
	CustomFieldListKeysResponseEntityCreditGrant    CustomFieldListKeysResponseEntity = "credit_grant"
	CustomFieldListKeysResponseEntityCustomer       CustomFieldListKeysResponseEntity = "customer"
	CustomFieldListKeysResponseEntityCustomerPlan   CustomFieldListKeysResponseEntity = "customer_plan"
	CustomFieldListKeysResponseEntityPlan           CustomFieldListKeysResponseEntity = "plan"
	CustomFieldListKeysResponseEntityProduct        CustomFieldListKeysResponseEntity = "product"
	CustomFieldListKeysResponseEntityBillableMetric CustomFieldListKeysResponseEntity = "billable_metric"
	CustomFieldListKeysResponseEntityCommit         CustomFieldListKeysResponseEntity = "commit"
)

type CustomFieldAddKeyParams struct {
	EnforceUniqueness param.Field[bool]                          `json:"enforce_uniqueness,required"`
	Entity            param.Field[CustomFieldAddKeyParamsEntity] `json:"entity,required"`
	Key               param.Field[string]                        `json:"key,required"`
}

func (r CustomFieldAddKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldAddKeyParamsEntity string

const (
	CustomFieldAddKeyParamsEntityCharge         CustomFieldAddKeyParamsEntity = "charge"
	CustomFieldAddKeyParamsEntityCreditGrant    CustomFieldAddKeyParamsEntity = "credit_grant"
	CustomFieldAddKeyParamsEntityCustomer       CustomFieldAddKeyParamsEntity = "customer"
	CustomFieldAddKeyParamsEntityCustomerPlan   CustomFieldAddKeyParamsEntity = "customer_plan"
	CustomFieldAddKeyParamsEntityPlan           CustomFieldAddKeyParamsEntity = "plan"
	CustomFieldAddKeyParamsEntityProduct        CustomFieldAddKeyParamsEntity = "product"
	CustomFieldAddKeyParamsEntityBillableMetric CustomFieldAddKeyParamsEntity = "billable_metric"
	CustomFieldAddKeyParamsEntityCommit         CustomFieldAddKeyParamsEntity = "commit"
)

type CustomFieldDeleteValuesParams struct {
	Entity   param.Field[CustomFieldDeleteValuesParamsEntity] `json:"entity,required"`
	EntityID param.Field[string]                              `json:"entity_id,required" format:"uuid"`
	Keys     param.Field[[]string]                            `json:"keys,required"`
}

func (r CustomFieldDeleteValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldDeleteValuesParamsEntity string

const (
	CustomFieldDeleteValuesParamsEntityCharge         CustomFieldDeleteValuesParamsEntity = "charge"
	CustomFieldDeleteValuesParamsEntityCreditGrant    CustomFieldDeleteValuesParamsEntity = "credit_grant"
	CustomFieldDeleteValuesParamsEntityCustomer       CustomFieldDeleteValuesParamsEntity = "customer"
	CustomFieldDeleteValuesParamsEntityCustomerPlan   CustomFieldDeleteValuesParamsEntity = "customer_plan"
	CustomFieldDeleteValuesParamsEntityPlan           CustomFieldDeleteValuesParamsEntity = "plan"
	CustomFieldDeleteValuesParamsEntityProduct        CustomFieldDeleteValuesParamsEntity = "product"
	CustomFieldDeleteValuesParamsEntityBillableMetric CustomFieldDeleteValuesParamsEntity = "billable_metric"
	CustomFieldDeleteValuesParamsEntityCommit         CustomFieldDeleteValuesParamsEntity = "commit"
)

type CustomFieldListKeysParams struct {
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// Optional list of entity types to return keys for
	Entities param.Field[[]CustomFieldListKeysParamsEntity] `json:"entities"`
}

func (r CustomFieldListKeysParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CustomFieldListKeysParams]'s query parameters as
// `url.Values`.
func (r CustomFieldListKeysParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomFieldListKeysParamsEntity string

const (
	CustomFieldListKeysParamsEntityCharge         CustomFieldListKeysParamsEntity = "charge"
	CustomFieldListKeysParamsEntityCreditGrant    CustomFieldListKeysParamsEntity = "credit_grant"
	CustomFieldListKeysParamsEntityCustomer       CustomFieldListKeysParamsEntity = "customer"
	CustomFieldListKeysParamsEntityCustomerPlan   CustomFieldListKeysParamsEntity = "customer_plan"
	CustomFieldListKeysParamsEntityPlan           CustomFieldListKeysParamsEntity = "plan"
	CustomFieldListKeysParamsEntityProduct        CustomFieldListKeysParamsEntity = "product"
	CustomFieldListKeysParamsEntityBillableMetric CustomFieldListKeysParamsEntity = "billable_metric"
	CustomFieldListKeysParamsEntityCommit         CustomFieldListKeysParamsEntity = "commit"
)

type CustomFieldRemoveKeyParams struct {
	Entity param.Field[CustomFieldRemoveKeyParamsEntity] `json:"entity,required"`
	Key    param.Field[string]                           `json:"key,required"`
}

func (r CustomFieldRemoveKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldRemoveKeyParamsEntity string

const (
	CustomFieldRemoveKeyParamsEntityCharge         CustomFieldRemoveKeyParamsEntity = "charge"
	CustomFieldRemoveKeyParamsEntityCreditGrant    CustomFieldRemoveKeyParamsEntity = "credit_grant"
	CustomFieldRemoveKeyParamsEntityCustomer       CustomFieldRemoveKeyParamsEntity = "customer"
	CustomFieldRemoveKeyParamsEntityCustomerPlan   CustomFieldRemoveKeyParamsEntity = "customer_plan"
	CustomFieldRemoveKeyParamsEntityPlan           CustomFieldRemoveKeyParamsEntity = "plan"
	CustomFieldRemoveKeyParamsEntityProduct        CustomFieldRemoveKeyParamsEntity = "product"
	CustomFieldRemoveKeyParamsEntityBillableMetric CustomFieldRemoveKeyParamsEntity = "billable_metric"
	CustomFieldRemoveKeyParamsEntityCommit         CustomFieldRemoveKeyParamsEntity = "commit"
)

type CustomFieldSetValuesParams struct {
	CustomFields param.Field[map[string]string]                `json:"custom_fields,required"`
	Entity       param.Field[CustomFieldSetValuesParamsEntity] `json:"entity,required"`
	EntityID     param.Field[string]                           `json:"entity_id,required" format:"uuid"`
}

func (r CustomFieldSetValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomFieldSetValuesParamsEntity string

const (
	CustomFieldSetValuesParamsEntityCharge         CustomFieldSetValuesParamsEntity = "charge"
	CustomFieldSetValuesParamsEntityCreditGrant    CustomFieldSetValuesParamsEntity = "credit_grant"
	CustomFieldSetValuesParamsEntityCustomer       CustomFieldSetValuesParamsEntity = "customer"
	CustomFieldSetValuesParamsEntityCustomerPlan   CustomFieldSetValuesParamsEntity = "customer_plan"
	CustomFieldSetValuesParamsEntityPlan           CustomFieldSetValuesParamsEntity = "plan"
	CustomFieldSetValuesParamsEntityProduct        CustomFieldSetValuesParamsEntity = "product"
	CustomFieldSetValuesParamsEntityBillableMetric CustomFieldSetValuesParamsEntity = "billable_metric"
	CustomFieldSetValuesParamsEntityCommit         CustomFieldSetValuesParamsEntity = "commit"
)
