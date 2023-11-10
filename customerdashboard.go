// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// CustomerDashboardService contains methods and other services that help with
// interacting with the example API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerDashboardService] method
// instead.
type CustomerDashboardService struct {
Options []option.RequestOption
}

// NewCustomerDashboardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerDashboardService(opts ...option.RequestOption) (r *CustomerDashboardService) {
  r = &CustomerDashboardService{}
  r.Options = opts
  return
}
