package core

import (
	"encoding/json"
	"fmt"
)

func LogStruct(d interface{}) {
	bytes, _ := json.Marshal(d)
	fmt.Println("$s\n", string(bytes))
}
