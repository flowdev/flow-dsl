package main

import (
	"fmt"

	"github.com/flowdev/flow-dsl/data"
	"github.com/flowdev/flow-dsl/draw"
)

func ConvertFlowsToDraw(flows []data.Flow) []*draw.Flow {
	drawFlows := make([]*draw.Flow, len(flows))
	for i := 0; i < len(flows); i++ {
		drawFlows[i] = convertFlowToDraw(flows[i])
	}
	return drawFlows
}

type drawConverter struct {
	registry draw.CompRegistry
}

func convertFlowToDraw(flow data.Flow) *draw.Flow {
	drawFlow := draw.NewFlow(flow.Name, draw.FlowModeNoLinks, 1500, false)
	dc := &drawConverter{registry: drawFlow}

	for _, start := range flow.StatementStarts {
		drawFlow.AddStart(dc.convertStatementStartToDraw(start))
	}
	return drawFlow
}

func (dc *drawConverter) convertStatementStartToDraw(startComp *data.StartComp) draw.StartComp {
	switch {
	case startComp.Comp != nil:
		return dc.convertCompToDraw(startComp.Comp)
	case startComp.PortName != "":
		return dc.convertStartPortToDraw(startComp.PortName, startComp.Output)
	default:
		panic("unsupported type of start component (neiter Component nor StartPort)")
	}
}

func (dc *drawConverter) convertArrowToDraw(arr *data.Arrow) *draw.Arrow {
	drawArr := draw.NewArrow(arr.SrcPort, arr.DstPort)
	for _, dt := range arr.DataTypes {
		drawArr.AddDataType(dt.Name, dt.Typ, dt.Link)
	}
	switch {
	case arr.EndComp.Comp != nil:
		return drawArr.AddDestination(dc.convertCompToDraw(arr.EndComp.Comp))
	case arr.EndComp.PortName != "":
		if arr.DstPort != "" { // arr.DstPort MUST BE ""
			panic(fmt.Sprintf("arrow MUST NOT have a destination port '%s', if it has an end port '%s'", arr.DstPort, arr.EndComp.PortName))
		}
		return drawArr.AddDestination(dc.convertEndPortToDraw(arr.EndComp.PortName))
	case arr.EndComp.LoopName != "":
		return drawArr.AddDestination(dc.convertLoopToDraw(arr.EndComp.LoopName, arr.EndComp.LoopPort, arr.EndComp.LoopLink))
	default:
		panic("unsupported type of arrow destination (neiter Component nor EndPort nor Loop)")
	}
}

func (dc *drawConverter) convertLoopToDraw(name, port, link string) draw.EndComp {
	return draw.NewLoop(name, port, link)
}

func (dc *drawConverter) convertStartPortToDraw(name string, output *data.Arrow) draw.StartComp {
	if output.SrcPort != "" { // output.SrcPort MUST BE ""
		panic(fmt.Sprintf("arrow MUST NOT have a source port '%s', if it has a start port '%s'", output.SrcPort, name))
	}
	return draw.NewStartPort(name).AddOutput(dc.convertArrowToDraw(output))
}

func (dc *drawConverter) convertEndPortToDraw(name string) draw.EndComp {
	return draw.NewEndPort(name)
}

func (dc *drawConverter) convertCompToDraw(comp *data.Comp) *draw.Comp {
	drawComp := draw.NewComp(comp.Name, comp.Typ, comp.Link, dc.registry)
	for _, pg := range comp.PluginGroups {
		drawComp.AddPluginGroup(dc.convertPluginGroupToDraw(pg))
	}
	for _, arr := range comp.Outputs {
		drawComp.AddOutput(dc.convertArrowToDraw(arr))
	}
	return drawComp
}

func (dc *drawConverter) convertPluginGroupToDraw(pg data.PluginGroup) *draw.PluginGroup {
	drawPG := draw.NewPluginGroup(pg.Interface)
	for _, p := range pg.Plugins {
		drawPG.AddPlugin(dc.convertPluginToDraw(p))
	}
	return drawPG
}

func (dc *drawConverter) convertPluginToDraw(p data.Plugin) *draw.Plugin {
	return draw.NewPlugin(p.Typ, p.Link)
}
