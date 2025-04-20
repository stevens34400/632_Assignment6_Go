# 632_Assignment6_Go

## Overview

This Go program demonstrates a simple concurrent task processing system using goroutines and channels. It simulates processing tasks and writes the results to an output file.

## Features

- **Task Processing**: The program creates a set of tasks and processes them concurrently using multiple workers.
- **Concurrency**: Utilizes goroutines and channels to handle tasks concurrently.
- **Synchronization**: Uses a mutex to synchronize writes to the output file.
- **Logging**: Logs the progress and completion of each task.

## How It Works

1. **Task Creation**: The program creates 20 tasks, each represented by a `Task` struct.
2. **Worker Initialization**: Four worker goroutines are started to process tasks concurrently.
3. **Task Processing**: Each worker processes tasks from the channel, simulating computation with a delay.
4. **Result Writing**: Processed task results are written to `output/results.txt`.
5. **Completion**: The program logs the completion of all tasks.

## Usage

1. **Run the Program**: Execute the program using the Go command:
   ```bash
   go run main.go
   ```
2. **Output**: Check the `output/results.txt` file for the processed task results.

## Requirements

- Go 1.16 or later

## Notes

- Ensure the `output` directory is writable, as the program creates and writes to `output/results.txt`.
- The program simulates task processing with a 500ms delay for demonstration purposes.
