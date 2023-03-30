package app

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

//create array of commands
var Commands = []cli.Command{

	{
		Name:  "fish",
		Usage: "Draw a fish",
		Action: func(c *cli.Context) error {
			Fish()
			return nil
		},
	},
	{
		Name:  "cat",
		Usage: "Draw a cat",
		Action: func(c *cli.Context) error {
			Cat()
			return nil
		},
	},
	{
		Name:  "dog",
		Usage: "Draw a dog",
		Action: func(c *cli.Context) error {
			Dog()
			return nil
		},
	},
	{
		Name:  "frog",
		Usage: "Draw a frog",
		Action: func(c *cli.Context) error {
			Frog()
			return nil
		},
	},
	{
		Name:  "lobster",
		Usage: "Draw a lobster",
		Action: func(c *cli.Context) error {
			Lobster()
			return nil
		},
	},
}


//create function to draw in the terminal the app
func Start() {
	app := cli.NewApp()
	app.Name = "CLI ZOO"
	app.Usage = "App to draw animals in the terminal"
	app.Version = "1.0.0"
	app.Commands = Commands
	app.Run(os.Args)
}

func handleError(err error) error {
	if err != nil {
		return err
	}
	return nil
}


func Fish() {

	fmt.Println("       /`-._")
	fmt.Println("     _/,.._/")
	fmt.Println("  ,-'   ,  `-:,.-')")
	fmt.Println(" : o ):';     _  {")
	fmt.Println("  `-.  `' _,.-\\`-.)")
	fmt.Println("     `\\\\``\\,.-'")

}

func Cat() {

	fmt.Println("  .       .")
	fmt.Println("  \\`-\"'\"-'/")
	fmt.Println("   } 6 6 {")
	fmt.Println("  =.  Y  ,=")
	fmt.Println("    /^^^\\  .")
	fmt.Println("   /     \\  )")
	fmt.Println("  (  )-(  )/")
	fmt.Println("   \"\"   \"\"")

	
}

func Dog() {
	fmt.Println("   / \\__")
	fmt.Println("  (    @\\___")
	fmt.Println("   /         O")
	fmt.Println("  /   (_____/")
	fmt.Println(" /_____/   U")
}

func Frog() {


	fmt.Println("           .--._.--.")
	fmt.Println("          ( O     O )")
	fmt.Println("          /   . .   \\")
	fmt.Println("         .`._______.'.")
	fmt.Println("        /(           )\\")
	fmt.Println("      _/  \\  \\   /  /  \\_")
	fmt.Println("   .~   `  \\  \\ /  /  '   ~.")
	fmt.Println("  {    -.   \\  V  /   .-    }")
	fmt.Println("_ _`.    \\  |  |  |  /    .'_ _")
	fmt.Println(" >_       _} |  |  | {_       _<")
	fmt.Println(" /. - ~ ,_-'  .^.  `-_, ~ - .\\")
	fmt.Println("         '-'|/   \\|`-`")


}

func Lobster(){

	fmt.Println("        .\"\".-._.-.\"\".")
	fmt.Println("        |   \\  |  /   |")
	fmt.Println("         \\   \\.T./   /")
	fmt.Println("          '-./   \\.-'")
	fmt.Println("            /     \\")
	fmt.Println("           ;       ;")
	fmt.Println("           |       |")
	fmt.Println("           |       |")
	fmt.Println("          /         \\")
	fmt.Println("          |    .    |")
	fmt.Println("       __.|    :    |.__")
	fmt.Println("   .-'`   |    :    |   `'-.")
	fmt.Println(" /`     .\"\\  0 : 0  /\".     `\\")
	fmt.Println("|     _/   './ : \\.'   \\_     |")
	fmt.Println("|    /      /`\"\"\"\"`\\      \\    |")
	fmt.Println(" \\   \\   .-'       '-._   /   /")
	fmt.Println("  '-._\\                 /_.-'")

}