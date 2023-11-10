// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// CustomerAlertService contains methods and other services that help with
// interacting with the example API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerAlertService] method
// instead.
type CustomerAlertService struct {
Options []option.RequestOption
}

// NewCustomerAlertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerAlertService(opts ...option.RequestOption) (r *CustomerAlertService) {
  r = &CustomerAlertService{}
  r.Options = opts
  return
}
