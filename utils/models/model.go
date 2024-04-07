package model

type (
	CreateUserReq struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		Mobile    string `json:"mobile"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	GenericRes struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta,omitempty"`
	}
	// Role struct {
	// 	RoleName       string           `json:"role_name"`
	// 	RolePermission []RolePermission `json:"role_permission"`
	// }
	// Permission struct {
	// 	Read   bool   `json:"read"`
	// 	Write  bool   `json:"write"`
	// 	Edit   bool   `json:"edit"`
	// 	Remove string `json:"remove"`
	// }
	// RolePermission struct {
	// 	RoleID       int         `json:"role_id"`
	// 	PermissionID int         `json:"permission_id"`
	// 	Permission   *Permission `json:"permission"`
	// 	ResourseID   int         `json:"resourse_id"`
	// }
)
