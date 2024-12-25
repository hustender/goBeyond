package main

import (
	"fmt"
	"github.com/yusufpapurcu/wmi"
)

type Win32_PhysicalMemory struct {
	Name         string
	Manufacturer string
	PartNumber   string
	SerialNumber string
}

func main() {
	var devices []Win32_PhysicalMemory

	query := "SELECT Name, Manufacturer, PartNumber, SerialNumber FROM Win32_PhysicalMemory"
	res := wmi.Query(query, &devices)
	if res != nil {
		return
	}
	fmt.Println(devices)
}
