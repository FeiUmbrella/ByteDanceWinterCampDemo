package model

import "gorm.io/gorm"

// app目录文件里面是RPC的服务端；rpc_gen是RPC的客户端
type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}
