package structs

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type Port struct {
	// Displayables
	Hotkey    rune   // May need to change
	LocalAddr string // Can be text like loopback
	Port      int
	Process   string // Not PID, just name (add pid later if we need it)
	// Other state
	LLMRes    Judgement // Machine-readable judgement
	LLMStatus JudgementProgress
	Eval      *Eval
}

func (p Port) String() string {
	return fmt.Sprintf("Port: %s:%d held by %s", p.LocalAddr, p.Port, p.Process)
}
func (p Port) ToPrompt() string {
	return fmt.Sprintf("%s:%d %s,", p.LocalAddr, p.Port, p.Process)
}
func (p Port) ToRow() table.Row {
	return table.Row{strconv.Itoa(p.Port), p.LocalAddr, p.Process}
}

type Eval struct {
	Concerns    string `json:"concerns"`
	Investigate string `json:"investigate"`
	Confidence  string `json:"confidence"`
	Port        int    `json:"port"`
}

func (e Eval) String() string {
	return fmt.Sprintf("concerns: %s,\n investigate: %s,\n confidence: %s,\n port: %d,",
		e.Concerns, e.Investigate, e.Confidence, e.Port)
}

type TotalEval struct {
	Overall    string `json:"overall"`
	Alert      string `json:"alert"`
	AlertShort string `json:"alert_short"`
	AlertLong  string `json:"alert_long"`
}

func (te TotalEval) String() string {
	return fmt.Sprintf("overall: %s,\n alert: %s,\n alert_short: %s,\nalert_long: %s",
		te.Overall, te.Alert, te.AlertShort, te.AlertLong)
}

type Alert struct {
	// Displayables
	Hotkey    rune
	ShortDesc string
	LongDesc  string
	Type      AlertType
}

func NewAlert(alert_short string, alert_long string, alert_type AlertType) Alert {
	var alert Alert
	alert.ShortDesc = alert_short
	alert.LongDesc = alert_long
	alert.Type = AlertType(alert_type)
	alert.Hotkey = 'A'
	return alert
}

type AlertType int

const (
	// Many systems ship with rules, worth a mention
	TablesAlready AlertType = iota
	// Need root to get proper resolution from ss or use nft
	NotRoot
	Red
	Yellow
)

func (a Alert) String() string {
	// TODO: repr type if useful need a case
	return fmt.Sprintf("Alert: %s\n    %s", a.ShortDesc, a.LongDesc)
}

// Should be made from structued LLM output
type Judgement int

const (
	Good Judgement = iota
	Attention
	Bad
)

// Need to know where we're at with LLM requests
type JudgementProgress int

const (
	Unsent JudgementProgress = iota
	Inflight
	Done
)
