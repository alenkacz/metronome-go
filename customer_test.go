// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/metronome/metronome-go"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/testutil"
  "time"
)

func TestCustomerNewWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.New(context.TODO(), metronome.CustomerNewParams{
    Name: metronome.F("Example, Inc."),
    BillingConfig: metronome.F(metronome.CustomerNewParamsBillingConfig{
      BillingProviderType: metronome.F(metronome.CustomerNewParamsBillingConfigBillingProviderTypeAwsMarketplace),
      BillingProviderCustomerID: metronome.F("string"),
      StripeCollectionMethod: metronome.F(metronome.CustomerNewParamsBillingConfigStripeCollectionMethodChargeAutomatically),
      AwsProductCode: metronome.F("string"),
      AwsRegion: metronome.F(metronome.CustomerNewParamsBillingConfigAwsRegionAfSouth1),
    }),
    ExternalID: metronome.F("team@example.com"),
    IngestAliases: metronome.F([]string{"x", "x", "x"}),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerGet(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.Get(context.TODO(), "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc")
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerListWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.List(context.TODO(), metronome.CustomerListParams{
    CustomerIDs: metronome.F([]string{"string", "string", "string"}),
    IngestAlias: metronome.F("string"),
    Limit: metronome.F(int64(1)),
    NextPage: metronome.F("string"),
    OnlyArchived: metronome.F(true),
    SalesforceAccountIDs: metronome.F([]string{"string", "string", "string"}),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerArchive(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.Archive(context.TODO(), metronome.CustomerArchiveParams{
    ID: metronome.F("8deed800-1b7a-495d-a207-6c52bac54dc9"),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerBillableMetricsWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.BillableMetrics(
    context.TODO(),
    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    metronome.CustomerBillableMetricsParams{
      Limit: metronome.F(int64(1)),
      NextPage: metronome.F("string"),
      OnCurrentPlan: metronome.F(true),
    },
  )
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerCostsWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.Costs(
    context.TODO(),
    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    metronome.CustomerCostsParams{
      EndingBefore: metronome.F(time.Now()),
      StartingOn: metronome.F(time.Now()),
      Limit: metronome.F(int64(1)),
      NextPage: metronome.F("string"),
    },
  )
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerSetIngestAliases(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  err := client.Customers.SetIngestAliases(
    context.TODO(),
    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    metronome.CustomerSetIngestAliasesParams{
      IngestAliases: metronome.F([]string{"team@example.com"}),
    },
  )
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerSetName(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Customers.SetName(
    context.TODO(),
    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    metronome.CustomerSetNameParams{
      Name: metronome.F("Example, Inc."),
    },
  )
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestCustomerUpdateConfigWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  err := client.Customers.UpdateConfig(
    context.TODO(),
    "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    metronome.CustomerUpdateConfigParams{
      SalesforceAccountID: metronome.F("0015500001WO1ZiABL"),
    },
  )
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
