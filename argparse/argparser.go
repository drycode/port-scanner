package argparse

import "fmt"

// ArgParser ...
type ArgParser struct {

}

func (ap *ArgParser) parseArgs(){
	fmt.Print(ap)
	fmt.Println("Argparser, standing by")
}
