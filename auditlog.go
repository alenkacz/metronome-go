// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/apijson"
  "time"
  "github.com/metronome/metronome-go/internal/param"
  "net/url"
  "github.com/metronome/metronome-go/internal/apiquery"
  "context"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/requestconfig"
)

// AuditLogService contains methods and other services that help with interacting
// with the metronome API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAuditLogService] method instead.
type AuditLogService struct {
Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r *AuditLogService) {
  r = &AuditLogService{}
  r.Options = opts
  return
}

// Retrieves a range of audit logs. If no further audit logs are currently
// available, the data array will be empty. As new audit logs are created,
// subsequent requests using the same next_page value will be in the returned data
// array, ensuring a continuous and uninterrupted reading of audit logs.
func (r *AuditLogService) GetAuditLogs(ctx context.Context, query AuditLogGetAuditLogsParams, opts ...option.RequestOption) (res *AuditLogGetAuditLogsResponse, err error) {
  opts = append(r.Options[:], opts...)
  path := "auditLogs"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
  return
}

type AuditLogGetAuditLogsResponse struct {
Data []AuditLogGetAuditLogsResponseData `json:"data,required"`
NextPage string `json:"next_page,required,nullable"`
JSON auditLogGetAuditLogsResponseJSON
}

// auditLogGetAuditLogsResponseJSON contains the JSON metadata for the struct
// [AuditLogGetAuditLogsResponse]
type auditLogGetAuditLogsResponseJSON struct {
Data apijson.Field
NextPage apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *AuditLogGetAuditLogsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type AuditLogGetAuditLogsResponseData struct {
ID string `json:"id,required"`
Timestamp time.Time `json:"timestamp,required" format:"date-time"`
Action string `json:"action"`
Actor AuditLogGetAuditLogsResponseDataActor `json:"actor"`
ResourceID string `json:"resource_id"`
ResourceType string `json:"resource_type"`
Status AuditLogGetAuditLogsResponseDataStatus `json:"status"`
JSON auditLogGetAuditLogsResponseDataJSON
}

// auditLogGetAuditLogsResponseDataJSON contains the JSON metadata for the struct
// [AuditLogGetAuditLogsResponseData]
type auditLogGetAuditLogsResponseDataJSON struct {
ID apijson.Field
Timestamp apijson.Field
Action apijson.Field
Actor apijson.Field
ResourceID apijson.Field
ResourceType apijson.Field
Status apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *AuditLogGetAuditLogsResponseData) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type AuditLogGetAuditLogsResponseDataActor struct {
ID string `json:"id,required"`
Name string `json:"name,required"`
Email string `json:"email"`
JSON auditLogGetAuditLogsResponseDataActorJSON
}

// auditLogGetAuditLogsResponseDataActorJSON contains the JSON metadata for the
// struct [AuditLogGetAuditLogsResponseDataActor]
type auditLogGetAuditLogsResponseDataActorJSON struct {
ID apijson.Field
Name apijson.Field
Email apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *AuditLogGetAuditLogsResponseDataActor) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

type AuditLogGetAuditLogsResponseDataStatus string

const (
  AuditLogGetAuditLogsResponseDataStatusSuccess AuditLogGetAuditLogsResponseDataStatus = "success"
  AuditLogGetAuditLogsResponseDataStatusFailure AuditLogGetAuditLogsResponseDataStatus = "failure"
  AuditLogGetAuditLogsResponseDataStatusPending AuditLogGetAuditLogsResponseDataStatus = "pending"
)

type AuditLogGetAuditLogsParams struct {
// Max number of results that should be returned
Limit param.Field[int64] `query:"limit"`
// Cursor that indicates where the next page of results should start.
NextPage param.Field[string] `query:"next_page"`
// RFC 3339 timestamp of the earliest audit log to return. Cannot be used with
// 'next_page'.
StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
}

// URLQuery serializes [AuditLogGetAuditLogsParams]'s query parameters as
// `url.Values`.
func (r AuditLogGetAuditLogsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatComma,
    NestedFormat: apiquery.NestedQueryFormatBrackets,
  })
}
