#!/bin/bash

# Project Name
PROJECT_NAME="scribe"

# Create top-level directory
mkdir -p $PROJECT_NAME
cd $PROJECT_NAME

# Create subdirectories
# mkdir -p backend/{src/{routes,models,controllers,services},tests}
mkdir -p frontend/{src/{components,pages,styles,utils},public}
mkdir -p docs
mkdir -p scripts
mkdir -p tests
mkdir -p shared

# Create top-level files
touch .gitignore README.md LICENSE

# Backend-specific files
# touch backend/{src/app.js,tests/.gitkeep,.env.example,package.json,Dockerfile}

# Frontend-specific files
touch frontend/{src/{App.js,index.js},public/index.html,package.json,webpack.config.js}

# Shared utilities
touch shared/{constants.js,utils.js}

echo "Project structure for $PROJECT_NAME created successfully!"
