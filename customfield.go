// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"

	"github.com/metronome/metronome-go/internal/apijson"
	"github.com/metronome/metronome-go/internal/param"
	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
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

// Add a key to the allow list for a given entity.
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
func (r *CustomFieldService) ListKeys(ctx context.Context, body CustomFieldListKeysParams, opts ...option.RequestOption) (res *CustomFieldListKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customFields/listKeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
// updates are not supported.
func (r *CustomFieldService) SetValues(ctx context.Context, body CustomFieldSetValuesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "customFields/setValues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomFieldListKeysResponse struct {
	Data     []CustomFieldListKeysResponseData `json:"data,required"`
	NextPage string                            `json:"next_page,required,nullable"`
	JSON     customFieldListKeysResponseJSON
}

// customFieldListKeysResponseJSON contains the JSON metadata for the struct
// [CustomFieldListKeysResponse]
type customFieldListKeysResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomFieldListKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomFieldListKeysResponseData struct {
	EnforceUniqueness bool                                  `json:"enforce_uniqueness,required"`
	Entity            CustomFieldListKeysResponseDataEntity `json:"entity,required"`
	Key               string                                `json:"key,required"`
	JSON              customFieldListKeysResponseDataJSON
}

// customFieldListKeysResponseDataJSON contains the JSON metadata for the struct
// [CustomFieldListKeysResponseData]
type customFieldListKeysResponseDataJSON struct {
	EnforceUniqueness apijson.Field
	Entity            apijson.Field
	Key               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomFieldListKeysResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomFieldListKeysResponseDataEntity string

const (
	CustomFieldListKeysResponseDataEntityCharge         CustomFieldListKeysResponseDataEntity = "charge"
	CustomFieldListKeysResponseDataEntityCreditGrant    CustomFieldListKeysResponseDataEntity = "credit_grant"
	CustomFieldListKeysResponseDataEntityCustomer       CustomFieldListKeysResponseDataEntity = "customer"
	CustomFieldListKeysResponseDataEntityCustomerPlan   CustomFieldListKeysResponseDataEntity = "customer_plan"
	CustomFieldListKeysResponseDataEntityPlan           CustomFieldListKeysResponseDataEntity = "plan"
	CustomFieldListKeysResponseDataEntityProduct        CustomFieldListKeysResponseDataEntity = "product"
	CustomFieldListKeysResponseDataEntityBillableMetric CustomFieldListKeysResponseDataEntity = "billable_metric"
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
)

type CustomFieldListKeysParams struct {
	// Optional list of entity types to return keys for
	Entities param.Field[[]CustomFieldListKeysParamsEntity] `json:"entities"`
}

func (r CustomFieldListKeysParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
)
