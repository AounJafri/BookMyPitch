# BookMyPitch - Online Cricket Ground Booking System

BookMyPitch is a web application that simplifies the process of booking cricket grounds online.  Users can browse available grounds, submit booking requests, and view their booking history. Administrators have a dedicated panel to manage grounds, approve bookings, and oversee users.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
    - [User Functionality](#user-functionality)
    - [Admin Functionality](#admin-functionality)
- [Installation & Setup](#installation--setup)
    - [Prerequisites](#prerequisites)
    - [Step-by-Step Setup](#step-by-step-setup)
- [API Documentation](#api-documentation)
    - [Authentication Routes](#authentication-routes)
    - [Admin Routes (Ground Management)](#admin-routes-ground-management)
    - [General Routes (Ground Access & Booking)](#general-routes-ground-access--booking)
- [Contribution Guidelines](#contribution-guidelines)
- [License](#license)

## Overview

BookMyPitch provides a user-friendly platform for booking cricket grounds.  It streamlines the booking process, making it easy for users to find and reserve their desired grounds. The admin panel empowers administrators with the tools to efficiently manage grounds and bookings.

## Features

### User Functionality

- User Registration & Login
- Browse available cricket grounds
- Request ground bookings
- View booking status and history
- Logout

### Admin Functionality

- Login as Admin
- Approve or reject ground bookings
- Create, Read, and Delete (CRD) cricket grounds
- View all users and their bookings

## Installation & Setup

### Prerequisites

- Install Go (1.18 or higher)
- Install MySQL

### Step-by-Step Setup

1. **Clone the Repository:**

```bash
git clone [https://github.com/yourusername/BookMyPitch.git](https://github.com/yourusername/BookMyPitch.git)  # Replace with your actual repository URL
cd BookMyPitch

2. **Install Dependencies:**

Bash

go mod tidy
Set Up the Database:
Run the SQL queries from Backend/pkg/config/queries.sql in your local MySQL database.  Create a database named bookmypitch_db (or adjust the DB_NAME in the .env file accordingly).

Configure Environment Variables:
Create a .env file in the /Backend/cmd/ directory. Add the following credentials, replacing the placeholders with your actual MySQL credentials:

DB_USER="your_mysql_username"
DB_PASSWORD="your_mysql_password"
DB_NAME="bookmypitch_db"
DB_HOST="localhost"
DB_PORT="3306"
Start the Application:
Bash

go run main.go
Or build and run:

Bash

go build -o bookmypitch
./bookmypitch
The application will run on localhost:8080.
