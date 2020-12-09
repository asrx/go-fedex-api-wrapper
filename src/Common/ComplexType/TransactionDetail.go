package ComplexType

type TransactionDetail struct {

	// 在回复中回显的自由格式文本。用于匹配请求和响应。
	CustomerTransactionId string `xml:"CustomerTransactionId,omitempty"`

	// Governs data payload language/translations (contrasted with ClientDetail.localization, which governs Notification.localizedMessage language selection).
	// 管理数据有效负载语言/翻译(与ClientDetail相比)。管理通知的本地化。localizedMessage语言选择)。
	Localization *Localization `xml:"Localization,omitempty"`
}
