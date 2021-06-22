package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/otiai10/copy"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// UpdateCatrina check the latest version available for catrina and update the files if is necessary.
func UpdateCatrina(version, url string) error {
	// Check version
	resp, err := http.Get(fmt.Sprintf("%v?version=%v&os=%v", url, version, runtime.GOOS))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dataUpdate UpdateResponse
	err = json.Unmarshal(data, &dataUpdate)
	if err != nil {
		return err
	}

	if dataUpdate.Error {
		return errors.New(dataUpdate.Msj)
	}

	if !dataUpdate.Update {
		return nil
	}

	// Update
	binDir, err := os.Executable()
	if err != nil {
		return err
	}
	dirBin := path.Dir(binDir)
	updateDir := filepath.Join(dirBin, ".update")

	err = os.Mkdir(updateDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	defer os.RemoveAll(updateDir)

	fmt.Printf("Downloading catrina %v...\n", dataUpdate.Version)
	err = downloadFile(filepath.Join(updateDir, "update.zip"), dataUpdate.Url)
	if err != nil {
		return err
	}

	fmt.Println("Extract files...")
	_, err = unzip(filepath.Join(updateDir, "update.zip"), filepath.Join(updateDir, dataUpdate.Version))
	if err != nil {
		return err
	}

	fmt.Println("Check integrity...")
	sum, err := md5Checksum(filepath.Join(updateDir, dataUpdate.Version, "catrina"))
	if err != nil {
		return err
	}

	if dataUpdate.Checksum != sum {
		return errors.New("the binary is corrupt")
	}

	fmt.Printf("installing %v...\n", dataUpdate.Version)

	err = c.Copy(filepath.Join(updateDir, dataUpdate.Version), dirBin, c.Options{
		Skip: func(src string) (bool, error) {
			if src == "catrina-update" {
				return true, nil
			}
			return false, nil
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("catrina %v is installed\n", dataUpdate.Version)
	return nil
}
