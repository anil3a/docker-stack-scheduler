# Docker Stack Scheduler (Go-based)
Manage Docker containers and stacks schedules

## Overview

Docker Stack Scheduler is a **lightweight, self-contained Docker application** that provides a **secure web UI** to:

* Discover running Docker containers and docker-compose stacks
* Group services by stack
* Schedule **start / stop / restart / pause** actions
* Automatically manage services during off-hours (e.g. night shutdown, morning startup)

The project is designed to be **native, minimal, fast, and secure**, with no reliance on host cron jobs or external orchestrators.

---

## Goals

* UI-first Docker lifecycle automation
* Stack-aware scheduling (docker-compose focused)
* Minimal runtime dependencies
* Small Docker image footprint
* Strong authentication & auditability
* Simple deployment via Docker

---

## Non-Goals

* Kubernetes support
* Multi-host orchestration (initially)
* Autoscaling
* Cloud-specific integrations
* CLI-only workflows

---

## Architecture

```
┌─────────────────┐
│   Web UI (SPA)  │  React or Vue (static build)
└────────┬────────┘
         │ HTTPS (JWT)
┌────────▼────────┐
│  Go Backend API │
│  + Scheduler    │
│  + Auth         │
└────────┬────────┘
         │ Docker Engine API
┌────────▼────────┐
│ Docker Socket   │
│ /var/run/docker │
└─────────────────┘
```

---

## Technology Stack

### Backend

* **Language:** Go
* **HTTP:** `net/http` (stdlib)
* **Scheduler:** `robfig/cron/v3`
* **Docker API:** `moby/docker/client`
* **Auth:** JWT + bcrypt
* **Database:** SQLite (embedded)

### Frontend

* React or Vue (Vite build)
* Static assets served by backend
* No runtime Node.js in production

### Storage

* SQLite database
* Single persistent volume

---

## Key Design Principles

* **Native first:** Direct Docker Engine API usage
* **Low dependency:** Avoid heavy frameworks
* **Single binary backend**
* **Secure by default**
* **Easy to self-host**

---

## Authentication & Security

### Authentication

* Local username/password authentication
* Passwords hashed using bcrypt
* JWT access tokens
* Token expiration and revocation

### Authorization (RBAC)

* `admin` – full access
* `operator` – manage schedules
* `viewer` – read-only access

### Security Features

* Mandatory authentication
* Audit logging
* Brute-force protection (rate limiting)
* Secure cookie / header handling

---

## Docker Discovery

### Container Detection

* Uses Docker Engine API
* Collects:

  * Container ID
  * Name
  * Image
  * State
  * Labels

### Stack Detection

Stacks are identified using docker-compose labels:

```
com.docker.compose.project
com.docker.compose.service
```

### Stack Status Logic

* **Running:** All containers running
* **Stopped:** All containers stopped
* **Partial:** Mixed state

---

## Scheduler

### Scheduler Engine

* Embedded cron engine
* Runs inside backend process
* Jobs persisted in database

### Supported Actions

* Start stack
* Stop stack
* Restart stack
* Pause / unpause containers

### Schedule Example

```
Stack: billing-stack
Action: STOP
Schedule: 0 23 * * *
Enabled: true
```

---

## Database Schema (Simplified)

### users

| Field         | Type     |
| ------------- | -------- |
| id            | integer  |
| username      | text     |
| password_hash | text     |
| role          | text     |
| created_at    | datetime |

### schedules

| Field       | Type              |
| ----------- | ----------------- |
| id          | integer           |
| target_type | stack / container |
| target_id   | text              |
| action      | start / stop      |
| cron        | text              |
| enabled     | boolean           |

### audit_logs

| Field     | Type     |
| --------- | -------- |
| id        | integer  |
| user      | text     |
| action    | text     |
| timestamp | datetime |

---

## API Endpoints (Initial)

### Authentication

```
POST /api/auth/login
POST /api/auth/logout
```

### Docker

```
GET /api/stacks
GET /api/containers
```

### Scheduler

```
GET    /api/schedules
POST   /api/schedules
PUT    /api/schedules/{id}
DELETE /api/schedules/{id}
```

---

## UI Features (MVP)

### Dashboard

* Stack status overview
* Upcoming scheduled actions

### Stack View

* Containers in stack
* Start / Stop buttons
* Assign schedules

### Scheduler UI

* Cron expression editor
* Enable / disable schedules
* Manual test execution

### Audit Log

* User action history
* Execution results

---

## Deployment

### Docker Compose

```yaml
services:
  stack-scheduler:
    image: stack-scheduler:latest
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - data:/data
    ports:
      - "8080:8080"
    restart: unless-stopped

volumes:
  data:
```

---

## Resource Targets

| Metric       | Target   |
| ------------ | -------- |
| Image size   | < 50 MB  |
| RAM usage    | < 50 MB  |
| Startup time | < 500 ms |
| Idle CPU     | ~0%      |

---

## Roadmap 

Look at the [ROADMAP.md](ROADMAP.md) file for detailed roadmap and milestones.


---

### 🎯 MVP Completion Definition

* Secure login
* Stack discovery
* Scheduled start/stop
* UI-driven management
* Single Docker container deployment

---

### Suggested Next Action

➡️ Start **Phase 1 → Backend Foundation + Docker Integration**
(This unlocks everything else)

---


## Project Name (Working)

* DockNap
* StackSleeper
* DockerCurfew
* Nocturne
* DockClock

---

## License

TBD (MIT or Apache 2.0 recommended)
