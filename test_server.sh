#!/bin/bash

# Test script for HTML to Markdown sidecar

echo "Testing HTML to Markdown Sidecar"
echo "================================="

# Test 1: Health check
echo -e "\n1. Testing health endpoint..."
curl -s http://localhost:7777/health
echo ""

# Test 2: Simple HTML conversion
echo -e "\n2. Testing simple HTML conversion..."
echo '<strong>Bold Text</strong>' | curl -s -X POST -d @- http://localhost:7777/convert
echo ""

# Test 3: Complex HTML conversion
echo -e "\n3. Testing complex HTML conversion..."
cat <<EOF | curl -s -X POST -d @- http://localhost:7777/convert
<h1>Main Heading</h1>
<p>This is a paragraph with <strong>bold</strong> and <em>italic</em> text.</p>
<ul>
  <li>Item 1</li>
  <li>Item 2</li>
  <li>Item 3</li>
</ul>
<a href="https://example.com">Link to example</a>
EOF
echo ""

echo -e "\n================================="
echo "Tests completed!"
