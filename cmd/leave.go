// In cmd/leave.go
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// Make sure this matches your module path
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Leave the current project's context",
	Run: func(cmd *cobra.Command, args []string) {
		previousDir := os.Getenv("PORTAL_PREVIOUS_DIR")
		activeProjectName := os.Getenv("PORTAL_ACTIVE_PROJECT")

		if activeProjectName == "" || previousDir == "" {
			fmt.Fprintln(os.Stderr, "Error: You are not in an active Portal session.")
			os.Exit(1)
		}

		// --- NEW LOGIC TO FIND AND READ CONFIG ---
		// Read global config to find where all projects live
		homeDir, _ := os.UserHomeDir()
		viper.AddConfigPath(filepath.Join(homeDir, ".config", "portal"))
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		_ = viper.ReadInConfig() // Ignore error if it doesn't exist yet
		projectsDir := viper.GetString("projects_dir")
		if strings.HasPrefix(projectsDir, "~/") {
			projectsDir = filepath.Join(homeDir, projectsDir[2:])
		}

		projectRoot := filepath.Join(projectsDir, activeProjectName)
		cfg, err := loadConfig(projectRoot) // Re-use the loadConfig function from enter.go
		// --- END NEW LOGIC ---

		var scriptBuilder strings.Builder

		// Add commands from the [leave] section of the TOML file
		if err == nil && cfg.Leave.Commands != nil {
			for _, command := range cfg.Leave.Commands {
				scriptBuilder.WriteString(command + "\n")
			}
		}

		// Add the standard cleanup commands
		scriptBuilder.WriteString(fmt.Sprintf("cd %s\n", previousDir))
		scriptBuilder.WriteString("unset PORTAL_ACTIVE_PROJECT\n")
		scriptBuilder.WriteString("unset PORTAL_PREVIOUS_DIR\n")
		scriptBuilder.WriteString("echo 'Left project context.'\n")

		fmt.Println(scriptBuilder.String())
	},
}

func init() {
	rootCmd.AddCommand(leaveCmd)
}
