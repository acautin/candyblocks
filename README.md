# CandyBlocks

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/37c0e2843cbf44a6bb44e2e343768f4d)](https://www.codacy.com/gh/acautin/candyblocks/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=acautin/candyblocks&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/37c0e2843cbf44a6bb44e2e343768f4d)](https://www.codacy.com/gh/acautin/candyblocks/dashboard?utm_source=github.com&utm_medium=referral&utm_content=acautin/candyblocks&utm_campaign=Badge_Coverage)

Candy themed blocks game developed using SDL and Go

## Development

To compile `candyblocks` you need to have [`Air`](https://github.com/cosmtrek/air) and the `SDL` library installed in your system.

*   `git clone https://github.com/acautin/candyblocks`
*   `cd candyblocks`
*   `air`

### Generated files

Don't forget to run `go generate ./...` if you extend the candy type definitions.

### Tests

Unit tests can be run via `gotest ./...`.
