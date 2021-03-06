/*
Copyright (c) 2016 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"github.com/stretchr/testify/assert"
	. "github.com/vmware/priam/testaid"
	"testing"
)

func TestHttpGet(t *testing.T) {
	h := func(t *testing.T, req *TstReq) *TstReply {
		assert.Empty(t, req.Input)
		return &TstReply{Output: "ok", ContentType: "text/plain"}
	}
	testpath, output, expected := "/testpath", "", "ok"
	srv := StartTstServer(t, map[string]TstHandler{"GET" + testpath: h})
	ctx := NewHttpContext(NewLogr(), srv.URL, "", "")
	err := ctx.Request("GET", testpath, nil, &output)
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}
