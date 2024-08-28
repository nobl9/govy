/*
  Formatter which works on cspell config file and:
  - Sorts the 'words' list.
  - Removes duplicates from 'words' list.
*/

import YAML from 'yaml';
import { readFileSync, writeFileSync } from 'fs';

const CSPELL_CONFIG = "cspell.yaml"

function format() {
  const f = readFileSync(CSPELL_CONFIG, 'utf8')
  const yaml = YAML.parseDocument(f, { keepSourceTokens: true })

  let words = yaml.get('words')
  words.items.sort()
  let set = new Set()
  words.items = words.items.filter((w) => {
    if (!set.has(w.value)) {
      set.add(w.value)
      return true
    }
    return false
  })

  writeFileSync(CSPELL_CONFIG, yaml.toString())
}

try { format() } catch (err) { console.error(err) }
