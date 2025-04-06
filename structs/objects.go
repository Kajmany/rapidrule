package structs

import "fmt"

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

type Alert struct {
	// Displayables
	Hotkey    rune
	ShortDesc string
	LongDesc  string
	Type      AlertType
}

type AlertType int

const (
	// Many systems ship with rules, worth a mention
	TablesAlready AlertType = iota
	// Need root to get proper resolution from ss or use nft
	NotRoot
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
