package main

type FDBWriteable interface {
	Subspace() []byte
}

type SignalEventValidator interface {
}

// Signal interface is encoded and decoded from JSON
//
type Signal interface {
	Id() UUID
	Name() string
	CollectTime() bool
	CollectGeo() bool
	Config() map[string]string
}

type SignalBaseType int
const (
	Bool SignalBaseType = iota
	Int SignalBaseType = iota
	String SignalBaseType = iota
)

type SignalValue interface {
	BaseType() SignalBaseType
}

type SignalEvent interface {
	SignalId UUID
	Value string
}

type Channel struct {
	Id UUID
	Name string
}

type Signal struct {
	Id UUID
	Name string
	ChannelId UUID
	CollectTime bool
	CollectGeo bool
	CollectData []string
}
