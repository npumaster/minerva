package config

type Experiments struct {
	FilestoreEnabled     bool
	UrlstoreEnabled      bool
	ShardingEnabled      bool `json:",omitempty"` // deprecated by autosharding
	GraphsyncEnabled     bool
	Libp2pStreamMounting bool
	P2pHttpProxy         bool //nolint
	StrategicProviding   bool
	AcceleratedDHTClient bool
}
