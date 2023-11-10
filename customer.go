// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
)

// CustomerService contains methods and other services that help with interacting
// with the example API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCustomerService] method instead.
type CustomerService struct {
Options []option.RequestOption
Plans *CustomerPlanService
Archive *CustomerArchiveService
BillableMetrics *CustomerBillableMetricService
Invoices *CustomerInvoiceService
BillingConfig *CustomerBillingConfigService
Dashboards *CustomerDashboardService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
  r = &CustomerService{}
  r.Options = opts
  r.Plans = NewCustomerPlanService(opts...)
  r.Archive = NewCustomerArchiveService(opts...)
  r.BillableMetrics = NewCustomerBillableMetricService(opts...)
  r.Invoices = NewCustomerInvoiceService(opts...)
  r.BillingConfig = NewCustomerBillingConfigService(opts...)
  r.Dashboards = NewCustomerDashboardService(opts...)
  return
}
