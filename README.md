# yagg

**Y**et **A**nother **G**it **G**ist is a light-weight no-database public git-gist. It allows operations like creating, editing, viewing, downloading and showing history of the gists created.

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

## Server Side Commands:

```
yagg  [-b] [-h] [-p] [-r]
```

## Options:


``` -b host-name```
    Used to set the hostname for the server. If nothing is specified, 0.0.0.0 will be used as the host.

``` -h ```
    Used for getting the usage of the flags.

``` -r path/to/save/repo ```
     Used to set the repository path where the git-gist repository will be saved. If nothing is specified, data folder in the present working directory will be set as the default path.

``` -p port-number ```
     Used to specify the port used by YAGG. Port 8080 will be used by default.

## How to use:

- Run the server at 0.0.0.0:8080 using command :

  ```
  $ yagg -b 0.0.0.0 -p 8080
  ```

- Let's **create** our first gist ``` main.go ```

    - ```http://0.0.0.0:8080/```

    ![Create-gist-image](https://github.com/gophergala2016/yagg/blob/master/images/create_gist.png "create")

- **View** 
    
    - ```http://0.0.0.0:8080/show/gist-name```
    - ```http://0.0.0.0:8080/show-by-id/gist-name/commit-id```

    ![View-gist-image](https://github.com/gophergala2016/yagg/blob/master/images/view_gist.png "view")

- **Edit**
    
    - ```http://0.0.0.0:8080/edit/gist-name```

    ![Edit-gist-image](https://github.com/gophergala2016/yagg/blob/master/images/edit_gist.png "edit")

- Track the **changes** 
    
    - ```http://0.0.0.0:8080/history/gist-name```

    ![History-gist-image](https://github.com/gophergala2016/yagg/blob/master/images/history_gist.png "history")

- Generate **raw** file
    
    - ```http://0.0.0.0:8080/raw/gist-name```
    
    - ```http://0.0.0.0:8080/raw-by-id/gist-name/commit-id```

    ![Raw-gist-image](https://github.com/gophergala2016/yagg/blob/master/images/raw_gist.png "raw")

- **Download** 

    -```http://0.0.0.0:8080/download/gist-name```
    
    -```http://0.0.0.0:8080/download-by-id/gist-name/commit-id```

- **Delete** 
    
    -```http://0.0.0.0:8080/download/gist-name```

- **View all**:

    -```http://0.0.0.0:8080/list```
    
    ![List-gist-image](https://github.com/gophergala2016/yagg/blob/master/images/list_gist.png "list")
    
## To-do:
- Sorting gists in the list page based on options like last updated, etc. Currently gists are sorted in the order of gist-name
- Search field to be included
- Option of uploading a file as a gist
- User specific
- User Interface to be enhanced

 ## References
- [Diff2Html](https://github.com/rtfpessoa/diff2html)
- [git2go](https://godoc.org/github.com/libgit2/git2go)
- [blog](https://blog.gopheracademy.com/advent-2014/git2go-tutorial/)
