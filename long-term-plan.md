# Milestone 0 — Project Bootstrap (Repo Skeleton & Basic Health)

## Tasks
- [ ] Initialize repository (`go mod init`, LICENSE, CONTRIBUTING, CODE_OF_CONDUCT)
- [ ] Create directory structure:
  - [ ] `/cmd/server`
  - [ ] `/cmd/worker`
  - [ ] `/internal/app`
  - [ ] `/internal/domain`
  - [ ] `/internal/ports`
  - [ ] `/internal/adapters`
  - [ ] `/internal/transport`
  - [ ] `/internal/auth`
  - [ ] `/pkg/sdk`
  - [ ] `/migrations`
  - [ ] `/deploy`
  - [ ] `/tests`
  - [ ] `/scripts`
- [ ] Create minimal server in `/cmd/server` with:
  - [ ] `/healthz` route
  - [ ] `/readyz` route
- [ ] Add server `Dockerfile`
- [ ] Create `docker-compose.yml` with:
  - [ ] Postgres
  - [ ] Redis
  - [ ] NATS
  - [ ] (Optional) Service for API server
- [ ] Add `.env.example`
- [ ] Add GitHub Actions workflow:
  - [ ] `go test ./...`
  - [ ] `golangci-lint run`
  - [ ] Container build step
- [ ] Add `Makefile` with commands:
  - [ ] `make run`
  - [ ] `make test`
  - [ ] `make build`
  - [ ] `make compose-up`
- [ ] Add unit test for health endpoint
- [ ] Write README:
  - [ ] Summary
  - [ ] Setup steps
  - [ ] Dev flow
  - [ ] Milestone references

## Acceptance Criteria
- [ ] Server starts locally
- [ ] `curl localhost:<port>/healthz` returns 200
- [ ] `go test ./...` passes
- [ ] GitHub Actions runs successfully
- [ ] `docker-compose up` boots all infra services

---

# Milestone 1 — Core Data Model & Admin API

## Tasks
- [ ] Write SQL schema for:
  - [ ] tenants
  - [ ] projects
  - [ ] environments
  - [ ] flags
  - [ ] segments
  - [ ] api_keys
  - [ ] audit_logs
  - [ ] flag_versions
- [ ] Add DB migrations (`migrations` folder)
- [ ] Add DB adapter implementing repository interfaces
- [ ] Add usecases for:
  - [ ] Create tenant
  - [ ] Create project
  - [ ] Create environment
  - [ ] Create/update flag
- [ ] Add transport layer routes for Admin API
- [ ] Validate payloads
- [ ] Add unit tests for usecases
- [ ] Add integration tests (optional) using docker-compose DB

## Acceptance Criteria
- [ ] Migrations run cleanly
- [ ] Admin API endpoints work (manual curl test)
- [ ] Usecases fully tested
- [ ] Can create tenant → project → env → flag end-to-end

---

# Milestone 2 — Redis Cache & SDK Bootstrap Endpoint

## Tasks
- [ ] Implement Redis cache adapter
- [ ] Define cache keys: `flags:env:<env_id>:v<version>`
- [ ] Add usecase: fetch flags for environment (DB → cache)
- [ ] Add signed JSON bootstrap payload logic
- [ ] Implement endpoint: `/.well-known/flags/<sdk_key>`
- [ ] Build minimal JS SDK:
  - [ ] Fetch bootstrap
  - [ ] Cache in memory
  - [ ] Evaluate simple boolean flag
- [ ] Build minimal Go SDK:
  - [ ] Fetch bootstrap
  - [ ] Local evaluation
- [ ] Add tests for cache behavior:
  - [ ] Cache miss → DB hit
  - [ ] Cache hit → no DB hit

## Acceptance Criteria
- [ ] Bootstrap endpoint returns signed config
- [ ] SDK fetches & evaluates locally
- [ ] Cache hit ratio measurable via logs/metrics
- [ ] Payload signature verification works

---

# Milestone 3 — Streaming Updates & Background Workers

## Tasks
- [ ] Integrate NATS JetStream
- [ ] Create topics:
  - [ ] `flag.updated`
  - [ ] `flag.deleted`
  - [ ] `webhook.dispatch`
- [ ] Publish events on flag change
- [ ] Worker service:
  - [ ] Redis cache update subscriber
  - [ ] Webhook dispatcher w/ retry
- [ ] Add SSE / WebSocket endpoint for SDK real-time updates
- [ ] Update SDK (JS/Go) to support:
  - [ ] Subscribe to streaming updates
  - [ ] Fallback to polling if disconnected

## Acceptance Criteria
- [ ] Updating a flag updates Redis via worker
- [ ] Streaming clients receive updates real-time
- [ ] Worker can restart without missing messages (durable)

---

# Milestone 4 — Multi-tenancy & Security

## Tasks
- [ ] Add API key authentication middleware
- [ ] Tenant resolution logic (from API key)
- [ ] RBAC model (roles: owner/admin/member)
- [ ] Add per-tenant rate limiting
- [ ] Secure SDK endpoints:
  - [ ] Validate SDK key → environment → tenant link
- [ ] Implement payload signing using HMAC/RSA
- [ ] Update SDKs with signature validation
- [ ] Add audit logging for all admin actions

## Acceptance Criteria
- [ ] SDK cannot access another tenant’s flags
- [ ] API admin actions require correct JWT/API key
- [ ] Rate limits enforced
- [ ] Audit log entries recorded

---

# Milestone 5 — Deployment, K8s & Autoscaling

## Tasks
- [ ] Create Helm chart for:
  - [ ] server
  - [ ] worker
  - [ ] redis (or external)
  - [ ] postgres (or external)
  - [ ] nats
- [ ] Add Kubernetes manifests:
  - [ ] Deployments + Services
  - [ ] ConfigMaps + Secrets
  - [ ] HPAs (CPU / request latency)
  - [ ] KEDA ScaledObject for NATS/Redis backlog
- [ ] Add Terraform example for:
  - [ ] Managed Postgres
  - [ ] Managed Redis
  - [ ] NATS cluster
  - [ ] Load balancer
  - [ ] Object storage
- [ ] Add GitHub Actions workflow:
  - [ ] Build image
  - [ ] Run migrations
  - [ ] Deploy to cluster
- [ ] Add secrets handling: sealed-secrets or SSM

## Acceptance Criteria
- [ ] Cluster can scale API pods via HPA
- [ ] Worker autoscaling via queue depth
- [ ] Zero-downtime deploy works
- [ ] Helm chart runs locally via kind/minikube

---

# Milestone 6 — Observability, Load Testing & Documentation

## Tasks
- [ ] Add Prometheus metrics:
  - [ ] request latency
  - [ ] flag evaluation latency
  - [ ] cache hit/miss ratio
  - [ ] db/redis/nats metrics
- [ ] Add OpenTelemetry tracing:
  - [ ] eval path
  - [ ] DB queries
  - [ ] cache operations
- [ ] Add Grafana dashboards
- [ ] Write K6 load tests:
  - [ ] heavy read
  - [ ] cold start
  - [ ] streaming tests
- [ ] Add docs:
  - [ ] Self-host guide
  - [ ] API reference
  - [ ] SDK usage (JS + Go)
  - [ ] Architecture diagrams
  - [ ] Contribution guide

## Acceptance Criteria
- [ ] End-to-end tracing works in Jaeger/Grafana
- [ ] Load tests show p95 < target thresholds
- [ ] Documentation complete enough for contributors
- [ ] Self-host setup validated on clean machine

