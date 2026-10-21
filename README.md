# SO-U | Sistem Operasi Usaha

A Web Application for operating Indonesian Business.

## Tech Stack

- **Frontend :** Nextjs Typescript
- **Frontend Runtime :** Node
- **Frontend Host** : Vercel
- **Backend :** Go Fiber
- **Backend Database Library** : pgx
- **Backend Caching :** Redis
- **Media Database :** Cloudinary
- **Primary Database :** Neon
- **Compatibility :** Docker
- **CI/CD :** Github Actions

## Getting Started For Developers

### 1. Clone Repository

```bash
git clone https://github.com/machyon/Pasok-Tani.git
cd Pasok-Tani
```

### 2. Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Set up environment variables (pull from Neon development branch or copy `.env.example`):
   ```bash
   # Install Neon CLI if not already installed
   npm i -g neon@latest
   neon auth
   neon env pull --project-id "plain-morning-01562545" # Replace with your own project-id
   ```
   Alternatively, copy the example template:
   ```bash
   cp .env.example .env
   ```
3. Install dependencies and run the server:
   ```bash
   go mod tidy
   go run cmd/server.go
   ```

### 3. Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the development server:
   ```bash
   npm run dev
   ```
   Open [http://localhost:3000](http://localhost:3000) in your browser.

