package models

/*
import (
	"errors"

	"github.com/jxeldotdev/url-backend/internal/db"
	pq "github.com/lib/pq"
	"gorm.io/gorm"
)

type ObjectAccess string

const (
	Any ObjectAccess = "Any"
	Own ObjectAccess = "Own"
)

type ActionAccess string

const (
	View     ActionAccess = "View"
	Create   ActionAccess = "Create"
	Update   ActionAccess = "Update"
	Delete   ActionAccess = "Delete"
	Redirect ActionAccess = "Redirect"
	Register ActionAccess = "Register"
	Login    ActionAccess = "Login"
)

type ApiPermission struct {
	ID           uint           `gorm:"primaryKey"`
	Name         string         `gorm:"unique"`
	ObjectAccess string         `gorm:"type:varchar(20);check:object_access IN ('Any','Own')"`
	ActionAccess pq.StringArray `gorm:"type:text[];not null"`
	Roles        []*Role        `gorm:"many2many:api_permission_roles;"`
}

type RoleName string

const (
	UnknownRole RoleName = "Unknown"
	UserRole    RoleName = "User"
	AdminRole   RoleName = "Admin"
)

type Role struct {
	Id          uint             `gorm:"primaryKey"`
	Name        string           `gorm:"unique"`
	Scope       string           `gorm:"not null"`
	Permissions []*ApiPermission `gorm:"many2many:api_permission_roles;"`
	UserIds		pq.Int64Array	 `gorm:j`
}

func createRoleIfNotPresent(role *Role) (*Role, error) {
	// crates role if it is not present

	var roleInDb Role
	err := db.Database.Where("name = ?", role.Name).First(&roleInDb).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create role
			err := db.Database.Create(&role).Error
			if err != nil {
				return role, nil
			}
			return &Role{}, nil
		}
		return &Role{}, err
	}
	return role, nil
}

func createPermissionIfNotPresent(perm *ApiPermission) (*ApiPermission, error) {
	// crates role if it is not present

	var permissionInDb ApiPermission
	err := db.Database.Where("name = ?", perm.Name).First(&permissionInDb).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create role
			err := db.Database.Create(&perm).Error
			if err != nil {
				return &ApiPermission{}, nil
			}
			return perm, nil
		}
		return perm, err
	}
	return &permissionInDb, nil
}

func CreateRoles() {
	anonPermission := &ApiPermission{
		Name:         "Access Any URL via redirect link",
		ObjectAccess: "Any",
		ActionAccess: []string{"Redirect", "Login", "Register"},
	}

	createPermissionIfNotPresent(anonPermission)

	anon := &Role{
		Name:        "Anonymous",
		Permissions: []*ApiPermission{*&anonPermission},
	}

	createRoleIfNotPresent(anon)

	userPermission := &ApiPermission{
		Name:         "Manage own URLs and account",
		ObjectAccess: "Own",
		ActionAccess: []string{"Create", "Update", "Delete", "View"},
	}

	createPermissionIfNotPresent(userPermission)

	user := &Role{
		Name:        "User",
		Permissions: []*ApiPermission{*&userPermission, *&anonPermission},
	}

	createRoleIfNotPresent(user)

	adminPermission := &ApiPermission{
		Name:         "Full permissions",
		ObjectAccess: "Any",
		ActionAccess: []string{"Redirect", "Create", "Update", "Delete", "View"},
	}

	createPermissionIfNotPresent(adminPermission)

	admin := &Role{
		Name:        "Admin",
		Permissions: []*ApiPermission{*&anonPermission, *&userPermission, *&adminPermission},
	}

	createRoleIfNotPresent(admin)

}

*/
