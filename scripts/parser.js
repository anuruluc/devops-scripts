// parser.js
const fs = require('fs');
const path = require('path');

function parseFile(filePath) {
  const fileContent = fs.readFileSync(filePath, 'utf8');
  const lines = fileContent.split('\n');
  const data = [];

  lines.forEach((line) => {
    const trimmedLine = line.trim();
    if (trimmedLine &&!trimmedLine.startsWith('#')) {
      const [key, value] = trimmedLine.split('=');
      data.push({ key, value });
    }
  });

  return data;
}

function parseConfigFile(filePath) {
  const configFileContent = fs.readFileSync(filePath, 'utf8');
  const config = {};

  configFileContent.split('\n').forEach((line) => {
    const trimmedLine = line.trim();
    if (trimmedLine &&!trimmedLine.startsWith('#')) {
      const [key, value] = trimmedLine.split('=');
      config[key] = value;
    }
  });

  return config;
}

function parseJsonFile(filePath) {
  const fileContent = fs.readFileSync(filePath, 'utf8');
  const jsonData = JSON.parse(fileContent);

  return jsonData;
}

module.exports = {
  parseFile,
  parseConfigFile,
  parseJsonFile,
};