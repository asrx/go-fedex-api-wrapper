package ComplexType

type VariableHandlingCharges struct {
	VariableHandlingCharge *Money `xml:"VariableHandlingCharge,omitempty"`

	FixedVariableHandlingCharge *Money `xml:"FixedVariableHandlingCharge,omitempty"`

	PercentVariableHandlingCharge *Money `xml:"PercentVariableHandlingCharge,omitempty"`

	TotalCustomerCharge *Money `xml:"TotalCustomerCharge,omitempty"`
}
