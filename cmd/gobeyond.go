//go:build windows

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yusufpapurcu/wmi"
)

// Declare the components
var cpu []Win32_Processor
var memory []Win32_PhysicalMemory
var gpu []Win32_VideoController

var allCommand = &cobra.Command{
	Use:     "all",
	Aliases: []string{"*"},
	Short:   "Shows the information about all the components.",
	Long:    "Shows the information about all the components.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, c := range cpu {
			c.Parse()
		}
		for _, m := range memory {
			m.Parse()
		}
		for _, g := range gpu {
			g.Parse()
		}
	},
}

var cpuCommand = &cobra.Command{
	Use:     "cpu",
	Aliases: []string{"processor"},
	Short:   "Shows the information about the CPU.",
	Long:    "Shows the information about the CPU.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, c := range cpu {
			c.Parse()
		}
	},
}

var memoryCommand = &cobra.Command{
	Use:     "memory",
	Aliases: []string{"ram"},
	Short:   "Shows the information about the memory.",
	Long:    "Shows the information about the memory.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, m := range memory {
			m.Parse()
		}
	},
}

var gpuCommand = &cobra.Command{
	Use:     "gpu",
	Aliases: []string{"graphicscard"},
	Short:   "Shows the information about the GPU.",
	Long:    "Shows the information about the GPU.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, g := range gpu {
			g.Parse()
		}
	},
}

func init() {
	// Load the information into memory
	LoadSpecs()

	rootCmd.AddCommand(allCommand)
	rootCmd.AddCommand(cpuCommand)
	rootCmd.AddCommand(memoryCommand)
	rootCmd.AddCommand(gpuCommand)
}

// loadInfoFor for all given components
func LoadSpecs() {
	loadInfoFor("cpu", &cpu)
	loadInfoFor("memory", &memory)
	loadInfoFor("gpu", &gpu)
}

// Loads the information for the given device into memory
func loadInfoFor(name string, dst interface{}) {
	query := getQueryFor(dst)
	res := wmi.Query(query, dst)
	if res != nil {
		fmt.Println("An error occured while trying to load information for " + name + ": \n" + res.Error())
	}
}

// Returns the query for the given component
func getQueryFor(component interface{}) string {
	switch component.(type) {
	case *[]Win32_Processor:
		return "SELECT Name, Manufacturer, CurrentClockSpeed, ProcessorId FROM Win32_Processor"
	case *[]Win32_PhysicalMemory:
		return "SELECT Name, Manufacturer, PartNumber, SerialNumber, Speed, DeviceLocator FROM Win32_PhysicalMemory"
	case *[]Win32_VideoController:
		return "SELECT Name, AdapterCompatibility, AdapterRAM, DriverVersion, CurrentHorizontalResolution, CurrentVerticalResolution, CurrentRefreshRate FROM Win32_VideoController"
	}
	return ""
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
	fmt.Printf("Speed: %d T/s\n", m.Speed)
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
	fmt.Printf("Refresh Rate: %d Hz\n", g.CurrentRefreshRate)
	fmt.Println()
}
