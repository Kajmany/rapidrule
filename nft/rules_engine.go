// Responsible for generating rules. Eventual string output should reult in a single well-formed Table
package nft

import (
	"fmt"
	"log"
	"strings"
	"text/template"
)

// We can't really do proper individual verification of these in here
// but these IDEALLY represent well-formed NFTable elements
type (
	table string
	chain struct {
		Name  string
		Type  string
		Rules []Rule
	}
	Rule string
)

const NFTTemplate = `table inet rapdidrule {
  chain {{.Name}} {
    {{.Type}}
    {{range .Rules -}}
    {{.}}
    {{end -}}
  }
}`

//func Rules(m Model) {
//	nftCommands := "nft flush ruleset\n nft add table inet filter\n nft add chain inet filter input\n nft add rule inet filter input counter drop"
//	for _, port := range m.Ports {
//		nftCommands += "\n nft add rule inet filter input tcp dport " + strconv.Itoa(port.Port) + " accept;"
//	}
//	log.Printf(nftCommands)
//}

func GenTable(chains []chain) table {
	var sb strings.Builder
	tmpl := template.Must(template.New("nft").Parse(NFTTemplate))

	for _, c := range chains {
		err := tmpl.Execute(&sb, c)
		if err != nil {
			// TODO: Msg or something
			panic(fmt.Sprintf("NFTTemplate execution failed: %v", err))
		}
	}
	output := sb.String()
	log.Printf("Parsed table from template:\n    %s", output)
	return table(output)
}

const ChainTemplate = `chain {{.Name}} {
    {{.Type}}
    {{range .Rules -}}
    {{.}}
    {{end -}}
}`

// TODO: This is kind of redundant panic code w.r.t what template takes
func GenChain(name string, chainType string, rules []Rule) chain {
	var sb strings.Builder
	tmpl := template.Must(template.New("chain").Parse(ChainTemplate))

	err := tmpl.Execute(&sb, chain{
		Name:  name,
		Type:  chainType,
		Rules: rules,
	})
	if err != nil {
		// TODO: Msg or something
		panic(fmt.Sprintf("ChainTemplate execution failed: %v", err))
	}

	// Log the generated chain for debugging
	generated := sb.String()
	log.Printf("Generated chain:\n%s", generated)

	return chain{
		Name:  name,
		Type:  chainType,
		Rules: rules,
	}
}
