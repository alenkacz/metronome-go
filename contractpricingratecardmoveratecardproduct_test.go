// File generated from our OpenAPI spec by Stainless.

package example_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/example/example-go"
  "github.com/example/example-go/option"
  "github.com/example/example-go/internal/testutil"
)

func TestContractPricingRateCardMoveRateCardProductUpdate(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.MoveRateCardProducts.Update(context.TODO(), example.ContractPricingRateCardMoveRateCardProductUpdateParams{
    ProductMoves: example.F([]example.ContractPricingRateCardMoveRateCardProductUpdateParamsProductMove{example.ContractPricingRateCardMoveRateCardProductUpdateParamsProductMove{
      ProductID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
      Position: example.F(0.000000),
    }, example.ContractPricingRateCardMoveRateCardProductUpdateParamsProductMove{
      ProductID: example.F("b086f2f4-9851-4466-9ca0-30d53e6a42ac"),
      Position: example.F(1.000000),
    }}),
    RateCardID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
