# Colors for output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

# Directories
BINARY_DIR := markdown-to-pdf
BINARY_NAME := markdown-to-pdf
OUTPUT_DIR := generated
CACHE_DIR := .cache
TRANSLATIONS_DIR := v1/translations

# Default target
.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@echo "$(GREEN)OWASP Top 10 for AI Agents - Make Targets$(RESET)"
	@echo "$(WHITE)Usage: make [target]$(RESET)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(YELLOW)%-30s$(RESET) %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the markdown-to-pdf binary
	@echo "$(GREEN)Building markdown-to-pdf...$(RESET)"
	@cd $(BINARY_DIR) && go build -o $(BINARY_NAME)
	@echo "$(GREEN)Build complete!$(RESET)"

.PHONY: install-fonts
install-fonts: ## Install required fonts
	@echo "$(GREEN)Installing fonts...$(RESET)"
	@mkdir -p $(HOME)/Library/Fonts
	@cp -f fonts/* $(HOME)/Library/Fonts/ 2>/dev/null || true
	@echo "$(GREEN)Fonts installed!$(RESET)"

.PHONY: generate-pdf
generate-pdf: build ## Generate PDF documents
	@echo "$(GREEN)Generating PDFs...$(RESET)"
	@mkdir -p $(OUTPUT_DIR)
	@./$(BINARY_DIR)/$(BINARY_NAME) -language en
	@echo "$(GREEN)PDF generation complete!$(RESET)"

.PHONY: generate-translations
generate-translations: build ## Generate translated PDFs
	@echo "$(GREEN)Generating translated PDFs...$(RESET)"
	@for lang in $(TRANSLATIONS_DIR)/*/ ; do \
		if [ -d "$$lang" ]; then \
			lang_code=$$(basename "$$lang"); \
			echo "Generating PDF for $$lang_code..."; \
			./$(BINARY_DIR)/$(BINARY_NAME) -language $$lang_code || true; \
		fi \
	done
	@echo "$(GREEN)Translation generation complete!$(RESET)"

.PHONY: generate-fr
generate-fr: build ## Generate French translation
	@echo "$(GREEN)Generating French translation...$(RESET)"
	@mkdir -p $(OUTPUT_DIR)
	@./$(BINARY_DIR)/$(BINARY_NAME) -language fr
	@echo "$(GREEN)French translation generated!$(RESET)"

.PHONY: pdf-fr
pdf-fr: generate-fr ## Generate French PDF
	@echo "$(GREEN)Generating French PDF...$(RESET)"
	@mkdir -p $(OUTPUT_DIR)
	@./$(BINARY_DIR)/$(BINARY_NAME) -language fr -pdf
	@echo "$(GREEN)French PDF generated!$(RESET)"

.PHONY: clean-build
clean-build: ## Clean build artifacts
	@echo "$(GREEN)Cleaning build artifacts...$(RESET)"
	@rm -rf $(OUTPUT_DIR)
	@rm -rf $(CACHE_DIR)
	@rm -f $(BINARY_DIR)/$(BINARY_NAME)
	@echo "$(GREEN)Cleanup complete!$(RESET)"

.PHONY: clean
clean: ## Clean repository safely
	@echo "$(GREEN)Cleaning repository safely...$(RESET)"
	@find . -type f -name "*.tmp" -delete
	@find . -type f -name "*.log" -delete
	@find . -type f -name ".DS_Store" -delete
	@find . -type d -name "__pycache__" -exec rm -rf {} + 2>/dev/null || true
	@rm -rf $(OUTPUT_DIR) 2>/dev/null || true
	@rm -rf $(CACHE_DIR) 2>/dev/null || true
	@rm -rf markdown-to-pdf/cache/ 2>/dev/null || true
	@echo "$(GREEN)Repository cleanup complete!$(RESET)"

.PHONY: test
test: ## Run tests
	@echo "$(GREEN)Running tests...$(RESET)"
	@cd $(BINARY_DIR) && go test -v ./...
	@echo "$(GREEN)Tests complete!$(RESET)"

