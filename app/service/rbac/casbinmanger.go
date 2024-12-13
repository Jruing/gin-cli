package rbac

import (
	"github.com/casbin/casbin/v2"
)

var (
	Enforcer *casbin.Enforcer
)

func init() {
	Enforcer, _ = casbin.NewEnforcer("config/rbac_model.conf", "config/policy.csv")
}

// 获取domain下面某个角色包含的所有用户名称
func GetUsersForRoleInDomain(rolename string, domain string) []string {
	data := Enforcer.GetUsersForRoleInDomain(rolename, domain)
	return data
}

// 获取domain下面某个用户所拥有的角色
func GetRolesForUserInDomain(username string, domain string) []string {
	data := Enforcer.GetRolesForUserInDomain(username, domain)
	return data
}

// 获取domain下面某个用户所拥有的权限
func GetPermissionsForUserInDomain(username string, domain string) [][]string {
	data := Enforcer.GetPermissionsForUserInDomain(username, domain)
	return data
}

// 将用户添加到某个角色下，若已添加则返回false
func AddRoleForUserInDomain(username string, rolename string, domain string) bool {
	data, err := Enforcer.AddRoleForUserInDomain(username, rolename, domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 删除domain下用户所拥有的某个角色,若用户没有该角色，则返回false
func DeleteRoleForUserInDomain(username string, rolename string, domain string) bool {
	data, err := Enforcer.DeleteRoleForUserInDomain(username, rolename, domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 删除domain下用户拥有的所有角色，如果用户没有角色，则返回false
func DeleteRolesForUserInDomain(username string, domain string) bool {
	data, err := Enforcer.DeleteRolesForUserInDomain(username, domain)
	if err != nil {
		return false
	} else {
		return data

	}
}

// 获取domain下面的所有用户
func GetAllUsersByDomain(domain string) []string {
	data, err := Enforcer.GetAllUsersByDomain(domain)
	if err != nil {
		return []string{}
	} else {
		return data
	}
}

// 删除domain下面的所有用户
func DeleteAllUsersByDomain(domain string) bool {
	data, err := Enforcer.DeleteAllUsersByDomain(domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 获取所有domain
func GetAllDomains() []string {
	domains, err := Enforcer.GetAllDomains()
	if err != nil {
		return []string{}
	} else {
		return domains
	}
}

// 获取用户具有的角色
func GetRolesForUser(username string, domain string) []string {
	data, err := Enforcer.GetRolesForUser(username, domain)
	if err != nil {
		return []string{}
	} else {
		return data
	}
}

// 获取用户下的角色
func GetUsersForRole(username string, domain string) []string {
	data, err := Enforcer.GetUsersForRole(username, domain)
	if err != nil {
		return []string{}
	} else {
		return data
	}
}

// 给用户分配角色
func AddRoleForUser(username string, rolename string, domain string) bool {
	data, err := Enforcer.AddRoleForUser(username, rolename, domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 给用户分配多个角色
func AddRolesForUser(username string, rolename []string, domain string) bool {
	data, err := Enforcer.AddRolesForUser(username, rolename, domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 删除用户的角色
func DeleteRoleForUser(username string, rolename string, domain string) bool {
	data, err := Enforcer.DeleteRoleForUser(username, rolename, domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 删除用户的所有角色
func DeleteRolesForUser(username string, domain string) bool {
	data, err := Enforcer.DeleteRolesForUser(username, domain)
	if err != nil {
		return false
	} else {
		return data
	}
}

// 删除domain下面的角色
func DeleteRole() {
}
