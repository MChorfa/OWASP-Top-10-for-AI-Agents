#!/bin/bash

# Exit on error
set -e

# Install dependencies if not already installed
if ! command -v vuepress &> /dev/null; then
    echo "Installing VuePress..."
    npm install -g @vuepress/core@next @vuepress/client@next
    npm install -g vuepress@next
fi

# Create docs directory structure
mkdir -p docs/.vuepress docs/guide docs/risks docs/glossary

# Create VuePress config if it doesn't exist
if [ ! -f docs/.vuepress/config.js ]; then
    cat > docs/.vuepress/config.js << EOL
import { defineUserConfig } from 'vuepress'
import { defaultTheme } from '@vuepress/theme-default'

export default defineUserConfig({
    lang: 'en-US',
    title: 'OWASP Top 10 for AI Agents',
    description: 'Security risks and best practices for AI Agents',
    theme: defaultTheme({
        repo: 'OWASP/Top-10-for-AI-Agents',
        docsDir: 'docs',
        editLink: true,
        navbar: [
            { text: 'Home', link: '/' },
            { text: 'Guide', link: '/guide/' },
            { text: 'Risks', link: '/risks/' },
            { text: 'Glossary', link: '/glossary/' }
        ]
    })
})
EOL
fi

# Copy content
cp README.md docs/index.md
cp -r v1/translations/en/* docs/risks/
cp -r v1/glossary/markdown/* docs/glossary/

# Start development server
echo "Starting local preview server..."
vuepress dev docs
