package sui

import (
	"fmt"
	"log"
	"os/exec"
)

func Sui() {
	sui_send := `AGGREGATOR=https://aggregator.walrus-testnet.walrus.space
PUBLISHER=https://publisher.walrus-testnet.walrus.space

curl -X PUT "$PUBLISHER/v1/store" -d "some string" # store the string 'some string' for 1 storage epoch
curl -X PUT "$PUBLISHER/v1/store?epochs=5" --upload-file "some/file" # store file 'some/file' for 5 storage epochs`

	cmd := exec.Command("bash", "-c", sui_send)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}
