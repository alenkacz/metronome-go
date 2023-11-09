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

func TestContractPricingProductUpdateUpdateProductWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricings.Products.Updates.UpdateProduct(context.TODO(), metronome.ContractPricingProductUpdateUpdateProductParams{
		ProductID:              metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		StartingAt:             metronome.F(time.Now()),
		BillableMetricID:       metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CompositeProductIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CompositeTags:          metronome.F([]string{"string", "string", "string"}),
		IsRefundable:           metronome.F(true),
		Name:                   metronome.F("My Updated Product"),
		NetsuiteInternalItemID: metronome.F("string"),
		NetsuiteOverageItemID:  metronome.F("string"),
		Tags:                   metronome.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
