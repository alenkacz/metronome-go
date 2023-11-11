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
    option.WithBearerToken("My Bearer Token"),
  )
  err := client.Ingest(context.TODO(), metronome.IngestParams{
    Body: metronome.F([]metronome.IngestParamsBody{metronome.IngestParamsBody{
      TransactionID: metronome.F("2021-01-01T00:00:00+00:00_cluster42"),
      CustomerID: metronome.F("team@example.com"),
      EventType: metronome.F("heartbeat"),
      Timestamp: metronome.F("2021-01-01T00:00:00+00:00"),
    }}),
  })
  if err != nil {
    t.Error(err)
  }
}
ageGetUsageParamsBillableMetric{
      ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      GroupBy: metronome.F(metronome.UsageGetUsageParamsBillableMetricsGroupBy{
        Key: metronome.F("string"),
        Values: metronome.F([]string{"x"}),
      }),
    }, metronome.UsageGetUsageParamsBillableMetric{
      ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      GroupBy: metronome.F(metronome.UsageGetUsageParamsBillableMetricsGroupBy{
        Key: metronome.F("string"),
        Values: metronome.F([]string{"x"}),
      }),
    }, metronome.UsageGetUsageParamsBillableMetric{
      ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      GroupBy: metronome.F(metronome.UsageGetUsageParamsBillableMetricsGroupBy{
        Key: metronome.F("string"),
        Values: metronome.F([]string{"x"}),
      }),
    }}),
    CustomerIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestUsageGetUsageGroupsWithOptionalParams(t *testing.T) {
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
  _, err := client.Usage.GetUsageGroups(context.TODO(), metronome.UsageGetUsageGroupsParams{
    BillableMetricID: metronome.F("222796fd-d29c-429e-89b2-549fabda4ed6"),
    CustomerID: metronome.F("04ca7e72-4229-4a6e-ab11-9f7376fccbcb"),
    WindowSize: metronome.F(metronome.UsageGetUsageGroupsParamsWindowSizeDay),
    Limit: metronome.F(int64(1)),
    NextPage: metronome.F("string"),
    CurrentPeriod: metronome.F(true),
    EndingBefore: metronome.F(time.Now()),
    GroupBy: metronome.F(metronome.UsageGetUsageGroupsParamsGroupBy{
      Key: metronome.F("region"),
      Values: metronome.F([]string{"US-East", "US-West", "EU-Central"}),
    }),
    StartingOn: metronome.F(time.Now()),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
