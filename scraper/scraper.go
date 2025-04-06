// Abstraction over Operating System to 'scrape' current state of
// ports/processes/nftable rules as useful
// These are and should probably only be run once at init (much more model logic is required to 'refresh')
package scraper

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/Kajmany/rapidrule/structs"
	tea "github.com/charmbracelet/bubbletea"
)

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
		// TODO: We can do more, but listening TCP socket only for now
		cmd := exec.Command("ss", "-tlnp")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("problem using ss to get port info")
			return PortScrapeError{Err: err}
		}

		ports, err := parseSSOutput(string(output))
		if err != nil {
			log.Println("problem parsing output of ss")
			return PortScrapeError{Err: err}
		}
		return PortsMsg{Ports: ports}
	}
}

// parseSSOutput parses the output of `ss -plant` into Port structs
func parseSSOutput(output string) ([]structs.Port, error) {
	var ports []structs.Port
	// Example line:
	// LISTEN  0        4096          127.0.0.54:53             0.0.0.0:*      users:(("systemd-resolve",pid=1288,fd=20))
	// Captures: LocalAddr, Port, Process-Name
	re := regexp.MustCompile(`^LISTEN\s+\d+\s+\d+\s+([^:]+):(\d+)\s+.*\"(.*)\"`)
	reader := strings.NewReader(output)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if matches := re.FindStringSubmatch(line); matches != nil {
			// Ensure we have enough matches
			if len(matches) >= 4 {
				portNum, err := strconv.Atoi(matches[2])
				if err != nil {
					return nil, err
				}
				port := structs.Port{
					LocalAddr: matches[1],
					Port:      portNum,
					Process:   matches[3],
				}
				ports = append(ports, port)
			}
		}
	}

	return ports, nil
}

// Sent when we have completed an alert check successfully
// Can be nil if we checked and there's no alert
type AlertMsg struct {
	HasAlert bool
	Alert    structs.Alert
}

// Sent when we error'd attempting to check if we should alert
type AlertError struct {
	Err error
}

func CheckIfRoot() tea.Cmd {
	return func() tea.Msg {
		// TODO: More portable way to do this?
		cmd := exec.Command("id", "-u")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return AlertError{Err: fmt.Errorf("failed to check user id: %v", err)}
		}

		uid, err := strconv.Atoi(strings.TrimSpace(string(output)))
		if err != nil {
			return AlertError{Err: fmt.Errorf("failed to parse user id: %v", err)}
		}

		if uid != 0 {
			return AlertMsg{
				HasAlert: true,
				Alert: structs.Alert{
					ShortDesc: "Running as non-root user",
					LongDesc:  "This program will not function without root access!",
				},
			}
		}
		return AlertMsg{HasAlert: false} // No alert if root
	}
}

func CheckTables() tea.Cmd {
	return func() tea.Msg {
		// TODO: Should this logic be in nft driver?
		cmd := exec.Command("nft", "list", "tables")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return AlertError{Err: fmt.Errorf("failed to list nftables: %v", err)}
		}

		tables := strings.TrimSpace(string(output))
		if tables != "" {
			return AlertMsg{
				HasAlert: true,
				Alert: structs.Alert{
					ShortDesc: "Existing nftables rules detected",
					LongDesc:  "System already has nftables rules which may conflict",
				},
			}
		}
		return AlertMsg{HasAlert: false} // No alert if no tables
	}
}
