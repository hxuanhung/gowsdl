package ipipeline

import (
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SubmitComparisonQuoteRequest struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 SubmitComparisonQuoteRequest"`

	Request *SubmitEnhancedProtectionComparisonQuoteRequestMessage `xml:"request,omitempty"`
}

type SubmitComparisonQuoteRequestResponse struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 SubmitComparisonQuoteRequestResponse"`

	SubmitComparisonQuoteRequestResult *SubmitComparisonQuoteResponseMessage `xml:"SubmitComparisonQuoteRequestResult,omitempty"`
}

type RetrieveComparisonQuoteResults struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 RetrieveComparisonQuoteResults"`

	Request *RetrieveComparisonQuoteResultsRequestMessage `xml:"request,omitempty"`
}

type RetrieveComparisonQuoteResultsResponse struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 RetrieveComparisonQuoteResultsResponse"`

	RetrieveComparisonQuoteResultsResult *RetrieveEnhancedProtectionComparisonQuoteResultResponseMessage `xml:"RetrieveComparisonQuoteResultsResult,omitempty"`
}

type RetrieveComparisonQuoteResultsWithInput struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 RetrieveComparisonQuoteResultsWithInput"`

	Request *RetrieveComparisonQuoteResultsRequestMessage `xml:"request,omitempty"`
}

type RetrieveComparisonQuoteResultsWithInputResponse struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 RetrieveComparisonQuoteResultsWithInputResponse"`

	RetrieveComparisonQuoteResultsWithInputResult *RetrieveEnhancedProtectionComparisonQuoteResultsWithInputResponseMessage `xml:"RetrieveComparisonQuoteResultsWithInputResult,omitempty"`
}

type RetrieveComparisonQuoteReport struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 RetrieveComparisonQuoteReport"`

	Request *RetrieveEnhancedProtectionComparisonQuoteReportRequestMessage `xml:"request,omitempty"`
}

type RetrieveComparisonQuoteReportResponse struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/APe2/V1 RetrieveComparisonQuoteReportResponse"`

	RetrieveComparisonQuoteReportResult *RetrieveComparisonQuoteReportResponseMessage `xml:"RetrieveComparisonQuoteReportResult,omitempty"`
}

type Char int32

type Duration *Duration

type Guid string

type Indexation string

const (
	IndexationAWEI Indexation = "AWEI"

	IndexationRPI Indexation = "RPI"

	IndexationLevel Indexation = "Level"

	IndexationOne Indexation = "One"

	IndexationTwo Indexation = "Two"

	IndexationThree Indexation = "Three"

	IndexationFour Indexation = "Four"

	IndexationFive Indexation = "Five"

	IndexationSix Indexation = "Six"

	IndexationSeven Indexation = "Seven"

	IndexationEight Indexation = "Eight"

	IndexationNine Indexation = "Nine"

	IndexationTen Indexation = "Ten"

	IndexationIncreasing Indexation = "Increasing"
)

type IntegratedIncomeCoverDeferredPeriod string

const (
	IntegratedIncomeCoverDeferredPeriodThirteen IntegratedIncomeCoverDeferredPeriod = "Thirteen"

	IntegratedIncomeCoverDeferredPeriodTwentySix IntegratedIncomeCoverDeferredPeriod = "TwentySix"

	IntegratedIncomeCoverDeferredPeriodFiftyTwo IntegratedIncomeCoverDeferredPeriod = "FiftyTwo"
)

type RequestPremiumType string

const (
	RequestPremiumTypeGuaranteed RequestPremiumType = "Guaranteed"

	RequestPremiumTypeReviewable RequestPremiumType = "Reviewable"

	RequestPremiumTypeAll RequestPremiumType = "All"
)

type TotalPermanentDisabilityType string

const (
	TotalPermanentDisabilityTypeOwnOccupation TotalPermanentDisabilityType = "OwnOccupation"

	TotalPermanentDisabilityTypeSpecifiedWorkTasks TotalPermanentDisabilityType = "SpecifiedWorkTasks"

	TotalPermanentDisabilityTypeActivitiesOfDailyLiving TotalPermanentDisabilityType = "ActivitiesOfDailyLiving"

	TotalPermanentDisabilityTypeFATs TotalPermanentDisabilityType = "FATs"

	TotalPermanentDisabilityTypeSuitedOccupation TotalPermanentDisabilityType = "SuitedOccupation"

	TotalPermanentDisabilityTypeAnyOccupation TotalPermanentDisabilityType = "AnyOccupation"

	TotalPermanentDisabilityTypeNone TotalPermanentDisabilityType = "None"
)

type WaiverOfPremium string

const (
	WaiverOfPremiumNone WaiverOfPremium = "None"

	WaiverOfPremiumFirstLife WaiverOfPremium = "FirstLife"

	WaiverOfPremiumSecondLife WaiverOfPremium = "SecondLife"

	WaiverOfPremiumBothLives WaiverOfPremium = "BothLives"
)

type IndependentLivesAssuredBasis string

const (
	IndependentLivesAssuredBasisFirstLife IndependentLivesAssuredBasis = "FirstLife"

	IndependentLivesAssuredBasisSecondLife IndependentLivesAssuredBasis = "SecondLife"
)

type DeferredPeriod string

