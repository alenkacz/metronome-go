// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// CustomerBillingConfigService contains methods and other services that help with
// interacting with the example API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerBillingConfigService]
// method instead.
type CustomerBillingConfigService struct {
Options []option.RequestOption
}

// NewCustomerBillingConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerBillingConfigService(opts ...option.RequestOption) (r *CustomerBillingConfigService) {
  r = &CustomerBillingConfigService{}
  r.Options = opts
  return
}
