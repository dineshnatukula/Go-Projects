Definiing the libraries outside of this project and accessing them in the project.
Observe the main.go file and go.mod files for better clarity.

Step 1:
Outside of this project, define the go.mod with the following line.
    github.com/dinesh.natukula/mylib
The structure of the mylib is as follows:
    go-libs
    |__mylib
        |__dblibs
            |__dblib.go
        |__mylib.go
        |__go.mod
From the above structure, your go.mod file should look like as follows:
<code>
    module github.com/dinesh.natukula/mylib
    go 1.21
</code>

