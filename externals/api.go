package externals

import (
	"fmt"
	"hp-api/configs"
)

func ConnectApi() error {
	err := configs.Load()
	if err != nil {
		fmt.Println("Error to load the config:", err)
		return err
	}

	return nil
}
