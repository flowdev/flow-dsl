package doc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"

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

func DocumentFlows(flowFiles []string, links bool, width int, dark bool, docFileNameTpl *template.Template) error {
	fmt.Printf("links: %t, width: %d, dark: %t\n", links, width, dark)
	errs := make([]error, 0, len(flowFiles))
	for _, flowFile := range flowFiles {
		if err := documentFlow(flowFile, links, width, dark, docFileNameTpl); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func documentFlow(flowFileName string, links bool, width int, dark bool, docFileNameTpl *template.Template) error {
	flowFile, err := parser.ParseFile(flowFileName, &linkGenerator{baseDir: "."})
	if err != nil {
		return fmt.Errorf("can't convert to data format: %w", err)
	}

	drawFlows := ConvertFlowsToDraw(flowFile)
	// flowFileExt := path.Ext(flowFileName)
	// baseFlowFile := flowFileName[0 : len(flowFileName)-len(flowFileExt)]
	// mdFile := baseFlowFile + "-links"

	flowMode := draw.FlowModeNoLinks
	if links {
		flowMode = draw.FlowModeMDLinks
	}

	for _, drawFlow := range drawFlows {
		imdFile := drawFlow.Name
		if docFileNameTpl != nil {
			buf := &bytes.Buffer{}
			if err := docFileNameTpl.Execute(buf, drawFlow); err != nil {
				return fmt.Errorf("unable to generate documentation file name: %w", err)
			}
			imdFile = buf.String()
		}
		drawFlow.ChangeConfig(imdFile, flowMode, width, dark)
		svgContents, mdContent, err := drawFlow.Draw()
		if err != nil {
			return fmt.Errorf("unexpected draw error: %w", err)
		}

		for svgFile, svgContent := range svgContents {
			err = os.WriteFile(svgFile, svgContent, 0666)
			if err != nil {
				return fmt.Errorf("unable to write file %q: %w", svgFile, err)
			}
		}

		err = os.WriteFile(imdFile+".md", mdContent, 0666)
		if err != nil {
			return fmt.Errorf("unable to write file %q: %w", imdFile+".md", err)
		}
	}
	return nil

}
