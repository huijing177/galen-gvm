package response

import "galen-gvm/model/system"

type RegisterResp struct {
	User system.SysUser `json:"user"`
}
