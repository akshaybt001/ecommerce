package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type supadminDatabase struct {
	DB *gorm.DB
}

func NewSupAdminRepository(DB *gorm.DB) interfaces.SupAdminRepository {
	return &supadminDatabase{DB}
}

// -------------------------- Login --------------------------//

func (c *supadminDatabase) SupAdminLogin(email string) (domain.SupAdmins, error) {
	var supadminData domain.SupAdmins
	err := c.DB.Raw("SELECT * FROM sup_admins WHERE email=?", email).Scan(&supadminData).Error
	return supadminData, err
}

// CreateAdmin implements interfaces.SupAdminRepository.
func (c *supadminDatabase) CreateAdmin(admin helper.CreateAdmin) (response.AdminData, error) {
	var newAdmin response.AdminData
	insertQuery := `INSERT INTO admins(name,email,password,created_at)VALUES($1,$2,$3,NOW()) RETURNING name,email,id`
	err := c.DB.Raw(insertQuery, admin.Name, admin.Email, admin.Password).Scan(&newAdmin).Error
	return newAdmin, err
}

// ListAllAdmins implements interfaces.SuperAdminRepository.
func (c *supadminDatabase) ListAllAdmins() ([]response.AdminDetails, error) {
	var admins []response.AdminDetails
	getAdmins := `SELECT * FROM admins`
	err := c.DB.Raw(getAdmins).Scan(&admins).Error
	return admins, err
}

// DisplayAdmin implements interfaces.SuperAdminRepository.

func (c *supadminDatabase) DisplayAdmin(id int) (response.AdminDetails, error) {
	var admin response.AdminDetails
	err := c.DB.Raw(`SELECT * FROM admins WHERE id=?`, id).Scan(&admin).Error
	if admin.Email == "" {
		return admin, fmt.Errorf("no admin found with given id")
	}
	return admin, err
}

// -------------------------- Block-User --------------------------//

func (c *supadminDatabase) BlockUser(body helper.BlockData) error {
	// Start a transaction
	tx := c.DB.Begin()
	//Check if the user is there
	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", body.UserId).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}

	// Execute the first SQL command (UPDATE)
	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = ?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Execute the second SQL command (INSERT)
	if err := tx.Exec("INSERT INTO user_block_infos (users_id, reason_for_blocking, blocked_at) VALUES (?, ?, NOW())", body.UserId, body.Reason).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	// If all commands were executed successfully, return nil
	return nil

}

// -------------------------- UnBlock-User --------------------------//

func (c *supadminDatabase) UnblockUser(id int) error {
	tx := c.DB.Begin()

	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND is_blocked=true)", id).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")
	}
	if err := tx.Exec("UPDATE users SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE user_block_infos SET reason_for_blocking=$1,blocked_at=NULL WHERE users_id=$2"
	if err := tx.Exec(query, "", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// -------------------------- Block-Admin --------------------------//

func (c *supadminDatabase) BlockAdmin(body helper.BlockAdminData) error {
	// Start a transaction
	tx := c.DB.Begin()
	//Check if the user is there
	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM admins WHERE id = $1)", body.AdminId).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such Admin")
	}

	// Execute the first SQL command (UPDATE)
	if err := tx.Exec("UPDATE admins SET is_blocked = true WHERE id = ?", body.AdminId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Execute the second SQL command (INSERT)
	if err := tx.Exec("INSERT INTO admin_block_infos (admin_id, reason_for_blocking, blocked_at) VALUES (?, ?, NOW())", body.AdminId, body.Reason).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	// If all commands were executed successfully, return nil
	return nil

}

// -------------------------- UnBlock-Admin --------------------------//

func (c *supadminDatabase) UnblockAdmin(id int) error {
	tx := c.DB.Begin()

	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM admins WHERE id = $1 AND is_blocked=true)", id).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such admin to unblock")
	}
	if err := tx.Exec("UPDATE admins SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE admin_block_infos SET reason_for_blocking=$1,blocked_at=NULL WHERE admin_id=$2"
	if err := tx.Exec(query, "", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
