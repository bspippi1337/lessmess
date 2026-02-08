
## 🧪 Feature Request: Fuzzer Mode (Discovery + Self-Serve Detection)

### Goal
Add a **Fuzzer Mode** for safe API discovery that helps users:
- discover undocumented endpoints
- detect self-serve API key flows
- auto-generate temporary credentials when explicitly supported

### Core ideas
- GET / OPTIONS / HEAD only by default
- Rate-limited, user-initiated discovery
- Clear separation between documented vs heuristic results

### Key capabilities
- Endpoint probing (`/health`, `/status`, `/auth`, `/token`, `/apikey`, `/v1`, `/beta`)
- Detection of self-serve API key creation flows
- Optional auto-generation of temporary keys
- Never overwrite user-provided credentials

### UX
- Dedicated **Fuzzer / Discovery** mode
- Progress + findings panel
- AI commentary explaining detected patterns

### Architecture notes
- `core/fuzzer`
- detector plugins (self-serve, openapi, hints)
- results feed into presets + request builder

### Ethics
Client-side discovery only. No brute force, no auth bypass, no write operations without consent.

