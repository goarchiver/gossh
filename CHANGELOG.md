# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.8.2]

### Added

- Subcommand `push` supports timeout for pushing each file/dir to each remote host by flag `--timetout.command`.  
  This feature solves the problem of the entire `gossh` task stuck if the network of a few remote servers is slow.

- Add more help examples for subcommand `push`, `exec`, `script`.

### Fixed

- Fix the problem that if pushing files/dirs fails, the temporary hidden files are not automatically deleted.

## [0.8.1]

### Fixed

Fix the problem of compression ratio of zip for improving files/dirs transmission efficiency.

## [0.8.0]

### Added

- Subcommand 'push' supports copying directories.  
  Also supports push files and directories efficiently at the same time. For efficient transmission, gossh adopts the method of first compressing locally and then decompressing files and directories on the remote server, so the `unzip` command is required on the remote server.

## [0.7.1]

### Changed

Optimize flag `-d/--dest-path` for subcommand `push` and `script`.
If the dest directory given by flag `-d` does not exist or does not have permission to write, output an easy-to-understand error message.

## [0.7.0]

### Added

- Subcommand `push`: keep mode and mtime of dest files and source files the same.

### Security

- For subcommand `push`: For security reasons, if the files to be copied already exists on the target hosts, error messages will be output. If you think it is safe to overwrite the files, you can specify `-F/--force` flag to force overwrite them.

- For subcommand `script`: For security reasons, if the script file already on the target hosts, error messages will
  be output. If you think it is safe to overwrite the script, you can specify `-F/--force` flag to force overwrite it.

## [0.6.1]

### Added

- Subcommand 'push' can push files, not only a file.

## [0.6.0]

### Added

- Provide the subcommand `config` to help users generate configuration file in easy way.

## [0.5.1]

### Added

- Supports three types of ssh tasks.  
  `exec`: Execute commands in remote hosts;  
  `script`: Execute a local script in remote hosts;  
  `push`: Push a local file to remote hosts.

- Supports using sudo to execute the commands or a script as other user(default is `root`).

- Supports specify i18n environment variable value while executing commands or a script to help keep the language of the outputs consistent. For example: zh_CN.UTF-8, en_US.UTF-8.

- Supports four authentication methods.  
  Priority: `ssh-agent authentication` -> `pubkey authentication` -> `password from command flag` -> `username:password from a file`.  
  If the user is not specified, the system environment variable `$USER` will be used by default.

- Supports two methods to specify target hosts. One is through command line arguments, input one or more target hosts, separated by space. The other is through command line flag or configuration file option to specify the hosts file. Both methods can be used at the same time.

- Supports three kinds of timeout:  
  Timeout for connecting each remote host (default `10` seconds);
  Timeout for executing commands or a script on each remote host;
  Timeout for the current gossh task.

- Supports outputting the execution results of ssh to a file or screen or to a file and screen at the same time. Supports specifying the format of output information as json. Supports outputting debug information. Supports silent output.

- High-performance and high-concurrency. You can specify number of concurrent connections (default `1`).

- For ease of use, it supports config file. You can write flags that are not frequently modified into the config file, so you don't need to laboriously specify these flags on the command line. If the flag in both command line and config file, flag that from command line takes precedence over the other. The default config file is: `~/.gossh.yaml`.

### Changed

### Removed

### Fixed

### Security