const (
	DeferredPeriodZero DeferredPeriod = "Zero"

	DeferredPeriodOne DeferredPeriod = "One"

	DeferredPeriodTwo DeferredPeriod = "Two"

	DeferredPeriodFour DeferredPeriod = "Four"

	DeferredPeriodEight DeferredPeriod = "Eight"

	DeferredPeriodThirteen DeferredPeriod = "Thirteen"

	DeferredPeriodTwentySix DeferredPeriod = "TwentySix"

	DeferredPeriodFiftyTwo DeferredPeriod = "FiftyTwo"

	DeferredPeriodFiftySix DeferredPeriod = "FiftySix"

	DeferredPeriodOneHundredAndFour DeferredPeriod = "OneHundredAndFour"

	DeferredPeriodOneHundredAndTwelve DeferredPeriod = "OneHundredAndTwelve"
)

type BenefitBasis string

const (
	BenefitBasisMaximum BenefitBasis = "Maximum"

	BenefitBasisMonthly BenefitBasis = "Monthly"
)

type BenefitIncreaseBasis string

const (
	BenefitIncreaseBasisPolicyInception BenefitIncreaseBasis = "PolicyInception"

	BenefitIncreaseBasisOnClaim BenefitIncreaseBasis = "OnClaim"
)

type IncomeProtectionIndexation string

const (
	IncomeProtectionIndexationRPI IncomeProtectionIndexation = "RPI"

	IncomeProtectionIndexationLevel IncomeProtectionIndexation = "Level"

	IncomeProtectionIndexationThree IncomeProtectionIndexation = "Three"

	IncomeProtectionIndexationFive IncomeProtectionIndexation = "Five"
)

type BenefitType string

const (
	BenefitTypeTerm BenefitType = "Term"

	BenefitTypeConvertibleTerm BenefitType = "ConvertibleTerm"

	BenefitTypeMortgageProtection BenefitType = "MortgageProtection"

	BenefitTypeFamilyIncomeBenefit BenefitType = "FamilyIncomeBenefit"

	BenefitTypeBusinessProtectionLevel BenefitType = "BusinessProtectionLevel"

	BenefitTypeBusinessProtectionDecreasing BenefitType = "BusinessProtectionDecreasing"

	BenefitTypeWholeOfLife BenefitType = "WholeOfLife"

	BenefitTypeWholeOfLifeOverFiftyPlans BenefitType = "WholeOfLifeOverFiftyPlans"

	BenefitTypeRelevantLife BenefitType = "RelevantLife"

	BenefitTypeGiftInterVivos BenefitType = "GiftInterVivos"

	BenefitTypeIncomeProtection BenefitType = "IncomeProtection"
)

type EmploymentStatus string

const (
	EmploymentStatusEmployed EmploymentStatus = "Employed"

	EmploymentStatusSelfEmployed EmploymentStatus = "SelfEmployed"

	EmploymentStatusHousePerson EmploymentStatus = "HousePerson"
)

type Residency string

const (
	ResidencyUK Residency = "UK"

	ResidencyJersey Residency = "Jersey"

	ResidencyGuernsey Residency = "Guernsey"

	ResidencyIsleOfMan Residency = "IsleOfMan"

	ResidencyOther Residency = "Other"
)

type QuotationBasis string

const (
	QuotationBasisSumAssured QuotationBasis = "SumAssured"

	QuotationBasisPremium QuotationBasis = "Premium"
)

type RefineOptions string

const (
	RefineOptionsNone RefineOptions = "None"

	RefineOptionsProviderExtranet RefineOptions = "ProviderExtranet"
)

type ResponsePremiumType string

const (
	ResponsePremiumTypeGuaranteed ResponsePremiumType = "Guaranteed"

	ResponsePremiumTypeReviewable ResponsePremiumType = "Reviewable"
)

type PeriodType string

const (
	PeriodTypeTerm PeriodType = "Term"

	PeriodTypeAge PeriodType = "Age"
)

type ResponseDefinitionOfIncapacity string

const (
	ResponseDefinitionOfIncapacityNotSupplied ResponseDefinitionOfIncapacity = "NotSupplied"

	ResponseDefinitionOfIncapacityOwnOccupation ResponseDefinitionOfIncapacity = "OwnOccupation"

	ResponseDefinitionOfIncapacityAnyOccupation ResponseDefinitionOfIncapacity = "AnyOccupation"

	ResponseDefinitionOfIncapacitySuitedOccupation ResponseDefinitionOfIncapacity = "SuitedOccupation"

	ResponseDefinitionOfIncapacityFATs ResponseDefinitionOfIncapacity = "FATs"

	ResponseDefinitionOfIncapacityActivitiesOfDailyLiving ResponseDefinitionOfIncapacity = "ActivitiesOfDailyLiving"
)

type EnhancedProtectionComparisonReportType string

const (
	EnhancedProtectionComparisonReportTypeSummaryWithCommission EnhancedProtectionComparisonReportType = "SummaryWithCommission"

	EnhancedProtectionComparisonReportTypeSummaryWithoutCommission EnhancedProtectionComparisonReportType = "SummaryWithoutCommission"
)

type SubmitEnhancedProtectionComparisonQuoteRequestMessage struct {
	*Message

	ComparisonQuoteRequest *EnhancedProtectionComparisonQuoteRequest `xml:"ComparisonQuoteRequest,omitempty"`

	OverrideOptions *OverrideOptions `xml:"OverrideOptions,omitempty"`
}

