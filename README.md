
# Dead gRPC

A minimal gRPC project to kickstart development.

## Contributing to the Project

Welcome to the project! If you'd like to contribute, follow these steps to get started:

### 1. Clone the Repository
Clone the repository to your local machine:
```bash
git clone https://github.com/AlmirSai/dead_gRPC.git
cd dead_gRPC
```

### 2. Install Go
Make sure you have [Go](https://go.dev/dl/) installed. This project uses **Go 1.23.4**.

### 3. Switch to the `developer` Branch
All contributions should be made to the `developer` branch. Switch to this branch using:
```bash
git checkout developer
```

If the `developer` branch does not exist locally, create it by fetching from the remote:
```bash
git fetch origin developer:developer
git checkout developer
```

### 4. Initialize the Project
If this is the first time setting up the project, initialize the Go module:
```bash
go mod init github.com/AlmirSai/dead_gRPC
```

Install any required dependencies as needed:
```bash
go mod tidy
```

### 5. Add Your Contribution
1. Create a new feature branch off the `developer` branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Write your code and commit your changes:
   ```bash
   git add .
   git commit -m "Description of your changes"
   ```

3. Push your feature branch to the remote repository:
   ```bash
   git push origin feature/your-feature-name
   ```

4. Open a pull request (PR) on GitHub, targeting the `developer` branch, and describe your changes.

### 6. Running the Project
To run the project locally, use:
```bash
go run main.go
```

If `main.go` doesn't exist yet, create it as the project entry point.

### 7. Collaboration Workflow
- Keep your feature branch updated with the latest changes from the `developer` branch:
  ```bash
  git pull origin developer
  ```

- Use `Issues` to report bugs or suggest new features.
- Review pull requests from other contributors.

---

## Requirements
- Go 1.23.4 or later
- Basic understanding of gRPC

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
```

---

### Changes Made:
1. **Default branch for contributions:** Updated all instructions to use the `developer` branch.
2. **Feature branches:** Emphasized creating feature branches off `developer`.
3. **Pull requests:** Clearly mentioned targeting the `developer` branch for PRs.
4. **Fetching `developer` branch:** Added commands to ensure the branch is available locally.
