package module 

import (
	"net/http"
	"base"
)

type MID string

type Counts struct {
	AcceptedCount uint64
	CalledCount uint64
	CompleteCount uint64
	HandingCount uint64
}
type CaculateScore func(couts Counts) uint64
type SummaryStruct struct {
	ID MID `json:"id"`
	Called uint64 `json:"called"`
	Accepted uint64 `josn:"accepted"`
	Handing uint64 `json:"handing"`
	Extra interface{} `json:"extra,omitempty"`
}

type Module interface {
	ID() MID
	Addr() string
	Score() uint64
	SetScore(score uint64)
	ScoreCalcutar() CaculateScore
	CalledCount() uint64
	AcceptedCout() uint64
	CompleteCount() uint64
	HandingNumber() uint64
	Counts() Counts
	Summary() SummaryStruct
} 

var midTemplate = "%s%d|%s"
type Type string

const (
	TYPE_DOWNLOADER Type = "downloader"
	TYPE_ANALYZER   Type = "analyzer"
	TYPE_PIPELINE   Type ="pipeline"
)

var legalTypeLetteMap = map[Type]string{TYPE_DOWNLOADER:"D",TYPE_ANALYZER:"A",TYPE_PIPELINE:"P"}

type SNGenertor interface  {
	Start( ) uint64
	Max() uint64
	CycleCount() uint64
	Get() uint64
} 

type Registrar interface {
	Reigster(module Module) (bool,error)
	Unregister(mid MID) (bool,error)
	Get(moduleType Type)(Module,error)
	GetAllByType(moduleType Type) (map[MID] Module,error)
	GetAll() map[MID] Module
	Clear()
}

