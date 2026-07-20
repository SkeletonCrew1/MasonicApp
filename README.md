# SkeletonCrew2 — Cult of the Tree

Minimal implementation. One database only: **PostgreSQL** (Redis removed from
docker-compose, since broadcast messages and ban state are just stored as
plain Postgres tables and the frontend polls/re-fetches — no pub/sub needed
for this scope).

## Run it

```bash
docker compose up --build
```

- Frontend (map + UI): http://localhost:5173
- Django API: http://localhost:8000/api
- ban-service: http://localhost:8081
- broadcast-service: http://localhost:8082

Django auto-runs `makemigrations` + `migrate` on container start, so no
manual migration files are checked in.

To use the Gold Mason panel, create a superuser-like account and flip
`is_gold_mason` in the Django admin (http://localhost:8000/admin, after
`docker compose exec backend-django python manage.py createsuperuser`).

## Tickets → what was built

- **SKEL2-9 — Implement Django backend**
  `backend-django/api`: `User` (with `status`: bronze/silver/golden and
  `is_gold_mason`), `Post`, `Invite` models + REST endpoints for auth,
  posts, user search/promotion, invite, and "delete all data".

- **SKEL2-8 — Create window with post info**
  `frontend/src/components/CreatePostWindow.vue` — modal for posting a new
  sighting (title, image URL, description) at the clicked map coordinates.
  Existing posts show their info (title, image, description, "Open Page")
  in a Leaflet popup on the map, styled after the reference site.

- **SKEL2-14 — Create ban window**
  `frontend/src/components/BanWindow.vue` — the "Ban agent" box: search
  field ("Find somebody by IP") + BAN button, matching the wireframe.

- **SKEL2-19 — Implement ban feature**
  `services-go/ban-service` — Go HTTP service, own `banned_ips` table in the
  shared Postgres DB. `POST /ban`, `GET /check?ip=`, `GET /bans`.

- **SKEL2-20 — Implement broadcast feature**
  `services-go/broadcast-service` — Go HTTP service, `broadcasts` table in
  the shared Postgres DB. `POST /broadcast` (message + target statuses:
  bronze/silver/golden checkboxes), `GET /broadcasts`.

The full **Gold mason control panel** wireframe (Broadcast / BAN / Promotion
/ invite / Delete all data) is wired up in
`frontend/src/components/GoldMasonPanel.vue`.
