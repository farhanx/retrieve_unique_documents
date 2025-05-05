# Media File Filter & Deduplicator

This Go program scans a specified folder for media files (e.g., images and videos), removes duplicates based on file content (using SHA-256 hashing), and moves the unique files into a subfolder called `filtered_files`.

## Features

* Supports filtering of common image and video formats.
* Uses SHA-256 hashing to detect and skip duplicate files.
* Organizes unique files into a new subdirectory.
* Provides detailed logs of the process.

## Supported File Types

* Images: `.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`
* Videos: `.mp4`, `.mov`, `.avi`, `.mkv`, `.webm`

## How It Works

1. Prompts the user for the path to a folder.
2. Recursively walks through the folder.
3. For each supported file:

   * Computes a SHA-256 hash of the file contents.
   * Skips files with duplicate hashes.
   * Moves unique files to the `filtered_files` subfolder.

## Usage

### Build the Program

```bash
go build -o media_filter
```

### Run the Program

```bash
./media_filter
```

When prompted, enter the **full path** to the folder you want to scan.

Example:

```
Enter the full path to your folder: /Users/yourname/Downloads/my_media
```

## Example Output

```
Moved 1 to filtered_files: photo1.jpg
Duplicate skipped: photo1_copy.jpg
Moved 2 to filtered_files: video1.mp4
```

## Error Handling

* If the specified directory doesn't exist, an error message is shown.
* File operation errors (e.g., moving or hashing failures) are logged but do not stop the program.

## License

This project is open-source and free to use under the [MIT License](https://opensource.org/licenses/MIT).

---

Would you like a `.gitignore` and `go.mod` file to go along with this?
