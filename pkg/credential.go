package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"utiloci/internal/data"
	"utiloci/models"
)

func GetCredential(path string) {
	var credential map[string]string
	var file *os.File
	var err error
	if path == "" {
		files, err := filepath.Glob("*.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		// if no have .josn
		if len(files) == 0 {
			panic("no file .json")
		}

		file, err = os.Open(files[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {

		file, err = os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	// read
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Convertir el JSON en una variable map
	err = json.Unmarshal(bytes, &credential)
	if err != nil {
		fmt.Println(err)
		return
	}

	data.Credential = models.MyCredential{
		UserOcid:        credential["user_ocid"],
		TenancyOcid:     credential["tenancy_ocid"],
		CompartmentOcid: credential["compartment_ocid"],
		PrivateKey:      credential["key_file"],
		Fingerprint:     credential["fingerprint"],
	}

}
