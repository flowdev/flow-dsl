package data

type FlowFile struct {
	Imports []string
	Flows   []*Flow
}

type Flow struct {
	Name            string
	StatementStarts []*StartComp
}

type DataType struct {
	Name string
	Typ  string
	Link string
}

type Arrow struct {
	DataTypes []DataType
	SrcPort   string
	DstPort   string
	SrcComp   *StartComp
	DstComp   *EndComp
}

// Shapes:
// Name     | Start | Mid | End
// ============================
// Arrow    | No    | Yes | No
// ----------------------------
// StartPort| Yes   | No  | No
// EndPort  | No    | No  | Yes
// Loop     | No    | No  | Yes
// Comp     | Yes   | Yes | Yes
// BreakEnd | Yes   | No  | No  // inserted automatically by draw package
// BreakStart No    | No  | Yes // inserted automatically by draw package

type StartComp struct {
	PortName string
	Output   *Arrow
	Comp     *Comp
}

type EndComp struct {
	LoopName string
	LoopPort string
	LoopLink string
	PortName string
	Input    *Arrow
	Comp     *Comp
}

type Comp struct {
	Name         string
	Typ          string
	Link         string
	PluginGroups []PluginGroup
	Inputs       []*Arrow
	Outputs      []*Arrow
}

type PluginGroup struct {
	Interface string
	Plugins   []Plugin
}

type Plugin struct {
	Typ  string
	Link string
}
