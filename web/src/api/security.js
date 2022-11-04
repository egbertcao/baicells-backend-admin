import service from '@/utils/request'

// @Tags Authentication
// @Summary 创建Authentication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authentication true "创建Authentication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mqtt/createAuthentication [post]
export const createAuthentication = (data) => {
  return service({
    url: '/deviceManagement/mqtt/createAuthentication',
    method: 'post',
    data
  })
}

// @Tags Authentication
// @Summary 删除Authentication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authentication true "删除Authentication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mqtt/deleteAuthentication [delete]
export const deleteAuthentication = (data) => {
  return service({
    url: '/deviceManagement/mqtt/deleteAuthentication',
    method: 'delete',
    data
  })
}

// @Tags Authentication
// @Summary 删除Authentication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Authentication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mqtt/deleteAuthentication [delete]
export const deleteAuthenticationByIds = (data) => {
  return service({
    url: '/deviceManagement/mqtt/deleteAuthenticationByIds',
    method: 'delete',
    data
  })
}

// @Tags Authentication
// @Summary 更新Authentication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authentication true "更新Authentication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mqtt/updateAuthentication [put]
export const updateAuthentication = (data) => {
  return service({
    url: '/deviceManagement/mqtt/updateAuthentication',
    method: 'put',
    data
  })
}

// @Tags Authentication
// @Summary 用id查询Authentication
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Authentication true "用id查询Authentication"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mqtt/findAuthentication [get]
export const findAuthentication = (params) => {
  return service({
    url: '/deviceManagement/mqtt/findAuthentication',
    method: 'get',
    params
  })
}

// @Tags Authentication
// @Summary 分页获取Authentication列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Authentication列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mqtt/getAuthenticationList [get]
export const getAuthenticationList = (params) => {
  return service({
    url: '/deviceManagement/mqtt/getAuthenticationList',
    method: 'get',
    params
  })
}


// @Tags Authorization
// @Summary 创建Authorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authorization true "创建Authorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mqtt/createAuthorization [post]
export const createAuthorization = (data) => {
  return service({
    url: '/deviceManagement/mqtt/createAuthorization',
    method: 'post',
    data
  })
}

// @Tags Authorization
// @Summary 删除Authorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authorization true "删除Authorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mqtt/deleteAuthorization [delete]
export const deleteAuthorization = (data) => {
  return service({
    url: '/deviceManagement/mqtt/deleteAuthorization',
    method: 'delete',
    data
  })
}

// @Tags Authorization
// @Summary 删除Authorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Authorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mqtt/deleteAuthorization [delete]
export const deleteAuthorizationByIds = (data) => {
  return service({
    url: '/deviceManagement/mqtt/deleteAuthorizationByIds',
    method: 'delete',
    data
  })
}

// @Tags Authorization
// @Summary 更新Authorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Authorization true "更新Authorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mqtt/updateAuthorization [put]
export const updateAuthorization = (data) => {
  return service({
    url: '/deviceManagement/mqtt/updateAuthorization',
    method: 'put',
    data
  })
}

// @Tags Authorization
// @Summary 用id查询Authorization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Authorization true "用id查询Authorization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mqtt/findAuthorization [get]
export const findAuthorization = (params) => {
  return service({
    url: '/deviceManagement/mqtt/findAuthorization',
    method: 'get',
    params
  })
}

// @Tags Authorization
// @Summary 分页获取Authorization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Authorization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mqtt/getAuthorizationList [get]
export const getAuthorizationList = (params) => {
  return service({
    url: '/deviceManagement/mqtt/getAuthorizationList',
    method: 'get',
    params
  })
}
