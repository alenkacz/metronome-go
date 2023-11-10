// File generated from our OpenAPI spec by Stainless.

package example_test

import (
  "testing"
  "github.com/example/example-go/internal/testutil"
  "github.com/example/example-go"
  "github.com/example/example-go/option"
  "context"
  "time"
)

func TestUsage(t *testing.T) {
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
  contractNewResponse, err := client.Contracts.New(context.TODO(), example.ContractNewParams{
    CustomerID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    StartingAt: example.F(time.Now()),
  })
  if err != nil {
    t.Error(err)
  }
  t.Logf("%+v\n", contractNewResponse.Data)
}
