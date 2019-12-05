package algorithm

import "github.com/profzone/service-id/internal/constants/types"

type GeneratorAlgorithm interface {
	GetAlgorithmID() types.GenerateAlgorithm
	InitGenerator() error
	GenerateUniqueID() (uint64, error)
}

type GeneratorContainer map[types.GenerateAlgorithm]GeneratorAlgorithm

func (a GeneratorContainer) RegisterAlgorithm(v GeneratorAlgorithm) {
	v.InitGenerator()
	a[v.GetAlgorithmID()] = v
}

func (a GeneratorContainer) GetAlgorithm(algorithm types.GenerateAlgorithm) GeneratorAlgorithm {
	if v, ok := a[algorithm]; ok {
		return v
	}
	return nil
}

var GeneratorContainerInstance GeneratorContainer

func init() {
	GeneratorContainerInstance = GeneratorContainer{}
}
