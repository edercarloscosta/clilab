### Clilab

This is a laboratory that implements a quick note by CLI Cobra. 

I just want to create a simple note tagging it by category.

### Execution Steps

I'm using SQLite3

As you can see on the code you need to start by creating the database first and executing the first two commands below

Inside the project root type on terminal

> go build .

After the build you will see a new file on the root project named `clilab`

Now just need to execute
> ./clilab init

You will see a new file named `sqlite-database.db` 

### Cli Commands

For general infos or useful guide

> ./clilab note

For create a new note on database

> ./clilab note new

For list all notes from database

> ./clilab note list

For delete a specific note 

> ./clilab note del