type EnhancedProtectionComparisonQuoteRequest struct {
	*ComparisonQuoteRequestWithProviders

	CommissionOptions *CommissionOptions `xml:"CommissionOptions,omitempty"`

	PolicyBasis *PolicyBasis `xml:"PolicyBasis,omitempty"`
}

type CommissionOptions struct {
	CommissionPercentage int32 `xml:"CommissionPercentage,omitempty"`

	CommissionType *CommissionType `xml:"CommissionType,omitempty"`

	IndemnityRequired bool `xml:"IndemnityRequired,omitempty"`
}

type PolicyBasis struct {
	Benefits *ArrayOfRequestBenefit `xml:"Benefits,omitempty"`

	LivesAssured *LivesAssured `xml:"LivesAssured,omitempty"`

	NumberOfBenefitsChosen int32 `xml:"NumberOfBenefitsChosen,omitempty"`

	QuotationBasis *QuotationBasis `xml:"QuotationBasis,omitempty"`
}

type ArrayOfRequestBenefit struct {
	RequestBenefit []*RequestBenefit `xml:"RequestBenefit,omitempty"`
}

type RequestBenefit struct {
	Benefit *RequestBaseBenefit `xml:"Benefit,omitempty"`

	BenefitType *BenefitType `xml:"BenefitType,omitempty"`
}

type RequestBaseBenefit struct {
}

type RequestTermBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllness bool `xml:"CriticalIllness,omitempty"`

	CriticalIllnessOptions *RequestCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	FractureCoverBasis *LivesAssuredBasisWithNone `xml:"FractureCoverBasis,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	IntegratedIncomeCoverOptions *IntegratedIncomeCoverOptions `xml:"IntegratedIncomeCoverOptions,omitempty"`

	Life bool `xml:"Life,omitempty"`

	LifeCoverBuyBack bool `xml:"LifeCoverBuyBack,omitempty"`

	LifeOrEarlierCriticalIllness bool `xml:"LifeOrEarlierCriticalIllness,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestCriticalIllnessOptions struct {
	CriticalIllnessBuyBack bool `xml:"CriticalIllnessBuyBack,omitempty"`

	CriticalIllnessCoverAmount int32 `xml:"CriticalIllnessCoverAmount,omitempty"`

	LifeOrEarlierCriticalIllnessCoverAmount int32 `xml:"LifeOrEarlierCriticalIllnessCoverAmount,omitempty"`
}

type IntegratedIncomeCoverOptions struct {
	FirstLife *IntegratedIncomeCoverPerson `xml:"FirstLife,omitempty"`

	LivesAssuredBasis *LivesAssuredBasisWithNone `xml:"LivesAssuredBasis,omitempty"`

	SecondLife *IntegratedIncomeCoverPerson `xml:"SecondLife,omitempty"`
}

type IntegratedIncomeCoverPerson struct {
	CoverAmount int32 `xml:"CoverAmount,omitempty"`

	IntegratedIncomeCoverDeferredPeriod *IntegratedIncomeCoverDeferredPeriod `xml:"IntegratedIncomeCoverDeferredPeriod,omitempty"`
}

type RequestConvertibleTermBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllness bool `xml:"CriticalIllness,omitempty"`

	CriticalIllnessOptions *RequestCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	FractureCoverBasis *LivesAssuredBasisWithNone `xml:"FractureCoverBasis,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	Life bool `xml:"Life,omitempty"`

	LifeCoverBuyBack bool `xml:"LifeCoverBuyBack,omitempty"`

	LifeOrEarlierCriticalIllness bool `xml:"LifeOrEarlierCriticalIllness,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestBusinessProtectionLevelBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllness bool `xml:"CriticalIllness,omitempty"`

	CriticalIllnessOptions *RequestCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	FractureCover bool `xml:"FractureCover,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	IntegratedIncomeCoverOptions *IntegratedIncomeCoverOptions `xml:"IntegratedIncomeCoverOptions,omitempty"`

	Life bool `xml:"Life,omitempty"`

	LifeCoverBuyBack bool `xml:"LifeCoverBuyBack,omitempty"`

	LifeOrEarlierCriticalIllness bool `xml:"LifeOrEarlierCriticalIllness,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestMortgageProtectionBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllness bool `xml:"CriticalIllness,omitempty"`

	CriticalIllnessOptions *RequestCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	FractureCoverBasis *LivesAssuredBasisWithNone `xml:"FractureCoverBasis,omitempty"`

	IntegratedIncomeCoverOptions *IntegratedIncomeCoverOptions `xml:"IntegratedIncomeCoverOptions,omitempty"`

	Life bool `xml:"Life,omitempty"`

	LifeCoverBuyBack bool `xml:"LifeCoverBuyBack,omitempty"`

	LifeOrEarlierCriticalIllness bool `xml:"LifeOrEarlierCriticalIllness,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LoanInterestRate float64 `xml:"LoanInterestRate,omitempty"`

	PolicyInterestRate float64 `xml:"PolicyInterestRate,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestBusinessProtectionDecreasingBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllness bool `xml:"CriticalIllness,omitempty"`

	CriticalIllnessOptions *RequestCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	FractureCover bool `xml:"FractureCover,omitempty"`

	IntegratedIncomeCoverOptions *IntegratedIncomeCoverOptions `xml:"IntegratedIncomeCoverOptions,omitempty"`

	Life bool `xml:"Life,omitempty"`

	LifeCoverBuyBack bool `xml:"LifeCoverBuyBack,omitempty"`

	LifeOrEarlierCriticalIllness bool `xml:"LifeOrEarlierCriticalIllness,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LoanInterestRate float64 `xml:"LoanInterestRate,omitempty"`

	PolicyInterestRate float64 `xml:"PolicyInterestRate,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestFamilyIncomeBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllness bool `xml:"CriticalIllness,omitempty"`

	CriticalIllnessOptions *RequestCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	FractureCoverBasis *LivesAssuredBasisWithNone `xml:"FractureCoverBasis,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	IntegratedIncomeCoverOptions *IntegratedIncomeCoverOptions `xml:"IntegratedIncomeCoverOptions,omitempty"`

	Life bool `xml:"Life,omitempty"`

	LifeCoverBuyBack bool `xml:"LifeCoverBuyBack,omitempty"`

	LifeOrEarlierCriticalIllness bool `xml:"LifeOrEarlierCriticalIllness,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestWholeOfLifeBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	WaiverOfPremium *WaiverOfPremium `xml:"WaiverOfPremium,omitempty"`
}

