package testdata

const NumericConstraint = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/NumericConstraint",
    "definitions": {
        "NumericConstraint": {
            "properties": {
                "int_val": {
                    "maximum": 2147483647,
                    "minimum": -2147483648,
                    "type": "integer"
                },
                "long_val": {
                    "type": "string"
                },
                "float_val": {
                    "type": "number"
                },
                "double_val": {
                    "type": "number"
                },
                "int_val_array": {
                    "items": {
                        "maximum": 2147483647,
                        "minimum": -2147483648,
                        "type": "integer"
                    },
                    "type": "array"
                },
                "long_val_array": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array"
                },
                "float_val_array": {
                    "items": {
                        "type": "number"
                    },
                    "type": "array"
                },
                "double_val_array": {
                    "items": {
                        "type": "number"
                    },
                    "type": "array"
                },
                "optional_int_val": {
                    "maximum": 2147483647,
                    "minimum": -2147483648,
                    "type": "integer"
                },
                "uint_val": {
                    "type": "integer"
                },
                "ulong_val": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Numeric Constraint"
        }
    }
}`