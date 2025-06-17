# Universal Project CLI

**Universal Project CLI** is a simple yet flexible Go-based command-line tool that scaffolds out a complete, ready-to-code project structure. Whether you're starting a Python script, a DevOps pipeline, or an infrastructure repo, this CLI gives you a clean, consistent starting point every time.

---

## 🚀 Features

- Creates a full project directory tree with:
  - Source code (`/src`)
  - Modules and reporting logic
  - Tests
  - Infrastructure-as-Code (`/infra`)
  - CI/CD pipelines
  - Scripts and configs
  - Markdown documentation
- Pre-populates essential files like:
  - `README.md`, `.gitignore`, `.env.example`, `LICENSE`
- CLI flags for naming and path control

---

## 🛠 Usage

1. **Build the CLI**
   ```bash
   go build -o ./executables/ ./cli.go

2. **Run the CLI**
    ```bash
    ./executables/cli --pname="my-project-name" --path="./projects"
    or
    ./executables/cli -pname "my-project-name" -path "./projects"
