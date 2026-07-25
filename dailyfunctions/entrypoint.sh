#!/bin/sh

# Start the password service in the background
./password-server &

# Start the inquisitor service in the background
./inquisitor-server &

# Wait for any background process to exit (keeps container alive)
wait