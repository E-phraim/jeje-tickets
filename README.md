# jeje-tickets
##API for ticketing platform

#Jeje Tickets API Documentation
Welcome to the Jeje Tickets API documentation. This API allows users to sign up, login, view events, purchase tickets, create events, and more. Below you will find details about the available routes and their functionalities.

##Base URL
###The base URL for all API endpoints is: https://api.jejetickets.com

Authentication Routes
Sign Up
POST /signup
Create a new user account.

Login
POST /login
Authenticate user and generate an access token.

Logout
POST /logout
Invalidate the current user's access token and log them out.

Event Routes
Get All Events
GET /events
Retrieve a list of all available events.

Get Event Details
GET /events/{eventId}
Retrieve detailed information about a specific event.

Purchase Ticket
POST /events/{eventId}/tickets
Purchase a ticket for a specific event.

Create Event
POST /events/create
Create a new event.

Update Event
PUT /events/{eventId}/update
Update an existing event.

Get Event Attendees
GET /events/{eventId}/attendees
Retrieve the list of attendees for a specific event.

Check-In Attendee
POST /events/{eventId}/check-in
Check-in an attendee for a specific event.

PDF Ticket Routes
Get PDF Ticket
GET /tickets/{ticketId}/pdf
Retrieve the PDF version of a purchased ticket.

Error Handling
The API follows standard HTTP status codes for indicating success or failure of requests. In case of errors, the response will include an error message with additional details.

Authentication and Authorization
This API uses JSON Web Tokens (JWT) for authentication and authorization. To access protected routes, include the access token in the Authorization header of your requests as follows: Authorization: Bearer {accessToken}.

Please refer to the API documentation for more detailed information on request payloads, response formats, and additional route-specific details.

Contributors
Your Name - Lead Developer
Feel free to contribute to this project by submitting bug reports or feature requests via GitHub Issues.

Happy ticketing!
