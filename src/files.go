package main

import(
    "path/filepath"
    "os"
)

func getData(searchDir string) []directory {
   
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

    return fileList
}
