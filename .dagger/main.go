// A generated module for OwaspTop10ForAiAgents functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/owasp-top-10-for-ai-agents/internal/dagger"
)

// GlossaryEntry represents a single glossary item
type GlossaryEntry struct {
	Term       string
	Definition string
	References []string
}

type OwaspTop10ForAiAgents struct{}

// Returns a container that echoes whatever string argument is provided
func (m *OwaspTop10ForAiAgents) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *OwaspTop10ForAiAgents) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

// AddGlossaryEntry adds a new entry to the glossary
func (m *OwaspTop10ForAiAgents) AddGlossaryEntry(ctx context.Context, entry GlossaryEntry) (*dagger.Container, error) {
	return dag.Container().
		From("alpine:latest").
		WithWorkdir("/work").
		WithExec([]string{"sh", "-c", "echo '" + entry.Term + ": " + entry.Definition + "' >> glossary.md"}), nil
}

// ConvertToPDF converts markdown files to PDF
func (m *OwaspTop10ForAiAgents) ConvertToPDF(ctx context.Context, inputDir *dagger.Directory) (*dagger.File, error) {
	return dag.Container().
		From("pandoc/core:latest").
		WithMountedDirectory("/work", inputDir).
		WithWorkdir("/work").
		WithExec([]string{"pandoc", "-f", "markdown", "-t", "pdf", "input.md", "-o", "output.pdf"}).
		File("output.pdf"), nil
}

// SignWithCosign signs the generated PDF with cosign
func (m *OwaspTop10ForAiAgents) SignWithCosign(ctx context.Context, pdfFile *dagger.File) (*dagger.File, error) {
	return dag.Container().
		From("gcr.io/projectsigstore/cosign:latest").
		WithMountedFile("/work/file.pdf", pdfFile).
		WithWorkdir("/work").
		WithExec([]string{"cosign", "sign", "file.pdf"}).
		File("file.pdf.sig"), nil
}

// CalculateHash calculates SHA256 hash of a file
func (m *OwaspTop10ForAiAgents) CalculateHash(ctx context.Context, file *dagger.File) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedFile("/work/file", file).
		WithWorkdir("/work").
		WithExec([]string{"sha256sum", "file"}).
		Stdout(ctx)
}

// ContentValidationResult represents validation results
type ContentValidationResult struct {
	Valid       bool
	Errors      []string
	Suggestions []string
}

// TranslationConfig represents translation settings
type TranslationConfig struct {
	SourceLang  string
	TargetLangs []string
	Glossary    map[string]map[string]string
}

func (m *OwaspTop10ForAiAgents) ValidateContent(ctx context.Context, contentDir *dagger.Directory) (*ContentValidationResult, error) {
	container := dag.Container().
		From("node:latest").
		WithMountedDirectory("/content", contentDir).
		WithWorkdir("/content").
		WithExec([]string{"npx", "markdownlint", "."}).
		WithExec([]string{"npx", "markdown-link-check", "."})

	_, err := container.Stdout(ctx)

	if err != nil {
		return &ContentValidationResult{Valid: false, Errors: []string{err.Error()}}, nil
	}

	return &ContentValidationResult{Valid: true}, nil
}

func (m *OwaspTop10ForAiAgents) TranslateContent(ctx context.Context, sourceDir *dagger.Directory, config TranslationConfig) (*dagger.Directory, error) {
	return dag.Container().
		From("python:3.9-slim").
		WithMountedDirectory("/source", sourceDir).
		WithWorkdir("/app").
		WithExec([]string{"pip", "install", "googletrans==3.1.0a0"}).
		WithExec([]string{
			"python", "-c",
			"from googletrans import Translator; translator = Translator()",
		}).
		Directory("/translations"), nil
}

func (m *OwaspTop10ForAiAgents) ValidateTranslations(ctx context.Context, sourceDir *dagger.Directory, translatedDir *dagger.Directory) (bool, error) {
	container := dag.Container().
		From("python:3.9-slim").
		WithMountedDirectory("/source", sourceDir).
		WithMountedDirectory("/translated", translatedDir).
		WithWorkdir("/app").
		WithExec([]string{"pip", "install", "deep-translator"}).
		WithExec([]string{"python", "scripts/validate_translations.py"})

	result, err := container.Stdout(ctx)
	if err != nil {
		return false, err
	}

	return result == "valid", nil
}

func (m *OwaspTop10ForAiAgents) GenerateContentMetrics(ctx context.Context, contentDir *dagger.Directory) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/content", contentDir).
		WithWorkdir("/content").
		WithExec([]string{
			"sh", "-c",
			`find . -name "*.md" -exec wc -w {} \; | awk '{total += $1} END {print "Total words: " total}'`,
		}).
		Stdout(ctx)
}
