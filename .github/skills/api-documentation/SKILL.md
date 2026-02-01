---
name: api-documentation
description: Automates API documentation creation for Go projects using godoc conventions and best practices. Use when creating, modifying, or reviewing documentation comments, or when the user mentions godoc, documentation, or API docs.
metadata:
  version: "1.0"
  activation:
    implicit: true
    priority: 2
    triggers:
      - "godoc"
      - "documentation"
      - "doc comment"
      - "api doc"
      - "package doc"
      - "document function"
    match:
      languages: ["go", "golang"]
      paths: ["pkg/**/*.go", "internal/**/*.go"]
      prompt_regex: "(?i)(godoc|documentation|doc comment|api doc|package doc|document)"
  usage:
    load_on_prompt: true
    autodispatch: true
---

# API Documentation

Instructions for AI coding agents on automating API documentation creation using godoc conventions and best practices in this Go project.

- [1. Benefits](#1-benefits)
- [2. Principles](#2-principles)
- [3. Conventions](#3-conventions)
  - [3.1. Package Documentation](#31-package-documentation)
  - [3.2. Function Documentation](#32-function-documentation)
  - [3.3. Type Documentation](#33-type-documentation)
  - [3.4. Constant and Variable Documentation](#34-constant-and-variable-documentation)
  - [3.5. Example Functions](#35-example-functions)
- [4. Workflow](#4-workflow)
- [5. Commands](#5-commands)
- [6. Style Guide](#6-style-guide)
- [7. Template](#7-template)
  - [7.1. Package Documentation Template](#71-package-documentation-template)
  - [7.2. Function Documentation Template](#72-function-documentation-template)
  - [7.3. Type Documentation Template](#73-type-documentation-template)
  - [7.4. Example Function Template](#74-example-function-template)
- [8. References](#8-references)

## 1. Benefits

- Discoverability
  > Well-documented APIs enable developers to quickly understand and use package functionality without reading implementation details.

- Consistency
  > Following godoc conventions ensures uniform documentation style across the codebase, improving readability and maintainability.

- Self-Documenting Code
  > Documentation comments serve as inline reference, reducing the need for external documentation and keeping docs in sync with code.

- Tooling Integration
  > Godoc-compatible comments integrate with Go tooling, enabling automatic documentation generation and IDE support.

- Onboarding
  > Comprehensive API documentation reduces onboarding time for new team members and external contributors.

## 2. Principles

Effective API documentation follows these core principles.

- Complete
  > Document all public APIs including classes, functions, parameters, return values, and exceptions. Private implementation details may be omitted.

- Contextual
  > Documentation provides context about usage patterns, performance characteristics, and thread safety guarantees.

- Consistent
  > Use a uniform style, format, terminology and structure throughout the API documentation using the patterns defined in this skill.

- Concise
  > Use clear, brief descriptions. Avoid redundant information that restates what is obvious from the signature.

- Concrete
  > Provide specific details about behavior, edge cases, and error conditions rather than vague statements.

- Convenient
  > Documentation should be easy to access and navigate, integrated with development tools and workflows.

- Accurate
  > Documentation must match the actual behavior. Update documentation whenever the implementation changes.

- Actionable
  > Include usage examples, preconditions, postconditions, and error handling to help developers use the API correctly.

## 3. Conventions

### 3.1. Package Documentation

Package documentation provides an overview of the package's purpose and usage.

- Location
  > Place package documentation in a `doc.go` file or at the top of the primary source file.

- Format
  > Start with `Package <name>` followed by a description of the package's functionality.

- Content
  > Include purpose, main types, key functions, usage examples, and any important notes.

### 3.2. Function Documentation

Function documentation describes what a function does, its parameters, and return values.

- First Sentence
  > Start with the function name followed by a verb describing its action.

- Parameters
  > Document non-obvious parameters and their expected values.

- Return Values
  > Describe what the function returns, including error conditions.

- Panics
  > Document any conditions that cause the function to panic.

### 3.3. Type Documentation

Type documentation describes the purpose and usage of types (structs, interfaces, aliases).

- First Sentence
  > Start with the type name followed by a description of what it represents.

- Fields
  > Document non-obvious struct fields inline.

- Methods
  > Document each method following function documentation conventions.

### 3.4. Constant and Variable Documentation

Constants and variables should be documented to explain their purpose and valid values.

- Grouped Constants
  > Document the group with a single comment, then document individual constants as needed.

- Sentinel Errors
  > Document error variables with the conditions that produce them.

### 3.5. Example Functions

Example functions demonstrate how to use package functionality.

- Naming
  > Name examples as `Example`, `ExampleFunctionName`, or `ExampleTypeName_MethodName`.

- Output Comments
  > Include `// Output:` comments to make examples testable.

- Completeness
  > Examples should be self-contained and runnable.

## 4. Workflow

1. Identify

    Identify exported functions, types, constants, and variables in `pkg/` or `internal/` that require documentation.

2. Add/Create

    Add documentation comments directly above the exported identifier.

3. Documentation Requirements

    Ensure documentation covers:
    - Purpose and functionality
    - Parameters and their expected values
    - Return values and error conditions
    - Usage examples for complex APIs
    - Edge cases and limitations

4. Apply Templates

    Structure all documentation using the [template](#7-template) patterns.

5. Validate

    Run `go doc` to verify documentation renders correctly.

## 5. Commands

| Command                        | Description                                    |
| ------------------------------ | ---------------------------------------------- |
| `go doc ./pkg/percent`         | View package documentation                     |
| `go doc ./pkg/percent.Percent` | View function documentation                    |
| `go doc -all ./pkg/percent`    | View all documentation including unexported    |
| `godoc -http=:6060`            | Start local documentation server               |
| `pkgsite -http=:8080`          | Start pkgsite server for pkg.go.dev style docs |

## 6. Style Guide

- License Header
  > All source files must include `// SPDX-License-Identifier: Apache-2.0` at the top.

- First Sentence
  > Start documentation with the identifier name (function, type, constant) followed by a verb.

- Complete Sentences
  > Write documentation as complete sentences with proper punctuation.

- Present Tense
  > Use present tense to describe what the code does (e.g., `Percent calculates` not `Percent will calculate`).

- Line Length
  > Keep documentation lines under 80 characters for readability in terminals.

- Paragraphs
  > Separate paragraphs with blank comment lines for complex documentation.

- Code Examples
  > Use indented lines (starting with a tab or spaces) for code examples within comments.

- Links
  > Reference other identifiers using their full path (e.g., `[errors.Is]`).

- Deprecation
  > Mark deprecated items with `Deprecated:` followed by migration instructions.

- No Redundancy
  > Avoid repeating the function signature in the documentation.

## 7. Template

Use these templates for documenting Go code. Replace placeholders with actual values.

### 7.1. Package Documentation Template

```go
// SPDX-License-Identifier: Apache-2.0

// Package <name> provides <brief description>.
//
// <Detailed description of the package's purpose and functionality.>
//
// # Usage
//
// <Brief usage example or overview.>
//
//	result, err := <name>.<Function>(args)
//	if err != nil {
//		// handle error
//	}
//
// # Key Functions
//
//   - [Function1]: <brief description>
//   - [Function2]: <brief description>
//
// # Notes
//
// <Any important notes, limitations, or considerations.>
package <name>
```

### 7.2. Function Documentation Template

```go
// <FunctionName> <verb describing action> <what it does>.
//
// <Additional details about behavior, algorithms, or implementation notes.>
//
// Parameters:
//   - param1: <description of first parameter>
//   - param2: <description of second parameter>
//
// Returns:
//   - <type>: <description of return value>
//   - error: <description of error conditions>
//
// Example:
//
//	result, err := <FunctionName>(arg1, arg2)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(result)
func <FunctionName>(param1 <type>, param2 <type>) (<returnType>, error) {
	// implementation
}
```

### 7.3. Type Documentation Template

```go
// <TypeName> represents <what it represents>.
//
// <Additional details about the type's purpose and usage.>
//
// Example:
//
//	var t <TypeName>
//	t.<Method>(args)
type <TypeName> struct {
	// Field1 is <description of field>.
	Field1 <type>

	// Field2 is <description of field>.
	Field2 <type>
}

// <MethodName> <verb describing action> <what it does>.
//
// <Additional details if needed.>
func (t *<TypeName>) <MethodName>() <returnType> {
	// implementation
}
```

### 7.4. Example Function Template

```go
// Example demonstrates basic usage of the <package> package.
func Example() {
	result, err := <Function>(args)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
	// Output: <expected output>
}

// Example<FunctionName> demonstrates usage of <FunctionName>.
func Example<FunctionName>() {
	result, err := <FunctionName>(arg1, arg2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("<format string>", result)
	// Output: <expected output>
}

// Example<FunctionName>_<suffix> demonstrates <specific scenario>.
func Example<FunctionName>_<suffix>() {
	// demonstrate specific use case
	// Output: <expected output>
}
```

## 8. References

- Go [Documentation Comments](https://go.dev/doc/comment) specification.
- Go [Effective Go - Commentary](https://go.dev/doc/effective_go#commentary) guide.
- Go [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) tool documentation.
- Go [pkgsite](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite) tool documentation.
