// In cmd/init.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pelletier/go-toml" // TOML library
	"github.com/spf13/cobra"
	"github.com/vedangit/portal/cache" // Your cache package
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project for Portal in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, _ := os.Getwd()
		projectName := filepath.Base(currentDir)

		fmt.Printf("✔ I've registered '%s' to this path: %s\n", projectName, currentDir)

		// --- Ask for enter commands ---
		var enterCommands []string
		for {
			var command string
			prompt := &survey.Input{
				Message: "Enter a command for 'portal enter' (or press Enter to finish):",
			}
			survey.AskOne(prompt, &command)

			if strings.TrimSpace(command) == "" {
				break
			}
			enterCommands = append(enterCommands, command)
		}

		// --- Ask for leave commands ---
		var leaveCommands []string
		// (You can add a similar loop for leave commands here if you want)

		// --- Save to .portal.toml ---
		tomlPath := filepath.Join(currentDir, ".portal.toml")
		config := map[string]interface{}{
			"enter": map[string]interface{}{"commands": enterCommands},
			"leave": map[string]interface{}{"commands": leaveCommands},
		}

		tomlBytes, _ := toml.Marshal(config)
		os.WriteFile(tomlPath, tomlBytes, 0644)
		fmt.Printf("✔ Configuration saved to %s!\n", tomlPath)

		// --- Save to global cache ---
		projectsCache, _ := cache.Read()
		projectsCache[projectName] = currentDir
		projectsCache.Write()
		fmt.Println("✔ Project saved to global cache. You can now run 'portal enter", projectName, "' from anywhere!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
