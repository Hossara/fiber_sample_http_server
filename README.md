# 🧪 Fiber HTTP Handler Example

This project is a simple HTTP server written in Go using the [Fiber](https://github.com/gofiber/fiber) web framework. It demonstrates how to:

- Serve a basic REST API (`GET /api/v1/items`)
- Handle success and error responses
- Use clean project architecture (handlers, models, stores)
- Write table-driven unit tests using `httptest` and Fiber’s testing tools

---

## 🗂 Project Structure

```
.
├── app/            # Server initialization and routing
├── cmd/            # Main entry point for the application
├── config/         # App configuration and loader
├── handlers/       # HTTP handlers and their tests
├── model/          # Data model definitions (Item)
├── store/          # In-memory data store implementation
├── config.json     # Example runtime configuration
├── Makefile        # Dev shortcuts (run, test)
```

---

## 🏁 Getting Started

### 1. Install Go

Ensure you have **Go 1.20+** installed:  
👉 https://golang.org/dl/

---

### 2. Clone and Run

```bash
git clone https://github.com/Hossara/fiber_sample_http_server.git
cd fiber_sample_http_server
make up
```

By default, the server runs at:  
👉 `http://localhost:8080/api/v1/items`

---

### 3. Example Response

```bash
curl http://localhost:8080/api/v1/items
```

Successful response (status `200`):
```json
[
  {
    "id": 1,
    "name": "Laptop",
    "description": "Thin and light"
  },
  {
    "id": 2,
    "name": "Phone",
    "description": "Flagship phone"
  }
]
```

Error response (status `404`):
```json
{
  "error": "no items found"
}
```

---

## ⚙️ Configuration

The app reads its configuration from a JSON file (default: `./config.json`).

Example `config.json`:

```json
{
  "server": {
    "port": 8080,
    "host": "127.0.0.1"
  }
}
```

You can override the config path using:

```bash
go run cmd/main.go --config=path/to/config.json
```

Or set the environment variable:

```bash
CONFIG_PATH=./config.sample.json make up
```

---

## 🧪 Running Tests

```bash
make test
```

This will run all tests under the `handlers/` directory, including:

- ✅ Items present: returns 200 OK with JSON array
- ❌ No items: returns 404 Not Found
- More cases can be easily added via table-driven tests

---

## ✨ Features & Highlights

- ✅ Clean, modular architecture
- 📦 Uses `fiber` for fast and expressive routing
- 🔁 Table-driven unit tests with `httptest`
- 📄 Uses custom `Store` interface for clean testability
- 🧪 Supports dependency injection for handlers

---

## 🧱 Tech Stack

- **Language:** Go (1.24.2)
- **Web Framework:** [Fiber](https://github.com/gofiber/fiber)
- **Testing:** `testing`, `httptest`
- **Config:** JSON + viper

---

## 📜 License

MIT License. See [LICENSE](./LICENSE) for details.