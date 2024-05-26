# Storg

This is Storg, a CLI program that helps to sort files in a directory into sub-directories based on their file extensions.

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