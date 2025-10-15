# 📝 postAPI

A lightweight Go REST API for managing forum posts with SQLite database support.

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![REST API](https://img.shields.io/badge/REST-API-green?style=for-the-badge)

</div>

<details>
<summary>📋 Table of Contents</summary>

- [📝 postAPI](#-postapi)
  - [🔍 Overview](#-overview)
  - [✨ Features](#-features)
  - [🛠 Tech Stack](#-tech-stack)
  - [⚙️ Prerequisites](#️-prerequisites)
  - [🚀 Installation](#-installation)
    - [Local Development](#local-development)
    - [Using Docker](#using-docker)
  - [📖 Usage](#-usage)
    - [Base URL](#base-url)
  - [🔌 API Reference](#-api-reference)
    - [Request Structure](#request-structure)
    - [Available Actions](#available-actions)
      - [1. Create Post](#1-create-post)
      - [2. Get One Post](#2-get-one-post)
      - [3. Get All Posts](#3-get-all-posts)
      - [4. Delete Post](#4-delete-post)
      - [5. Update Post](#5-update-post)
  - [📁 Project Structure](#-project-structure)
  - [🐳 Docker Support](#-docker-support)
    - [Environment Variables](#environment-variables)
  - [🧪 Testing](#-testing)
    - [Manual Testing with cURL](#manual-testing-with-curl)
      - [Create a Post](#create-a-post)
      - [Get One Post](#get-one-post)
      - [Get All Posts](#get-all-posts)
      - [Delete a Post](#delete-a-post)
    - [Expected Responses](#expected-responses)
  - [📦 Dependencies](#-dependencies)
  - [🤝 Contributing](#-contributing)

</details>

## 🔍 Overview

postAPI is a RESTful API service designed for handling post operations in a forum application. It provides endpoints for creating, reading, updating, and deleting posts with support for categories, likes/dislikes, and user associations.

## ✨ Features

- Create posts with multiple categories
- Retrieve single or multiple posts
- Delete posts
- Update posts
- Like/Dislike functionality
- SQLite database integration
- Docker support
- JSON request/response format

## 🛠 Tech Stack

- **Language**: Go 1.20
- **Database**: SQLite3
- **Driver**: [go-sqlite3](https://github.com/mattn/go-sqlite3)
- **Containerization**: Docker

## ⚙️ Prerequisites

Before running this project, ensure you have the following installed:

- Go 1.20 or higher
- GCC (for SQLite3 compilation)
- Docker (optional, for containerized deployment)

## 🚀 Installation

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd postAPI
   ```

2. **Install dependencies**
   ```bash
   go get github.com/mattn/go-sqlite3
   # or
   go mod download
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8082`
   
   > **Note**: The database will be automatically initialized on first run

### Using Docker

1. **Build the Docker image**
   ```bash
   docker build -t postapi .
   ```

2. **Run the container**
   ```bash
   docker run -p 8082:8082 -v $(pwd)/databases:/app/databases postapi
   ```

## 📖 Usage

The API accepts POST requests with a JSON body containing an `action` field and an optional `body` field with the request data.

### Base URL
```
http://localhost:8082/
```

## 🔌 API Reference

### Request Structure

All requests follow this general structure:

```json
{
  "action": "<action_name>",
  "body": {
    // action-specific data
  }
}
```

### Available Actions

#### 1. Create Post

**Action**: `createPost`

**Request Body**:
```json
{
  "action": "createPost",
  "body": {
    "userID": 1,
    "categorie": ["Technology", "Programming", "Go"],
    "content": "Your post content here",
    "img": "optional-image-url"
  }
}
```

**Response**:
- **Status**: `201 Created`
- **Body**: `"New post created"`

---

#### 2. Get One Post

**Action**: `getOne`

**Request Body**:
```json
{
  "action": "getOne",
  "body": {
    "id": 1
  }
}
```

**Response**:
- **Status**: `200 OK`
- **Body**: Post object
```json
{
  "postID": 1,
  "userID": 1,
  "nickname": "username",
  "categorie": ["Technology", "Programming"],
  "likedBy": [],
  "dislikedBy": [],
  "content": "Post content",
  "img": "",
  "nbrLike": 0,
  "nbrDislike": 0,
  "createAt": "2025-10-15T10:30:00Z"
}
```

---

#### 3. Get All Posts

**Action**: `getAll`

**Request Body**:
```json
{
  "action": "getAll"
}
```

**Response**:
- **Status**: `200 OK`
- **Body**: Array of post objects

---

#### 4. Delete Post

**Action**: `delete`

**Request Body**:
```json
{
  "action": "delete",
  "body": {
    "id": 1
  }
}
```

**Response**:
- **Status**: `200 OK`
- **Body**: `"Post well deleted"`

---

#### 5. Update Post

**Action**: `update`

**Request Body**:
```json
{
  "action": "update",
  "body": {
    "postID": 1,
    "content": "Updated content",
    "categorie": ["Updated", "Categories"]
  }
}
```

**Response**:
- **Status**: `200 OK`
- **Body**: Success message

## 📁 Project Structure

```
postAPI/
├── config/
│   └── constants.go          # Configuration constants
├── databases/
│   └── sqlRequests/           # SQL query files
│       ├── createTable.sql
│       └── insertNewPost.sql
├── internals/
│   ├── dbManager/             # Database initialization
│   │   └── initDB.go
│   ├── handlers/              # HTTP request handlers
│   │   ├── createHandler.go
│   │   ├── deleteHandler.go
│   │   ├── getHandler.go
│   │   ├── mainHandler.go
│   │   └── update.go
│   └── tools/                 # Utility functions
│       └── utils.go
├── models/
│   ├── post.go               # Post model and methods
│   └── request.go            # Request structure
├── scripts/
│   ├── init.sh               # Initialization script
│   └── push.sh               # Deployment script
├── Dockerfile                # Docker configuration
├── go.mod                    # Go module file
├── go.sum                    # Dependency checksums
├── main.go                   # Application entry point
└── README.md                 # This file
```

## 🐳 Docker Support

The project includes a Dockerfile for easy deployment:

- **Base Image**: `golang:1.20-alpine`
- **Exposed Port**: `8082`
- **Volume**: `/app/databases` (for persistent data)

### Environment Variables

- `CGO_ENABLED=1` (required for SQLite3)

## 🧪 Testing

### Manual Testing with cURL

#### Create a Post
```bash
curl -X POST http://localhost:8082/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "createPost",
    "body": {
      "userID": 1,
      "categorie": ["Manga", "Anime", "Berserk"],
      "content": "I am the black swordsman"
    }
  }'
```

#### Get One Post
```bash
curl -X POST http://localhost:8082/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "getOne",
    "body": {
      "id": 1
    }
  }'
```

#### Get All Posts
```bash
curl -X POST http://localhost:8082/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "getAll"
  }'
```

#### Delete a Post
```bash
curl -X POST http://localhost:8082/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "delete",
    "body": {
      "id": 1
    }
  }'
```

### Expected Responses

| Action | Status Code | Response Body |
|--------|-------------|---------------|
| createPost | 201 Created | "New post created" |
| getOne | 200 OK | Post object (JSON) |
| getAll | 200 OK | Array of posts (JSON) |
| delete | 200 OK | "Post well deleted" |

## 📦 Dependencies

- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for Go

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

**Note**: This API is designed to work as part of a larger forum application ecosystem. Ensure proper authentication and authorization mechanisms are implemented in production environments

---

<div align="center">

**⭐ Star this repository if you found it helpful! ⭐**

Made with ❤️ from 🇸🇳

</div>