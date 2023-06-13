**REPO STATEMENT** 
---

This repo is prepared for basic and intermediate understanding of Go Programming Language. The workspace of the codes written here is Visual Studio Code.
You can find the download link in the KEY POINTS section. You also need to download to install Go in Visual Studio Code. Its link is also in the KEY POINTS section.



**REPO CONTENTS**
---

- [variables](https://github.com/a1zuws/StudysGolang/tree/main/variables)
- [conditionals](https://github.com/a1zuws/StudysGolang/tree/main/conditionals)
- [loops](https://github.com/a1zuws/StudysGolang/tree/main/loops)
- [arrays](https://github.com/a1zuws/StudysGolang/tree/main/arrays)
- [slices](https://github.com/a1zuws/StudysGolang/tree/main/slices)
- [functions](https://github.com/a1zuws/StudysGolang/tree/main/functions)
- [maps](https://github.com/a1zuws/StudysGolang/tree/main/maps)
- [for_range](https://github.com/a1zuws/StudysGolang/tree/main/for_range)
- [pointers](https://github.com/a1zuws/StudysGolang/tree/main/pointers)
- [structs](https://github.com/a1zuws/StudysGolang/tree/main/structs)
- [goroutines](https://github.com/a1zuws/StudysGolang/tree/main/goroutines)
- [channels](https://github.com/a1zuws/StudysGolang/tree/main/channels)
- [interfaces](https://github.com/a1zuws/StudysGolang/tree/main/interfaces)
- [defer_statement](https://github.com/a1zuws/StudysGolang/tree/main/defer_statement)
- [error_handling](https://github.com/a1zuws/StudysGolang/tree/main/error_handling)
- [string_functions](https://github.com/a1zuws/StudysGolang/tree/main/string_functions)
- [restful](https://github.com/a1zuws/StudysGolang/tree/main/restful)
- [project](https://github.com/a1zuws/StudysGolang/tree/main/project)



DOWNLOAD LINKS
---

- Go Download Link: [Golang](https://go.dev/dl/)
  - Click on the one appropriate for your computer and it will download.
  
- Visual Studio Code Download Link: [Visual Studio Code](https://code.visualstudio.com/download)
  - Click on the one appropriate for your computer and it will download.
  


**KEY POINTS**
---

Module Creation :

``` 
E:\StudysGolang> go mod init <module name> 
```
- *NOTE =>* You can replace <module name> with any module name you want.
  
Create A Folder : 
``` 
C:\> mkdir goWorks --> E:\> cd goWorks 
  ```

Terminal Output : 
``` 
C:\StudysGolang> go run main.go 
  ```

Install JSON Server :
``` 
C:\StudysGolang> npm install -g json-server
  ```
  - *ATTENTION! =>* If you get an error in JSON Server installation, install it from [Node.js](https://nodejs.org) , close and reopen your workspace.

Running The JSON Server :
```
C:\StudysGolang> json-server --watch db.json
  ```
  - *ATTENTION! =>* If it fails after running it, type `npx` followed by a space and then `json-server --watch db.json` and run it again.
  - So like this : 
```
C:\StudysGolang> npx json-server --watch db.json
  ```
