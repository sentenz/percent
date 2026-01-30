# `skills/`

Agent skills are modular capabilities that AI agents can utilize to perform specific tasks within a project. Skills enhance the functionality of AI agents by providing them with specialized context knowledge and tools.

- [1. Details](#1-details)
  - [1.1. Agent Skills](#11-agent-skills)
- [2. References](#2-references)

## 1. Details

### 1.1. Agent Skills

Skills are documented in individual `SKILL.md` files located in appropriate subdirectories following the [Agent Skills](https://agentskills.io/specification) specification, containing metadata and descriptions of their purpose, usage, and integration within the project.

- [Unit Testing](unit-testing/SKILL.md)
  > Unit test creation using consistent software testing patterns for Go projects.

- [Fuzz Testing](fuzz-testing/SKILL.md)
  > Fuzz test creation using Go's native fuzzing engine with coverage-guided testing.

- [Benchmark Testing](benchmark-testing/SKILL.md)
  > Benchmark test creation for performance measurement and optimization.

## 2. References

- [AGENTS.md](../../AGENTS.md) page.
- Agent Skills [Specification](https://agentskills.io/specification) page.
