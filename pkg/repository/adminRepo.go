package repository

import (
	"fmt"

	"gorm.io/gorm"
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
	qury := `SELECT users.id,
			 users.name,
			 users.email, 
			 users.mobile, 
			 users.is_blocked, 
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

	getUsers := `SELECT users.id,
				users.name,
				users.email, 
				users.mobile, 
				users.is_blocked,
				block_infos.blocked_at,
				block_infos.reason_for_blocking 
				FROM users as users 
				FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id;`

	err := c.DB.Raw(getUsers).Scan(&userDatas).Error
	return userDatas, err
}
