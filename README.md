<!-- SPDX-License-Identifier: CC-BY-4.0 -->
# slip CLI

> A minimal, trustless, system-wide package manager driven by a single declarative config file. Written in Go, configured with Lua.

---

## 📌 Overview

`slip` is a declarative package and configuration manager for people who:
- Want full control of their system state
- Don't want to learn an obscure DSL 
- Hate lockfiles, YAML, and emotional tooling

It installs and removes packages based on your `slip.lua` — and nothing else.

---

## 💡 Philosophy

- **`slip.lua` is the source of truth** — all state is declared there
- **No runtime config edits** — the CLI does not mutate `slip.lua`
- **No daemons, no registries, no metadata** — just logic and files
- **Predictable by default** — reproducible infra without ceremony

---

## 🔧 Installation

```sh
curl -sSL https://slip.run/install.sh | sh
```

Or build manually:

```sh
git clone https://github.com/slippkg/cli
cd cli
go build -o slip ./cmd/slip
```

Requires Go 1.21+

---

## 🚀 Commands

### `slip sync`
Read `slip.lua`, install packages that are missing, and remove packages that are present but not declared.

Always verbose.

```sh
slip sync
```

Example output:

```text
[+] Installing: starship (v1.18.2)
[-] Removing: htop (not declared)
[✓] ripgrep already installed (v14.0.0)
```

---

### `slip sync --dry-run`
Preview what would change without touching the system.

```sh
slip sync --dry-run
```

```text
[+] Would install: starship (v1.18.2)
[-] Would remove: htop
[✓] ripgrep already installed (v14.0.0)
No changes made (dry-run)
```

---

### `slip up`
Check which packages have updates available.
Does not install or modify anything.

```sh
slip up
```

Example output:

```text
starship   v1.16.0 → v1.18.2   (tag: v1.18.2)
ripgrep    v13.0.0 → v14.0.0   (tag: v14.0.0)
```

---

## 📁 Structure

| Directory      | Purpose                          |
|----------------|----------------------------------|
| `cmd/slip/`    | CLI entrypoint and flags         |
| `core/`        | Install/uninstall logic          |
| `parser/`      | Lua config reading               |
| `internal/`    | Utilities, helpers               |

All package state is managed via `.slip/bin/` in the user's home directory.

---

## 📜 License

Licensed under **AGPL-3.0**.  
If you commercialize it, publish your source — or don’t use it.

---

## 🌐 Related Projects

- [slippkg/pkgs](https://github.com/slippkg/pkgs) – Shared slipfiles (optional package install scripts)
- [slippkg/docs](https://github.com/slippkg/docs) – Documentation and usage patterns

---

## 📬 Contact

Issue reports or ideas? Open an issue or reach out:  
📧 `hello [at] slip [dot] run`


