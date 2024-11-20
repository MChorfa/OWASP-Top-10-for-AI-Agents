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
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yourorg/hitl"
)

// Global logger
var log = logrus.New()

func init() {
	// Configure logger
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

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
		return &ContentValidationResult{Valid: false, Errors: []string{err.Error()}}, err
	}

	return &ContentValidationResult{Valid: true}, nil
}

// TranslateContent uses a local model or API to translate content
func (m *OwaspTop10ForAiAgents) TranslateContent(ctx context.Context, sourceDir *dagger.Directory, config TranslationConfig) (*dagger.Directory, error) {
	return dag.Container().
		From("local/translation-model:latest").
		WithMountedDirectory("/source", sourceDir).
		WithWorkdir("/app").
		WithExec([]string{
			"./local_translate",
			"--source", "/source",
			"--target", "/translations",
			"--languages", strings.Join(config.TargetLangs, ","),
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

// DocumentFormat represents supported output formats
type DocumentFormat string

const (
	FormatPDF  DocumentFormat = "pdf"
	FormatDOCX DocumentFormat = "docx"
	FormatHTML DocumentFormat = "html"
)

// DocumentOptions represents configuration for document conversion
type DocumentOptions struct {
	Format          DocumentFormat
	Template        string
	TableOfContents bool
	Stylesheet      string
	Metadata        map[string]string
}

// ConvertDocument converts markdown to various document formats with enhanced options
func (m *OwaspTop10ForAiAgents) ConvertDocument(ctx context.Context, inputDir *dagger.Directory, opts DocumentOptions) (*dagger.File, error) {
	return dag.Container().
		From("pandoc/core:latest").
		WithMountedDirectory("/work", inputDir).
		WithMountedDirectory("/templates", dag.Directory().WithNewFile("/templates/custom.tex", opts.Template)).
		WithWorkdir("/work").
		WithExec([]string{
			"pandoc",
			"--from", "markdown",
			"--to", string(opts.Format),
			"--template", "/templates/custom.tex",
			opts.Stylesheet,
			(opts.TableOfContents && "--toc") || "",
			"--metadata", formatMetadata(opts.Metadata),
			"-o", "output." + string(opts.Format),
			"*.md",
		}).
		File("output." + string(opts.Format)), nil
}

// BuildBook combines multiple markdown files into a single document
func (m *OwaspTop10ForAiAgents) BuildBook(ctx context.Context, inputDir *dagger.Directory, opts DocumentOptions) (*dagger.File, error) {
	return dag.Container().
		From("pandoc/core:latest").
		WithMountedDirectory("/work", inputDir).
		WithWorkdir("/work").
		WithExec([]string{
			"sh", "-c",
			`find . -name "*.md" -type f -print0 | sort -z | xargs -0 cat > combined.md && \
             pandoc combined.md \
             --from markdown \
             --to ` + string(opts.Format) + ` \
             --toc \
             --number-sections \
             --highlight-style tango \
             -V geometry:margin=1in \
             -o output.` + string(opts.Format),
		}).
		File("output." + string(opts.Format)), nil
}

// Helper function to format metadata for pandoc
func formatMetadata(metadata map[string]string) string {
	result := ""
	for k, v := range metadata {
		result += k + "=" + v + ","
	}
	return result
}

// LanguageConfig represents language-specific settings
type LanguageConfig struct {
	Code      string            // ISO language code (e.g., "en", "es", "fr")
	Name      string            // Language name
	Direction string            // "ltr" or "rtl"
	Font      string            // Default font for the language
	Metadata  map[string]string // Language-specific metadata
}

// TranslationRequest represents a single language translation request
func (m *OwaspTop10ForAiAgents) TranslateToLanguage(ctx context.Context, sourceDir *dagger.Directory, lang LanguageConfig) (*dagger.Directory, error) {
	return dag.Container().
		From("local/translation-model:latest").
		WithMountedDirectory("/source", sourceDir).
		WithWorkdir("/app").
		WithExec([]string{
			"./local_translate",
			"--source", "/source",
			"--target", "/translations/" + lang.Code,
			"--language", lang.Code,
		}).
		Directory("/translations/" + lang.Code), nil
}

// ConvertLanguageDocument converts documents with language-specific settings
func (m *OwaspTop10ForAiAgents) ConvertLanguageDocument(ctx context.Context, inputDir *dagger.Directory, lang LanguageConfig, opts DocumentOptions) (*dagger.File, error) {
	// Add language-specific options
	langOpts := []string{
		"--variable=lang=" + lang.Code,
		"--variable=dir=" + lang.Direction,
		fmt.Sprintf("--variable=mainfont=%s", lang.Font),
	}

	return dag.Container().
		From("pandoc/core:latest").
		WithMountedDirectory("/work", inputDir).
		WithWorkdir("/work").
		WithExec(append([]string{
			"pandoc",
			"--from", "markdown",
			"--to", string(opts.Format),
			"--pdf-engine=xelatex",
			opts.TableOfContents && "--toc" || "",
			"--metadata", formatMetadata(mergeMaps(opts.Metadata, lang.Metadata)),
		}, langOpts...)).
		File("output-" + lang.Code + "." + string(opts.Format)), nil
}

// Helper function to merge metadata maps
func mergeMaps(m1, m2 map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

// ProjectConfig represents the main configuration structure
type ProjectConfig struct {
	Languages     []LanguageConfig    `yaml:"languages"`
	Documentation DocumentationConfig `yaml:"documentation"`
	Templates     map[string]string   `yaml:"templates"`
	Validation    ValidationConfig    `yaml:"validation"`
	BuildOptions  BuildOptions        `yaml:"build"`
	Versioning    VersioningConfig    `yaml:"versioning"`
	Lifecycle     LifecycleConfig     `yaml:"lifecycle"`
}

// DocumentationConfig represents documentation settings
type DocumentationConfig struct {
	OutputFormats []DocumentFormat  `yaml:"outputFormats"`
	BasePath      string            `yaml:"basePath"`
	Structure     []DocumentSection `yaml:"structure"`
	Metadata      map[string]string `yaml:"metadata"`
}

// DocumentSection represents a section in the documentation
type DocumentSection struct {
	Title    string            `yaml:"title"`
	Path     string            `yaml:"path"`
	Children []DocumentSection `yaml:"children,omitempty"`
}

// ValidationConfig represents content validation settings
type ValidationConfig struct {
	Rules        []string `yaml:"rules"`
	Exclusions   []string `yaml:"exclusions"`
	MinWordCount int      `yaml:"minWordCount"`
	MaxWordCount int      `yaml:"maxWordCount"`
}

// BuildOptions represents build configuration
type BuildOptions struct {
	Parallel     bool   `yaml:"parallel"`
	CacheEnabled bool   `yaml:"cacheEnabled"`
	OutputDir    string `yaml:"outputDir"`
}

// VersionInfo represents version metadata
type VersionInfo struct {
	Version     string         `yaml:"version"`
	ReleaseDate string         `yaml:"releaseDate"`
	Stage       string         `yaml:"stage"` // draft, review, published
	Changes     []ChangeEntry  `yaml:"changes"`
	Reviewers   []string       `yaml:"reviewers"`
	Status      DocumentStatus `yaml:"status"`
}

// ChangeEntry tracks document changes
type ChangeEntry struct {
	Date        string   `yaml:"date"`
	Author      string   `yaml:"author"`
	Type        string   `yaml:"type"` // add, modify, remove
	Description string   `yaml:"description"`
	Files       []string `yaml:"files"`
}

// DocumentStatus represents document lifecycle states
type DocumentStatus string

const (
	StatusDraft     DocumentStatus = "draft"
	StatusReview    DocumentStatus = "review"
	StatusApproved  DocumentStatus = "approved"
	StatusPublished DocumentStatus = "published"
	StatusArchived  DocumentStatus = "archived"
)

// VersioningConfig represents versioning settings
type VersioningConfig struct {
	Strategy  string   `yaml:"strategy"` // semver, calver
	Branches  []string `yaml:"branches"` // main, develop, release/*
	TagPrefix string   `yaml:"tagPrefix"`
	Changelog bool     `yaml:"generateChangelog"`
}

// LifecycleConfig represents document lifecycle settings
type LifecycleConfig struct {
	RequiredReviewers int      `yaml:"requiredReviewers"`
	Stages            []string `yaml:"stages"`
	AutoArchive       bool     `yaml:"autoArchive"`
	RetentionPeriod   string   `yaml:"retentionPeriod"`
}

// LoadConfig loads project configuration from yaml
func (m *OwaspTop10ForAiAgents) LoadConfig(ctx context.Context, configFile *dagger.File) (*ProjectConfig, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedFile("/work/config.yaml", configFile).
		WithWorkdir("/work").
		WithExec([]string{"cat", "config.yaml"}).
		WithExec([]string{
			"sh", "-c",
			`yq eval . config.yaml`,
		}).
		Stdout(ctx)

}

// BuildFromConfig builds documentation using project configuration
func (m *OwaspTop10ForAiAgents) BuildFromConfig(ctx context.Context, config *ProjectConfig) error {
	for _, lang := range config.Languages {
		opts := createDocumentOptions(config, lang)
		if err := m.processSections(ctx, config.Documentation.Structure, lang, opts); err != nil {
			return err
		}
	}
	return nil
}

// createDocumentOptions creates document options for a language
func createDocumentOptions(config *ProjectConfig, lang LanguageConfig) DocumentOptions {
	return DocumentOptions{
		Format:          FormatPDF,
		TableOfContents: true,
		Template:        config.Templates[lang.Code],
		Metadata:        mergeMaps(config.Documentation.Metadata, lang.Metadata),
	}
}

// processSections processes multiple document sections
func (m *OwaspTop10ForAiAgents) processSections(ctx context.Context, sections []DocumentSection, lang LanguageConfig, opts DocumentOptions) error {
	for _, section := range sections {
		if err := m.processSection(ctx, section, lang, opts); err != nil {
			return err
		}
	}
	return nil
}

// processSection processes a single document section
func (m *OwaspTop10ForAiAgents) processSection(ctx context.Context, section DocumentSection, lang LanguageConfig, opts DocumentOptions) error {
	inputDir := dag.Directory().WithFile(section.Path)

	translatedDir, err := m.TranslateToLanguage(ctx, inputDir, lang)
	if err != nil {
		return err
	}

	if _, err := m.ConvertLanguageDocument(ctx, translatedDir, lang, opts); err != nil {
		return err
	}

	if len(section.Children) > 0 {
		return m.processSections(ctx, section.Children, lang, opts)
	}
	return nil
}

// Version management functions
func (m *OwaspTop10ForAiAgents) CreateRelease(ctx context.Context, version string, changes []ChangeEntry) (*VersionInfo, error) {
	return dag.Container().
		From("alpine/git:latest").
		WithMountedDirectory("/work", dag.Host().Directory(".")).
		WithWorkdir("/work").
		WithExec([]string{
			"sh", "-c",
			fmt.Sprintf(`
                git tag -a "v%s" -m "Release %s"
                git push origin "v%s"
            `, version, version, version),
		}).
		Sync(ctx)
}

// Lifecycle management functions
func (m *OwaspTop10ForAiAgents) UpdateDocumentStatus(ctx context.Context, docPath string, status DocumentStatus) error {
	container := dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/work", dag.Host().Directory(".")).
		WithWorkdir("/work").
		WithExec([]string{
			"sh", "-c",
			fmt.Sprintf(`yq eval '.status = "%s"' -i "%s"`, status, docPath),
		}).
		WithExec([]string{"git", "add", docPath}).
		WithExec([]string{"git", "commit", "-m", fmt.Sprintf("Update document status to %s", status)})

	return container.Sync(ctx)
}

// Archive management
func (m *OwaspTop10ForAiAgents) ArchiveVersion(ctx context.Context, version string) (*dagger.Container, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/work", dag.Host().Directory(".")).
		WithWorkdir("/work").
		WithExec([]string{
			"sh", "-c",
			fmt.Sprintf(`
                mkdir -p archives/v%s
                cp -r docs/* archives/v%s/
                git add archives/v%s
                git commit -m "Archive version %s"
            `, version, version, version, version),
		}).
		Sync(ctx)
}

// CriticalOperation performs a sensitive task that requires human approval
func (m *OwaspTop10ForAiAgents) CriticalOperation(ctx context.Context, params ...interface{}) (result interface{}, err error) {
	approved, err := hitl.RequestApproval("Execute CriticalOperation", ctx)
	if err != nil {
		log.WithError(err).Error("Approval process failed")
		return nil, fmt.Errorf("approval process failed: %v", err)
	}
	if !approved {
		log.Warn("Operation not approved by human operator")
		return nil, fmt.Errorf("operation not approved by human operator")
	}

	// ...existing code...

	log.Info("CriticalOperation executed successfully")
	return result, nil
}
