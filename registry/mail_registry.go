package registry

import "github.com/ecea-nitt/ecea-server/services"

func (r *registry) NewMailService() services.MailService {
	return services.NewMailService()
}
