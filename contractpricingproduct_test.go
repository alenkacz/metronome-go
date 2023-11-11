// File generated from our OpenAPI spec by Stainless.

package example_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/example/example-go"
  "github.com/example/example-go/option"
  "github.com/example/example-go/internal/testutil"
  "time"
)

func TestContractPricingProductNewWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.ContractPricing.Products.New(context.TODO(), example.ContractPricingProductNewParams{
    Name: example.F("My Product"),
    Type: example.F(example.ContractPricingProductNewParamsTypeUsage),
    BillableMetricID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    CompositeProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
    CompositeTags: example.F([]string{"string", "string", "string"}),
    IsRefundable: example.F(true),
    NetsuiteInternalItemID: example.F("12345"),
    NetsuiteOverageItemID: example.F("12346"),
    StartingAt: example.F(time.Now()),
    Tags: example.F([]string{"string", "string", "string"}),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingProductUpdateWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.ContractPricing.Products.Update(context.TODO(), example.ContractPricingProductUpdateParams{
    ProductID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    StartingAt: example.F(time.Now()),
    BillableMetricID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
    CompositeProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
    CompositeTags: example.F([]string{"string", "string", "string"}),
    IsRefundable: example.F(true),
    Name: example.F("My Updated Product"),
    NetsuiteInternalItemID: example.F("string"),
    NetsuiteOverageItemID: example.F("string"),
    Tags: example.F([]string{"string", "string", "string"}),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingProductListWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.ContractPricing.Products.List(context.TODO(), example.ContractPricingProductListParams{
    Body: example.F[any](map[string]interface{}{}),
    Limit: example.F(int64(1)),
    NextPage: example.F("string"),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingProductGet(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.ContractPricing.Products.Get(context.TODO(), example.ContractPricingProductGetParams{
    ID: example.F("d84e7f4e-7a70-4fe4-be02-7a5027beffcc"),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
