// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// CustomerBillableMetricService contains methods and other services that help with
// interacting with the example API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerBillableMetricService]
// method instead.
type CustomerBillableMetricService struct {
Options []option.RequestOption
}

// NewCustomerBillableMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerBillableMetricService(opts ...option.RequestOption) (r *CustomerBillableMetricService) {
  r = &CustomerBillableMetricService{}
  r.Options = opts
  return
}
