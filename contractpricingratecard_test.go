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

func TestContractPricingRateCardNewWithOptionalParams(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.New(context.TODO(), example.ContractPricingRateCardNewParams{
    Name: example.F("My Rate Card"),
    Description: example.F("My Rate Card Description"),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingRateCardUpdateWithOptionalParams(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.Update(context.TODO(), example.ContractPricingRateCardUpdateParams{
    RateCardID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    Description: example.F("My Updated Rate Card Description"),
    Name: example.F("My Updated Rate Card"),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingRateCardListWithOptionalParams(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.List(context.TODO(), example.ContractPricingRateCardListParams{
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

func TestContractPricingRateCardAddRateWithOptionalParams(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.AddRate(context.TODO(), example.ContractPricingRateCardAddRateParams{
    Entitled: example.F(true),
    Price: example.F(100.000000),
    ProductID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    RateCardID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    RateType: example.F(example.ContractPricingRateCardAddRateParamsRateTypeFlat),
    StartingAt: example.F(time.Now()),
    EndingBefore: example.F(time.Now()),
    UseListPrices: example.F(true),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingRateCardGet(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.Get(context.TODO(), example.ContractPricingRateCardGetParams{
    ID: example.F("f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractPricingRateCardMoveRateCardProducts(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.MoveRateCardProducts(context.TODO(), example.ContractPricingRateCardMoveRateCardProductsParams{
    ProductMoves: example.F([]example.ContractPricingRateCardMoveRateCardProductsParamsProductMove{example.ContractPricingRateCardMoveRateCardProductsParamsProductMove{
      ProductID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
      Position: example.F(0.000000),
    }, example.ContractPricingRateCardMoveRateCardProductsParamsProductMove{
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

func TestContractPricingRateCardSetRateCardProductsOrder(t *testing.T) {
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
  _, err := client.ContractPricing.RateCards.SetRateCardProductsOrder(context.TODO(), example.ContractPricingRateCardSetRateCardProductsOrderParams{
    ProductOrder: example.F([]string{"13117714-3f05-48e5-a6e9-a66093f13b4d", "b086f2f4-9851-4466-9ca0-30d53e6a42ac"}),
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
