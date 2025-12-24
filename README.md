### 🧑🏻‍💻 Snap Stash
---
Snap-Stash is a backend-focused media storage system built with Go, React, MySQL, Redis, and a S3-compatible object storage.

### 🧑‍🔧 Architecture Overview
---
* React frontend (Vite) communicates with Go REST API
* Go backend handles auth, media upload, retrieval, and deletion
* MySQL stores media metadata
* MinIO (S3-compatible) stores media files
* Redis caches recently accessed media

### ⭐️ Features
---
* 💨 Caching: Uses Redis to cache recently viewed media for rapid retrieval
* 📲 Upload: Allows user to upload media
* 🗑️ Delete: Allows user to delete media
* 👮‍♂️ Security: Account creation and login required to access user media
* 🖌️ Simple UI: Easy to navigate UI + a clean and aesthetic touch.
* 💽 Efficent Storage: Media is stored in S3-compatible object storage (MinIO), not the database
* 🚀 Streaming: Large media files are streamed directly from object storage
* 🧰 Auto Setup: Infrastructure and storage buckets are initialized automatically at startup

### 📝 Requirements
---
* Docker (Docker Engine + Docker Compose)
* Go 1.21+
* Node.js 18+

### 💻 Demo Instructions
---
#### Local Setup
```bash
git clone <REPO_URL>
cd Snap-Stash
cd infra
docker compose up -d
```
Open a new terminal and run:
```bash
cd Snap-Stash
cd backend
go run ./cmd/api
```
Open another terminal and run:
```bash
cd Snap-Stash
cd frontend
npm install
npm run dev
```

### Demo Walkthrough
---
1. Open the link in the frontend terminal
2. Create an account
3. Upload a photo or a video
4. View media in the gallery
5. Click media to go fullscreen
6. Delete media by clicking the X on the top right of the photo/video

### Stopping the Demo
---
Stop the backend & frontend with control + C in their terminals.

```bash
docker compose down
```
To wipe all data
```bash
docker compose down -v
```

### 📸 Preview
<p align="center">
  <img width="578" height="974" alt="Snap Stash Demo" src="https://github.com/user-attachments/assets/17c5c46a-6bc6-4218-8100-b1afedd6ca17" />
</p>p>


