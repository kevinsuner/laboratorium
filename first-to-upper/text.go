/*
* SPDX-License-Identifier: GPL-3.0-only
* Copyright (C) 2024 Kevin Suñer <ksuner@pm.me>
 */

package text

import "unicode"

// Uppercases the first letter in every word.
func FirstLetterToUpper(text string) string {
    out := make([]rune, 0, len(text))
    for i := 0; i < len(text); i++ {
        if i == 0 {
            out = append(out, unicode.ToUpper(rune(text[i])))
            continue
        }

        if 'a' <= text[i] && text[i] <= 'z' || text[i] == ' ' {
            if text[i] == ' ' && i < len(text)-1 { // avoid index out of bounds
                out = append(out, ' ')
                out = append(out, unicode.ToUpper(rune(text[i+1])))
                i++ // skips the next one to avoid duplicates
                
                continue
            }

            out = append(out, rune(text[i]))
        }
    }

    return string(out)
}
