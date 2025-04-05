package scraper

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/charmbracelet/bubbletea"
)

// Abstraction over Operating System to 'scrape' current state of
// ports/processes/nftable rules as useful

// PortsMsg is sent when port scraping completes successfully
type PortsMsg struct {
	Ports []structs.Port
}

// PortScrapeError is sent when port scraping fails
type PortScrapeError struct {
	Err error
}

func (e PortScrapeError) Error() string {
	return fmt.Sprintf("port scrape error: %v", e.Err)
}

func GetPorts() tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("ss", "-tulnp")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return PortScrapeError{Err: err}
		}

		ports, err := parseSSOutput(string(output))
		if err != nil {
			return PortScrapeError{Err: err}
		}
		return PortsMsg{Ports: ports}
	}
}

// parseSSOutput parses the output of `ss -plant` into Port structs
func parseSSOutput(output string) ([]structs.Port, error) {
	var ports []structs.Port
	outputs
	// TODO: Implement proper regex parsing
	re := regexp.MustCompile(`.+?\s+.+?\s+.+?\s+.+?\s+(.+?):(\d+)\s+.+?\s+users:\(\(\"(.+?)\"`)
	matches := re.FindAllStringSubmatch(output, -1)

	for _, match := range matches {
		if len(match) < 4 {
			continue
		}
		ports = append(ports, structs.Port{
			LocalAddr: match[1],
			Port:      match[2],
			Process:   match[3],
		})
	}
	return ports, nil
}
