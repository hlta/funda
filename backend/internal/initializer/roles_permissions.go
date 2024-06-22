package initializer

import (
	"funda/configs"
	"funda/internal/logger"
	"funda/internal/model"

	"gorm.io/gorm"
)

// LoadPredefinedRolesAndPermissions loads predefined roles and permissions from configuration and updates the database.
func LoadPredefinedRolesAndPermissions(db *gorm.DB, config configs.RolesPermissionsConfig, log logger.Logger) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := loadPermissions(tx, config.Permissions, log); err != nil {
			return err
		}
		if err := loadRoles(tx, config.Roles, log); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("Failed to load predefined roles and permissions: ", err)
	}
}

func loadPermissions(tx *gorm.DB, permissions []configs.PermissionConfig, log logger.Logger) error {
	for _, permConfig := range permissions {
		var permission model.Permission
		err := tx.Where("name = ?", permConfig.Name).First(&permission).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Error("Failed to check permission: ", err)
			return err
		}
		if permission.ID == 0 {
			err = tx.Create(&model.Permission{Name: permConfig.Name}).Error
			if err != nil {
				log.Error("Failed to create permission: ", err)
				return err
			}
		}
	}
	return nil
}

func loadRoles(tx *gorm.DB, roles []configs.RoleConfig, log logger.Logger) error {
	for _, roleConfig := range roles {
		var role model.Role
		err := tx.Where("name = ?", roleConfig.Name).Preload("Permissions").First(&role).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Error("Failed to check role: ", err)
			return err
		}

		permissions, err := getPermissions(tx, roleConfig.Permissions, log)
		if err != nil {
			return err
		}

		if role.ID == 0 {
			role = model.Role{Name: roleConfig.Name, Permissions: permissions}
			err = tx.Create(&role).Error
			if err != nil {
				log.Error("Failed to create role: ", err)
				return err
			}
		} else {
			// Merge existing and new permissions without duplicating
			role.Permissions = mergePermissions(role.Permissions, permissions)
			err = tx.Model(&role).Association("Permissions").Replace(role.Permissions)
			if err != nil {
				log.Error("Failed to update role permissions: ", err)
				return err
			}
		}
	}
	return nil
}

func getPermissions(tx *gorm.DB, permissionNames []string, log logger.Logger) ([]model.Permission, error) {
	var permissions []model.Permission
	for _, permName := range permissionNames {
		var permission model.Permission
		err := tx.Where("name = ?", permName).First(&permission).Error
		if err != nil {
			log.Error("Failed to get permission: ", err)
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func mergePermissions(existing []model.Permission, new []model.Permission) []model.Permission {
	permissionMap := make(map[uint]model.Permission)
	for _, perm := range existing {
		permissionMap[perm.ID] = perm
	}
	for _, perm := range new {
		if _, exists := permissionMap[perm.ID]; !exists {
			existing = append(existing, perm)
		}
	}
	return existing
}
