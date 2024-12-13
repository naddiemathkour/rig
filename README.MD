# HTTP Servers in Go

This directory provides simple steps for creating HTTP servers in Go.

## Overview
The implementations are barebones MVPs for several approaches to creating HTTP 
servers. The goal is to familiarize yourself with the steps required to create 
a backend server in Go. Many principles demonstrated here are transferable to 
other programming languages.

### Focus
The primary focus of these servers is to provide REST API functionality for
CRUD operations in a project.

---

## Getting Started

### Starting the Server
1. **Build the Go program**:  
   Use the `go build` command to compile your code.
2. **Run the executable**:  
   Execute the compiled file to start the server.

---

## Testing the Server

To test the server, use the included `http.sh` file, which contains shell
commands for pinging the running server at `localhost:8080`.

### Steps to Create and Use `http.sh`

1. **Install httpie**:  
  ```bash
    sudo apt install httpie
  ```
2. **Create the `http.sh` file**:
  ```bash
    touch http.sh
  ```
3. **Add a `shebang` at top of the file**:
  ```
    #!/bin/bash
  ```
4. **Add your `httpie` test calls**:
  ```
    http GET http://localhost:8080/
    http GET http://localhost:8080/users
    etc...
  ```
5. **Make the `http.sh` file executable**:
  ```bash
    chmod +x http.sh
  ```
6. **Run the `http.sh` file**:
  ```bash
    ./http.sh
  ```
