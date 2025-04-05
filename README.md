# gguf-info

## What

gguf-info reads a local `gguf` format LLM file and prints out details.

e.g.

```
$ gguf-info nomic-embed-text-v1.5.f32.gguf

Header:
        Magic:            GGUF
        Version:          3
        Tensor Count:     112
        Metadata Entries: 22
        Metadata:
                * general.architecture=nomic-bert
                * general.name=nomic-embed-text-v1.5
                * nomic-bert.block_count=12
                * nomic-bert.context_length=2048
                * nomic-bert.embedding_length=768
                * nomic-bert.feed_forward_length=3072
                * nomic-bert.attention.head_count=12
                * nomic-bert.attention.layer_norm_epsilon=0.000000
                * general.file_type=0
                * nomic-bert.attention.causal=false
                * nomic-bert.pooling_type=1
                * nomic-bert.rope.freq_base=1000.000000
                * tokenizer.ggml.token_type_count=2
                * tokenizer.ggml.bos_token_id=101
                * tokenizer.ggml.eos_token_id=102
                * tokenizer.ggml.model=bert
                * tokenizer.ggml.tokens=...truncated...
                * tokenizer.ggml.scores=...truncated...
                * tokenizer.ggml.token_type=...truncated...
                * tokenizer.ggml.unknown_token_id=100
                * tokenizer.ggml.seperator_token_id=102
                * tokenizer.ggml.padding_token_id=0
        Tensors:
                * name:       token_embd_norm.bias
                  dimensions: 1, [768]
                  type:       GGML_TYPE_F32
                  offset:     0
                * name:       token_embd_norm.weight
                  dimensions: 1, [768]
                  type:       GGML_TYPE_F32
                  offset:     3072
                * name:       token_types.weight
                  ...
```

To view a specific metadata key (in full), use `--metadata-key` or `-m`, e.g. 
```
$ gguf-info nomic-embed-text-v1.5.f32.gguf -m general.architecture
nomic-bert
```

Metadata with large array values is returned as JSON, e.g.

```
$ gguf-info models/nomic-embed-text-v1.5.f32.gguf -m tokenizer.ggml.tokens | json_pp | less
[
   "[PAD]",
   "[unused0]",
   "[unused1]",
   ...
```
