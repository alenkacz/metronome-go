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

func TestContractPricingRateCardAddRateAddRateWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricings.RateCards.AddRates.AddRate(context.TODO(), metronome.ContractPricingRateCardAddRateAddRateParams{
		Entitled:      metronome.F(true),
		Price:         metronome.F(100.000000),
		ProductID:     metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		RateCardID:    metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RateType:      metronome.F(metronome.ContractPricingRateCardAddRateAddRateParamsRateTypeFlat),
		StartingAt:    metronome.F(time.Now()),
		EndingBefore:  metronome.F(time.Now()),
		UseListPrices: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
