package main

import (
	"context"
	"os"
	"utiloci/internal/data"
	defaultregion "utiloci/internal/defaultRegion"
	"utiloci/internal/output"

	"utiloci/pkg"

	"github.com/oracle/oci-go-sdk/identity"
	"github.com/oracle/oci-go-sdk/v58/common"
)

func main() {

	//Resq for default region
	defaultregion.Resq()

	//clear
	pkg.RemoveLinesUp(39)

	//set path
	pathCredential := ""

	if len(os.Args) > 1 {
		pathCredential = os.Args[1]
	}

	//set credential
	pkg.GetCredential(pathCredential)

	configProvider := common.NewRawConfigurationProvider(
		data.Credential.TenancyOcid, data.Credential.UserOcid, data.DefaultRegion, data.Credential.Fingerprint, string(data.Credential.PrivateKey), nil,
	)

	client, _ := identity.NewIdentityClientWithConfigurationProvider(configProvider)
	resList, _ := client.ListRegions(context.Background())

	output.Output(resList)

}
