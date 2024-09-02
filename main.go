package main

import (
	"fmt"
	"github.com/songgao/water"
)

// NewConfig 创建Config
func NewConfig(DeviceType int) water.Config {
	config := water.Config{
		DeviceType: water.DeviceType(DeviceType),
		PlatformSpecificParams: water.PlatformSpecificParams{
			InterfaceName: "tap0",
		},
	}
	return config
}

func main() {
	config := NewConfig(water.TAP)

	ifce, err := water.New(config)

	if err != nil {
		fmt.Println("创建TAP异常,", err)
		return
	}

}
