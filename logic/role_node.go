package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/qingcc/goblog/model"
)

type RoleNodeLogic struct{}

var DefaultRoleNode = RoleNodeLogic{}

//region Remark: 获取角色 Author:Qing
func (self RoleNodeLogic) FindRoleNodes(c *gin.Context, role_id int64) []*model.RoleNode {
	objLog := GetLogger(c)
	nodes := make([]*model.RoleNode, 0)
	err := databases.Orm.Where("role_id = ?", role_id).Find(&nodes)
	if err != nil {
		objLog.Errorf("RoleNodeLogic find errof:", err)
		return nil
	}

	return nodes
}

//endregion
