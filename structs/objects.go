package structs

type Port struct {
  // Displayables
  Hotkey rune, // May need to change
  LocalAddr string, // Can be text like loopback
  Port int,
  Process string, // Not PID, just name (add pid later if we need it)
  // Other state
  LLMDescription string, // Multiline human-readable elaboration
  LLMRes Judgement, // Machine-readable judgement

}

type Alert struct {
  // Displayables
  Hotkey rune,
  ShortDesc string,
  LongDesc string,
  LLMDescription string, // Multiline human-readable elaboration
  // Other state
  LLMRes Judgement, // Machine-readable judgement
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
