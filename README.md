# api-scenario-runner

ScenarioRunner is a CLI tool for **scenario-based API testing and validation**.

Instead of testing individual API requests, you define an end-to-end **user flow** in YAML (steps + rules), then run it locally or in CI.

## Status

This repository currently contains a **runnable skeleton (stubs)**:

- **CLI**: `scenario version`, `scenario validate`, `scenario run`
- **Scenario loader**: parses YAML into internal structs
- **Engine**: executes steps sequentially (supports `repeat`)
- **Reporter**: prints a console report

HTTP execution and rule evaluation are currently stubbed (no network calls; rules always pass).

## Requirements

- Go **1.22+**

## Quickstart

From the repo root:

```bash
go run ./cmd/scenario version
go run ./cmd/scenario validate example.yaml
go run ./cmd/scenario run example.yaml
```

## Build

```bash
go build -o scenario ./cmd/scenario
```

Then run:

```bash
./scenario version
./scenario validate example.yaml
./scenario run example.yaml
```

## CLI Usage

### Validate a scenario

```bash
scenario validate <file>
```

### Run a scenario

```bash
scenario run <file>
```

### Print version

```bash
scenario version
```

## Scenario YAML example

See `example.yaml` in the repo root.

```yaml
scenario: cart-discount

vars:
  base_url: https://api.shop.com

steps:
  - name: login
    request:
      method: POST
      url: ${base_url}/login
      body:
        username: demo
        password: demo

  - name: add_item
    repeat: 10
    request:
      method: POST
      url: ${base_url}/cart/items
      body:
        product_id: 123

rules:
  - name: discount_rule
    if: cart_items >= 10
    expect:
      discount_applied: true
```

## Example output

```text
Running scenario: cart-discount

Step login .......... OK (200)
Step add_item .......... OK (200)
...

Rule discount_rule .. PASS

Scenario result: SUCCESS
```

## Notes / current limitations

- `internal/executor` currently returns stubbed results (does not call the network).
- `internal/rules` currently marks all rules as PASS.
- Variable interpolation (e.g. `${base_url}`) is not implemented yet.
