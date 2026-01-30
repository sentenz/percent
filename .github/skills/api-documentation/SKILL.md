---
name: api-documentation
description: Adds godoc-compatible documentation comments to Go source files following Go documentation conventions.
metadata:
  version: "1.0"
  activation:
    implicit: true
    priority: 0
    triggers:
      - "godoc"
      - "API docs"
      - "API documentation"
      - "document the API"
      - "add doc comments"
      - "package documentation"
    match:
      languages: ["go", "golang"]
      paths: ["pkg/**/*.go", "internal/**/*.go"]
      prompt_regex: "(?i)(api doc|godoc|document the API|API docs|doc comments|package documentation)"
  usage:
    load_on_prompt: true
    autodispatch: true
---

# API Documentation

Instructions for AI coding agents on adding godoc-compatible documentation comments to Go source files.

> [!NOTE]
> This skill is for documenting Go packages, types, functions, and methods following Go documentation conventions.

- [1. Benefits](#1-benefits)
- [2. Principles](#2-principles)
- [3. Patterns](#3-patterns)
  - [3.1. Package Documentation](#31-package-documentation)
  - [3.2. Type Documentation](#32-type-documentation)
  - [3.3. Function Documentation](#33-function-documentation)
  - [3.4. Method Documentation](#34-method-documentation)
  - [3.5. Constant and Variable Documentation](#35-constant-and-variable-documentation)
  - [3.6. Examples](#36-examples)
- [4. Workflow](#4-workflow)
- [5. Style Guide](#5-style-guide)
- [6. Templates](#6-templates)
  - [6.1. Package Template](#61-package-template)
  - [6.2. Type Template](#62-type-template)
  - [6.3. Function Template](#63-function-template)
  - [6.4. Method Template](#64-method-template)
  - [6.5. Constant/Variable Template](#65-constantvariable-template)
  - [6.6. Example Template](#66-example-template)
- [7. References](#7-references)

## 1. Benefits

- Discoverability
  > Well-documented APIs enable developers to quickly understand and use packages without reading implementation details.

- Maintainability
  > Documentation embedded in source code stays synchronized with implementation, reducing drift between code and documentation.

- Tooling Integration
  > Documentation comments are automatically extracted by `go doc` and displayed in IDEs, editors, and pkg.go.dev.

- Standards Compliance
  > Following Go documentation conventions ensures consistency across the Go ecosystem.

## 2. Principles

Effective API documentation follows these core principles aligned with Go conventions.

- Complete
  > Document all exported identifiers (packages, types, functions, methods, constants, and variables). Unexported identifiers may have documentation for internal clarity.

- Contextual
  > Documentation provides context about usage patterns, performance characteristics, and concurrency safety.

- Consistent
  > Use a uniform style and format throughout following Go documentation conventions defined in this skill.

- Concise
  > Use clear, brief descriptions. The first sentence should be a complete sentence that starts with the name being documented.

- Concrete
  > Provide specific details about behavior, edge cases, and error conditions rather than vague statements.

- Accurate
  > Documentation must match the actual behavior. Update documentation whenever the implementation changes.

- Actionable
  > Include usage examples, preconditions, postconditions, and error handling to help developers use the API correctly.

## 3. Patterns

### 3.1. Package Documentation

Package documentation appears in a comment immediately before the package clause or in a separate `doc.go` file.

- Format
  > Package documentation starts with "Package <name>" followed by a description.

- Content
  > Describes the package's purpose, main concepts, and typical usage patterns.

- Location
  > Can be in any file in the package (typically the main file) or in a dedicated `doc.go` file.

### 3.2. Type Documentation

Type documentation describes the purpose and usage of a type.

- Format
  > Starts with the type name followed by a description.

- Content
  > Explains what the type represents, its responsibilities, and any invariants.

- Related Types
  > May reference related types and their relationships.

### 3.3. Function Documentation

Function documentation describes what the function does and its contract.

- Format
  > Starts with the function name followed by a description.

- Parameters
  > Parameters are described in the main description when needed, not as separate annotations.

- Return Values
  > Return values are described in the main description, including error conditions.

- Errors
  > Document what errors can be returned and under what conditions.

### 3.4. Method Documentation

Method documentation is similar to function documentation.

- Format
  > Starts with the method name followed by a description.

- Receiver
  > The receiver type is mentioned when relevant to understanding the method.

- Context
  > Describes how the method relates to the type's overall functionality.

### 3.5. Constant and Variable Documentation

Constants and variables are documented with their purpose and usage.

- Format
  > Starts with the identifier name followed by a description.

- Grouped Declarations
  > Groups of related constants/variables can share a single comment before the group.

### 3.6. Examples

Examples demonstrate typical usage patterns.

- Format
  > Example functions follow the naming convention `Example`, `Example_suffix`, or `ExampleType_Method`.

- Output
  > Examples can include `// Output:` comments to specify expected output for testing.

- Documentation
  > Examples appear in godoc as executable code samples.

## 4. Workflow

1. Identify

    Identify exported identifiers in `pkg/` or `internal/` that need documentation.

2. Review Existing

    Check if documentation exists and if it needs updating or improvement.

3. Add Documentation

    Add or update documentation comments following the templates and style guide.

4. Verify Format

    Ensure comments follow Go conventions:
    - Start with the identifier name
    - Form complete sentences
    - Use proper punctuation
    - Are placed immediately before the declaration

5. Test with go doc

    Run `go doc <package>` or `go doc <package>.<identifier>` to verify documentation displays correctly.

6. Add Examples (Optional)

    Add example functions to demonstrate typical usage when appropriate.

## 5. Style Guide

- Comment Style
  > Use `//` comment style, not `/* */` for documentation comments.

- First Sentence
  > The first sentence should be a complete, standalone summary that starts with the name being documented.

- Capitalization
  > Start with the identifier name (capitalized as it appears in code), then continue with normal sentence capitalization.

- Punctuation
  > Use proper punctuation. End sentences with periods.

- Blank Lines
  > No blank lines within a documentation comment block.

- Positioning
  > Place the comment immediately before the declaration with no blank line in between.

- Code References
  > Reference other identifiers without special markup (godoc will automatically link them).

- Links
  > Use URLs directly in documentation; godoc will render them as links.

- Formatting
  > Indented text is displayed as preformatted (code blocks).

- Paragraphs
  > Separate paragraphs with blank comment lines (lines containing only `//`).

- Lists
  > Use simple text lists; no special markup is needed.

## 6. Templates

### 6.1. Package Template

Use this template for package documentation. Place in the main package file or in `doc.go`.

```go
// Package <packagename> provides <brief description of package purpose>.
//
// <Extended description of what the package does, main concepts,
// and typical usage patterns.>
//
// Example usage:
//
//	<code example showing typical usage>
//
// <Additional notes about concurrency, performance, or other important
// considerations.>
package <packagename>
```

### 6.2. Type Template

Use this template for documenting types (structs, interfaces, type aliases).

```go
// <TypeName> <describes what the type represents and its purpose>.
//
// <Extended description of responsibilities, invariants, and usage patterns.
// Include information about thread safety, lifecycle, or other important
// characteristics.>
//
// Example:
//
//	<code example showing typical usage of the type>
type <TypeName> struct {
	// <FieldName> <describes the field's purpose>
	<FieldName> <type>
}
```

### 6.3. Function Template

Use this template for documenting functions.

```go
// <FunctionName> <describes what the function does>.
//
// <Extended description including parameter meanings, return value
// descriptions, and any error conditions.>
//
// The function <describes behavior, preconditions, postconditions>.
//
// Returns <description of return value(s)>. If an error occurs,
// it returns <description of error conditions>.
//
// Example:
//
//	<code example showing typical usage>
func <FunctionName>(<params>) (<returns>) {
	// implementation
}
```

### 6.4. Method Template

Use this template for documenting methods.

```go
// <MethodName> <describes what the method does>.
//
// <Extended description including how it relates to the receiver type,
// parameter meanings, return value descriptions, and any error conditions.>
//
// The method <describes behavior, preconditions, postconditions>.
//
// Returns <description of return value(s)>. If an error occurs,
// it returns <description of error conditions>.
func (r <ReceiverType>) <MethodName>(<params>) (<returns>) {
	// implementation
}
```

### 6.5. Constant/Variable Template

Use this template for documenting constants and variables.

```go
// <IdentifierName> <describes the constant/variable and its purpose>.
const <IdentifierName> = <value>

// <GroupDescription> describes this group of related constants.
const (
	// <ConstName1> <describes this specific constant>
	<ConstName1> = <value1>
	
	// <ConstName2> <describes this specific constant>
	<ConstName2> = <value2>
)
```

### 6.6. Example Template

Use this template for creating example functions.

```go
// Example demonstrates basic usage of <PackageName>.
func Example() {
	// Setup
	<setup code>
	
	// Usage
	<example usage code>
	
	// Output:
	// <expected output>
}

// Example<Type>_<method> demonstrates usage of <Type>.<method>.
func Example<Type>_<method>() {
	// Example implementation
	<example code>
	
	// Output:
	// <expected output>
}
```

## 7. References

- [Effective Go - Commentary](https://go.dev/doc/effective_go#commentary)
- [Go Doc Comments](https://go.dev/doc/comment)
- [Godoc: documenting Go code](https://go.dev/blog/godoc)
- [Go Code Review Comments - Doc Comments](https://github.com/golang/go/wiki/CodeReviewComments#doc-comments)
