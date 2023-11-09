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

func TestContractCreateNewContractWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Creates.NewContract(context.TODO(), metronome.ContractCreateNewContractParams{
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		StartingAt: metronome.F(time.Now()),
		Commits: metronome.F([]metronome.ContractCreateNewContractParamsCommit{{
			Type:      metronome.F(metronome.ContractCreateNewContractParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:               metronome.F(0.000000),
			Description:          metronome.F("string"),
			RolloverFraction:     metronome.F(0.000000),
			ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableTags:       metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			Type:      metronome.F(metronome.ContractCreateNewContractParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:               metronome.F(0.000000),
			Description:          metronome.F("string"),
			RolloverFraction:     metronome.F(0.000000),
			ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableTags:       metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			Type:      metronome.F(metronome.ContractCreateNewContractParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}, {
					Amount:       metronome.F(0.000000),
					StartingAt:   metronome.F(time.Now()),
					EndingBefore: metronome.F(time.Now()),
				}}),
			}),
			InvoiceSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:               metronome.F(0.000000),
			Description:          metronome.F("string"),
			RolloverFraction:     metronome.F(0.000000),
			ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableTags:       metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}}),
		Discounts: metronome.F([]metronome.ContractCreateNewContractParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractCreateNewContractParamsDiscountsSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractCreateNewContractParamsDiscountsSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractCreateNewContractParamsDiscountsSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsDiscountsScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}}),
		EndingBefore:         metronome.F(time.Now()),
		Name:                 metronome.F("string"),
		NetPaymentTermsDays:  metronome.F(0.000000),
		NetsuiteSalesOrderID: metronome.F("1234"),
		Overrides: metronome.F([]metronome.ContractCreateNewContractParamsOverride{{
			ProductID:    metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Tags:         metronome.F([]string{"string", "string", "string"}),
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractCreateNewContractParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractCreateNewContractParamsOverridesOverwriteRate{
				RateType:      metronome.F(metronome.ContractCreateNewContractParamsOverridesOverwriteRateRateTypeFlat),
				Price:         metronome.F(0.000000),
				UseListPrices: metronome.F(true),
			}),
		}, {
			ProductID:    metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Tags:         metronome.F([]string{"string", "string", "string"}),
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractCreateNewContractParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractCreateNewContractParamsOverridesOverwriteRate{
				RateType:      metronome.F(metronome.ContractCreateNewContractParamsOverridesOverwriteRateRateTypeFlat),
				Price:         metronome.F(0.000000),
				UseListPrices: metronome.F(true),
			}),
		}, {
			ProductID:    metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Tags:         metronome.F([]string{"string", "string", "string"}),
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractCreateNewContractParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractCreateNewContractParamsOverridesOverwriteRate{
				RateType:      metronome.F(metronome.ContractCreateNewContractParamsOverridesOverwriteRateRateTypeFlat),
				Price:         metronome.F(0.000000),
				UseListPrices: metronome.F(true),
			}),
		}}),
		RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		ResellerRoyalties: metronome.F([]metronome.ContractCreateNewContractParamsResellerRoyalty{{
			ResellerType:          metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("string"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("string"),
				AwsPayerReferenceID: metronome.F("string"),
				AwsOfferID:          metronome.F("string"),
			}),
			GcpOptions: metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("string"),
				GcpOfferID:   metronome.F("string"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("string"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("string"),
				AwsPayerReferenceID: metronome.F("string"),
				AwsOfferID:          metronome.F("string"),
			}),
			GcpOptions: metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("string"),
				GcpOfferID:   metronome.F("string"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("string"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("string"),
				AwsPayerReferenceID: metronome.F("string"),
				AwsOfferID:          metronome.F("string"),
			}),
			GcpOptions: metronome.F(metronome.ContractCreateNewContractParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("string"),
				GcpOfferID:   metronome.F("string"),
			}),
		}}),
		SalesforceOpportunityID: metronome.F("006opportunity1"),
		ScheduledCharges: metronome.F([]metronome.ContractCreateNewContractParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesSchedule{
				ScheduleItems: metronome.F([]metronome.ContractCreateNewContractParamsScheduledChargesScheduleScheduleItem{{
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}, {
					UnitPrice: metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Amount:    metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
				}}),
				RecurringSchedule: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractCreateNewContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}}),
		TotalContractValue: metronome.F(0.000000),
		Transition: metronome.F(metronome.ContractCreateNewContractParamsTransition{
			Type:           metronome.F(metronome.ContractCreateNewContractParamsTransitionTypeSupersede),
			FromContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		UsageFilter: metronome.F(metronome.ContractCreateNewContractParamsUsageFilter{
			GroupKey:    metronome.F("string"),
			GroupValues: metronome.F([]string{"string", "string", "string"}),
			StartingAt:  metronome.F(time.Now()),
		}),
		UsageInvoiceSchedule: metronome.F(metronome.ContractCreateNewContractParamsUsageInvoiceSchedule{
			Frequency: metronome.F(metronome.ContractCreateNewContractParamsUsageInvoiceScheduleFrequencyMonthly),
		}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
