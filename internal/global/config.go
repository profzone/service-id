package global

import "github.com/profzone/service-id/internal/constants/types"

type SnowflakeConfig struct {
	Epoch    int64
	NodeID   int64
	NodeBits uint8
	StepBits uint8
}

var Config = struct {
	GenerateAlgorithm types.GenerateAlgorithm
	SnowflakeConfig   SnowflakeConfig
}{}
