package system

import (
	"errors"

	"galen-gvm/global"
	"galen-gvm/utils"

	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId uint      `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Enable      int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                      //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}

func GetUserbyUsername(name string) (SysUser, error) {
	var user SysUser
	err := global.GVA_DB.Where("username = ?", name).First(&user).Error
	return user, err
}

func Register(u *SysUser) (*SysUser, error) {
	var userTmp SysUser
	err := global.GVA_DB.Where("username = ?", u.Username).First(&userTmp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 密码加密存储，1、防止监守自盗。2、防止数据信息泄露后，密码也被得知
			u.Password = utils.BcryptHash(u.Password)
			u.UUID = uuid.Must(uuid.NewV4())
			// TODO zap.Any能否正常打印user信息
			global.GVA_LOG.Info("Register user:", zap.Any("", u))
			return u, global.GVA_DB.Create(&u).Error
		}
		global.GVA_LOG.Error("Register err:", zap.Error(err))
		return nil, err
	}
	return nil, errors.New("用户名已存在")
}

func ChangePassword(u *SysUser, newPassword string) (*SysUser, error) {
	var userTmp SysUser
	err := global.GVA_DB.Where("id = ?", u.ID).First(&userTmp).Error
	if err != nil {
		global.GVA_LOG.Error("ChangePassword err:", zap.Error(err))
		return nil, err
	}
	if utils.BcryptCheck(u.Password, userTmp.Password) {
		global.GVA_LOG.Error("原密码错误", zap.Any("user", u))
		return u, errors.New("原密码错误")
	}
	u.Password = utils.BcryptHash(newPassword)
	global.GVA_LOG.Info("ChangePassword user:", zap.Any("user", u))
	return u, global.GVA_DB.Save(&u).Error
}
