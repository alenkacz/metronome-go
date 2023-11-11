// File generated from our OpenAPI spec by Stainless.

package example_test

import (
  "testing"
  "context"
  "net/http/httputil"
  "github.com/example/example-go"
  "github.com/example/example-go/option"
  "github.com/example/example-go/internal/testutil"
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
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Contracts.New(context.TODO(), example.ContractNewParams{
    CustomerID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    StartingAt: example.F(time.Now()),
    Commits: example.F([]example.ContractNewParamsCommit{example.ContractNewParamsCommit{
      Type: example.F(example.ContractNewParamsCommitsTypePrepaid),
      Name: example.F("x"),
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: example.F(example.ContractNewParamsCommitsAccessSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsCommitsAccessScheduleScheduleItem{example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: example.F(example.ContractNewParamsCommitsInvoiceSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: example.F(0.000000),
      Description: example.F("string"),
      RolloverFraction: example.F(0.000000),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: example.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractNewParamsCommit{
      Type: example.F(example.ContractNewParamsCommitsTypePrepaid),
      Name: example.F("x"),
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: example.F(example.ContractNewParamsCommitsAccessSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsCommitsAccessScheduleScheduleItem{example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: example.F(example.ContractNewParamsCommitsInvoiceSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: example.F(0.000000),
      Description: example.F("string"),
      RolloverFraction: example.F(0.000000),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: example.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractNewParamsCommit{
      Type: example.F(example.ContractNewParamsCommitsTypePrepaid),
      Name: example.F("x"),
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: example.F(example.ContractNewParamsCommitsAccessSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsCommitsAccessScheduleScheduleItem{example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractNewParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: example.F(example.ContractNewParamsCommitsInvoiceSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: example.F(0.000000),
      Description: example.F("string"),
      RolloverFraction: example.F(0.000000),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: example.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: example.F("string"),
    }}),
    Discounts: example.F([]example.ContractNewParamsDiscount{example.ContractNewParamsDiscount{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractNewParamsDiscountsSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsDiscountsScheduleScheduleItem{example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsDiscountsScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractNewParamsDiscount{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractNewParamsDiscountsSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsDiscountsScheduleScheduleItem{example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsDiscountsScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractNewParamsDiscount{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractNewParamsDiscountsSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsDiscountsScheduleScheduleItem{example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsDiscountsScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }}),
    EndingBefore: example.F(time.Now()),
    Name: example.F("string"),
    NetPaymentTermsDays: example.F(0.000000),
    NetsuiteSalesOrderID: example.F("1234"),
    Overrides: example.F([]example.ContractNewParamsOverride{example.ContractNewParamsOverride{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      Entitled: example.F(true),
      Type: example.F(example.ContractNewParamsOverridesTypeOverwrite),
      Multiplier: example.F(0.000000),
      OverwriteRate: example.F(example.RateParam{
        RateType: example.F(example.RateRateTypeFlat),
        Price: example.F(0.000000),
        UseListPrices: example.F(true),
      }),
    }, example.ContractNewParamsOverride{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      Entitled: example.F(true),
      Type: example.F(example.ContractNewParamsOverridesTypeOverwrite),
      Multiplier: example.F(0.000000),
      OverwriteRate: example.F(example.RateParam{
        RateType: example.F(example.RateRateTypeFlat),
        Price: example.F(0.000000),
        UseListPrices: example.F(true),
      }),
    }, example.ContractNewParamsOverride{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      Entitled: example.F(true),
      Type: example.F(example.ContractNewParamsOverridesTypeOverwrite),
      Multiplier: example.F(0.000000),
      OverwriteRate: example.F(example.RateParam{
        RateType: example.F(example.RateRateTypeFlat),
        Price: example.F(0.000000),
        UseListPrices: example.F(true),
      }),
    }}),
    RateCardID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    ResellerRoyalties: example.F([]example.ContractNewParamsResellerRoyalty{example.ContractNewParamsResellerRoyalty{
      ResellerType: example.F(example.ContractNewParamsResellerRoyaltiesResellerTypeAws),
      Fraction: example.F(0.000000),
      NetsuiteResellerID: example.F("string"),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      ResellerContractValue: example.F(0.000000),
      AwsOptions: example.F(example.ContractNewParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: example.F("string"),
        AwsPayerReferenceID: example.F("string"),
        AwsOfferID: example.F("string"),
      }),
      GcpOptions: example.F(example.ContractNewParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: example.F("string"),
        GcpOfferID: example.F("string"),
      }),
    }, example.ContractNewParamsResellerRoyalty{
      ResellerType: example.F(example.ContractNewParamsResellerRoyaltiesResellerTypeAws),
      Fraction: example.F(0.000000),
      NetsuiteResellerID: example.F("string"),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      ResellerContractValue: example.F(0.000000),
      AwsOptions: example.F(example.ContractNewParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: example.F("string"),
        AwsPayerReferenceID: example.F("string"),
        AwsOfferID: example.F("string"),
      }),
      GcpOptions: example.F(example.ContractNewParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: example.F("string"),
        GcpOfferID: example.F("string"),
      }),
    }, example.ContractNewParamsResellerRoyalty{
      ResellerType: example.F(example.ContractNewParamsResellerRoyaltiesResellerTypeAws),
      Fraction: example.F(0.000000),
      NetsuiteResellerID: example.F("string"),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      ResellerContractValue: example.F(0.000000),
      AwsOptions: example.F(example.ContractNewParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: example.F("string"),
        AwsPayerReferenceID: example.F("string"),
        AwsOfferID: example.F("string"),
      }),
      GcpOptions: example.F(example.ContractNewParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: example.F("string"),
        GcpOfferID: example.F("string"),
      }),
    }}),
    SalesforceOpportunityID: example.F("006opportunity1"),
    ScheduledCharges: example.F([]example.ContractNewParamsScheduledCharge{example.ContractNewParamsScheduledCharge{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractNewParamsScheduledChargesSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsScheduledChargesScheduleScheduleItem{example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractNewParamsScheduledCharge{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractNewParamsScheduledChargesSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsScheduledChargesScheduleScheduleItem{example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractNewParamsScheduledCharge{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractNewParamsScheduledChargesSchedule{
        ScheduleItems: example.F([]example.ContractNewParamsScheduledChargesScheduleScheduleItem{example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractNewParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }}),
    TotalContractValue: example.F(0.000000),
    Transition: example.F(example.ContractNewParamsTransition{
      Type: example.F(example.ContractNewParamsTransitionTypeSupersede),
      FromContractID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
    }),
    UsageFilter: example.F(example.ContractNewParamsUsageFilter{
      GroupKey: example.F("string"),
      GroupValues: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
    }),
    UsageInvoiceSchedule: example.F(example.ContractNewParamsUsageInvoiceSchedule{
      Frequency: example.F(example.ContractNewParamsUsageInvoiceScheduleFrequencyMonthly),
    }),
  })
  if err != nil {
    var apierr *example.Error
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
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Contracts.List(context.TODO(), example.ContractListParams{
    CustomerID: example.F("9b85c1c1-5238-4f2a-a409-61412905e1e1"),
    IncludeLedgers: example.F(true),
  })
  if err != nil {
    var apierr *example.Error
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
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Contracts.Amend(context.TODO(), example.ContractAmendParams{
    ContractID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    StartingAt: example.F(time.Now()),
    Commits: example.F([]example.ContractAmendParamsCommit{example.ContractAmendParamsCommit{
      Type: example.F(example.ContractAmendParamsCommitsTypePrepaid),
      Name: example.F("x"),
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: example.F(example.ContractAmendParamsCommitsAccessSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsCommitsAccessScheduleScheduleItem{example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: example.F(example.ContractAmendParamsCommitsInvoiceSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: example.F(0.000000),
      Description: example.F("string"),
      RolloverFraction: example.F(0.000000),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: example.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractAmendParamsCommit{
      Type: example.F(example.ContractAmendParamsCommitsTypePrepaid),
      Name: example.F("x"),
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: example.F(example.ContractAmendParamsCommitsAccessSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsCommitsAccessScheduleScheduleItem{example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: example.F(example.ContractAmendParamsCommitsInvoiceSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: example.F(0.000000),
      Description: example.F("string"),
      RolloverFraction: example.F(0.000000),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: example.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractAmendParamsCommit{
      Type: example.F(example.ContractAmendParamsCommitsTypePrepaid),
      Name: example.F("x"),
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      AccessSchedule: example.F(example.ContractAmendParamsCommitsAccessSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsCommitsAccessScheduleScheduleItem{example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsAccessScheduleScheduleItem{
          Amount: example.F(0.000000),
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
        }}),
      }),
      InvoiceSchedule: example.F(example.ContractAmendParamsCommitsInvoiceSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsCommitsInvoiceScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      Amount: example.F(0.000000),
      Description: example.F("string"),
      RolloverFraction: example.F(0.000000),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      ApplicableTags: example.F([]string{"string", "string", "string"}),
      NetsuiteSalesOrderID: example.F("string"),
    }}),
    Discounts: example.F([]example.ContractAmendParamsDiscount{example.ContractAmendParamsDiscount{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractAmendParamsDiscountsSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsDiscountsScheduleScheduleItem{example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsDiscountsScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractAmendParamsDiscount{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractAmendParamsDiscountsSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsDiscountsScheduleScheduleItem{example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsDiscountsScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractAmendParamsDiscount{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractAmendParamsDiscountsSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsDiscountsScheduleScheduleItem{example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsDiscountsScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsDiscountsScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }}),
    NetsuiteSalesOrderID: example.F("1234"),
    Overrides: example.F([]example.ContractAmendParamsOverride{example.ContractAmendParamsOverride{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      Entitled: example.F(true),
      Type: example.F(example.ContractAmendParamsOverridesTypeOverwrite),
      Multiplier: example.F(0.000000),
      OverwriteRate: example.F(example.RateParam{
        RateType: example.F(example.RateRateTypeFlat),
        Price: example.F(0.000000),
        UseListPrices: example.F(true),
      }),
    }, example.ContractAmendParamsOverride{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      Entitled: example.F(true),
      Type: example.F(example.ContractAmendParamsOverridesTypeOverwrite),
      Multiplier: example.F(0.000000),
      OverwriteRate: example.F(example.RateParam{
        RateType: example.F(example.RateRateTypeFlat),
        Price: example.F(0.000000),
        UseListPrices: example.F(true),
      }),
    }, example.ContractAmendParamsOverride{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Tags: example.F([]string{"string", "string", "string"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      Entitled: example.F(true),
      Type: example.F(example.ContractAmendParamsOverridesTypeOverwrite),
      Multiplier: example.F(0.000000),
      OverwriteRate: example.F(example.RateParam{
        RateType: example.F(example.RateRateTypeFlat),
        Price: example.F(0.000000),
        UseListPrices: example.F(true),
      }),
    }}),
    ResellerRoyalties: example.F([]example.ContractAmendParamsResellerRoyalty{example.ContractAmendParamsResellerRoyalty{
      ResellerType: example.F(example.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
      Fraction: example.F(0.000000),
      NetsuiteResellerID: example.F("string"),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      ResellerContractValue: example.F(0.000000),
      AwsOptions: example.F(example.ContractAmendParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: example.F("string"),
        AwsPayerReferenceID: example.F("string"),
        AwsOfferID: example.F("string"),
      }),
      GcpOptions: example.F(example.ContractAmendParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: example.F("string"),
        GcpOfferID: example.F("string"),
      }),
    }, example.ContractAmendParamsResellerRoyalty{
      ResellerType: example.F(example.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
      Fraction: example.F(0.000000),
      NetsuiteResellerID: example.F("string"),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      ResellerContractValue: example.F(0.000000),
      AwsOptions: example.F(example.ContractAmendParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: example.F("string"),
        AwsPayerReferenceID: example.F("string"),
        AwsOfferID: example.F("string"),
      }),
      GcpOptions: example.F(example.ContractAmendParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: example.F("string"),
        GcpOfferID: example.F("string"),
      }),
    }, example.ContractAmendParamsResellerRoyalty{
      ResellerType: example.F(example.ContractAmendParamsResellerRoyaltiesResellerTypeAws),
      Fraction: example.F(0.000000),
      NetsuiteResellerID: example.F("string"),
      ApplicableProductIDs: example.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
      StartingAt: example.F(time.Now()),
      EndingBefore: example.F(time.Now()),
      ResellerContractValue: example.F(0.000000),
      AwsOptions: example.F(example.ContractAmendParamsResellerRoyaltiesAwsOptions{
        AwsAccountNumber: example.F("string"),
        AwsPayerReferenceID: example.F("string"),
        AwsOfferID: example.F("string"),
      }),
      GcpOptions: example.F(example.ContractAmendParamsResellerRoyaltiesGcpOptions{
        GcpAccountID: example.F("string"),
        GcpOfferID: example.F("string"),
      }),
    }}),
    SalesforceOpportunityID: example.F("006opportunity1"),
    ScheduledCharges: example.F([]example.ContractAmendParamsScheduledCharge{example.ContractAmendParamsScheduledCharge{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractAmendParamsScheduledChargesSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsScheduledChargesScheduleScheduleItem{example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractAmendParamsScheduledCharge{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractAmendParamsScheduledChargesSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsScheduledChargesScheduleScheduleItem{example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }, example.ContractAmendParamsScheduledCharge{
      ProductID: example.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
      Name: example.F("x"),
      Schedule: example.F(example.ContractAmendParamsScheduledChargesSchedule{
        ScheduleItems: example.F([]example.ContractAmendParamsScheduledChargesScheduleScheduleItem{example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }, example.ContractAmendParamsScheduledChargesScheduleScheduleItem{
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          Timestamp: example.F(time.Now()),
        }}),
        RecurringSchedule: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringSchedule{
          StartingAt: example.F(time.Now()),
          EndingBefore: example.F(time.Now()),
          Frequency: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
          UnitPrice: example.F(0.000000),
          Quantity: example.F(0.000000),
          Amount: example.F(0.000000),
          AmountDistribution: example.F(example.ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
        }),
      }),
      NetsuiteSalesOrderID: example.F("string"),
    }}),
    TotalContractValue: example.F(0.000000),
  })
  if err != nil {
    var apierr *example.Error
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
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Contracts.Get(context.TODO(), example.ContractGetParams{
    ContractID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    IncludeLedgers: example.F(true),
  })
  if err != nil {
    var apierr *example.Error
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
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  err := client.Contracts.SetUsageFilter(context.TODO(), example.ContractSetUsageFilterParams{
    ContractID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    GroupKey: example.F("business_subscription_id"),
    GroupValues: example.F([]string{"ID-1", "ID-2"}),
    StartingAt: example.F(time.Now()),
  })
  if err != nil {
    var apierr *example.Error
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
  client := example.NewClient(
    option.WithBaseURL(baseURL),
    option.WithBearerToken("My Bearer Token"),
  )
  _, err := client.Contracts.UpdateEndDate(context.TODO(), example.ContractUpdateEndDateParams{
    ContractID: example.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
    CustomerID: example.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    EndingBefore: example.F(time.Now()),
  })
  if err != nil {
    var apierr *example.Error
    if errors.As(err, &apierr) {
      t.Log(string(apierr.DumpRequest(true)))
    }
    t.Fatalf("err should be nil: %s", err.Error())
  }
}
