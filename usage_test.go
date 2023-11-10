// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
  "testing"
  "github.com/metronome/metronome-go/internal/testutil"
  "github.com/metronome/metronome-go"
  "github.com/metronome/metronome-go/option"
  "context"
)

func TestUsage(t *testing.T) {
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
  contractPricingProductGetGetProductResponse, err := client.ContractPricings.Products.Gets.GetProduct(context.TODO(), metronome.ContractPricingProductGetGetProductParams{
    ID: metronome.F("REPLACE_ME"),
  })
  if err != nil {
    t.Error(err)
  }
  t.Logf("%+v\n", contractPricingProductGetGetProductResponse.Data)
}
