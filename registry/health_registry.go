package registry

import "github.com/wmaldonadoc/academy-go-q42021/interface/controller"

// NewHealthController - Creates an instance of controller.
func (r *registry) NewHealthController() controller.HealthController {
	return controller.NewHealthController()
}
