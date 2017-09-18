package alipay

type ExtendParam struct {
	/**
	* 使用花呗分期要进行的分期数
	 */
	HbFqNum string
	/**
	* 使用花呗分期需要卖家承担的手续费比例的百分值，传入100代表100%
	 */
	HbFqSellerPercent string
	/**
	* 系统商编号
	*	该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	 */
	SysServiceProviderId string
}
type InvoiceKeyInfo struct {
	/**
	 * 开票商户名称：商户品牌简称|商户门店简称
	 */

	InvoiceMerchantName string

	/**
	 * 该交易是否支持开票
	 */

	IsSupportInvoice bool

	/**
	 * 税号
	 */

	TaxNum string
}
type InvoiceInfo struct {
	/**
	 * 开票内容
	注：json数组格式
	*/
	Details string

	/**
	 * 开票关键信息
	 */
	KeyInfo InvoiceKeyInfo
}

type RoyaltyDetailInfos struct {

	/**
	 * 分账的金额，单位为元
	 */

	Amount int64

	/**
	 * 分账的比例，值为20代表按20%的比例分账
	 */

	AmountPercentage string

	/**
		 * 分账批次号
	分账批次号。
	目前需要和转入账号类型为bankIndex配合使用。
	*/

	BatchNo string

	/**
	 * 分账描述信息
	 */

	Desc string

	/**
		 * 商户分账的外部关联号，用于关联到每一笔分账信息，商户需保证其唯一性。
	如果为空，该值则默认为“商户网站唯一订单号+分账序列号”
	*/

	OutRelationId string

	/**
	 * 分账序列号，表示分账执行的顺序，必须为正整数
	 */

	SerialNo int64

	/**
		 * 如果转入账号类型为userId，本参数为接受分账金额的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	&#61548	如果转入账号类型为bankIndex，本参数为28位的银行编号（商户和支付宝签约时确定）。
	如果转入账号类型为storeId，本参数为商户的门店ID。
	*/

	TransIn string

	/**
		 * 接受分账金额的账户类型：
	&#61548	userId：支付宝账号对应的支付宝唯一用户号。
	&#61548	bankIndex：分账到银行账户的银行编号。目前暂时只支持分账到一个银行编号。
	storeId：分账到门店对应的银行卡编号。
	默认值为userId。
	*/

	TransInType string

	/**
	 * 如果转出账号类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	 */

	TransOut string

	/**
		 * 要分账的账户类型。
	目前只支持userId：支付宝账号对应的支付宝唯一用户号。
	默认值为userId。
	*/

	TransOutType string
}

type RoyaltyInfo struct {
	/**
	 * 分账明细的信息，可以描述多条分账指令，json数组。
	 */

	RoyaltyDetailInfos map[string]RoyaltyDetailInfos

	/**
		 * 分账类型
	* 卖家的分账类型，目前只支持传入ROYALTY（普通分账类型）。
	*/

	RoyaltyType string
}
type SubMerchant struct {
	/**
	 * 二级商户的支付宝id
	 */
	MerchantId string
}
type PayModel struct {

	/**
	* 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	 */
	AuthToken string

	/**
	* Iphone6 16G
	 */
	Body string

	/**
	* 禁用渠道，用户不可用指定渠道支付
	* 当有多个渠道时用“,”分隔
	* 注，与enable_pay_channels互斥
	 */
	DisablePayChannels string

	/**
	* 可用渠道，用户只能在指定渠道范围内支付
	* 当有多个渠道时用“,”分隔
	* 注，与disable_pay_channels互斥
	 */
	EnablePayChannels string

	/**
	* 业务扩展参数
	 */
	ExtendParams ExtendParam

	/**
	* 商品主类型 :0-虚拟类商品,1-实物类商品
	 */
	GoodsType string

	/**
	* 开票信息
	 */
	InvoiceInfo InvoiceInfo

	/**
	 * 商户网站唯一订单号
	 */
	OutTradeNo string

	/**
	 * 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	 */
	PassbackParams string

	/**
	 * 销售产品码，商家和支付宝签约的产品码
	 */
	ProductCode string

	/**
	* 优惠参数
	* 注：仅与支付宝协商后可用
	 */
	PromoParams string

	/**
	 * 用户付款中途退出返回商户网站的地址
	 */
	QuitUrl string

	/**
	 * 描述分账信息，Json格式，详见分账参数说明
	 */
	RoyaltyInfo RoyaltyInfo

	/**
	 * 收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	 */
	SellerId string

	/**
	* 指定渠道，目前仅支持传入pcredit
	* 若由于用户原因渠道不可用，用户可选择是否用其他渠道支付。
	* 注：该参数不可与花呗分期参数同时传入
	 */
	SpecifiedChannel string

	/**
	 * 商户门店编号
	 */
	StoreId string

	/**
	 * 间连受理商户信息体，当前只对特殊银行机构特定场景下使用此字段
	 */
	SubMerchant SubMerchant

	/**
	 * 商品的标题/交易标题/订单标题/订单关键字等。
	 */
	Subject string

	/**
	 * 绝对超时时间，格式为yyyy-MM-dd HH:mm。
	 */
	TimeExpire string

	/**
	 * 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	 */
	TimeoutExpress string

	/**
	 * 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	 */
	TotalAmount string
}
