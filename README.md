# Gear Framework

A lightweight Go CLI framework for rapid web application development with built-in database migrations, HTTP server management, and auto-recompilation features.

## Features

- **Database Migrations**: Automatic database schema migration using GORM
- **Development Server**: Hot-reload development server with file change detection
- **Production Server**: Production-ready HTTP server
- **CLI Interface**: Intuitive command-line interface using Cobra
- **Git Integration**: Automatic change detection using Git status
- **Testing Support**: Built-in test runner integration
- **SQLite Support**: Ready-to-use SQLite database driver

## Installation

```bash
go get github.com/sajad-dev/gear
```

## Quick Start

### 1. Initialize Your Project

```go
package main

import (
    "github.com/sajad-dev/gear/startup"
    "github.com/sajad-dev/gear/config"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Define your models
type User struct {
    gorm.Model
    Name  string
    Email string
}

func main() {
    // Database setup
    db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // HTTP server handler
    httpHandler := func(port int, db *gorm.DB, stop chan struct{}) error {
        // Your HTTP server implementation here
        // Use frameworks like Gin, Echo, or standard net/http
        return nil
    }

    // Bootstrap the application
    startup.Boot(config.ConfigSt{
        Db:          db,
        Tables:      []any{&User{}},
        Http:        httpHandler,
        APP_NAME:    "MyApp",
        DESCRIPTION: "My awesome application",
    })
}
```

### 2. Available Commands

After setting up your project, you can use the following commands:

#### Run Database Migrations
```bash
./your-app migration
```

#### Start Development Server (with hot-reload)
```bash
./your-app run dev -p 8080
```

#### Start Production Server
```bash
./your-app run production -p 8080
```

## Project Structure

```
your-project/
├── cmd/
│   └── main.go          # Your application entry point
├── internal/
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Database models
│   └── services/        # Business logic
├── migrations/          # Database migrations (optional)
└── main.go             # Gear bootstrap file
```

## Configuration

The `config.ConfigSt` struct accepts the following parameters:

- `Db`: GORM database instance
- `Tables`: Slice of model structs for auto-migration
- `Http`: HTTP server function handler
- `APP_NAME`: Application name for CLI
- `DESCRIPTION`: Application description

## Development Features

### Auto-Recompilation
The development server automatically detects file changes using Git status and restarts your application:

- Monitors file changes every 2 seconds
- Automatically stages changes with `git add .`
- Restarts the server when changes are detected
- Provides colored output for better debugging

### Database Migrations
Simple database migration system:

```go
// Your models will be automatically migrated
type User struct {
    gorm.Model
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
}

// Add to Tables slice in config
Tables: []any{&User{}, &Post{}, &Comment{}},
```

### HTTP Server Integration
Implement your HTTP server handler:

```go
func httpHandler(port int, db *gorm.DB, stop chan struct{}) error {
    // Example with Gin
    r := gin.Default()
    
    r.GET("/users", func(c *gin.Context) {
        var users []User
        db.Find(&users)
        c.JSON(200, users)
    })
    
    srv := &http.Server{
        Addr:    fmt.Sprintf(":%d", port),
        Handler: r,
    }
    
    go func() {
        <-stop
        srv.Shutdown(context.Background())
    }()
    
    return srv.ListenAndServe()
}
```

## Dependencies

- **GORM**: Database ORM
- **Cobra**: CLI framework
- **SQLite**: Database driver (default)
- **Color**: Terminal colors for better output

## Testing

Run tests for your application:

```bash
go test ./...
```

The framework includes built-in test utilities and supports standard Go testing patterns.

## CLI Flags

### run command
- `-p, --port`: Specify the port number (default: 8080)

### Common Usage Examples

```bash
# Start development server on port 3000
./myapp run dev -p 3000

# Run migrations
./myapp migration

# Start production server
./myapp run production -p 8080

# Get help
./myapp --help
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For issues and questions:
- Create an issue on GitHub
- Check the documentation
- Review the example implementations

---

**Note**: This framework requires Go 1.24.4 or higher and Git to be installed on your system for development features to work properly.
