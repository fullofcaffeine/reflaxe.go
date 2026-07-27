#!/usr/bin/env node

import fs from "node:fs";
import Ajv from "ajv";

const [schemaPath, reportPath] = process.argv.slice(2);
if (!schemaPath || !reportPath) {
  console.error("usage: validate_surface_contract_schema.mjs <schema> <report|->");
  process.exit(2);
}

const schema = JSON.parse(fs.readFileSync(schemaPath, "utf8"));
const report = JSON.parse(
  reportPath === "-" ? fs.readFileSync(0, "utf8") : fs.readFileSync(reportPath, "utf8"),
);
const ajv = new Ajv({ allErrors: true, strict: true });
const validate = ajv.compile(schema);

if (!validate(report)) {
  console.error(ajv.errorsText(validate.errors, { separator: "\n" }));
  process.exit(1);
}
