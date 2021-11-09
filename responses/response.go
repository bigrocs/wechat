/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package responses

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/bigrocs/wechat/config"
	"github.com/bigrocs/wechat/requests"
	"github.com/bigrocs/wechat/util"
	"github.com/clbanning/mxj"
)

const (
	CLOSED     = "CLOSED"     // -1 订单关闭
	USERPAYING = "USERPAYING" // 0	订单支付中
	SUCCESS    = "SUCCESS"    // 1	订单支付成功
	WAITING    = "WAITING"    // 2	系统执行中请等待
)

// CommonResponse 公共回应
type CommonResponse struct {
	Config      *config.Config
	Request     *requests.CommonRequest
	httpContent []byte
	json        string
}

type Map *mxj.Map

// NewCommonResponse 创建新的请求返回
func NewCommonResponse(config *config.Config, request *requests.CommonRequest) (response *CommonResponse) {
	c := &CommonResponse{}
	c.Config = config
	c.Request = request
	return c
}

// GetHttpContentJson 获取 JSON 数据
func (res *CommonResponse) GetHttpContentJson() string {
	return res.json
}

// GetHttpContentMap 获取 MAP 数据
func (res *CommonResponse) GetHttpContentMap() (mxj.Map, error) {
	return mxj.NewMapJson([]byte(res.json))
}

// GetVerifySignDataMap 获取 GetVerifySignDataMap 校验后数据数据
func (res *CommonResponse) GetVerifySignDataMap() (m mxj.Map, err error) {
	r, err := res.GetHttpContentMap()
	if err != nil {
		return r, err
	}
	if r["sign"] != nil {
		if util.VerifySign(r, r["sign"].(string), res.Config.ApiKey, res.Config.SignType) {
			return res.GetSignDataMap()
		} else {
			return r, errors.New("sign verification failed")
		}
	} else {
		return r, errors.New("sign is not")
	}
}

// SetHttpContent 设置请求信息
func (res *CommonResponse) SetHttpContent(httpContent []byte, dataType string) {
	res.httpContent = httpContent
	switch dataType {
	case "xml":
		mv, _ := mxj.NewMapXml(res.httpContent) // unmarshal
		var str interface{}
		if _, ok := mv["xml"]; ok { //去掉 xml 外层
			str = mv["xml"]
		} else {
			str = mv
		}
		jsonStr, _ := json.Marshal(str)
		res.json = string(jsonStr)
	case "string":
		res.json = string(res.httpContent)
	}
}

// data{
// 	channel			//	通道内容		alipay、wechat、icbc
// 	content			//	第三方返回内容 	{}
// 	return_code		//	返回代码 		SUCCESS
// 	return_msg		//	返回消息		支付失败
// 	status			//	下单状态 		【SUCCESS成功、CLOSED关闭、USERPAYING等待用户付款、WAITING系统繁忙稍后查询】
// 	total_fee		//  订单金额		88
// 	refund_fee 		//  退款金额		10
// 	trade_no 		// 	渠道交易编号 	2013112011001004330000121536
// 	out_trade_no	// 	商户订单号		T1024501231476
//  out_refund_no	//  商户退款单号	T1024501231476_T
// 	wechat_open_id		//  微信openid		[oUpF8uN95-Ptaags6E_roPHg7AG
//  wechat_is_subscribe 	//  微信是否微信关注公众号
// 	alipay_logon_id  //	支付宝账号		158****1562
//  alipay_user_id  //	买家在支付宝的用户id	2088101117955611
// 	time_end		//  支付完成时间	20141030133525
// }

// GetSignDataMap 获取 MAP 数据
func (res *CommonResponse) GetSignDataMap() (mxj.Map, error) {
	data := mxj.New()
	content, err := res.GetHttpContentMap()
	if err != nil {
		return nil, err
	}
	if res.Request.ApiName == "pay.micropay" {
		data = res.handerWechatTradePay(content)
	}
	if res.Request.ApiName == "pay.orderquery" {
		data = res.handerWechatTradeQuery(content)
	}
	if res.Request.ApiName == "pay.refund" {
		data = res.handerWechatTradeRefund(content)
	}
	if res.Request.ApiName == "pay.refundquery" {
		data = res.handerWechatTradeRefundQuery(content)
	}
	if res.Request.ApiName == "tools.authcodetoopenid" {
		data = res.handerWechatQueryOpenId(content)
	}

	data["channel"] = "wechat" //渠道
	data["content"] = content
	return data, err
}

// handerWechatTradePay
func (res *CommonResponse) handerWechatTradePay(content mxj.Map) mxj.Map {
	data := mxj.New()
	data["return_msg"] = content["return_msg"]
	if content["return_code"] == "SUCCESS" {
		if content["result_code"] == "SUCCESS" {
			data["return_code"] = SUCCESS
			data["status"] = SUCCESS
			total_fee, _ := strconv.ParseInt(content["total_fee"].(string), 10, 64)
			data["total_fee"] = total_fee
			if v, ok := content["cash_fee"]; ok { // 用户实际扣减金额
				i, _ := strconv.ParseInt(v.(string), 10, 64)
				data["buyer_pay_fee"] = i
			} else {
				data["buyer_pay_fee"] = total_fee
			}
			data["trade_no"] = content["transaction_id"]
			data["out_trade_no"] = content["out_trade_no"]
			data["wechat_is_subscribe"] = content["is_subscribe"]
			data["wechat_open_id"] = content["openid"]
			data["time_end"] = content["time_end"]
		} else {
			data["return_code"] = "FAIL"
			data["return_msg"] = content["err_code_des"]
			if content["err_code"] == "USERPAYING" {
				data["status"] = USERPAYING
			}
		}

	} else {
		data["return_code"] = "FAIL"
	}
	return data
}

