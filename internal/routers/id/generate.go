package id

import (
	"errors"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/profzone/service-id/internal/constants/models"
	"github.com/profzone/service-id/internal/global"
	"github.com/profzone/service-id/internal/modules/algorithm"
)

func (p *ID) Generate(req *interface{}) (resp *models.IDGenerateResponse, status *erpc.Status) {
	generator := algorithm.GeneratorContainerInstance.GetAlgorithm(global.Config.GenerateAlgorithm)
	if generator == nil {
		status = erpc.NewStatus(1, "generator algorithm is invalid", errors.New("generator algorithm is invalid"))
		return
	}
	id, err := generator.GenerateUniqueID()
	if err != nil {
		status = erpc.NewStatus(1, err.Error(), err)
		return
	}
	resp = &models.IDGenerateResponse{
		ID: id,
	}
	return
}
