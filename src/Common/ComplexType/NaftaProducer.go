package ComplexType

type NaftaProducer struct {
	Id string `xml:"Id,omitempty"`

	Producer *Party `xml:"Producer,omitempty"`
}
