# JIBClaw Project

## Project Overview

**JIBClaw** is a fork of [picoclaw](https://github.com/sipeed/picoclaw.git) that has been rebranded and extended with additional capabilities.

### Origin

- **Source:** Forked from `https://github.com/sipeed/picoclaw.git` (main branch)
- **Rebrand:** picoclaw → JIBClaw

### Custom Enhancements

1. **Rebranding:** Changed all branding from "picoclaw" to "JIBClaw" including:
   - User-facing messages
   - `USER.md`
   - `SOUL`
   - `IDENTITY`

2. **Blockchain Integration:** Added **SKILL** for connecting to **JIBChain V2** blockchain

## Rules

1. **Communication Language:** Communicate and respond in **Thai**
2. **Code Comments:** Write code comments in **English**
3. **Commit Messages:** Write commit messages in **English**
4. **No Commit/Push:** Do not commit or push code without explicit permission

## Upstream Sync Strategy

We will regularly pull updates from `https://github.com/sipeed/picoclaw.git` to keep our codebase current.

### Merge Strategy

To minimize conflicts during merges:
- Keep rebranding changes **isolated** to specific files (messages, `USER.md`, `SOUL`, `IDENTITY`)
- Avoid modifying core logic files when possible
- Use **minimal patches** for rebranding to reduce merge conflicts
- When pulling upstream updates, review and re-apply rebrand patches carefully

## Directory Structure

```
jibclaw/
├── QWEN.md          # Context file for AI assistant (this file)
├── USER.md          # User configuration (rebranded)
├── SOUL             # Agent soul/identity (rebranded)
├── IDENTITY         # Agent identity (rebranded)
├── SKILL/           # Blockchain skills including JIBChain V2
└── ...              # Other picoclaw source files
```

---

*Last updated: March 22, 2026*
