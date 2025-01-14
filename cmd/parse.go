package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

type BundleName struct {
	Name string `yaml:"name"`
}

type Workspace struct {
	Host string `yaml:"host"`
}

type Target struct {
	Mode      string    `yaml:"mode"`
	Default   bool      `yaml:"default,omitempty"`
	Workspace Workspace `yaml:"workspace"`
	Resources
}

type Task struct {
	TaskKey         string `yaml:"task_key"`
	SparkPythonTask *struct {
		PythonFile string `yaml:"python_file"`
	} `yaml:"spark_python_task,omitempty"`
	NotebookTask *struct {
		NotebookPath string `yaml:"notebook_path"`
	} `yaml:"notebook_task,omitempty"`
}

type WebhookNotification struct {
	ID string `yaml:"id"`
}

type WebhookNotifications struct {
	OnFailure []WebhookNotification `yaml:"on_failure,omitempty"`
}
type Schedule struct {
	PauseStatus string `yaml:"pause_status,omitempty"`
}

type Job struct {
	Name                 string                `yaml:"name"`
	Tasks                []Task                `yaml:"tasks"`
	Schedule             *Schedule             `yaml:"schedule,omitempty"`
	WebhookNotifications *WebhookNotifications `yaml:"webhook_notifications,omitempty"`
}

type Resources struct {
	Jobs map[string]Job `yaml:"jobs"`
}

type BundleConfig struct {
	Bundle  BundleName        `yaml:"bundle"`
	Include []string          `yaml:"include"`
	Targets map[string]Target `yaml:"targets"`
	Resources
}

type TargetResources struct {
	Jobs map[string]Job `yaml:"jobs"`
}

type LintConfig struct {
	NotificationsInProd bool `toml:"notifications_in_prod"`
}

func ParseBundleConfig(path string) (*BundleConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read bundle config: %w", err)
	}

	var config BundleConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse bundle config: %w", err)
	}

	if config.Include == nil {
		config.Include = []string{}
	}

	baseDir := filepath.Dir(path)

	for _, includePath := range config.Include {
		absIncludePath := filepath.Join(baseDir, includePath)
		includedConfig, err := ParseBundleConfig(absIncludePath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse included config '%s': %w", includePath, err)
		}
		mergeBundleConfig(&config, includedConfig)
	}

	return &config, nil
}

func mergeBundleConfig(mainConfig, includedConfig *BundleConfig) {
	for key, target := range includedConfig.Targets {
		mainConfig.Targets[key] = target
	}
	mainConfig.Include = append(mainConfig.Include, includedConfig.Include...)

	if mainConfig.Jobs == nil {
		mainConfig.Jobs = make(map[string]Job)
	}

	for key, job := range includedConfig.Jobs {
		mainConfig.Jobs[key] = job
	}
}

func mergeTargetResources(mainResources, includedResources *TargetResources) {
	if mainResources == nil || includedResources == nil {
		return
	}

	for key, job := range includedResources.Jobs {
		if existingJob, exists := mainResources.Jobs[key]; exists {
			mergeJob(&existingJob, &job)
			mainResources.Jobs[key] = existingJob
		} else {
			mainResources.Jobs[key] = job
		}
	}
}

func mergeJob(mainJob, includedJob *Job) {
	if includedJob.Schedule != nil {
		mainJob.Schedule = includedJob.Schedule
	}
	if includedJob.WebhookNotifications != nil {
		mainJob.WebhookNotifications = includedJob.WebhookNotifications
	}
}

func ParseLintConfig(path string) (*LintConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read package config: %w", err)
	}

	var config struct {
		Tool struct {
			BundleLint *LintConfig `toml:"bundlelint"`
		} `toml:"tool"`
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse package config: %w", err)
	}

	if config.Tool.BundleLint == nil {
		return nil, fmt.Errorf("[tool.bundlelint] section not found in package config")
	}

	return config.Tool.BundleLint, nil
}
