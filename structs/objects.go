package structs

import "fmt"

type Port struct {
	// Displayables
	Hotkey    rune   // May need to change
	LocalAddr string // Can be text like loopback
	Port      int
	Process   string // Not PID, just name (add pid later if we need it)
	// Notice no shortdesc
	LongDesc       string // Our own templated string
	LLMDescription string // Multiline human-readable elaboration
	// Other state
	LLMRes    Judgement // Machine-readable judgement
	LLMStatus JudgementProgress
}

func (p Port) String() string {
	return fmt.Sprintf("Port: %s:%d held by %s", p.LocalAddr, p.Port, p.Process)
}

type Alert struct {
	// Displayables
	Hotkey         rune
	ShortDesc      string
	LongDesc       string
	LLMDescription string // Multiline human-readable elaboration
	// Other state
	LLMRes    Judgement // Machine-readable judgement
	LLMStatus JudgementProgress
}

func (a Alert) String() string {
	return fmt.Sprintf("Alert: %s", a.ShortDesc)
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
