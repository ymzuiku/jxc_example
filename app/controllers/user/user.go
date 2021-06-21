package user

import "gewu_jxc/app/controllers/user/userControllers"

func UserInit() {
	userControllers.Register()
	userControllers.CheckSimCode()
	userControllers.Delete()
}
