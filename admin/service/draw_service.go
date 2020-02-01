package service

import (
	"demo/model"
	"demo/utils"
	"github.com/go-xorm/xorm"
)

type DrawService interface {
	AddDrawActive (data model.LuckyDrawActive) utils.ReturnData
	DeleteDrawActive (ActiveId int64) utils.ReturnData
}

func NewDrawService(engine *xorm.Engine) DrawService {
	return &drawService{
		Engine: engine,
	}
}

type drawService struct {
	Engine *xorm.Engine
}

func (uc *drawService) DeleteDrawActive(ActiveId int64) utils.ReturnData {
	var active model.LuckyDrawActive
	rowNum, err := uc.Engine.Id(ActiveId).Delete(&active)
	if err != nil {
		return utils.ReturnData{
			Message: err.Error(),
			Data:    nil,
			Code:    1,
		}
	}
	if rowNum == 0 {
		return utils.ReturnData{
			Message: "数据添加失败",
			Data:    nil,
			Code:    2,
		}
	}
	return utils.ReturnData{
		Message: "success",
		Data:    nil,
		Code:    0,
	}
}

func (uc *drawService) AddDrawActive(data model.LuckyDrawActive) utils.ReturnData {
	rowNum, err := uc.Engine.Insert(&data)
	if err != nil {
		return utils.ReturnData{
			Message: err.Error(),
			Data:    nil,
			Code:    1,
		}
	}
	if rowNum == 0 {
		return utils.ReturnData{
			Message: "数据添加失败",
			Data:    nil,
			Code:    2,
		}
	}
	return utils.ReturnData{
		Message: "success",
		Data:    nil,
		Code:    0,
	}
}

