package charge

var Cmb_wallet = map[string]interface{}{
	// 必须，交易完成跳转的地址。
	"result_url": "http://example.com/success",
	/**
	 * 对于 p_no, seq , m_uid , mobile 这几个参数：
	 * 1. 这几个参数是用户自定义的。
	 * 2. 对于同一个终端用户每次请求 charge 务必使用同一套参数（确保每个参数都不变），
	 * 任意参数变更都会导致用户重新签约，同一个用户和招行重新签约的次数有限制，超限制就会无法签约 ，导致用户无法使用。
	 */
	// 必须，客户协议号，不超过 30 位的纯数字字符串。
	"p_no": "your p_no",
	// 必须，协议开通请求流水号，不超过 20 位的纯数字字符串，请保证系统内唯一。
	"seq": "your seq",
	// 必须，协议用户 ID，不超过 20 位的纯数字字符串。
	"m_uid": "your m_uid",
	// 必须，协议手机号，11 位数字。
	"mobile": "your mobile",
}