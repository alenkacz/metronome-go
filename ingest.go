// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// IngestService contains methods and other services that help with interacting
// with the example API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewIngestService] method instead.
type IngestService struct {
Options []option.RequestOption
}

// NewIngestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIngestService(opts ...option.RequestOption) (r *IngestService) {
  r = &IngestService{}
  r.Options = opts
  return
}
