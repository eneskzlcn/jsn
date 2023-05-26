### About Project
jsn is a cli tool that takes a json as string, then formats it. 
After format, the application copies the output to your clipboard while 
printing it to the your console.

This tool designed to work only on MacOS operating system.

### How To Install
After you download the latest binary from releases section, you
can add it to your PATH to use from anywhere in you computer.

#### Ex:
```shell
chmod +x jsn
./jsn '{"name": "sth", age: 53, projects: ["jsn"]}'
```

Then you will see the output on  terminal
```shell

{
    "name": "sth",
    "age": 53,
    "projects": [
        "jsn"
    ]
}

```
and also you can cmd+v or 

```shell
pbpaste
```
to paste the formatted output anywhere.

### How To Use

You need to give your json input inside two `'`{json_data} `'`. Otherwise,
the program will not be able to understand and parse the json.

After you install the latest binary and give execute permission to it, you
can directly use like the below example.