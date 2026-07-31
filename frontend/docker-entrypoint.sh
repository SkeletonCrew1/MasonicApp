#!/bin/sh
{
    echo "window.__APP_CONFIG__ = {"
    echo "  API_BASE_URL: \"${API_BASE_URL}\","
    echo "  BIRDWATCHING_URL: \"${BIRDWATCHING_URL}\""
    echo "};"
} > /app/dist/config.js

exec "$@"