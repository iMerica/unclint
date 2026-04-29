# Unclint

![Unc](unc.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/yourname/unc)](https://goreportcard.com/report/github.com/yourname/unc) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Tired of being a chopped Unc? 

Listen. You're sitting at your desk. You're trying to write a simple sentence, and suddenly you type "driving alignment on high leverage synergies". You're a sick guy. You need help.

Or worse, you're a product manager trying to use "rizz" in a slide deck. It's tough to watch. 

Unclint is a Go CLI that catches your bad copy before anyone else sees it. It flags corporate jargon, stale millennial slang, boomer framing, and fake youth garbage. It keeps you honest.

## The Problem

Your copy stinks. Here is what Unclint catches:

**The Tryhard:**
> "We leverage creator rizz to unlock authentic engagement."
*Error: 'leverage' as a verb reads corporate. This mixes corporate jargon with youth slang. Please stop.*

**The LinkedIn Lunatic:**
> "Thrilled to announce we're circling back to drill down on low-hanging fruit to move the needle."
*Error: Corporate noun pile detected. 'circle back' is filler. Tell us what you actually do.*

**The Stale Millennial:**
> "I did a thing! This new feature is lowkey fire and it's giving me all the feels."
*Error: 'I did a thing' is stale. 'lowkey fire' is tryhard. Grow up.*

## Install

Get it via Homebrew.

```sh
brew tap iMerica/homebrew-tap
brew install unc
```

## Usage

Point it at your docs, your code, or your sad little blog.

```sh
unc lint ./content
unc lint ./copy.md --json
unc lint ./app --include "**/*.{md,tsx}"
```

Explain why a string sucks:

```sh
unc explain "We need to operationalize our synergies."
```

## Configuration

Drop an `.uncrc.yml` in your project. Tune it so it doesn't hurt your feelings too bad.

```yaml
version: 1

# Max score before failing the build. Default 0.
max_score: 15

# Minimum severity to flag (0=info, 1=warn, 2=error)
min_severity: 1

include:
  - "**/*.md"
  - "**/*.tsx"
  - "**/*.txt"

exclude:
  - "node_modules/**"
  - "vendor/**"

rules:
  corporate: true
  stale: true
  tryhard: true
  millennial: true
  boomer: true

overrides:
  - path: "docs/**"
    context: docs
    max_score: 50 # Docs are inherently boring, give them slack

allow:
  terms:
    - "Cloudflare Workers" # Don't flag technical terms

disable:
  - "corporate/verb-leverage" # If you really love the word leverage
```

## Ignore the haters

Sometimes you just gotta say it. Suppress the linter inline.

```md
<!-- unc-disable-next-line corporate/phrase-circle-back -->
I need to circle back on this.
```

Or disable a whole file:

```md
<!-- unc-disable -->
This whole document is a corporate wasteland and I accept that.
```

## Build it yourself

```sh
git clone https://github.com/iMerica/unc.git
cd unc
make build
./bin/unc --help
```

To run the tests with beautiful formatted output:

```sh
make test
```
