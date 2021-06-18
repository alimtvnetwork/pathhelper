package pathinsfmt

type SymbolicLink struct {
	Src           string `json:"Src"`
	Dst           string `json:"Dst"`
	IsForce       bool   `json:"IsForce"`
	IsSkipOnError bool   `json:"IsSkipOnError"`
}
