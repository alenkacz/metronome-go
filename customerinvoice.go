// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// CustomerInvoiceService contains methods and other services that help with
// interacting with the example API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomerInvoiceService] method
// instead.
type CustomerInvoiceService struct {
Options []option.RequestOption
}

// NewCustomerInvoiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerInvoiceService(opts ...option.RequestOption) (r *CustomerInvoiceService) {
  r = &CustomerInvoiceService{}
  r.Options = opts
  return
}
