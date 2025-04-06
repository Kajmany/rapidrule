package tea

import (
	"log"
	"strconv"

)

func Rules(m Model) {
	nftCommands := "nft flush ruleset\n nft add table inet filter\n nft add chain inet filter input\n nft add rule inet filter input counter drop"
		for _, port:= range m.Ports {
			nftCommands+="\n nft add rule inet filter input tcp dport "+strconv.Itoa(port.Port)+" accept;"
		}
	log.Printf(nftCommands)
	
}

