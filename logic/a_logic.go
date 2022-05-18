package logic

import (
	"fmt"

	"github.com/eryajf-world/go-ldap-admin/public/tools"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))

	Api          = &ApiLogic{}
	User         = &UserLogic{}
	Group        = &GroupLogic{}
	Role         = &RoleLogic{}
	Menu         = &MenuLogic{}
	OperationLog = &OperationLogLogic{}
	Base         = &BaseLogic{}
)