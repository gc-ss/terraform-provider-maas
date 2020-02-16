package api

import (
	"github.com/roblox/terraform-provider-maas/pkg/api/params"
	"github.com/roblox/terraform-provider-maas/pkg/maas/entity"
)

type Subnets interface {
	Get() ([]entity.Subnet, error)
	Post(*params.Subnet) (*entity.Subnet, error)
}
