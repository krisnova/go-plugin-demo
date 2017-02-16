package main

import (
	"plugin"
	"fmt"
	"os"
)

// PluginNumber defines a unique Go plugin to run our program with.
// Here we define which plugin we want to use with our main program.
// Simply switch the PluginNumber to change the plugin we will be using.
const PluginNumber = "1"

// PluginName is just a path to the plugin we will use. Note that the name
// will be string formatted immediately in the main() function.
var PluginName = "plugins/plugin%s.so"

// main is a basic function to demonstrate how Go plugins work.
func main() {

	// Here we demonstrate how we can switch plugins at runtime
	// without having to recompile our binary! How cool!
	num := os.Getenv("PLUGIN_NUMBER")
	if num == "" {
		PluginName = fmt.Sprintf(PluginName, PluginNumber)
	} else {
		PluginName = fmt.Sprintf(PluginName, num)
	}

	// Basic wrapper for the demonstration.
	fmt.Printf("(main)    Running with plugin name: %s\n", PluginName)
	err := runPlugins()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// runPlugins in procedural logic for demonstrating, and manipulating Go plugins at runtime.
func runPlugins() error {

	// Example of running a plugin with no arguments
	Run, err := getSymbol(PluginName, "Run")
	if err != nil {
		return err
	}
	Run.(func())()

	// Example of getting a variable from a plugin
	I, err := getSymbol(PluginName, "I")
	if err != nil {
		return err
	}

	// Now manipulate the variable
	i := *I.(*int)
	fmt.Printf("(main)    Manipulating a variable (%d + 1)\n", i)
	i = i + 1
	*I.(*int) = i

	// Reference the new variable from the plugin
	PrintI, err := getSymbol(PluginName, "PrintI")
	if err != nil {
		return err
	}
	PrintI.(func())()
	return nil
}

// getSymbol is a convenience wrapper for pulling a symbol out of a Go plugin.
func getSymbol(so, name string) (plugin.Symbol, error) {

	// Get a pointer to a plugin type
	// *Plugin: https://tip.golang.org/pkg/plugin/#Plugin
	pluginPointer, err := plugin.Open(so)
	if err != nil {
		return nil, err
	}

	// Get the symbol for RunPlugin
	// Symbol: https://tip.golang.org/src/plugin/plugin.go?s=1914:1937#L60
	symbol, err := pluginPointer.Lookup(name)
	if err != nil {
		return nil, err
	}
	return symbol, nil

}
