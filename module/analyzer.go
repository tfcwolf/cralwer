package module

type ParseResponse func(httpResp *http.Response,respDepth uint32) ([]Data,[] error)
type Analyzer interface {
	Module
	RespParsers() [] ParseResponse
	Analyzer(resp * Response)([] Data,[] error)
}

