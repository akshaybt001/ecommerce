package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type supadminDatabase struct {
	DB *gorm.DB
}

func NewSupAdminRepository(DB *gorm.DB)interfaces.SupAdminRepository{
	return &supadminDatabase{DB}
}

// -------------------------- Login --------------------------//

func (c *supadminDatabase) SupAdminLogin(email string) (domain.SupAdmins, error) {
	var supadminData domain.SupAdmins
	err := c.DB.Raw("SELECT * FROM sup_admins WHERE email=?", email).Scan(&supadminData).Error
	return supadminData, err
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