.PHONY: validate-translations
validate-translations: ## Validate translation files
	@echo "$(GREEN)Validating translations...$(RESET)"
	@for lang in $(TRANSLATIONS_DIR)/*/ ; do \
		if [ -d "$$lang" ]; then \
			if [ ! -f "$$lang/metadata.yaml" ]; then \
				echo "$(YELLOW)Warning: Missing metadata.yaml in $$lang$(RESET)"; \
			fi; \
		fi \
	done
	@echo "$(GREEN)Translation validation complete!$(RESET)"

.PHONY: setup-arabic
setup-arabic: ## Setup Arabic translation environment
	@echo "$(GREEN)Setting up Arabic translation environment...$(RESET)"
	@mkdir -p $(TRANSLATIONS_DIR)/ar
	@mkdir -p $(TRANSLATIONS_DIR)/styles
	@echo "$(GREEN)Arabic translation environment setup complete!$(RESET)"

.PHONY: validate-rtl
validate-rtl: ## Validate RTL content
	@echo "$(GREEN)Validating RTL content...$(RESET)"
	@for file in $(TRANSLATIONS_DIR)/ar/*.md; do \
		if [ -f "$$file" ]; then \
			echo "Checking $$file..."; \
			grep -l "direction: rtl" "$$file" > /dev/null || \
			echo "$(YELLOW)Warning: Missing RTL direction in $$file$(RESET)"; \
		fi \
	done
	@echo "$(GREEN)RTL validation complete!$(RESET)"

# Dagger targets
.PHONY: dagger-init
dagger-init: ## Initialize Dagger development environment
	@echo "$(GREEN)Initializing Dagger development environment...$(RESET)"
	@cd .dagger && go mod tidy
	@echo "$(GREEN)Dagger environment initialized!$(RESET)"

.PHONY: dagger-validate-glossary
dagger-validate-glossary: ## Validate glossary using Dagger
	@echo "$(GREEN)Validating glossary...$(RESET)"
	@cd .dagger && go run main.go validate-glossary
	@echo "$(GREEN)Glossary validation complete!$(RESET)"

.PHONY: dagger-export-markdown
dagger-export-markdown: ## Export glossary to markdown using Dagger
	@echo "$(GREEN)Exporting glossary to markdown...$(RESET)"
	@cd .dagger && go run main.go export-markdown
	@echo "$(GREEN)Markdown export complete!$(RESET)"

.PHONY: dagger-update-timestamp
dagger-update-timestamp: ## Update glossary timestamp using Dagger
	@echo "$(GREEN)Updating glossary timestamp...$(RESET)"
	@cd .dagger && go run main.go update-timestamp
	@echo "$(GREEN)Timestamp update complete!$(RESET)"

.PHONY: dagger-sign-glossary
dagger-sign-glossary: ## Sign and hash the glossary file
	@echo "$(GREEN)Signing glossary...$(RESET)"
	@cd .dagger && go run main.go sign-glossary
	@echo "$(GREEN)Glossary signing complete!$(RESET)"

.PHONY: dagger-verify-glossary
dagger-verify-glossary: ## Verify glossary hash and signature
	@echo "$(GREEN)Verifying glossary...$(RESET)"
	@cd .dagger && go run main.go verify-glossary
	@echo "$(GREEN)Glossary verification complete!$(RESET)"

.PHONY: dagger-all
dagger-all: dagger-init dagger-validate-glossary dagger-export-markdown dagger-sign-glossary dagger-verify-glossary ## Run all Dagger tasks

# Content signing and verification
.PHONY: dagger-sign-content
dagger-sign-content:
	go run .dagger/main.go sign-content

.PHONY: dagger-verify-content
dagger-verify-content:
	go run .dagger/main.go verify-content

.PHONY: dagger-sign-translation
dagger-sign-translation:
	go run .dagger/main.go sign-translation

.PHONY: dagger-verify-translation
dagger-verify-translation:
	go run .dagger/main.go verify-translation

# Content and translation updates
.PHONY: dagger-update-content
dagger-update-content:
	go run .dagger/main.go update-content
	go run .dagger/main.go sign-content

.PHONY: dagger-update-translation
dagger-update-translation:
	go run .dagger/main.go update-translation
	go run .dagger/main.go sign-translation

# Combined update and verify
.PHONY: dagger-update-all
dagger-update-all: dagger-update-content dagger-update-translation dagger-verify-content dagger-verify-translation
	@echo "$(GREEN)All content updated and verified!$(RESET)"

# Content lifecycle management
.PHONY: dagger-process-content
dagger-process-content:
	@echo "$(YELLOW)Processing content...$(RESET)"
	GIT_USER=$$(git config user.name) go run .dagger/main.go process-content
	@echo "$(GREEN)Content processed!$(RESET)"

.PHONY: dagger-content-review
dagger-content-review: dagger-process-content
	@echo "$(YELLOW)Checking content review status...$(RESET)"
	@echo "$(GREEN)Content review status checked!$(RESET)"

.PHONY: dagger-content-approve
dagger-content-approve: dagger-process-content
	@echo "$(YELLOW)Approving content...$(RESET)"
	GIT_USER=$$(git config user.name) go run .dagger/main.go approve-content
	@echo "$(GREEN)Content approved!$(RESET)"

.PHONY: dagger-content-publish
dagger-content-publish: dagger-process-content
	@echo "$(YELLOW)Publishing content...$(RESET)"
	GIT_USER=$$(git config user.name) go run .dagger/main.go publish-content
	@echo "$(GREEN)Content published!$(RESET)"

.PHONY: dagger-content-archive
dagger-content-archive:
	@echo "$(YELLOW)Archiving content...$(RESET)"
	GIT_USER=$$(git config user.name) go run .dagger/main.go archive-content
	@echo "$(GREEN)Content archived!$(RESET)"

# Translation lifecycle management
.PHONY: dagger-translation-sync
dagger-translation-sync:
	@echo "$(YELLOW)Synchronizing translations...$(RESET)"
	go run .dagger/main.go update-translation
	@echo "$(GREEN)Translations synchronized!$(RESET)"

.PHONY: dagger-translation-approve
dagger-translation-approve:
	@echo "$(YELLOW)Approving translations...$(RESET)"
	GIT_USER=$$(git config user.name) go run .dagger/main.go approve-translation
	@echo "$(GREEN)Translations approved!$(RESET)"

# Combined lifecycle management
.PHONY: dagger-lifecycle-status
dagger-lifecycle-status:
	@echo "$(YELLOW)Checking content lifecycle status...$(RESET)"
	go run .dagger/main.go lifecycle-status
	@echo "$(GREEN)Status check complete!$(RESET)"
