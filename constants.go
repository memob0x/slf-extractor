package main

// NOTE: see SLF.md

const STRING_ENCODING string = "ascii"

const INT_BUFFER_STRING_LENGTH int = 256

const INT_BUFFER_NUMBER_LENGTH int = 4

const INT_SLF_BUFFER_ENTRY_LENGTH int = 280

// placed in slf header after lib name (string) and lib path (string)
const INT_SLF_BUFFER_OFFSET_START_ENTRIES_COUNT int = INT_BUFFER_STRING_LENGTH * 2
