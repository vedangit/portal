# Portal 🚀

**Portal is a command-line tool that "learns" your project setup so you can jump into any development context instantly.**

It remembers your project locations and setup commands, allowing you to switch between complex environments with a single command, from anywhere on your system.



---

## Features

* **🧠 Smart Project Discovery**: Portal learns where your projects are located simply by using it once with `portal init`.
* **⚡️ One-Command Setup**: Run multi-step environment setups (e.g., activating virtualenvs, starting Docker, setting Node versions) with a single command like `portal enter my-project`.
* **💨 Effortless Context Switching**: Jump between any of your registered projects from any directory in your terminal.
* **🧹 Clean Teardown**: Gracefully shut down services and run cleanup scripts with `portal leave`.
* **✅ Zero-Config by Default**: No need to create manual configuration files. The tool adapts to your workflow.

---

## Installation

The installation is a simple, one-time process. Choose your operating system below for detailed instructions.

<details>
<summary><strong> macOS Installation</strong></summary>

### 1. Download the Binary

Go to the [**Latest Release Page**](https://github.com/vedangit/portal/releases/latest) and download the file for your Mac's architecture:
* `portal-cli_darwin_arm64` for Apple Silicon (M1/M2/M3).
* `portal-cli_darwin_amd64` for Intel-based Macs.

### 2. Install the Binary

Open your terminal and run the following commands to make the tool available everywhere on your system. This example is for an Apple Silicon Mac.

```bash
# Navigate to your Downloads folder
cd ~/Downloads

# Rename the file to just 'portal-cli'
mv portal-cli_darwin_arm64 portal-cli

# Make the file executable
chmod +x portal-cli

# Move it to a location in your system's PATH
sudo mv portal-cli /usr/local/bin/
```
*(You may be prompted for your administrator password for the `sudo` command.)*

### 3. Configure Your Shell

This is the final, one-time step to give Portal permission to change your directory.

```bash
# Add the required function to your Zsh configuration file
echo '
# Portal: The Project Context Switcher
portal() {
  case "$1" in
    enter|leave)
      local script_to_run=$(portal-cli "$@")
      if [ $? -eq 0 ]; then eval "$script_to_run"; fi;;
    *)
      portal-cli "$@";;
  esac
}
' >> ~/.zshrc

# Reload your shell to activate the command
source ~/.zshrc
```
**Done!** The `portal` command is now fully installed.

</details>

<details>
<summary><strong>🐧 Linux Installation</strong></summary>

### 1. Download the Binary

Go to the [**Latest Release Page**](https://github.com/vedangit/portal/releases/latest) and download the `portal-cli_linux_amd64` file. You can do this directly from your terminal with `curl`.

```bash
curl -L -o portal-cli_linux_amd64 [https://github.com/vedangit/portal/releases/download/v1.0.0/portal-cli_linux_amd64](https://github.com/vedangit/portal/releases/download/v1.0.0/portal-cli_linux_amd64)
# Note: Replace the URL with the link to your specific release file.
```

### 2. Install the Binary

Next, make the file executable and move it to a standard `PATH` directory so it can be run from anywhere.

```bash
# Make the file executable
chmod +x portal-cli_linux_amd64

# Move it to /usr/local/bin and rename it
sudo mv portal-cli_linux_amd64 /usr/local/bin/portal-cli
```
*(You will likely be prompted for your password for the `sudo` command.)*

### 3. Configure Your Shell

This is the final, one-time step. Add the required helper function to your shell's configuration file (e.g., `.bashrc` or `.zshrc`).

```bash
# Add the required function to your Bash configuration file
echo '
# Portal: The Project Context Switcher
portal() {
  case "$1" in
    enter|leave)
      local script_to_run=$(portal-cli "$@")
      if [ $? -eq 0 ]; then eval "$script_to_run"; fi;;
    *)
      portal-cli "$@";;
  esac
}
' >> ~/.bashrc

# Reload your shell to activate the command
source ~/.bashrc
```
**Done!** The `portal` command is now fully installed.

</details>

<details>
<summary><strong>❖ Windows Installation</strong></summary>

### 1. Download the Binary

Go to the [**Latest Release Page**](https://github.com/vedangit/portal/releases/latest) and download the `portal-cli_windows_amd64.exe` file.

### 2. Prepare the System `PATH`

We'll ensure you have a dedicated folder for command-line tools and that it's in your system's `PATH`.

Open **PowerShell as an Administrator** and run these commands:

```powershell
# Create a folder for your scripts if it doesn't exist
if (-not (Test-Path -Path $HOME\Scripts)) {
    New-Item -Path $HOME\Scripts -ItemType Directory
}

# Move the downloaded file from your Downloads folder and rename it
Move-Item -Path $HOME\Downloads\portal-cli_windows_amd64.exe -Destination $HOME\Scripts\portal-cli.exe

# Add your Scripts folder to the User PATH (this is a permanent change)
$userPath = [System.Environment]::GetEnvironmentVariable('PATH', 'User')
if (-not ($userPath -like "*$HOME\Scripts*")) {
    $newPath = "$userPath;$HOME\Scripts"
    [System.Environment]::SetEnvironmentVariable('PATH', $newPath, 'User')
    Write-Host "PATH updated. Please restart your terminal to apply changes."
}
```
**Important**: You must **close and reopen** your PowerShell terminal after this step for the `PATH` changes to take effect.

### 3. Configure Your PowerShell Profile

This is the final, one-time step. You'll add a helper function to your PowerShell profile (the equivalent of `.bashrc`).

```powershell
# Create a profile if you don't have one
if (-not (Test-Path -Path $PROFILE)) {
    New-Item -Path $PROFILE -ItemType File -Force
}

# Add the Portal function to your profile
$function_code = @"

# Portal: The Project Context Switcher
function portal {
    switch (`$args[0]) {
        'enter' {
            `$script = portal-cli.exe `$args
            if (`$LASTEXITCODE -eq 0) { Invoke-Expression `$script }
        }
        'leave' {
            `$script = portal-cli.exe `$args
            if (`$LASTEXITCODE -eq 0) { Invoke-Expression `$script }
        }
        default {
            portal-cli.exe `$args
        }
    }
}
"@
Add-Content -Path $PROFILE -Value $function_code

# Reload your profile to activate the command
. $PROFILE
```
**Done!** The `portal` command is now fully installed in PowerShell.

</details>

---

## Usage

Portal works by learning your projects. You only ever have to teach it about a project once.

### 1. Teaching Portal a New Project

For the very first time you want to use Portal with a project, navigate to that project's root directory and run the interactive setup.

```bash
# Navigate to your project folder
cd ~/dev/my-cool-project

# Run the interactive init command
portal init
```
The tool will ask you for the setup commands you want to run every time you enter this project. It then automatically remembers the project's name (`my-cool-project`) and its location, and creates a `.portal.toml` file for you.

### 2. Entering and Leaving a Project

Once a project has been initialized, you can enter it from **anywhere** on your system.

```bash
# From any directory...
portal enter my-cool-project

# Your terminal is now inside ~/dev/my-cool-project and your setup commands have run.
# To exit, simply run:
portal leave
```

### Example `.portal.toml`

Here is an example of the configuration file that `portal init` creates for a Python project.

```toml
[enter]
# Commands to run when you enter the project
commands = [
  "pyenv local 3.11.4",
  "source .venv/bin/activate",
  "docker-compose up -d postgres-db"
]
message = "Python 3.11.4 venv activated. PostgreSQL container is running."

[leave]
# Commands to run when you leave
commands = [
  "docker-compose down"
]
```
You can manually edit this file at any time to add or change commands.