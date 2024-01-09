package output

import (
	"fmt"
	"strings"
	"sync"
	"utiloci/internal/data"
	"utiloci/models"
	"utiloci/pkg"

	"github.com/fatih/color"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/oracle/oci-go-sdk/v58/common"
)

func Output(resList identity.ListRegionsResponse) {
	c := color.New(color.FgCyan).Add(color.BgBlue)
	var wg sync.WaitGroup
	wg.Add(len(resList.Items))
	for _, resList := range resList.Items {
		//fmt.Println("\n...checking", *resList.Name)
		//time.Sleep(2 * time.Second)
		region := *resList.Name
		go pkg.ListVMs(common.NewRawConfigurationProvider(
			data.Credential.TenancyOcid, data.Credential.UserOcid, region, data.Credential.Fingerprint, string(data.Credential.PrivateKey), nil,
		), data.Credential.CompartmentOcid, &wg)
	}
	wg.Wait()

	for region, lines := range models.OutputVM {
		fmt.Println("\n" + region)
		for i, line := range lines {

			if i == 0 {
				max := pkg.MaxLen(lines)

				c.Println(line + strings.Repeat(" ", max-len(line)))
				continue
			}

			if strings.Contains(line, "TERMINATED") {
				color.Blue(line)
			} else {
				color.Green(line)
			}
		}
	}

}
