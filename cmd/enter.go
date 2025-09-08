// In cmd/enter.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vedangit/portal/cache" // Your cache package
)

var enterCmd = &cobra.Command{
	Use:   "enter [PROJECT_NAME]",
	Short: "Enter a project's context by its registered name",
	Args:  cobra.ExactArgs(1), // This command now requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Find the project in the cache
		projectsCache, err := cache.Read()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading cache:", err)
			os.Exit(1)
		}

		projectRoot, found := projectsCache[projectName]
		if !found {
			fmt.Fprintf(os.Stderr, "Error: Project '%s' not found. Have you run 'portal init' in its directory yet?\n", projectName)
			os.Exit(1)
		}

		// Load the config from the path we found in the cache
		cfg, err := loadConfig(projectRoot) // This function is in cmd/utils.go
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error loading config in '"+projectRoot+"':", err)
			os.Exit(1)
		}

		originalDir, _ := os.Getwd()
		var scriptBuilder strings.Builder

		scriptBuilder.WriteString(fmt.Sprintf("cd %s\n", projectRoot))
		scriptBuilder.WriteString(fmt.Sprintf("export PORTAL_ACTIVE_PROJECT=%s\n", filepath.Base(projectRoot)))
		scriptBuilder.WriteString(fmt.Sprintf("export PORTAL_PREVIOUS_DIR=%s\n", originalDir))

		for _, command := range cfg.Enter.Commands {
			scriptBuilder.WriteString(command + "\n")
		}

		if cfg.Enter.Message != "" {
			safeMessage := strings.ReplaceAll(cfg.Enter.Message, "'", `'\''`)
			scriptBuilder.WriteString(fmt.Sprintf("echo '🚀 %s'\n", safeMessage))
		}

		fmt.Println(scriptBuilder.String())
	},
}

func init() {
	rootCmd.AddCommand(enterCmd)
}
