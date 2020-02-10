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
	dataType    string
}

// NewCommonResponse 创建新的请求返回
func NewCommonResponse() (response *CommonResponse) {
	return &CommonResponse{}
}

func (req *CommonResponse) GetHttpContent() string {
	switch req.dataType {
	case "xml":
		mv, _ := mxj.NewMapXml(req.httpContent) // unmarshal
		var str interface{}
		if _, ok := mv["xml"]; ok { //去掉 xml 外层
			str = mv["xml"]
		} else {
			str = mv
		}
		jsonStr, _ := json.Marshal(str)
		return string(jsonStr)
	case "string":
		return string(req.httpContent)
	}
	return ""
}

func (req *CommonResponse) SetHttpContent(httpContent []byte, dataType string) {
	req.httpContent = httpContent
	req.dataType = dataType
}
