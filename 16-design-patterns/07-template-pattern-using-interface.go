package main

import (
	"fmt"
	"strings"
)

// ============= TRADITIONAL TEMPLATE METHOD (Limited in Go) =============
// Go doesn't have inheritance, so we use composition + interfaces

type DataPipeline interface {
	ExtractData() []Record
	TransformData([]Record) []Record
	LoadData([]Record) error
	SendNotification() // Hook method (optional)
}

type Record struct {
	ID    string
	Data  map[string]interface{}
	Valid bool
}

// BasePipeline provides common template method
type BasePipeline struct {
	pipeline DataPipeline
}

func NewBasePipeline(pipeline DataPipeline) *BasePipeline {
	return &BasePipeline{pipeline: pipeline}
}

// Template method
func (b *BasePipeline) Execute() error {
	fmt.Println("Starting pipeline execution...")

	// Step 1: Extract
	data := b.pipeline.ExtractData()
	fmt.Printf("Extracted %d records\n", len(data))

	// Step 2: Transform
	data = b.pipeline.TransformData(data)
	fmt.Printf("Transformed to %d valid records\n", len(data))

	// Step 3: Load
	if err := b.pipeline.LoadData(data); err != nil {
		return err
	}

	// Step 4: Notify (hook)
	b.pipeline.SendNotification()

	fmt.Println("Pipeline completed successfully")
	return nil
}

// ============= CONCRETE IMPLEMENTATION =============

type CSVToPostgresPipeline struct {
	csvFile   string
	dbConnStr string
	tableName string
}

func NewCSVToPostgresPipeline(csvFile, dbConnStr, tableName string) *CSVToPostgresPipeline {
	return &CSVToPostgresPipeline{
		csvFile:   csvFile,
		dbConnStr: dbConnStr,
		tableName: tableName,
	}
}

func (c *CSVToPostgresPipeline) ExtractData() []Record {
	fmt.Printf("Reading CSV from %s\n", c.csvFile)
	// Simulate CSV reading
	return []Record{
		{ID: "1", Data: map[string]interface{}{"name": "John", "age": 30}, Valid: true},
		{ID: "2", Data: map[string]interface{}{"name": "", "age": -5}, Valid: false},
		{ID: "3", Data: map[string]interface{}{"name": "Jane", "age": 25}, Valid: true},
	}
}

func (c *CSVToPostgresPipeline) TransformData(data []Record) []Record {
	fmt.Println("Transforming and validating data...")
	var valid []Record
	for _, record := range data {
		if c.validateRecord(record) {
			record = c.normalizeRecord(record)
			valid = append(valid, record)
		}
	}
	return valid
}

func (c *CSVToPostgresPipeline) validateRecord(record Record) bool {
	name, ok := record.Data["name"].(string)
	if !ok || name == "" {
		return false
	}
	age, ok := record.Data["age"].(int)
	if !ok || age < 0 {
		return false
	}
	return true
}

func (c *CSVToPostgresPipeline) normalizeRecord(record Record) Record {
	if name, ok := record.Data["name"].(string); ok {
		record.Data["name"] = strings.ToLower(name)
	}
	return record
}

func (c *CSVToPostgresPipeline) LoadData(data []Record) error {
	fmt.Printf("Loading %d records to PostgreSQL table %s\n", len(data), c.tableName)
	for _, record := range data {
		fmt.Printf("  INSERT INTO %s: %+v\n", c.tableName, record.Data)
	}
	return nil
}

func (c *CSVToPostgresPipeline) SendNotification() {
	fmt.Println("Sending completion notification via email")
}

// ============= FUNCTIONAL APPROACH (Idiomatic Go) =============

type PipelineStep func([]Record) ([]Record, error)

type FunctionalPipeline struct {
	steps []PipelineStep
	hooks []func() // Optional hooks
}

func NewFunctionalPipeline() *FunctionalPipeline {
	return &FunctionalPipeline{
		steps: make([]PipelineStep, 0),
		hooks: make([]func(), 0),
	}
}

func (f *FunctionalPipeline) AddStep(step PipelineStep) *FunctionalPipeline {
	f.steps = append(f.steps, step)
	return f
}

func (f *FunctionalPipeline) AddHook(hook func()) *FunctionalPipeline {
	f.hooks = append(f.hooks, hook)
	return f
}

func (f *FunctionalPipeline) Execute(initialData []Record) ([]Record, error) {
	data := initialData
	var err error

	for i, step := range f.steps {
		fmt.Printf("Executing step %d...\n", i+1)
		data, err = step(data)
		if err != nil {
			return nil, fmt.Errorf("step %d failed: %w", i+1, err)
		}
	}

	// Execute hooks
	for _, hook := range f.hooks {
		hook()
	}

	return data, nil
}

// Pipeline step functions
func ExtractFromCSV(filename string) PipelineStep {
	return func(_ []Record) ([]Record, error) {
		fmt.Printf("Extracting from %s\n", filename)
		return []Record{
			{ID: "1", Data: map[string]interface{}{"name": "John", "age": 30}},
			{ID: "2", Data: map[string]interface{}{"name": "Jane", "age": 25}},
		}, nil
	}
}

func FilterInvalid(records []Record) ([]Record, error) {
	fmt.Println("Filtering invalid records...")
	var valid []Record
	for _, r := range records {
		if name, ok := r.Data["name"].(string); ok && name != "" {
			valid = append(valid, r)
		}
	}
	return valid, nil
}

func NormalizeData(records []Record) ([]Record, error) {
	fmt.Println("Normalizing data...")
	for i := range records {
		if name, ok := records[i].Data["name"].(string); ok {
			records[i].Data["name"] = strings.ToLower(name)
		}
	}
	return records, nil
}

func LoadToDatabase(table string) PipelineStep {
	return func(records []Record) ([]Record, error) {
		fmt.Printf("Loading %d records to table %s\n", len(records), table)
		for _, r := range records {
			fmt.Printf("  INSERT: %+v\n", r.Data)
		}
		return records, nil
	}
}

func main() {
	fmt.Println("=== Traditional Template Method ===")
	csvPipeline := NewCSVToPostgresPipeline("data.csv", "postgres://localhost", "users")
	basePipeline := NewBasePipeline(csvPipeline)
	basePipeline.Execute()

	fmt.Println("\n=== Functional Pipeline (Idiomatic Go) ===")
	pipeline := NewFunctionalPipeline().
		AddStep(ExtractFromCSV("data.csv")).
		AddStep(FilterInvalid).
		AddStep(NormalizeData).
		AddStep(LoadToDatabase("users")).
		AddHook(func() {
			fmt.Println("Pipeline completed! Sending notification...")
		})

	result, err := pipeline.Execute(nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Processed %d records successfully\n", len(result))
	}
}
