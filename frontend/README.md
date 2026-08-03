# What is and how to run?
## What is?
This code is implementation of our frontend website that features sighting map and multiple other pages.
* `Dockerfile` is used to run our application in container, it includes installation of all requirements and building of our application.
* `.json`, `.config.js` and `.config.ts` files are requirements that are used during build process.
* `src/` folder contains our application source code.
* * `src/App.vue` is an entrypoint for our application, allows us to have multiple pages inside our application.
* * `src/pages` folder, obviously, is a folder for our application pages.
* * `src/pages/Home.vue` is our main page that features our map.
* * `src/pages/SightingDetail.vue` is a page that features detailed post information, such as time of discovery and description.
* * `src/pages/GoldMasonPanel.vue` is a control panel page used by golden mason to perform certain actions.
* * `src/pages/AddSighting.vue` is a page that allows users to add sightings they spotted.
* * `src/router/index.ts` is a code that allows us to register a page and link it to a certain URL in browser.
* * `src/styles is a folder` for CSS styles, names of CSS files corespond to names of files in src/pages/.
* * `src/api/client.js` is a JS code used to let GoldenMasonPanel speak to Django backend services.
## How to run?
* First download the repository.
* `cd` into repository folder
* Run `docker compose up --build` inside main repository folder (not frontend)
* Go to browser, type `http://localhost:3000` and enjoy the website
## Prerequisites
* Have Git installed
* Have `Docker` and `Docker Compose` installed