package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/particledecay/semver/pkg/semver"
)

var (
	verbose bool
	outType string
	headers bool
)

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "semver interacts with SemVer-compliant versions",
	Long: `semverdiff allows you to compare two SemVer-compliant versions
		   and perform various actions with the resulting information.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		log.Debug().Msg("Debug messaging turned on")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("You must supply an action to this command")
		}
		return nil
	},
}

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "returns a diff between the first SemVer and second SemVer",
	Long: `compares both provided SemVer-compliant versions and reports on
		   drift between the two.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("You must supply two version strings (%d found)", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		firstVer, err := semver.Parse(args[0])
		if err != nil {
			log.Error().Msgf("'%s' is not a valid SemVer string", args[0])
		}
		secondVer, err := semver.Parse(args[1])
		if err != nil {
			log.Error().Msgf("'%s' is not a valid SemVer string", args[1])
		}
		results := firstVer.Diff(secondVer)
		semver.OutputDiff(results, outType, headers)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display debug messages")
	rootCmd.PersistentFlags().StringVarP(&outType, "output", "o", "table", "output format (table, json)")
	rootCmd.PersistentFlags().BoolVar(&headers, "no-headers", false, "suppress headers from output (table only)")
}

// Execute combines all of the available command functions
func Execute() {
	rootCmd.AddCommand(diffCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msgf("Error during execution: %v", err)
	}
}
