package urlp

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func CleanFilterParam(allowedKey []string, filterParamMap map[string]interface{}) error {
	if len(allowedKey) < 0 {
		return fmt.Errorf("allowed key is blank")
	}

	if len(filterParamMap) < 0 {
		return fmt.Errorf("filter param is null")
	}

	for k, _ := range filterParamMap {
		for _, aK := range allowedKey {
			if k != aK {
				return fmt.Errorf("some filter param is not allowed, %s", k)
			}
		}
	}

	return nil
}

func DecodeEncodedString(encodedFilterString string, v interface{}) (err error) {
	encodedFilterString = strings.ReplaceAll(encodedFilterString, ".", "+")
	encodedFilterString = strings.ReplaceAll(encodedFilterString, "_", "/")
	encodedFilterString = strings.ReplaceAll(encodedFilterString, "-", "=")

	decodedByte, err := base64.StdEncoding.DecodeString(encodedFilterString)
	if err != nil {
		return err
	}

	err = json.Unmarshal(decodedByte, &v)
	if err != nil {
		return err
	}

	return nil
}
