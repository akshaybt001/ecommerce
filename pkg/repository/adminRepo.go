package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRespository {
	return &adminDatabase{DB}
}

//-------------------------- Login --------------------------//

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}

//-------------------------- Show-Single-User --------------------------//

func (c *adminDatabase) ShowUser(userID int) (response.UserDetails, error) {
	var userData response.UserDetails
	qury := `SELECT users.name,
			 users.email, 
			 users.mobile, 
			 users.is_blocked, 
			 block_infos.blocked_by,
			 block_infos.blocked_at,
			 block_infos.reason_for_blocking 
			 FROM users as users 
			 FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id
			 WHERE users.id = $1;`

	err := c.DB.Raw(qury, userID).Scan(&userData).Error
	if err != nil {
		return response.UserDetails{}, err
	}
	if userData.Email == "" {
		return response.UserDetails{}, fmt.Errorf("no such user")
	}
	return userData, nil
}

//-------------------------- Show-All-Users --------------------------//

func (c *adminDatabase) ShowAllUser() ([]response.UserDetails, error) {
	var userDatas []response.UserDetails

	getUsers := `SELECT users.name,
				users.email, 
				users.mobile, 
				users.is_blocked, 
				block_infos.blocked_by,
				block_infos.blocked_at,
				block_infos.reason_for_blocking 
				FROM users as users 
				FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id;`

	err := c.DB.Raw(getUsers).Scan(&userDatas).Error
	return userDatas, err
}

// -------------------------- Block-User --------------------------//

func (c *adminDatabase) BlockUser(body helper.BlockData, adminId int) error {
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
	if err := tx.Exec("INSERT INTO user_block_infos (users_id, reason_for_blocking, blocked_at, blocked_by) VALUES (?, ?, NOW(), ?)", body.UserId, body.Reason, adminId).Error; err != nil {
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

func (c *adminDatabase) UnblockUser(id int) error {
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
	query := "UPDATE user_block_infos SET reason_for_blocking=$1,blocked_at=NULL,blocked_by=$2 WHERE users_id=$3"
	if err := tx.Exec(query, "", 0, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
