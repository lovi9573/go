package main

import(
    "path/filepath"
    "os"
    "fmt"
    "strconv"
)

//var displaySize int64

func getData(searchDir string) []directory {
   
//user can pass a directory in the current 
//directory or pass nothing in and the 
//program will use the current directory
//    switch len(os.Args) {
    
//    case 1:
//	displaySize = 0
//    case 2:
//	other, err := strconv.Atoi(os.Args[1])
//	displaySize = int64(other)
//
//	if err != nil {
//	    fmt.Println("TODO")
//	}
//    default:
//	fmt.Println("Incorrect Number of Arguments")
//	os.Exit(1)
//    }

    fileList := []directory{}

    //walks through the directories starting at the searchDir var
    filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error{
	    if f.IsDir(){					    //if a directory is found
		fileList = append(fileList, directory{0, path})    //add it to the list
	    }

	return nil
    })

    //goes through the found directories to calculate the total size 
    for i := 0; i<len(fileList); i++ {
        filepath.Walk(fileList[i].path, func(path string, f os.FileInfo, err error) error{
	    fileList[i].size += f.Size()
	    return nil
	})
    }

    pmergesort(fileList)

//    for i := len(fileList)-1; i >= 0; i-- { 
	
//	f := fileList[i]
	
//	if f.size >= displaySize {
//	    fmt.Println(f.path, f.size)
//	}
//    }

    return fileList
}
