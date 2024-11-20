# OWASP Top 10 for AI Agents

## Overview
This project documents the top 10 security risks specifically related to AI Agents, representing a comprehensive analysis of vulnerabilities unique to autonomous AI systems. The document provides detailed descriptions, examples, and mitigation strategies for each risk, helping organizations secure their AI agent deployments effectively.

## Purpose
As AI agents become increasingly prevalent because GenAI models, understanding and mitigating their security risks becomes crucial. This guide aims to:
- Identify and explain the most critical security risks in AI agent systems
- Provide practical mitigation strategies for each identified risk
- Help organizations implement secure AI agent architectures
- Promote best practices in AI agent security

## Project Structure
The documentation is organized into top ten main security risks, each covering a specific risk category:
1. Agent Authorization and Control Hijacking
2. Agent Critical Systems Interaction
3. Agent Goal and Instruction Manipulation
4. Agent Hallucination Exploitation
5. Agent Impact Chain and Blast Radius
6. Agent Memory and Context Manipulation
7. Agent Orchestration and Multi-Agent Exploitation
8. Agent Resource and Service Exhaustion
9. Agent Supply Chain and Dependency Attacks
10. Agent Knowledge Base Poisoning

# Setup and Installation

## Prerequisites
- dagger
- go
- git
- make
- markdownlint

init dagger
```bash
dagger init --name=owasp-top-10-for-ai-agents  --sdk=go
```

## Development Setup
To set up your development environment:

1. **Install Go**: Ensure you have Go version 1.22 or higher installed.
2. **Install Docker and Docker Compose**: Required for containerization.
3. **Install Dagger CLI**: Follow the official [Dagger installation guide](https://docs.dagger.io/engine/install).
4. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourorg/OWASP-Top-10-for-AI-Agents.git
   cd OWASP-Top-10-for-AI-Agents
   ```

5. **Initialize the Project**:

   ```bash
   make dagger-init
   ```

6. **Run Tests**:

   ```bash
   make test
   ```

## Human-In-The-Loop (HITL)

Critical operations such as release creation and deployment now require manual approval to ensure security and compliance. Only users with the appropriate roles can approve these operations.

## Security and Compliance

We have enhanced the repository to meet enterprise-grade security standards, including:

- **Role-Based Access Control (RBAC)**: Permissions are defined per role in `config.yaml`.
- **Audit Logging**: All operations are logged for compliance purposes.
- **Compliance Standards**: Adherence to ISO27001 and SOC2 standards.

## Contribution Guidelines

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on our code of conduct and the process for submitting pull requests.

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Contributors

### Editors
- **Vishwas Manral**: Initial document framework and early contributions
- **Ken Huang**: Overall editing and conversion of initial document to OWASP format
- **Akram Sheriff**: Orchestration Loop, Planner Agentic security, Multi-modal agentic security
- **Arunesh Salhotra**: Technical review and content organization

### Authors
- **Anton Chuvakin**: DoS and Capitalize overfitting sections
- **Akram Sheriff**: Planner security, Orchestration Loop, Multi-modal agentic security, Confused Deputy
- **Aradhna Chetal**: Agent Supply Chain
- **Ken Huang**: Document structure and OWASP standardization
- **Raj Badhwar**: Capitalize Agentic Overfitting, Model extraction
- **Govindraj Palanisamy**: Alignment of sections to OWASP TOP 10 Mapping, Threat Mapping
- **Mateo Rojas-Carulla**: Data poisoning at scale from untrusted sources, Overreliance and lack of oversight
- **Matthias Kraft**: Data poisoning at scale from untrusted sources, Overreliance and lack of oversight
- **Royce Lu**: Stealth Propagation Agent Threats, Agent Memory Exploitation
- **Anatoly Chikanov**: Technical contributions
- **Alex Schulman-Peleg**: Security analysis
- **Alok Talgaonkar**: Content review
- **Sunil Arora**: Technical input
- **S M Zia Ur Rashid**: Content contributions

## Organizational Support
This project has been made possible through the support and contributions of professionals from leading organizations including:
- Jacobs
- Cisco Systems
- GSK
- Palo Alto Networks
- Precize
- Lakera
- EY
- Google
- Humana
- GlobalPayments
- TIAA

## License
This project is part of OWASP and follows OWASP's licensing terms.

## How to Contribute
We welcome contributions from the security community. Please see our contribution guidelines for more information on how to participate in this project.

## Contact
For questions, suggestions, or concerns, please open an issue in this repository or contact the project maintainers.

## Acknowledgments
Special thanks to all contributors who have dedicated their time and expertise to make this project possible, and to the organizations that have supported their participation in this important security initiative.

---

*This document is maintained by the OWASP community and represents a collaborative effort to improve security in AI agent systems.*
