# URL Shortener

## Introduction

This project is a URL shortener application that allows users to convert long URLs into shorter, more manageable links. It features a minimalistic frontend developed with React (TypeScript) and a "PROPER" backend powered by Golang. The project utilizes Tailwind CSS for styling and interacts with a PostgreSQL database to store URL mappings and track click statistics.

## Tech Stack

- **Backend:**
  - Golang
  - PostgreSQL for Database
  - Cookie-based authentication using Firebase (Google Provider)

- **Frontend:**
  - React with TypeScript
  - Tailwind CSS

## Features

- **URL Shortening:** Convert long URLs into shortened links for easier sharing.
- **Click Tracking:** Keep track of the number of clicks on each shortened link.
- **Secure:** Utilize cookie-based authentication with Google provider via Firebase for secure access to user data.

## Usage

1. **Shorten a URL:**
   - Enter a long URL into the input field.
   - Click the "Generate" button to generate a shortened URL. 
   - You can eneter a custom path if you want customization

2. **Share Shortened URLs:**
   - Copy the generated shortened URL and share it with others.

3. **Track Clicks:**
   - Monitor the number of clicks on each shortened URL.

## Authentication

The application ensures secure authentication by employing cookie-based authentication with Google provider. Each request requires the inclusion of the authentication cookie to make an API request, thereby safeguarding user privacy and data integrity.

## Deployment

The project is currently hosted at [shrink.bhaskaraa45.me](https://shrink.bhaskaraa45.me).

## Setup Instructions

1. **Backend:**
   - Clone the repository and navigate to the backend directory.
   - Build the backend using `make build`.
   - Configure the PostgreSQL database connection details in the backend.
   - Run the backend server with `make run`.

2. **Frontend:**
   - Navigate to the frontend directory.
   - Install dependencies with `npm install`.
   - Configure Firebase authentication with Google provider.
   - Start the frontend development server using `npm start`.

## Contributing

Contributions to the project are welcome! If you have any improvements or additional features to suggest, feel free to submit pull requests.

### Contact

For any inquiries or issues, please contact me at [shrink@bhaskaraa45.me](mailto:shrink@bhaskaraa45.me).
