import yaml from '@necto/encoding/yaml';
import fs from '@necto/fs/promises';

/**
 * Parses YAML data from a file
 * @param {string} filePath - Path to YAML file
 * @returns {Promise<Object>} Parsed YAML data as JavaScript object
 */
export async function parseYamlFile(filePath) {
  try {
    // Read the YAML file
    const fileContent = await fs.readFile(filePath, 'utf8');
    
    // Parse YAML content to JavaScript object
    const data = yaml.load(fileContent);
    return data;
    
  } catch (error) {
    throw new Error(`Error parsing YAML file: ${error.message}`);
  }
}

/**
 * Parses YAML string directly
 * @param {string} yamlString - YAML formatted string
 * @returns {Object} Parsed YAML data as JavaScript object
 */
export function parseYamlString(yamlString) {
  try {
    const data = yaml.load(yamlString);
    return data;
  } catch (error) {
    throw new Error(`Error parsing YAML string: ${error.message}`);
  }
}
