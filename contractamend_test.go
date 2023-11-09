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

func TestContractAmendAmendContractWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Amends.AmendContract(context.TODO(), metronome.ContractAmendAmendContractParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		StartingAt: metronome.F(time.Now()),
		Commits: metronome.F([]metronome.ContractAmendAmendContractParamsCommit{{
			Type:      metronome.F(metronome.ContractAmendAmendContractParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsCommitsAccessScheduleScheduleItem{{
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
			InvoiceSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:               metronome.F(0.000000),
			Description:          metronome.F("string"),
			RolloverFraction:     metronome.F(0.000000),
			ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableTags:       metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			Type:      metronome.F(metronome.ContractAmendAmendContractParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsCommitsAccessScheduleScheduleItem{{
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
			InvoiceSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:               metronome.F(0.000000),
			Description:          metronome.F("string"),
			RolloverFraction:     metronome.F(0.000000),
			ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableTags:       metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			Type:      metronome.F(metronome.ContractAmendAmendContractParamsCommitsTypePrepaid),
			Name:      metronome.F("x"),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsCommitsAccessScheduleScheduleItem{{
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
			InvoiceSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			Amount:               metronome.F(0.000000),
			Description:          metronome.F("string"),
			RolloverFraction:     metronome.F(0.000000),
			ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableTags:       metronome.F([]string{"string", "string", "string"}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}}),
		Discounts: metronome.F([]metronome.ContractAmendAmendContractParamsDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsDiscountsScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsDiscountsScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsDiscountsScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}}),
		NetsuiteSalesOrderID: metronome.F("1234"),
		Overrides: metronome.F([]metronome.ContractAmendAmendContractParamsOverride{{
			ProductID:    metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Tags:         metronome.F([]string{"string", "string", "string"}),
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractAmendAmendContractParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractAmendAmendContractParamsOverridesOverwriteRate{
				RateType:      metronome.F(metronome.ContractAmendAmendContractParamsOverridesOverwriteRateRateTypeFlat),
				Price:         metronome.F(0.000000),
				UseListPrices: metronome.F(true),
			}),
		}, {
			ProductID:    metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Tags:         metronome.F([]string{"string", "string", "string"}),
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractAmendAmendContractParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractAmendAmendContractParamsOverridesOverwriteRate{
				RateType:      metronome.F(metronome.ContractAmendAmendContractParamsOverridesOverwriteRateRateTypeFlat),
				Price:         metronome.F(0.000000),
				UseListPrices: metronome.F(true),
			}),
		}, {
			ProductID:    metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Tags:         metronome.F([]string{"string", "string", "string"}),
			StartingAt:   metronome.F(time.Now()),
			EndingBefore: metronome.F(time.Now()),
			Entitled:     metronome.F(true),
			Type:         metronome.F(metronome.ContractAmendAmendContractParamsOverridesTypeOverwrite),
			Multiplier:   metronome.F(0.000000),
			OverwriteRate: metronome.F(metronome.ContractAmendAmendContractParamsOverridesOverwriteRate{
				RateType:      metronome.F(metronome.ContractAmendAmendContractParamsOverridesOverwriteRateRateTypeFlat),
				Price:         metronome.F(0.000000),
				UseListPrices: metronome.F(true),
			}),
		}}),
		ResellerRoyalties: metronome.F([]metronome.ContractAmendAmendContractParamsResellerRoyalty{{
			ResellerType:          metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("string"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("string"),
				AwsPayerReferenceID: metronome.F("string"),
				AwsOfferID:          metronome.F("string"),
			}),
			GcpOptions: metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("string"),
				GcpOfferID:   metronome.F("string"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("string"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("string"),
				AwsPayerReferenceID: metronome.F("string"),
				AwsOfferID:          metronome.F("string"),
			}),
			GcpOptions: metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("string"),
				GcpOfferID:   metronome.F("string"),
			}),
		}, {
			ResellerType:          metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesResellerTypeAws),
			Fraction:              metronome.F(0.000000),
			NetsuiteResellerID:    metronome.F("string"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			StartingAt:            metronome.F(time.Now()),
			EndingBefore:          metronome.F(time.Now()),
			ResellerContractValue: metronome.F(0.000000),
			AwsOptions: metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("string"),
				AwsPayerReferenceID: metronome.F("string"),
				AwsOfferID:          metronome.F("string"),
			}),
			GcpOptions: metronome.F(metronome.ContractAmendAmendContractParamsResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("string"),
				GcpOfferID:   metronome.F("string"),
			}),
		}}),
		SalesforceOpportunityID: metronome.F("006opportunity1"),
		ScheduledCharges: metronome.F([]metronome.ContractAmendAmendContractParamsScheduledCharge{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsScheduledChargesScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsScheduledChargesScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}, {
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:      metronome.F("x"),
			Schedule: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesSchedule{
				ScheduleItems: metronome.F([]metronome.ContractAmendAmendContractParamsScheduledChargesScheduleScheduleItem{{
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
				RecurringSchedule: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringSchedule{
					StartingAt:         metronome.F(time.Now()),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					UnitPrice:          metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					Amount:             metronome.F(0.000000),
					AmountDistribution: metronome.F(metronome.ContractAmendAmendContractParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
				}),
			}),
			NetsuiteSalesOrderID: metronome.F("string"),
		}}),
		TotalContractValue: metronome.F(0.000000),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
