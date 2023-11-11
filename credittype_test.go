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

func TestCreditTypeListWithOptionalParams(t *testing.T) {
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
  _, err := client.CreditTypes.List(context.TODO(), metronome.CreditTypeListParams{
    Limit: metronome.F(int64(1)),
    NextPage: metronome.F("string"),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
