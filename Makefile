# Version
VERSION := 1.0.0

# Core configuration
ROOT_DIR := $(shell pwd)
DAGGER_DIR := $(ROOT_DIR)/.dagger
PROJECT_NAME := owasp-top-10-for-ai-agents

# Colors and formatting
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

# Validation helpers
DOCKER_CMD := $(shell command -v docker 2> /dev/null)
DAGGER_CMD := $(shell command -v dagger 2> /dev/null)
GO_CMD := $(shell command -v go 2> /dev/null)

# Dagger configuration
DAGGER_PLAN := cd $(ROOT_DIR) && go run $(DAGGER_DIR)/main.go
DAGGER_OPTS := --log-format=$(DAGGER_LOG_FORMAT) --log-level=$(DAGGER_LOG_LEVEL)
DAGGER_LOG_FORMAT ?= plain
DAGGER_LOG_LEVEL ?= info
export DAGGER_LOG_FORMAT
export DAGGER_LOG_LEVEL

# Check required tools
check-requirements:
ifndef DOCKER_CMD
	$(error "Docker is not installed. Please install Docker first.")
endif
ifndef GO_CMD
	$(error "Go is not installed. Please install Go first.")
endif

# ==================== Primary Targets ====================

.PHONY: all
all: check-requirements setup build test ## Run complete build pipeline

# ==================== Development Targets ====================

.PHONY: dev
dev: check-requirements ## Start development environment
	@echo "$(GREEN)Starting development environment$(RESET)"
	@$(DAGGER_PLAN) dev $(DAGGER_OPTS)

.PHONY: setup
setup: check-requirements ## Initialize development environment
	@echo "$(GREEN)Setting up development environment$(RESET)"
	@$(MAKE) dagger-init
	@test -f .env || cp .env.example .env
	@go mod download

# ==================== Build Targets ====================

.PHONY: build
build: check-requirements ## Build project
	@echo "$(GREEN)Building project$(RESET)"
	@$(DAGGER_PLAN) build $(DAGGER_OPTS)

.PHONY: test
test: check-requirements ## Run tests
	@echo "$(GREEN)Running tests$(RESET)"
	@$(DAGGER_PLAN) test $(DAGGER_OPTS)

# ==================== Documentation Targets ====================

.PHONY: docs
docs: generate pdf ## Generate all documentation

.PHONY: generate
generate: check-requirements ## Generate documentation
	@echo "$(GREEN)Generating documentation$(RESET)"
	@$(DAGGER_PLAN) generate $(DAGGER_OPTS) --output-dir=$(ROOT_DIR)/$(OUTPUT_DIR)

.PHONY: pdf
pdf: check-requirements ## Generate PDFs
	@echo "$(GREEN)Generating PDFs$(RESET)"
	@$(DAGGER_PLAN) generate-pdf $(DAGGER_OPTS) --output-dir=$(ROOT_DIR)/$(OUTPUT_DIR)

# ==================== Translation Targets ====================

.PHONY: i18n
i18n: translate validate-translations ## Handle all translation tasks

.PHONY: translate
translate: ## Manage translations using Dagger
	$(DAGGER_PLAN) translate

.PHONY: validate-translations
validate-translations: ## Validate translations using Dagger
	$(DAGGER_PLAN) validate-translations

# ==================== Security Targets ====================

.PHONY: security
security: sign verify ## Run all security tasks

.PHONY: sign
sign: ## Sign content using Dagger
	$(DAGGER_PLAN) sign

.PHONY: verify
verify: ## Verify signatures using Dagger
	$(DAGGER_PLAN) verify

# ==================== Cleanup Targets ====================

.PHONY: clean
clean: ## Clean build artifacts
	@echo "$(GREEN)Cleaning project$(RESET)"
	@$(DAGGER_PLAN) clean $(DAGGER_OPTS) --path=$(ROOT_DIR)

.PHONY: purge
purge: clean ## Deep clean including dependencies
	@echo "$(GREEN)Deep cleaning project$(RESET)"
	@rm -rf $(ROOT_DIR)/node_modules
	@rm -rf $(ROOT_DIR)/vendor
	@rm -f $(ROOT_DIR)/.env

# ==================== Monitoring Targets ====================

.PHONY: metrics
metrics: ## Collect and report metrics
	@echo "$(GREEN)Collecting metrics...$(RESET)"
	@$(DAGGER_PLAN) metrics $(DAGGER_OPTS)

.PHONY: monitor
monitor: ## Start monitoring stack
	@echo "$(GREEN)Starting monitoring stack...$(RESET)"
	@docker-compose up -d prometheus otel-collector

.PHONY: monitor-down
monitor-down: ## Stop monitoring stack
	@echo "$(GREEN)Stopping monitoring stack...$(RESET)"
	@docker-compose down

# ==================== Helper Targets ====================

.PHONY: version
version: ## Show version information
	@echo "$(PROJECT_NAME) version $(VERSION)"

.PHONY: help
help: ## Show this help message
	@echo "$(GREEN)$(PROJECT_NAME) v$(VERSION)$(RESET)"
	@echo "$(WHITE)Usage: make [target]$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"; printf "$(YELLOW)Targets:$(RESET)\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  $(YELLOW)%-15s$(RESET) %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# Ensure all commands run from root directory
.PHONY: %
%:
	@$(MAKE) -C $(ROOT_DIR) $@

# Set default target
.DEFAULT_GOAL := help

# Remove targets that are now handled by Dagger
.PHONY: remove-old-targets
	@echo "Removed old non-Dagger targets for clarity"

# Pipeline targets
.PHONY: pipeline
pipeline: generate sign verify ## Run full pipeline using Dagger

# Content management
.PHONY: content
content: ## Manage content using Dagger
	$(DAGGER_PLAN) content

.PHONY: glossary
glossary: ## Manage glossary using Dagger
	$(DAGGER_PLAN) glossary

# Dagger specific targets
.PHONY: dagger-init
dagger-init: ## Initialize Dagger
	@echo "$(GREEN)Initializing Dagger in $(DAGGER_DIR)$(RESET)"
	@cd $(DAGGER_DIR) && go mod tidy

.PHONY: dagger-debug
dagger-debug: ## Run Dagger with debug logging
	@DAGGER_LOG_LEVEL=debug $(MAKE) pipeline

.PHONY: dagger-plan
dagger-plan: ## Show Dagger execution plan
	@$(DAGGER_PLAN) plan
