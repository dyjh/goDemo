package controller

import (
	"demo/admin/service"
	"demo/model"
	"demo/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type DrawController struct {
	Ctx iris.Context
	DrawService service.DrawService
	Session *sessions.Session
}

func (dc *DrawController) PostAddDrawActive() mvc.Result {
	var active model.LuckyDrawActive
	if err := dc.Ctx.ReadForm(&active); err != nil {
		panic(err.Error())
	} else {
		errString := utils.ValidateData(active)
		if errString != "success" {
			return utils.Api(utils.RECODE_FAIL, errString, map[string]interface{}{})
		}
	}
	result := dc.DrawService.AddDrawActive(active)
	if result.Code == 0 {
		return utils.Api(utils.RECODE_FAIL, "success", map[string]interface{}{})
	} else {
		return utils.Api(utils.RECODE_OK, result.Message, map[string]interface{}{})
	}
}


func (dc *DrawController) GetDeleteDrawActive() mvc.Result {
	activeId, err := dc.Ctx.Params().GetInt64("active_id")
	if err != nil {
		//panic(err.Error())
		return utils.Api(utils.RECODE_OK, err.Error(), map[string]interface{}{})
	}
	result := dc.DrawService.DeleteDrawActive(activeId)
	if result.Code == 0 {
		return utils.Api(utils.RECODE_FAIL, "success", map[string]interface{}{})
	} else {
		return utils.Api(utils.RECODE_OK, result.Message, map[string]interface{}{})
	}
}


func (dc *DrawController) GetDrawActiveInfo() mvc.Result {
	activeId, err := dc.Ctx.Params().GetInt64("active_id")
	if err != nil {
		//panic(err.Error())
		return utils.Api(utils.RECODE_OK, err.Error(), map[string]interface{}{})
	}
	result := dc.DrawService.DeleteDrawActive(activeId)
	if result.Code == 0 {
		return utils.Api(utils.RECODE_FAIL, "success", map[string]interface{}{})
	} else {
		return utils.Api(utils.RECODE_OK, result.Message, map[string]interface{}{})
	}
}