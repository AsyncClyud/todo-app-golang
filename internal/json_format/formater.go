package jsonformat

import "encoding/json"

func FormatRawDataIntoJSON(Data any) string {
	json_result, err := json.MarshalIndent(Data, "", " ")
	if err != nil {
		panic(err)
	}

	return string(json_result)
}
