/*
 * Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package swagger

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	xmlBody, err := jsonToXML(r)
	if err != nil {
		log.Print("Error converting JSON request payload to XML", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &RequestHandlerResponseBody{
		HeadersToReplace: &map[string]string{
			"content-type": "application/xml",
		},
		Body: xmlBody,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Print("Error parsing to RequestHandlerResponseBody", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(respBytes); err != nil {
		log.Print("Error writing response body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func jsonToXML(r *http.Request) (string, error) {
	reqBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var reqBody RequestHandlerRequestBody
	if err = json.Unmarshal(reqBodyBytes, &reqBody); err != nil {
		return "", err
	}
	decodedBody, err := base64.StdEncoding.DecodeString(reqBody.RequestBody)
	if err != nil {
		return "", err
	}

	var xmlBody Book
	if err := json.Unmarshal(decodedBody, &xmlBody); err != nil {
		return "", err
	}
	xmlStr := fmt.Sprintf("<name>%s</name>", xmlBody.Name)
	return base64.StdEncoding.EncodeToString([]byte(xmlStr)), nil
}
