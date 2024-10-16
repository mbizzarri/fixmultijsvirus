//
// fixonefile - fix one file - remove virus
// related to fixmultifiles - this is the single file version
// This is used to remove virus from .js files under WordPress
// note the virus signatures are hard coded into 4 variables
//
//	substr := ";if(typeof zqxq==="
//	midstr := "M:0xe1},H=Q,V=d,N=navigator,M=document,P=screen,v=window,k=M[V(X.N)"
//      midstr2 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/="
//	endstr := "(C.N)"
// default starts with current directory
// if a directory is used as a parameter, the search and destroy starts there
//
//
// written by Maurice Bizzarri, October 2024
// First release: October, 2024
//
package main

import (
        "bufio"
	"strings"
	"io"
	"fmt"
	"os"

)


//
// main function - default to . unless parameter given
// this routine walks down the directories
// all work done by fixfiles
//
func main() {
        fmt.Println("fixonefile - fix virus in one file\nVersion 0.2\n")

        if (len(os.Args) <  2) {
		fmt.Printf("Command line parameter must be file name!\n")
		os.Exit(-1)
	}

	path := os.Args[1]    // first command line parameter	
        fixfiles(path)






}
//
// fixfiles - get rid of virus at end of javascript file
//

func fixfiles(files string) {
	fmt.Println("filename: "+files)
	//
	//open up file
	//

	// first we need to read in one file
	// then we need to look for start of virus string inside of file
	// then we either hit end of file, or the start of the virus string
	//
	// if both strings are preset, the code in between is removed from the f
	//

	// make a buffer to keep chunks that are read
	//
	buf := make([]byte, 1600000)

	//* open up file
	fi, err := os.Open(files)
	if err != nil {
		panic(err)
    }
    // make a read buffer
    r := bufio.NewReader(fi)
    
    // read a chunk
      n, err := r.Read(buf)
      if err != nil && err != io.EOF {
            panic(err)
       }
	// nothing read?
       if n == 0 {
           return
       }
     // close fi and check for its returned error
    
        if err := fi.Close(); err != nil {
            panic(err)
        
	}


	    //
	    // check buf for string
	    //

         //
         // substrings to look for
         // these are all virus signatures

	substr := ";if(typeof zqxq==="
	midstr := "M:0xe1},H=Q,V=d,N=navigator,M=document,P=screen,v=window,k=M[V(X.N)"
        midstr2 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/="
	endstr := "(C.N)"

  	strbuf := string(buf)
        index := strings.Index(strbuf, substr)
	index2 :=strings.Index(strbuf, endstr)
	index3 := strings.Index(strbuf,midstr)
	index4 := strings.Index(strbuf,midstr2)


	if index4 == -1 {
		fmt.Println("Virus signature not found (middle 1)\n")
		return
	}

	if index3 == -1 {
		fmt.Println("Virus signature not found (middle 2)\n")
		return
	}
	if index2 == -1 {
		fmt.Println("Virus Signature (end) not found\n")
		return
	}
        if index == -1 {
		fmt.Println("Virus Signature (beginning) not found \n")
	    return
        } else
	{
		//clean up the file

		    fmt.Println(index)
	    }
        // write a chunk
	// open output file
	fo, err := os.Create(files)
		if err != nil {
			panic(err)
		}
	// make a write buffer
	w := bufio.NewWriter(fo)
	//
	// write the buffer out
	//
       if _, err := w.Write(buf[0:index]); err != nil {
            panic(err)
        }

    if err = w.Flush(); err != nil {
        panic(err)
    }
// close output file and check for its returned error

        if err := fo.Close(); err != nil {
            panic(err)
        }


}
