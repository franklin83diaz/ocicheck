package pkg

import (
	"context"
	"fmt"
	"sync"
	"utiloci/models"

	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/core"
)

func ListVMs(configProvider common.ConfigurationProvider, compartmentOcid string, wg *sync.WaitGroup) {

	computeClient, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Println("Error create Compute:", err)
		return
	}

	// Create resq for list instance
	request := core.ListInstancesRequest{
		CompartmentId: common.String(compartmentOcid),
	}

	// Call API list instance with the resques
	resp, err := computeClient.ListInstances(context.Background(), request)

	if err != nil {
		wg.Done()
		return
	}
	//c := color.New(color.FgWhite).Add(color.Underline)

	region, _ := configProvider.Region()
	if len(resp.Items) > 0 {
		//	max := maxLen(resp.Items) + 15
		//	c.Println(region, strings.Repeat(" ", max-len(region)))
	} else {
		wg.Done()
		return
	}
	var lines []string
	//head
	lines = append(lines, fmt.Sprintf("%-45s: %5s \t [%s] (%s)", "Display Name", "State", "Tags", "Created"))

	for _, instance := range resp.Items {

		if instance.LifecycleState == "TERMINATED" {
			lines = append(lines, fmt.Sprintf("%-45s: %5s \t [%s] (%s) \n", *instance.DisplayName, instance.LifecycleState, GetTagInfo(instance), instance.TimeCreated.Time))

		} else {
			lines = append(lines, fmt.Sprintf("%-45s: %5s \t [%s] (%s) \n", *instance.DisplayName, instance.LifecycleState, GetTagInfo(instance), instance.TimeCreated.Time))
		}
	}

	models.OutputVM[region] = lines

	wg.Done()

}
