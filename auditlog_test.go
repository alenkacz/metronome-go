// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/metronome/metronome-go"
	"github.com/metronome/metronome-go/internal/testutil"
	"github.com/metronome/metronome-go/option"
)

func TestAuditLogGetAuditLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.AuditLogs.GetAuditLogs(context.TODO(), metronome.AuditLogGetAuditLogsParams{
		Limit:      metronome.F(int64(1)),
		NextPage:   metronome.F("string"),
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
