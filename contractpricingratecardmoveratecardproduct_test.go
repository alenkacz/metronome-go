// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/metronome/metronome-go"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/testutil"
)

func TestContractPricingRateCardMoveRateCardProductMoveRateCardProducts(t *testing.T) {
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
  _, err := client.ContractPricings.RateCards.MoveRateCardProducts.MoveRateCardProducts(context.TODO(), metronome.ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParams{
    ProductMoves: metronome.F([]metronome.ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParamsProductMove{metronome.ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParamsProductMove{
      ProductID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
      Position: metronome.F(0.000000),
    }, metronome.ContractPricingRateCardMoveRateCardProductMoveRateCardProductsParamsProductMove{
      ProductID: metronome.F("b086f2f4-9851-4466-9ca0-30d53e6a42ac"),
      Position: metronome.F(1.000000),
    }}),
    RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
