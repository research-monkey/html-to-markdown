# HTML to Markdown Sidecar Service

A lightweight HTTP service that converts HTML to Markdown using the [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) library.

## Overview

This sidecar service runs alongside the main WebSearchAgent application to provide HTML-to-Markdown conversion capabilities via a simple HTTP API.

## Endpoints

### POST /convert
Converts HTML to Markdown.

**Request:**
- Method: `POST`
- Content-Type: `text/html` or `application/x-www-form-urlencoded`
- Body: Raw HTML content

**Response:**
- Content-Type: `text/plain; charset=utf-8`
- Body: Converted Markdown text

**Example:**
```bash
echo '<strong>Bold Text</strong>' | curl -X POST -d @- http://localhost:7777/convert
# Output: **Bold Text**
```

### GET /health
Health check endpoint.

**Response:**
- Status: `200 OK`
- Body: `OK`

## Building

### Local Build
```bash
go build -o html-to-markdown main.go
```

### Docker Build
```bash
docker build -t html-to-markdown-sidecar .
```

## Running

### Local
```bash
./html-to-markdown
# Server starts on port 7777
```

### Docker
```bash
docker run -p 7777:7777 html-to-markdown-sidecar
```

## Testing

Run the test script:
```bash
./test_server.sh
```

Or test manually:
```bash
# Health check
curl http://localhost:7777/health

# Convert HTML
curl -X POST -d '<h1>Hello World</h1>' http://localhost:7777/convert
```

## Integration with WebSearchAgent

The sidecar is designed to run alongside the WebSearchAgent container in Kubernetes:

```yaml
containers:
  - name: websearch-agent
    image: websearch-agent:latest
    # ... main container config
  
  - name: html-to-markdown
    image: html-to-markdown-sidecar:latest
    ports:
      - containerPort: 7777
    livenessProbe:
      httpGet:
        path: /health
        port: 7777
      initialDelaySeconds: 5
      periodSeconds: 30
```

From the main application, you can call the sidecar at `http://localhost:7777/convert`.

## Features

- Fast HTML to Markdown conversion
- Supports complex HTML structures
- Preserves formatting (bold, italic, lists, links, etc.)
- Lightweight Alpine-based Docker image
- Health check endpoint for Kubernetes
- No external dependencies in runtime

## Performance

- Built with Go for high performance
- Multi-stage Docker build for minimal image size
- Stateless design for easy horizontal scaling