type RequestWholeOfLifeOverFiftyPlansBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *IndependentLivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`
}

type RequestGiftInterVivosBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	LivesAssuredBasis *IndependentLivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`
}

type RequestRelevantLifeBenefit struct {
	*RequestBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	ExcludeLowStart bool `xml:"ExcludeLowStart,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *IndependentLivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`
}

type RequestIncomeProtectionBenefit struct {
	*RequestBaseBenefit

	AdditionalDeferredPeriodDetails *AdditionalDeferredPeriodDetails `xml:"AdditionalDeferredPeriodDetails,omitempty"`

	AnnualDividendIncome float64 `xml:"AnnualDividendIncome,omitempty"`

	BenefitBasis *BenefitBasis `xml:"BenefitBasis,omitempty"`

	BenefitIncreaseBasis *BenefitIncreaseBasis `xml:"BenefitIncreaseBasis,omitempty"`

	ExcludeAgeBanded bool `xml:"ExcludeAgeBanded,omitempty"`

	ExistingBenefit float64 `xml:"ExistingBenefit,omitempty"`

	FractureCoverBasis *LivesAssuredBasisWithNone `xml:"FractureCoverBasis,omitempty"`

	IncludeLimitedPaymentPlans bool `xml:"IncludeLimitedPaymentPlans,omitempty"`

	Indexation *IncomeProtectionIndexation `xml:"Indexation,omitempty"`

	InitialDeferredPeriodDetails *InitialDeferredPeriodDetails `xml:"InitialDeferredPeriodDetails,omitempty"`

	LivesAssuredBasis *IndependentLivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *RequestPremiumType `xml:"PremiumType,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`
}

type AdditionalDeferredPeriodDetails struct {
	BenefitAmount float64 `xml:"BenefitAmount,omitempty"`

	DeferredPeriod *DeferredPeriod `xml:"DeferredPeriod,omitempty"`
}

type InitialDeferredPeriodDetails struct {
	BenefitAmount float64 `xml:"BenefitAmount,omitempty"`

	DeferredPeriod *DeferredPeriod `xml:"DeferredPeriod,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`
}

type LivesAssured struct {
	FirstLife *Person `xml:"FirstLife,omitempty"`

	SecondLife *Person `xml:"SecondLife,omitempty"`
}

type Person struct {
	DateOfBirth time.Time `xml:"DateOfBirth,omitempty"`

	EmploymentStatus *EmploymentStatus `xml:"EmploymentStatus,omitempty"`

	Forename string `xml:"Forename,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	Occupation string `xml:"Occupation,omitempty"`

	Postcode string `xml:"Postcode,omitempty"`

	Residency *Residency `xml:"Residency,omitempty"`

	Salary int32 `xml:"Salary,omitempty"`

	Smoker bool `xml:"Smoker,omitempty"`

	Surname string `xml:"Surname,omitempty"`

	Title string `xml:"Title,omitempty"`
}

type RetrieveEnhancedProtectionComparisonQuoteResultResponseMessage struct {
	*Message

	EnhancedProtectionComparisonQuoteResult *EnhancedProtectionComparisonQuoteResult `xml:"EnhancedProtectionComparisonQuoteResult,omitempty"`
}

type EnhancedProtectionComparisonQuoteResult struct {
	*ComparisonQuoteResult

	EnhancedProtectionQuoteResults *ArrayOfEnhancedProtectionQuoteResult `xml:"EnhancedProtectionQuoteResults,omitempty"`
}

type ArrayOfEnhancedProtectionQuoteResult struct {
	EnhancedProtectionQuoteResult []*EnhancedProtectionQuoteResult `xml:"EnhancedProtectionQuoteResult,omitempty"`
}

type EnhancedProtectionQuoteResult struct {
	*QuoteResult

	AvailableOptions *AvailableOptions `xml:"AvailableOptions,omitempty"`

	Benefits *ArrayOfResponseBenefit `xml:"Benefits,omitempty"`

	Commission *ResultCommission `xml:"Commission,omitempty"`

	PolicyFee float64 `xml:"PolicyFee,omitempty"`

	ProductName string `xml:"ProductName,omitempty"`

	TotalPremium float64 `xml:"TotalPremium,omitempty"`
}

type AvailableOptions struct {
	ExtranetApply bool `xml:"ExtranetApply,omitempty"`

	RefineOption *RefineOptions `xml:"RefineOption,omitempty"`
}

type ArrayOfResponseBenefit struct {
	ResponseBenefit []*ResponseBenefit `xml:"ResponseBenefit,omitempty"`
}

type ResponseBenefit struct {
	Benefit *ResponseBaseBenefit `xml:"Benefit,omitempty"`

	BenefitType *BenefitType `xml:"BenefitType,omitempty"`
}

type ResponseBaseBenefit struct {
}

type ResponseTermBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllnessOptions *ResponseCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	ProviderQuotedOnDifferentOccupation bool `xml:"ProviderQuotedOnDifferentOccupation,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseCriticalIllnessOptions struct {
	CriticalIllnessCoverAmount int32 `xml:"CriticalIllnessCoverAmount,omitempty"`

	LifeOrEarlierCriticalIllnessCoverAmount int32 `xml:"LifeOrEarlierCriticalIllnessCoverAmount,omitempty"`
}

type ResponseConvertibleTermBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllnessOptions *ResponseCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	ProviderQuotedOnDifferentOccupation bool `xml:"ProviderQuotedOnDifferentOccupation,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseBusinessProtectionLevelBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllnessOptions *ResponseCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseMortgageProtectionBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllnessOptions *ResponseCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PolicyInterestRate float64 `xml:"PolicyInterestRate,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	ProviderQuotedOnDifferentOccupation bool `xml:"ProviderQuotedOnDifferentOccupation,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseBusinessProtectionDecreasingBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllnessOptions *ResponseCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PolicyInterestRate float64 `xml:"PolicyInterestRate,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseFamilyIncomeBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	CriticalIllnessOptions *ResponseCriticalIllnessOptions `xml:"CriticalIllnessOptions,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	ProviderQuotedOnDifferentOccupation bool `xml:"ProviderQuotedOnDifferentOccupation,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`

	TotalPermanentDisability *TotalPermanentDisabilityType `xml:"TotalPermanentDisability,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseWholeOfLifeBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	WaiverExpiryAge int32 `xml:"WaiverExpiryAge,omitempty"`
}

type ResponseWholeOfLifeOverFiftyPlansBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LifeCoverDeferredForMonths int32 `xml:"LifeCoverDeferredForMonths,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumCeaseAtAge int32 `xml:"PremiumCeaseAtAge,omitempty"`
}

type ResponseGiftInterVivosBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`
}

