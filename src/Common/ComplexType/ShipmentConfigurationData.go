package ComplexType

type ShipmentConfigurationData struct {

	// 指定运输中危险品包装的常见数据。当装运包含相同的危险货物商品的包装时，就会填充此值。
	DangerousGoodsPackageConfigurations []*DangerousGoodsDetail `xml:"DangerousGoodsPackageConfigurations,omitempty"`
}
