package config

import "time"

const DefaultReproviderInterval = time.Hour * 22
const DefaultReproviderStrategy = "all"

type Reprovider struct {
	Interval *OptionalDuration `json:",omitempty"` // Time period to reprovide locally stored objects to the network
	Strategy *OptionalString   `json:",omitempty"` // Which keys to announce
}