type ResponseRelevantLifeBenefit struct {
	*ResponseBaseBenefit

	BenefitAmount int32 `xml:"BenefitAmount,omitempty"`

	Indexation *Indexation `xml:"Indexation,omitempty"`

	LivesAssuredBasis *LivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	LowStart bool `xml:"LowStart,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	TermYears int32 `xml:"TermYears,omitempty"`

	ToAge int32 `xml:"ToAge,omitempty"`
}

type ResponseIncomeProtectionBenefit struct {
	*ResponseBaseBenefit

	AdditionalBenefitAmount float64 `xml:"AdditionalBenefitAmount,omitempty"`

	AdditionalDeferredPeriod *DeferredPeriod `xml:"AdditionalDeferredPeriod,omitempty"`

	AgeBanded bool `xml:"AgeBanded,omitempty"`

	BenefitAmount float64 `xml:"BenefitAmount,omitempty"`

	BenefitFrequency *PremiumFrequency `xml:"BenefitFrequency,omitempty"`

	BenefitIncreaseRate *Indexation `xml:"BenefitIncreaseRate,omitempty"`

	BenefitPeriod *BenefitOrCoverPeriod `xml:"BenefitPeriod,omitempty"`

	CoverPeriod *BenefitOrCoverPeriod `xml:"CoverPeriod,omitempty"`

	DeferredPeriod *DeferredPeriod `xml:"DeferredPeriod,omitempty"`

	DefinitionOfIncapacity *ResponseDefinitionOfIncapacity `xml:"DefinitionOfIncapacity,omitempty"`

	IsLimitedPaymentPlan bool `xml:"IsLimitedPaymentPlan,omitempty"`

	LivesAssuredBasis *IndependentLivesAssuredBasis `xml:"LivesAssuredBasis,omitempty"`

	PremiumAmount float64 `xml:"PremiumAmount,omitempty"`

	PremiumFrequency *PremiumFrequency `xml:"PremiumFrequency,omitempty"`

	PremiumType *ResponsePremiumType `xml:"PremiumType,omitempty"`

	ProviderQuotedOnDifferentOccupation bool `xml:"ProviderQuotedOnDifferentOccupation,omitempty"`

	Renewable bool `xml:"Renewable,omitempty"`
}

type BenefitOrCoverPeriod struct {
	PeriodType *PeriodType `xml:"PeriodType,omitempty"`

	PeriodValue int32 `xml:"PeriodValue,omitempty"`
}

type RetrieveEnhancedProtectionComparisonQuoteResultsWithInputResponseMessage struct {
	*Message

	EnhancedProtectionComparisonQuoteInputData *EnhancedProtectionComparisonQuoteInputData `xml:"EnhancedProtectionComparisonQuoteInputData,omitempty"`

	EnhancedProtectionComparisonQuoteResult *EnhancedProtectionComparisonQuoteResult `xml:"EnhancedProtectionComparisonQuoteResult,omitempty"`
}

type EnhancedProtectionComparisonQuoteInputData struct {
	CommissionOptions *CommissionOptions `xml:"CommissionOptions,omitempty"`

	PolicyBasis *PolicyBasis `xml:"PolicyBasis,omitempty"`
}

type RetrieveEnhancedProtectionComparisonQuoteReportRequestMessage struct {
	*Message

	ComparisonQuoteReportRequest *ComparisonQuoteReportRequest `xml:"ComparisonQuoteReportRequest,omitempty"`
}

type ComparisonQuoteReportRequest struct {
	*ComparisonQuoteReportRequest

	ReportType *EnhancedProtectionComparisonReportType `xml:"ReportType,omitempty"`
}

type BusinessType string

const (
	BusinessTypeDesignatedInvestment BusinessType = "DesignatedInvestment"

	BusinessTypeGeneralInsurance BusinessType = "GeneralInsurance"
)

type IllustrationBasis string

const (
	IllustrationBasisQuote IllustrationBasis = "Quote"

	IllustrationBasisQuoteAndPrint IllustrationBasis = "QuoteAndPrint"
)

type CommissionType string

const (
	CommissionTypeFull CommissionType = "Full"

	CommissionTypeModified CommissionType = "Modified"

	CommissionTypeNil CommissionType = "Nil"
)

type LivesAssuredBasisWithNone string

const (
	LivesAssuredBasisWithNoneNone LivesAssuredBasisWithNone = "None"

	LivesAssuredBasisWithNoneFirstLife LivesAssuredBasisWithNone = "FirstLife"

	LivesAssuredBasisWithNoneSecondLife LivesAssuredBasisWithNone = "SecondLife"

	LivesAssuredBasisWithNoneBothLives LivesAssuredBasisWithNone = "BothLives"
)

type LivesAssuredBasis string

const (
	LivesAssuredBasisFirstLife LivesAssuredBasis = "FirstLife"

	LivesAssuredBasisSecondLife LivesAssuredBasis = "SecondLife"

	LivesAssuredBasisJointLifeFirstDeath LivesAssuredBasis = "JointLifeFirstDeath"

	LivesAssuredBasisJointLifeSecondDeath LivesAssuredBasis = "JointLifeSecondDeath"
)

type PremiumFrequency string

const (
	PremiumFrequencyMonthly PremiumFrequency = "Monthly"

	PremiumFrequencyAnnually PremiumFrequency = "Annually"
)

type Gender string

const (
	GenderMale Gender = "Male"

	GenderFemale Gender = "Female"
)

type ComparisonQuoteSubmissionStatus string

const (
	ComparisonQuoteSubmissionStatusSuccess ComparisonQuoteSubmissionStatus = "Success"

	ComparisonQuoteSubmissionStatusServiceNotAvailable ComparisonQuoteSubmissionStatus = "ServiceNotAvailable"

	ComparisonQuoteSubmissionStatusMessageValidationFailure ComparisonQuoteSubmissionStatus = "MessageValidationFailure"
)

type ComparisonQuoteStatus string

const (
	ComparisonQuoteStatusCompleted ComparisonQuoteStatus = "Completed"

	ComparisonQuoteStatusPending ComparisonQuoteStatus = "Pending"

	ComparisonQuoteStatusError ComparisonQuoteStatus = "Error"

	ComparisonQuoteStatusNoQuotesGenerated ComparisonQuoteStatus = "NoQuotesGenerated"
)

type ProductUnavailableReasonCode string

const (
	ProductUnavailableReasonCodeIncorrectProductType ProductUnavailableReasonCode = "IncorrectProductType"

	ProductUnavailableReasonCodeProviderUnavailable ProductUnavailableReasonCode = "ProviderUnavailable"

	ProductUnavailableReasonCodeTechnicalError ProductUnavailableReasonCode = "TechnicalError"

	ProductUnavailableReasonCodeFilteredOut ProductUnavailableReasonCode = "FilteredOut"

	ProductUnavailableReasonCodeInProgress ProductUnavailableReasonCode = "InProgress"

	ProductUnavailableReasonCodeBusinessError ProductUnavailableReasonCode = "BusinessError"
)

type QuoteStatus string

const (
	QuoteStatusAwaitingResponse QuoteStatus = "AwaitingResponse"

	QuoteStatusBusinessError QuoteStatus = "BusinessError"

	QuoteStatusProviderUnavailable QuoteStatus = "ProviderUnavailable"

	QuoteStatusSucceeded QuoteStatus = "Succeeded"

	QuoteStatusQuoteWithWarning QuoteStatus = "QuoteWithWarning"

	QuoteStatusInternalError QuoteStatus = "InternalError"

	QuoteStatusComplete QuoteStatus = "Complete"
)

type ReportDataFormat string

const (
	ReportDataFormatNonZipped ReportDataFormat = "NonZipped"

	ReportDataFormatZipped ReportDataFormat = "Zipped"
)

type Message struct {
	Header *Header `xml:"Header,omitempty"`
}

type Header struct {
	Space       string `xml:"xmlns,attr"`
	HungAttr    xml.Attr
	Application string `xml:"Application,omitempty"`
}

type ComparisonQuoteRequestWithProviders struct {
	*ComparisonQuoteRequest

	ProvidersRequested *ArrayOfProvider `xml:"ProvidersRequested,omitempty"`
}

type ComparisonQuoteRequest struct {
	*ComparisonQuoteRequest

	TransactionHeaders *TransactionHeaders `xml:"TransactionHeaders,omitempty"`
}

type TransactionHeaders struct {
	BusinessType *BusinessType `xml:"BusinessType,omitempty"`

	IllustrationBasis *IllustrationBasis `xml:"IllustrationBasis,omitempty"`

	QuoteReferences *QuoteReferences `xml:"QuoteReferences,omitempty"`
}

type QuoteReferences struct {
	AdviceProcessId *Guid `xml:"AdviceProcessId,omitempty"`

	ComparisonQuoteParentId *Guid `xml:"ComparisonQuoteParentId,omitempty"`

	ExternalReferenceId string `xml:"ExternalReferenceId,omitempty"`

	SessionId *Guid `xml:"SessionId,omitempty"`

	TransactionId *Guid `xml:"TransactionId,omitempty"`
}

type ArrayOfProvider struct {
	Provider []*Provider `xml:"Provider,omitempty"`
}

type Provider struct {
	Code string `xml:"Code,omitempty"`
}

type OverrideOptions struct {
	UserOverrideOptions *UserOverrideOptions `xml:"UserOverrideOptions,omitempty"`
}

type UserOverrideOptions struct {
	AgencyAddressLine1 string `xml:"AgencyAddressLine1,omitempty"`

	AgencyAddressLine2 string `xml:"AgencyAddressLine2,omitempty"`

	AgencyCode string `xml:"AgencyCode,omitempty"`

	AgencyTown string `xml:"AgencyTown,omitempty"`

	CompanyName string `xml:"CompanyName,omitempty"`

	FirmFSAOverride string `xml:"FirmFSAOverride,omitempty"`

	FirmFSARef string `xml:"FirmFSARef,omitempty"`

	IndividualFSARef string `xml:"IndividualFSARef,omitempty"`

	PanelIdentifier string `xml:"PanelIdentifier,omitempty"`

	Postcode string `xml:"Postcode,omitempty"`

	PrincipleFSARef string `xml:"PrincipleFSARef,omitempty"`

	WholeOfMarket bool `xml:"WholeOfMarket,omitempty"`
}

type SubmitComparisonQuoteResponseMessage struct {
	*Message

	SubmitComparisonQuoteResponse *SubmitComparisonQuoteResponse `xml:"SubmitComparisonQuoteResponse,omitempty"`
}

type SubmitComparisonQuoteResponse struct {
	Status *ComparisonQuoteSubmissionStatus `xml:"Status,omitempty"`

	TransactionHeaders *TransactionHeaders `xml:"TransactionHeaders,omitempty"`
}

type RetrieveComparisonQuoteResultsRequestMessage struct {
	*Message

	ComparisonQuoteResultsRequest *ComparisonQuoteResultsRequest `xml:"ComparisonQuoteResultsRequest,omitempty"`
}

type ComparisonQuoteResultsRequest struct {
	ComparisonId *Guid `xml:"ComparisonId,omitempty"`
}

type ComparisonQuoteResult struct {
	QuotesComplete int32 `xml:"QuotesComplete,omitempty"`

	QuotesGenerated int32 `xml:"QuotesGenerated,omitempty"`

	Status *ComparisonQuoteStatus `xml:"Status,omitempty"`

	TransactionHeaders *TransactionHeaders `xml:"TransactionHeaders,omitempty"`

	UnavailableProducts *ArrayOfUnavailableProduct `xml:"UnavailableProducts,omitempty"`
}

type ArrayOfUnavailableProduct struct {
	UnavailableProduct []*UnavailableProduct `xml:"UnavailableProduct,omitempty"`
}

type UnavailableProduct struct {
	Consequence *Consequence `xml:"Consequence,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	ProductId int32 `xml:"ProductId,omitempty"`

	ProviderShortName string `xml:"ProviderShortName,omitempty"`

	Reasons *ArrayOfProductUnavailableReason `xml:"Reasons,omitempty"`
}

