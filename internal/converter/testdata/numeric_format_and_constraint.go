package testdata

const NumericFormatAndConstraint = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/NumericConstraint",
    "definitions": {
        "NumericConstraint": {
            "properties": {
                "int_val": {
                    "maximum": 2147483647,
                    "minimum": -2147483648,
                    "type": "integer",
                    "format": "int32"
                },
                "long_val": {
                    "maximum": 9223372036854775807,
                    "minimum": -9223372036854775808,
                    "type": "string",
                    "format": "int64"
                },
                "float_val": {
                    "type": "number",
                    "format": "float"
                },
                "double_val": {
                    "type": "number",
                    "format": "double"
                },
                "int_val_array": {
                    "items": {
                        "maximum": 2147483647,
                        "minimum": -2147483648,
                        "type": "integer",
                        "format": "int32"
                    },
                    "type": "array"
                },
                "long_val_array": {
                    "items": {
                        "maximum": 9223372036854775807,
                        "minimum": -9223372036854775808,
                        "type": "string",
                        "format": "int64"
                    },
                    "type": "array"
                },
                "float_val_array": {
                    "items": {
                        "type": "number",
                        "format": "float"
                    },
                    "type": "array"
                },
                "double_val_array": {
                    "items": {
                        "type": "number",
                        "format": "double"
                    },
                    "type": "array"
                },
                "optional_int_val": {
                    "maximum": 2147483647,
                    "minimum": -2147483648,
                    "type": "integer",
                    "format": "int32"
                },
                "uint_val": {
                    "type": "integer",
                    "format": "int32"
                },
                "ulong_val": {
                    "type": "string",
                    "format": "int64"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Numeric Constraint"
        }
    }
}`
