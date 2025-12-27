# SBOM Generate

A GitHub composite action for generating Software Bill of Materials (SBOM) using Syft.

## Overview

This action generates machine-readable SBOM files in multiple standard formats (CycloneDX, SPDX) to support dependency management and supply chain security initiatives.

## Features

- Generate SBOM using [Syft](https://github.com/anchore/syft)
- Support for multiple SBOM formats (CycloneDX JSON, SPDX JSON, Syft JSON)
- Automatic artifact upload to GitHub Actions
- Comprehensive SBOM statistics and summaries
- Configurable output paths and artifact names

## Inputs

| Name | Description | Required | Default |
|------|-------------|----------|---------|
| `source` | Source path or directory to scan for SBOM generation | No | `.` |
| `format` | SBOM output format (cyclonedx-json, spdx-json, syft-json) | No | `cyclonedx-json` |
| `output-file` | Output file name for the SBOM | No | `sbom.json` |
| `artifact-name` | Name of the artifact to upload | No | `sbom` |
| `upload-artifact` | Whether to upload the SBOM as a GitHub artifact | No | `true` |
| `syft-version` | Version of Syft to use | No | `v1.20.0` |

## Outputs

| Name | Description |
|------|-------------|
| `sbom-path` | Path to the generated SBOM file |
| `sbom-format` | Format of the generated SBOM |

## Usage

### Basic Usage

```yaml
- name: Generate SBOM
  uses: ./.github/actions/sbom-generate
  with:
    source: '.'
    format: 'cyclonedx-json'
    output-file: 'sbom.json'
```

### Generate Multiple Formats

```yaml
- name: Generate SBOM (CycloneDX)
  uses: ./.github/actions/sbom-generate
  with:
    source: '.'
    format: 'cyclonedx-json'
    output-file: 'sbom-cyclonedx.json'
    artifact-name: 'sbom-cyclonedx'

- name: Generate SBOM (SPDX)
  uses: ./.github/actions/sbom-generate
  with:
    source: '.'
    format: 'spdx-json'
    output-file: 'sbom-spdx.json'
    artifact-name: 'sbom-spdx'
```

### Without Artifact Upload

```yaml
- name: Generate SBOM
  uses: ./.github/actions/sbom-generate
  with:
    source: '.'
    upload-artifact: 'false'
```

## Supported Formats

- **cyclonedx-json**: CycloneDX JSON format (OWASP standard)
- **spdx-json**: SPDX JSON format (Linux Foundation standard)
- **syft-json**: Syft native JSON format

## Example Workflow

```yaml
name: Generate SBOM

on:
  push:
    branches: [main]
  release:
    types: [published]

jobs:
  sbom:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v6

      - name: Generate SBOM
        uses: ./.github/actions/sbom-generate
        with:
          source: '.'
          format: 'cyclonedx-json'
```

## License

Apache-2.0
