import service from '@/utils/request'

// @Tags Template
// @Summary 创建Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "创建Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/createTemplate [post]
export const createTemplate = (data) => {
  return service({
    url: '/deviceManagement/template/createTemplate',
    method: 'post',
    data
  })
}

// @Tags Template
// @Summary 删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteTemplate [delete]
export const deleteTemplate = (data) => {
  return service({
    url: '/deviceManagement/template/deleteTemplate',
    method: 'delete',
    data
  })
}

// @Tags Template
// @Summary 删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /device/deleteTemplate [delete]
export const deleteTemplateByIds = (data) => {
  return service({
    url: '/deviceManagement/template/deleteTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags Template
// @Summary 更新Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Template true "更新Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /device/updateTemplate [put]
export const updateTemplate = (data) => {
  return service({
    url: '/deviceManagement/template/updateTemplate',
    method: 'put',
    data
  })
}

// @Tags Template
// @Summary 用id查询Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Template true "用id查询Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /device/findTemplate [get]
export const findTemplate = (params) => {
  return service({
    url: '/deviceManagement/template/findTemplate',
    method: 'get',
    params
  })
}

// @Tags Template
// @Summary 分页获取Template列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Template列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /device/getTemplateList [get]
export const getTemplateList = (params) => {
  return service({
    url: '/deviceManagement/template/getTemplateList',
    method: 'get',
    params
  })
}
