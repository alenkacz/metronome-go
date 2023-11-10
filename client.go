// File generated from our OpenAPI spec by Stainless.

package example

import (
  "context"
  "github.com/example/example-go/option"
  "github.com/example/example-go/internal/shared"
)

// Client creates a struct with services and top level methods that help with
// interacting with the example API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
Options []option.RequestOption
ContractPricing *ContractPricingService
Contracts *ContractService
Alerts *AlertService
CustomerAlerts *CustomerAlertService
Plans *PlanService
Credits *CreditService
CreditTypes *CreditTypeService
Customers *CustomerService
Ingest *IngestService
Usage *UsageService
AuditLogs *AuditLogService
CustomFields *CustomFieldService
}

// NewClient generates a new client with the default option read from the
// environment (EXAMPLE_BEARER_TOKEN). The option passed in as arguments are
// applied after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
  defaults := []option.RequestOption{option.WithEnvironmentProduction()}
  if o, ok := os.LookupEnv("EXAMPLE_BEARER_TOKEN"); ok {
    defaults = append(defaults, option.WithBearerToken(o))
  }
  opts = append(defaults, opts...)

  r = &Client{Options: opts}

  r.ContractPricing = NewContractPricingService(opts...)
  r.Contracts = NewContractService(opts...)
  r.Alerts = NewAlertService(opts...)
  r.CustomerAlerts = NewCustomerAlertService(opts...)
  r.Plans = NewPlanService(opts...)
  r.Credits = NewCreditService(opts...)
  r.CreditTypes = NewCreditTypeService(opts...)
  r.Customers = NewCustomerService(opts...)
  r.Ingest = NewIngestService(opts...)
  r.Usage = NewUsageService(opts...)
  r.AuditLogs = NewAuditLogService(opts...)
  r.CustomFields = NewCustomFieldService(opts...)

  return
}
