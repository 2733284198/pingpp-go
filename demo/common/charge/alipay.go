package charge

var Alipay = map[string]interface{}{
	// 可选，开放平台返回的包含账户信息的 token（授权令牌，商户在一定时间内对支付宝某些服务的访问权限）。通过授权登录后获取的  alipay_open_id ，作为该参数的  value ，登录授权账户即会为支付账户，32 位字符串。
	// "extern_token":""
	// 可选，是否发起实名校验，T 代表发起实名校验；F 代表不发起实名校验。
	"rn_checK": "T",
}