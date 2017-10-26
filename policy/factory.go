package policy

import (
	"goweb/models"

	"strings"
)

type Factory struct {
}

func (this *Factory) FindPolicy(jobinfo *models.JobInfo) Policy {

	switch jobinfo.Type {
	case "PRIORITY":
		return &PriorityPolicy{urls: strings.Split(jobinfo.Urls, ","), retryMaxTimes: 3}
	default:
		return &RandomPolicy{urls: strings.Split(jobinfo.Urls, ","), retryMaxTimes: 3}

	}
}
