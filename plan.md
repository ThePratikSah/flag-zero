# 🚀 Minimal Feature-Flag MVP — Checklist (Using Full Clean Architecture Folder Structure)

# Folder Structure Used
cmd/
  server/
  worker/         (empty for MVP)
internal/
  app/            (usecases)
  domain/         (entities/models)
  ports/          (interfaces)
  adapters/       (mysql, redis implementations)
  transport/      (http handlers)
  auth/           (optional for future)
pkg/
  sdk/            (empty for MVP)
migrations/
deploy/
tests/
scripts/
config/

---

# Milestone 0 — Basic Project Setup

## Tasks
- [x] Initialize Go module (`go mod init`)
- [x] Create full folder structure (clean architecture)
- [x] Add `.env.example` with:
  - [x] MYSQL_DSN
  - [x] REDIS_ADDR
  - [x] SERVER_PORT
- [x] Add config loader in `/config`
- [x] Implement minimal HTTP server in `/cmd/server`:
  - [x] `/health` route
  - [x] `/ready` route
- [ ] Add MySQL connection in `/internal/adapters/mysql`
- [ ] Add Redis client in `/internal/adapters/redis`
- [ ] Add `docker-compose.yml` with:
  - [ ] MySQL
  - [ ] Redis
- [ ] Add `Makefile` with:
  - [ ] `make run`
  - [ ] `make test`
  - [ ] `make compose-up`
- [ ] Add root-level README:
  - [ ] Overview
  - [ ] Running locally
  - [ ] Folder structure

## Acceptance Criteria
- [ ] Server runs locally
- [ ] MySQL & Redis boot via docker-compose
- [ ] `/health` returns 200

---

# Milestone 1 — Domain & Repository Layer

## Tasks
- [ ] Create `Flag` domain model in `/internal/domain`
  - id
  - key
  - value (json/string)
  - enabled (bool)
  - updated_at (timestamp)
- [ ] Create MySQL schema migration:
  - [ ] `migrations/001_create_flags_table.sql`
- [ ] Define repository interface in `/internal/ports`:
  - [ ] `CreateFlag`
  - [ ] `UpdateFlag`
  - [ ] `GetFlagByKey`
  - [ ] `GetAllFlags`
  - [ ] `DeleteFlag`
- [ ] Implement repository in `/internal/adapters/mysql`
- [ ] Write basic tests under `/tests` (optional DB or mock)

## Acceptance Criteria
- [ ] MySQL migrations run cleanly
- [ ] Repo functions work (manual testing or test DB)

---

# Milestone 2 — Service Layer (Usecases)

## Tasks
- [ ] Create usecase/service layer in `/internal/app/services`
  - [ ] `CreateFlagService`
  - [ ] `UpdateFlagService`
  - [ ] `GetFlagService`
  - [ ] `ListFlagsService`
  - [ ] `DeleteFlagService`
- [ ] Validation logic:
  - [ ] key must not be empty
  - [ ] value must be valid JSON or string
- [ ] Integrate Redis cache in `/internal/adapters/redis`
  - [ ] Cache key: `flags:all`
  - [ ] Write-through on create/update/delete:
    - [ ] Update DB
    - [ ] Refresh Redis cache
  - [ ] Read-through for list:
    - [ ] Try Redis → fallback to DB → update Redis
- [ ] Add optional counters or logs

## Acceptance Criteria
- [ ] Caching works (verified by logs/prints)
- [ ] Cache miss hits DB only once
- [ ] Cache updates after flag write operations

---

# Milestone 3 — HTTP API (Admin CRUD)

## Tasks
- [ ] Add Gin handlers under `/internal/transport/http`
- [ ] Route definitions:
  - [ ] `POST /api/flags`
  - [ ] `PUT /api/flags/:key`
  - [ ] `GET /api/flags`
  - [ ] `GET /api/flags/:key`
  - [ ] `DELETE /api/flags/:key`
- [ ] Bind request JSON
- [ ] Call service layer methods
- [ ] Map errors → HTTP responses
- [ ] Add minimal validation middleware (optional)

## Acceptance Criteria
- [ ] All CRUD operations work via Postman/curl
- [ ] Redis shows updated results after CRUD
- [ ] DB and cache remain in sync

---

# Milestone 4 — Client-Facing Endpoint (Read-Only)

## Tasks
- [ ] Add read-only endpoint for clients:
  - [ ] `GET /client/flags`
- [ ] Behavior:
  - [ ] Fetch from Redis → fallback DB → write to Redis
- [ ] Response:
  - [ ] `{ "flags": [ ... ] }`
- [ ] No auth, no keys (MVP only)

## Acceptance Criteria
- [ ] Read latency is low (<5ms on cache hit)
- [ ] Only the first request populates cache if empty
- [ ] Clients get updated flags immediately after admin CRUD

---

# Milestone 5 — Minimal Deployment Flow

## Tasks
- [ ] Create `Dockerfile` for server
- [ ] Add `docker-compose.prod.yml`
- [ ] Add `.env.prod.example`
- [ ] Add documentation:
  - [ ] “How to run migrations”
  - [ ] “How to start server”
  - [ ] “How to configure MySQL/Redis”
- [ ] Add basic logger config in `/internal/adapters/logging`

## Acceptance Criteria
- [ ] App can be deployed via docker-compose
- [ ] Flags persist across restarts (MySQL)
- [ ] Cache warms cleanly on first request

---

# Milestone 6 — Polishing the MVP

## Tasks
- [ ] Add simple basic-auth (optional)
- [ ] Add rate limiting (optional)
- [ ] Add error-wrapping helper utilities
- [ ] Improve README with:
  - [ ] API examples
  - [ ] Curl samples
  - [ ] Architecture summary
- [ ] Add seed script in `/scripts/seed.go`

## Acceptance Criteria
- [ ] MVP stable & simple to self-host
- [ ] Minimal docs ready for contributors
