# Snap Stash: Requirements Specification

## 1. Functional Requirements
These requirements define the behaviours and features of the application.

### 1.1. User Authentication
- Implement account creation & login.

### 1.2. Photo/Video Upload
- Allow users to upload photos and videos to the system storage.

### 1.3. Photo/Video Deletion
- Allow users to delete their uploaded photos and videos.

### 1.4. Timeline Sorting
- Organize and display media by date, functioning similarly to Snapchat memories.

### 1.5. Caching 
- Utilize Redis to cache frequently accessed data for faster retrieval.

### 1.6. Object Storage
- Utilize an object storage to store BLOB data.

## 2. Non-Functional Requirements
These requirements define the constraints for the application.

### 2.1. Frontend/Usability
- The frontend will be written in HTML, CSS, and Javascript(React).
- Ensure a clean and easy to use user interface.

### 2.2. Backend
- The backend will be written in Go to ensure high performance.

### 2.3. Storage Infrastructure
- Media files are stored in S3-compatible object storage (MinIO for local development and demos).
- The relational database (MySQL) stores metadata.

### 2.4. Performance
- The system will utilize Redis for caching to ensure low latency and fast retrieval.

### 2.5. Maintainability
- The codebase should aim for high cohesion and low coupling to ensure future updates are easy to manage.

### 2.6. Deployment & Environment
- Application must be fully containerized via Docker to ensure the system can run locally on all machines.

