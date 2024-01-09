package pkg

import (
	"strings"

	"github.com/oracle/oci-go-sdk/v58/core"
)

func GetTagInfo(instance core.Instance) (uc string) {

	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	uc = strings.ReplaceAll(strings.Split(instance.Metadata["ssh_authorized_keys"], " ")[2], "\n", "")
	return uc
}
