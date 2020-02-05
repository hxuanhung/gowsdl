package ipipeline

import (
	"encoding/xml"
	"flag"
	"os"
	"testing"

	"github.com/hooklift/gowsdl/soap"
)

const (
	wsURL = "https://reapext.assureweb.co.uk/Ape2/EnhancedProtectionComparisonService.svc/UNP"
)

var (
	runIntegration = flag.Bool("integration", false, "test t7t integration (T7T_LICENSE_KEY env var required)")
	licenseKey     = os.Getenv("T7T_LICENSE_KEY")
)

func PrettyXML(payload interface{}) string {
	b, err := xml.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "Failed to generate json"
	}
	return string(b)
}

// TestIntegration issues requests to live twenty7tec API.
// E.g., T7T_LICENSE_KEY=**REDACTED** go test -v -run TestIntegration -integration
func TestIntegration(t *testing.T) {
	// t.Parallel()
	// if !*runIntegration {
	// 	t.Skip()
	// }

	level := IndexationLevel
	livesAssuredBasisFirstLife := LivesAssuredBasisFirstLife
	premiumFrequencyMonthly := PremiumFrequencyMonthly
	requestPremiumTypeGuaranteed := RequestPremiumTypeGuaranteed
	totalPermanentDisabilityTypeNone := TotalPermanentDisabilityTypeNone
	waiverOfPremiumNone := WaiverOfPremiumNone
	livesAssuredBasisWithNone := LivesAssuredBasisWithNoneNone
	requestBaseBenefit := &RequestBaseBenefit{}
	benefit := &RequestTermBenefit{
		RequestBaseBenefit:     requestBaseBenefit,
		BenefitAmount:          123000,
		CriticalIllness:        false,
		CriticalIllnessOptions: nil,
		ExcludeLowStart:        true,
		FractureCoverBasis:     nil,
		Indexation:             &level,
		IntegratedIncomeCoverOptions: &IntegratedIncomeCoverOptions{
			FirstLife:         nil,
			LivesAssuredBasis: &livesAssuredBasisWithNone,
		},
		Life:                         true,
		LifeCoverBuyBack:             true,
		LifeOrEarlierCriticalIllness: false,
		LivesAssuredBasis:            &livesAssuredBasisFirstLife,
		PremiumAmount:                0,
		PremiumFrequency:             &premiumFrequencyMonthly,
		PremiumType:                  &requestPremiumTypeGuaranteed,
		Renewable:                    false,
		TermYears:                    11,
		ToAge:                        25,
		TotalPermanentDisability:     &totalPermanentDisabilityTypeNone,
		WaiverOfPremium:              &waiverOfPremiumNone,
	}
	t.Log(benefit)

	// can replace this type to string in xml
	var guid Guid = "b2034a0b-2d0d-4846-b9be-0ee005425890"

	req := RetrieveComparisonQuoteResults{
		Request: &RetrieveComparisonQuoteResultsRequestMessage{
			Space: "http://assureweb.co.uk/Schema/Shared/V2",
			Message: &Message{
				Header: &Header{
					Application: "Acre",
				},
			},
			ComparisonQuoteResultsRequest: &ComparisonQuoteResultsRequest{
				ComparisonId: &guid,
			},
		},
	}
	auth := soap.NewWSSSecurityHeader("web67752", "star477Hobbs", "", "1")
	soapClient := soap.NewClient(wsURL, soap.WithSOAPVersion(soap.SOAP12))

	soapClient.AddHeader(auth)
	type Action struct {
		XMLName        xml.Name `xml:"http://www.w3.org/2005/08/addressing Action"`
		MustUnderstand string   `xml:"mustUnderstand,attr"`
		Data           string   `xml:",chardata"`
	}
	soapClient.AddHeader(Action{
		Data: "RetrieveComparisonQuoteResults", MustUnderstand: "1"})

	client := NewIEnhancedProtectionComparisonService(soapClient)
	res, err := client.RetrieveComparisonQuoteResults(&req)
	t.Log(PrettyXML(res))
	t.Log(err)
}
