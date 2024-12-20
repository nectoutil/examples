import * as fs from '@necto/fs';
import * as path from '@necto/path';
import * as haml from '@necto/encoding/haml';

// Function to read a HAML file and convert it to HTML
function parseHamlFile(filePath: string): string {
    try {
        const hamlContent = fs.readFileSync(filePath, 'utf8');

        const htmlContent = haml.render(hamlContent);

        return htmlContent;
    } catch (error) {
        console.error('Error parsing HAML:', error);
        throw error;
    }
}

// Example hamle file:
// !!! 5
// %html
//   %head
//     %title Sample HAML
//   %body
//     %h1 Hello, HAML!
//     %p This is a sample HAML file.
const hamlFilePath = path.join(__dirname, 'example.haml');
const htmlOutput = parseHamlFile(hamlFilePath);
console.log(htmlOutput);
