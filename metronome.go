// File generated from our OpenAPI spec by Stainless.

package metronome

import (
  "github.com/metronome/metronome-go/internal/param"
  "github.com/metronome/metronome-go/internal/apijson"
)

type IngestParams struct {
Body param.Field[[]IngestParamsBody] `json:"body,required"`
}

func (r IngestParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r.Body)
}

type IngestParamsBody struct {
CustomerID param.Field[string] `json:"customer_id,required"`
EventType param.Field[string] `json:"event_type,required"`
// RFC 3339 formatted
Timestamp param.Field[string] `json:"timestamp,required"`
TransactionID param.Field[string] `json:"transaction_id,required"`
Properties param.Field[map[string]interface{}] `json:"properties"`
}

func (r IngestParamsBody) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}
