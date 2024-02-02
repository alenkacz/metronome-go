// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/internal/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
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
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *AuditLogListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auditLogs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuditLogListResponse struct {
	Data     []AuditLogListResponseData `json:"data,required"`
	NextPage string                     `json:"next_page,required,nullable"`
	JSON     auditLogListResponseJSON   `json:"-"`
}

// auditLogListResponseJSON contains the JSON metadata for the struct
// [AuditLogListResponse]
type auditLogListResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogListResponseData struct {
	ID           string                         `json:"id,required"`
	Timestamp    time.Time                      `json:"timestamp,required" format:"date-time"`
	Action       string                         `json:"action"`
	Actor        AuditLogListResponseDataActor  `json:"actor"`
	ResourceID   string                         `json:"resource_id"`
	ResourceType string                         `json:"resource_type"`
	Status       AuditLogListResponseDataStatus `json:"status"`
	JSON         auditLogListResponseDataJSON   `json:"-"`
}

// auditLogListResponseDataJSON contains the JSON metadata for the struct
// [AuditLogListResponseData]
type auditLogListResponseDataJSON struct {
	ID           apijson.Field
	Timestamp    apijson.Field
	Action       apijson.Field
	Actor        apijson.Field
	ResourceID   apijson.Field
	ResourceType apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AuditLogListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogListResponseDataActor struct {
	ID    string                            `json:"id,required"`
	Name  string                            `json:"name,required"`
	Email string                            `json:"email"`
	JSON  auditLogListResponseDataActorJSON `json:"-"`
}

// auditLogListResponseDataActorJSON contains the JSON metadata for the struct
// [AuditLogListResponseDataActor]
type auditLogListResponseDataActorJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseDataActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogListResponseDataStatus string

const (
	AuditLogListResponseDataStatusSuccess AuditLogListResponseDataStatus = "success"
	AuditLogListResponseDataStatusFailure AuditLogListResponseDataStatus = "failure"
	AuditLogListResponseDataStatusPending AuditLogListResponseDataStatus = "pending"
)

type AuditLogListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// RFC 3339 timestamp of the earliest audit log to return. Cannot be used with
	// 'next_page'.
	StartingOn param.Field[time.Time] `query:"starting_on" format:"date-time"`
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
