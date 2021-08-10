# Backpacker
A small HTML5 Game Wrapper

## How to use the wrapper
The Backpacker is a small HTML5 Game Wrapper. Because some tools like electron and tauri are way overpowered for the usual HTML5 Game, Backpacker packs all files in the public folder into the app. By running the app, Backpacker creates a embedded webserver for localhost on port 3000, also Backpacker uses the webview library to open a webview window to display the content of the embedded webserver in a controllable window.

To build you own Backpack instance you have to follow these steps:
1. Copy you HTML5 Game files into the public folder (you can delete the .gitkeep file)
2. You install all dependencies
3. You build the wrapper for your system (currently there is only linux support)

## Install dependencies
You must have golang installed. 

Also you have to install the pkger dependency by running:

```bash
$ go get github.com/markbates/pkger/cmd/pkger
```

## Build on Linux
To build the wrapper on Linux, simply run make. This will pack the public folder and build the go app

```bash
$ make
```

### Used projects
* https://github.com/markbates/pkger
* https://github.com/webview/webview