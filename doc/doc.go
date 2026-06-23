package doc

import (
	// "fmt"
	"errors"
	"fmt"
	"os"

	"github.com/flowdev/flow-dsl/draw"
	"github.com/flowdev/flow-dsl/parser"
)

type linkGenerator struct {
	baseDir string
}

func (lg linkGenerator) GenerateLink(imprt, typ string, isData bool) (link string, isLinkToFlow bool) {
	if isData {
		if typ == "data.Data" {
			return "data/data.go#L3-L8", true
		}
		return "", false
	}
	switch typ {
	case "validateMIWU":
		// log.Printf("DEBUG: returning: %q + '.md', true", typ)
		return typ + ".md", true
	case "checkBasicMIWU":
		// log.Printf("DEBUG: returning: %q + '.md', true", typ)
		return typ + ".md", true
	}
	return "", false
}

func DocumentFlows(flowFiles []string) error {
	errs := make([]error, 0, len(flowFiles))
	for _, flowFile := range flowFiles {
		if err := documentFlow(flowFile); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func documentFlow(flowFileName string) error {
	flowFile, err := parser.ParseFile(flowFileName, &linkGenerator{baseDir: "."})
	if err != nil {
		fmt.Errorf("can't convert to data format: %w", err)
		return err
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
		// imdFile := fmt.Sprintf("%s-no-links", drawFlow.Name)
		imdFile := drawFlow.Name
		drawFlow.ChangeConfig(imdFile, flowMode, width, darkMode)
		svgContents, mdContent, err := drawFlow.Draw()
		if err != nil {
			fmt.Errorf("unexpected draw error: %w", err)
			return err
		}

		for svgFile, svgContent := range svgContents {
			err = os.WriteFile(svgFile, svgContent, 0666)
			if err != nil {
				fmt.Errorf("unable to write file %q: %w", svgFile, err)
				return err
			}
		}

		err = os.WriteFile(imdFile+".md", mdContent, 0666)
		if err != nil {
			fmt.Errorf("unable to write file %q: %w", imdFile+".md", err)
			return err
		}
	}
	return nil

}