type ArrayOfProductUnavailableReason struct {
	ProductUnavailableReason []*ProductUnavailableReason `xml:"ProductUnavailableReason,omitempty"`
}

type ProductUnavailableReason struct {
	Code *ProductUnavailableReasonCode `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type QuoteResult struct {
	*QuoteResult

	ExpiryDate time.Time `xml:"ExpiryDate,omitempty"`

	IllustrationBasis *IllustrationBasis `xml:"IllustrationBasis,omitempty"`

	ProductId int32 `xml:"ProductId,omitempty"`

	ProviderDisplayName string `xml:"ProviderDisplayName,omitempty"`

	ProviderQuoteStatus *ProviderQuoteStatus `xml:"ProviderQuoteStatus,omitempty"`

	ProviderShortName string `xml:"ProviderShortName,omitempty"`

	QuoteId *Guid `xml:"QuoteId,omitempty"`
}

type ProviderQuoteStatus struct {
	Errors *Errors `xml:"Errors,omitempty"`

	Explanations *Explanations `xml:"Explanations,omitempty"`

	Status *QuoteStatus `xml:"Status,omitempty"`

	Warnings *Warnings `xml:"Warnings,omitempty"`
}

type Errors struct {
	Error []string `xml:"Error,omitempty"`
}

type Explanations struct {
	Explanation []string `xml:"Explanation,omitempty"`
}

type Warnings struct {
	Warning []string `xml:"Warning,omitempty"`
}

type ResultCommission struct {
	Amount float64 `xml:"Amount,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type HungComparisonQuoteReportRequest struct {
	ComparisonId *Guid `xml:"ComparisonId,omitempty"`

	RemoveBranding bool `xml:"RemoveBranding,omitempty"`

	ReportFormat *ReportDataFormat `xml:"ReportFormat,omitempty"`
}

