package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// ============= TARGET INTERFACE =============

type JSONDataProcessor interface {
	ProcessJSON(data string) error
}

// ============= ADAPTEE (Legacy System) =============

type LegacyXMLProcessor struct {
	version string
}

func NewLegacyXMLProcessor() *LegacyXMLProcessor {
	return &LegacyXMLProcessor{version: "1.0"}
}

func (l *LegacyXMLProcessor) HandleXMLData(xmlData string) error {
	fmt.Printf("Processing XML (v%s): %s\n", l.version, xmlData)
	return nil
}

// ============= ADAPTER =============

type XMLToJSONAdapter struct {
	xmlProcessor *LegacyXMLProcessor
}

func NewXMLToJSONAdapter(processor *LegacyXMLProcessor) *XMLToJSONAdapter {
	return &XMLToJSONAdapter{xmlProcessor: processor}
}

func (a *XMLToJSONAdapter) ProcessJSON(jsonData string) error {
	// Convert JSON to XML
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return err
	}

	xmlData, err := a.convertToXML(data)
	if err != nil {
		return err
	}

	return a.xmlProcessor.HandleXMLData(xmlData)
}

func (a *XMLToJSONAdapter) convertToXML(data map[string]interface{}) (string, error) {
	type XMLData struct {
		XMLName xml.Name
		Data    map[string]interface{} `xml:",any"`
	}

	xmlObj := XMLData{
		XMLName: xml.Name{Local: "data"},
		Data:    data,
	}

	xmlBytes, err := xml.Marshal(xmlObj)
	if err != nil {
		return "", err
	}

	return string(xmlBytes), nil
}

// ============= FUNCTIONAL ADAPTER (Go Style) =============

// Adapter using function composition
func AdaptXMLToJSON(xmlHandler func(string) error) JSONDataProcessor {
	return &funcAdapter{
		processFn: func(jsonData string) error {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
				return err
			}

			// Simple conversion for demo
			xmlData := fmt.Sprintf("<data>%s</data>", jsonData)
			return xmlHandler(xmlData)
		},
	}
}

type funcAdapter struct {
	processFn func(string) error
}

func (f *funcAdapter) ProcessJSON(data string) error {
	return f.processFn(data)
}

func main() {
	// Traditional adapter
	legacyProcessor := NewLegacyXMLProcessor()
	adapter := NewXMLToJSONAdapter(legacyProcessor)

	jsonData := `{"user":"john","age":30}`
	adapter.ProcessJSON(jsonData)

	fmt.Println("---")

	// Functional adapter
	xmlHandler := func(xmlData string) error {
		fmt.Printf("Handling XML: %s\n", xmlData)
		return nil
	}

	funcAdapter := AdaptXMLToJSON(xmlHandler)
	funcAdapter.ProcessJSON(jsonData)
}
