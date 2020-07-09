package analyzer

import (
	"fmt"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("base_analyzer")

type BaseAnalyzer struct {
	config
}

func main() {
	fmt.Println("vim-go")
}
