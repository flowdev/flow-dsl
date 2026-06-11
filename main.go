package main

import (
	"log"
	"os"

	"github.com/flowdev/flow-dsl/draw"
	"github.com/flowdev/flow-dsl/parser"
)

type linkGenerator struct {
	baseDir string
}

func (lg linkGenerator) GenerateLink(imprt, lastImportPart, typ string, isData bool) (link string, isLinkToFlow bool) {
	if isData {
		return "", false
	}
	if typ == "validateMiwu" {
		// log.Printf("DEBUG: returning: %q + '.md', true", typ)
		return typ + ".md", true
	}
	// log.Printf("DEBUG: imprt: %q, lastImportPart: %q, typ: %q, isData: %t", imprt, lastImportPart, typ, isData)
	file := "create_miwu.go"
	switch typ {
	case "convertCreateMiwuRequestToData":
		return file + "#L23-L26", false
	case "convertMiwuToDb":
		return file + "#L28-L30", false
	case "createMiwuInDb":
		return file + "#L32-L35", false
	case "convertDbMiwuToResponse":
		return file + "#L37-L39", false
	case "checkBasicMiwu":
		return file + "#L41-L44", false
	case "checkMiwuFeaturesExist":
		return file + "#L46-L49", false
	case "checkMiwuSetExists":
		return file + "#L51-L54", false
	}
	return file, false
}

func main() {
	if len(os.Args) < 2 {
		log.Println("FATAL: Flow file needed as input")
		os.Exit(1)
	}
	flowFileName := os.Args[1]
	flowFile, err := parser.ParseFile(flowFileName, &linkGenerator{baseDir: "."})
	if err != nil {
		log.Printf("ERROR: Can't convert to data format: %v", err)
		os.Exit(2)
	}

	drawFlows := ConvertFlowsToDraw(flowFile)
	// flowFileExt := path.Ext(flowFileName)
	// baseFlowFile := flowFileName[0 : len(flowFileName)-len(flowFileExt)]
	// mdFile := baseFlowFile + "-links"

	// flowMode := draw.FlowModeNoLinks
	flowMode := draw.FlowModeMDLinks
	width := 1900
	darkMode := false

	for _, drawFlow := range drawFlows {
		// imdFile := fmt.Sprintf("%s-%s", mdFile, drawFlow.Name)
		imdFile := drawFlow.Name
		drawFlow.ChangeConfig(imdFile, flowMode, width, darkMode)
		svgContents, mdContent, err := drawFlow.Draw()
		if err != nil {
			log.Printf("unexpected draw error: %v", err)
			os.Exit(3)
		}

		for svgFile, svgContent := range svgContents {
			err = os.WriteFile(svgFile, svgContent, 0666)
			if err != nil {
				log.Printf("ERROR: unable to write file %q: %v", svgFile, err)
				os.Exit(4)
			}
		}

		err = os.WriteFile(imdFile+".md", mdContent, 0666)
		if err != nil {
			log.Printf("ERROR: unable to write file %q: %v", imdFile+".md", err)
			os.Exit(5)
		}
	}

}
