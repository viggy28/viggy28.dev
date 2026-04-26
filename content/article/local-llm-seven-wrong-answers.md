---
date: 2026-04-25T00:00:00-04:00
description: "Seven attempts, seven different wrong answers — lessons from setting up a local LLM."
tags: ["llm", "local-llm", "ollama", "ai"]
title: "I Asked My Local LLM to Add 23 Numbers. I Got Seven Different Wrong Answers."
toc: true
---

*Seven attempts, seven different wrong answers — lessons from setting up a local LLM.*

---

It's tax season, which means I've been staring at a notes file full of stock sales — 23 transactions across the year that I needed to total up. The kind of data I'd rather not paste into a chat window I don't control.

I'd been meaning to set up a local LLM anyway, and this seemed like the perfect low-stakes test. M3 Max, 64GB of unified memory, plenty of headroom for a real model. I installed Ollama, pulled Qwen 2.5 Coder, pasted the list, asked the question.

"947 shares sold so far," it told me. Confidently.

The actual answer was 1,884. Over the next five hours and seven attempts, my local LLM gave me 2,333. Then 1,994. Then 2,364. Then 859. Twice it produced no number at all. Eventually, finally, 1,884.

*(For the actual filing I used Python. This is a post about exploring local LLMs, not doing taxes with one. Don't do taxes with one.)*

What I thought would be a five-minute test became the cleanest tour through the modern AI stack I've ever stumbled into. By the end I understood what every layer — model, inference engine, orchestrator, harness — actually *does*, because I'd watched each one fail in turn.

## The data

```
these are the stocks of Chegg i have sold.
250 shares of Chegg at 35
100 shares of Chegg at 42
88 shares of Chegg at 112
50 shares of $CHGG@78
40 shares of $CHGG @42. Cost basis $112. Sold in a loss.
80 shares at 22
145 shares at 8
... (23 transactions total, prices spanning $35 down to $0.80)
how many I have sold so far?
```

Some lines say "Chegg," some say "$CHGG," some are bare numbers. Real answer: **1,884**.

## Attempt 1: Ollama desktop, 7B → **947**

Pasted into the chat. Got "947 shares so far." The model had silently dropped half the input — listed only 12 of 23 transactions. Worse, even those 12 don't sum to 947; they total 1,179. Two compounding failures, one confident answer.

*Lesson: small models will produce list-shaped output that omits items, then total the omitted version without acknowledging the omission.*

## Attempt 2: `ollama run`, 7B → **2,333**

Same model, raw CLI. This time it identified all 23 transactions and wrote out the expression:

```
250 + 100 + 100 + 101 + 80 + 88 + 60 + 80 + 50 + 70 + 29 +
40 + 51 + 80 + 80 + 50 + 60 + 50 + 145 + 70 + 68 + 82 + 100 = 2,333
```

The expression is correct. The answer isn't.

*Lesson: transformers don't compute arithmetic. After the `=`, the model is pattern-matching what number looks plausible, not running addition. Sometimes right, often not, never reliable past a few terms.*

## Attempt 3: Open Interpreter → **never executed**

This is what should fix it. Open Interpreter is a CLI harness — model writes code, harness runs it in a Python sandbox. Pointed it at Ollama:

```bash
interpreter --model ollama/qwen2.5-coder:7b
```

The model produced:

```
{"name": "execute", "arguments":{"language": "python", "code": "..."}}
```

…and Open Interpreter just printed it as text. No "Run this? (y/n)" prompt. The model emitted JSON-shaped *text* that looked like a tool call but wasn't a structured tool call the harness recognized. Same result with the 32B.

*Lesson: tool-calling has two skills — knowing you should call a tool, and emitting the exact tokens that signal one. Smaller open-weights models do the first reliably and fumble the second. Frontier models are heavily post-trained on structured output; smaller models aren't. The handshake fails.*

## Attempt 4: `ollama run`, 32B → **1,994**

Maybe a bigger model handles longer arithmetic. It tried:

```
Chegg subtotal (8 entries):   859     ✓ correct
$CHGG subtotal (15 entries):  1,135   ✗ actual 1,025
Total:                        1,994   ← off by 110
```

Better organization — the 32B split by label and computed subtotals. Got the 8-number sum exactly right. Got the 15-number sum wrong by 110.

*Lesson: bigger model, better organization, same arithmetic floor — it just shifts up. Even Sonnet and Opus get long sums wrong without code execution.*

## Attempt 5: Open WebUI without Code Interpreter → **2,364**

Switched to Open WebUI — localhost web UI with a built-in Python sandbox. Got 2,364. I'd forgotten to enable Code Interpreter for that chat. Without the toggle, Open WebUI is just a chat UI.

*Lesson: a harness that can run code but isn't told to is no harness at all.*

## Attempt 6: Open WebUI + Code Interpreter → **859**, then **1,884** ✓

Code Interpreter on. Model wrote Python. The sandbox executed it: `STDOUT: 859`. Real execution, verified math — but only the 8 rows literally labeled "Chegg." The 15 rows labeled "$CHGG" or unlabeled were excluded.

This is genuinely ambiguous input. The list has *the word* "Chegg" on some lines and *the ticker* "$CHGG" on others. A careful interpreter handling financial data should default to the narrow reading. The model wasn't being dumb; it was being conservative.

I clarified: *"Chegg's ticker is $CHGG. All 23 transactions are Chegg. Sum them all."* Re-asked. Watched the model write Python with all 23 numbers, watched the sandbox execute: **`STDOUT: 1884`**.

*Lesson: when the harness works, your remaining failures shift from computation to interpretation. That's a much better failure mode — interpretation errors are debuggable through clearer prompts.*

## The table

| Setup | Answer | What broke |
|---|---|---|
| Ollama desktop, 7B | 947 | Dropped data + bad arithmetic |
| `ollama run`, 7B | 2,333 | Correct expression, predicted wrong sum |
| Open Interpreter, 7B | (never ran) | Tool-call format mismatch |
| `ollama run`, 32B | 1,994 | Better thinking, still wrong arithmetic |
| Open Interpreter, 32B | (never ran) | Same protocol mismatch |
| Open WebUI, no Code Interpreter | 2,364 | Inline arithmetic |
| Open WebUI + Code Interpreter, narrow prompt | 859 | Right code, narrow scope |
| **Open WebUI + Code Interpreter + clear prompt** | **1,884** ✓ | **Nothing** |

## What this stack actually is

The lesson the failures forced on me: **every working AI product is four layers stacked.**

- **Model** — Qwen, Llama, GPT, Claude. Predicts tokens. Doesn't compute, doesn't act.
- **Inference engine** — llama.cpp, vLLM. Loads weights, runs the math, exposes an API.
- **Orchestrator** — Ollama, OpenAI's API. Manages models, lifecycle, the HTTP surface.
- **Harness** — Open WebUI, Cline, Claude Code itself. Wraps the model with tools, code execution, an agent loop. *This is where reliability lives.*

ChatGPT and Claude.ai give you correct answers to "add these 23 numbers" because they're not just exposing the model — they have built-in code interpreters that fire on math. The model writes a script, the script runs, you see the result. The model never computes.

For local LLMs, this means: **a chat UI that just talks to Ollama is roughly useless for anything involving computation.** Pick a harness with real code execution. Verify it actually fires — watch for `STDOUT` markers or approval prompts. If you don't see code running, you're getting mental math.

Seven wrong answers in one evening was the longest way to learn this: the model is one of four moving parts, and the harness is where reliability lives. Somehow also the only way it stuck.
