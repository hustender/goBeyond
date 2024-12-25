package main

import (
	"fmt"
	"github.com/yusufpapurcu/wmi"
)

var queryDict = map[string]string{
	"cpu":    "SELECT Name, Manufacturer, CurrentClockSpeed, ProcessorId FROM Win32_Processor",
	"memory": "SELECT Name, Manufacturer, PartNumber, SerialNumber, Speed, DeviceLocator FROM Win32_PhysicalMemory",
	"gpu":    "SELECT Name, AdapterCompatibility, AdapterRAM, DriverVersion, CurrentHorizontalResolution, CurrentVerticalResolution, CurrentRefreshRate FROM Win32_VideoController",
}

func main() {
	var memory []Win32_PhysicalMemory
	var cpu []Win32_Processor
	var gpu []Win32_VideoController

	if getInfo("cpu", &cpu) {
		for _, c := range cpu {
			c.Parse()
		}
	}
	if getInfo("memory", &memory) {
		for _, m := range memory {
			m.Parse()
		}
	}
	if getInfo("gpu", &gpu) {
		for _, g := range gpu {
			g.Parse()
		}
	}
}

func getInfo(name string, dst interface{}) bool {
	query := queryDict[name]
	res := wmi.Query(query, dst)
	if res != nil {
		fmt.Println("An error occured while trying to read information for " + name + ": \n" + res.Error())
		return false
	}
	return true
}

// CPU
type Win32_Processor struct {
	Name              string // 12th Gen Intel(R) Core(TM) i3-12100F
	Manufacturer      string // Manufacturer
	CurrentClockSpeed int    // Base clock speed (MHz)
	ProcessorId       string // Model
}

func (c Win32_Processor) Parse() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Manufacturer: %s\n", c.Manufacturer)
	fmt.Printf("Base clock speed: %d (MHz)\n", c.CurrentClockSpeed)
	fmt.Printf("Model: %s\n", c.ProcessorId)
	fmt.Println()
}

// Memory
type Win32_PhysicalMemory struct {
	Name          string // Name
	Manufacturer  string // Manufacturer
	Speed         int    // Speed (T/s)
	PartNumber    string // Model
	SerialNumber  string // Serial Number
	DeviceLocator string // Channel
}

func (m Win32_PhysicalMemory) Parse() {
	fmt.Printf("Name: %s\n", m.Name)
	fmt.Printf("Manufacturer: %s\n", m.Manufacturer)
	fmt.Printf("Speed: %d (T/s)\n", m.Speed)
	fmt.Printf("Model: %s\n", m.PartNumber)
	fmt.Printf("Serial Number: %s\n", m.SerialNumber)
	fmt.Printf("Channel: %s\n", m.DeviceLocator)
	fmt.Println()
}

// Graphics card
type Win32_VideoController struct {
	Name                        string // Name
	AdapterCompatibility        string // Manufacturer
	AdapterRAM                  uint64 // Video Memory (bytes)
	DriverVersion               string // Driver Version
	CurrentHorizontalResolution uint32 // Horizontal Resolution
	CurrentVerticalResolution   uint32 // Vertical Resolution
	CurrentRefreshRate          uint32 // Refresh Rate (Hz)
}

func (g Win32_VideoController) Parse() {
	fmt.Printf("Name: %s\n", g.Name)
	fmt.Printf("Manufacturer: %s\n", g.AdapterCompatibility)
	fmt.Printf("Video Memory: %d (bytes)\n", g.AdapterRAM)
	fmt.Printf("Driver Version: %s\n", g.DriverVersion)
	fmt.Printf("Horizontal Resolution: %d\n", g.CurrentHorizontalResolution)
	fmt.Printf("Vertical Resolution: %d\n", g.CurrentVerticalResolution)
	fmt.Printf("Refresh Rate: %d (Hz)\n", g.CurrentRefreshRate)
	fmt.Println()
}
