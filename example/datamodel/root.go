package root

import (
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/datamodel/config"
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/datamodel/nexus"
)

type Root struct {
	nexus.Node
	DisplayName string
	Config      config.Config `nexus:"child"`
	CustomBar   Bar
}

type Bar struct {
	Name string
}
