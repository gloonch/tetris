# Terminal Tetris (Go)

## About
This project is a simple Tetris clone written in pure Go, inspired by the 2023 Tetris movie and the story of Alexey Pajitnov. After watching the movie, I thought: if Alexey could build Tetris in C, I can definitely build it in Go without any third-party libraries! So, I started this project the same day, aiming for a clean, dependency-free, terminal-based Tetris experience.

## Features
- 100% Go, no external libraries
- Runs in your terminal
- Classic Tetris gameplay: moving, rotating, and dropping pieces
- Scoring system with combos for clearing multiple lines
- Pause and resume functionality

## How to Run
1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd tetris
   ```
2. **Build and run:**
   ```sh
   go run .
   ```

## Controls
- **Left Arrow**: Move piece left
- **Right Arrow**: Move piece right
- **Down Arrow**: Move piece down faster
- **R**: Rotate piece 90° clockwise
- **P**: Pause/Resume the game
- **Q**: Quit the game
