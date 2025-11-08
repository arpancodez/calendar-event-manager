# 📅 Calendar Event Manager

A simple and elegant calendar event manager built with **Go** backend and **HTML/JavaScript** frontend. This application provides a complete CRUD (Create, Read, Update, Delete) interface for managing calendar events.

## ✨ Features

- **Create Events**: Add new calendar events with title, description, start/end times, and location
- **View Events**: Display all events in a beautiful card-based layout
- **Delete Events**: Remove events with confirmation
- **RESTful API**: Clean REST API built with Go and Gorilla Mux
- **Responsive UI**: Modern, gradient-styled interface that works on all devices
- **In-Memory Storage**: Fast event storage using Go's concurrent-safe maps
- **CORS Enabled**: Frontend and backend can run on different ports

## 🛠️ Tech Stack

### Backend
- **Go** (Golang)
- **Gorilla Mux** - HTTP router and URL matcher
- **Native HTTP** - Go's standard library

### Frontend
- **HTML5**
- **CSS3** - Modern styling with gradients and animations
- **Vanilla JavaScript** - Fetch API for REST calls

## 📋 Prerequisites

Before running this application, make sure you have:

- **Go 1.21+** installed ([Download Go](https://golang.org/dl/))
- A modern web browser
- Terminal/Command prompt

## 🚀 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/arpancodez/calendar-event-manager.git
cd calendar-event-manager
```

### 2. Install Dependencies

```bash
go mod download
```

This will download the required Gorilla Mux package.

### 3. Run the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

You should see:
```
Server starting on port :8080...
```

### 4. Access the Application

Open your web browser and navigate to:
```
http://localhost:8080
```

You'll see the Calendar Event Manager interface!

## 📚 API Endpoints

The application exposes the following REST API endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/events` | Get all events |
| `GET` | `/api/events/{id}` | Get a specific event |
| `POST` | `/api/events` | Create a new event |
| `PUT` | `/api/events/{id}` | Update an existing event |
| `DELETE` | `/api/events/{id}` | Delete an event |

### Event Object Structure

```json
{
  "id": "1699876543",
  "title": "Team Meeting",
  "description": "Weekly team sync",
  "start_time": "2024-11-20T10:00:00Z",
  "end_time": "2024-11-20T11:00:00Z",
  "location": "Conference Room A"
}
```

## 💡 Usage Examples

### Create an Event via API

```bash
curl -X POST http://localhost:8080/api/events \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Team Meeting",
    "description": "Weekly sync",
    "start_time": "2024-11-20T10:00:00Z",
    "end_time": "2024-11-20T11:00:00Z",
    "location": "Conference Room"
  }'
```

### Get All Events

```bash
curl http://localhost:8080/api/events
```

### Delete an Event

```bash
curl -X DELETE http://localhost:8080/api/events/1699876543
```

## 🏗️ Project Structure

```
calendar-event-manager/
├── main.go          # Go backend server with API routes
├── go.mod           # Go module dependencies
├── index.html       # Frontend UI
├── .gitignore       # Git ignore rules
└── README.md        # This file
```

## 🔧 How It Works

1. **Backend**: The Go server uses Gorilla Mux for routing and maintains an in-memory event store
2. **Frontend**: HTML/JS interface makes fetch calls to the REST API
3. **CORS**: Middleware enables cross-origin requests
4. **Concurrency**: Thread-safe event storage using `sync.RWMutex`

## 🎨 UI Features

- **Gradient Background**: Beautiful purple gradient
- **Card Layout**: Events displayed in responsive grid
- **Form Validation**: Required fields marked
- **Hover Effects**: Interactive card animations
- **Emoji Icons**: Visual event details
- **Confirmation Dialogs**: Safe deletion with prompts

## 🐛 Troubleshooting

### Port Already in Use
If port 8080 is already occupied, you can change it in `main.go`:
```go
port := ":8080"  // Change to :8081 or any other port
```

### CORS Issues
If you're running frontend separately, ensure CORS middleware is enabled (already configured in the code).

### Module Not Found
Run `go mod download` to install dependencies.

## 🔮 Future Enhancements

- [ ] Event editing functionality
- [ ] Database persistence (SQLite/PostgreSQL)
- [ ] User authentication
- [ ] Event categories and tags
- [ ] Search and filter events
- [ ] Email reminders
- [ ] Calendar view (month/week/day)
- [ ] Event recurrence patterns

## 📝 License

This project is open source and available for learning purposes.

## 👤 Author

**Arpan**
- GitHub: [@arpancodez](https://github.com/arpancodez)

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

## ⭐ Show Your Support

Give a ⭐️ if you like this project!

---

**Happy Event Managing! 📅✨**
