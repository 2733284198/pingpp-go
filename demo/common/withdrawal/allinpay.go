package withdrawal

var Allinpay = map[string]interface{}{
	"account": "6214850288888888",    //必填项 收款人银行卡号或者存折号。
    "name": "张三",                    //必填项 收款人姓名。
    "open_bank_code": "0308",         //必填项 开户银行编号。
    "business_code": "09900",         //选填项 业务代码，根据通联业务人员提供，不填使用通联提供默认值 09900。
    "card_type": 0,                   //选填项 银行卡号类型，0：银行卡；1：存折。
}