\# ScenarioRunner – Scenario-Based API Testing Engine



\*\*Product \& Technical Design Document (MVP)\*\*



\---



\# 1. Overview



\*\*ScenarioRunner\*\* is an open-source CLI tool for \*\*scenario-based API testing and validation\*\*.

Instead of testing individual API requests (as traditional API tools do), ScenarioRunner allows developers to describe \*\*complete user flows\*\* and validate \*\*system behavior through rules\*\*.



The system executes sequences of API calls, maintains runtime state, and evaluates \*\*business rules\*\* on the results.



Example scenario:



```

User logs in

User adds 10 items to cart

System should apply discount

```



The tool is designed to:



\* Run locally on developer machines

\* Run inside CI/CD pipelines

\* Execute complex API workflows

\* Validate business rules

\* Produce structured test reports



\---



\# 2. Goals



\## Primary Goals



1\. \*\*Scenario-first API testing\*\*

2\. \*\*Business rule validation\*\*

3\. \*\*CLI-first developer workflow\*\*

4\. \*\*CI/CD integration\*\*

5\. \*\*Single executable binary\*\*



\## Non-Goals (MVP)



The MVP intentionally excludes:



\* GUI

\* Postman collection import

\* OAuth flows

\* GraphQL support

\* API mocking

\* Visual dashboards



These can be added after the core engine is stable.



\---



\# 3. Core Concept



Most tools focus on \*\*individual API calls\*\*.



ScenarioRunner focuses on \*\*system behavior\*\*.



Traditional approach:



```

POST /login

POST /cart/items

GET /cart

```



ScenarioRunner approach:



```

Scenario: checkout-flow



login

add item x10

verify discount rule

```



\---



\# 4. Example Scenario File



Scenario files are written in \*\*YAML\*\*.



Example:



```yaml

scenario: cart-discount



vars:

&#x20; base\_url: https://api.shop.com



steps:



&#x20; - name: login

&#x20;   request:

&#x20;     method: POST

&#x20;     url: ${base\_url}/login

&#x20;     body:

&#x20;       username: demo

&#x20;       password: demo



&#x20; - name: add\_item

&#x20;   repeat: 10

&#x20;   request:

&#x20;     method: POST

&#x20;     url: ${base\_url}/cart/items

&#x20;     body:

&#x20;       product\_id: 123



rules:



&#x20; - name: discount\_rule

&#x20;   if: cart\_items >= 10

&#x20;   expect:

&#x20;     discount\_applied: true

```



\---



\# 5. CLI Usage



Run a scenario:



```

scenario run cart-discount.yaml

```



Validate syntax:



```

scenario validate cart-discount.yaml

```



Output example:



```

Running scenario: cart-discount



Step login .......... OK (200)

Step add\_item x10 ... OK



Rule discount\_rule .. PASS



Scenario result: SUCCESS

```



\---



\# 6. System Architecture



The system contains \*\*five core modules\*\*.



```

CLI

&#x20;│

&#x20;├── Scenario Loader

&#x20;│

&#x20;├── Scenario Engine

&#x20;│

&#x20;├── HTTP Executor

&#x20;│

&#x20;├── Rule Engine

&#x20;│

&#x20;└── Reporter

```



\---



\# 7. Module Design



\## 7.1 CLI Module



Responsible for:



\* command parsing

\* command execution

\* user interaction



Commands:



```

scenario run <file>

scenario validate <file>

scenario version

```



Suggested library:



```

cobra

```



Directory:



```

cmd/

&#x20;  root.go

&#x20;  run.go

&#x20;  validate.go

```



\---



\## 7.2 Scenario Loader



Responsible for:



\* loading YAML

\* validating structure

\* mapping YAML → internal structs



Directory:



```

internal/scenario

```



Data model:



```

Scenario

&#x20;├── Name

&#x20;├── Variables

&#x20;├── Steps

&#x20;└── Rules

```



Example struct:



```

type Scenario struct {

&#x20;   Name string

&#x20;   Vars map\[string]string

&#x20;   Steps \[]Step

&#x20;   Rules \[]Rule

}

```



\---



\## 7.3 Scenario Engine



The \*\*orchestration layer\*\*.



Responsibilities:



\* execute steps sequentially

\* manage runtime state

\* store responses

\* handle loops (repeat)

\* call rule engine



Execution flow:



```

Load Scenario

Initialize Context

Execute Steps

Evaluate Rules

Generate Report

```



Directory:



```

internal/engine

```



Runtime context example:



```

Context

&#x20;├── Variables

&#x20;├── StepResults

&#x20;└── RuntimeState

```



\---



\## 7.4 HTTP Executor



Responsible for making HTTP requests.



Responsibilities:



\* build request

\* send request

\* capture response

