package models

// TransactionTemplate comment
type TransactionTemplate struct {
	Inputs   []Input `json:"input"`
	Output   Output  `json:"output"`
	Locktime int     `json:"locktime,omitonempty"`
}

// Input comment
type Input struct {
	Txid     string `json:"txid"`
	Vout     int    `json:"vout"`
	Sequence int    `json:"sequence,omitonempty"`
}

// Output comment
type Output struct {
	Data string `json:"data,omitonempty"`
}

// AddOutput comment
func (tt *TransactionTemplate) AddOutput(output Output) {
	tt.Output = output
}

// AddInput comment
func (tt *TransactionTemplate) AddInput(input Input) {
	tt.Inputs = append(tt.Inputs, input)
}

// IsValid checks if the transaction template is valid
func (tt TransactionTemplate) IsValid() (valid bool) {
	if len(tt.Inputs) < 1 {
		return false
	}
	return true
}
