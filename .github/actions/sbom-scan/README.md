# SBOM Scan

A GitHub composite action for scanning SBOM files for known vulnerabilities using Software Composition Analysis (SCA).

## Overview

This action uses [Grype](https://github.com/anchore/grype) to scan SBOM files for known security vulnerabilities in dependencies, providing comprehensive vulnerability reports and security insights.

## Features

- Scan SBOM files for known vulnerabilities using Grype
- Support for multiple output formats (table, JSON, SARIF, CycloneDX)
- Configurable severity thresholds for build failures
- Detailed vulnerability statistics (critical, high, medium, low)
- Automatic artifact upload for vulnerability reports
- Integration with GitHub Security tab via SARIF

## Inputs

| Name | Description | Required | Default |
|------|-------------|----------|---------|
| `sbom-file` | Path to the SBOM file to scan | Yes | - |
| `fail-on-severity` | Fail on vulnerability severity (negligible, low, medium, high, critical) | No | `high` |
| `output-format` | Output format for vulnerability report (table, json, sarif, cyclonedx) | No | `table` |
| `output-file` | Output file name for the vulnerability report | No | `vulnerability-report.txt` |
| `grype-version` | Version of Grype to use | No | `v0.104.3` |
| `upload-artifact` | Whether to upload the vulnerability report as a GitHub artifact | No | `true` |

## Outputs

| Name | Description |
|------|-------------|
| `report-path` | Path to the vulnerability report |
| `vulnerabilities-found` | Number of vulnerabilities found |
| `critical-count` | Number of critical vulnerabilities |
| `high-count` | Number of high vulnerabilities |

## Usage

### Basic Usage

```yaml
- name: Scan SBOM for Vulnerabilities
  uses: ./.github/actions/sbom-scan
  with:
    sbom-file: 'sbom/sbom.json'
    fail-on-severity: 'high'
```

### With Custom Output Format

```yaml
- name: Scan SBOM
  uses: ./.github/actions/sbom-scan
  with:
    sbom-file: 'sbom/sbom.json'
    fail-on-severity: 'critical'
    output-format: 'json'
    output-file: 'vulnerabilities.json'
```

### Using Outputs

```yaml
- name: Scan SBOM
  id: scan
  uses: ./.github/actions/sbom-scan
  with:
    sbom-file: 'sbom/sbom.json'

- name: Check Results
  run: |
    echo "Vulnerabilities found: ${{ steps.scan.outputs.vulnerabilities-found }}"
    echo "Critical: ${{ steps.scan.outputs.critical-count }}"
    echo "High: ${{ steps.scan.outputs.high-count }}"
```

## Severity Levels

- **negligible**: Minor issues with minimal impact
- **low**: Low severity vulnerabilities
- **medium**: Medium severity vulnerabilities
- **high**: High severity vulnerabilities (default threshold)
- **critical**: Critical vulnerabilities requiring immediate attention

## Output Formats

- **table**: Human-readable table format
- **json**: Machine-readable JSON format
- **sarif**: SARIF format for GitHub Code Scanning
- **cyclonedx**: CycloneDX VEX format

## Example Workflow

```yaml
name: Security Scan

on:
  push:
    branches: [main]
  schedule:
    - cron: '0 6 * * *'  # Daily at 6 AM UTC

jobs:
  scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v6

      - name: Generate SBOM
        uses: ./.github/actions/sbom-generate
        with:
          format: 'cyclonedx-json'

      - name: Scan for Vulnerabilities
        uses: ./.github/actions/sbom-scan
        with:
          sbom-file: 'sbom/sbom.json'
          fail-on-severity: 'high'
```

## License

Apache-2.0
