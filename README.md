# Catrina
A framework for web apps. Catrina includes his self libraries CSS and javascript. The API search to be easy and make light final files.

In this moment is only available for Linux systems, but you can compile for your OS whit the source in **cmd** directory.

## Install (Linux only)

Execute **install.sh** file in /

```bash
# ./install.sh
```

This script edit the **/etc/bash/bash.bashrc** file and add catrina to $PATH.

## Start Project

```shell
$ catrina new myProject
```

This command make a new directory with the name of the project and the libraries includes in catrina.

### Configuration

The catrina's configuration will read to **catrina.config.json** file. You can set this configuration using the wizard setup was run after `catrina new` command or make yourself the file.

The file's structure is the next (the values are example):

```json
{
  "serverPort": ":9095",
  "inputFileJS": "./src/main.js",
  "inputFileCSS": "./src/styles.css",
  "deployPath": "./deploy",
  "finalFileJS": "main.js",
  "finalFileCSS": "styles.css"
}
```



## Version
1.1.0.6

