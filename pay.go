package main

import (
	"encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	"github.com/pingplusplus/pingpp-go/pingpp/event"
	"github.com/pingplusplus/pingpp-go/pingpp/redEnvelope"
	"github.com/pingplusplus/pingpp-go/pingpp/refund"
	"log"
)

func init() {
	pingpp.Key = "YOUR-KEY"
	fmt.Println(pingpp.Version())
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleCharge_new() {
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	params := &pingpp.ChargeParams{
		Order_no:  "sdaffsadfdfsweqrad34xg13",
		App:       pingpp.App{Id: "YOUR-APP-ID"},
		Amount:    100,
		Channel:   "wx_pub",
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   "Your Subject",
		Body:      "Your Body",
		Extra:     pingpp.Extra{Open_id: "your openid"},
		Metadata:  metadata,
	}
	a, _ := json.Marshal(params)
	fmt.Println(string(a))
	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	ch, err := charge.New(params)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(ch)
	chs, err := json.Marshal(ch)
	if err != nil {
		fmt.Println(err)
	} else {
		//打印 charge json
		fmt.Println(string(chs))
	}
}

func ExampleCharge_get() {
	ch, err := charge.Get("ch_id")
	if err != nil {
		log.Fatal(err)
	}
	chstring, _ := json.Marshal(ch)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(chstring))
}

func ExampleCharge_list() {

	params := &pingpp.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := charge.List(params)
	for i.Next() {
		c := i.Charge()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	// fmt.Println(i)
}

func ExampleRefund_new() {
	params := &pingpp.RefundParams{
		Amount:      1,
		Description: "12345",
	}
	re, err := refund.New("ch_id", params)

	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(re)
	log.Printf("%v\n", string(restring))
}

func ExampleRefund_get() {
	re, err := refund.Get("ch_id", "re_yb5KiTHqnrT8GqrTOSuv1u5S")
	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(re)

	log.Printf("%v\n", string(restring))
}

func ExampleRefund_list() {
	params := &pingpp.RefundListParams{}
	// params.Filters.AddFilter("limit", "", "1")
	// //设置是不是只需要之前设置的 limit 这一个查询参数
	// params.Single = true
	i := refund.List("ch_id", params)
	for i.Next() {
		c := i.Refund()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	fmt.Println(i)
}

func ExampleRedEnvelope_new() {
	redenvelopeParams := &pingpp.RedEnvelopeParams{
		App:         pingpp.App{Id: "YOUR-APP-ID"},
		Channel:     "wx_pub",
		Order_no:    "987654345",
		Amount:      100,
		Currency:    "cny",
		Recipient:   "your openid",
		Subject:     "Your Subject",
		Body:        "Your Body",
		Description: "Your Description",
		Extra:       pingpp.RedEnvelopeExtra{Nick_name: "nick name", Send_name: "nick name"},
	}
	redEnvelope, err := redEnvelope.New(redenvelopeParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(redEnvelope)
}

func ExampleRedEnvelope_get() {
	redEnvelope, err := redEnvelope.Get("red_id")
	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(redEnvelope)
	log.Printf("%v\n", string(restring))
}

func ExampleRedEnvelope_list() {
	params := &pingpp.RedEnvelopeListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := redEnvelope.List(params)
	for i.Next() {
		c := i.RedEnvelope()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
}

func ExampleEvent_get() {
	eve, err := event.Get("evt_id")
	if err != nil {
		log.Fatal(err)
	}
	// eve, _ := json.Marshal(event)
	// log.Printf("%v\n", string(eve))
	fmt.Println(eve)
}

func ExampleEvent_list() {
	params := &pingpp.EventListParams{}
	params.Filters.AddFilter("type", "", "event_type")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := event.List(params)
	for i.Next() {
		c := i.Event()
		fmt.Println(c)
	}
}

func main() {
	ExampleCharge_new()
}
