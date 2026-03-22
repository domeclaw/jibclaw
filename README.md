---

<div align="center">
  <img src="assets/jibclaw-logo.png" alt="JIBClaw" width="512">

  <h1>JIBClaw: PicoClaw Fork</h1>

  <h3>Ultra-Efficient AI Assistant with JIBChain V2 Integration</h3>

</div>

---

> **JIBClaw** is a fork of PicoClaw by JIB Computer. It is written entirely in **Go** with added blockchain integration for JIBChain V2.

🦐 JIBClaw is an ultra-lightweight personal AI Assistant, refactored from PicoClaw and enhanced with JIBChain V2 blockchain connectivity.

⚡️ Runs on <10MB RAM with native JIB token support for gas fees.

## 📦 JIBClaw Editions

JIBClaw comes in three editions to fit your needs:

### 1. **JIBClaw Desktop** (Windows / macOS)
Install directly on your Windows or macOS machine for local AI assistant capabilities.

### 2. **JIBClaw Box**
Dedicated hardware appliance for 24/7 AI assistant operations.

![JIBClaw Box](assets/jibbox.png)

- Plug-and-play setup
- Always-on AI assistant
- Local JIBChain V2 node integration
- Perfect for home or office use

### 3. **JIBClaw Cloud Edition**
Hosted solution for enterprises and teams.

- Managed infrastructure
- Multi-user support
- Team collaboration features
- Enterprise-grade security

## ✨ Features

🪶 **Ultra-Lightweight**: <10MB Memory footprint

⚡️ **Lightning Fast**: Boot in <1 second

🔌 **MCP Support**: Native Model Context Protocol integration

🧠 **Smart Routing**: Rule-based model routing

🔗 **JIBChain V2**: Native blockchain integration with JIB token for gas fees

## 🔗 JIBChain V2 Integration

JIBClaw includes built-in support for JIBChain V2 blockchain:

- **JIB Token**: Use JIB as gas for blockchain transactions
- **Smart Contracts**: Interact with JIBChain V2 smart contracts
- **Wallet Integration**: Built-in wallet management
- **Skill System**: Extend capabilities via blockchain-based skills

See `SKILL/` directory for available blockchain skills.

## 🚀 Quick Start

### 1. Initialize

```bash
jibclaw onboard
```

### 2. Configure

Edit `~/.jibclaw/config.json`:

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.jibclaw/workspace",
      "model_name": "gpt-4",
      "max_tokens": 8192,
      "temperature": 0.7
    }
  },
  "model_list": [
    {
      "model_name": "gpt-4",
      "model": "openai/gpt-4",
      "api_key": "your-api-key"
    }
  ],
  "blockchain": {
    "jibchain": {
      "enabled": true,
      "network": "mainnet",
      "wallet_path": "~/.jibclaw/wallet"
    }
  }
}
```

### 3. Chat

```bash
jibclaw agent -m "What is 2+2?"
```

## 💬 Supported Channels

Connect JIBClaw to your favorite messaging platforms:

- **Telegram** - Bot integration
- **Discord** - Full server support
- **WhatsApp** - Native connection
- **WeChat** - Personal account integration
- **LINE** - Official account support
- **WeCom** - Enterprise AI Bot

## 🖥️ CLI Reference

| Command                   | Description                    |
| ------------------------- | ------------------------------ |
| `jibclaw onboard`         | Initialize config & workspace  |
| `jibclaw agent -m "..."`  | Chat with the agent            |
| `jibclaw agent`           | Interactive chat mode          |
| `jibclaw gateway`         | Start the gateway              |
| `jibclaw status`          | Show status                    |
| `jibclaw version`         | Show version info              |
| `jibclaw skills list`     | List installed skills          |
| `jibclaw skills install`  | Install a skill                |
| `jibclaw auth login`      | Authenticate with providers    |

## 📚 Documentation

For detailed guides, see the `docs/` directory.

## 🤝 Contribute

PRs welcome! JIBClaw is open source under MIT License.

---

<div align="center">
  <p><strong>JIBClaw</strong> - Powered by <strong>JIB Computer</strong></p>
  <p><em>Building the future of AI + Blockchain</em></p>
</div>
