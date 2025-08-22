package sdk

import "encoding/json"

func ConvertType(input any, output any) error {
	in, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(in, output)
}