// handerWechatTradeQuery
func (res *CommonResponse) handerWechatTradeQuery(content mxj.Map) mxj.Map {
	// 	SUCCESS--支付成功
	// REFUND--转入退款
	// NOTPAY--未支付
	// CLOSED--已关闭
	// REVOKED--已撤销(刷卡支付)
	// USERPAYING--用户支付中
	// PAYERROR--支付失败(其他原因，如银行返回失败)
	// ACCEPT--已接收，等待扣款
	// 支付状态机请见下单API页面
	data := mxj.New()
	data["return_msg"] = content["return_msg"]
	if content["return_code"] == "SUCCESS" {
		if content["result_code"] == "SUCCESS" {
			data["return_code"] = SUCCESS
			switch content["trade_state"] {
			case "SUCCESS":
				data["status"] = SUCCESS
			case "REFUND":
				data["status"] = SUCCESS
			case "NOTPAY":
				data["status"] = USERPAYING
			case "CLOSED":
				data["status"] = CLOSED
			case "REVOKED":
				data["status"] = CLOSED
			case "USERPAYING":
				data["status"] = USERPAYING
			case "PAYERROR":
				data["status"] = CLOSED
			case "ACCEPT":
				data["status"] = WAITING
			}
			total_fee := int64(0)
			if v, ok := content["total_fee"]; ok { // 用户实际扣减金额
				total_fee, _ = strconv.ParseInt(v.(string), 10, 64)
				data["total_fee"] = total_fee
			}
			if v, ok := content["cash_fee"]; ok { // 用户实际扣减金额
				i, _ := strconv.ParseInt(v.(string), 10, 64)
				data["buyer_pay_fee"] = i
			} else {
				data["buyer_pay_fee"] = total_fee
			}
			data["trade_no"] = content["transaction_id"]
			data["out_trade_no"] = content["out_trade_no"]
			data["wechat_is_subscribe"] = content["is_subscribe"]
			data["wechat_open_id"] = content["openid"]
			data["wechat_sub_open_id"] = content["sub_openid"]
			data["time_end"] = content["time_end"]
		} else {
			data["return_code"] = "FAIL"
			data["return_msg"] = content["err_code_des"]
			if content["err_code"] == "ORDERNOTEXIST" {
				data["return_code"] = SUCCESS
				data["status"] = CLOSED
			}
		}

	} else {
		data["return_code"] = "FAIL"
	}
	return data
}

// handerWechatTradeRefund
func (res *CommonResponse) handerWechatTradeRefund(content mxj.Map) mxj.Map {
	data := mxj.New()
	data["return_msg"] = content["return_msg"]
	if content["return_code"] == "SUCCESS" {
		if content["result_code"] == "SUCCESS" {
			data["return_code"] = SUCCESS
			data["status"] = USERPAYING
			total_fee, _ := strconv.ParseInt(content["total_fee"].(string), 10, 64)
			data["total_fee"] = total_fee
			refund_fee, _ := strconv.ParseInt(content["refund_fee"].(string), 10, 64)
			data["refund_fee"] = refund_fee
			data["trade_no"] = content["transaction_id"]
			data["out_trade_no"] = content["out_trade_no"]
			data["out_refund_no"] = content["out_refund_no"]
		} else {
			data["return_code"] = "FAIL"
			data["return_msg"] = content["err_code_des"]
			data["status"] = USERPAYING
		}
	} else {
		data["return_code"] = "FAIL"
	}
	return data
}

// handerWechatTradeRefundQuery
func (res *CommonResponse) handerWechatTradeRefundQuery(content mxj.Map) mxj.Map {
	data := mxj.New()
	data["return_msg"] = content["return_msg"]
	if content["return_code"] == "SUCCESS" {
		if content["result_code"] == "SUCCESS" {
			data["return_code"] = SUCCESS
			switch content["refund_status_0"] {
			case "SUCCESS":
				data["status"] = SUCCESS
			case "PROCESSING":
				data["status"] = USERPAYING
			case "REFUNDCLOSE":
				data["status"] = CLOSED
			case "CHANGE":
				data["status"] = CLOSED
			}
			total_fee, _ := strconv.ParseInt(content["total_fee"].(string), 10, 64)
			data["total_fee"] = total_fee
			refund_fee, _ := strconv.ParseInt(content["refund_fee"].(string), 10, 64)
			data["refund_fee"] = refund_fee
			data["trade_no"] = content["transaction_id"]
			data["out_trade_no"] = content["out_trade_no"]
			data["out_refund_no"] = content["out_refund_no"]
		} else {
			data["return_code"] = "FAIL"
			data["return_msg"] = content["err_code_des"]
			if content["err_code"] == "REFUNDNOTEXIST" {
				data["return_code"] = SUCCESS
				data["status"] = CLOSED
			}
		}
	} else {
		data["return_code"] = "FAIL"
	}
	return data
}

// handerWechatQueryOpenId
func (res *CommonResponse) handerWechatQueryOpenId(content mxj.Map) mxj.Map {
	data := mxj.New()
	data["return_msg"] = content["return_msg"]
	if content["return_code"] == "SUCCESS" {
		if content["result_code"] == "SUCCESS" {
			data["return_code"] = SUCCESS
			data["wechat_open_id"] = content["openid"]
			data["wechat_sub_open_id"] = content["sub_openid"]
		} else {
			data["return_code"] = "FAIL"
			data["return_msg"] = content["err_code_des"]
		}
	} else {
		data["return_code"] = "FAIL"
	}
	return data
}
