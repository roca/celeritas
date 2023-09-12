# celeritas

## Available commands:
```sh
./celeritas help # show the help commands
./celeritas version # print application version
./celeritas migrate # run all up migrations that have not been run previously
./celeritas migrate down # reverses the most recent migration
./celeritas migrate reset # runs all down migrations in reverse order, and then runs all up migrations
./celeritas make migration <name> # create two new up and down migrations in the migrations folder
./celeritas make auth # creates and runs migrations for the authentication tables, and creates models and middleware
./celeritas make handler <name> # creates a stub handler in the handlers directory
./celeritas make model <name> # creates a new model in the data directory
./celeritas make session # creates a table in the database as a session store
./celeritas make mail <name> # creates two starter mail templates in the mail directory
```


## ToDo:

- [Add file storage feature using viant/afs](https://github.com/viant/afs)
- Laravel Broadcasting
- Lavarel Dusk
- Laravel Queues


[Laravel Docs](https://laravel.com/docs/10.x)