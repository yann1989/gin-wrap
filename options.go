// Author: yann
// Date: 2022/5/20
// Desc: ginwrap

package ginwrap

type EngineWrapOption func()

var (
	EngineWrapLoggerOption EngineWrapOption = func() {
		isPrint = true
	}
	PrintReqParamsOption EngineWrapOption = func() {
		printReq = true
	}
	PrintRespParamsOption EngineWrapOption = func() {
		printResp = true
	}
)
