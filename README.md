# progress-tracker
A progress tracker go binary 

# Features
- **Add Books**: Add resource with title,chapter,pages and total page count
- **Update Books** : Updated resource read with a summary if necessary.
- **View Progress** : Track progress of resource being utilized.
- **Data Persistance** : Data stored on a json file, that can be converted to a detailed file.
- **Statistics** : Track last reads,start dates and reading duration.

# Installation
1. Install Go > 1.22.2
2. Clone the project.
```bash
    git clone https://github.com/JOOwino/progress-tracker.git
```
3. Navigate to the project directory.

# Running
```bash
make start
```

# Program options
1. Add New Book
2. Update Reading Progress
3. List All Books
4. Print Report
5. Exit