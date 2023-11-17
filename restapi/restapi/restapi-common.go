package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func makeResponse(name string, route string) []byte {
	resp := make(map[string]string)
	resp["status"] = "ok"
	resp["message"] = "09-https-server/" + name
	resp["route"] = route
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return jsonResp
}

func JsonPrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func JsonPrettyPrint(j any) {
	var buffer bytes.Buffer

	err := JsonPrettyEncode(j, &buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer.String())
}
