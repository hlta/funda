package initializer

import (
	"funda/configs"
	"funda/internal/logger"
	"funda/internal/model"

	"gorm.io/gorm"
)

// LoadPredefinedRolesAndPermissions loads predefined roles and permissions from configuration and updates the database.
func LoadPredefinedRolesAndPermissions(db *gorm.DB, config configs.RolesPermissionsConfig, log logger.Logger) {
	for _, permission := range config.Permissions {
		var existingPermission model.Permission
		err := db.Where("name = ?", permission.Name).First(&existingPermission).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Error("Failed to check permission: ", err)
			continue
		}
		if existingPermission.ID == 0 {
			db.Create(&model.Permission{Name: permission.Name})
		} else {
			existingPermission.Name = permission.Name
			db.Save(&existingPermission)
		}
	}

	for _, role := range config.Roles {
		var existingRole model.Role
		err := db.Where("name = ?", role.Name).First(&existingRole).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Error("Failed to check role: ", err)
			continue
		}

		var permissions []model.Permission
		for _, permName := range role.Permissions {
			var permission model.Permission
			db.Where("name = ?", permName).First(&permission)
			permissions = append(permissions, permission)
		}

		if existingRole.ID == 0 {
			db.Create(&model.Role{Name: role.Name, Permissions: permissions})
		} else {
			existingRole.Name = role.Name
			existingRole.Permissions = permissions
			db.Save(&existingRole)
		}
	}
}
