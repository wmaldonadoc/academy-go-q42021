package registry

import "github.com/wmaldonadoc/academy-go-q42021/interface/controller"

func (r *registry) NewHealthController() controller.HealthController {
	return controller.NewHealthController()
}
