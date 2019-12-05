package types

//go:generate eden generate enum --type-name=GenerateAlgorithm
// api:enum
type GenerateAlgorithm uint8

// ID生成算法
const (
	GENERATE_ALGORITHM_UNKNOWN    GenerateAlgorithm = iota
	GENERATE_ALGORITHM__SNOWFLAKE                   // Snowflake
)
