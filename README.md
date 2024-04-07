# Funda: Open-Source Accounting Software

Funda is a pioneering open-source accounting platform designed to meet the comprehensive needs of Australian businesses, with a vision to expand and adapt to global accounting standards. Developed with the flexibility to accommodate different financial regulations and reporting standards, Funda aims to become the go-to accounting solution for businesses around the world.

## Key Features

- **Australian Financial Compliance**: Tailored to meet the Australian accounting standards and tax regulations, including GST, BAS reporting, and superannuation management.
- **Global Scalability**: Architecturally designed to support extensions for different countries' accounting standards, making it adaptable for global use.
- **Multi-Currency Transactions**: Equipped to handle transactions in various currencies, facilitating international business operations.
- **Cloud-Based**: Ensures secure, real-time access from anywhere, enabling efficient financial management and collaboration.

## Getting Started

These instructions will help you set up Funda for development and testing purposes.

### Prerequisites

Ensure you have the following installed:

- Go (version 1.15 or later)
- Node.js and npm (or Yarn)
- Docker and Docker Compose (for containerized environments)

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/hlta/funda.git
cd funda
```

2. **Start the Backend**

Navigate to the backend directory, install Go dependencies, and run the server:

```bash
cd backend
go mod tidy
go run cmd/funda/main.go
```

3. **Launch the Frontend**

In a new terminal, set up the frontend:

```bash
cd frontend
npm install
npm run dev
```

Visit `http://localhost:3000` to view the app.

### Docker Setup (Optional)

Use Docker Compose to simplify the setup:

```bash
docker-compose up --build
```

## Contributing

We welcome contributions to Funda! Whether you're interested in adding features, fixing bugs, or extending support for other countries, please read our [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute.

## License

Funda is released under the GNU General Public License v3.0, ensuring it remains free and open-source. See [LICENSE](LICENSE) for full details.

## Acknowledgments

- Special thanks to all contributors for their support and dedication to the Funda project.
- Open-source projects and tools that have made Funda possible.
```
