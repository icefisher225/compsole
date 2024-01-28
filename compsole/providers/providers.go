package providers

import (
	"encoding/json"
	"fmt"

	"github.com/BradHacker/compsole/compsole/providers/openstack"
	"github.com/BradHacker/compsole/compsole/utils"
	"github.com/BradHacker/compsole/ent"
	"github.com/icefisher225/compsole/providers/proxmox"
)

type CompsoleProvider interface {
	ID() string
	Name() string
	Author() string
	Version() string
	GetConsoleUrl(vmObject *ent.VmObject, consoleType utils.ConsoleType) (string, error)
	GetPowerState(vmObject *ent.VmObject) (utils.PowerState, error)
	ListVMs() ([]*ent.VmObject, error)
	RestartVM(vmObject *ent.VmObject, rebootType utils.RebootType) error
	PowerOnVM(vmObject *ent.VmObject) error
	PowerOffVM(vmObject *ent.VmObject) error
}

func NewProvider(providerType string, config string) (provider CompsoleProvider, err error) {
	switch providerType {
	case openstack.ID:
		return openstack.NewOpenstackProvider(config)
	case proxmox.ID:
		return proxmox.NewProxmoxProvider(config)
	default:
		err = fmt.Errorf("invalid provider type")
		return
	}
}

func ValidateConfig(providerType string, config string) error {
	switch providerType {
	case openstack.ID:
		var openstackConfig openstack.OpenstackConfig
		return json.Unmarshal([]byte(config), &openstackConfig)
	case proxmox.ID:
		var proxmoxConfig proxmox.ProxmoxConfig
		return json.Unmarshal([]byte(config), &proxmoxConfig)
	default:
		return fmt.Errorf("invalid provider type")
	}
}
