# Storg

This is Storg, a CLI program that helps to sort files in a directory into sub-directories based on their file extensions.

## Installation

- Download the binary compatible with your machine from the release page.
- Add the download location to the system path variable, or move the binary to a location already in the path.
- Rename the binary to _storg_ for convenience.

## Usage

```sh
storg [command] [flags]
```

## Commands

- sort
    - 
    This is the command that is used to sort files based on their extensions.

    ```sh
    storg sort [flags]
    ```

    ### Flags

    - srcDir
        -
        This is used to specify the directory where the files to be sorted are located. Its default value is the working directory from which the command is being used. This flag is optional.
    - tarDir
        -
        This is used to specify the directory in which the sub-directories which will be created should be located. It also defaults to the working directory from which the command is being used. This flag is also optional.