package module
type Scheduler interface {
	Init(requestArgs RequestArgs,dataArgs DataArgs,moduleArgs ModuleArgs)(err error)
	Start(firstHTTPReq *http.Request) (err error)
	Stop() (err error)
	Status() Status
	ErrorChan() <-chan error
	Idle() bool
	Summary() SchedSummary 
}

type RequestArgs struct {
	AcceptedDomains [] string `json:"accpted_primary_domains"`
	MaxDepth uint32 `json:"max_depth"`

}

type DataArgs struct {
	ReqBufferCap uint32 `json:"req_buffer_cap"`
	ReMaxBufferNumber uint32 `json:"req_max_buffer_number"`
	RespBufferCap uint32 `json:"resp_buffer_cap`
	RespMaxBufferNumber uint32 `json:"resp_max_buffer_number"`
	ItemBufferCap uint32 `json:"item_buffer_cap"`
	ItemMaxBufferNumber uint32 `json:"item_max_buffer_number"`
	ErrorBufferCap uint32 `json:"error_buffer_cap"`
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

type ModuleArgs struct {
	Downloader []module.Downloader
	Analyzer [] module.Analyzer
	Pipelinse [] module.Pipeline
}

type Args interface {
	Check() error
}

const (
	SCHED_STATUS_UNINITIALIZED Status =0
	SCHED_STATUS_INITIALIZING Status =1
	SCHED_STATUS_INITIALIZED Status =2
	SHCED_STATUS_STARTING Status =3
	SHCED_STATUS_STARTED Status=4
	SHCED_STATUS_STOPPING Status =5
	SHCED_STATUS_STOPPED Status=6
)
