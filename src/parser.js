const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
  }

  readFileSync() {
    try {
      return fs.readFileSync(this.filePath, 'utf8');
    } catch (error) {
      throw new Error(`Failed to read file: ${error.message}`);
    }
  }

  parseJson() {
    const fileContent = this.readFileSync();
    try {
      return JSON.parse(fileContent);
    } catch (error) {
      throw new Error(`Failed to parse JSON: ${error.message}`);
    }
  }

  parseYaml() {
    const yaml = require('js-yaml');
    const fileContent = this.readFileSync();
    try {
      return yaml.load(fileContent);
    } catch (error) {
      throw new Error(`Failed to parse YAML: ${error.message}`);
    }
  }
}

module.exports = Parser;