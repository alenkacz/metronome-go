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

func TestContractPricingProductCreateNewProductWithOptionalParams(t *testing.T) {
  baseURL := "http://localhost:4010"
  if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
    baseURL = envURL
  }
  if !testutil.CheckTestServer(t, baseURL) {
    return
  }
  client := metronome.NewClient(
    option.WithBaseURL(baseURL),
    option.WithAPIKey("My API Key"),
  )
  _, err := client.ContractPricings.Products.Creates.NewProduct(context.TODO(), metronome.ContractPricingProductCreateNewProductParams{
    Name: metronome.F("My Product"),
    Type: metronome.F(metronome.ContractPricingProductCreateNewProductParamsTypeUsage),
    BillableMetricID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    CompositeProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
    CompositeTags: metronome.F([]string{"string", "string", "string"}),
    IsRefundable: metronome.F(true),
    NetsuiteInternalItemID: metronome.F("12345"),
    NetsuiteOverageItemID: metronome.F("12346"),
    StartingAt: metronome.F(time.Now()),
    Tags: metronome.F([]string{"string", "string", "string"}),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
