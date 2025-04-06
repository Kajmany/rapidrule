package structs

type Port struct {
  // Displayables
  Hotkey rune, // May need to change
  LocalAddr string, // Can be text like loopback
  Port int,
  Process string, // Not PID, just name (add pid later if we need it)
  // Notice no shortdesc
  LongDesc string, // Our own templated string
  LLMDescription string, // Multiline human-readable elaboration
  // Other state
  LLMRes Judgement, // Machine-readable judgement
  LLMStatus JudgementProgress,

}

type Alert struct {
  // Displayables
  Hotkey rune,
  ShortDesc string,
  LongDesc string,
  LLMDescription string, // Multiline human-readable elaboration
  // Other state
  LLMRes Judgement, // Machine-readable judgement
  LLMStatus JudgementProgress,
}

// Should be made from structued LLM output
const Judgement (
  Good = iota
  Attention
  Bad
)

// Need to know where we're at with LLM requests
const JudgementProgress (
  Unsent = iota
  Inflight
  Done
)
