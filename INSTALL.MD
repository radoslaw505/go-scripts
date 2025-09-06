# Installing Go 1.25.1 on Linux

> Download and install Go from the official source.

## Steps

1. **Download the Go tarball**

    ```sh
    wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz
    # Downloads the Go 1.25.1 archive for Linux (64-bit)
    ```

2. **Remove any previous Go installation and extract the archive**

    ```sh
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.25.1.linux-amd64.tar.gz
    # Extracts Go to /usr/local directory
    ```

3. **Set up environment variables**

    Add the following lines to your `~/.profile`, `~/.bashrc`, or `~/.zshrc`:

    ```sh
    export PATH=$PATH:/usr/local/go/bin
    # Adds Go binary to your PATH
    ```

    Then reload your shell configuration:

    ```sh
    source ~/.profile
    # Or use source ~/.bashrc or source ~/.zshrc depending on your shell
    ```

4. **Verify the installation**

    ```sh
    go version
    # Should output: go version go1.25.1 linux/amd64
    ```

> For more details, visit the [official Go installation guide](https://go.dev/doc/install).
