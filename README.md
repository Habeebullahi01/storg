# Storg

This is Storg, a CLI program that helps to sort files in a directory into sub-directories based on their media type and file extensions.

## Use Case

Imagine a directory containing files of various media types; images, audios, videos and so on. Now you'd like to sort the contents  of this directory into sub-directories based on their media types, so you'd have all the images neatly separated in one sub-directory and videos in another sub-directory, and (equally) so on. This tool helps you do exactly that.

## Installation

- Download the binary compatible with your machine from the [releases](https://github.com/Habeebullahi01/storg/releases/) page.
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
        ```sh
        storg sort --srcDir|-s
        ```
        This is used to specify the directory where the files to be sorted are located. Its default value is the working directory from which the command is being used. This flag is optional.
    - tarDir
        -
        ```sh
        storg sort --tarDir|-t
        ```
        This is used to specify the directory in which the sub-directories which will be created should be located. It also defaults to the working directory from which the command is being used. This flag is also optional.

- rename
    - 
    This command allows users to rename all files in a specified directory by numbering them sequentially. Additionally, users can add an optional prefix to the filenames, providing greater flexibility and organization. This feature is particularly useful for managing large sets of files, ensuring consistent and meaningful naming conventions.
    ```sh
    storg rename [flags]
    ```

    ### Flags

    - srcDir(-s)
        -
        ```sh
        storg rename --srcDir|-s
        ```
        This is the directory containing the files to be renamed. If unspecified it defaults to the location from which the command is being used.
    - tarDir(-t)
        -
        ```sh
        storg rename --tarDir|-t
        ```
        This is the directory into which the renamed files should be copied. The specified directory will be created if it does not already exist. If unspecified, it defaults to a sub-directory called 'renamed files' inside the source directory.
    - prefix(-p)
        -
        ```sh
        storg rename --prefix|-p
        ```
        This is an optional prefix to be added to the renamed files for greater organisation.
