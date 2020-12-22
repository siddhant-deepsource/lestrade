package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type analyzerTomlConfig struct {
	Category            string   `toml:"category"`
	Name                string   `toml:"name"`
	Shortcode           string   `toml:"shortcode"`
	Version             string   `toml:"version"`
	Description         string   `toml:"description"`
	Logo                string   `toml:"logo"`
	DocumentationURL    string   `toml:"documentation_url"`
	DiscussURL          string   `toml:"discuss_url"`
	AnalysisCommand     string   `toml:"analysis_command"`
	AutofixCommand      string   `toml:"autofix_command"`
	Trigger             string   `toml:"trigger"`
	Processors          []string `toml:"processors"`
	MINCPULimit         int      `toml:"min_cpu_limit"`
	MAXCPULimit         int      `toml:"max_cpu_limit"`
	MINMEMORYLimit      int      `toml:"min_memory_limit"`
	MAXMEMORYLimit      int      `toml:"max_memory_limit"`
	MetaSchemaPath      string   `toml:"meta_schema_path"`
	SupportedFiles      string   `toml:"supported_files"`
	DefaultTestPatterns []string `toml:"default_test_patterns"`
	Metrics             []string `toml:"metrics"`
	ExampleConfig       string   `toml:"example_config"`
	IssuesMeta          string   `toml:"issues_meta"`
	StarIssues          []string `toml:"star_issues"`
}

func main() {
	githubWorkspacePath := os.Getenv("GITHUB_WORKSPACE")
	analyzerTomlPath := path.Join(githubWorkspacePath, "analyzer.toml")
	var config analyzerTomlConfig
	if _, err := toml.DecodeFile(analyzerTomlPath, &config); err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Analyzer TOML Parsed.")

	// checking the category field
	switch config.Category {
	case "lang":
	case "conf":
	case "covg":
	default:
		log.Fatalln("Error in the \"category\" Field of analyzer.toml")
	}

	// checking the triggers
	switch config.Trigger {
	case "code":
	case "data":
	default:
		log.Fatalln("Error in the \"trigger\" Field of analyzer.toml. Acceptable values are - \"code\" and \"data\"")
	}

	// checking the processors
	for _, processor := range config.Processors {
		switch processor {
		case "source_code_load":
		case "skip_cq":
		default:
			log.Fatalln("Error in the \"processors\" Field of analyzer.toml. Acceptable values are - \"source_code_load\" and \"skip_cq\"")
		}
	}

	// checking the metrics
	for _, metric := range config.Metrics {
		switch metric {
		case "DCV":
		case "DDP":
		case "IDP":
		case "TCV":
		default:
			log.Fatalln("Error in the \"metrics\" Field of analyzer.toml. Acceptable values are - \"DCV\", \"DDP\", \"IDP\" and \"TCV\"")
		}
	}
}
