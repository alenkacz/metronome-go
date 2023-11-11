// File generated from our OpenAPI spec by Stainless.

package metronome

import (
	"github.com/metronome/metronome-go/internal/apierror"
	"github.com/metronome/metronome-go/internal/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type Commit = shared.Commit

// This is an alias to an internal type.
type CommitProduct = shared.CommitProduct

// This is an alias to an internal type.
type CommitType = shared.CommitType

// This is an alias to an internal value.
const CommitTypePrepaid = shared.CommitTypePrepaid

// This is an alias to an internal value.
const CommitTypePostpaid = shared.CommitTypePostpaid

// Only valid for "PREPAID" commits: The schedule that the customer will gain
// access to the credits purposed with this commit.
//
// This is an alias to an internal type.
type CommitAccessSchedule = shared.CommitAccessSchedule

// This is an alias to an internal type.
type CommitAccessScheduleScheduleItem = shared.CommitAccessScheduleScheduleItem

// This is an alias to an internal type.
type CommitLedger = shared.CommitLedger

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSegmentStartLedgerEntry = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitSegmentStartLedgerEntryType = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart = shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction = shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitRolloverLedgerEntry = shared.CommitLedgerPrepaidCommitRolloverLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitRolloverLedgerEntryType = shared.CommitLedgerPrepaidCommitRolloverLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover = shared.CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitExpirationLedgerEntry = shared.CommitLedgerPrepaidCommitExpirationLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitExpirationLedgerEntryType = shared.CommitLedgerPrepaidCommitExpirationLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration = shared.CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCanceledLedgerEntry = shared.CommitLedgerPrepaidCommitCanceledLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCanceledLedgerEntryType = shared.CommitLedgerPrepaidCommitCanceledLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled = shared.CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCreditedLedgerEntry = shared.CommitLedgerPrepaidCommitCreditedLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPrepaidCommitCreditedLedgerEntryType = shared.CommitLedgerPrepaidCommitCreditedLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited = shared.CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance = shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction = shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitTrueupLedgerEntry = shared.CommitLedgerPostpaidCommitTrueupLedgerEntry

// This is an alias to an internal type.
type CommitLedgerPostpaidCommitTrueupLedgerEntryType = shared.CommitLedgerPostpaidCommitTrueupLedgerEntryType

// This is an alias to an internal value.
const CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup = shared.CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup

// This is an alias to an internal type.
type CommitRolledOverFrom = shared.CommitRolledOverFrom

// This is an alias to an internal type.
type ContractWithoutAmendments = shared.ContractWithoutAmendments

// This is an alias to an internal type.
type ContractWithoutAmendmentsResellerRoyalty = shared.ContractWithoutAmendmentsResellerRoyalty

// This is an alias to an internal type.
type ContractWithoutAmendmentsResellerRoyaltiesResellerType = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerType

// This is an alias to an internal value.
const ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws

// This is an alias to an internal value.
const ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp = shared.ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp

// This is an alias to an internal type.
type ContractWithoutAmendmentsTransition = shared.ContractWithoutAmendmentsTransition

// This is an alias to an internal type.
type ContractWithoutAmendmentsTransitionsType = shared.ContractWithoutAmendmentsTransitionsType

// This is an alias to an internal value.
const ContractWithoutAmendmentsTransitionsTypeSupersede = shared.ContractWithoutAmendmentsTransitionsTypeSupersede

// This is an alias to an internal value.
const ContractWithoutAmendmentsTransitionsTypeRenewal = shared.ContractWithoutAmendmentsTransitionsTypeRenewal

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageInvoiceSchedule = shared.ContractWithoutAmendmentsUsageInvoiceSchedule

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageInvoiceScheduleFrequency = shared.ContractWithoutAmendmentsUsageInvoiceScheduleFrequency

// This is an alias to an internal value.
const ContractWithoutAmendmentsUsageInvoiceScheduleFrequencyMonthly = shared.ContractWithoutAmendmentsUsageInvoiceScheduleFrequencyMonthly

// This is an alias to an internal value.
const ContractWithoutAmendmentsUsageInvoiceScheduleFrequencyQuarterly = shared.ContractWithoutAmendmentsUsageInvoiceScheduleFrequencyQuarterly

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageFilter = shared.ContractWithoutAmendmentsUsageFilter

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageFilterCurrent = shared.ContractWithoutAmendmentsUsageFilterCurrent

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageFilterInitial = shared.ContractWithoutAmendmentsUsageFilterInitial

// This is an alias to an internal type.
type ContractWithoutAmendmentsUsageFilterUpdate = shared.ContractWithoutAmendmentsUsageFilterUpdate

// This is an alias to an internal type.
type Discount = shared.Discount

// This is an alias to an internal type.
type DiscountProduct = shared.DiscountProduct

// This is an alias to an internal type.
type ID = shared.ID

// This is an alias to an internal type.
type Override = shared.Override

// This is an alias to an internal type.
type OverrideType = shared.OverrideType

// This is an alias to an internal value.
const OverrideTypeOverwrite = shared.OverrideTypeOverwrite

// This is an alias to an internal value.
const OverrideTypeMultiplier = shared.OverrideTypeMultiplier

// This is an alias to an internal type.
type OverrideProduct = shared.OverrideProduct

// This is an alias to an internal type.
type Rate = shared.Rate

// This is an alias to an internal type.
type RateRateType = shared.RateRateType

// This is an alias to an internal value.
const RateRateTypeFlat = shared.RateRateTypeFlat

// This is an alias to an internal value.
const RateRateTypePercentage = shared.RateRateTypePercentage

// This is an alias to an internal type.
type RateParam = shared.RateParam

// This is an alias to an internal type.
type SchedulePointInTime = shared.SchedulePointInTime

// This is an alias to an internal type.
type SchedulePointInTimeScheduleItem = shared.SchedulePointInTimeScheduleItem

// This is an alias to an internal type.
type ScheduledCharge = shared.ScheduledCharge

// This is an alias to an internal type.
type ScheduledChargeProduct = shared.ScheduledChargeProduct
