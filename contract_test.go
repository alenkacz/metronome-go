// File generated from our OpenAPI spec by Stainless.

package metronome_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/metronome/metronome-go"
  "github.com/metronome/metronome-go/option"
  "github.com/metronome/metronome-go/internal/testutil"
  "time"
)

func TestContractNewWithOptionalParams(t *testing.T) {
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
  _, err := client.Contracts.New(context.TODO(), metronome.ContractNewParams{
    CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    StartingAt: metronome.F(time.Now()),
    Commits: metronome.F([]metronome.ContractNewParamsCommit{metronome.ContractNewParamsCommit{
      Type: metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
      Name: metronome.F("x"),
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: metronome.F(0.000000),
      Description: metronome.F("string"),
      RolloverFraction: metronome.F(0.000000),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: metronome.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractNewParamsCommit{
      Type: metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
      Name: metronome.F("x"),
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: metronome.F(0.000000),
      Description: metronome.F("string"),
      RolloverFraction: metronome.F(0.000000),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: metronome.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractNewParamsCommit{
      Type: metronome.F(metronome.ContractNewParamsCommitsTypePrepaid),
      Name: metronome.F("x"),
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: metronome.F(metronome.ContractNewParamsCommitsAccessSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: metronome.F(0.000000),
      Description: metronome.F("string"),
      RolloverFraction: metronome.F(0.000000),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: metronome.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: metronome.F("string"),
    }}),
    Discounts: metronome.F([]metronome.ContractNewParamsDiscount{metronome.ContractNewParamsDiscount{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractNewParamsDiscount{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractNewParamsDiscount{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractNewParamsDiscountsSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsDiscountsScheduleScheduleItem{metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }}),
    EndingBefore: metronome.F(time.Now()),
    Name: metronome.F("string"),
    NetPaymentTermsDays: metronome.F(0.000000),
    NetsuiteSalesOrderID: metronome.F("1234"),
    Overrides: metronome.F([]metronome.ContractNewParamsOverride{metronome.ContractNewParamsOverride{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      Entitled: metronome.F(true),
      Type: metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
      Multiplier: metronome.F(0.000000),
      OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
        RateType: metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
        Price: metronome.F(0.000000),
        UseListPrices: metronome.F(true),
      }),
    }, metronome.ContractNewParamsOverride{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      Entitled: metronome.F(true),
      Type: metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
      Multiplier: metronome.F(0.000000),
      OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
        RateType: metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
        Price: metronome.F(0.000000),
        UseListPrices: metronome.F(true),
      }),
    }, metronome.ContractNewParamsOverride{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      Entitled: metronome.F(true),
      Type: metronome.F(metronome.ContractNewParamsOverridesTypeOverwrite),
      Multiplier: metronome.F(0.000000),
      OverwriteRate: metronome.F(metronome.ContractNewParamsOverridesOverwriteRate{
        RateType: metronome.F(metronome.ContractNewParamsOverridesOverwriteRateRateTypeFlat),
        Price: metronome.F(0.000000),
        UseListPrices: metronome.F(true),
      }),
    }}),
    RateCardID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    ResellerRoyalties: metronome.F([]metronome.ContractNewParamsResellerRoyalty{metronome.ContractNewParamsResellerRoyalty{
      ResellerType: metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
      Fraction: metronome.F(0.000000),
      NetsuiteResellerID: metronome.F("string"),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      ResellerContractValue: metronome.F(0.000000),
      AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: metronome.F("string"),
        AwsPayerReferenceID: metronome.F("string"),
        AwsOfferID: metronome.F("string"),
      }),
      GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: metronome.F("string"),
        GcpOfferID: metronome.F("string"),
      }),
    }, metronome.ContractNewParamsResellerRoyalty{
      ResellerType: metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
      Fraction: metronome.F(0.000000),
      NetsuiteResellerID: metronome.F("string"),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      ResellerContractValue: metronome.F(0.000000),
      AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: metronome.F("string"),
        AwsPayerReferenceID: metronome.F("string"),
        AwsOfferID: metronome.F("string"),
      }),
      GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: metronome.F("string"),
        GcpOfferID: metronome.F("string"),
      }),
    }, metronome.ContractNewParamsResellerRoyalty{
      ResellerType: metronome.F(metronome.ContractNewParamsResellerRoyaltiesResellerTypeAws),
      Fraction: metronome.F(0.000000),
      NetsuiteResellerID: metronome.F("string"),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      ResellerContractValue: metronome.F(0.000000),
      AwsOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: metronome.F("string"),
        AwsPayerReferenceID: metronome.F("string"),
        AwsOfferID: metronome.F("string"),
      }),
      GcpOptions: metronome.F(metronome.ContractNewParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: metronome.F("string"),
        GcpOfferID: metronome.F("string"),
      }),
    }}),
    SalesforceOpportunityID: metronome.F("006opportunity1"),
    ScheduledCharges: metronome.F([]metronome.ContractNewParamsScheduledCharge{metronome.ContractNewParamsScheduledCharge{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractNewParamsScheduledCharge{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractNewParamsScheduledCharge{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractNewParamsScheduledChargesSchedule{
        ScheduleItems: metronome.F([]metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }}),
    TotalContractValue: metronome.F(0.000000),
    Transition: metronome.F(metronome.ContractNewParamsTransition{
      Type: metronome.F(metronome.ContractNewParamsTransitionTypeSupersede),
      FromContractID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
    }),
    UsageFilter: metronome.F(metronome.ContractNewParamsUsageFilter{
      GroupKey: metronome.F("string"),
      GroupValues: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
    }),
    UsageInvoiceSchedule: metronome.F(metronome.ContractNewParamsUsageInvoiceSchedule{
      Frequency: metronome.F(metronome.ContractNewParamsUsageInvoiceScheduleFrequencyMonthly),
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

func TestContractListWithOptionalParams(t *testing.T) {
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
  _, err := client.Contracts.List(context.TODO(), metronome.ContractListParams{
    CustomerID: metronome.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
    IncludeLedgers: metronome.F(true),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractAmendWithOptionalParams(t *testing.T) {
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
  _, err := client.Contracts.Amend(context.TODO(), metronome.ContractAmendParams{
    ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    StartingAt: metronome.F(time.Now()),
    Commits: metronome.F([]metronome.ContractAmendParamsCommit{metronome.ContractAmendParamsCommit{
      Type: metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
      Name: metronome.F("x"),
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: metronome.F(0.000000),
      Description: metronome.F("string"),
      RolloverFraction: metronome.F(0.000000),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: metronome.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractAmendParamsCommit{
      Type: metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
      Name: metronome.F("x"),
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: metronome.F(0.000000),
      Description: metronome.F("string"),
      RolloverFraction: metronome.F(0.000000),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: metronome.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractAmendParamsCommit{
      Type: metronome.F(metronome.ContractAmendParamsCommitsTypePrepaid),
      Name: metronome.F("x"),
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: metronome.F(metronome.ContractAmendParamsCommitsAccessSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: metronome.F(0.000000),
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: metronome.F(0.000000),
      Description: metronome.F("string"),
      RolloverFraction: metronome.F(0.000000),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: metronome.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: metronome.F("string"),
    }}),
    Discounts: metronome.F([]metronome.ContractAmendParamsDiscount{metronome.ContractAmendParamsDiscount{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractAmendParamsDiscount{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractAmendParamsDiscount{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractAmendParamsDiscountsSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsDiscountsScheduleScheduleItem{metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }}),
    NetsuiteSalesOrderID: metronome.F("1234"),
    Overrides: metronome.F([]metronome.ContractAmendParamsOverride{metronome.ContractAmendParamsOverride{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      Entitled: metronome.F(true),
      Type: metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
      Multiplier: metronome.F(0.000000),
      OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
        RateType: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
        Price: metronome.F(0.000000),
        UseListPrices: metronome.F(true),
      }),
    }, metronome.ContractAmendParamsOverride{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      Entitled: metronome.F(true),
      Type: metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
      Multiplier: metronome.F(0.000000),
      OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
        RateType: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
        Price: metronome.F(0.000000),
        UseListPrices: metronome.F(true),
      }),
    }, metronome.ContractAmendParamsOverride{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: metronome.F([]string{"string", "string", "string"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      Entitled: metronome.F(true),
      Type: metronome.F(metronome.ContractAmendParamsOverridesTypeOverwrite),
      Multiplier: metronome.F(0.000000),
      OverwriteRate: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRate{
        RateType: metronome.F(metronome.ContractAmendParamsOverridesOverwriteRateRateTypeFlat),
        Price: metronome.F(0.000000),
        UseListPrices: metronome.F(true),
      }),
    }}),
    ResellerRoyalties: metronome.F([]metronome.ContractAmendParamsResellerRoyalty{metronome.ContractAmendParamsResellerRoyalty{
      ResellerType: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
      Fraction: metronome.F(0.000000),
      NetsuiteResellerID: metronome.F("string"),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      ResellerContractValue: metronome.F(0.000000),
      AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: metronome.F("string"),
        AwsPayerReferenceID: metronome.F("string"),
        AwsOfferID: metronome.F("string"),
      }),
      GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: metronome.F("string"),
        GcpOfferID: metronome.F("string"),
      }),
    }, metronome.ContractAmendParamsResellerRoyalty{
      ResellerType: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
      Fraction: metronome.F(0.000000),
      NetsuiteResellerID: metronome.F("string"),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      ResellerContractValue: metronome.F(0.000000),
      AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: metronome.F("string"),
        AwsPayerReferenceID: metronome.F("string"),
        AwsOfferID: metronome.F("string"),
      }),
      GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: metronome.F("string"),
        GcpOfferID: metronome.F("string"),
      }),
    }, metronome.ContractAmendParamsResellerRoyalty{
      ResellerType: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
      Fraction: metronome.F(0.000000),
      NetsuiteResellerID: metronome.F("string"),
      ApplicableProductIDs: metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: metronome.F(time.Now()),
      EndingBefore: metronome.F(time.Now()),
      ResellerContractValue: metronome.F(0.000000),
      AwsOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: metronome.F("string"),
        AwsPayerReferenceID: metronome.F("string"),
        AwsOfferID: metronome.F("string"),
      }),
      GcpOptions: metronome.F(metronome.ContractAmendParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: metronome.F("string"),
        GcpOfferID: metronome.F("string"),
      }),
    }}),
    SalesforceOpportunityID: metronome.F("006opportunity1"),
    ScheduledCharges: metronome.F([]metronome.ContractAmendParamsScheduledCharge{metronome.ContractAmendParamsScheduledCharge{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractAmendParamsScheduledCharge{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: metronome.F("string"),
    }, metronome.ContractAmendParamsScheduledCharge{
      ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: metronome.F("x"),
      Schedule: metronome.F(metronome.ContractAmendParamsScheduledChargesSchedule{
        ScheduleItems: metronome.F([]metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }, metronome.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          Timestamp: metronome.F(time.Now()),
        }}),
        RecurringSchedule: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: metronome.F(time.Now()),
          EndingBefore: metronome.F(time.Now()),
          Frequency: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: metronome.F(0.000000),
          Quantity: metronome.F(0.000000),
          Amount: metronome.F(0.000000),
          AmountDistribution: metronome.F(metronome.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
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

func TestContractGetWithOptionalParams(t *testing.T) {
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
  _, err := client.Contracts.Get(context.TODO(), metronome.ContractGetParams{
    ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    IncludeLedgers: metronome.F(true),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractSetUsageFilter(t *testing.T) {
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
  err := client.Contracts.SetUsageFilter(context.TODO(), metronome.ContractSetUsageFilterParams{
    ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    GroupKey: metronome.F("business_subscription_id"),
    GroupValues: metronome.F([]string{"ID-1", "ID-2"}),
    StartingAt: metronome.F(time.Now()),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}

func TestContractUpdateEndDateWithOptionalParams(t *testing.T) {
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
  _, err := client.Contracts.UpdateEndDate(context.TODO(), metronome.ContractUpdateEndDateParams{
    ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    EndingBefore: metronome.F(time.Now()),
  })
  if err != nil {
    var apierr *metronome.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
