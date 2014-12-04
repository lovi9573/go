package main

import(
    "fmt"
    "os"
    "strconv"
)

var size int64 = 0
var start string = ""

func main(){

    //there is either no arguments or 2.
    //starting location first followed by size
    switch len(os.Args){

    case 1:
	start = "."
	size = 0
    case 3:
	start = os.Args[1]
	j, err := strconv.Atoi(os.Args[2])

	if err != nil {
	    fmt.Println("Incorrect size")
	    os.Exit(1)
	}

	size = int64(j)

    default:
	fmt.Println("Incorrect arguments")
	fmt.Println("driver or driver <start location> <files sizes greater then>")
	os.Exit(1)
    }

    directory := getData(start)

	//building a webpage using string concatenation...
	html := "<html><head>"
	html += "<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css\">"
	html += "<script src=\"//ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js\"></script>"
	html += "<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js\">"
	html += "</script><link rel=\"stylesheet\" href=\"main.css\"></head><body>"

	html += "<div class=\"container\"><div class=\"top-bar\"><div data-spy=\"affix\" class=\"white-back\"><div class=\"page-header\">"
	html += "<h1>Quota<span class=\"lessbold\">nizer</span><small> A friendly disk usage utility</small></h1></div></div></div>"
	html += "<form action='/submit' method='post'>"
	html += "<div class='form-group'>"
	html += "<label for='dir'>Directory to analyze </label>"
	html += "<input type='text' name='dir' placeholder='Directory'>"
	html += "</div>"
	html += "<div class='form-group'>"
	html += "<label for='minSize'>Minimum file size (bytes) </label>"
	html += "<input type='text' name='minSize' value='0'>"
	html += "</div>"
	html += "<div class='form-group'>"
	html += "<input class='btn btn-info' type='submit' value='Submit'>"
	html += "</div>"	
	html += "</form>"
	html += "<table class=\"table\"><tr><th>Path</th><th>Size</th><th>Options</th></tr>"


	flag := false
	for i := len(directory)-1; i >= 0; i-- {
	dir := directory[i]

		//only adds items to the list that are larger than the size passed as command line arg...
		if dir.size >= size {
		flag = true;
		count :=0
		for ; dir.size >=1000; count++ {
			dir.size /= 1000
		}

		output := strconv.FormatInt(dir.size, 10)
		switch count {
		case 0:
			output += " bytes"
		case 1:
			output += " KB"
		case 2:
			output += " MB"
		case 3:
			output += " GB"
		}

		html += "<tr><td>" + dir.path  + "</td><td>" + output
		html += "</td><td><a href=\"javascript:void(0)\">DELETE</a></td></tr>"
		}
    }

	if !flag {
	html += "<tr><td>No items fit the specified search critera. Please try again.</tr></td>"
	}

	html += "</table></div></body></html>"
	fmt.Println(html);

}
