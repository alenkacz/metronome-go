// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
	"context"
	"os"
	"testing"

	"github.com/metronome/metronome-go"
	"github.com/metronome/metronome-go/internal/testutil"
	"github.com/metronome/metronome-go/option"
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
			CustomerID:    metronome.F("team@example.com"),
			EventType:     metronome.F("heartbeat"),
			Timestamp:     metronome.F("2021-01-01T00:00:00+00:00"),
		}}),
	})
	if err != nil {
		t.Error(err)
	}
}
