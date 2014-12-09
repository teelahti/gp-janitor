Group policy janitor
====================

This is a simple program to clean Windows domain group policy cruft from the machine. What should be cleaned is highly subjective based on user, and the applied group policy, but this project provides a simple framework to make any kind of cleanups. Assumes that executing user has administrative privileges to the computer. 

There are three kind of triggers for cleaning up computer:  

1. At program (service launch)
2. At timer interval
3. With file system watcher

Currently only #2 is implemented, as that was enough to fix the most annoying policy of forced IE home page.

Installing
----------

Install [Go](https://golang.org/), if not yet installed. Then run

    go get github.com/teelahti/gp-janitor

This will download sources for this project, compile them, and "install" this package - in other words place the sources into your $GOPATH$/Src folder, and the resulting binary into your $GOPATH$/bin folder.

*For more information about the Go command line tool check [official documentation](https://golang.org/cmd/go/).*

Now open an administrative command prompt, and 

    cd $GOPATH$/bin
    
    # Run from within command prompt, do not attempt
	# to run as Windows service. Use this to test 
    # operation before installing for good
    gp-janitor -interactive

    # Install as Windows service
    gp-janitor -register

    # Remove windows service registration
    gp-janitor -unregister

**Note!**

As this service changes Internet Explorer registry values under the current user registry node, it cannot run as LocalSystem (Windows service default). You need to use sc.exe or open local services manager, and change the service user to your personal user account you want to keep clean. This is annoying, but there is no easy other way. Another option is to use this tool as startup command instead of a Windows service.