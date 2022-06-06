// Author: yann
// Date: 2022/5/20
// Desc: ginwrap

package ginwrap

type EngineWrapOption func(wrap *engineWrap)

func EngineWrapLoggerOption() EngineWrapOption {
	return func(wrap *engineWrap) {
		wrap.Use(Print)
	}
}

func PrintReqParamsOption() EngineWrapOption {
	return func(wrap *engineWrap) {
		printReq = true
	}
}

func PrintRespParamsOption() EngineWrapOption {
	return func(wrap *engineWrap) {
		printResp = true
	}
}
