package proxmox

import (
	"encoding/json"
	"fmt"

	"github.com/BradHacker/compsole/compsole/utils"
	"github.com/BradHacker/compsole/ent"
)

// #########
// # TYPES #
// #########
type CompsoleProviderProxmox struct {
	Config ProxmoxConfig
}

type ProxmoxConfig struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	URL         string `json:"url"`
	NoVNCConfig string `json:"novnc_config"`
}

type NoVNCConfig struct {
	Autoconnect    string `json:"autoconnect"`     // Automatically connect as soon as the page has finished loading.
	Reconnect      string `json:"reconnect"`       // If noVNC should automatically reconnect if the connection is dropped.
	ReconnectDelay string `json:"reconnect_delay"` // How long to wait in milliseconds before attempting to reconnect.
	Host           string `json:"host"`            // The WebSocket host to connect to.
	Port           string `json:"port"`            // The WebSocket port to connect to.
	Encrypt        string `json:"encrypt"`         // If TLS should be used for the WebSocket connection.
	Path           string `json:"path"`            // The WebSocket path to use.
	Password       string `json:"password"`        // The password sent to the server, if required.
	RepeaterID     string `json:"repeater_id"`     // The repeater ID to use if a VNC repeater is detected.
	Shared         string `json:"shared"`          // If other VNC clients should be disconnected when noVNC connects.
	Bell           string `json:"bell"`            // If the keyboard bell should be enabled or not.
	ViewOnly       string `json:"view_only"`       // If the remote session should be in non-interactive mode.
	ViewClip       string `json:"view_clip"`       // If the remote session should be clipped or use scrollbars if it cannot fit in the browser.
	Resize         string `json:"resize"`          // How to resize the remote session if it is not the same size as the browser window. Can be one of off, scale and remote.
	Quality        string `json:"quality"`         // The session JPEG quality level. Can be 0 to 9.
	Compression    string `json:"compression"`     // The session compression level. Can be 0 to 9.
	ShowDot        string `json:"show_dot"`        // If a dot cursor should be shown when the remote server provides no local cursor, or provides a fully-transparent (invisible) cursor.
	Logging        string `json:"logging"`         // The console log level. Can be one of error, warn, info or debug.
}

// ############
// # METADATA #
// ############
const (
	ID      string = "NoVNC"
	Name    string = "Ryan Cheevers-Brown"
	Author  string = "ma1ist4ir3"
	Version string = "v0.1"
)

func (provider CompsoleProviderProxmox) ID() string      { return ID }
func (provider CompsoleProviderProxmox) Name() string    { return Name }
func (provider CompsoleProviderProxmox) Author() string  { return Author }
func (provider CompsoleProviderProxmox) Version() string { return Version }

// #############
// # FUNCTIONS #
// #############

// NewNoVNCProvider creates a provider for noVNC
func NewProxmoxProvider(config string) (provider CompsoleProviderProxmox, err error) {
	var providerConfig ProxmoxConfig
	err = json.Unmarshal([]byte(config), &providerConfig)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal Proxmox Config: %v", err)
		return
	}
	provider = CompsoleProviderProxmox{
		Config: providerConfig,
	}
	return
}

func GetConsoleUrl(vmObject *ent.VmObject, consoleType utils.ConsoleType) (string, error) {
	// NYI
	return "Not yet implemented", nil
}

func GetPowerState(vmObject *ent.VmObject) (utils.PowerState, error){
	return "Not yet implemented", nil
}

func ListVMs() ([]*ent.)
