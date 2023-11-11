// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"context"
	"net/http"
	"os"

	"github.com/metronome/metronome-go/internal/requestconfig"
	"github.com/metronome/metronome-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the metronome API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options         []option.RequestOption
	ContractPricing *ContractPricingService
	Contracts       *ContractService
	Alerts          *AlertService
	CustomerAlerts  *CustomerAlertService
	Plans           *PlanService
	Credits         *CreditService
	CreditTypes     *CreditTypeService
	Customers       *CustomerService
	Dashboards      *DashboardService
	Usage           *UsageService
	AuditLogs       *AuditLogService
	CustomFields    *CustomFieldService
}

// NewClient generates a new client with the default option read from the
// environment (METRONOME_BEARER_TOKEN). The option passed in as arguments are
// applied after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("METRONOME_BEARER_TOKEN"); ok {
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
	r.Dashboards = NewDashboardService(opts...)
	r.Usage = NewUsageService(opts...)
	r.AuditLogs = NewAuditLogService(opts...)
	r.CustomFields = NewCustomFieldService(opts...)

	return
}

// Send usage events to Metronome. The body of this request is expected to be a
// JSON array of between 1 and 100 usage events. Compressed request bodies are
// supported with a `Content-Encoding: gzip` header. See
// [Getting usage into Metronome](https://docs.metronome.com/getting-usage-data-into-metronome/overview)
// to learn more about usage events.
func (r *Client) Ingest(ctx context.Context, body IngestParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}
