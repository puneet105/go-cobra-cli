##Create your own command command using cobra-cli

###Create workspace
```
mkdir go-cobra-cli && cd go-cobra-cli
go mod init github.com/<USER_NAME>/go-cobra-cli
```
###Install cobra-cli  on your machine
####cobra-cli init command will generate your initial application code for you.
####You can use few flags while running below command
```cobra-cli init --author "Puneet Sharma abc@gmail.com" ```

```cobra-cli init --license apache ```
####Viper is a companion to Cobra intended to provide easy handling of environment variables and config files and seamlessly connecting them to the application flags.
``` cobra-cli init --viper ```

```
go install github.com/spf13/cobra-cli@latest
go get -u github.com/spf13/cobra
cobra-cli init 
```

####You can verify whether everything is working fine or not by running below command.
```go run main.go```

Output will look like this.
```
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
```

####Now open **root.go** file and uncomment the Run: field and add below print statement inside function.

Also, you can see the name of your command from your root.go file. Below is the reference.
**Use:   "go-cobra-cli",**

```
Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello From First Cobra Example")
	}
```

####Again run your main.go file by giving your command name as an argument.
```
go run main.go go-cobra-cli

Output : Hello From First Cobra Example
```

####Once a cobra application is initialized you can continue to use the Cobra generator to add additional commands to your application. 
####The command to do this is ```cobra-cli add <COMMAND_NAME>```

####Let's create another command to print fibonacci series based on the number we provide as an argument. Run the below command to add another cobra-cli command.

```
cobra-cli add FibonacciSeries

Output : FibonacciSeries created at C:\go\src\github.com\puneet105\go-cobra-cli
```

####It will create **FibonacciSeries.go** file under cmd/ . You can make below changes in the file.

```
/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// FibonacciSeriesCmd represents the FibonacciSeries command
var FibonacciSeriesCmd = &cobra.Command{
	Use:   "FibonacciSeries",
	Short: "Command to print a fibonacci series based on the number we provide as an argument",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("FibonacciSeries called....!!")
		printFibonacci(args[0])
	},
}

func printFibonacci(num string){
	x := 0
	y := 1
	numInt,_ := strconv.Atoi(num)
	fmt.Printf("Fibonacci Series is : %d %d", x,y)
	for i:=2;i<numInt;i++{
		z := x + y
		fmt.Printf(" %d", z)
		x = y
		y = z
	}
}

func init() {
	rootCmd.AddCommand(FibonacciSeriesCmd)

}

```

#### Now run main.go file.

```
go run main.go --help

Output : A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  go-cobra-cli [flags]
  go-cobra-cli [command]

Available Commands:
  FibonacciSeries Command to print a fibonacci series based on the number we provide as an argument
  completion      Generate the autocompletion script for the specified shell
  help            Help about any command

Flags:
  -h, --help     help for go-cobra-cli
  -t, --toggle   Help message for toggle

Use "go-cobra-cli [command] --help" for more information about a command.

```

```
go run main.go FibonacciSeries 5

Output : 
FibonacciSeries called....!!
Fibonacci Series is : 0 1 1 2 3

```