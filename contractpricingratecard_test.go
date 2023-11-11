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

func TestContractPricingRateCardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricing.RateCards.New(context.TODO(), metronome.ContractPricingRateCardNewParams{
		Name:        metronome.F("My Rate Card"),
		Description: metronome.F("My Rate Card Description"),
	})
	if err != nil {
		var apierr *metronome.Error
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
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.ContractPricing.RateCards.Get(context.TODO(), metronome.ContractPricingRateCardGetParams{
		ID: metronome.F("f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"),
	})
	if err != nil {
		var apierr *metronome.Error
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
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.ContractPricing.RateCards.Update(context.TODO(), metronome.ContractPricingRateCardUpdateParams{
		RateCardID:  metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Description: metronome.F("My Updated Rate Card Description"),
		Name:        metronome.F("My Updated Rate Card"),
	})
	if err != nil {
		var apierr *metronome.Error
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
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.ContractPricing.RateCards.List(context.TODO(), metronome.ContractPricingRateCardListParams{
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

func TestContractPricingRateCardAddRateWithOptionalParams(t *testing.T) {
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
	_, err := client.ContractPricing.RateCards.AddRate(context.TODO(), metronome.ContractPricingRateCardAddRateParams{
		Entitled:      metronome.F(true),
		Price:         metronome.F(100.000000),
		ProductID:     metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		RateCardID:    metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		RateType:      metronome.F(metronome.ContractPricingRateCardAddRateParamsRateTypeFlat),
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

func TestContractPricingRateCardMoveRateCardProducts(t *testing.T) {
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
	_, err := client.ContractPricing.RateCards.MoveRateCardProducts(context.TODO(), metronome.ContractPricingRateCardMoveRateCardProductsParams{
		ProductMoves: metronome.F([]metronome.ContractPricingRateCardMoveRateCardProductsParamsProductMove{metronome.ContractPricingRateCardMoveRateCardProductsParamsProductMove{
			ProductID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
			Position:  metronome.F(0.000000),
		}, metronome.ContractPricingRateCardMoveRateCardProductsParamsProductMove{
			ProductID: metronome.F("b086f2f4-9851-4466-9ca0-30d53e6a42ac"),
			Position:  metronome.F(1.000000),
		}}),
		RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
	})
	if err != nil {
		var apierr *metronome.Error
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
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.ContractPricing.RateCards.SetRateCardProductsOrder(context.TODO(), metronome.ContractPricingRateCardSetRateCardProductsOrderParams{
		ProductOrder: metronome.F([]string{"13117714-3f05-48e5-a6e9-a66093f13b4d", "b086f2f4-9851-4466-9ca0-30d53e6a42ac"}),
		RateCardID:   metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
