#!/bin/bash

# Euclidean Rhythm Generator - Build and Run Script
# This script helps you easily run different components of the generator

set -e

echo "🎵 Euclidean Rhythm Generator"
echo "============================="

case "${1:-help}" in
    "generate"|"gen")
        echo "Generating default Euclidean rhythm..."
        go run main.go
        ;;
    "demo")
        echo "Showing rhythm pattern demonstrations..."
        go run demo.go
        ;;
    "examples")
        echo "Generating collection of example rhythms..."
        go run examples.go
        ;;
    "clean")
        echo "Cleaning generated files..."
        rm -f *.wav main
        rm -rf examples/
        echo "✓ Cleaned"
        ;;
    "build")
        echo "Building main generator..."
        go build main.go
        echo "✓ Built 'main' executable"
        ;;
    "help"|*)
        cat << EOF

Usage: $0 [command]

Commands:
  generate, gen  Generate default Euclidean rhythm (main.go)
  demo          Show interactive pattern demonstrations (demo.go)  
  examples      Generate collection of example rhythms (examples.go)
  build         Build the main generator executable
  clean         Remove generated files and directories
  help          Show this help message

Examples:
  $0 generate     # Create euclid.wav with default 16/6 pattern
  $0 demo         # See patterns from world music traditions
  $0 examples     # Generate multiple rhythm examples
  $0 clean        # Clean up generated files

EOF
        ;;
esac