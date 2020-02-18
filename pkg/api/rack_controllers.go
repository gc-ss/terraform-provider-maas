package api

import (
	"github.com/roblox/terraform-provider-maas/pkg/api/params"
	"github.com/roblox/terraform-provider-maas/pkg/maas/entity"
)

type RackControllers interface {
	Get(*params.RackControllerSearch) ([]entity.RackController, error)
}
