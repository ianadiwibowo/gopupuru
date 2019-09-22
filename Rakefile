#!/usr/bin/env rake

ENV['APP_ENV'] ||= 'development'

require 'dotenv'
require 'standalone_migrations'

if ENV['APP_ENV'] == 'development'
  Dotenv.load
end

StandaloneMigrations::Tasks.load_tasks
