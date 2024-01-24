# webHooks

This simple GoLang application provides a local testing environment for services utilizing webhooks. It allows users to test webhook functionality without the need for external public services. Particularly useful when working with sensitive data, this server ensures that testing is secure and private.

## Features:

- **Local Testing Environment**: Run the server locally to simulate webhook events within a controlled environment.
- **Secure Testing with Sensitive Data**: Keep your sensitive data secure by avoiding the use of external public services for webhook testing.

## Getting Started:

### Prerequisites:

- [GoLang](https://golang.org/) installed on your machine.

### Installation:

1. Clone the repository:

   ```bash
   git clone https://github.com/rkiwi/webHooks.git
   cd webHooks
   ```
2. Build the executable:

   ```bash
   go build webHooks.go
   ```

### Usage:

1. #### Run the server (with flags if you want to customise output):
   ```bash
   ./webHooks
   ```

   The server provides several options that can be configured through command-line flags:

   -answer (bool)
   ```text
   Enable/disable answer from service (Status: OK, Data received) (default true)
   ```
   <img width="370" alt="image" src="https://github.com/rkiwi/webHooks/assets/68079296/fe5c0124-bc32-4f00-9a62-215ee84fac11">
   
   
   -headers  (bool)
   ```text
   Enable/disable technical information of request (default true)
   ```
   <img width="429" alt="image" src="https://github.com/rkiwi/webHooks/assets/68079296/034d9331-768b-452c-a2d6-3cf27f2384ec">


   -port int
   ```text
   Set listening port of webhook catcher (default 8080)
   ```


   -pretty  (bool)
   ```text
   Enable/disable pretty JSON formatting (default true)
   ```
   <img width="1390" alt="image" src="https://github.com/rkiwi/webHooks/assets/68079296/7c67dc8b-d899-4aa7-aff9-243750e6107b">

3. #### Sending Requests:

   Send your webhook requests to:
   ```text
   https://localhost:<your_port>
   ```
   Example:
   ```bash
   curl -X POST http://127.0.0.1:8080 \                                                                                                                                                                                        ─╯
   -d '{"key": "value"}'
   ```
4. #### Tracking Requests:

   Observe incoming webhook requests directly in the terminal where the server is running.

