# Qrepo Package Manager

Qrepo is a lightweight and intuitive package and build manager designed to streamline the process of managing project dependencies and scripts. Whether you're working on a small script or a complex software project, Qrepo provides a simple yet powerful set of commands to enhance your development workflow.

## Installation

### Linux and macOS

Copy the following command on your terminal to smoothly install Qrepo to your system:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/nthnn/Qrepo/master/support/install.sh)"
```

### Windows

For Windows, download the executable file from [here](https://github.com/nthnn/Qrepo/releases).

## Getting Started

1. Initialize your project as a Qrepo repository: `qrepo init`
2. Define scripts in your project's configuration file.
3. Run scripts using: `qrepo run <script-name>`
4. View repository information: `qrepo log`

## Command

1. `init`: Initialize the current directory as a Qrepo repository project.
    Example:
    ```bash
    qrepo init
    ```

2. `run`: Run a defined script in the project folder.
    Example:
    ```bash
    qrepo run build
    ```

3. `log`: Display detailed information on the current repository.
    Example:
    ```bash
    qrepo log
    ```

4. `help`: Show help message for a specific command.
    Example:
    ```bash
    qrepo help
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.