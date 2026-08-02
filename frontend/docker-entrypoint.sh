#!/bin/sh
{ # Group output streams
    echo "window.__APP_CONFIG__ = {" # Create config object
    echo "  API_BASE_URL: \"${API_BASE_URL}\"," # Map main API URL
    echo "  BIRDWATCHING_URL: \"${BIRDWATCHING_URL}\"" # Map Birdwatching URL
    echo "};" # End config obj
} > /app/dist/config.js # Save to the JS file

exec "$@" # Execute Docker args
