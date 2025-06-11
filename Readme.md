# govim

`govim` is a minimal Vim-like text editor written in Go. It runs in the terminal and supports basic editing, saving, and quitting commands.

## Features

* Open and edit files from the terminal
* Insert and command modes
* Basic cursor movement
* Line numbering
* Save files with `:w`
* Save and quit with `:wq`
* Quit without saving with `:q`

## Installation

### Build from Source

```bash
git clone https://github.com/yourusername/govim.git
cd govim
go build -o govim main.go
```

This will generate a binary named `govim` (or `govim.exe` on Windows) in the project directory.

### Windows Installation (Using Prebuilt Executable)

If you don’t want to build from source, you can use the prebuilt executable:

#### Step 1: Download

* Go to the [Releases](https://github.com/adhit-420/govim/releases) section of this repository.
* Download the latest `govim.exe` file.

#### Step 2: Move the Executable (Optional)

Move `govim.exe` to a directory of your choice, e.g.:

```
C:\govim\
```

#### Step 3: Add to System PATH (Optional but Recommended)

To run `govim` from any terminal window:

1. Press `Win + S`, search for **Environment Variables**, and open it.
2. Click on **"Environment Variables..."**
3. Under **System variables**, find and select `Path`, then click **Edit**.
4. Click **New** and add the path where `govim.exe` is located (e.g., `C:\govim`).
5. Click OK to save and close all dialogs.

#### Step 4: Run the Editor

Open **Command Prompt** or **PowerShell**, then run:

```bash
govim filename.txt
```

If the file does not exist, it will be created upon saving.

## Usage

```bash
govim filename.txt
```

## Commands

* `:w` – Save the current file
* `:q` – Quit the editor
* `:wq` – Save and quit

## Development

This project is written in Go using low-level terminal control. It avoids external GUI libraries and works directly in the terminal window.

## Contributing

Contributions are welcome. Please open issues or submit pull requests for new features, bug fixes, or suggestions.

## License

This project is licensed under the MIT License.


Made with ❤️ and ChatGPT