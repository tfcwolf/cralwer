package base
import (
	"net/http"
)
type Request struct {
	rq *http.Request
	depth unit32
}

func NewRequest(httReq *http.Request, depth unit32) *Request {
	return &Request{httpReq:httpReq,depth:depth}
}

func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

func (req *Request) Depth() unit32 {
	return req.*depth
}

type Response struct {
	httpResp *http.Response
	depth unit32
}

func NewResponse(httpResp *http.Response,depth unit32) *Response {
	return &Response{httpResp:httpResp,depth:depth}
}

func (rep *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

func (resp *Response) Depth() unit32 {
	return resp.depth
}

type Item map[string]interface{}
type Data interface {
	Valid() bool
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) Valid() bool {
	return item !=nil
}

type CrawleError interface {
	Type() ErrorType
	Error() string
}

type ErrorType string;

type myCrawlerError struct {
	errType ErrorType
	errMsg string
	fullErrMsg string
}

const (
	ERROR_TYPE_DOWNLOADER ErrorType ="downloader error"
	ERROR_TYPE_ANALYZER   ErrorType = "analyzer error"
	ERROR_TYPE_SCHEDULR   ErrorType = "scheduler error"
	ERROR_TYPE_PIPELINE   ErrorType = "pipeline error"
)

func NewCrawlerError(errType ErrorType,errMsg string) CrawlerError {
	return &myCrawlerError{
		errType:errType,
		errMsg:strings.TrimSpace(errMsg),
	}
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType 
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg ==""{
		ce.genFullerrMsg()
	}
	return ce.fullErrMsg
	
}

func (ce *myCrawlerError) genFullerrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error:")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(ce.errMsg)
	ce.genFullerrMsg = fmt.Sprintf("%s",buffer.String())
	return
}