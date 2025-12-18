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

### 1.5. Album Management 
- Allow users to create albums to organize their media.

### 1.6. Privacy Storage
- Allow users to upload photos/videos to a restricted area, separate from the main gallery.

### 1.7. Caching 
- Utilize Redis to cache frequently accessed data for faster retrieval.

## 2. Non-Functional Requirements
These requirements define the constraints for the application.

### 2.1. Frontend/Usability
- The frontend will be written in React and Tailwind.
- Ensure a clean and easy to use user interface.

### 2.2. Backend
- The backend will be written in Go to ensure high performance.

### 2.3. Storage Infrastructure
- All media will be stored in a relational database (MySQL).

### 2.4. Performance
- The system will utilize Redis for caching to ensure low latency and fast retrieval.

### 2.5. Maintainability
- The codebase should aim for high cohesion and low coupling to ensure future updates are easy to manage.

### 2.6. Security
- Access control must be enforced for privacy storage feature and through OAuth for account verification.

### 2.7. Deployment & Environment
- Application must be fully containerized via Docker to ensure the system can run locally on all machines.

