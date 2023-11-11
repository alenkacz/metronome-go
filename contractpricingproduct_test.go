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

func TestContractPricingProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricing.Products.New(context.TODO(), metronome.ContractPricingProductNewParams{
		Name:                   metronome.F("My Product"),
		Type:                   metronome.F(metronome.ContractPricingProductNewParamsTypeUsage),
		BillableMetricID:       metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CompositeProductIDs:    metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		CompositeTags:          metronome.F([]string{"string", "string", "string"}),
		IsRefundable:           metronome.F(true),
		NetsuiteInternalItemID: metronome.F("12345"),
		NetsuiteOverageItemID:  metronome.F("12346"),
		StartingAt:             metronome.F(time.Now()),
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

func TestContractPricingProductGet(t *testing.T) {
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
	_, err := client.ContractPricing.Products.Get(context.TODO(), metronome.ContractPricingProductGetParams{
		ID: metronome.F("d84e7f4e-7a70-4fe4-be02-7a5027beffcc"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractPricingProductUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricing.Products.Update(context.TODO(), metronome.ContractPricingProductUpdateParams{
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

func TestContractPricingProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricing.Products.List(context.TODO(), metronome.ContractPricingProductListParams{
		Body:     metronome.F[any](map[string]interface{}{}),
		Limit:    metronome.F(int64(1)),
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
