# OWASP Top 10 for AI Agents - Glossary

This directory contains the multilingual glossary for the OWASP Top 10 for AI Agents project. The glossary ensures consistent terminology across all translations and provides clear definitions for key concepts.

## Structure

- `glossary.yaml`: Main glossary file containing all terms and their translations
- `schema.json`: JSON Schema for validating the glossary structure
- `markdown/`: Generated markdown files for each language
  - `glossary_en.md`: English glossary
  - `glossary_ar.md`: Arabic glossary
  - `glossary_fr.md`: French glossary

## Glossary Format

Each term in the glossary follows this structure:

```yaml
- id: "term_id"
  en:
    term: "Term in English"
    definition: "Definition in English"
    context: "Usage context"
    notes: "Optional additional notes"
    examples:
      - "Example 1"
      - "Example 2"
  ar:
    term: "Term in Arabic"
    definition: "Definition in Arabic"
    context: "Usage context in Arabic"
  fr:
    term: "Term in French"
    definition: "Definition in French"
    context: "Usage context in French"
```

## Managing the Glossary

Use the `scripts/manage_glossary.py` script to manage the glossary:

```bash
# Validate the glossary against the schema
python scripts/manage_glossary.py --validate

# Update the last_updated field
python scripts/manage_glossary.py --update

# Export to markdown files
python scripts/manage_glossary.py --export-md
```

## Contributing

When adding new terms or updating existing ones:

1. Add the term to `glossary.yaml`
2. Follow the established format
3. Provide translations for all supported languages
4. Run validation to ensure correctness
5. Submit a pull request

## Guidelines for Translations

- Maintain consistency with existing translations
- Preserve technical terms where appropriate
- Consider cultural context
- Include usage examples when helpful
- Provide clear and concise definitions

## Signing and Verification

The glossary is cryptographically signed using [Cosign](https://github.com/sigstore/cosign) to ensure integrity and authenticity. Each update to the glossary generates:

1. A SHA-256 hash stored in both:
   - The YAML file itself
   - A separate `glossary.sha256` file

2. Cosign signatures for:
   - The glossary file (`glossary.yaml.sig`)
   - The hash file (`glossary.sha256.sig`)

### Signing Process

To sign the glossary, you need:
1. A Cosign key pair
2. The private key password

```bash
# Generate a new key pair if you don't have one
cosign generate-key-pair

# Set the required environment variables
export COSIGN_PRIVATE_KEY=$(cat cosign.key)
export COSIGN_PASSWORD='your-key-password'

# Sign the glossary
make dagger-sign-glossary
```

### Verification Process

To verify the signatures, you need the public key:

```bash
# Set the public key
export COSIGN_PUBLIC_KEY=$(cat cosign.pub)

# Verify the signatures
make dagger-verify-glossary
```

The verification process checks:
1. The integrity of the glossary content against its stored hash
2. The authenticity of both the glossary and hash files using Cosign signatures

### Security Notes

- Keep the private key secure and never commit it to the repository
- The public key should be distributed to all contributors who need to verify signatures
- All changes to the glossary should be signed by authorized maintainers
- The signing process is integrated into the CI/CD pipeline

## Validation

The glossary is automatically validated:
- On pull requests
- When generating translations
- Before creating releases

The validation checks:
- YAML syntax
- Schema compliance
- Required translations
- Format consistency
