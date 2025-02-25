# Algorithms for computation of the Levenshtein distance

This repository contains four implementations of the algorithm along with tests and benchmarks for each of them:
- Recursive algorithm
- Memoised recursive algorithm
- Algorithm utilizing full matrix approach
- Algorithm utilizing two rows approach

  ### Run tests and benchmarks locally:
  - Clone project
  ```bash
  git clone https://github.com/C-C-IM-31/levenshtein-distance.git
  ```
  - Navigate to project folder
  ```bash
  cd levenshtein-distance
  ```
  - Run tests
  ```bash
  go test
  ```
  - Run benchmarks
  ```bash
  go test -bench .
  ```
