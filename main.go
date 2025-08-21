package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ahsanwtc/gator/internal/config"
	"github.com/ahsanwtc/gator/internal/database"
	"github.com/ahsanwtc/gator/internal/services"
	_ "github.com/lib/pq"
)

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		fmt.Println("could not connect to the database")
		os.Exit(1)
	}

	dbQueries := database.New(db)
	state := State{
		config: &cfg,
		db: dbQueries,
		userService: services.NewUserService(dbQueries),
	}

	commands := Commands{
		commands: make(map[string]CommandHandler),
	}

	// commands
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAggregate)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	
	args := os.Args

	if len(args) < 2 {
		fmt.Println("too few arguments, user `help`")
		os.Exit(1)
	}

	command := args[1]
	parameters := args[2:]
	
	var commandError error
	switch command {
		case "login":
			commandError = commands.run(&state, Command{
				name: "login",
				parameters: parameters,
			})
		case "register":
			commandError = commands.run(&state, Command{
				name: "register",
				parameters: parameters,
			})
		case "reset":
			commandError = commands.run(&state, Command{
				name: "reset",
				parameters: parameters,
			})
		case "users":
			commandError = commands.run(&state, Command{
				name: "users",
				parameters: parameters,
			})
		case "agg":
			commandError = commands.run(&state, Command{
				name: "agg",
				parameters: parameters,
			})
		case "addfeed":
			commandError = commands.run(&state, Command{
				name: "addfeed",
				parameters: parameters,
			})
		case "feeds":
			commandError = commands.run(&state, Command{
				name: "feeds",
				parameters: parameters,
			})
		case "follow":
			commandError = commands.run(&state, Command{
				name: "follow",
				parameters: parameters,
			})
		case "following":
			commandError = commands.run(&state, Command{
				name: "following",
				parameters: parameters,
			})
		default:
			fmt.Printf("`%s`: unknown command\n", command)
			os.Exit(1)
	}

	if commandError != nil {
		fmt.Println("Error: ", commandError)
		os.Exit(1)
	}
}