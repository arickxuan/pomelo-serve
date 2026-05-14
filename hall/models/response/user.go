package response

import (
	"pomeloServe/common"
	"pomeloServe/hall/models/request"
)

type UpdateUserAddressRes struct {
	common.Result
	UpdateUserData request.UpdateUserAddressReq `json:"updateUserData"`
}