type RetrieveComparisonQuoteReportResponseMessage struct {
	*Message

	ComparisonQuoteReport *ComparisonQuoteReport `xml:"ComparisonQuoteReport,omitempty"`
}

type ComparisonQuoteReport struct {
	Report []byte `xml:"Report,omitempty"`
}

type Service struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/Schema/Faults/V1 ServiceFault"`

	MessageText string `xml:"MessageText,omitempty"`

	Data *ArrayOfstring `xml:"Data,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type Duplicate struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/Schema/Faults/V1 DuplicateFault"`

	ID *Guid `xml:"ID,omitempty"`
}

type InvalidService struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/Schema/Faults/V1 InvalidServiceFault"`

	ServiceExpected string `xml:"ServiceExpected,omitempty"`

	ServiceRequested string `xml:"ServiceRequested,omitempty"`
}

type TransactionIdentifier struct {
	XMLName xml.Name `xml:"http://assureweb.co.uk/Schema/Faults/V1 TransactionIdentifierFault"`

	TransactionID *Guid `xml:"TransactionID,omitempty"`
}

type ArrayOfstring struct {
	String []string `xml:"string,omitempty"`
}

type ValidationFault struct {
	Details *ArrayOfValidationDetail `xml:"Details,omitempty"`
}

type ArrayOfValidationDetail struct {
	ValidationDetail []*ValidationDetail `xml:"ValidationDetail,omitempty"`
}

type ValidationDetail struct {
	Key string `xml:"Key,omitempty"`

	Message string `xml:"Message,omitempty"`

	Tag string `xml:"Tag,omitempty"`
}

type Consequence string

const (
	ConsequenceNone Consequence = "None"

	ConsequenceCoarseFiltered Consequence = "CoarseFiltered"

	ConsequenceFineFiltered Consequence = "FineFiltered"
)

type IEnhancedProtectionComparisonService interface {

	// Error can be either of the following types:
	//
	//   - ServiceFault
	//   - DuplicateFault
	//   - ValidationFault

	SubmitComparisonQuoteRequest(request *SubmitComparisonQuoteRequest) (*SubmitComparisonQuoteRequestResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidServiceFault
	//   - TransactionIdentifierFault
	//   - ServiceFault

	RetrieveComparisonQuoteResults(request *RetrieveComparisonQuoteResults) (*RetrieveComparisonQuoteResultsResponse, error)

	// Error can be either of the following types:
	//
	//   - InvalidServiceFault
	//   - TransactionIdentifierFault
	//   - ServiceFault

	RetrieveComparisonQuoteResultsWithInput(request *RetrieveComparisonQuoteResultsWithInput) (*RetrieveComparisonQuoteResultsWithInputResponse, error)

	// Error can be either of the following types:
	//
	//   - ServiceFault
	//   - TransactionIdentifierFault

	RetrieveComparisonQuoteReport(request *RetrieveComparisonQuoteReport) (*RetrieveComparisonQuoteReportResponse, error)
}

type iEnhancedProtectionComparisonService struct {
	client *soap.Client
}

func NewIEnhancedProtectionComparisonService(client *soap.Client) IEnhancedProtectionComparisonService {
	return &iEnhancedProtectionComparisonService{
		client: client,
	}
}

func (service *iEnhancedProtectionComparisonService) SubmitComparisonQuoteRequest(request *SubmitComparisonQuoteRequest) (*SubmitComparisonQuoteRequestResponse, error) {
	response := new(SubmitComparisonQuoteRequestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iEnhancedProtectionComparisonService) RetrieveComparisonQuoteResults(request *RetrieveComparisonQuoteResults) (*RetrieveComparisonQuoteResultsResponse, error) {
	response := new(RetrieveComparisonQuoteResultsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iEnhancedProtectionComparisonService) RetrieveComparisonQuoteResultsWithInput(request *RetrieveComparisonQuoteResultsWithInput) (*RetrieveComparisonQuoteResultsWithInputResponse, error) {
	response := new(RetrieveComparisonQuoteResultsWithInputResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *iEnhancedProtectionComparisonService) RetrieveComparisonQuoteReport(request *RetrieveComparisonQuoteReport) (*RetrieveComparisonQuoteReportResponse, error) {
	response := new(RetrieveComparisonQuoteReportResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
