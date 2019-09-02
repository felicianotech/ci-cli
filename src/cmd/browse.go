package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Open this project's CI or VCS page in the browser",
	RunE: func(cmd *cobra.Command, args []string) error {

		output, err := exec.Command("git", "remote", "get-url", "origin").Output()
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(output)

		// eventually we need to loop this
		// / provide some logic for multiple remotes
		fetchURL, err := buf.ReadString(' ')
		fetchURL = strings.TrimSpace(fetchURL)
		fetchURL = strings.TrimPrefix(fetchURL, "https://")
		fetchURL = strings.TrimPrefix(fetchURL, "http://")
		fetchURL = strings.TrimSuffix(fetchURL, ".git")

		if strings.HasPrefix(fetchURL, "git@") {
			fetchURL = strings.TrimPrefix(fetchURL, "git@")
			fetchURL = strings.Replace(fetchURL, ":", "/", 1)
		}

		pathParts := strings.Split(fetchURL, "/")

		var openURL string
		vcsMode, err := cmd.Flags().GetBool("vcs")
		if err != nil {
			return err
		}

		if vcsMode {
			openURL = "https://" + fetchURL
		} else {
			openURL = "https://circleci.com/gh/" + pathParts[1] + "/" + pathParts[2]
		}

		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", openURL).Start()
		case "darwin":
			err = exec.Command("open", openURL).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", openURL).Start()
		default:
			err = fmt.Errorf("Unsupported platform.")
		}

		if err != nil {
			fmt.Println("Opening in browser didn't work.")
		} else {
			fmt.Println("Opening in browser.")
		}

		return nil
	},
}

func init() {

	rootCmd.AddCommand(browseCmd)
	browseCmd.Flags().Bool("vcs", false, "Open repository in the VCS website instead of CI provider")
}
