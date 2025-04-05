package gguf_info

const (
	GGML_TYPE_F32 Type = iota
	GGML_TYPE_F16
	GGML_TYPE_Q4_0
	GGML_TYPE_Q4_1
	GGML_TYPE_Q4_2 // support has been removed
	GGML_TYPE_Q4_3 // support has been removed
	GGML_TYPE_Q5_0
	GGML_TYPE_Q5_1
	GGML_TYPE_Q8_0
	GGML_TYPE_Q8_1
	GGML_TYPE_Q2_K
	GGML_TYPE_Q3_K
	GGML_TYPE_Q4_K
	GGML_TYPE_Q5_K
	GGML_TYPE_Q6_K
	GGML_TYPE_Q8_K
	GGML_TYPE_IQ2_XXS
	GGML_TYPE_IQ2_XS
	GGML_TYPE_IQ3_XXS
	GGML_TYPE_IQ1_S
	GGML_TYPE_IQ4_NL
	GGML_TYPE_IQ3_S
	GGML_TYPE_IQ2_S
	GGML_TYPE_IQ4_XS
	GGML_TYPE_I8
	GGML_TYPE_I16
	GGML_TYPE_I32
	GGML_TYPE_I64
	GGML_TYPE_F64
	GGML_TYPE_IQ1_M
	GGML_TYPE_COUNT
)

func (t Type) String() string {
	switch t {
	case GGML_TYPE_F32:
		return "GGML_TYPE_F32"
	case GGML_TYPE_F16:
		return "GGML_TYPE_F16"
	case GGML_TYPE_Q4_0:
		return "GGML_TYPE_Q4_0"
	case GGML_TYPE_Q4_1:
		return "GGML_TYPE_Q4_1"
	case GGML_TYPE_Q4_2:
		return "GGML_TYPE_Q4_2"
	case GGML_TYPE_Q4_3:
		return "GGML_TYPE_Q4_3"
	case GGML_TYPE_Q5_0:
		return "GGML_TYPE_Q5_0"
	case GGML_TYPE_Q5_1:
		return "GGML_TYPE_Q5_1"
	case GGML_TYPE_Q8_0:
		return "GGML_TYPE_Q8_0"
	case GGML_TYPE_Q8_1:
		return "GGML_TYPE_Q8_1"
	case GGML_TYPE_Q2_K:
		return "GGML_TYPE_Q2_K"
	case GGML_TYPE_Q3_K:
		return "GGML_TYPE_Q3_K"
	case GGML_TYPE_Q4_K:
		return "GGML_TYPE_Q4_K"
	case GGML_TYPE_Q5_K:
		return "GGML_TYPE_Q5_K"
	case GGML_TYPE_Q6_K:
		return "GGML_TYPE_Q6_K"
	case GGML_TYPE_Q8_K:
		return "GGML_TYPE_Q8_K"
	case GGML_TYPE_IQ2_XXS:
		return "GGML_TYPE_IQ2_XXS"
	case GGML_TYPE_IQ2_XS:
		return "GGML_TYPE_IQ2_XS"
	case GGML_TYPE_IQ3_XXS:
		return "GGML_TYPE_IQ3_XXS"
	case GGML_TYPE_IQ1_S:
		return "GGML_TYPE_IQ1_S"
	case GGML_TYPE_IQ4_NL:
		return "GGML_TYPE_IQ4_NL"
	case GGML_TYPE_IQ3_S:
		return "GGML_TYPE_IQ3_S"
	case GGML_TYPE_IQ2_S:
		return "GGML_TYPE_IQ2_S"
	case GGML_TYPE_IQ4_XS:
		return "GGML_TYPE_IQ4_XS"
	case GGML_TYPE_I8:
		return "GGML_TYPE_I8"
	case GGML_TYPE_I16:
		return "GGML_TYPE_I16"
	case GGML_TYPE_I32:
		return "GGML_TYPE_I32"
	case GGML_TYPE_I64:
		return "GGML_TYPE_I64"
	case GGML_TYPE_F64:
		return "GGML_TYPE_F64"
	case GGML_TYPE_IQ1_M:
		return "GGML_TYPE_IQ1_M"
	case GGML_TYPE_COUNT:
		return "GGML_TYPE_COUNT"
	default:
		return "Unknown"
	}
}

const (
	// The value is a 8-bit unsigned integer.
	GGUF_METADATA_VALUE_TYPE_UINT8 ValueType = iota
	// The value is a 8-bit signed integer.
	GGUF_METADATA_VALUE_TYPE_INT8
	// The value is a 16-bit unsigned little-endian integer.
	GGUF_METADATA_VALUE_TYPE_UINT16
	// The value is a 16-bit signed little-endian integer.
	GGUF_METADATA_VALUE_TYPE_INT16
	// The value is a 32-bit unsigned little-endian integer.
	GGUF_METADATA_VALUE_TYPE_UINT32
	// The value is a 32-bit signed little-endian integer.
	GGUF_METADATA_VALUE_TYPE_INT32
	// The value is a 32-bit IEEE754 floating point number.
	GGUF_METADATA_VALUE_TYPE_FLOAT32
	// The value is a boolean.
	// 1-byte value where 0 is false and 1 is true.
	// Anything else is invalid, and should be treated as either the model being invalid or the reader being buggy.
	GGUF_METADATA_VALUE_TYPE_BOOL
	// The value is a UTF-8 non-null-terminated string, with length prepended.
	GGUF_METADATA_VALUE_TYPE_STRING
	// The value is an array of other values, with the length and type prepended.
	///
	// Arrays can be nested, and the length of the array is the number of elements in the array, not the number of bytes.
	GGUF_METADATA_VALUE_TYPE_ARRAY
	// The value is a 64-bit unsigned little-endian integer.
	GGUF_METADATA_VALUE_TYPE_UINT64
	// The value is a 64-bit signed little-endian integer.
	GGUF_METADATA_VALUE_TYPE_INT64
	// The value is a 64-bit IEEE754 floating point number.
	GGUF_METADATA_VALUE_TYPE_FLOAT64
)