\* measure latency



Directory:



```

internal/executor

```



Output structure:



```

StepResult

&#x20;├── StatusCode

&#x20;├── Headers

&#x20;├── Body

&#x20;└── Duration

```



Uses Go standard library:



```

net/http

```



\---



\## 7.5 Rule Engine



This module evaluates business rules.



Example rule:



```

cart\_items >= 10

```



Or:



```

response.status == 200

```



Responsibilities:



\* parse rule expressions

\* evaluate conditions

\* return PASS / FAIL



Directory:



```

internal/rules

```



Possible implementation:



```

expression evaluator

```



Later improvements:



\* complex expressions

\* logical operators

\* rule actions



\---



\## 7.6 Reporter



Responsible for output.



Types of output:



```

console

JSON

```



Example console report:



```

Scenario: cart-discount



Step login ............ OK

Step add\_item x10 ..... OK



Rule discount\_rule .... PASS



Result: SUCCESS

```



Directory:



```

internal/report

```



\---



\# 8. Project Structure



Suggested Go project layout:



```

scenario-runner/



cmd/

&#x20;  root.go

&#x20;  run.go



internal/



&#x20;  scenario/

&#x20;     loader.go

&#x20;     model.go



&#x20;  engine/

&#x20;     runner.go



&#x20;  executor/

&#x20;     http.go



&#x20;  rules/

&#x20;     engine.go



&#x20;  report/

&#x20;     console.go

```



\---



\# 9. Development Roadmap



\## Phase 1 – Core CLI



Tasks:



\* initialize Go project

\* implement CLI commands

\* implement scenario loader

\* validate YAML structure



Estimated time: \*\*2 days\*\*



\---



\## Phase 2 – Scenario Engine



Tasks:



\* create scenario execution engine

\* add runtime context

\* implement step loop

\* store results



Estimated time: \*\*3–4 days\*\*



\---



\## Phase 3 – HTTP Execution



Tasks:



\* implement HTTP client

\* build request builder

\* capture response

\* measure duration



Estimated time: \*\*2–3 days\*\*



\---



\## Phase 4 – Rule Engine



Tasks:



\* implement rule model

\* implement condition evaluator

\* connect engine → rule engine

\* return PASS / FAIL



Estimated time: \*\*3–4 days\*\*



\---



\## Phase 5 – Reporting



Tasks:



\* console output formatter

\* JSON output

\* error handling



Estimated time: \*\*2 days\*\*



\---



\# 10. Initial Task List



\## Project Setup



\* \[ ] Create GitHub repository

\* \[ ] Initialize Go module

\* \[ ] Add CLI framework

\* \[ ] Create project directory structure



\---



\## Scenario Module



\* \[ ] Define scenario model structs

\* \[ ] Implement YAML loader

\* \[ ] Validate scenario schema



\---



\## Engine Module



\* \[ ] Implement scenario runner

\* \[ ] Implement step execution loop

\* \[ ] Implement variable store

\* \[ ] Implement repeat step logic



\---



\## HTTP Executor



\* \[ ] Implement request builder

\* \[ ] Implement HTTP client

\* \[ ] Capture response data

\* \[ ] Measure response time



\---



\## Rule Engine



\* \[ ] Define rule struct

\* \[ ] Implement condition evaluator

\* \[ ] Implement result comparison

\* \[ ] Integrate rule engine with scenario engine



\---



\## Reporter



\* \[ ] Implement console report

\* \[ ] Implement JSON output

\* \[ ] Add scenario summary



\---



\# 11. First Milestone



\*\*Goal:\*\* Working CLI scenario execution.



Command:



```

scenario run example.yaml

```



System should:



1\. Load scenario

2\. Execute steps

3\. Evaluate rules

4\. Print results



\---



\# 12. Future Features (Post-MVP)



Potential roadmap:



\### Simulation Engine



```

simulate 100 users

```



\### Parallel Execution



```

parallel steps

```



\### API Authentication



```

OAuth

JWT

API Keys

```



\### Scenario Graph Visualization



Generate flow diagrams.



\### Web UI



Optional web interface.



\---



\# 13. Success Criteria



The MVP is successful if:



\* Developers can define API scenarios in YAML

\* Scenarios execute reliably

\* Rules validate system behavior

\* Tool runs in CI/CD pipelines

\* Single binary distribution works



\---



\# 14. License



Recommended:



```

MIT License

```



Encourages open source adoption.



\---



\# 15. Summary



ScenarioRunner introduces a \*\*scenario-based model for API testing\*\*.



Key innovations:



\* behavior-driven API scenarios

\* rule-based validation

\* CLI-first architecture

\* CI/CD friendly design



The MVP can realistically be implemented in \*\*2–3 weeks\*\* with a focused development scope.



\---



