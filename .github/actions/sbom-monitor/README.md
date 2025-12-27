# SBOM Monitor

A GitHub composite action for monitoring SBOM changes to track dependency updates and maintain supply chain security.

## Overview

This action monitors Software Bill of Materials (SBOM) files by comparing them against baseline versions to detect new, removed, or updated dependencies. It helps maintain visibility into supply chain changes and potential security risks.

## Features

- Compare current SBOM against baseline to detect changes
- Track new, removed, and updated dependencies
- Generate comprehensive monitoring reports in multiple formats
- Automatic baseline SBOM download from artifacts
- Supply chain security insights
- Support for markdown, JSON, and text report formats

## Inputs

| Name | Description | Required | Default |
|------|-------------|----------|---------|
| `current-sbom` | Path to the current SBOM file | Yes | - |
| `baseline-sbom` | Path to the baseline SBOM file to compare against | No | `''` |
| `artifact-name` | Name of the baseline SBOM artifact to download | No | `sbom-baseline` |
| `report-format` | Output format for the monitoring report (markdown, json, text) | No | `markdown` |
| `output-file` | Output file name for the monitoring report | No | `sbom-monitor-report.md` |
| `upload-artifact` | Whether to upload the monitoring report as a GitHub artifact | No | `true` |

## Outputs

| Name | Description |
|------|-------------|
| `report-path` | Path to the monitoring report |
| `new-dependencies` | Number of new dependencies added |
| `removed-dependencies` | Number of dependencies removed |
| `updated-dependencies` | Number of dependencies updated |

## Usage

### Basic Usage

```yaml
- name: Monitor SBOM Changes
  uses: ./.github/actions/sbom-monitor
  with:
    current-sbom: 'sbom/sbom.json'
```

### With Custom Baseline

```yaml
- name: Monitor SBOM
  uses: ./.github/actions/sbom-monitor
  with:
    current-sbom: 'sbom/sbom.json'
    baseline-sbom: 'sbom/baseline.json'
    report-format: 'markdown'
```

### Using Outputs

```yaml
- name: Monitor SBOM
  id: monitor
  uses: ./.github/actions/sbom-monitor
  with:
    current-sbom: 'sbom/sbom.json'

- name: Check Changes
  run: |
    echo "New dependencies: ${{ steps.monitor.outputs.new-dependencies }}"
    echo "Removed dependencies: ${{ steps.monitor.outputs.removed-dependencies }}"
    echo "Updated dependencies: ${{ steps.monitor.outputs.updated-dependencies }}"
    
    if [ "${{ steps.monitor.outputs.new-dependencies }}" -gt "0" ]; then
      echo "⚠️ New dependencies detected. Review for security."
    fi
```

## Report Formats

- **markdown**: GitHub-flavored markdown format (suitable for PR comments)
- **json**: Machine-readable JSON format
- **text**: Plain text format

## Example Workflow

```yaml
name: SBOM Monitoring

on:
  pull_request:
    branches: [main]

jobs:
  monitor:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v6

      - name: Generate Current SBOM
        uses: ./.github/actions/sbom-generate
        with:
          format: 'cyclonedx-json'

      - name: Monitor Changes
        uses: ./.github/actions/sbom-monitor
        with:
          current-sbom: 'sbom/sbom.json'
          artifact-name: 'sbom-baseline'

      - name: Comment on PR
        uses: marocchino/sticky-pull-request-comment@v2
        with:
          path: reports/sbom-monitor-report.md
```

## Use Cases

### Supply Chain Security

Monitor for:
- Unexpected new dependencies
- Removal of critical dependencies
- Version downgrades
- Dependency substitutions

### Compliance and Auditing

Track:
- Complete dependency history
- License changes
- Vulnerability introduction points
- Component provenance

### Development Workflow

Enable:
- Automated dependency review in PRs
- Baseline comparison for releases
- Continuous monitoring of dependency drift
- Security-focused code review

## License

Apache-2.0
