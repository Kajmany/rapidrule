// Interface to nft system utility - unlike other execs that could be more portable
// our app should use this interface which is also made to be scriptable
package nft

// TODO: HEY!! this passes the sniff test but it needs thorough testing

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	// This is our file. 'enable' will include it into main
	nftIncludeFile = "rapidrule.nft"
	// This is where the nftables stuff should be in a sensible system
	nftIncludeLine = "include \"/etc/nftables/" + nftIncludeFile + "\""
	// This should be the base nftables file in a sensible system
	nftConfigPath = "/etc/nftables.conf"
)

type NFTMsg struct {
	// TODO: What else matters here? v0v we'll find out
	Changed bool
	Output  string
}

type NFTErr struct {
	Err error
}

// EnsureEnabled makes sure the include line exists in nftables.conf
// the systemd unit is responsible for actually loading this in
func EnsureEnabled() tea.Cmd {
	return func() tea.Msg {
		// Read current config
		content, err := os.ReadFile(nftConfigPath)
		if err != nil {
			return NFTErr{Err: err}
		}

		// Check if include already exists
		if strings.Contains(string(content), nftIncludeLine) {
			return NFTMsg{Output: "include already present", Changed: false}
		}

		// Append include line
		f, err := os.OpenFile(nftConfigPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return NFTErr{Err: err}
		}
		defer f.Close()

		if _, err := f.WriteString("\n" + nftIncludeLine + "\n"); err != nil {
			return NFTErr{Err: err}
		}

		return NFTMsg{Output: "added include to config", Changed: true}
	}
}

// EnsureDisabled removes the include line from nftables.conf
func EnsureDisabled() tea.Cmd {
	return func() tea.Msg {
		// Read current config
		content, err := os.ReadFile(nftConfigPath)
		if err != nil {
			return NFTErr{Err: err}
		}

		lines := strings.Split(string(content), "\n")
		var newLines []string
		found := false

		// Filter out our include line
		// TODO: This'd be better as a regex but I prompted this and it's ok
		for _, line := range lines {
			if strings.TrimSpace(line) != nftIncludeLine {
				newLines = append(newLines, line)
			} else {
				found = true
			}
		}

		if !found {
			return NFTMsg{Output: "include not present", Changed: false}
		}

		// Write back modified config
		if err := os.WriteFile(nftConfigPath, []byte(strings.Join(newLines, "\n")), 0644); err != nil {
			return NFTErr{Err: err}
		}

		return NFTMsg{Output: "removed include from config", Changed: true}
	}
}

// WriteRule writes the ruleset to the include file and loads it
func WriteRule(ruleset string) tea.Cmd {
	return func() tea.Msg {
		// Ensure directory exists
		if err := os.MkdirAll(filepath.Dir(nftIncludeFile), 0755); err != nil {
			return NFTErr{Err: err}
		}

		// Write rules to file
		fullPath := filepath.Join("/etc/nftables", nftIncludeFile)
		if err := os.WriteFile(fullPath, []byte(ruleset), 0644); err != nil {
			return NFTErr{Err: err}
		}

		// Load the rules
		cmd := exec.Command("nft", "-f", fullPath)
		if output, err := cmd.CombinedOutput(); err != nil {
			return NFTErr{Err: fmt.Errorf("%v: %s", err, string(output))}
		}

		return NFTMsg{Output: "rules written and loaded successfully", Changed: true}
	}
}
