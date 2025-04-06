# RapidRule – A Fast, TUI-Based Semi-Automatic Firewall Starter for Small Home and Office Environments

<img src="https://github.com/Kajmany/rapidrule/blob/main/RapidRulesLogo.png" width="150" alt="RapidRule logo">

**GitHub Repository**: [https://github.com/Kajmany/rapidrule](https://github.com/Kajmany/rapidrule)



---

## Overview

**RapidRule** is a lightweight, terminal user interface (TUI) application written in Go, designed to simplify the process of configuring firewalls in small office and home office (SOHO) environments. It integrates Google's Gemini language model for intelligent configuration analysis and suggestions of various `nftable` configurations and rule management. In addition, it provides real time security posture analysis and alerts the user to suspicious open ports.

RapidRule aims to provide secure and customizable firewall configurations in a fraction of the time typically required, while increasing insight into esoteric `nftable` rulesets.

---

## Key Features

- Realtime port and security posture analysis
- Fast generation of custom firewall templates with secure templatization  
- Clean and fast Text Based User interface (TUI)
- Integration with Gemini for detection qualification and explanation of vulnerabilities and misconfigurations  
- Detection of compromised systems and support for automatic mitigation  

---

## Motivation

Networking and `nftables` tend to be a rather esoteric subject. All of the information can be confusing and overwhelming. With RapidRule we work to explain vulnerabilities and misconfigurations while keeping a clean UI to avoid user overload.

In addition, configuring firewalls can be a painful task for anyone, but especially individuals and small teams lacking the resources or expertise to navigate complex tools. RapidRule was built to fill this gap—providing an efficient, semi-automated system that gives users full visibility and comphensive suggestions and explanations. It combines the speed and simplicity of preconfigured templates with the intelligence of modern language models to assist in real-time decision-making and evaluation.

---

## How It Works

1. Launch RapidRule in the terminal.  
2. Quickly see open ports and security posture analysis.  
3. Dive deeper into individual ports with automatic analysis
4. Generate and apply a tailored `nftables` ruleset.  
5. Apply suggestions and monitor for compromised system behavior.

All configurations can be previewed, modified, or exported before deployment.

---

## Intended Audience

- IT administrators managing small or home networks  
- Freelancers and independent developers hosting their own infrastructure  
- Security-conscious users who prefer minimal, auditable tools  
- Educational and lab environments requiring reproducible setups

---

## Technical Stack

- **Language**: Go  
- **Firewall Backend**: `nftables`  
- **AI Integration**: Google Gemini API  
- **Interface**: Terminal UI (TUI)  

---

## Roadmap

- Add more features to become a complete security suite
- Complete outbound traffic anaylsis
- SUI blockchain based threat analysis sharing through Walrus
- Virustotal integration for analyzing suspicious processes  
- Create and integrate Ghidra with a Gemini plugin to add code comments and relevent variable names
- Multisystem admin management interface

---

## Contributing

We welcome suggestions and contributions from users and the community. Whether you're interested in improving firewall templates, expanding detection logic, or enhancing the user experience, visit the [GitHub repository](https://github.com/Kajmany/rapidrule) to get involved. Issues and pull requests are always appreciated.
