package config

import (
	"gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/datamodel/config/gns"
        "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/common-library.git/pkg/nexus"
	"net/http"
)

var BarCustomCodesResponses = nexus.HTTPCodesResponse{
	http.StatusBadRequest: nexus.HTTPResponse{Description: "Bad Request"},
}

var BarCustomMethodsResponses = nexus.HTTPMethodsResponses{
	http.MethodPatch: BarCustomCodesResponses,
}

type Config struct {
	nexus.Node
	GNS gns.Gns `nexus:"child"`
}
