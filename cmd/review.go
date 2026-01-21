/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reviewCmd represents the review command
var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Run a code review on a specified branch",
	Long: `Run a code review on a specified branch against a target commit.
This command analyzes the changes in the review branch and provides feedback.`,
	Run: func(cmd *cobra.Command, args []string) {
		folder, _ := cmd.Flags().GetString("folder")
		reviewBranch, _ := cmd.Flags().GetString("review-branch")
		targetCommit, _ := cmd.Flags().GetString("target-commit")

		fmt.Printf("Review called with:\n")
		fmt.Printf("  Folder: %s\n", folder)
		fmt.Printf("  Review Branch: %s\n", reviewBranch)
		fmt.Printf("  Target Commit: %s\n", targetCommit)
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)

	// Define flags for the review command
	reviewCmd.Flags().StringP("folder", "f", ".", "Working folder to run the review in")
	reviewCmd.Flags().StringP("review-branch", "r", "", "Branch to run code review on")
	reviewCmd.Flags().StringP("target-commit", "t", "", "Target commit for merge comparison")

	// Mark required flags
	reviewCmd.MarkFlagRequired("review-branch")
	reviewCmd.MarkFlagRequired("target-commit")
}
