# Gravity Simulation

## Overview
The Gravity Simulation project is a Go-based application designed to simulate gravitational interactions between celestial bodies. It provides tools for modeling, detecting collisions, and simulating the behavior of planets, dust, and other objects in a 3D space.

## Features
- **Collision Detection**: Detect collisions between celestial bodies.
- **Planet Factory**: Generate planets with customizable properties.
- **Vector Operations**: Perform 3D vector calculations.
- **Octree Implementation**: Efficient spatial partitioning for large-scale simulations.

## Project Structure
```
.
├── go.mod                # Go module definition
├── go.sum                # Dependency checksums
├── main.go               # Entry point of the application
├── collision/            # Collision detection logic
├── commands/             # Command-line utilities for simulation
├── models/               # Core data models (e.g., Body, Octree)
├── planets/              # Planet generation logic
├── test/                 # Experimental and test scripts
└── utils/                # Utility functions (e.g., vector operations)
```

## Getting Started

### Prerequisites
- Go 1.20 or later

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/fernandomedin/gravity-simulation.git
   cd gravity-simulation
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Simulation
To start the simulation, run:
```bash
go run main.go
```

## Ref

- https://www.raylib.com/cheatsheet/cheatsheet.html
- https://www.raylib.com/cheatsheet/raymath_cheatsheet.html
- https://github.com/gen2brain/raylib-go
- https://www.youtube.com/watch?v=q_edsSpDzHg
- https://en.wikipedia.org/wiki/Barnes%E2%80%93Hut_simulation
- The road to reality book, by Roger Penrose