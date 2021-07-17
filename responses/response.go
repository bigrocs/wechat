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

	"github.com/bigrocs/wechat/config"
	"github.com/bigrocs/wechat/requests"
	"github.com/bigrocs/wechat/util"
	"github.com/clbanning/mxj"
)

const (
	CLOSED     = "CLOSED"     // -1 订单关闭
	USERPAYING = "USERPAYING" // 0	订单支付中
	SUCCESS    = "SUCCESS"    // 1	订单支付成功
	REFUND     = "REFUND"     // 2 退款
	WAIT       = "WAIT"       // 3	系统执行中请等待
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

// GetSignDataMap 获取 MAP 数据
func (res *CommonResponse) GetSignDataMap() (mxj.Map, error) {
	data := mxj.New()
	content, err := res.GetHttpContentMap()
	data["content"] = content
	// 下单
	// 查询 trade_state
	// SUCCESS--支付成功
	// REFUND--转入退款
	// NOTPAY--未支付
	// CLOSED--已关闭
	// REVOKED--已撤销(刷卡支付)
	// USERPAYING--用户支付中
	// PAYERROR--支付失败(其他原因，如银行返回失败)
	// ACCEPT--已接收，等待扣款
	// 支付状态机请见下单API页面
	data["return_msg"] = content["err_code_des"]
	if content["return_code"] == "SUCCESS" {
		if content["result_code"] == "SUCCESS" {
			data["return_code"] = SUCCESS
			switch content["trade_state"] {
			case "SUCCESS":
				data["stauts"] = SUCCESS
			case "REFUND":
				data["stauts"] = REFUND
			case "NOTPAY":
				data["stauts"] = USERPAYING
			case "CLOSED":
				data["stauts"] = CLOSED
			case "REVOKED":
				data["stauts"] = CLOSED
			case "PAYERROR":
				data["stauts"] = CLOSED
			case "ACCEPT":
				data["stauts"] = WAIT
			}
		} else {
			data["return_code"] = "FAIL"
			if content["err_code"] == "ORDERNOTEXIST" { // 订单不存在关闭
				data["stauts"] = CLOSED
			}
		}
	} else {
		data["return_code"] = "FAIL"
	}
	return data, err
}

// GetVerifySignDataMap 获取 GetVerifySignDataMap 校验后数据数据
func (res *CommonResponse) GetVerifySignDataMap() (m mxj.Map, err error) {
	r, err := res.GetHttpContentMap()
	if err != nil {
		return m, err
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
