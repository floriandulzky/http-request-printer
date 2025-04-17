# HTTP Request Printer

HTTP Request Printer is a terminal-based application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) that allows you to start an HTTP server and view incoming HTTP requests in real-time. The app provides a simple interface to inspect request details such as method, URL, headers, and body.

## Features

- Start an HTTP server on port 8000.
- View incoming HTTP requests in a user-friendly terminal interface.
- Navigate through multiple requests using keyboard controls.
- Supports real-time updates for incoming requests.

## Requirements

- Go 1.24 or later
- A terminal emulator

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/floriandulzky/http-request-printer.git
   cd http-request-printer
   ```

2. Build the application:

   ```bash
   go build -o http-request-printer
   ```

3. Run the application:

   ```bash
   ./http-request-printer
   ```

## Usage

1. Start the application by running the binary:

   ```bash
   ./http-request-printer
   ```

2. Press **Enter** to start the HTTP server on port 8000.

3. Send HTTP requests to `http://localhost:8000` using tools like `curl`, Postman, or your browser.

4. View the details of incoming requests in the terminal interface.

5. Use the following keyboard controls:
    - **Right Arrow**: Navigate to the next request.
    - **Left Arrow**: Navigate to the previous request.
    - **Ctrl+C**: Quit the application.

## Download Prebuilt Binary

If you prefer not to build the application yourself, you can download a prebuilt binary from the [Releases](https://github.com/floriandulzky/http-request-printer/releases) page.

1. Download the binary for your operating system.
2. Make the binary executable:

   ```bash
   chmod +x http-request-printer
   ```

3. Run the application:

   ```bash
   ./http-request-printer
   ```

---