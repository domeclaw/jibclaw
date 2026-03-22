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
├── AGENTS.md        # Context file for AI assistant (this file)
├── USER.md          # User configuration (rebranded)
├── SOUL             # Agent soul/identity (rebranded)
├── IDENTITY         # Agent identity (rebranded)
├── SKILL/           # Blockchain skills including JIBChain V2
└── ...              # Other picoclaw source files
```

## Rebranding Changes

This section documents all files that need to be modified when syncing from upstream picoclaw. These changes rebrand "picoclaw" to "jibclaw".

### 1. CLI Binary Name & Banner

**File:** `cmd/picoclaw/main.go`

| Line(s) | Original | Changed To |
|---------|----------|------------|
| Header comment | `PicoClaw - Ultra-lightweight...` | `JIBClaw - Ultra-lightweight...` |
| Copyright | `PicoClaw contributors` | `JIBClaw contributors` |
| Function name | `NewPicoclawCommand()` | `NewJibclawCommand()` |
| Short description | `"picoclaw - Personal AI Assistant"` | `"jibclaw - Personal AI Assistant"` |
| Use command | `Use: "picoclaw"` | `Use: "jibclaw"` |
| Example | `Example: "picoclaw version"` | `Example: "jibclaw version"` |
| Function call | `NewPicoclawCommand()` | `NewJibclawCommand()` |
| Banner (line 57-63) | `PICOCLAW` ASCII art | `JIBCLAW` ASCII art |

**ASCII Art Banner Change:**
```go
// Original (PICOCLAW):
colorBlue + "██████╗ ██╗ ██████╗ ██████╗ " + colorRed + " ██████╗██╗      █████╗ ██╗    ██╗\n" +
colorBlue + "██╔══██╗██║██╔════╝██╔═══██╗" + colorRed + "██╔════╝██║     ██╔══██╗██║    ██║\n" +
...

// Changed to (JIBCLAW):
colorBlue + "     ██╗██╗██████╗  ██████╗" + colorRed + "██╗      █████╗ ██╗    ██╗\n" +
colorBlue + "     ██║██║██╔══██╗██╔════╝" + colorRed + "██║     ██╔══██╗██║    ██║\n" +
...
```

### 2. Version Command Output

**File:** `cmd/picoclaw/internal/version/command.go`

| Line | Original | Changed To |
|------|----------|------------|
| 26 | `fmt.Printf("%s picoclaw %s\n", ...)` | `fmt.Printf("%s jibclaw %s\n", ...)` |

### 3. Web Frontend Rebranding

**Files:** `web/frontend/*`

#### 3.1 Page Title
**File:** `web/frontend/index.html`

| Line | Original | Changed To |
|------|----------|------------|
| 11 | `<title>PicoClaw</title>` | `<title>JIBClaw</title>` |

#### 3.2 English Translations
**File:** `web/frontend/src/i18n/locales/en.json`

| Line | JSON Path | Change |
|------|-----------|--------|
| 342 | `agent.skills.description` | `PicoClaw` → `JIBClaw` |
| 443 | `settings.autostart_hint` | `PicoClaw Web` → `JIBClaw Web` |
| 447 | `settings.server_port_hint` | `PicoClaw Web` → `JIBClaw Web` |

#### 3.3 Chinese Translations
**File:** `web/frontend/src/i18n/locales/zh.json`

| Line | JSON Path | Change |
|------|-----------|--------|
| 342 | `agent.skills.description` | `PicoClaw` → `JIBClaw` |
| 443 | `settings.autostart_hint` | `PicoClaw Web` → `JIBClaw Web` |
| 447 | `settings.server_port_hint` | `PicoClaw Web` → `JIBClaw Web` |

### 4. Workspace Configuration Files

**Files:** `workspace/*.md` (already updated)

- `workspace/IDENTITY.md` - Name: JIBClaw 🦞
- `workspace/SOUL.md` - "I am jibclaw..."
- `workspace/USER.md` - Name: JIBClaw

### Post-Sync Checklist

When pulling updates from upstream, run through this checklist:

1. [ ] Check if `cmd/picoclaw/main.go` was modified by upstream
   - If yes, re-apply all rebranding changes from section 1 above
2. [ ] Check if `cmd/picoclaw/internal/version/command.go` was modified
   - If yes, re-apply change from section 2 above
3. [ ] Check if `web/frontend/index.html` was modified
   - If yes, re-apply `<title>JIBClaw</title>`
4. [ ] Check if translation files were modified
   - `web/frontend/src/i18n/locales/en.json` - Re-apply 3 strings
   - `web/frontend/src/i18n/locales/zh.json` - Re-apply 3 strings
5. [ ] Verify `pkg/auth/oauth.go` exists (was accidentally deleted once)
4. [ ] Build binaries and test:
   ```bash
   go build -ldflags="-s -w" -o jibclaw ./cmd/picoclaw
   CGO_ENABLED=1 go build -ldflags="-s -w" -o jibclaw-launcher ./web/backend
   ./jibclaw version
   ./jibclaw --help
   ```

### Build Commands

After syncing and applying rebranding patches:

```bash
# Generate embedded files
go generate ./...

# Build main binary
go build -ldflags="-s -w" -o jibclaw ./cmd/picoclaw

# Build launcher (requires frontend build first)
cd web/frontend && pnpm install && pnpm build:backend
cd ../..
CGO_ENABLED=1 go build -ldflags="-s -w" -o jibclaw-launcher ./web/backend
```

---

*Last updated: March 22, 2026 (added web/frontend rebranding)*
