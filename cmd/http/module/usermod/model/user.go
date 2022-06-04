package model

type User struct {
	Id        string `json:"id" gorm:"column:id;not null;type:varchar(64); primary key;comment:'id'"` //id
	PeerId    string `json:"peerId" gorm:"column:peer_id;unique;type:varchar(64);"`                   //peerId
	Name      string `json:"name" gorm:"column:name;type:varchar(128);"`                              //名字
	Phone     string `json:"phone" gorm:"column:phone;type:varchar(64);"`                             //手机号
	Sex       int64  `json:"sex" gorm:"column:sex;type:int(10);"`                                     //性别 0 未知 1 男 2 女
	Ptime     int64  `json:"ptime" gorm:"column:ptime;type:int(10);"`                                 //时间
	Utime     int64  `json:"utime" gorm:"column:utime;type:int(10);"`                                 //时间
	Nickname  string `json:"nickname" gorm:"column:nickname;type:varchar(128);"`                      //昵称
	Img       string `json:"img" gorm:"column:img;type:varchar(128);"`                                //头像图片
	Role      string `json:"role" gorm:"column:role;type:varchar(8);"`                                //角色 1 人工客服  2 普通用户
	InviterId string `json:"inviter_id" gorm:"column:inviter_id;type:varchar(64);default:''"`         //邀请人ID
	Lang      string `json:"lang" gorm:"column:lang;type:varchar(8);default:'ZH';"`                   //语言
	Status    int64  `json:"status" gorm:"column:status;type:int(10);not null;default:0"`             //状态（0启用 -1禁用）
	Profile   string `json:"profile" gorm:"column:profile;type:varchar(256);default:''"`              //个人简介
}

func (User) TableName() string {
	return "user"
}
