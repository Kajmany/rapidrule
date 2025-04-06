// Abstraction over Operating System to 'scrape' current state of
// ports/processes/nftable rules as useful
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
