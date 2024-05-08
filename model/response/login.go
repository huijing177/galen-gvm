package response

import "galen-gvm/model/system"

type LoginResp struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
