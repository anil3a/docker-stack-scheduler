
## Roadmap (Checklist with Breakdown)

---

### ✅ Phase 1 — MVP (Core & Usable)

#### 🔹 Backend Foundation

* [x] Initialize Go module
* [x] Basic HTTP server (`net/http`)
* [x] Health check endpoint
* [x] Config loader (env + file)
* [x] Graceful shutdown handling

#### 🔹 Docker Integration

* [ ] Connect to Docker socket
* [ ] List containers
* [ ] Read container labels
* [ ] Detect docker-compose stacks
* [ ] Group containers by stack
* [ ] Determine stack status (running / stopped / partial)

#### 🔹 Authentication & Security

* [ ] User model
* [ ] Password hashing (bcrypt)
* [ ] Login endpoint
* [ ] JWT generation & validation
* [ ] Auth middleware
* [ ] Role-based access (admin / operator / viewer)
* [ ] Initial admin bootstrap

#### 🔹 Scheduler Core

* [ ] Embedded cron engine
* [ ] Schedule model
* [ ] Persist schedules in SQLite
* [ ] Start stack action
* [ ] Stop stack action
* [ ] Manual run (test execution)
* [ ] Enable / disable schedules

#### 🔹 UI (Minimal but Functional)

* [ ] Login screen
* [ ] Dashboard (stack list + status)
* [ ] Stack detail view
* [ ] Start / Stop buttons
* [ ] Schedule creation form
* [ ] Schedule list

#### 🔹 Persistence & Audit

* [ ] SQLite integration
* [ ] DB migrations
* [ ] Audit log table
* [ ] Log user actions
* [ ] Log scheduler executions

#### 🔹 Packaging

* [ ] Single binary build
* [ ] Minimal Dockerfile
* [ ] docker-compose example
* [ ] Volume for data persistence

---

### 🟡 Phase 2 — Advanced Scheduling & UX

#### 🔹 Scheduling Enhancements

* [ ] Time window schedules (sleep mode)
* [ ] Restart schedules
* [ ] Pause / unpause containers
* [ ] Stack dependency ordering
* [ ] Conflict detection (overlapping jobs)

#### 🔹 UX Improvements

* [ ] Human-readable cron editor
* [ ] Schedule preview (“next run”)
* [ ] Bulk stack actions
* [ ] Dark mode
* [ ] UI notifications (success/failure)

#### 🔹 Observability

* [ ] Execution history UI
* [ ] Error surfacing in UI
* [ ] Stack health indicators
* [ ] Execution duration metrics

#### 🔹 Security Hardening

* [ ] Rate limiting
* [ ] Account lockout
* [ ] Token revocation
* [ ] Configurable password policy

---

### 🔵 Phase 3 — Power Features & Scale

#### 🔹 Platform Capabilities

* [ ] Multiple Docker host support
* [ ] Remote Docker (TLS)
* [ ] Host grouping
* [ ] Per-host schedules

#### 🔹 Reliability

* [ ] HA mode (leader election)
* [ ] Scheduler state recovery
* [ ] Missed job reconciliation
* [ ] Backup & restore schedules

#### 🔹 Integrations

* [ ] Webhooks
* [ ] Email notifications
* [ ] Slack / Discord alerts
* [ ] API tokens for automation

#### 🔹 Enterprise (Optional)

* [ ] LDAP / SSO
* [ ] 2FA (TOTP)
* [ ] Fine-grained RBAC
* [ ] Immutable audit logs

---

### 🧭 Phase 4 — Polishing & Open Source Readiness

* [ ] Documentation site
* [ ] Example configs
* [ ] Security audit
* [ ] Contribution guide
* [ ] CI/CD pipeline
* [ ] Automated tests
* [ ] Release artifacts

---
