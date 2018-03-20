package endpoints

import (
	"fmt"

	"github.com/TimWoolford/podrick/internal/config"
	"github.com/TimWoolford/podrick/internal/util"
)

type K8sEndpoint struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	address     string
	port        int32
	config      *config.Config
}


func (ep K8sEndpoint) StatusUrl(statusPath string) string {
	sp := util.FirstPopulated(statusPath, ep.Annotations[ep.config.StatusPathAnnotation], ep.config.DefaultStatusPath)
	return fmt.Sprintf("http://%s:%d%s", ep.address, ep.port, sp)
}
