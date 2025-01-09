# gRPC Tutorial
Check the main link: [gRPC Tutorial](https://habr.com/ru/articles/774796/)

## Quick Tutorial

### How to collaborate with this project:

0. **Add remote repository**
```bash
git remote add origin git@github.com:AlmirSai/dead_gRPC.git
```
1. **Clone this repository**
```bash
git clone git@github.com:AlmirSai/dead_gRPC.git
```
2. **Install dependencies**
```bash
go mod tidy
```
3.	Git add files
- Add all files to git
```bash
git add.
```
- Add a specific file to git
```bash
git add <file_name>
```
4.	Git commit
```bash
git commit -m "commit message"
```
5.	Git push
```bash
git push -u origin developer
```
### (Note: Always push to the developer branch. Do not push directly to main or stage.)

6.	Update your local branch
```bash
git fetch
```

Branching Workflow:
	•	developer: For active development. Always create your feature branches from developer.
	•	stage: For staging and testing. Only Pull Requests from developer should be merged here.
	•	main: For production. Only Pull Requests from stage should be merged into main.
