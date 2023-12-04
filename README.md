# Qrepo Package Manager

![Build Action](https://github.com/nthnn/Qrepo/actions/workflows/build.yml/badge.svg)
[![GitHub Issues](https://img.shields.io/github/issues/nthnn/Qrepo.svg)](https://github.com/nthnn/Qrepo/issues)
[![GitHub Stars](https://img.shields.io/github/stars/nthnn/Qrepo.svg)](https://github.com/nthnn/Qrepo/stargazers)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/nthnn/Qrepo/blob/main/LICENSE)
<a href="https://www.buymeacoffee.com/nthnn"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" height="20px"></a>

Qrepo is a lightweight and intuitive package and build manager designed to streamline the process of managing project dependencies and scripts. Whether you're working on a small script or a complex software project, Qrepo provides a simple yet powerful set of commands to enhance your development workflow.

<p align="center">
    <img src="./assets/screenshot.png" />
</p>

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

## Contributing and Contributors

All contributions are welcome to make Qrepo even better. Whether you want to report a bug, suggest new features, or contribute code, your contributions are highly appreciated.

### Issue Reporting

If you come across a bug, have a feature request, or wish to propose enhancements, we kindly encourage you to initiate the process by opening an issue on our [GitHub Issue Tracker](https://github.com/nthnn/Qrepo/issues).

### Pull Requests

If you're eager to get involved and contribute your coding expertise to Qrepo, thrilled to have you on board! To ensure a smooth and collaborative process, here's an outlined the following steps that you can follow:

1. Fork the Qrepo repository to your GitHub account. And then clone it to your local machine.

    ```bash
    git clone https://github.com/<your username>/Qrepo.git
    ```

2. Create a new branch for your changes:

    ```bash
    git checkout -b feature/<your feature name>
    ```

3. You can now make changes to the repository.
4. Commit your changes:

    ```bash
    git add -A
    git commit -m "Add your meaningful commit message here"
    ```

5. Push your changes to your forked repository:

    ```bash
    git push origin feature/<your feature name>
    ```

6. Create a pull request (PR) from your branch to the main branch of the Qrepo repository.
7. Your PR will be reviewed, and any necessary changes will be discussed and implemented.
8. Once your PR is approved, it will be merged into the main branch, and your contribution will be part of Qrepo.

### Contributors

- [Nathanne Isip](https://github.com/nthnn) — Original Author, Developer
- [Bryan Carpizo](https://github.com/Hirazoem) — He built Qrepo for Windows

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.