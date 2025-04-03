## Overview

This project is a full-stack application for managing virtual machines. It includes:
- A **Go-based backend** for handling API requests, WebSocket connections, and database operations.
- A **Vue 3-based frontend** for user interaction, including VM management and real-time notifications.

---


## How to run the project

### Docker Setup

This project includes a docker-compose.yml file for local development you need to have Docker on your machine.

Create a .env file in the root directory with the following variables

| Variable                  | Description                        |
|---------------------------|------------------------------------|
| `DB_URL`                  | Database connection URL           |
| `JWT_SECRET`              | Secret key for JWT generation     |
| `USER_PASSW`             | Admin and client user password               |
| `POSTGRES_PASSWORD` | User password for Postgres    |

You can find a .env.example file to clarify it.

---

### Services

- **Backend**: Runs the Go API on port `3000`.
- **Frontend**: Runs the Vue app on port `80`.
- **Database**: PostgreSQL database on port `5432`.

### Usage

1. **Start all services**:
   ```sh
   docker-compose up --build
   ```

After this step go to `localhost` and login as an admin user `admin@test.com` or a client user `client@test.com`. Both users has the same password that you define in the env file.


2. **Stop all services**:
   ```sh
   docker-compose down
   ```

---

## Access the deployed project

Go to [https://d1w3z9l8n19jy.cloudfront.net/](https://d1w3z9l8n19jy.cloudfront.net/) and login with the follow credentials:

* `Admin user`: admin@test.com - prueba1234
* `Client user`: client@test.com - prueba1234

After the login you'll see the list of virtual machine creates, if you are an admin you can click the button in the top rigth corner to create a new one. If you click one vm of the list it will redirect you to the detail view, as a client you can't do more, but as a an admin you can edit or delete the vm.




## Roles and Permissions

- **Admin Role:** Can create, update, and delete VMs.
- **Client Role:** Can view VMs and their details.

---

## Folder Structure

```
.
├── api/                # Backend (Go)
│   ├── src/            # Source code
│   ├── Dockerfile      # Docker configuration for the backend
│   ├── go.mod          # Go module dependencies
│   └── fly.toml        # Fly.io deployment configuration
├── app/                # Frontend (Vue 3)
│   ├── src/            # Source code
│   ├── Dockerfile      # Docker configuration for the frontend
│   ├── vite.config.js  # Vite configuration
│   └── package.json    # Node.js dependencies
├── docker-compose.yml  # Docker Compose configuration
├── .env                # Env file for Docker
└── .github/            # GitHub Actions workflows
```

---

## Backend (API)

### Features

- **Authentication**: JWT-based authentication with role-based access control.
- **Virtual Machine Management**: Create, update, delete, and list VMs.
- **WebSocket Notifications**: Real-time updates for VM events (created, updated, deleted).
- **Database**: PostgreSQL integration using GORM.

### Endpoints

| Method | Endpoint         | Description                          | Auth Required | Role Required |
|--------|------------------|--------------------------------------|---------------|---------------|
| POST   | `/api/login`     | Authenticate user and return a token | No            | N/A           |
| GET    | `/api/vms`       | List all VMs                         | Yes           | Any           |
| GET    | `/api/vms/{id}`  | Get details of a specific VM         | Yes           | Any           |
| POST   | `/api/vms`       | Create a new VM                      | Yes           | Admin         |
| PUT    | `/api/vms/{id}`  | Update a VM                          | Yes           | Admin         |
| DELETE | `/api/vms/{id}`  | Delete a VM                          | Yes           | Admin         |

### WebSocket

- **Endpoint**: `/ws`
- **Description**: Provides real-time notifications for VM events.
- **Message Format**:
  ```json
  {
    "event_type": "created | updated | deleted",
    "message": "string"
  }
  ```

---

## Frontend (App)

### Features (Frontend)

- **User Authentication**: Login functionality with JWT token storage.
- **Dashboard**: Displays a list of VMs with details.
- **VM Management**: Create, update, and delete VMs (admin only).
- **Real-Time Notifications**: Toast notifications for VM events using WebSocket.
- **Responsive Design**: Built with Bootstrap for a modern UI.


### Available Scripts

Run these commands in the app directory:

- **Install dependencies**:
  ```sh
  npm install
  ```
- **Start development server**:
  ```sh
  npm run dev
  ```
- **Build for production**:
  ```sh
  npm run build
  ```

---

## Deployment

### Backend Deployment

The backend is configured to deploy on **Fly.io**. The deployment process is automated using the GitHub Actions workflow in deploy-back.yml.

### Frontend Deployment

The frontend is configured to deploy to **AWS S3** and **CloudFront**. The deployment process is automated using the GitHub Actions workflow in deploy-front.yml.

---
