# RapidRule – A Fast, TUI-Based Semi-Automatic Firewall Starter for Small Home and Office Environments

**GitHub Repository**: [https://github.com/Kajmany/rapidrule](https://github.com/Kajmany/rapidrule)

---

## Overview

**RapidRule** is a lightweight, terminal user interface (TUI) application written in Go, designed to simplify the process of configuring firewalls in small office and home office (SOHO) environments. It integrates Google's Gemini language model for intelligent configuration analysis and suggestions of various `nftable` configurations and rule management. In addition, it provides real time security posture analysis and alerts the user to suspicious open ports.

RapidRule aims to provide secure and customizable firewall configurations in a fraction of the time typically required, while increasing insight into esoteric `nftable` rulesets.

---

## Key Features

- Realtime port and security posture analysis
- Fast generation of custom firewall templates with secure defaults  
- Semi-automated configuration via an intuitive text-based UI  
- Integration with Gemini for detection of vulnerabilities and misconfigurations  
- Detection of compromised systems and support for automatic mitigation  
- Flexible rule management with support for user-defined blocking criteria  

---

## Motivation

Configuring firewalls remains a pain point for individuals and small teams lacking the resources or expertise to navigate complex tools. RapidRule was built to fill this gap—providing an efficient, semi-automated system that still gives users full visibility and control. It combines the speed and simplicity of preconfigured templates with the intelligence of modern language models to assist in real-time decision-making and evaluation.

---

## How It Works

1. Launch RapidRule in the terminal.  
2. Follow guided prompts or use auto-detected defaults.  
3. Generate and apply a tailored `nftables` rule set.  
4. Optionally run vulnerability analysis powered by Gemini.  
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

- SUI blockchain based threat analysis sharing
- IPv6 configuration and analysis support  
- Rule effectiveness tracking and feedback  
- Real-time log parsing and reactive rule adjustment  
- Template sharing and versioning  
- Expanded Gemini-based recommendation modules

---

## Contributing

We welcome contributions from the community. Whether you're interested in improving firewall templates, expanding detection logic, or enhancing the user experience, visit the [GitHub repository](https://github.com/Kajmany/rapidrule) to get involved. Issues and pull requests are always appreciated.
