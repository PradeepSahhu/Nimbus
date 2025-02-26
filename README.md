# Nimbus - Self-Hosted File Backup Server

## Project Overview
Nimbus is a **self-hosted local file backup server** designed for **personal use, small teams, and developers** who want an easy, reliable, and locally managed backup solution. Users can upload files via a web interface, and the system tracks them in a database while ensuring the actual file storage remains persistent on disk.

### Key Features (MVP)
- **File Upload & Management**: Upload, view, and manage files from a modern web UI.
- **Local Storage Support**: Files are stored locally on the host machine or an attached volume.
- **User Authentication**: Secure login system with roles (admin/user).
<!-- - **Scheduled Backups**: Users can configure automated backups. -->
- **Activity Logs**: Tracks uploads, deletions, and system events for easy debugging.
- **Dockerized Deployment**: Runs in isolated containers for easy setup and portability.

---

## 🚀 Current Development Progress
### **Backend (Golang + SQLite)** ✅
- [x] Golang backend initialized.
- [ ] API endpoints for authentication and file handling.
- [ ] Dockerized backend with persistent volume support.
- [ ] Implement file versioning.
- [ ] Implement scheduled backups.

### **Frontend (Next.js + Tailwind CSS)** ✅
- [x] Next.js frontend initialized.
- [ ] UI designed for file uploads and management.
- [ ] Connected API to frontend.
- [ ] Improve UI/UX (drag-and-drop upload, file preview, etc.).

### **Docker & Deployment** ✅
- [x] Dockerized backend & frontend.
- [x] Docker Compose setup for local development.
- [x] Persistent storage via Docker volumes.
- [ ] CI/CD Pipeline for automated builds & deployments.

---

## Getting Started
### **1. Clone the Repository**
```bash
 git clone https://github.com/yashpatil74/nimbus.git
 cd nimbus
```

### **2. Run with Docker**
```bash
 docker-compose up --build
```

### **3. Access the Web UI**
Open your browser and go to:
```
http://localhost:8081  # Frontend (User Interface)
http://localhost:8050  # Backend API
```

---

## Tech Stack
- **Frontend:** Next.js, Tailwind CSS
- **Backend:** Golang, SQLite, GORM
- **Storage:** Local file system (via Docker volumes)
- **Deployment:** Docker, Docker Compose

---

## Future Roadmap
- **File versioning system** (restore previous file versions).
- **Advanced user roles & multi-user support**.
- **Encryption for stored files** (optional security feature).
- **Cloud backup integrations (S3, Google Drive)**.
- **Mobile-friendly UI improvements**.

---

## Contributions & Feedback
This is an open-source project in active development. Feel free to submit issues or suggestions!

**Stay tuned for updates and feature improvements!**

