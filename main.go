package ultimateAmiDestroyer

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Byron's famous AMI Crusher package. No more annoying AMI's. Bye bye!")
}

// CrushAMIs removes all of your accounts AMIs
// Account ID is passed as a string parameter
// Nothing is returned because it's guaranteed to remove all of your AMI's
func CrushAMIs(acctID string) {
	fmt.Printf("Thanks for flying Air AMI Crusher. Total AMIs for account number %s: 0, duh! :P Have a nice AMI free day\n", acctID)
}
