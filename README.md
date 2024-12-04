# Password Generator

A simple password generator application built with Go (backend) and Vue.js (frontend).

## Features

- Customizable password length
- Options to include/exclude numbers, symbols, lowercase, and uppercase letters
- Copy to clipboard functionality
- Modern and responsive UI

## Running with Docker

The easiest way to run the application is using Docker Compose:

```bash
docker-compose up --build
```

The application will be available at `http://localhost:5173`

## Manual Setup

### Prerequisites

- Go 1.21 or later
- Node.js and npm

### Backend

1. Install Go dependencies:

```bash
go mod tidy
```

1. Run the backend server:

```bash
go run main.go
```

The server will start at `http://localhost:8080`

### Frontend

1. Navigate to the frontend directory:

```bash
cd frontend
```

1. Install dependencies:

```bash
npm install
```

1. Run the development server:

```bash
npm run dev
```

The frontend will be available at `http://localhost:5173`

## Usage

1. Open `http://localhost:5173` in your browser
1. Adjust the password length using the slider
1. Select which character types to include in your password
1. Click "Generate Password" to create a new password
1. Use the "Copy" button to copy the password to your clipboard
