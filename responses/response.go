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

	"github.com/clbanning/mxj"
)

// CommonResponse 公共回应
type CommonResponse struct {
	httpContent []byte
	json        string
}

// NewCommonResponse 创建新的请求返回
func NewCommonResponse() (response *CommonResponse) {
	return &CommonResponse{}
}

// GetHttpContentJson 获取 JSON 数据
func (req *CommonResponse) GetHttpContentJson() string {
	return req.json
}

// GetHttpContentMap 获取 MAP 数据
func (req *CommonResponse) GetHttpContentMap() (mxj.Map, error) {
	return mxj.NewMapJson([]byte(req.json))
}

// SetHttpContent 设置请求信息
func (req *CommonResponse) SetHttpContent(httpContent []byte, dataType string) {
	req.httpContent = httpContent
	switch dataType {
	case "xml":
		mv, _ := mxj.NewMapXml(req.httpContent) // unmarshal
		var str interface{}
		if _, ok := mv["xml"]; ok { //去掉 xml 外层
			str = mv["xml"]
		} else {
			str = mv
		}
		jsonStr, _ := json.Marshal(str)
		req.json = string(jsonStr)
	case "string":
		req.json = string(req.httpContent)
	}
}
