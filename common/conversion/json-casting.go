package conversion

import (
	"encoding/json"

	errors "github.com/GabrielAchumba/Simple-Banking-App/common/errors"
)

func Conversion(source interface{}, destination interface{}) error {
	sourceConverted, err := json.Marshal(source)
	if err != nil {
		return errors.Error("Message: " + err.Error())
	}

	err = json.Unmarshal([]byte(sourceConverted), destination)
	if err != nil {
		return errors.Error("Message: " + err.Error())
	}

	return nil

}
