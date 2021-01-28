// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"reflect"
	"wizz-home-page/app/dao/internal"
	"wizz-home-page/app/model"

	"github.com/gogf/gf/frame/g"
)

// userRoleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type userRoleDao struct {
	*internal.UserRoleDao
}

var (
	// UserRole is globally public accessible object for table user_role operations.
	UserRole = &userRoleDao{
		internal.UserRole,
	}
)

// Fill with you ideas below.

func GetRolesByUser(user *model.User) []model.Role {
	var roles []model.Role
	g.Log().Line().Debug(reflect.TypeOf(roles))

	if roleIds, err := g.DB().Table("user_role").Array("role_id", "user_id = ", user.Id); err != nil {
		g.Log().Line().Error(err)
		return nil
	} else {
		// g.Log().Line().Debug(roleIds)
		if err := g.DB().Table("role").Where("id IN (?)", roleIds).Structs(&roles); err != nil {
			g.Log().Line().Error(err)
			return nil
		}
		// g.Log().Line().Debug(roles)
		return roles
	}
}

func HasRole(user *model.User, roleName string) bool {
	roles := GetRolesByUser(user)
	for _, r := range roles {
		if roleName == r.Name {
			return true
		}
	}
	return false
}