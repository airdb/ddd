//Package infra generated by 'freedom new-project airdb.io/airdb/ddd'
package infra

import (
	"strconv"

	"encoding/json"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

// JSONResponse .
type JSONResponse struct {
	Code        int
	Error       error
	contentType string
	content     []byte
	Object      interface{}
}

// Dispatch This is the middleware for HTTP output.
func (jrep JSONResponse) Dispatch(ctx context.Context) {
	jrep.contentType = "application/json"
	var repData struct {
		Code  int         `json:"code"`
		Error string      `json:"error"`
		Data  interface{} `json:"data,omitempty"`
	}

	repData.Data = jrep.Object
	repData.Code = jrep.Code
	if jrep.Error != nil {
		repData.Error = jrep.Error.Error()
	}
	if repData.Error != "" && repData.Code == 0 {
		repData.Code = 400
	}
	ctx.Values().Set("code", strconv.Itoa(repData.Code))

	jrep.content, _ = json.Marshal(repData)
	ctx.Values().Set("response", string(jrep.content))
	hero.DispatchCommon(ctx, 0, jrep.contentType, jrep.content, nil, nil, true)
}
