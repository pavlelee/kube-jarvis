/*
* Tencent is pleased to support the open source community by making TKEStack
* available.
*
* Copyright (C) 2012-2019 Tencent. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the “License”); you may not use
* this file except in compliance with the License. You may obtain a copy of the
* License at
*
* https://opensource.org/licenses/Apache-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an “AS IS” BASIS, WITHOUT
* WARRANTIES OF ANY KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations under the License.
 */
package example

import (
	"context"
	"fmt"
	"testing"
	"tkestack.io/kube-jarvis/pkg/plugins"
	"tkestack.io/kube-jarvis/pkg/plugins/diagnose"
	"tkestack.io/kube-jarvis/pkg/translate"
)

func TestDiagnostic_StartDiagnose(t *testing.T) {
	cases := []struct {
		message string
	}{
		{
			message: "mes1",
		},
		{
			message: "mes2",
		},
	}

	for _, cs := range cases {
		t.Run(fmt.Sprintf("%+v", cs), func(t *testing.T) {
			s := NewDiagnostic(&diagnose.MetaData{
				MetaData: plugins.MetaData{
					Translator: translate.NewFake(),
				},
			}).(*Diagnostic)

			s.Message = cs.message
			result, err := s.StartDiagnose(context.Background(), diagnose.StartDiagnoseParam{})
			if err != nil {
				t.Fatalf(err.Error())
			}

			for {
				res, ok := <-result
				if !ok {
					break
				}
				t.Logf("%+v", res)
			}
		})
	}
}
