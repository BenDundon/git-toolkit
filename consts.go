package main

const (
	feature  string = "feature"
	fix      string = "fix"
	docs     string = "docs"
	style    string = "style"
	refactor string = "refactor"
	test     string = "test"
	chore    string = "chore"
	perf     string = "perf"
	hotfix   string = "hotfix"
)

const (
	featDesc     string = "FEATURE (Introducing new features)"
	fixDesc      string = "FIX (Bug fix)"
	docsDesc     string = "DOCS (Writing docs)"
	styleDesc    string = "STYLE (Improving structure/format of the code)"
	refactorDesc string = "REFACTOR (Refactoring code)"
	testDesc     string = "TEST (When adding missing tests)"
	choreDesc    string = "CHORE (Changing CI/CD)"
	perfDesc     string = "PERF (Improving performance)"
	hotfixDesc   string = "HOTFIX (Bug fix urgently)"
)

const commitMessageCheckPattern = `^(feat|fix|docs|style|refactor|test|chore|perf|hotfix)\((\S.*)\):\s(\S.*)|^Merge.*`

const commitMessageCheckFailedMsg = `
╭──────────────────────────────────────────────────╮
│                                                  │
│    ✗ The commit message is not standardized.     │
│    ✗ It must match the regular expression:       │
│                                                  │
│    ^(feat|fix|docs|style|refactor|test|chore|    │
│     perf|hotfix)\((\S.*)\):\s(\S.*)|^Merge.*     │
│                                                  │
╰──────────────────────────────────────────────────╯`

const editorKey = "//edit"
const newLineKey = `\\n `

const luckyCommitEnv = "MMH_LUCKY_COMMIT"
const strictHostKey = "MMH_STRICT_HOST_KEY"
