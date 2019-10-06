# Scout Regiment

Pupuru Go microservices boilerplate.

# Development

- Clone this repository: `git clone https://github.com/ianadiwibowo/gopupuru`
- Copy `.env` from `.env.template`: `cp .env.template .env`
- Setup database (only for first time): `rake db:create`
- Migrate database: `rake db:migrate`
- Install dependencies: `make dep`
- Build the binary: `make build`
- Run the service: `make run`

# Create New DB Migration

- Create the migration file: `rake db:new_migration name=create_cats`
- Edit the file: `vim db/migrate/20190922192846_create_cats.rb`
- Follow the guideline: https://guides.rubyonrails.org/active_record_migrations.html
