const { readFileSync } = require("node:fs");
const Ajv2020 = require("ajv/dist/2020").default;
const addFormats = require("ajv-formats");

const { schema, inputs } = JSON.parse(readFileSync(0, "utf8"));
const ajv = new Ajv2020({
  strict: true,
  // Govy permits partial type trees, required-only branches, and open tuples.
  strictTypes: false,
  strictRequired: false,
  strictTuples: false,
  allErrors: true,
});
addFormats(ajv, { mode: "full", keywords: false });
ajv.addKeyword("x-govy-omittedRules");
const validate = ajv.compile(schema);
const results = inputs.map((input) => {
  const valid = validate(input);
  return { valid, errors: validate.errors };
});

process.stdout.write(JSON.stringify(results));
