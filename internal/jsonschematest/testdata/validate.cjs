const { readFileSync } = require("node:fs");
const Ajv2020 = require("ajv/dist/2020").default;

const { schema, inputs } = JSON.parse(readFileSync(0, "utf8"));
const ajv = new Ajv2020({ strict: true, allErrors: true });
const validate = ajv.compile(schema);
const results = inputs.map((input) => {
  const valid = validate(input);
  return { valid, errors: validate.errors };
});

process.stdout.write(JSON.stringify(results));
