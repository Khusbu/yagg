# yagg

Yet-Another-Git-Gist(YAGG) is a light-weight no-database git gist. It allows operations like creating, editing, viewing, downloading and showing history of the gists created.

## Installation


- Install cmake from brew or apt-get package manager. To build from source, follow [this](https://cmake.org/install/) link

- Install libgit2 as follows :
    ```
      $ wget https://github.com/libgit2/libgit2/archive/v0.23.4.tar.gz
      $ tar xzf v0.23.4.tar.gz
      $ cd libgit2-0.23.4/
      $ cmake .
      $ make
      $ sudo make install
    ```

- Install go packages as folloes :
    ```
      $ go get github.com/libgit2/git2go
    ```

- Build the project using :
    ```
      $ go build
    ```


- Troubleshooting:-  
    ```ImportError: libgit2.so.0: cannot open shared object file: No such file or directory```  
         This happens for instance in Ubuntu, the libgit2 library is installed within the /usr/local/lib directory, but the linker does not look for it there.
         To fix this call
    ```
      $ sudo ldconfig
    ```

## Server Side Usages:

```
servidor  [-b] [-h] [-p] [-r]
```

## Options:


``` -b host-name```
    Used to set the hostname for the server. If nothing is specified, 0.0.0.0 will be used as the host.

``` -h ```
    Used for getting the usage of the flags.

``` -r path/to/save/repo ```
     Used to set the repository path where the git-gist repository will be saved. If nothing is specified, present working directory will be set as the default path.

``` -p port-number ```
     Used to specify the port used by YAGG. Port 8080 will be used by default.

## Examples:

**$ YAGG**

Runs the server at 0.0.0.0:8080

**$ yagg -b 0.0.0.0 -p 8080**


## Acknowledgement
- [Diff2Html](https://github.com/rtfpessoa/diff2html)
- [git2go](https://godoc.org/github.com/libgit2/git2go)
- [blog]https://blog.gopheracademy.com/advent-2014/git2go-tutorial/
