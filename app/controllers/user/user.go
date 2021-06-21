package user

import "gewu_jxc/app/controllers/user/userControllers"

func Init() {
	userControllers.Register()
	userControllers.CheckSimCode()
	userControllers.Delete()
}
