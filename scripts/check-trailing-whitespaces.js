/*
  Linter which checks if all files under git control do not contain any trailing
  white spaces (both spaces and tabs characters), moreover non-text files are
  excluded from check based on extension from array fileExtensionsToIgnore
  Requires git available in PATH and can be run only in a repository
*/

import {readFile} from 'fs'
import {spawnSync} from 'child_process'

const fileExtensionsToIgnore = ['.ico', '.png']

// get all files under git control
const gitListFiles = spawnSync('git', ['ls-tree', '-r', 'HEAD', '--name-only'])
if (gitListFiles.stderr.toString() !== "") {
    console.error(`Unexpected error occurred: ${gitListFiles.stderr.toString()}`)
    process.exit(2)
}

const filesToCheck = gitListFiles.stdout.toString().split('\n').filter(file =>
    fileExtensionsToIgnore.every(extension => !file.endsWith(extension))
)

const noTrailingWhitespaces = new RegExp(/[ \t]+$/gm)
filesToCheck.forEach(file => {
    readFile(file, 'utf8', (_, content) => {
        const match = noTrailingWhitespaces.exec(content);
        if (match) {
            console.error(`Found trailing whitespaces: ${file}:${findLineNumberByIndex(match.input, match.index)}`)
            process.exitCode = 1
        }
    })
})

function findLineNumberByIndex(input, index) {
    const lines = input.split('\n');
    let currentIndex = 0;

    for (let lineNumber = 0; lineNumber < lines.length; lineNumber++) {
        const line = lines[lineNumber];
        const lineLength = line.length + 1; // Add 1 for the newline character

        if (currentIndex + lineLength > index) {
            return lineNumber + 1; // Line numbers are usually 1-based
        }

        currentIndex += lineLength;
    }

    return -1; // Index is out of bounds
}