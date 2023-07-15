# jeje-tickets
##API for ticketing platform

# Jeje Tickets API Documentation

Welcome to the Jeje Tickets API documentation. This API allows users to sign up, login, view events, purchase tickets, create events, and more. Below you will find details about the available routes and their functionalities.

## Base URL

The base URL for all API endpoints is: `https://jeje-tickets.onrender.com/api/jejetickets`

## Authentication Routes

### Sign Up

- `POST /auth/register`

Create a new user account.

### Login

- `POST /auth/login`

Authenticate the user and generate an access token.

### Logout

- `POST /auth/logout`

Invalidate the current user's access token and log them out.

## Event Routes

### Get All Events

- `GET /api/events?page=1&limit=10`

Retrieve a list of all available events.

### Get Event Details

- `GET /events/{eventId}`

Retrieve detailed information about a specific event.

### Create Event

- `POST /events/create`

Create a new event.

### Update Event

- `PUT /events/{eventId}/update`

Update an existing event.

## Error Handling

The API follows standard HTTP status codes for indicating success or failure of requests. In case of errors, the response will include an error message with additional details.

## Authentication and Authorization

This API uses JSON Web Tokens (JWT) for authentication and authorization. To access protected routes, include the access token in the `Authorization` header of your requests as follows: `Authorization: Bearer {accessToken}`.

Please refer to the API documentation for more detailed information on request payloads, response formats, and additional route-specific details.

## Contributors

- [Kodeforce](https://github.com/E-phraim) - Lead Developer

Feel free to contribute to this project by submitting bug reports or feature requests via GitHub Issues.

Happy ticketing!
